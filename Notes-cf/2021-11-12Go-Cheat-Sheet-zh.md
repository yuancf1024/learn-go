# Go-Cheat-Sheet-zh

# Readme

1. 可用于快速复习Go语言的核心语法;

2. 用于

An overview of Go syntax and features.

Go语法和功能概述。

原作者Github repository：[golang-cheat-sheet](https://github.com/a8m/golang-cheat-sheet#go-cheat-sheet)

# 索引

[toc]

## Credits

Most example code taken from [A Tour of Go](http://tour.golang.org/), which is an excellent introduction to Go. If you're new to Go, do that tour. Seriously.

大多数示例代码摘自A Tour of Go，这是一个很好的Go介绍。如果你是新手，那就去看看。认真对待。

## Go in a Nutshell

* 命令型语言
* 静态类型
* 语法特性类似于C（但是更少的括号、没有分号）和Oberon-2的结构
* 编译到本机代码（无JVM）
* 没有类，但是具有方法的结构体
* 接口
* 没有实现继承。但是有嵌入类型
* 函数是一等公民
* 函数能返回多值
* 有闭包
* 指针，但没有指针算术
* 内置并发原语：Goroutines 和 Channels

# 基础语法

## Hello World

文件 `hello.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello Go")
}
```

`$ go run hello.go`

## Operators

### Arithmetic

| Operator | Description|
|-|-|
| `+` | 加
| `-` | 减
| `*` | 乘
| `/` | 除
| `%` | 取余
| `&` | 按位与
| `\|` | 按位或
| `^` | 按位异或（01、10为真，11、00为假）
| `&^` | 按位与非 即z = x &^ y运算相当于先把y取反（针对y的每个bit：0变成1，1变成0），然后再和x进行&运算 *把共同存在的位清除掉*
| `<<` | 左移位
| `>>` | 右移位

### Comparison

| Operator | Description|
|-|-|
| `==` | 相等
| `!=` | 不等
| `<` | 小于
| `<=` | 小于等于
| `>` | 大于
| `>=` | 大于等于

### Logical

| Operator | Description|
|-|-|
| `&&` | 逻辑与
| `\|\|` | 逻辑或
| `!` | 逻辑非

### Other

| Operator | Description|
|-|-|
| `&` | 创建指针/的地址
| `*` | 引用指针
| `<-` | 发送/接收操作符（参见下面的`Channels`)

## 声明

类型在标识符后面！

```go
var foo int // 没有初始化的声明
var foo int = 42 // 初始化的声明
var foo, bar int = 42, 1302 // 一次声明并给多个变量赋值
var foo = 42 // 类型省略，将被推断出来
foo := 42 // 简写，仅出现在函数体，省略 var 关键字，类型总是隐式的
const constant = "This is a constant"

// iota 可用于递增数字，从0开始
const (
    _ = iota
    a
    b
    c = 1 << iota
    d
)
    fmt.Println(a, b) // 1 2 (跳过0)
    fmt.Println(c, d) // 8 16 (2^3, 2^4)
```

## 函数

```go
// 一个简单的函数
func functionName() {}

// 带参函数（再强调一次，类型在标识符后面）
func functionName(param1 string, param2 int) {}

// 相同类型的多参数
func functionName(param1, param2 int) {}

// 返回类型声明
func functionName() int {
    return 42
}

// 能一次返回多值
func returnMulti() (int, string) {
    return 42, "foobar"
}
var x, str = returnMulti()

// 简单的通过return返回多个已命名的结果
func returnMulti2() (n int, s string) {
    n = 42
    s = "foobar"
    // n 和 s 将会被返回
    return
}
var x, str = returnMulti2()
```

### 函数作为值和闭包

```go
func main() {
    // 赋给函数一个名称
    add := func(a, b int) int {
        return a + b
    }
    // 使用函数名调用函数
    fmt.Println(add(3, 4))
}

// 闭包，作用域：函数可以在定义时访问`scope`内的值
func scope() func() int {
    outer_var := 2
    foo := func() int {return outer_var}
    return foo
}

func another_scope() func() int {
    // 不能编译，因为 outer_var 和 foo 没有在该函数范围内定义
    outer_var = 444
    return foo
}

// 闭包
func other() (func() int, int) {
    outer_var := 2
    inner := func() int {
        outer_var += 99 // 来自外部范围的 outer_var 是可变的
        return outer_var
    }
    inner()
    return inner, outer_var // 返回inner函数并改变outer_var的值为101
}
```

### 变参函数

```go
func main() {
    fmt.Println(adder(1, 2, 3)) // 6
    fmt.Println(adder(9, 9)) // 18

    nums := []int{10, 20, 30}
    fmt.Println(adder(nums...)) // 60
}

// 通过使用...在最后一个参数的类型名称之前，您可以指示它需要零个或多个参数。
//像任何其他函数一样调用函数，除了我们可以通过我们想要的许多参数传递。
func adder(args ...int) int {
    total := 0
    for _, v := range args { // 无论多少数字都会迭代参数。
        total += v
    }
    return total
}
```

## 内置类型

```go
bool

string

int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8的别名
rune // int32的别名 〜= 一个字符（Unicode代码点） - very Viking

float32 float64

complex64 complex128
```

所有GO的预划分的标识符都定义在[内置](https://golang.org/pkg/builtin/)包中。

## 类型转换

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// 替代语法
i := 42
f := float64(i)
u := uint(f)
```

## 包

* 包声明在每个源文件的顶部
* 可执行文件在package `main`中
* 惯例：包名称==导入路径的最后一个名字（import path `math/rand` => package `rand`）
* 大写标识符：可被导出（其他包可见）
* 小写标识符：私有（其他包不可见）

## 结构控制

### if

> 真实代码中的 `if` 语句内部不应该出现`return`，示例代码是不可运行的，将相应`return` 语句改为 `fmt.Println()` 即可。

```go
func main() {
    // 基本
    if x > 10 {
        return x
    } else if x == 10 {
        return 10
    } else {
        return -x
    }

    // 可以在条件前放一个声明
    if a := b + c; a < 42 {
        return a
    } else {
        return a - 42
    }

    // if 内部的类型断言
    var val interface{} = "foo"
    if str, ok := val.(string); ok {
        fmt.Println(str)
    }
}
```

### Loops

```go
// Go 中只有 `for`, 没有 `while` `until`
for i := 1; i < 10; i++ {
}
for ; i < 10; { // while - loop
}
for i < 10 { // 如果只有一个条件，您可以省略分号
}
for { // 可以省略条件 ~ while (true)
}

// 在当前循环使用 break/continue
// 在外圈循环带着标签使用 break/continue

here:
    for i := 0; i < 2; i++ {
        for j := i + 1; j < 3; j++ {
            if i == 0 {
                continue here
            }
            fmt.Println(j)
            if j == 2 {
                break
            }
        }
    }

there:
    for i := 0; i < 2; i++ {
        for j := i + 1; j < 3; j++ {
            if j == 1 {
                continue
            }
            fmt.Println(j)
            if j == 2 {
                break there
            }
        }
    }
```

### Switch

```go
// switch 声明
switch operatingSystem {
    case "darwin":
        fmt.Println("Mac OS Hipster")
        // 案例自动跳出，默认情况下没有 fallthrough
    case "linux":
        fmt.Println("Linux Geek")
    default:
        // Windows, BSD, ...
        fmt.Println("Other")
}

// 与 for 和 if 一样，您可以在 switch 值之前具有一条分配声明语句
switch os := runtime.GOOS; os {
    case "darwin": ...
}

// 您还可以在 switch 案例中进行比较
number := 42
switch {
    case number < 42:
        fmt.Println("Smaller")
    case number == 42:
        fmt.Println("Equal")
    case number > 42:
        fmt.Println("Greater")
}

// 案例可以在逗号分隔的列表中呈现
var char byte = '?'
switch char {
    case ' ', '?', '&', '=', '#', '+', '%':
        fmt.Println("Should escape")
}
```

## Arrays, Slices, Ranges

### 数组

```go
var a [10]int // 声明一个长度为10的整型数组，数组长度是类型的一部分
a[3] = 42 // 设置元素值
i := a[3] // 读取元素值

// 声明和初始化
var a = [2]int{1, 2}
a := [2]int{1, 2} // 简写
a := [...]int{1, 2} // elipsis - >编译器识别输出数组长度
```

### 切片

```go
var a []int // 声明一个切片，类似于一个数组的声明，但是长度不确定
var a = []int{1, 2, 3, 4} // 声明并初始化一个切片（由隐式给出的数组备份）
a := []int{1, 2, 3, 4} // 简写
chars := []string{0: "a", 2: "c", 1: "b"} // ["a", "b", "c"]

var b = a[lo:hi] // 创建一个索引从lo到hi-1的切片（数组的视角）
var b = a[1:4] // 索引从1到3的切片，含左不含右
var b = a[:3] // 缺少低维索引意味着从0开始
var b = a[3:] // 缺少高维索引意味着到len(a)结束
a = append(a, 17, 3) // 向切片a增加新值
c := append(a,b...) // 连接切片a和b

// 用make创建切片
a = make([]byte, 5, 5) // 第一个5代表长度，第二个5代表容量
a = make([]byte, 5) // 容量是可选的

// 从一个数组创建一个切片
x := [3]string{"Лайка", "Белка", "Стрелка","喜欢"，"贝尔卡"，"箭头"}
s := x[:] // 切片引用x的存储
```

### 数组和切片的操作

`len(a)` 为您提供数组/切片的长度。它是一个内置函数，而不是数组上的属性/方法。

```go
// 在一个数组或切片上循环
for i, e := range a {
    // i 是索引，e是元素
}

// 如果你只需要元素
for _, e := range a {
    // e 是元素
}

// 如果你仅仅需要索引
for i := range a {
}

// 在go pre-1.4中，如果您不使用i和e，您将获得编译器错误。
// Go 1.4介绍了一种无变量的形式，这样您就可以这样做了
for range time.Tick(time.Second) {
    // 每一秒钟执行一次
}
```

## 映射

```go
m := make(map[string]int)
m["key"] = 42
fmt.Println(m["key"])

delete(m, "key")

elem, ok := m["key"] // 测试是否存在键 "key" 并检索它，如果存在的话

// 映射的语法
var m = map[string]Vertex {
    "Bell Labs": {40.68433, -74.39967},
    "Google": {37.42202, -122.08408},
}

// 遍历映射内容
for key, value := range m {
}
```

## 结构体

没有类，只有结构体。结构体可以有方法。

```go
// 结构是一种类型。它也是一个领域的集合

// 声明
type Vertex struct {
    X, Y int
}

// 创建
var v = Vertex{1, 2}
var v = Vertex{X: 1, Y: 2} // 通过使用键定义值来创建结构体

// 访问成员
v.X = 4

// 您可以在结构上声明方法。 您要在 `func` 关键字和方法名称之间介绍（接收类型）的结构。在每个方法调用上结构被复制（！）
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 调用方法
v.Abs()

// 对于变化的方法，您需要将指针（请参阅下面）作为结构的类型。 有了这个，没有为方法调用而复制结构值。
func (v *Vertex) add(n float64) {
    v.X += n
    v.Y += n
}
```

**匿名结构体** ： 比起使用 `map[string]interface` 更快更安全。

```go
point := struct {
    X, Y int
} {1, 2}
```

## 指针

```go
p := Vertex{1, 2} // p是一个Vertex结构体
q := &p // q 是一个指向结构体 Vertex 的指针
r := &Vertex{1, 2} // r也是一个指向结构体 Vertex 的指针

// 指向 Vertex 的指针是 *Vertex

var s *Vertex = new(Vertex) // 新建指向新结构实例的指针
```

## 接口

```go
// 接口声明
type Awesomizer interface {
    Awesomize() string
}

// 类型*不是*声明来实现接口
type Foo struct {}

// 相反，如果它们实现所有必需的方法，则类型隐式地满足接口
func (foo Foo) Awesomize() string {
    return "Awesome!"
}
```

## Embedding

GO语言没有子类。相反，存在接口和结构体嵌入。

```go
// Readwriter实现必须满足 Reader 和 Writer
type ReadWriter interface {
    Reader
    Writer
}

// Server 公开了 Logger 具有的所有方法
type Server struct {
    Host string
    Port int
    *log.Logger
}

// 通常初始化嵌入式类型的方式
server := &Server{"localhost", 80, log.New(...)}

// 在嵌入式结构上实现的方法被省略
server.Log(...) // 调用 `server.Logger.Log(...)`

// 嵌入式类型的字段名称是其类型名称（在本例 Logger中）
var logger *log.Logger = server.Logger
```

## 错误

没有异常处理。相反，可能产生错误的函数只是声明 `error` 类型的额外返回值。这是 `error` 接口：

```go
// 错误内置接口类型是表示错误条件的传统接口，使用nil值表示没有错误。
type eror interface {
    Error() string
}
```

这里是一个例子：

```go
func sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, errors.New("negative value")
    }
    return math.Sqrt(x), nil
}

