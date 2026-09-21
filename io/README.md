# GitLab Post

原帖查看 <https://about.gitlab.com/blog/compose-readers-and-writers-in-golang-applications/>

这里做了删减和融入一些我自己的理解

---

## 接口 (Interface)

> 首先先上点私货， 讲解一下 **接口到底是什么, 以及它的应用场景**

接口是 golang 当中极其重要的概念. 我们知道 `struct` **一般包含存储真实数据和状态的属性以及方法**

理论上使用 struct 就可以实现所有的需求， 那么为什么还需要接口这个东西呢？

我个人认为， 接口是一种抽象， 这种抽象存在的价值在于

- 隐藏不必要的复杂度
- 进行不同模块的解耦
- 规定统一的用途， 使用相同的代码进行处理

我们举一个非常简单的例子

就是拷贝文件和传输文件, 网络传输的实现肯定和本地的文件拷贝是不同的实现。

但是它们拥有共同点 **都拥有数据的来源和目标**, 核心的操作就是 **读和写**

如果我们单纯使用结构体进行传输， 就需要使用不同的调用和处理方式， 遇到不同的情况时需要记忆各种方式，并且还要写两套不同的拷贝实现。

但是我们如果统一暴露 Read 和 Write 接口， 也就是数据的来源不关心内容是什么， 来自哪里， 调用方只需要知道， 读取数据调用它的 Read 方法即可. 同样的我们定义 Write 方法， 不关心内部是写入本地文件还是传输网络内容， 只提供相应的数据进去， 数据如何打包， 压缩， 处理， 我们不关心。

这样我们只需要使用 Read 和 Write 这两个方法就可以将网络的文件传输和本地的文件拷贝统一起来。

举个例子

```go
type NetConn struct {
    URL string
    // 其他的相关内容...
}

func (nc *NetConn) Read()

func (nc *NetConn) Write()

type LocalFile struct {
    FileName string
    // 其他的相关内容...
}

func (lf *LocalFile) Read()

func (lf *LocalFile) Write()
```

这两个结构体都实现了 Read 和 Write 方法， 然后我们希望对它们进行统一处理， 比如说我根据它们的 Read 和 Write 实现写了一个新的方法专门用来把数据从一个地方拷贝到另一个地方

这个时候就需要提供数据的来源这个对象以及数据的目的地这个对象

但是 `NetConn` 和 `LocalFile` 是不同的结构体， 如何实现统一的传入和调用呢？

这个时候就需要一种新的类型, 也是 golang 当中极其核心的类型 `interface`

```go
type Writer interface {
    Write()
}

type Reader interface {
    Read()
}
```

然后我们进行 Copy 函数的定义

```go

func Copy(dst Writer, src Reader)
```

这里需要注意， 这里不需要写 *Writer 或者 *Reader*, 因为 interface 本身就是一种指针

相应的， **传入具体的结构体的时候， 应该传入指针而不是结构体本身**

这样一来我们的 Copy 函数就不需要关心 数据到底怎么来， 如何写入和传输， 实现了 **复杂度的分离**

这也就是接口最为典型的应用, 我们后面会提到 **依赖注入** 的方法， 可以非常方便地实现解耦

需要提到的是， golang 的接口是鸭子类型， 所谓鸭子类型， 也就是只要实现了这个方法， 这个结构体就可以被视作这个接口， 不需要 进行 **显式的继承**. 实现方法就是实现接口

---

> 在golang程序中使用 Reader 和 Writer

## Reader & Writer

Golang 的 `io` 报提供了 `Reader` 和 `Writer` 接口对于 I/O 的不同功能性实现进行了抽象

`Reader` 是一个只包含 `Read` 方法的接口

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

- 这里的 p 是一个缓冲区， 读取到的内容会覆盖掉这个缓冲区的内容
- n 是实际写入缓冲区的长度， 如果缓冲区满了， n = len(p), 如果剩下的数据没有办法充满缓冲区， 则返回一个更小的 n
- err 读取过冲中的错误， 如果全部读取完毕， 会返回 `io.EOF`

`Writer` 是一个只包含 `Write` 方法的接口

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

- p 是缓冲区， Writer 会尝试将这里的数据全部写入到目标
- n 是有效写入的长度， 正常会全部写入 n = len(p), 写入失败的时候， 会返回一个更小的 n 值
- err 是写入过程中返回的错误，根据具体的实现而不同

---

比如 `os` package 提供了一种实现了文件的读取. `File` 这个结构体实现了 `Reader` 和 `Writer` 接口 (只需要实现 `Read` 和 `Write`)

我们这里举一个简单的例子， 使用 Write 和 Read 方法来实现从文件中读取内容， 并将这些内容打印出来

首先先准备一个文件

```bash
cat > data.txt <<EOF
Hello Go
Hello golang
Hello interface
EOF
```

然后创建一个 `main.go` 文件

```go
func main() {
    // 打开 data.txt 文件
    file, err := os.Open("data.txt")
    if err != nil {
        log.Fatal(err)
    }
    // 在函数结束的时候， 自动关闭文件
    defer file.Close()

    // 创建一个缓冲区, 大小为 32K
    p := make([]byte, 32 * 1024)
    for {
        // 从文件中读取
        n, err := file.Read(p)

        // 将缓冲区的有效内容写入到 stdout (也就是直接打印出来)
        _, errW := os.Stdout.Write(p[:n])
        if errW != nil {
            log.Fatal(errW)
        }

        if err != nil {
            if errors.Is(err, io.EOF) {
                break
            }

            log.Fatal(err)
        }
    }
}

```

