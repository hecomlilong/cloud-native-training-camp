# 安装repo
```
helm repo add grafana https://grafana.github.io/helm-charts
```
# 安装grafana和prometheus
```
helm upgrade --install loki grafana/loki-stack --set grafana.enabled=true,prometheus.enabled=true,prometheus.alertmanager.persistentVolume.enabled=false,prometheus.server.persistentVolume.enabled=false
```

# 获取密码
```
kubectl get secret loki-grafana -oyaml
在admin-password字段中获取xxx
```
# base64解密
```
echo 'xxx' | base64 -d
```
