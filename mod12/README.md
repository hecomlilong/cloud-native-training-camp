# 2022-03-27 作业
# 编译服务，打包镜像
```
MOD=mod12 SURFIX=service-1 MAIN_FILE=service1.go make build
MOD=mod12 SURFIX=service-1 MAIN_FILE=service1.go make docker-build
MOD=mod12 SURFIX=service-1 MAIN_FILE=service1.go DOCKER_ACCOUNT=swiftabc make docker-push
MOD=mod12 SURFIX=service-2 MAIN_FILE=service2.go make build
MOD=mod12 SURFIX=service-2 MAIN_FILE=service2.go make docker-build
MOD=mod12 SURFIX=service-2 MAIN_FILE=service2.go DOCKER_ACCOUNT=swiftabc make docker-push
MOD=mod12 SURFIX=service-3 MAIN_FILE=service3.go make build
MOD=mod12 SURFIX=service-3 MAIN_FILE=service3.go make docker-build
MOD=mod12 SURFIX=service-3 MAIN_FILE=service3.go DOCKER_ACCOUNT=swiftabc make docker-push
```
# 进入mod12
```
cd mod12
```
# 生成cert
```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=highking.tech/O=highking" -addext "subjectAltName = DNS:highking.tech"
```
# 创建secret
```
sudo kubectl create secret tls cncamp-tls --cert=./tls.crt --key=./tls.key
```
# 安装istio
```
curl -L https://istio.io/downloadIstio | sh -
sudo cp istio-1.13.2/bin/istioctl /usr/local/bin
sudo kubectl create namespace istio-system
istioctl install --set profile=demo -y
```
# 创建命名空间
```
sudo kubectl create ns istio-gw
sudo kubectl label ns istio-gw istio-injection=enabled
```
# 创建微服务的deploy 和 svc
```
sudo kubectl apply -f service/service1.yml -n istio-gw
sudo kubectl apply -f service/service2.yml -n istio-gw
sudo kubectl apply -f service/service3.yml -n istio-gw
```
# 配置七层路由规则 && tls 保证安全
```
sudo kubectl apply -f istio-specs.yml -n istio-gw
```

# 验证Open tracing
# 安装jaeger
```
sudo kubectl apply -f jaeger.yml
sudo kubectl edit configmap istio -n istio-system
```
# 设置采样频率tracing.sampling=100

# 查看ingress ip
```
INGRESS_IP=$(sudo kubectl get svc istio-ingressgateway -n istio-system -o=jsonpath='{.spec.clusterIP}')
```
# 客户端访问service1
```
for((i=1;i<=100;i++));
do
    curl https://$INGRESS_IP/service1/hello -k;
done
```
# 查看jaeger面板
```
istioctl dashboard jaeger
```
![jaeger 面板](https://github.com/hecomlilong/cloud-native-training-camp/mod12/jaeger.png "jaeger dashboard")
