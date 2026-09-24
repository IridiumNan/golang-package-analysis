# bytes.Reader 介绍

`bytes` 这个 package 当中定义了一个 `Reader` 结构体

主要是用于将原始的 `[]byte` 类型的数据包装成一个可以供多个函数方便调用的接口

`bytes.Reader` 是下列接口的实现

- `io.Reader`
- `io.ReadAt`
- `io.WriteTo`
- `io.Seeker`
- `io.ByteScanner`
- `io.RuneScanner`

并且 `bytes.Reader` 当中的数据是只读的

---

## 主要方法

- 查看长度

```go
// 获取剩余未读的 []byte 长度
func (r *Reader) Len() int

// 获取数据的总长度
func (r *Reader) Size() int64
```

- 标准接口方法

```go
// [io.Reader] 实现
func (r *Reader) Read(b []byte) (n int, err error)

// [io.ReadAt] 实现
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)

// [io.ByteReader] 实现
func (r *Reader) ReadByte() (byte, error)

// ...

// [io.WriteTo] 实现
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
```

- 重置数据指针

```go
// 将指向当前数据的指针重置到开头的位置， 可以反复读取数据
func (r *Reader) Reset(b []byte)
```