func main() {
    val, err := sqrt(-1)
    if err != nil {
        // 处理错误
        fmt.Println(err) // negative value
        return
    }
    // All is good, use `val`
    fmt.Println(val)
}

```

# 并发

## Goroutines

Goroutines 是轻量级线程（由Go管理，而不是OS线程）。`go f(a, b)` 启动新的goroutine 运行 `f`（给定 `f` 是函数）。

```go
// 只是一个函数（稍后始于 goroutine）
func doStuff(s string) {
}

func main() {
    // 在一个 goroutine 中使用一个命名函数
    go doStuff("foobar")

    // 在一个 goroutine 中使用一个匿名函数
    go func (x int) {
        // 函数体从这里开始
    } (42)
}
```

## Channels

```go
ch := make(chan int) // 创建一个整型的 channel
ch <- 42 // 向 channel ch 发送一个值
v := <- ch // 从 channel ch 接收一个值 

// Non-buffered channels block. Read blocks when no value is available, write blocks until there is a read.
// 无缓冲 channel 块。没有值可用时，读取块，直到有值读取时写入块。

// Create a buffered channel. Writing to a buffered channels does not block if less than <buffer size> unread values have been written.
// 创建一个缓冲 channel, 写入缓冲通道不会阻塞，如果小于缓冲区大小的未读值已写入。
ch := make(chan int, 100)

