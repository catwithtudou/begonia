package client

import (
	"github.com/MashiroC/begonia/dispatch"
	"github.com/MashiroC/begonia/logic"
	"context"
)

// starter.go something

// BootStartByCenter 根据center cluster模式启动
func BootStartByCenter(optionMap map[string]interface{}) *rClient {

	ctx, cancel := context.WithCancel(context.Background())
	c := &rClient{
		ctx:    ctx,
		cancel: cancel,
	}

	// TODO:给dispatch初始化

	var addr string
	if addrIn, ok := optionMap["managerAddr"]; ok {
		addr = addrIn.(string)
	}

	var dp dispatch.Dispatcher
	dp = dispatch.NewByDefaultCluster()
	dp.Link(addr)

	c.lg = logic.NewClient(dp)

	//TODO: 发一个包，拉取配置

	/*

		先不去拉配置 后面再加

		// 假设这个getConfig是sub service的一个远程函数
		var getConfig = func(...interface{}) (interface{}, error) {
			return map[string]interface{}{}, nil
		}

		// 假设m就是拿到的远程配置
		m, err := getConfig()

		// TODO:根据拿到的远程配置来修改配置
		// do some thing
		fmt.Println(m, err)
		// 修改配置之前的一系列调用全部都是按默认配置来的
	*/
	return c
}



