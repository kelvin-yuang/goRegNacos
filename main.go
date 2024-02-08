package main

import (
	"fmt"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func main() {
	// 创建
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "nacos.yuang-group.cn",
			Port:   8848,
		},
	}

	//创建 ClientConfig
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId("f82b2dbf-af7b-4ffb-af84-b117c2aca9ad"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// 创建 Nacos 客户端
	client, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		fmt.Println("创建 Nacos 客户端失败:", err)
		return
	}
	fmt.Println("client:", client)

	// 注销服务
	_, err = client.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          "127.0.0.2",
		Port:        8900,
		ServiceName: "goDemo01",
		GroupName:   "DEFAULT_GROUP",
		Cluster:     "clusterA",
	})
	if err != nil {
		fmt.Println("注销服务失败:", err)
		return
	}
	fmt.Println("注销服务成功")

	// 注册服务
	_, err = client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "127.0.0.2",
		Port:        8900,
		ServiceName: "goDemo01",
		GroupName:   "DEFAULT_GROUP",
		ClusterName: "clusterA",
		Weight:      10,
		Enable:      false,
		Healthy:     true,
		Metadata:    map[string]string{"idc": "shanghai"},
	})
	if err != nil {
		fmt.Println("注册服务失败:", err)
		return
	}
	fmt.Println("注册服务成功")

	// configClient, err := clients.CreateConfigClient(map[string]interface{}{
	// 	"serverConfigs": serverConfigs,
	// 	"clientConfig":  clientConfig,
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// content, err := configClient.GetConfig(vo.ConfigParam{
	// 	DataId: "userServer",
	// 	Group:  "DEFAULT_GROUP",
	// })

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(content) //字符串 - yaml
	// err = configClient.ListenConfig(vo.ConfigParam{
	// 	DataId: "userServer",
	// 	Group:  "DEFAULT_GROUP",
	// 	OnChange: func(namespace, group, dataId, data string) {
	// 		fmt.Println("配置文件发生了变化...")
	// 		fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
	// 	},
	// })

	time.Sleep(300 * time.Second)

	fmt.Print("hi")

}
