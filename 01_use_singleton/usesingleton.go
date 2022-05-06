package main

import (
	singleton "go-design-pattern/01singleton"
)

func main() {
	singleton.GetInstance().Connect("127.0.0.1:2000")

	singleton.GetInstance().GetConnectInfo()

	singleton.GetInstance().Connect("333.0.0.1:4444")

	singleton.GetInstance().GetConnectInfo()
}
