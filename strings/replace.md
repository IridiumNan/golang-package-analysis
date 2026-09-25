# String Replacer

`strings.Replacer` 是一个方便的并发安全的字符串替换结构体

使用方式也非常简单

- 创建新的 Replacer

```go
// 输入的参数是 (旧1, 新1, 旧2, 新2), 也就是说总参数的个数必须是偶数， 第奇数位的参数会被替换成下一位的字符
r := strings.NewReplacer("{name}", "cai", "{language}", "Chinese")
```

- 调用替换函数

```go
// 直接返回字符串
newStr := r.Replace(tmpl)


// 将内容写入到 [io.Writer] 当中
_, err := r.WriteString(&b, tmpl)
```

具体的案例可以查看[replace test](./replace_test.go)
