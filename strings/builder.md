# String Builder

`strings.Builder` 的作用是在构造 和拼接字符串的时候减少分配的次数进而提升性能

当我们的字符串拥有多个部分的时候， 使用预分配空间再进行拼接能够大大减小开销

核心的运作逻辑是

- `Grow` 函数实现空间预分配
- `Write` 或者 `WriteString` 写入需要添加的字符串部分
- `String` 函数获得最终的字符串

具体的示例查看 [builder test](./builder_test.go)
