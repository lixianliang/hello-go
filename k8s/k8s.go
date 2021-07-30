package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type Config struct {
	Host  string
	Token string
	Port  int
}

func main() {
	clientSet, _ := NewKubernetesClient(&Config{
		Host:  "172.16.6.46",
		Port:  6443,
		Token: "MIICXAIBAAKBgQC3dH+crByJTz3LY/SQfSU7TuO3VVeWFGnpUoz1fFnEnXlSjtDTfP13k6MwrktiNESDBXoRWhUrtPD+y/jzJsZUc4K2ZxxGkPZ21rNuwfFWCeTyqM9pU+HdDB20T32533ncBcmXe36SFV7AOv8sI8zI48I08+JxfkL5e/4jNXjs8wIDAQABAoGAfYv6cQvQE5/pGDH1gpRCUI4yhJqg8BJUUNqvoKhS/p0OFBOska8t/xFIUt5UtIY0hL3QxeMyLdEMRDLu0egtPSAsZ7OsQ5ER8nVsy0TCLd41IFgLC2OXGouM8siiZRV2PXeiJPxJQSYnvK7IDkjLnI5EMWR9UM/DSTEELdESQMECQQDEea+igrJ9Z5iICOZnqdrrIqDiX4lV2+BH4Im7kvm18pePIHL1aAiR7nuPej2+Uh8NyxeWlOWXqZLO5Rjdp5KxAkEA7wj7QqdZqldVQ4o7uII2OTFESU/D9nmGtjNKKvP/wrilosMk928382CLOnD1eCfysTx9m4D0uefbTG6hrPH64wJAGGQi8cHX9smNnhW8xNHJY7eA0ZmaqxYI2eN+NdMhPP1I43Pb0auApN0+aal7UM1RHZ1A6GjDt/hNSXIXjCzpgQJBAJEO10PUwPJBe2m4SpOm2XcNsc33jQlXKGwLZhf46J3nZgUG/bj2knKshPFbOWvIelwaRHOI53ql/Iw+mviUBFcCQHT93NN6PSnMwkBiQ13a5w+jvnIgDvFiNPJZcTT4aAjg7pnOqNyByKEiaBnykNBom4MMPZU2aRk1BgK1eTO0XZs=",
	})

	namespaceList, err := clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		//panic(log.Log.Info(err.Error()))
		panic(err.Error())
		return
	}
	//namespaces := GetAllNamespace(clientSet)
	var namespaces []string
	fmt.Println("******************")

	for _, nsList := range namespaceList.Items {
		fmt.Println(nsList.Name)
		namespaces = append(namespaces, nsList.Name)
	}

	fmt.Println("******************")

}

func NewKubernetesClient(c *Config) (*kubernetes.Clientset, error) {
	kubeConf := &rest.Config{
		Host:        fmt.Sprintf("%s:%d", c.Host, c.Port),
		BearerToken: c.Token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	return kubernetes.NewForConfig(kubeConf)
}

// GetAllNamespace get all namespace in cluster.
func GetAllNamespace(clientset *kubernetes.Clientset) []string {
	var namespaces []string
	namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Log.Info("***************err*****************")
		log.Log.Info(err.Error())
	} else {
		//fmt.Printf(namespaces[0])
		for _, nsList := range namespaceList.Items {
			namespaces = append(namespaces, nsList.Name)
		}
	}

	return namespaces
}
