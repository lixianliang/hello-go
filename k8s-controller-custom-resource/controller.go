package main

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"

	samplecrdv1 "github.com/lixianliang/hello-go/k8s-controller-custom-resource/pkg/apis/samplecrd/v1"
	clientset "github.com/lixianliang/hello-go/k8s-controller-custom-resource/pkg/client/clientset/versioned"
	networkscheme "github.com/lixianliang/hello-go/k8s-controller-custom-resource/pkg/client/clientset/versioned/scheme"
	informers "github.com/lixianliang/hello-go/k8s-controller-custom-resource/pkg/client/informers/externalversions/samplecrd/v1"
	listers "github.com/lixianliang/hello-go/k8s-controller-custom-resource/pkg/client/listers/samplecrd/v1"
)

const controllerAgentName = "network-controller"

const (
	SuccessSynced = "Synced"

	MessageResourceSynced = "Network synced successfully"
)

type Controller struct {
	kubeclientset    kubernetes.Interface
	networkclientset clientset.Interface

	networksLister listers.NetworkLister
	networksSynced cache.InformerSynced

	//workqueue workqueue.RatelimitingInterface
	workqueue workqueue.RateLimitingInterface
	recorder  record.EventRecorder
}

func NewController(kubeclientset kubernetes.Interface, networkclientset clientset.Interface, networkInformer informers.NetworkInformer) *Controller {
	utilruntime.Must(networkscheme.AddToScheme(scheme.Scheme))
	glog.V(4).Info("Create event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &Controller{
		kubeclientset:    kubeclientset,
		networkclientset: networkclientset,
		networksLister:   networkInformer.Lister(),
		networksSynced:   networkInformer.Informer().HasSynced,
		workqueue:        workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Networks"),
		recorder:         recorder,
	}

	glog.Info("Setting up event handlers")
	networkInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueNetwork,
		UpdateFunc: func(old, new interface{}) {
			oldNetwork := old.(*samplecrdv1.Network)
			newNetwork := new.(*samplecrdv1.Network)
			if oldNetwork.ResourceVersion == newNetwork.ResourceVersion {
				return
			}
			controller.enqueueNetwork(new)
		},
		DeleteFunc: controller.enqueueNetworkForDelete,
	})

	return controller
}

func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
	defer runtime.HandleCrash()
	defer c.workqueue.ShutDown()

	glog.Info("Starting Network control loop")

	glog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.networksSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	glog.Infof("Starting workers")
	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	glog.Info("Started workers")
	<-stopCh
	glog.Info("Shutting down workers")

	return nil
}

func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
	}
}

func (c *Controller) processNextWorkItem() bool {
	obj, shutdown := c.workqueue.Get()

	if shutdown {
		return false
	}

	err := func(obj interface{}) error {
		defer c.workqueue.Done(obj)
		var key string
		var ok bool

		if key, ok = obj.(string); !ok {
			c.workqueue.Forget(obj)
			runtime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}

		if err := c.syncHandler(key); err != nil {
			return fmt.Errorf("error syncing %s:%s", key, err.Error())
		}

		c.workqueue.Forget(obj)
		glog.Infof("Successfully synced %s", key)
		return nil
	}(obj)

	if err != nil {
		runtime.HandleError(err)
		return true
	}

	return true
}

func (c *Controller) syncHandler(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		runtime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	network, err := c.networksLister.Networks(namespace).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			glog.Warningf("Network: %s/%s does not exist in local cache, will delete it from Neutron ...",
				namespace, name)
			glog.Infof("[Neutron] Deleting network: %s/%s ...", namespace, name)
			// FIX ME: call Neutron API to delete this network by name.
			// neutron.Delete(namespace, name)
			return nil
		}

		runtime.HandleError(fmt.Errorf("failed to list network by: %s/%s", namespace, name))
		return err
	}

	glog.Infof("[Neutron] Try to process network: %#v ...", network)

	// FIX ME: Do diff().
	// actualNetwork, exists := neutron.Get(namespace, name)
	// if !exists {
	// 	neutron.Create(namespace, name)
	// } else if !reflect.DeepEqual(actualNetwork, network) {
	// 	neutron.Update(namespace, name)
	// }

	c.recorder.Event(network, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)
	return nil
}

func (c *Controller) enqueueNetwork(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		runtime.HandleError(err)
		return
	}

	c.workqueue.AddRateLimited(key)
}

func (c *Controller) enqueueNetworkForDelete(obj interface{}) {
	var key string
	var err error
	key, err = cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
		return
	}

	c.workqueue.AddRateLimited(key)
}