然后使用 `go run ./main.go` 执行即可

其中 os.Stdout 是一个特殊的Writer, 往里面写入的内容会直接显示到 标准输出当中

它对应的是 Linux 系统当中的 `/dev/stdout` 这个抽象的文件

**Writer 和 Reader 接口是 golang io 生态的绝对核心**

其小接口的设计也是接口设计的典范 (所谓小接口就是方法很少， 通用性很强， 抽象层次高)

---

标准库还基于这两个接口， 提供了一些非常实用的方法。

我们自己也可以实现 Write 和 Read 方法进而方便地调用这些工具函数。

但是应该要注意的是， 函数的签名 (也就是函数的输入和输出， 类型必须一致， 并且要保持含义的一致性， 否则可能不会实现预期的效果)

## Copy

非常典型的就是我们刚才提到的 `Copy` 函数
它将 Reader 当中的所有内容都写入到 Writer 当中

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

这里的 written 就是实际有效的写入值, 这样我们就可以避免每次都手动截断和判断是否结束

我们刚才也看到了， Write 和 Read 方法的输入都是一个 []byte 类型， 也就是数据的缓冲区， 在 golang 的标准 `io.Copy` 方法当中， 它的大小是 `32K`. 当我们希望增大这个缓冲区或者减小这个缓冲区的大小 (或者说在多次的 Copy 调用中复用同一个缓冲区), 我们可以使用 `io.CopyBuffer`函数

```go
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
```

多了一个参数是 buf, 也就是自己手动传入的缓冲区， 这个缓冲区在 拷贝到过程中会被不断地写入和读取

---

## MultiWriter

刚才我们提到的是简单的 1 Reader -> 1 Writer 的模式

而 golang 还提供了一个很好的结构体帮助我们实现 1 Reader -> n Writer 的模式

它就是 `MultiWriter`

```go
func MultiWriter(writers ...Writer) Writer
```

这个函数接收多个 Writer 之后返回一个 Writer, **每次调用这个返回的新的 Writer 的 Write 进行写入的时候, 会将内容写入到这个 创建这个 MultiWriter 的传入的每一个 Writer 里面**

比如说我们希望将日志同时写入远程的服务器 (用 nfs 挂载文件) 以及本地的文件， 打开两个文件之后构建一个 MultiWriter， 然后每次写入日志的时候只需要调用它的 Write 方法即可

> 当然了几乎在任何时候， 我们都不需要自己手动调用 Write 方法， 而是只需要提供这个 Writer,很多标准库和第三方库会自动调用这个方法进行写入

还有一个特殊的应用就是实现像 `tee` 一样的功能， 将输出打印到标准输出， 并写入到文件当中

牢记标准输出也被操作系统抽象成了一个文件， 所以 `os.Stdout` 也是一个 Writer

这是一个简单的示例代码

```go
func tee(r io.Reader) {
    file, err := os.CreateTemp(".", "tmp-*.out")
    if err != nil {
        log.Fatal(err)
    }

    mw := io.MultiWriter(file, os.Stdout)

    written, err := io.Copy(mw, r)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("written size: %d", written)

    fmt.Printf("see tmp file: %s", file.Name())
}

func main() {
    r := bytes.NewBufferString("hello world\n")

    tee(r)
}
```

这里的 NewBufferString 就是把一个字符串变成一个拥有 Read 方法的 Buffer， 然后传入作为 tee 的参数

---

## MultiReader

MultiWriter 为我们提供了方便的多 dst 写入的方式

自然的就想到 如果有多个数据源， 是否也可以含方便的进行整合

golang 也提供了一个非常有用的结构体， 也就是 `MultiReader`

```go
func MultiReader(readers ...Reader) Reader
```

但是我们每次调用 Read 方法从一个地方读取， 所以它并不是同时读取。 而是将多个数据源进行排列， 线性进行读取。 就像队列一样， 一个数据源读空了， 就换下一个

那么我们可以使用这个方式来对多个数据源进行整合和综合解析

---

## MultiReader -> MultiWriter

因为 MultiWriter 和 MultiReader 实现的是标准的 Write 和 Read 方法， 所以我们也可以直接使用 io.Copy 方法将一个 MultiReader 当作数据源， 写入到 MultiWriter 当中

```go
written, err := io.Copy(io.MultiWriter(w1, w2, w3), io.MultiReader(r1, r2, r3))
```

---

## TeeReader

这是一个特殊的 Reader

正如它的命名一样， 它的可以实现类似 Tee 的效果

构造一个新的 teeReader 需要传入一个 Reader 和 Writer

```go
func io.TeeReader(r io.Reader, w io.Writer) io.Reader
```

然后返回这个 Reader, 注意原来的内容不会马上写入到 Writer 当中， 而是在 teeReader 的 Read 的方法被调用的时候才同时写入

可以这样理解， 本来的 Read 开了一条固定的支流将数据同步写入到 初始化传入的 Writer 中

常用的场景有 同时上传文件和计算哈希值, 最终进行哈希校验

原帖子当中还提到一个场景是 将接受到的资源同时传送给用户和使用云端对象存储 来备份， 当然使用 MultiWriter 也可以实现相同的功能

---
