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
