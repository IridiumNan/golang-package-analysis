# bytes.go

下面是常用的函数

- 判断两个 byte 切片是否相等

```go
func Equal(a, b []byte) bool
```

- 比较两个切片的大小

```go
func Compare(a, b []byte) int
```

- 通过部分切片的出现次数

```go
func Count(s, sep []byte) int
```

- 判断是否存在， 返回bool, 支持切片，rune, string, 自定义函数等

```go
func Contains(b, subslice []byte) bool

func ContainsAny(b []byte, chars string) bool

func ContainsRune(b []byte, r rune) bool

func ContainsFunc(b []byte, f func(rune) bool) bool

func IndexByte(b []byte, c byte) int
```

- 寻找目标的下标索引, 依旧支持各种类型, 可以查找第一个或者最后一个的下标

```go
func IndexByte(b []byte, c byte) int

func LastIndex(s, sep []byte) int

func LastIndexByte(s []byte, c byte) int

func IndexRune(s []byte, r rune) int

func IndexAny(s []byte, chars string) int

func LastIndexAny(s []byte, chars string) int

func Index(s, sep []byte) int
```

- 按照分隔符分割

```go

// 根据分隔符 sep 分割直到不包含分隔符
func Split(s, sep []byte) [][]byte


// 分局分隔符 sep 分割， 分割的次数最多为 n
// 比如说要分割成两个部分就 指定 n = 2
func SplitN(s, sep []byte, n int) [][]byte


// 根据分隔符 sep 分割， 最终的子切片中会包含分隔符
func SplitAfter(s, sep []byte) [][]byte


func SplitAfterN(s, sep []byte, n int) [][]byte
```

- 去除空白部分并分割

> 会被去除的有
> '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0

```go
func Fields(s []byte) [][]byte

// 支持自定义选择跳过的内容, 遇到跳过的内容时候应该返回 true, 其他情况返回false
func FieldsFunc(s []byte, f func(rune) bool) [][]byte
```

- 拼接并指定分隔符

```go
// sep 为分隔符， s 是包含多个byte切片的， 也就是需要拼接的内容
func Join(s [][]byte, sep []byte) []byte
```

- 前缀后缀判断

```go
// 判断是否存在前缀
func HasPrefix(s, prefix []byte) bool

// 判断是否存在后缀
func HasSuffix(s, suffix []byte) bool
```

- 映射

```go
// 使用函数将原始的 rune 映射到新的 rune, 产生一个全新的 []byte 
func Map(mapping func(r rune) rune, s []byte) []byte
```

- 重复

```go
// 将原始的输入重复几次之后返回一个新的 []byte
func Repeat(b []byte, count int) []byte
```

- 大小写转换

```go
// 将所有的字符全部大写
func ToUpper(s []byte) []byte

// 全部小写
func ToLower(s []byte) []byte

// 生成标题
func ToTitle(s []byte) []byte
```

- 裁剪

```go
// 同时对左侧和右侧进行裁剪， 去掉 cutset 当中的字符
func Trim(s []byte, cutset string) []byte

// 只裁剪左侧
func TrimLeft(s []byte, cutset string) []byte

// 只裁剪右侧
func TrimRight(s []byte, cutset string) []byte

// 根据规则裁剪， 函数返回 true 的内容会被裁剪掉
func TrimFunc(s []byte, f func(r rune) bool) []byte


// 裁剪掉前缀
func TrimPrefix(s, prefix []byte) []byte

// 裁剪后缀
func TrimSuffix(s, suffix []byte) []byte


// 裁剪掉空白的字符
func TrimSpace(s []byte) []byte
```

- 替换

```go
// 替换最多指定次数
func Replace(s, old, new []byte, n int) []byte

// 全部替换
func ReplaceAll(s, old, new []byte) []byte
```

- 剪切

```go
// 根据分隔符将第一次出现分隔符之前的内容和之后的内容分开
// 使用 SplitN(s, sep, 2) 可以实现相同的功能
// 如果没有找到则 found = false
func Cut(s, sep []byte) (before, after []byte, found bool)


// 前缀裁剪
func CutPrefix(s, prefix []byte) (after []byte, found bool)

// 后缀裁剪
func CutSuffix(s, suffix []byte) (before []byte, found bool)
```

- 克隆

```go

// 克隆一份原始的数据并返回
func Clone(b []byte) []byte
```
