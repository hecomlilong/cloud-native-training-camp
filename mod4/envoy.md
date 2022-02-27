## 启动一个 Envoy Deployment。
```
sudo kubectl create -f envoy-deploy.yaml
```
## 要求 Envoy 的启动配置从外部的配置文件 Mount 进 Pod。
```
sudo kubectl create configmap envoy-config --from-file=envoy.yaml
sudo kubectl create -f envoy-deploy.yaml
```
## 进入 Pod 查看 Envoy 进程和配置。
```
envoy_pod=$(sudo kubectl get pods -o=jsonpath='{.items[?(@.metadata.labels.run=="envoy")].metadata.name}')
sudo kubectl exec -it $envoy_pod -- /bin/bash
ps aux | grep envoy
cat /etc/envoy/envoy.yaml
```
## 更改配置的监听端口并测试访问入口的变化。
## 通过非级联删除的方法逐个删除对象。
```
sudo kubectl delete deploy envoy --cascade=orphan
envoy_rs=$(sudo kubectl get rs -o=jsonpath='{.items[?(@.metadata.labels.run=="envoy")].metadata.name}')
sudo kubectl delete rs $envoy_rs --cascade=orphan
envoy_pod=$(sudo kubectl get pods -o=jsonpath='{.items[?(@.metadata.labels.run=="envoy")].metadata.name}')
sudo kubectl delete po $envoy_pod
```
