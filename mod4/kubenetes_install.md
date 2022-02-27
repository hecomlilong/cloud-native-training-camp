# 用 Kubeadm 安装 Kubernetes 集群

## 操作系统：Centos 7

## 操作步骤：

### 1、安装Kubeadm

#### 1.1 允许 iptables 检查桥接流量

```
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF

cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
sudo sysctl --system
```

#### 1.2 检查所需端口

[必要的端口](https://kubernetes.io/zh/docs/reference/ports-and-protocols/)

#### 1.3 安装容器运行时

容器运行环境是负责运行容器的软件。

Kubernetes 支持多个容器运行环境: [Docker](https://kubernetes.io/zh/docs/reference/kubectl/docker-cli-to-kubectl/)、 [containerd](https://containerd.io/docs/)、[CRI-O](https://cri-o.io/#what-is-cri-o) 以及任何实现 [Kubernetes CRI (容器运行环境接口)](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/container-runtime-interface.md)。

docker在centos中的安装命令如下，内容来自：https://docs.docker.com/engine/install/

```
sudo yum install -y yum-utils
sudo yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
sudo yum install docker-ce docker-ce-cli containerd.io
sudo systemctl start docker
sudo systemctl enable docker
```

#### 1.4 配置kubenetes的repo

```
cat <<EOF | sudo tee /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-\$basearch
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
exclude=kubelet kubeadm kubectl
EOF
```

如果不能翻墙请换成阿里云的源，即:

```
cat <<EOF | sudo tee /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=http://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=http://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg http://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
exclude=kubelet kubeadm kubectl
EOF
```

#### 1.5 将 SELinux 设置为 permissive 模式（相当于将其禁用）

```
sudo setenforce 0
sudo sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config
```

#### 1.6 安装kubelet，kubeadm，kubectl

```
sudo yum install -y kubelet kubeadm kubectl --disableexcludes=kubernetes
```

#### 1.7 设置kubelet开机启动

```
sudo systemctl enable --now kubelet
```





### 2、安装单个控制平面的 Kubernetes 集群

#### 2.1  配置 cgroup 驱动程序，docker需要和kubelet保持一致，本文修改docker驱动，和kubelet保持一致，修改后的配置文件/etc/docker/daemon.json 

```
{
  "exec-opts": [
    "native.cgroupdriver=systemd"
  ],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2",
  "bip":"172.12.0.1/24"
}
```

重启docker

```
systemctl restart docker
```

#### 2.2  创建kubenetes集群

```
sudo kubeadm init --pod-network-cidr=172.26.0.0/16
```

不能翻墙执行以下命令

```
sudo kubeadm init  --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --pod-network-cidr=172.26.0.0/16
```

配置不同的--pod-network-cidr参数防止跟其他网络冲突

#### 2.3 非 root 用户可以运行 kubectl，请运行以下命令

```
 mkdir -p $HOME/.kube
 sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
 sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

root用户可以执行以下命令

```
export KUBECONFIG=/etc/kubernetes/admin.conf
```

#### 2.4 安装pod网络附加组件

```
wget https://docs.projectcalico.org/manifests/tigera-operator.yaml
wget https://docs.projectcalico.org/manifests/custom-resources.yaml
```

将custom-resources.yaml文件中的cidr改成172.26.0.0/16即和kubeadm init时的参数一致

然后安装 Pod 网络附加组件

```
 sudo kubectl apply -f tigera-operator.yaml
 sudo kubectl apply -f custom-resources.yaml
```

查看pod变化

```
 sudo watch kubectl get pods -n calico-system
```

### 3、排错经验

#### 3.1、 缺少conntrack

错误内容

```
Error: Package: kubelet-1.22.3-0.x86_64 (kubernetes)
           Requires: conntrack
Error: Package: kubelet-1.22.3-0.x86_64 (kubernetes)
           Requires: socat
```

解决方案

```
sudo yum install -y yum-utils
sudo yum-config-manager \
    --add-repo \
    http://mirrors.aliyun.com/repo/epel-7.repo
sudo yum-config-manager \
    --add-repo \
    http://mirrors.aliyun.com/repo/Centos-7.repo
sudo yum install epel-release
sudo yum install conntrack-tools
```


