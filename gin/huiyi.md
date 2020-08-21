3+2 * 12

电信合作技术成本:
提供目前已有的所有玩法
人力成本：
60w: 所有玩法经过一年上下的开发时间
成本按GPU机器按1台每月5000算（具体看算法部署量）

https://algo.rthdo.com/algo/jimu/v1/make_video

test@max-master:~/kube-prometheus/manifests$ kubectl get --raw /apis/custom.metrics.k8s.io/v1beta1/namespaces/default/services/rabbitmq-exporter/rabbitmq_queue_messages_ready?metricLabelSelector=queue%3Dtest-1 | jq .
{
  "kind": "MetricValueList",
  "apiVersion": "custom.metrics.k8s.io/v1beta1",
  "metadata": {
    "selfLink": "/apis/custom.metrics.k8s.io/v1beta1/namespaces/default/services/rabbitmq-exporter/rabbitmq_queue_messages_ready"
  },
  "items": [
    {
      "describedObject": {
        "kind": "Service",
        "namespace": "default",
        "name": "rabbitmq-exporter",
        "apiVersion": "/v1"
      },
      "metricName": "rabbitmq_queue_messages_ready",
      "timestamp": "2020-02-21T09:19:41Z",
      "value": "0",
      "selector": null
    }
  ]
}

1028  kubectl get --raw="/apis/custom.metrics.k8s.io/v1beta1/namespaces/default/pods/*/nsq_channel_depth?metricLabelSelector=topic%3Dvideo_post_vid" | jq
1029  kubectl get --raw="/apis/custom.metrics.k8s.io/v1beta1/namespaces/default/services/nsq-exporter-svc/nsq_channel_depth?metricLabelSelector=topic%3Dvideo_post_vid" | jq


apiVersion: autoscaling/v2beta2
kind: HorizontalPodAutoscaler
metadata:
  name: video-post-h5-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: video-post-h5-deployment # 需要扩缩容的deployment名字
  minReplicas: 1
  maxReplicas: 4
  metrics:
  - type: Pods
    pods:
      metricName: nsq_channel_depth
      targetAverageValue: 10

 [ValidationError(HorizontalPodAutoscaler.spec.metrics[0].pods): unknown field "metricName" in io.k8s.api.autoscaling.v2beta2.PodsMetricSource, ValidationError(HorizontalPodAutoscaler.spec.metrics[0].pods): unknown field "targetAverageValue" in io.k8s.api.autoscaling.v2beta2.PodsMetricSource, ValidationError(HorizontalPodAutoscaler.spec.metrics[0].pods): missing required field "metric" in io.k8s.api.autoscaling.v2beta2.PodsMetricSource, ValidationError(HorizontalPodAutoscaler.spec.metrics[0].pods): missing required field "target" in io.k8s.api.autoscaling.v2beta2.PodsMetricSource]; if you choose to ignore these errors, turn validation off with --validate=false

## Xxx
1. 响应慢(没有及时反馈)
2. 如果jira多久可以多久恢复 === 多久反馈
3. 回答问题原因，说下一步解决问题的人(简化问题回复)
