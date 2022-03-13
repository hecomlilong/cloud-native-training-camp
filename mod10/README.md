# 编译httpserver
```
MOD=mod10 make build
```
# 打包镜像
```
MOD=mod10 make docker-build
```

# 推送镜像
```
MOD=mod10 DOCKER_ACCOUNT=swiftabc make docker-push
```

# 创建pod
```
sudo kubectl apply -f mod10/httpserver.yaml
```

#请求接口
```
pod_ip=$(sudo kubectl get po -o=jsonpath='{.items[?(@.metadata.labels.app=="httpserver")].status.podIP}')

curl $pod_ip:1880/hello
curl $pod_ip:1880/hello
curl $pod_ip:1880/hello
curl $pod_ip:1880/hello
curl $pod_ip:1880/hello
```
