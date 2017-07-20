# go_sample
## goをMacにインストール(How to install Go language)

`brew install go --cross-compile-common`
<br>クロスコンパイル可能な状態でGoをインストールすることで，MacからLinuxにデプロイすることが可能となる

### 実行
hello.goを作って実行して見る

`go run hello.go`

次にコンパイルするには．．．

```
go build hello.go
./hello
```
### import
```
import "fmt"
import (
   "fmt"
   "math"
)
```

### Print文
```
// 改行
fmt.Println("hello world")

```
### iota
- 連続した数値を定義する
```
const (
    foo = iota   // foo = 0
    bar          // bar = 1
    baz          // bar = 2
)
```

### for
```
a []int = {1,2,3,4}
for i, v := a {
  
}
```
