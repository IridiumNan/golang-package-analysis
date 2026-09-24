# Bytes Buffer 介绍

> 相应的示例代码 [`buffer_test.go`](./buffer_test.go)

## 基础结构体

`bytes.Buffer` 是一个实用的 结构体

不同于 `bytes.Reader` 的只读， 它是可读可写并且同样实现了各种接口

`bytes.Buffer` 的核心构成非常简单, 就是一个简单的存放原始数据的 `[]byte`, 一个标记当前位置的 `off`, 以及标记上一次读取的偏移位置的 `lastRead`

```go
type Buffer struct {
    buf      []byte // contents are the bytes buf[off : len(buf)]
    off      int    // read at &buf[off], write at &buf[len(buf)]
    lastRead readOp // last read operation, so that Unread* can work correctly.

    // Copying and modifying a non-zero Buffer is prone to error,
    // but we cannot employ the noCopy trick used by WaitGroup and Mutex,
    // which causes vet's copylocks checker to report misuse, as vet
    // cannot reliably distinguish the zero and non-zero cases.
    // See #26462, #25907, #47276, #48398 for history.
}

type readOp int8
```

---

## 常用方法

- 基础访问

```go
// 直接访问尚未读取的剩余数据的 []byte
func (b *Buffer) Bytes() []byte

// 将为读取的数据内容转化为 String 类型之后返回
func (b *Buffer) String() string


// 返回尚未读取的数据切片长度
func (b *Buffer) Len() int

// 返回总共分配的切片大小， 包括已经读取的部分
func (b *Buffer) Cap() int

```

- 接口实现

```go
func (b *Buffer) Write(p []byte) (n int, err error)

func (b *Buffer) Read(p []byte) (n int, err error)

// 从 Reader 当中读取所有内容， 添加到原始的数据后方
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)

// 将所有尚未读取的数据写入到 Writer 当中
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)
```

- 进阶控制

```go

// 预分配空间, 在原来的基础上添加 n 个字节的长度
func (b *Buffer) Grow(n int)

// 清空内容, 清空之后旧的数据就是非法空间， 不能重复访问, 主要用于 buffer 的复用
func (b *Buffer) Reset()

```
