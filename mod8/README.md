# 服务镜像
docker.io/swiftabc/cncamp-lilong-2022-01-13-95b5088
# 加载configmap
sudo kubectl create cm web-config --from-file=config/web-config.ini
# 执行yml
sudo kubectl apply -f mod8/deploy.yml