close(ch) // 关闭 channel (只有发送方应该关闭)

// 从 channel 读取并测试它是否已关闭
v, ok := <- ch

// 如果ok为false,那么channel已经被关闭

// 从 channel 读取知道它被关闭
for i := range ch {
    fmt.Println(i)
}

// 在多个 channel 操作中选择块，如果其中一个块停止阻塞，则执行相应的case
func soStuff(channelOut, channelIn chan int) {
    select {
        case channelOut <- 42:
            fmt.Println("We could write to channelOut!")
        case x := <- channelIn:
            fmt.Pirntln("We could read from channelIn")
        case <- time.After(time.Second * 1):
            fmt.Println("timeout")
    }
}
```

### Channel Axioms

**nil channel 无论是接收还是发送都会永久阻塞**

* A send to a nil channel blocks forever

```go
var c chan string
c <- "Hello, World!"
// 致命错误:所有的 goroutines 都在休眠 - 死锁!
```

* A receive from a nil channel blocks forever

```go
var c chan string
fmt.Println(<-c)
// 致命错误:所有的 goroutines 都在休眠 - 死锁!
```

* A send to a closed channel panics

```go
var c = make(chan string, 1)
c <- "Hello, World!"
close(c)
c <- "Hello, Panic!"
// panic: 向关闭的 channel 发送值
```

* A receive from a closed channel returns the zero value immediately

```go
var c = make(chan int, 2)
c <- 1
c <- 2
close(c)
for i := 0; i < 3; i++ {
    fmt.Println("%d ", <-c)
}
// 1 2 0
```

## Printing

```go
fmt.Println("Hello, 你好, नमस्ते, Привет, ᎣᏏᏲ") // 基本的输出打印, 加上换行
p := struct {X, Y int}{17, 2 }
fmt.Println("My point:", p, "x coord=", p.X) // 打印结构体,整型等
s := fmt.Sprintln("My point:", p, "x coord=", p.X) // 打印字符变量

