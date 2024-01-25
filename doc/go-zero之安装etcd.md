本文只涉及 Linux 上的安装。

### 二进制安装

* **下载二进制安装包**

```sh
#ETCD_VER=v3.4.28
ETCD_VER=v3.5.10
DOWNLOAD_URL=https://github.com/etcd-io/etcd/releases/download
INSTALL_DIR=/tmp

rm -f ${INSTALL_DIR}/etcd-${ETCD_VER}-linux-amd64.tar.gz
rm -rf ${INSTALL_DIR}/etcd-download-test && mkdir -p /tmp/etcd-download-test

curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o ${INSTALL_DIR}/etcd-${ETCD_VER}-linux-amd64.tar.gz
```

下载地址示例：

```
https://github.com/etcd-io/etcd/releases/download/v3.5.10/etcd-v3.5.10-linux-amd64.tar.gz
```

* **解压二进制安装包**

```sh
tar xzvf ${INSTALL_DIR}/etcd-${ETCD_VER}-linux-amd64.tar.gz -C ${INSTALL_DIR}/etcd-download-test --strip-components=1
```

* **删除二进制安装包**

 ```sh
rm -f ${INSTALL_DIR}/etcd-${ETCD_VER}-linux-amd64.tar.gz
```

* **版本检查**

```sh
${INSTALL_DIR}/etcd-download-test/etcd --version
${INSTALL_DIR}/etcd-download-test/etcdctl version
```

* **启动 etcd**

```sh
${INSTALL_DIR}/etcd-download-test/etcd
```

* **往 etcd 写读数据**

```
${INSTALL_DIR}/etcd-download-test/etcdctl --endpoints=localhost:2379 put foo bar
${INSTALL_DIR}/etcd-download-test/etcdctl --endpoints=localhost:2379 get foo
```

### Docker 安装

```sh
INSTALL_DIR=/tmp
rm -rf ${INSTALL_DIR}/etcd-data.tmp && mkdir -p ${INSTALL_DIR}/etcd-data.tmp && \
  docker rmi gcr.io/etcd-development/etcd:v3.4.28 || true && \
  docker run \
  -p 2379:2379 \
  -p 2380:2380 \
  --mount type=bind,source=${INSTALL_DIR}/etcd-data.tmp,destination=/etcd-data \
  --name etcd-gcr-v3.4.28 \
  gcr.io/etcd-development/etcd:v3.4.28 \
  /usr/local/bin/etcd \
  --name s1 \
  --data-dir /etcd-data \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://0.0.0.0:2379 \
  --listen-peer-urls http://0.0.0.0:2380 \
  --initial-advertise-peer-urls http://0.0.0.0:2380 \
  --initial-cluster s1=http://0.0.0.0:2380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --log-level info \
  --logger zap \
  --log-outputs stderr

docker exec etcd-gcr-v3.4.28  /usr/local/bin/etcd --version
docker exec etcd-gcr-v3.4.28  /usr/local/bin/etcdctl version
docker exec etcd-gcr-v3.4.28  /usr/local/bin/etcdctl endpoint health
docker exec etcd-gcr-v3.4.28  /usr/local/bin/etcdctl put foo bar
docker exec etcd-gcr-v3.4.28  /usr/local/bin/etcdctl get foo
```

etcd 主要使用 Google 容器注册表（gcr.io）下的 gcr.io/etcd-development/etcd 仓库来存储其容器镜像。作为次要选项，它还使用 Quay.io 下的 quay.io/coreos/etcd 仓库。这两个注册表都提供 etcd 容器镜像，可用于在类似 Kubernetes 的容器化环境中部署 etcd。

grc: Google Container Registry

### 安装参考

[https://github.com/etcd-io/etcd/releases/](https://github.com/etcd-io/etcd/releases/)

### 使用证书

假设 CACert 证书文件为 etcd-6ecm89rt-CAcert，Cert 证书文件为 etcd-6ecm89rt-Cert，私钥文件为 etcd-6ecm89rt-Key，etcd 的访问地址为 https://192.168.10.17:2023，取 key 为 foo 的值：

```sh
etcdctl --endpoints=https://192.168.10.17:2023 \
        --cacert=etcd-6ecm89rt-CAcert \
        --cert=etcd-6ecm89rt-Cert \
        --key=etcd-6ecm89rt-Key \
        get foo
```
