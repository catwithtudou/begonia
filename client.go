package begonia

import (
	"github.com/MashiroC/begonia/app/client"
	"github.com/MashiroC/begonia/app/option"
)

// Client 客户端的接口
type Client interface {
	Service(name string) (client.Service, error)
	FunSync(serviceName, funName string) (client.RemoteFunSync, error)
	FunAsync(serviceName, funName string) (client.RemoteFunAsync, error)
	Wait()
	Close()
}

// NewClient 初始化，获得一个service对象，传入一个mode参数，以及一个option的不定参数
func NewClient(mode string, optionFunc ...option.WriteFunc) (cli Client) {
	optionMap := defaultClientConfig()

	for _, f := range optionFunc {
		f(optionMap)
	}

	switch mode {
	case "center":
		cli = client.BootStartByCenter(optionMap)
		// TODO:其他的模式和模式出问题的判断
	}

	return
}

func defaultClientConfig() map[string]interface{} {
	m := make(map[string]interface{})

	// TODO:加入默认配置

	return m
}