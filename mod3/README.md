# mod3

# 1、构建本地镜像
执行 `make docker-build`
# 2、编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化
执行 `make docker-build`
# 3、将镜像推送至 docker 官方镜像仓库
执行 `DOCKER_ACCOUNT=swiftabc make docker-push`
# 4、通过 docker 命令本地启动 httpserver
执行 `make docker-run`
# 5、通过 nsenter 进入容器查看 IP 配置
需要按以下步骤执行命令
1. 执行命令 `docker ps | grep cncamp` 获取容器ID
2. 执行命令 `PID=$(docker inspect --format "{{ .State.Pid}}" <容器ID>)` 获取容器pid
3. 执行命令 `nsenter -t $PID -n ip a` 查看容器IP配置
