package app

import etcdClient "go.etcd.io/etcd/client/v3"

func GetRegistryEtcdClient() (*etcdClient.Client, error) {
	client, err := etcdClient.New(etcdClient.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	return client, err
}

func GetRegistryEtcdClientWithDefalut() *etcdClient.Client {
	client, err := etcdClient.New(etcdClient.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	return client
}
