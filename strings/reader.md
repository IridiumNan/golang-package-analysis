# String Reader

`strings.Reader` 将原始的字符串包装成一个 只读的 `io.Reader` 结构体

当我们需要传入一个 内容为当前字符串的 Reader 供其他函数调用的时候， 就可以使用 `strings.NewReader` 来创建新的结构体

它实现了多种接口, 包括

```go
// [io.Reader] 接口
func (r *Reader) Read(b []byte) (n int, err error)

// [io.ReadAt] 接口
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)

// [io.WriterTo] 接口
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
```
