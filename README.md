# cloud-native-training-camp

# 编译
执行 make mod=[number] 编译指定模块的代码，例如编译模块2的代码执行以下命令即可
make mod=2
# 执行
执行 make run mod=[number] 执行指定模块编译好的二进制文件，例如执行模块2编译好的二进制文件，执行以下命令即可
make run mod=2
# 测试
执行 make test 执行单元测试 指定环境变量DEBUG=1 可打印调试信息, 例如：
DEBUG=1 make test
