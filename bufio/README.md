# Bufio Package

参考的帖子 <https://medium.com/@emusbeny/mastering-bufio-in-go-the-art-of-buffered-i-o-17cae584ee4b>

---

> [运行案例](./bufio_test.go)

## 什么是 bufio

`bufio` 是一个 golang 的标准库， 它的核心作用是 提供**缓冲式的 IO 操作**。

这会让你的 I/O 操作更加高效， 可以借助缓冲来减少系统调用 (我们都知道 IO 是相当耗时的， 所以减少 IO 的次数就可以让程序跑得更快).

举个非常简单的例子， 比如说有一个文件总共有 5 行， 然后程序内部需要一行一行对数据进行处理。 而我们如果每次只读取一行， 就需要 5 次 IO. 所以更好的方式是一次读取多行， 比如说把 5 行全部读取进来， 每次从缓冲区拿出一行进行相应的操作.

---

## bufio.Reader

Reader 的作用就是一次性读取更多的内容保存到缓冲区当中 (数据直接存储在内存当中), 然后通过各种接口对内部的内容进行读取，如果缓冲区空了， 会继续调用初始化时传入的 Reader 的 Read 方法获取下一批的数据.

一般来说，Reader 会被用于超大文件的处理。对于逐行处理的超大文件， 每次读取一行都需要一次系统调用， 太慢了， 而因为文件本身的体积太大， 我们无法将其直接加载到内存当中， 所以我们会使用 `bufio.Reader` 进行缓冲。

这样一来我们减少了多次的系统调用并始终将内存的占用维持在较为可控的水平。

最常用的方法是 `ReadSlice`， 跟它类似的有 `ReadBytes`, 但是后者会拷贝一份数据， 所以前者性能更好

```go
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
```

这里的delim 就是 delimiter, 也就是分隔符。 一般用 `,` `\n` 之类的， 根据需求使用

**Reader** 同样实现了 `io.Reader` 的 Read 方法， 所以它的灵活性很高。

并且提供了 `ReadString`, `ReadRune` 等方法

> 注意 `ReadLine` 是一个比较底层的方法， 一般不调用， 相同的逻辑应该使用 `ReadSlice(deli)` 或者 `ReadBytes(deli)`, `Scanner` 等实现

举个例子

```go
func TestReaderReadSlice(t *testing.T) {
    file, err := os.Open("../hugeFile.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    r := bufio.NewReader(file)

    sum = 0
    for {
        line, err := r.ReadSlice('\n')
        if err != nil {
            break
        }

        doSomeWork(line)

    }


    t.Logf("The length of all byte data: %d", sum)
}
```

读取到结尾之后， ReadSlice 会返回 `io.EOF`, 所以直接 break 即可， 当然在实际的使用过程中应该判断是否是 `io.EOF`

---

## bufio.Scanner

Scanner 是一个功能强大的分割处理工具， 可以实现各种自定义的分割和扫描

核心操作是 `Scan` -> `Bytes()`/`Text()`

其中 Scan 是扫描文本并将内部的索引存储往前推进， 将分割之前的部分加入到可以通过接口访问的地方

之后 Bytes 或者 Text 函数就是取出上一次扫描获得的数据， 区别在于数据的类型

简单的使用例子

```go
func TestScanner(t *testing.T) {
    file, err := os.Open("../hugeFile.txt")
    if err != nil {
    log.Fatal(err)
    }

    defer file.Close()

    s := bufio.NewScanner(file)

    sum = 0
    for s.Scan() {
        if s.Err() != nil {
            break
        }
        sum += len(s.Bytes())
    }
    t.Logf("The length of all byte data: %d", sum)
}
```

**使用Split方法来制定分割的规则,默认使用的是按行分割**

> 有个小技巧就是标准库已经写了按行分割的函数实现 `Scanner.ScanLine`, 当我们需要指定自定义的分隔符的时候， 直接复制源码当中的主体部分进行修改即可
>
> 关于 Reader 和 Scanner 的选择， 可以查看讨论 <https://www.reddit.com/r/golang/comments/ba387a/how_to_choose_between_bufionewscanner_and/>

---

## 标准输入读取

对于标准输入 `os.Stdin` 的读取， 也就是读取用户在终端当中输入的内容， 一般使用 `Scanner`

我们知道 `os.Stdin` 是一个 `*os.File` 类型， 也就是说我们直接像操作文件一样操作标准输入。 所以当我们初始化的时候将 `os.Stdin` 当成 Reader 传入的之后。 就可以使用 Scanner 来读取终端当中输入的内容。

```go
func main() {
    s := bufio.NewScanner(os.Stdin)

    fmt.Println("starting scanning...")

    s.Scan()

    fmt.Println("blocking...(This will not printed until ENTER is pressed)")

    fmt.Printf("The scanned text: %s", s.Text())
}
```

在上面的例子当中， s.Scan 会阻塞程序直到完成分割

**使用,作为分隔符，直到读取到`,`才取消阻塞**

直接从标准库当中复制函数出来然后改一下

```go

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
    if len(data) > 0 && data[len(data)-1] == '\r' {
        return data[0 : len(data)-1]
    }
    return data
}

func ScanComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
    if atEOF && len(data) == 0 {
        return 0, nil, nil
    }
    if i := bytes.IndexByte(data, ','); i >= 0 {
    // We have a full newline-terminated line.
        return i + 1, dropCR(data[0:i]), nil
    }
    // If we're at EOF, we have a final, non-terminated line. Return it.
    if atEOF {
        return len(data), dropCR(data), nil
    }
    // Request more data.
    return 0, nil, nil
}

func main() {
    s := bufio.NewScanner(os.Stdin)

    fmt.Println("starting scanning for comma...")

    s.Split(ScanComma)
    s.Scan()

    fmt.Println("blocking...(This will not printed until ENTER is pressed and comma appear)")

    fmt.Printf("The scanned text: %s", s.Text())
}
```

然后进行一个测试

```text
starting scanning for comma...
Hello world
Hello Linux
Hello Nixos
,
blocking...(This will not printed until ENTER is pressed and comma appear)
The scanned text: Hello world
Hello Linux
Hello Nixos
```

---

## bufio.Writer

如果你希望将内容写到文件当中或者进行时间消耗较大的网络传输。 应该使用 Writer， 先将内容写入到缓冲区当中， 直到缓冲区被写满或者手动调用 Flush 方法进行刷入数据。

```go

func TestWriter(t *testing.T) {
    file, err := os.Create("output.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    w := bufio.NewWriter(file)
    n, err := w.WriteString("Hello world, this is bufio.Writer")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%d bytes are written", n)
    err = w.Flush()
    if err != nil {
        log.Fatal(err)
    }
}
```

可以使用 `NewWriterSize` 来直接指定缓冲区大小， 默认的是 4096 (4 K)

**不要经常调用 Flush 方法**, 这样就没有办法充分发挥缓冲写入的效果。

只在需要保证全部数据都成功写入 (一般在Write 方法已经全部执行结束)的时候，调用 Flush 保证数据不丢即可
