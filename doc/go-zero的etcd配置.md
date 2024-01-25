实现代码在 core/discov/config.go 文件中：

```go
type EtcdConf struct {
	Hosts              []string
	Key                string
	ID                 int64  `json:",optional"`
	User               string `json:",optional"`
	Pass               string `json:",optional"`
	CertFile           string `json:",optional"`
	CertKeyFile        string `json:",optional=CertFile"`
	CACertFile         string `json:",optional=CertFile"`
	InsecureSkipVerify bool   `json:",optional"`
}
```

配置示例：

```
# cat etc/add.yaml 
Name: add.rpc
ListenOn: 0.0.0.0:2023
Etcd:
  Hosts:
  - 192.168.10.17:2379
  Key: add.rpc
  CertFile: /tmp/etcd-6ecm89rt-Cert
  CertKeyFile: /tmp/etcd-6ecm89rt-Key
  CACertFile: /tmp/etcd-6ecm89rt-CAcert
```
