package singleton

import (
	"fmt"
	"sync"
)

//DBClient 通过该接口可以避免 GetInstance 返回一个包私有类型的指针
type DBClient interface {
	Connect(msg string)
	GetConnectInfo() string
}

type dbClient struct {
	addr string
}

func (d *dbClient) Connect(addr string) {
	d.addr = addr
}

func (d *dbClient) GetConnectInfo() string {
	fmt.Println("msg=", d.addr)
	return d.addr
}

var (
	instance *dbClient //单例对象
	once     sync.Once
)

//GetInstance 获取单例对象
func GetInstance() DBClient {
	once.Do(func() {
		instance = &dbClient{
			addr: "",
		}
	})

	return instance
}
