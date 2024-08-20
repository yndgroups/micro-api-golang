# 测试用列

## 代码实例
- 创建demo_test.go写入下面代码

```
package test

import (
	"fmt"
	"testing"
)

func DemoHello(t *testing.T) {
	fmt.Println("DemoHello")
}

func DemoWorld(t *testing.T) {
	fmt.Println("DemoWorld")
}

```

## 运行测试
- 测试整个文件
```
go test -v demo_test.go
```

- 测试单个函数
```
go test -v demo_test.go -test.run TestYouAge
```