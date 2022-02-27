## 按照课上讲解的方法在本地构建一个单节点的基于 HTTPS 的 etcd 集群
```
ETCD_VER=v3.5.2
DOWNLOAD_URL=https://github.com/etcd-io/etcd/releases/download
rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
rm -rf /tmp/etcd-download-test && mkdir -p /tmp/etcd-download-test
curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/etcd-download-test --strip-components=1
sudo cp /tmp/etcd-download-test/etcd* /usr/bin/
etcd --listen-client-urls 'http://localhost:12379' \
 --advertise-client-urls 'http://localhost:12379' \
 --listen-peer-urls 'http://localhost:12380' \
 --initial-advertise-peer-urls 'http://localhost:12380' \
 --initial-cluster 'default=http://localhost:12380'
etcdctl member list --write-out=table --endpoints=localhost:12379
```
## 写一条数据
```
etcdctl --endpoints=localhost:12379 put /key val1
```
## 查看数据细节
```
etcdctl --endpoints=localhost:12379 get /key -wjson
```
## 删除数据
```
etcdctl --endpoints=localhost:12379 del /key
```
