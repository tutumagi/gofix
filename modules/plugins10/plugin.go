package main

import (
	"fmt"
	"plugin_fix/route"
)

// go build -buildmode=plugin
func InitPlugin(r *route.Route) {
	fmt.Printf("register router\n")

	r.Register("error", DoErr)
	r.Register("info", DoInfo)
}

func DoErr() {
	fmt.Println("this is err 我的天啊")
}

func DoInfo() {
	fmt.Println("this is info 你的地啊")
}
