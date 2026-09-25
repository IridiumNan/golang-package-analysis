# Strconv 介绍

`strconv` 用于进行字符串和其他类型之间的相互转换

## String -> Other

从 String 到其他类型的转换使用的是 `strconv.Parsexxx` 类的函数

返回值中第一个是相应的类型， 第二个是错误

```go
// 有个很常用的函数, 专门用来转int类型
// 效果等价于 ParseInt(s, 10, 0)
func Atoi(s string) (int, error)

func ParseBool(str string) (bool, error)

func ParseComplex(s string, bitSize int) (complex128, error)

func ParseFloat(s string, bitSize int) (float64, error)

func ParseUint(s string, base int, bitSize int) (uint64, error)

func ParseInt(s string, base int, bitSize int) (i int64, err error)

```

## Other -> String

从其他类型格式化成相应的 String 使用 `strconv.Formatxxx` 类的函数

返回值是一个字符串

```go

// 常用的函数， 将 int 转 String (十进制), 等价于 FormatInt(i, 10)
func Itoa(i int) string

func FormatBool(b bool) string

func FormatComplex(c complex128, fmt byte, prec, bitSize int) string

func FormatFloat(f float64, fmt byte, prec, bitSize int) string

func FormatUint(i uint64, base int) string

func FormatInt(i int64, base int) string
```

> 如果你希望将格式化出来的内容直接追加到 []byte 当中， 使用 `strconv.Appendxxx` 类的函数

具体的案例查看 [strconv test](./strconv_test.go)
