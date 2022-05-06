package singleton

import (
	"sync"
	"testing"
)

const (
	CurrentCount = 100 //并发数
)

//测试顺序获取单例对象
func TestOderGetInstance(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

//测试并发获取单例对象
func TestCurrentGetInstance(t *testing.T) {
	startCh := make(chan struct{})
	instances := [CurrentCount]DBClient{} //初始化长度为100的接口数组

	var wg sync.WaitGroup
	for i := 0; i < CurrentCount; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			<-startCh //阻塞协程的执行，直到通道关闭
			t.Log("go routine index=", index)
			instances[index] = GetInstance()
		}(i)
	}

	close(startCh) //关闭channel，所有协程同时开始运行
	wg.Wait()

	for index := range instances {
		if index+1 < len(instances) && instances[index] != instances[index+1] {
			t.Fatal("instance is not equal")
		}
	}

}

//结论：无论什么时候调用获取单例对象接口，无论调用多少次获取单例对象接口，返回的都是同一个 对象
