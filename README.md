# go_sample
## 参考
- [参考1](http://www.geocities.jp/m_hiroi/golang/index.html)
- [参考2](https://oohira.github.io/gobyexample-jp/)

## goをMacにインストール(How to install Go language)

  [b4b5a062]: http://www.geocities.jp/m_hiroi/golang/index.html "title"

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
### 並行プログラミング
「並行 (concurrent) プログラミング」は複数のプログラムを並行に実行しますが、このとき複数の CPU で同時に動かす場合と、ひとつの CPU で複数のプログラムを動かす場合があります。一般的には、前者を「並列 (parallel) プログラミング」といい、複数のハードウェアを並列に実行することによる処理速度の向上が主な目的となります。

後者の場合、一定時間毎に実行するプログラムを切り替えることで、複数のプログラムを並列に実行しているかのように見せることができます。この処理を「時分割 (time sharing) 」もしくは「タイム・スライス (time slice) 」といいます。一般に、タイム・スライスは OS でサポートされている機能です。OS が実行するプログラムのことを「プロセス (process) 」または「タスク (task) 」といいます。

並列的に動作するプログラムをひとつのプロセスだけで作るのはけっこう大変です。そこで、プロセス内では逐次的な処理にとどめ、複数のプロセス間で情報交換を行うことにより、全体で並列的な動作を実現することを考えます。このほうが自然にプログラムを記述できる場合があるのです。これが後者の主な目的となります。

プロセスは互いに独立したプログラムですが、OS にはプロセス間でデータをやり取りする機能 (プロセス間通信) が用意されています。たとえば、UNIX ライクな OS では「パイプ (pipe) 」を使って複数のプログラム (コマンド) を連結することができます。この場合、パイプを通してデータがプログラムに送られ、各プログラムは並行に動作することになります。入出力処理で待たされるコマンドがあったとしても、その間に他のコマンドが実行されるので、各コマンドを逐次的に実行するよりも、効率的に処理することが可能です。

最近では、ひとつのプログラムの中で独立した複数の処理を記述できるようになりました。この機能を「スレッド (thread) 」とか「マルチスレッド」いいます。スレッドは「縫い糸」という意味ですが、プログラムでは「制御の流れ」という意味で使われています。並列的な動作をプログラムする場合、逐次的な処理をひとつのスレッドに割り当て、複数のスレッドを並行に動作させることにより、全体で並列的な動作を実現するわけです。

一般に、スレッドは一定時間毎に実行するスレッドを強制的に切り替えます。このとき、スレッドのスケジューリングは処理系が行います。これを「プリエンプティブ (preemptive) 」といいます。これに対し、Ruby のファイバーや Lua のコルーチンは、プログラムの実行を一定時間毎に切り替えるものではありません。他のプログラムが実行できるよう自主的に処理を中断する、といった協調的な動作を行わせることで、複数のプログラムを並行に動作させています。これを「ノンプリエンプティブ (nonpreemptive) 」といいます。

スレッドは同じプロセス内に存在するので、メモリ空間を共有することができます。これを「共有メモリ」といいます。スレッド間の通信は共有メモリを使って簡単に行うことができますが、プリエンプティブなスレッドの場合、共有メモリのアクセス時に発生する「競合」が問題になります。このため、プリエンプティブなマルチスレッドをサポートしているプログラミング言語では、競合を回避するための仕組みが用意されています。

Go 言語の goroutine はプリエンプティブなマルチスレッドに近いもので、goroutine は OS ではなく Go 言語が管理します。Go 言語の goroutine は共有メモリを使って通信することができますが、それは推薦されていないようです。そのかわり「チャネル (channel) 」という通信路が用意されいて、それを使って goroutine 間でデータの送受信を行うことができます。他のプログラミング言語では、関数型言語 Erlang が同様な方法を採用しています。Erlang のプロセスと同じように、Go 言語の goroutine は軽量で高速に動作するといわれています。

なお、Go 言語は「並列プログラミング」にも対応していますが、ver 1.2 のデフォルト設定ではマルチ CPU を利用することはできません。今回は並行プログラミングについて簡単に説明し、並列プログラミングについては次回以降に取り上げることにします。