fmt.Pritnf("%d hex:%x bin:%b fp:%f sci:%e", 17, 17, 17, 17.0, 17.0) // C系格式
s2 := fmt.Sprintf("%d %f", 17, 17.0) // 格式化打印字符串变量

hellomsg := `
"Hello" in Chinese is 你好 ('Ni Hao')
"Hello" in Hindi is नमस्ते ('Namaste')
`
// 多行字符串语法，在开始和结束时使用back-tick(`)
```

## Reflection

### Type Switch

Type Switch 类似于常规的switch语句，但是 Type Switch 中的情况指定了类型(而不是值)，并且这些值将与给定接口值所持有的值的类型进行比较。

```go
func do(i interface{}) {
    switch v := i.(type) {
        case int:
            fmt.Printf("Twice %v is %v\n", v, v*2)
        case string:
            fmt.Printf("%q is %v bytes long\n", v, len(v))
        default:
            fmt.Printf("I don't know about type %T!\n", v)
    }
}

func main() {
    do(21)
    do("hello")
    do(true)
}
```

# 小抄

## Files Embedding

go程序可以使用 `"embed"` package 嵌入静态文件，如下所示：

```go
package main

import (
    "embed"
    "log"
    "net/http"
)

// 内容包含静态内容（2个文件）或Web服务器。 
// go: embed a.txt b.txt

var content embed.FS

func main() {
    http.Handle("/", http.FileSever(http.FS(content)))
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

[Full Playground Example](https://play.golang.org/p/pwWxdrQSrYv)

## HTTP Server

```go
package main

import (
    "fmt"
    "net/http"
)

// 为响应定义一种类型
type Hello struct {}

// 让该类型实现ServeHTTP方法(在http.Handler接口中定义)
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello!")
}

func main() {
    var h Hello
    http.ListenAndServe("localhost:4000", h)
}

// Here's the method signature of http.ServeHTTP:
// type Handler interface {
//     ServeHTTP(w http.ResponseWriter, r *http.Request)
// }
```