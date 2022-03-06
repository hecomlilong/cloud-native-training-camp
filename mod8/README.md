# 服务镜像
```
docker.io/swiftabc/cncamp-lilong-2022-01-13-95b5088
```
# 加载configmap
```
sudo kubectl create cm web-config --from-file=config/web-config.ini
```
# 执行yml
```
sudo kubectl apply -f mod8/deploy.yml
```


# 2022-03-06 作业
# 生成cert
```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=highking.tech/O=highking" -addext "subjectAltName = DNS:highking.tech"
```
# 创建secret
```
sudo kubectl create secret tls cncamp-tls --cert=./tls.crt --key=./tls.key
```
# 创建service
```
sudo kubectl apply -f service.yml
```
# 创建ingress
```
sudo kubectl apply -f ingress.yml
```
