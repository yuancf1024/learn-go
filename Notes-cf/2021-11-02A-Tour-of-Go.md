# 2021-11-02A-Tour-of-Go

[toc]

## 学习记录

- [x] 2021-11-02 包、变量和函数
- [x] 2021-11-03+08 流程控制语句+更多类型-映射
- [x] 2021-11-09 修改映射~方法和接口-选择值或指针作为接收者
- [x] 2021-11-11 方法和接口该节*后面全部*
- [x] 2021-11-11 完成并发部分 *完结,撒花花!*

> 这是Go官网的一份入门教程，可以快速了解Go语言的基本语法和常见的高级主题，对于从其他语言转向Go的开发者而言将会是一份快速入门Go的教程。
>
> 学习不是一件容易的事情，很早就开始看到并接触了这份教程，但是直到此刻，才下决心完成这份教程。没有来不来得及，现在开始，就来得及。Have fun！

## 1 Welcome

### Hello，世界

欢迎来到 Go 编程语言指南。

点击页面左上角的 Go 指南可以访问本指南的模块列表。

你可以随时点击页面右上角的菜单来浏览内容列表。

本指南由一系列幻灯片贯穿始终，其中有一些练习需要你来完成。

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello，世界")
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\hello\hello.go"
Hello，世界

### Go 本地化

Go 中文版 https://tour.go-zh.org/
Go 英文版 https://tour.golang.org/list

### Go 离线版教程

本指南也可作为独立的程序使用，这样你无需访问互联网就能运行它。 独立的 tour 更快，它会在你自己的机器上构建并运行代码示例。 要在本地安装并运行此教程的中文版，请在命令行执行： `go get -u github.com/Go-zh/tour tour`

该程序会打开一个浏览器并显示你本地版本的教程。 当然，你也可以继续在线学习本教程。 * Go 练习场 本指南构建在 Go 练习场上，它是一个运行在 golang.org 服务器上的一个 Web 服务。 该服务会接收一个 Go 程序，然后在沙盒中编译、链接并运行它，最后返回输出。 在练习场中运行的程序有一些限制： - 练习场中的时间始于 2009-11-10 23:00:00 UTC（此日期的含义留给读者自己去发现）。 赋予程序确定的输出能让缓存程序更加容易。 - 程序的执行时间、CPU 和内存的使用同样也有限制，并且程序不能访问外部网络中的主机。 练习场使用最新发布的 Go 的稳定版本。 参阅 Go 练习场的内部机制了解更多信息。

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Welcome to the playground!")

    fmt.Println("The time is", time.Now())
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\sandbox\sandbox.go"
Welcome to the playground!
The time is 2021-11-02 15:56:53.212646 +0800 CST m=+0.007297401

## 2 包、变量和函数

### 包

每个 Go 程序都是由包构成的。

程序从 `main` 包开始运行。

本程序通过导入路径 `"fmt" `和 `"math/rand"` 来使用这两个包。

*按照约定，包名与导入路径的最后一个元素一致*。例如，`"math/rand"` 包中的源码均以 `package rand` 语句开始。

*注意：* 此程序的运行环境是固定的，因此 `rand.Intn` 总是会返回相同的数字。 （要得到不同的数字，需为生成器提供不同的种子数，参见 `rand.Seed`。 练习场中的时间为常量，因此你需要用其它的值作为种子数。）

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\packages\packages.go"
My favorite number is 1

### 导入

此代码用圆括号组合了导入，这是“分组”形式的导入语句。

当然你也可以编写多个导入语句，例如：

```go
import "fmt"
import "math"
```

不过使用分组导入语句是更好的形式。

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\imports\imports.go"
Now you have 2.6457513110645907 problems.

### 导出名

在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。例如，`Pizza` 就是个已导出名，`Pi` 也同样，它导出自 `math` 包。

`pizza` 和 `pi` 并未以大写字母开头，所以它们是未导出的。

*在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。*

执行代码，观察错误输出。

然后将 `math.pi` 改名为 `math.Pi` 再试着执行一次。

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\exported-names\exported-names.go"
3.141592653589793

### 函数

函数可以没有参数或接受多个参数。

在本例中，`add` 接受两个 `int` 类型的参数。

注意类型在变量名 **之后**。

（参考 这篇关于 Go 语法声明的文章http://blog.go-zh.org/gos-declaration-syntax 了解这种类型声明形式出现的原因。）

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\functions\functions.go"  
55

*当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。*

在本例中，

`x int, y int`

被缩写为

`x, y int`

```go
func add(x, y int) int {
	return x + y
}
```

### 多值返回

函数可以返回任意数量的返回值。

`swap` 函数返回了两个字符串。

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\multiple-results\multiple-results.go"
world hello

### 命名返回值

Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。

返回值的名称应当具有一定的意义，它可以作为文档使用。

没有参数的 `return` 语句返回已命名的返回值。也就是 `直接` 返回。

*直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。*

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9 // Go的除法自动实现四舍五入
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\named-results\named-results.go"
7 10

### 变量

`var` 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。

就像在这个例子中看到的一样，var 语句可以出现在包或函数级别。

```go
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\variables\variables.go"  
0 false false false

### 变量的初始化

变量声明可以包含初始值，每个变量对应一个。

*如果初始化值已存在，则可以省略类型*；变量会从初始值中获得类型。

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\variables-with-initializers\variables-with-initializers.go"
1 2 true false no!

**短变量声明**

在函数中，简洁赋值语句 `:=` 可在类型明确的地方代替 var 声明。

*函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。*

```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\short-variable-declarations\short-variable-declarations.go"
1 2 3 true false no!

### 基本类型

Go 的基本类型有

```GO
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
    // 表示一个 Unicode 码点

float32 float64

complex64 complex128
```

本例展示了几种类型的变量。 *同导入语句一样，变量声明也可以“分组”成一个语法块。*

`int, uint` 和 `uintptr` 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。 当你需要一个整数值时应使用 `int` 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。

```GO
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\basic-types\basic-types.go"
Type: bool Value: false
Type: uint64 Value: 18446744073709551615
Type: complex128 Value: (2+3i)

### 零值

没有明确初始值的变量声明会被赋予它们的 **零值**。

零值是：

- 数值类型为 0，
- 布尔类型为 false，
- 字符串为 ""（空字符串）。

```go
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q", i, f, b, s)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\zero\zero.go"
0 0 false ""

### 类型转换

*表达式 `T(v)` 将值 `v` 转换为类型 `T`。*

一些关于数值的转换：

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

或者，更加简单的形式：

```go
i := 42
f := float64(i)
u := uint(f)
```

*与 C 不同的是，Go 在不同类型的项之间赋值时需要显式转换。*试着移除例子中 float64 或 uint 的转换看看会发生什么。

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Println(f)
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\type-conversions\type-conversions.go"
5
3 4 5

### 类型推导

*在声明一个变量而不指定其类型时（即使用不带类型的 `:=` 语法或 `var =` 表达式语法），变量的类型由右值推导得出。*

当右值声明了类型时，新变量的类型与其相同：

```go
var i int
j := i // j 也是一个 int
```

不过当右边包含未指明类型的数值常量时，新变量的类型就可能是 `int, float64` 或 `complex128` 了，这取决于常量的精度：

```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

尝试修改示例代码中 `v` 的初始值，并观察它是如何影响类型的。

```go
package main

import "fmt"

func main() {
	// v := 42
	// v := 3.142
	v := 0.867 + 0.5i
	fmt.Printf("v is of type %T\n", v)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\type-inference\type-inference.go"
v is of type int
PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\type-inference\type-inference.go"
v is of type float64
PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\type-inference\type-inference.go"
v is of type complex128

### 常量

常量的声明与变量类似，只不过是使用 `const` 关键字。

*常量可以是字符、字符串、布尔值或数值。*

> 常量不能用 := 语法声明。

```go
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\constants\constants.go"
Hello 世界
Happy 3.14 Day
Go rules? true

### 数值常量

数值常量是高精度的 **值**。

*一个未指定类型的常量由上下文来决定其类型。*

再尝试一下输出 `needInt(Big)` 吧。

（int 类型最大可以存储一个 64 位的整数，有时会更小。）

（int 可以存放最大64位的整数，根据平台不同有时会更少。）

```go
package main

import "fmt"

const (
	// 将1左移100位来创建一个非常大的数字，2^100
	// 即这个数的二进制是1后面跟着100个0
	Big = 1 << 100
	// 再往右移99位，即Small = 1 << 1, 或者说Small = 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x*0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	// fmt.Println("Big = %e", Big)
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big)) // overflows int
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\numeric-constants\numeric-constants.go"
21
0.2
1.2676506002282295e+29

## 3 流程控制语句：for、if、else、switch和defer

### for

Go 只有一种循环结构：`for` 循环。

基本的 `for` 循环由三部分组成，它们用分号隔开：

* 初始化语句：在第一次迭代前执行
* 条件表达式：在每次迭代前求值
* 后置语句：在每次迭代的结尾执行

初始化语句通常为一句短变量声明，该变量声明仅在 for 语句的作用域中可见。

一旦条件表达式的布尔值为 false，循环迭代就会终止。

> **注意**：和 C、Java、JavaScript 之类的语言不同，Go 的 for 语句后面的三个构成部分外没有小括号， 大括号 { } 则是必须的。

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\for\for.go"
45

*初始化语句和后置语句是可选的。*

```go
package main

import "fmt"

/*
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
*/

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\for\for.go"
1024

### for 是Go中的 "while"

此时你可以去掉分号，因为 C 的 `while` 在 Go 中叫做 `for`。

```go
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\for-is-gos-while\for-is-gos-while.go"
1024

### 无限循环

如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑。

```go
package main

func main() {
	for {
		
	}
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\forever\forever.go"
exit status 0xc000013a

### if

*Go 的 `if` 语句与 `for` 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的.*

```go
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\if\if.go"        
1.4142135623730951 2i

### if 的简短语句

**同 `for` 一样， `if` 语句可以在条件表达式前执行一个简单的语句。**

*该语句声明的变量作用域仅在 `if` 之内。*

（在最后的 `return` 语句处使用 `v` 看看。）

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\if-with-a-short-statement\if-with-a-short-statement.go"
9 20

### if 和 else

在 `if` 的简短语句中声明的变量同样可以在任何对应的 `else` 块中使用。

（在 `main` 的 `fmt.Println` 调用开始前，两次对 `pow` 的调用均已执行并返回其各自的结果。）

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用v了
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 1, 3),
		pow(3, 3, 20),
	)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\if-and-else\if-and-else.go"
3 >= 3
27 >= 20
9 3 20

> Go语言中`Println`调用函数执行顺序和输出顺序不一致？

### 练习：循环与函数

为了练习函数与循环，我们来实现一个平方根函数：用牛顿法实现平方根函数。

计算机通常使用循环来计算 x 的平方根。从某个猜测的值 z 开始，我们可以根据 z² 与 x 的近似度来调整 z，产生一个更好的猜测：

`z -= (z*z - x) / (2*z)`

重复调整的过程，猜测的结果会越来越精确，得到的答案也会尽可能接近实际的平方根。

在提供的 `func Sqrt` 中实现它。无论输入是什么，对 z 的一个恰当的猜测为 1。 要开始，请重复计算 10 次并随之打印每次的 z 值。观察对于不同的值 x（1、2、3 ...）， 你得到的答案是如何逼近结果的，猜测提升的速度有多快。

> 提示：用类型转换或浮点数语法来声明并初始化一个浮点数值：

```go
z := 1.0
z := float64(1)
```

然后，修改循环条件，使得当值停止改变（或改变非常小）的时候退出循环。观察迭代次数大于还是小于 10。 尝试改变 z 的初始猜测，如 x 或 x/2。你的函数结果与标准库中的 math.Sqrt 接近吗？

（*注：* 如果你对该算法的细节感兴趣，上面的 z² − x 是 z² 到它所要到达的值（即 x）的距离， 除以的 2z 为 z² 的导数，我们*通过 z² 的变化速度来改变 z 的调整量。 这种通用方法叫做牛顿法*。 它对很多函数，特别是平方根而言非常有效。）

```go
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	// z := 1.0
	z := x/2
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Printf("第%v次计算的z值: %v\n", i, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(1))
	fmt.Println("***************")
	fmt.Println(Sqrt(2))
	fmt.Println("***************")
	fmt.Println(Sqrt(3))
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\exercise-loops-and-functions\exercise-loops-and-functions.go"
第1次计算的z值: 1.25
第2次计算的z值: 1.025
第3次计算的z值: 1.0003048780487804
第4次计算的z值: 1.0000000464611474
第5次计算的z值: 1.000000000000001
第6次计算的z值: 1
第7次计算的z值: 1
第8次计算的z值: 1
第9次计算的z值: 1
第10次计算的z值: 1
1
***************
第1次计算的z值: 1.5
第2次计算的z值: 1.4166666666666667
第3次计算的z值: 1.4142156862745099
第4次计算的z值: 1.4142135623746899
第5次计算的z值: 1.4142135623730951
第6次计算的z值: 1.414213562373095
第7次计算的z值: 1.4142135623730951
第8次计算的z值: 1.414213562373095
第9次计算的z值: 1.4142135623730951
第10次计算的z值: 1.414213562373095
1.414213562373095
***************
第1次计算的z值: 1.75
第2次计算的z值: 1.7321428571428572
第3次计算的z值: 1.7320508100147276
第4次计算的z值: 1.7320508075688772
第5次计算的z值: 1.7320508075688774
第6次计算的z值: 1.7320508075688772
第7次计算的z值: 1.7320508075688774
第8次计算的z值: 1.7320508075688772
第9次计算的z值: 1.7320508075688774
第10次计算的z值: 1.7320508075688772
1.7320508075688772

### switch

`switch` 是编写一连串 `if - else` 语句的简便方法。*它运行第一个值等于条件表达式的 case 语句*。

Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过*Go 只运行选定的 case，而非之后所有的 case*。 实际上，*Go 自动提供了在这些语言中每个 case 后面所需的 break 语句*。 除非以 `fallthrough` 语句结束，否则分支会自动终止。 *Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。*

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s. \n", os)
	}
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\switch\switch.go"
Go runs on windows.

### switch的求值顺序

switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。

（例如，

```go
switch i {
case 0:
case f():
}
```

在 i==0 时 f 不会被调用。）

> *注意：* Go 练习场中的时间总是从 2009-11-10 23:00:00 UTC 开始，该值的意义留给读者去发现。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Friday?")
	today := time.Now().Weekday()
	// switch time.Saturday {
	switch time.Friday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\switch-evaluation-order\switch-evaluation-order.go"
When's Friday?
In two days.

### 没有条件的switch

没有条件的 switch 同 switch true 一样。

这种形式能将一长串 if-then-else 写得更加清晰。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\switch-with-no-condition\switch-with-no-condition.go"
Good morning!

### defer

*defer 语句会将函数推迟到外层函数返回之后执行。*

*推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。*

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\defer\defer.go"  
hello
world

### defer 栈

*推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。*

更多关于 defer 语句的信息，请阅读此博文。

```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\defer-multi\defer-multi.go"
counting
done
9
8
7
6
5
4
3
2
1
0

## 4 更多类型：struct、slice和映射

### 指针

*Go 拥有指针。指针保存了**值的内存地址**。*

类型 `*T` 是指向 `T` 类型值的指针。其零值为 `nil`。

`var p *int`

*`&` 操作符会生成一个指向其操作数的指针。*

```go
i := 42
p = &i
```

*`*` 操作符表示指针指向的底层值。*

```go
fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
```

*这也就是通常所说的“间接引用”或“重定向”。*

> 与 C 不同，Go 没有指针运算。
> `&` 生成指针
> `*` 操作指针

```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i // 指向i，*生成一个指针*
	fmt.Println(*p) // 通过指针读取i的值
	*p = 21 // 通过指针设置i的值
	fmt.Println(i) // 查看i的值

	p = &j // 指向j
	*p = *p / 37 // 通过指针对j进行除法运算
	fmt.Println(j) // 查看j的值
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\pointers\pointers.go"

42
21
73

### 结构体

一个结构体（`struct`）就是一组字段（field）。

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\structs\structs.go"
{1 2}

### 结构体字段

*结构体字段使用点号来访问。*

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\struct-fields\struct-fields.go"
4

### 结构体指针

*结构体字段可以通过结构体指针来访问。*

如果我们有一个指向结构体的指针 `p`，那么可以通过 (*p).X 来访问其字段 X。不过这么写太啰嗦了，所以语言也允许我们使用*隐式间接引用*，直接写 `p.X` 就可以。

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v // 指向结构体v
	p.X = 1e9 // 使用结构体指针操作结构体v中的变量X
	fmt.Println(v)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\struct-pointers\struct-pointers.go"
{1000000000 2}

### 结构体文法

*结构体文法通过直接列出字段的值来新分配一个结构体。*

使用 `Name:` 语法可以仅列出部分字段。（字段名的顺序无关。）

*特殊的前缀 `&` 返回一个指向结构体的指针。*

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2} // 创建一个Vertex类型的结构体
	v2 = Vertex{X: 1} // Y:0 被隐式地赋予
	v3 = Vertex{} // X:0 Y:0
	p = &Vertex{1, 2} // 创建一个*Vertex类型地结构体（指针）
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\struct-literals\struct-literals.go"
{1 2} &{1 2} {1 0} {0 0}

### 数组

*类型 `[n]T` 表示拥有 n 个 T 类型的值的数组。*

表达式

```go
var a [10]int
```

会将变量 a 声明为拥有 10 个整数的数组。

*数组的长度是其类型的一部分，因此数组不能改变大小。*这看起来是个限制，不过没关系，Go 提供了更加便利的方式来使用数组。

```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\array\array.go"      
Hello World
[Hello World]
[2 3 5 7 11 13]

### 切片

每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。*在实践中，切片比数组更常用。*

*类型 `[]T` 表示一个元素类型为 T 的切片。*

切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：

```go
a[low : high]
```

*它会选择一个半开区间，包括第一个元素，但排除最后一个元素。*

以下表达式创建了一个切片，它包含 `a` 中下标从 1 到 3 的元素：

```go
a[1:4]
```

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\slices\slices.go"    
[3 5 7]

### 切片就像数组的引用

切片并不存储任何数据，它只是描述了底层数组中的一段。

更改切片的元素会修改其底层数组中对应的元素。

与它共享底层数组的切片都会观测到这些修改。

```go
package main

import "fmt"

func main() {
	names := [4]string {
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names) 
	// 数组指向的其实是内存地址，修改值，相当于修改了对应内存地址
	// 会影响其他变量的结果，切片就像数组的引用
	// 修改切片的值，同时会改变原数组的值
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\slices-pointers\tempCodeRunnerFile.go"
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo]

### 切片文法

*切片文法类似于没有长度的数组文法。*

这是一个数组文法：

```go
[3]bool{true, true, false}
```

下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：

```go
[]bool{true, true, false}
```

```go
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	} {
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\slice-literals\slice-literals.go"
[2 3 5 7 11 13]
[true false true true false true]
[{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]

### 切片的默认行为

在进行切片时，你可以利用它的默认行为来忽略上下界。

*切片下界的默认值为 0，上界则是该切片的长度。*

对于数组

```go
var a [10]int
```

来说，以下切片是等价的：

```go
a[0:10]
a[:10]
a[0:]
a[:]
```

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s) // [3, 5, 7]

	s = s[:2]
	fmt.Println(s) // [2, 3]

	s = s[1:]
	fmt.Println(s) // [3, 5, 7, 11, 13]
	// 没有注意到切片后赋值给了s，用的和被切片数组是一个变量
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\slice-bounds\slice-bounds.go"
[3 5 7]
[3 5]
[5]

> 切片的时候注意赋值，不要使用和元数组相同的变量名称，容易覆盖掉原数组的值

### 切片的长度与容量

切片拥有 **长度 和 容量**。

切片的长度就是它所包含的元素个数。

切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

切片 `s` 的长度和容量可通过表达式 `len(s)` 和 `cap(s)` 来获取。

你可以通过重新切片来扩展一个切片，给它提供足够的容量。试着修改示例程序中的切片操作，向外扩展它的容量，看看会发生什么。

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为0
	s = s[:0]
	printSlice(s) // []

	// 拓展其长度
	s = s[:4]
	printSlice(s) // [2, 3, 5, 7]

	// 舍弃前两个值
	s = s[2:]
	printSlice(s) // [5, 7]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

> 切片改变的是数组的长度，容量不变。

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\slice-len-cap\slice-len-cap.go"
len=6 cap=6 [2 3 5 7 11 13]
len=0 cap=6 []
len=4 cap=6 [2 3 5 7]
len=2 cap=4 [5 7]

### nil 切片

切片的零值是 `nil`。

*nil 切片的长度和容量为 0 且没有底层数组。*

```go
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\nil-slices\nil-slices.go"
[] 0 0
nil!

### 用make创建切片

*切片可以用内建函数 `make` 来创建，这也是你创建动态数组的方式。*

*`make` 函数会分配一个元素为**零值**的数组并返回一个引用了它的切片：*

```go
a := make([]int, 5)  // len(a)=5
```

*要指定它的容量，需向 make 传入第三个参数：*

```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

```go
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)
	
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
s, len(x), cap(x), x)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\making-slices\making-slices.go"
a len=5 cap=5 [0 0 0 0 0]
b len=0 cap=5 []
c len=2 cap=5 [0 0]
d len=3 cap=3 [0 0 0]


### 切片的切片

*切片可包含任何类型，甚至包括其它的切片。*

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 创建一个井字板（经典游戏）
	board := [][]string {
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
		// fmt.Println(len(board))
	}
}
```

> 切片的切片，按照3行输出，每一行都是一个子切片。

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\slices-of-slice\slices-of-slice.go"
X _ X
O _ X
_ _ O

### 向切片追加元素

为切片追加新的元素是种常用的操作，为此 Go 提供了内建的 `append` 函数。内建函数的文档对此函数有详细的介绍。

`func append(s []T, vs ...T) []T`

`append` 的第一个参数 `s` 是一个元素类型为 `T` 的切片，其余类型为 `T` 的值将会追加到该切片的末尾。

*`append` 的结果是一个包含原切片所有元素加上新添加元素的切片。*

当 `s` 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。

（要了解关于切片的更多内容，请阅读文章 Go 切片：用法和本质。）https://blog.go-zh.org/go-slices-usage-and-internals


```go
package main

import (
	"fmt"
)

func main() {
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\append\append.go"
len=0 cap=0 []
len=1 cap=1 [0]        
len=2 cap=2 [0 1]      
len=5 cap=6 [0 1 2 3 4]

### Range

*`for` 循环的 `range` 形式可遍历切片或映射。*

*当使用 `for` 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。*

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\range\range.go"       
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128

可以将下标或值赋予 `_` 来忽略它。

```go
for i, _ := range pow
for _, value := range pow
```

*若你只需要索引，忽略第二个变量即可。*

```go
for i := range pow
```

```go
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
		// fmt.Println(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```

PS D:\gocf> go run "d:\gocf\src\A-Tour-of-Go\range-continued\range-continued.go"
1
2
4
8
16
32
64
128
256
512

### 练习：切片

实现 `Pic`。它应当返回一个长度为 `dy` 的切片，其中每个元素是一个长度为 `dx`，元素类型为 `uint8` 的切片。当你运行此程序时，它会将每个整数解释为灰度值（好吧，其实是蓝度值）并显示它所对应的图像。

图像的选择由你来定。几个有趣的函数包括 `(x+y)/2, x*y, x^y, x*log(y) 和 x%(y+1)`。

（提示：需要使用循环来分配 `[][]uint8` 中的每个 `[]uint8`；请使用 `uint8(intValue)` 在类型之间转换；你可能会用到 `math` 包中的函数。）

```go
package main

import (
	// "math"
	"golang.org/x/tour/pic"
)

/*分析：
* 外层切片的长度为dy；
* 内层切片的长度为dx；
* 内层切片中的每个元素值为(x+y)/2, x*y, x^y, x*log(y) 和 x%(y+1)
* 使用嵌套循环的方式计算颜色值
*/

func Pic(dx, dy int) [][]uint8 {
	// temp := make([][]uint8,dy)
	// for i := 1; i < dy; i++ {
	// 	temp[i] = []uint8
	// 	append(temp, [])

	a := make([][]uint8, dy) // 外层切片
	for x := range a{
		b := make([]uint8, dx) // 内层切片
		for y := range b{
			//b[y] = uint8((x+y)/2) // 给内层里的每一个元素赋值。
			//b[y] = uint8(x*y) // 给内层里的每一个元素赋值。
			//b[y] = uint8(x^y) // 给内层里的每一个元素赋值。
			//b[y] = uint8(float64(x)*math.Log(float64(y))) // 给内层里的每一个元素赋值。
			b[y] = uint8(x%(y+1)) // 给内层里的每一个元素赋值。
		}
		a[x] = b // 给外层切片里的每一个元素赋值
	}
	return a
	}

func main() {
	pic.Show(Pic)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\exercise-slices\exercise-slices.go"
IMAGE:iVBORw0KGgoAAAANSUhEUgAAAQAAAAEACAIAAADTED8xAAACaUlEQVR42uzVMRGAAAzAwLSHf8tgAAf95QVkyVNvNRN50FWBl10V6ABa0AFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIB6ADqEAHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdAA6gBZ0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIBSAcgHYB0ANIB6AAq0AFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgHQA0gFIByAdgA6gAh2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADSAUgHIB2AdADyxy8AAP//YSoDD5pLB7MAAAAASUVORK5CYII=

运行上述代码需要额外的包：

```shell
go get -u golang.org/x/tour
# 进入项目目录

# 生成 go.mod 文件
go mod init

# 拉取依赖，会进行指定性拉取（更新），并不会更新所依赖的其它模块。
go get golang.org/x/tour/pic

```

> **存在的问题，暂时我不会使用Go显示图片，因此，终端出现的是图象编码！**

### 映射

映射将键映射到值。

*映射的零值为 `nil` 。`nil` 映射既没有键，也不能添加键。*

*`make` 函数会返回给定类型的映射，并将其初始化备用。*

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\maps\maps.go"
{40.68433 -74.39967}

### 映射的文法

映射的文法与结构体相似，不过必须有键名。

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex {
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\map-literals\map-literals.go"
map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]

*若顶级类型只是一个类型名，你可以在文法的元素中省略它。*

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\map-literals-continued\map-literals-continued.go"
map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]

### 修改映射

在映射 `m` 中插入或修改元素：

`m[key] = elem`

获取元素：

`elem = m[key]`

删除元素：

`delete(m, key)`

*通过双赋值检测某个键是否存在：*

`elem, ok = m[key]`

*若 key 在 m 中，ok 为 true ；否则，ok 为 false。*

*若 key 不在映射中，那么 elem 是该映射元素类型的零值。*

同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。

> **注 ：若 elem 或 ok 还未声明，你可以使用短变量声明：**

`elem, ok := m[key]`

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\mutating-maps\mutating-maps.go"
The value: 42
The value: 48
The value: 0
The value: 0 Present? false

### 练习：映射

实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。

你会发现 strings.Fields 很有帮助。

> 如："I am a boy ha ha "
> 用strings.Fields()来获取分割的信息，以[]string的形式返回，再分析单词出现的次数

```go
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int) // map对象
	c := strings.Fields(s) // []string
	for _, v := range c {
		m[v] += 1 // 如果v没有在map中，m[v]的值为初始值0
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\exercise-maps\exercise-maps.go"
PASS
 f("I am learning Go!") =
  map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
PASS
 f("The quick brown fox jumped over the lazy dog.") =
  map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
PASS
 f("I ate a donut. Then I ate another donut.") =
  map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
PASS
 f("A man a plan a canal panama.") =
  map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}

### 函数值

函数也是值。它们可以像其它值一样传递。

*函数值可以用作函数的参数或返回值。*

```go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\function-values\function-values.go"
13
5
81

### 函数的闭包

Go 函数可以是一个闭包。*闭包是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。*

例如，函数 `adder` 返回一个闭包。每个闭包都被绑定在其各自的 `sum` 变量上。

```go
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\function-closures\function-closures.go"
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90

### 练习：斐波那契闭包

让我们用函数做些好玩的事情。

实现一个 `fibonacci` 函数，它返回一个函数（闭包），该闭包返回一个斐波纳契数列 `(0, 1, 1, 2, 3, 5, ...)`。

```go
package main

import "fmt"

/*
斐波那契数，亦称之为斐波那契数列（意大利语： Successione di Fibonacci)，
又称黄金分割数列、费波那西数列、费波拿契数、费氏数列，指的是这样一个数列：
0、1、1、2、3、5、8、13、21、……在数学上，斐波纳契数列以如下被以递归的方法定义：
F0=0，F1=1，Fn=Fn-1+Fn-2（n>=2，n∈N*），用文字来说，就是斐波那契数列列由
 0 和 1 开始，之后的斐波那契数列系数就由之前的两数相加。
*/

// // 返回一个“返回int的函数”
// func fibonacci() func(int) int {
// 	var a [10]int
// 	a[0] = 0
// 	a[1] = 1
// 	return func(x int) int {
// 		a[x] = a[x] + a[x-1]
// 		return a[x]
// 	}
// }

func fibonacci() func() int {
	back1, back2 := 0, 1 // 预先定义好前两个值

	return func() int {
		// 记录（back1）的值
		temp := back1

		// 重新赋值（这个就是核心代码）
		back1, back2 = back2, (back1 + back2)

		// 返回temp
		return temp
	}
}

func main() {
	f := fibonacci() // 返回一个闭包函数
	for i := 0; i < 10; i++ { // 检测下前10个值
		fmt.Println(f())
	}
}

```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\exercise-fibonacci-closure\exercise-fibonacci-closure.go"
0
1
1
2
3
5
8
13
21
34

## 5 方法和接口

### 方法

*Go 没有类。不过你可以为结构体类型定义方法。*

方法就是一类带特殊的 **接收者** 参数的函数。

方法接收者在它自己的参数列表内，位于 `func` 关键字和方法名之间。

在此例中，Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\methods\methods.go"
5

### 方法即函数

**记住：方法只是个带接收者参数的函数。**

现在这个 `Abs` 的写法就是个正常的函数，功能并没有什么变化。

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\methods-funcs\methods-funcs.go"
5

你也可以为非结构体类型声明方法。

在此例中，我们看到了一个带 `Abs` 方法的数值类型 `MyFloat`。

*你只能为在同一包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 `int` 之类的内建类型）的接收者声明方法。*

（译注：就是*接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法*。）

```go
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2) // const math.Sqrt2 untyped float = 1.41421
	fmt.Println(f.Abs())
}

```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\methods-continued\methods-continued.go"
1.4142135623730951

### 指针接收者

你可以为指针接收者声明方法。

这意味着对于某类型 `T`，接收者的类型可以用 `*T` 的文法。（此外，`T` 不能是像 `*int` 这样的指针。）

例如，这里为 `*Vertex` 定义了 `Scale` 方法。

*指针接收者的方法可以修改接收者指向的值*（就像 `Scale`在这做的）。由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。

试着移除第 16 行 `Scale` 函数声明中的 `*`，观察此程序的行为如何变化。

若使用值接收者，那么 `Scale` 方法会对原始 `Vertex` 值的副本进行操作。（对于函数的其它参数也是如此。）*`Scale` 方法必须用指针接受者来更改 `main` 函数中声明的 `Vertex` 的值。*

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\methods-pointers\methods-pointers.go"
50

### 指针与函数

现在我们要把 `Abs` 和 `Scale` 方法重写为函数。

同样，我们先试着移除掉第 16 的 *。你能看出为什么程序的行为改变了吗？要怎样做才能让该示例顺利通过编译？

（若你不确定，继续往下看。）

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\methods-pointers-explained\methods-pointers-explained.go"
50

### 方法与指针重定向

*比较前两个程序，你大概会注意到带指针参数的函数必须接受一个指针：*

```go
var v Vertex
ScaleFunc(v, 5)  // 编译错误！
ScaleFunc(&v, 5) // OK
```

*而以指针为接收者的方法被调用时，接收者既能为值又能为指针：*

```go
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

对于语句 `v.Scale(5)`，即便 `v` 是个值而非指针，带指针接收者的方法也能被直接调用。 也就是说，由于 `Scale` 方法有一个指针接收者，为方便起见，Go 会将语句 `v.Scale(5)` 解释为 `(&v).Scale(5)`。

```go
package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\indirection\indirection.go"
{60 80} &{96 72}

**同样的事情也发生在相反的方向。**

*接受一个值作为参数的函数必须接受一个指定类型的值：*

```go
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // 编译错误！
```

*而以值为接收者的方法被调用时，接收者既能为值又能为指针：*

```go
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```

> 这种情况下，方法调用 p.Abs() 会被解释为 (*p).Abs()。

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v)) 
	// 接受一个值作为参数的函数必须接受一个指定类型的值：

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
	// 以值为接收者的方法被调用时，接收者既能为值又能为指针：
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\indirection-values\indirection-values.go"
5
5
5
5

### 选择值或指针作为接收者

使用指针接收者的原因有二：

*首先，方法能够修改其接收者指向的值。*

*其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。*

在本例中，`Scale` 和 `Abs` 接收者的类型为 `*Vertex`，即便 `Abs` 并不需要修改其接收者。

**通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。**（我们会在接下来几页中明白为什么。)

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y *f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\methods-with-pointer-receivers\methods-with-pointer-receivers.go"     
Before scaling: &{X:3 Y:4}, Abs: 5
After scaling: &{X:15 Y:20}, Abs: 25

### 接口

**接口类型** 是由*一组方法签名定义的集合*。

*接口类型的变量可以保存任何实现了这些方法的值。*

> 注意: 示例代码的 22 行存在一个错误。由于 Abs 方法只为 *Vertex （指针类型）定义，因此 Vertex（值类型）并未实现 Abser。

```go
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat 实现了Abser
	a = &v // a *Vertex 实现了Abser

	// 下面一行，v是一个Vertex（而不是*Vertex）
	// 所以没有实现Abser。
	// a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\interfaces\interfaces.go"
5

### 接口与隐式实现

*类型通过实现一个接口的所有方法来实现该接口。*既然无需专门显式声明，也就没有“implements”关键字。

*隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。*

*因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。*

```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型T实现了接口I，但我们无需显式声明此事。

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\interfaces-are-satisfied-implicitly\interfaces-are-satisfied-implicitly.go"
hello

### 接口值

接口也是值。它们可以像其它值一样传递。

*接口值可以用作函数的参数或返回值。*

*在内部，接口值可以看做包含值和具体类型的元组：*

`(value, type)`

*接口值保存了一个具体底层类型的具体值。*

*接口值调用方法时会执行其底层类型的同名方法。*

```go
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型T实现了接口I，但我们无需显式声明此事。
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\interface-values\interface-values.go"
(&{Hello}, *main.T)
Hello
(3.141592653589793, main.F)
3.141592653589793

### 底层值为nil的接口值

即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。

在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）。

> 注意: 保存了 nil 具体值的接口其自身并不为 nil。

```go
package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型T实现了接口I，但我们无需显式声明此事。
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\interface-values-with-nil\interface-values-with-nil.go"
(<nil>, *main.T)
<nil>
(&{hello}, *main.T)
hello

### nil接口值

*nil 接口值既不保存值也不保存具体类型。*

***为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 **具体** 方法的类型。***

```go
package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\nil-interface-values\nil-interface-values.go"
(<nil>, <nil>)
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x0 addr=0x0 pc=0xcbda87]

goroutine 1 [running]:
main.main()
        d:/gocf/src/A-Tour-of-Go/nil-interface-values/nil-interface-values.go:12 +0x67
exit status 2

### 空接口

指定了零个方法的接口值被称为 *空接口：*

`interface{}`

*空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）*

*空接口被用来处理未知类型的值。* 例如，`fmt.Print` 可接受类型为 `interface{}` 的任意数量的参数。

```go
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\empty-interface\empty-interface.go"
(<nil>, <nil>)
(42, int)
(hello, string)

### 类型断言

**类型断言** 提供了访问接口值底层具体值的方式。

`t := i.(T)`

该语句断言接口值 `i` 保存了具体类型 `T`，并将其底层类型为 `T` 的值赋予变量 `t`。

若 `i` 并未保存 `T` 类型的值，该语句就会触发一个恐慌。

为了 **判断** 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。

`t, ok := i.(T)`

若 `i` 保存了一个 `T`，那么 `t` 将会是其底层值，而 `ok` 为 true。

否则，`ok` 将为 `false` 而 `t` 将为 `T` 类型的零值，程序并不会产生恐慌。

> 请注意这种语法和读取一个映射时的相同之处。

```go
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // 报错(panic)
	fmt.Println(f)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\type-assertions\type-assertions.go"
hello
hello true
0 false
panic: interface conversion: interface {} is string, not float64  

goroutine 1 [running]:
main.main()
        d:/gocf/src/A-Tour-of-Go/type-assertions/type-assertions.go:17 +0x165
exit status 2

### 类型选择

**类型选择** 是一种按顺序从几个类型断言中选择分支的结构。

类型选择与一般的 switch 语句相似，不过类型选择中的 case 为类型（而非值）， 它们针对给定接口值所存储的值的类型进行比较。

```go
switch v := i.(type) {
case T:
    // v 的类型为 T
case S:
    // v 的类型为 S
default:
    // 没有匹配，v 与 i 的类型相同
}
```

类型选择中的声明与类型断言 i.(T) 的语法相同，只是具体类型 T 被替换成了关键字 type。

此选择语句判断接口值 i 保存的值类型是 T 还是 S。在 T 或 S 的情况下，变量 v 会分别按 T 或 S 类型保存 i 拥有的值。在默认（即没有匹配）的情况下，变量 v 与 i 的接口类型和值相同。

```go
package main

import "fmt"

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

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\type-switches\type-switches.go"
Twice 21 is 42
"hello" is 5 bytes long
I don't know about type bool!

### Stringer

`fmt` 包中定义的 `Stringer` 是最普遍的接口之一。

```go
type Stringer interface {
    String() string
}
```

*Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。*

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\stringer\stringer.go"
Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)

### 练习：Stringer

通过让 `IPAddr` 类型实现 `fmt.Stringer` 来打印点号分隔的地址。

例如，`IPAddr{1, 2, 3, 4}` 应当打印为 `"1.2.3.4"`。

```go
package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: 给IPAddr 添加一个 "String() string" 方法

// 实现1
// func (ip IPAddr) String() string{
// 	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
// }

// 实现2
func (ip IPAddr) String() string {
	lastIndex := len(ip) - 1
	var s string
	for i, v := range ip {
		s += strconv.Itoa(int(v)) // 数字文字间相互转用strconv类
		if i != lastIndex {
			s += "."
		}
	}
	return s 
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\exercise-stringer\exercise-ise-stringer\exercise-stringer.go"
loopback: 127.0.0.1
googleDNS: 8.8.8.8

### 错误

Go 程序使用 `error` 值来表示错误状态。

与 `fmt.Stringer` 类似，`error` 类型是一个内建接口：

```go
type error interface {
    Error() string
}
```

（与 `fmt.Stringer` 类似，`fmt` 包在打印值时也会满足 `error`。）

通常函数会返回一个 `error` 值，调用的它的代码应当判断这个错误是否等于 `nil` 来进行错误处理。

```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```
`error` 为 nil 时表示成功；非 nil 的 `error` 表示失败。

```go
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\errors\errors.go"
at 2021-11-11 11:23:05.4082596 +0800 CST m=+0.006843501, it didn't 
work

### 练习：错误

从之前的练习中复制 `Sqrt` 函数，修改它使其返回 `error` 值。

`Sqrt` 接受到一个负数时，应当返回一个非 nil 的错误值。复数同样也不被支持。

创建一个新的类型

`type ErrNegativeSqrt float64`

并为其实现

`func (e ErrNegativeSqrt) Error() string`

方法使其拥有 `error` 值，通过 `ErrNegativeSqrt(-2).Error()` 调用该方法应返回 `"cannot Sqrt negative number: -2"`。

**注意**: 在 `Error` 方法内调用 `fmt.Sprint(e)` 会让程序陷入死循环。可以通过先转换 `e` 来避免这个问题：`fmt.Sprint(float64(e))`。这是为什么呢？

修改 `Sqrt` 函数，使其接受一个负数时，返回 `ErrNegativeSqrt` 值。

```go
package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	// z := 1.0
	if x < 0 {
		// return x, ErrNegativeSqrt(x).Error()
		return 0, ErrNegativeSqrt(x) // 注意这里的用法
	}
	z := x / 2
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		// fmt.Printf("第%v次计算的z值: %v\n", i, z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\exercise-errors\exercise-errors.go"
1.414213562373095 <nil>
0 cannot Sqrt negative number: -2

### Reader

*`io` 包指定了 `io.Reader` 接口，它表示从数据流的末尾进行读取。*

*Go 标准库包含了该接口的许多实现，包括文件、网络连接、压缩和加密等等。*

`io.Reader` 接口有一个 `Read `方法：

```go
func (T) Read(b []byte) (n int, err error)
```

`Read` 用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 `io.EOF` 错误。

示例代码创建了一个 `strings.Reader` 并以每次 `8` 字节的速度读取它的输出。

```go
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\reader\reader.go"
n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
b[:n] = "Hello, R"
n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
b[:n] = "eader!"
n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
b[:n] = ""

### 练习：Reader

实现一个 `Reader` 类型，它产生一个 ASCII 字符 `'A'` 的无限流。

`io` 包指定了 `io.Reader` 接口，它表示从数据流的末尾进行读取。

`Read` 用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 `io.EOF` 错误。

```go
package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: 给 MyReader 添加一个Read([]byte) (int, error) 方法
// func Read([]byte) (int, error) {
// 	return fmt.Sprintf("%v", )
// }

func (r MyReader) Read(b []byte) (int, error) {
	// 赋值并返回
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\exercise-reader\exercise-reader.go"
OK!

validate.go

```go
// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reader // import "golang.org/x/tour/reader"

import (
	"fmt"
	"io"
	"os"
)

func Validate(r io.Reader) {
	b := make([]byte, 1024, 2048)
	i, o := 0, 0
	for ; i < 1<<20 && o < 1<<20; i++ { // test 1mb
		n, err := r.Read(b)
		for i, v := range b[:n] {
			if v != 'A' {
				fmt.Fprintf(os.Stderr, "got byte %x at offset %v, want 'A'\n", v, o+i)
				return
			}
		}
		o += n
		if err != nil {
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
			return
		}
	}
	if o == 0 {
		fmt.Fprintf(os.Stderr, "read zero bytes after %d Read calls\n", i)
		return
	}
	fmt.Println("OK!")
}
```

### 练习：rot13Reader

有种常见的模式是一个 `io.Reader` 包装另一个 `io.Reader`，然后通过某种方式修改其数据流。

例如，`gzip.NewReader` 函数接受一个 `io.Reader`（已压缩的数据流）并返回一个同样实现了 `io.Reade`r 的 `*gzip.Reader`（解压后的数据流）。

编写一个实现了 `io.Reader` 并从另一个 `io.Reader` 中读取数据的 `rot13Reader`，通过应用 `rot13` 代换密码对数据流进行修改。

`rot13Reader` 类型已经提供。实现 `Read` 方法以满足 `io.Reader`。

**关键信息**

 ROT13  是过去在古罗马开发的**凯撒加密**的一种变体。

 套用ROT13到一段文字上仅仅只需要检查字元字母顺序并取代它在13位之后的对应字母 ，有需要超过时则重新绕回26英文字母开头即可[2]。A换成N、B换成O、依此类推到M换成Z，然后序列反转：N换成A、O换成B、最后Z换成M。只有这些出现在英文字母里头的字元受影响；数字、符号、空白字元以及所有其他字元都不变。 因为只有在英文字母表里头只有26个，并且26=2×13，ROT13函数是它自己的逆反： 

```go
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// 转换byte，前进13位/后退13位
func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M' || 'a' <= b && b <= 'm':
		b += 13
	case 'M' < b && b <= 'Z' || 'm' < b && b <= 'z':
		b -= 13
	}
	return b
}

// rot13的更精简写法
// func rot13(b byte) byte {
// 	if ('A' <= b && b <= 'M') || ('a' <= b && b <= 'm') {
// 		b += 13
// 	} else if ('N' <= b && b <= 'Z') || ('n' <= b && b <= 'z') {
// 		b -= 13
// 	}
// 	return b
// }

// rot13再精简的写法
// func rot13(b byte) byte {
// 	switch {
// 	case 'A' <= b && b <= 'Z':
// 		b = (b-'A'+13)%26 + 'A'
// 	case 'a' <= b && b <= 'z':
// 		b = (b-'a'+13)%26 + 'a'
// 	}
// }

// 重写Read方法
func (mr rot13Reader) Read(b []byte) (int, error) {
	n, err := mr.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\exercise-rot-reader\exercise-rot-reader.go"
You cracked the code!

### 图像

`image` 包定义了 `Image` 接口：

```go
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

> 注意: `Bounds` 方法的返回值 `Rectangle` 实际上是一个 `image.Rectangle`，它在 `image` 包中声明。

（请参阅文档https://go-zh.org/pkg/image/#Image了解全部信息。）

`color.Color` 和 `color.Model` 类型也是接口，但是通常因为直接使用预定义的实现 `image.RGBA` 和 `image.RGBAModel` 而被忽视了。这些接口和类型由 `image/color` 包定义。

```go
package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\images\images.go"
(0,0)-(100,100)
0 0 0 0

### 练习：图像

还记得之前编写的图片生成器 吗？我们再来编写另外一个，不过这次它将会返回一个 `image.Image` 的实现而非一个数据切片。

定义你自己的 `Image` 类型，实现必要的方法https://go-zh.org/pkg/image/#Image并调用 `pic.ShowImage`。

`Bounds` 应当返回一个 `image.Rectangle` ，例如 `image.Rect(0, 0, w, h)`。

`ColorModel` 应当返回 `color.RGBAModel`。

`At` 应当返回一个颜色。上一个图片生成器的值 `v` 对应于此次的 `color.RGBA{v, v, 255, 255}`。

**关键信息**

```go
// 1 先了解官方的image结构

type Image interface {

	// 颜色模式
    ColorModel() color.Model

	// 图片边界
    Bounds() Rectangle

	// 某个点的颜色
    At(x, y int) color.Color
}

// 2 重点是实现Image的方法，这样就可以自己使用了
```

```go
package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// 1 新建构造体

type Image struct{}

// 2 实现官方image的三个方法

func (i Image) ColorModel() color.Model {
	// 实现Image包中颜色模式的方法
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	// 实现Image包中生成图片边界的方法
	// 这里的200 （宽、高） 我写死了 仅仅是展示作用 正确的做法是从i中获取
	return image.Rect(0, 0, 200, 200)
}

func (i Image) At(x, y int) color.Color {
	// 实现Image包中生成图像某个点的方法
	return color.RGBA{uint8(x), uint8(y), uint8(255), uint8(255)}
}
func main() {
	// 可以自己设置宽高，传递进去
	m := Image{}
	// 3 调用
	pic.ShowImage(m)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\exercise-images\exercise-images.go"
IMAGE:iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAIAAAAiOjnJAAABgElEQVR42uzSUQkAMAxDwQzq3/LGNDSfd4QKKG+SexKz7uYfaBMWwkJYCMsTEBbCQlggLISFsEBYCAthgbAQFsICYSEshAXCQlgIC4SFsBAWCAthISwQFsJCWCAshIWwQFgIC2GBsBAWwgJhISyEBcJCWAgLhIWwEBYIC2EhLBAWwkJYICyEhbBAWAgLYYGwEBbCAmEhLIQFwkJYCAuEhbAQFggLYSEsEBbCQlggLISFsEBYCAthgbAQFsJCWMJCWAgLYYGwEBbCAmEhLIQFwkJYCAuEhbAQFggLYSEsEBbCQlggLISFsEBYCAthgbAQFsICYSEshAXCQlgIC4SFsBAWCAthISwQFsJCWCAshIWwQFgIC2GBsBAWwgJhISyEBcJCWAgLhIWwEBYIC2EhLBAWwkJYICyEhbBAWAgLYYGwEBbCAmEhLIQFwkJYCAthCQthISyEBcJCWAgLhIWwEBYIC2EhLBAWwkJYICyEhbBAWAgLYYGwEBbCgo0XAAD//8rVBa5B2OnCAAAAAElFTkSuQmCC

这个图片base64后的输出结果，要想查看图片展示，要借助这个网站 ，在解码之前，要把IMAGE: 改成data:image/png;base64,
即：

```shell
data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAIAAAAiOjnJAAABgElEQVR42uzSUQkAMAxDwQzq3/LGNDSfd4QKKG+SexKz7uYfaBMWwkJYCMsTEBbCQlggLISFsEBYCAthgbAQFsICYSEshAXCQlgIC4SFsBAWCAthISwQFsJCWCAshIWwQFgIC2GBsBAWwgJhISyEBcJCWAgLhIWwEBYIC2EhLBAWwkJYICyEhbBAWAgLYYGwEBbCAmEhLIQFwkJYCAuEhbAQFggLYSEsEBbCQlggLISFsEBYCAthgbAQFsJCWMJCWAgLYYGwEBbCAmEhLIQFwkJYCAuEhbAQFggLYSEsEBbCQlggLISFsEBYCAthgbAQFsICYSEshAXCQlgIC4SFsBAWCAthISwQFsJCWCAshIWwQFgIC2GBsBAWwgJhISyEBcJCWAgLhIWwEBYIC2EhLBAWwkJYICyEhbBAWAgLYYGwEBbCAmEhLIQFwkJYCAthCQthISyEBcJCWAgLhIWwEBYIC2EhLBAWwkJYICyEhbBAWAgLYYGwEBbCgo0XAAD//8rVBa5B2OnCAAAAAElFTkSuQmCC
```

## 6 并发

### Goroutine

Go 程（goroutine）是由 Go 运行时管理的轻量级线程。

`go f(x, y, z)`

会启动一个新的 Go 程并执行

`f(x, y, z)`

`f, x, y` 和 `z` 的*求值*发生在当前的 Go 程中，而 f 的*执行*发生在新的 Go 程中。

*Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。sync 包提供了这种能力，不过在 Go 中并不经常用到，因为还有其它的办法（见下一页）。*

```go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\goroutines\goroutines.go"
hello
world
world
hello
hello
world
world
hello
hello

### channel

信道是带有类型的管道，你可以通过它用**信道操作符** `<-` 来发送或者接收值。

```go
ch <- v    // 将 v 发送至信道 ch。
v := <-ch  // 从 ch 接收值并赋予 v。
```

（“箭头”就是数据流的方向。）

**和映射与切片一样，信道在使用前必须创建：**

```go
ch := make(chan int)
```

*默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。*

以下示例对切片中的数进行求和，将任务分配给两个 Go 程。一旦两个 Go 程完成了它们的计算，它就能算出最终的结果。

```go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c,<-c // 从c中接收

	fmt.Println(x, y, x+y)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\channels\channels.go"
-5 17 12

### 带缓冲的channel

**信道可以是 *带缓冲* 的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道：**

```go
ch := make(chan int, 100)
```

*仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。*

修改示例填满缓冲区，然后看看会发生什么。

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\buffered-channels\buffered-channels.go"
1
2

### range 和 close

发送者可通过 `close` 关闭一个信道来表示没有需要发送的值了。接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭，那么在执行完

```go
v, ok := <-ch
```

之后 `ok` 会被设置为 `false`。

循环 `for i := range c` 会不断从信道接收值，直到它被关闭。

*注意：* 只有发送者才能关闭信道，而接收者不能。*向一个已经关闭的信道发送数据会引发程序恐慌（panic）。*

*还要注意：* **信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个 range 循环。**

```go
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0;i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c) // cap() 可以用来查看数组或slice的容量
	for i := range c {
		fmt.Println(i)
	}
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\range-for-close\range-for-close.go"
0
1
1
2
3
5
8
13
21
34

### select 语句 ***示例代码没有看懂

*`select` 语句使一个 Go 程可以等待多个通信操作。*

*`select` 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。*

```go
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\select\select.go"
0
1
1
2
3
5
8
13
21
34
quit

### 默认选择

当 `select` 中的其它分支都没有准备好时，`default` 分支就会执行。

为了在尝试发送或者接收时不发生阻塞，可使用 `default` 分支：

```go
select {
case i := <-c:
    // 使用 i
default:
    // 从 c 中接收会阻塞时执行
}
```

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("	.")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\default-selection\default-selection.go"
        .
        .
tick.
        .
        .
tick.
        .
        .
tick.
        .
tick.
        .
        .
BOOM!

### 练习：等价二叉查找树1

不同二叉树的叶节点上可以保存相同的值序列。例如，以下两个二叉树都保存了序列 `1，1，2，3，5，8，13`。

```go
		3
	  /   \
	 1	   8
	/ \   / \
   1   2 5   13

		8
	   / \
	  3   13
	 / \
	1   5
   / \
  1   2
```

在大多数语言中，检查两个二叉树是否保存了相同序列的函数都相当复杂。 我们将使用 Go 的并发和信道来编写一个简单的解法。

本例使用了 tree 包，它定义了类型：

```go
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
```

点击下一页继续。

### 练习：等价二叉查找树2

1. 实现 `Walk` 函数。

2. 测试 `Walk` 函数。

函数 `tree.New(k)` 用于构造一个随机结构的已排序二叉查找树，它保存了值 `k, 2k, 3k, ..., 10k`。

创建一个新的信道 `ch` 并且对其进行步进：

`go Walk(tree.New(1), ch)`

然后从信道中读取并打印 `10` 个值。应当是数字 `1, 2, 3, ..., 10`。

3. 用 `Walk` 实现 `Same` 函数来检测 `t1` 和 `t2` 是否存储了相同的值。

4. 测试 `Same` 函数。

`Same(tree.New(1), tree.New(1))` 应当返回 `true`，而 `Same(tree.New(1), tree.New(2))` 应当返回 `false``。

`Tree` 的文档可在这里找到。https://godoc.org/golang.org/x/tour/tree#Tree

**关键信息**

> 开启子线程，把value发送到channel中，在主线程中接受并处理

```go
package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// 发送value，结束后关闭channel
// Walk 步进 tree t 将所有的值从tree 发送到channel ch
func Walk(t *tree.Tree, ch chan int) {
	snedValue(t, ch)
	close(ch)
}

// 递归向channel传值
func snedValue(t *tree.Tree, ch chan int) {
	if t != nil {
		snedValue(t.Left, ch)
		ch <- t.Value
		snedValue(t.Right, ch)
	}
}

// 使用写好的walk函数来确定两个tree对象 是否一样 原理还是判断value值
// Same 检测树t1 和 t2是否含有相同的值
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <- ch2 {
			return false
		}
	}
	return true
}

func main() {
	// 打印tree.New(1)的值
	var ch = make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}

	// 比较两个tree的value值是否相等
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\exercise-equivalent-binary-trees\exercise-equivalent-binary-trees.go"     
1
2
3
4
5
6
7
8
9
10
true
false

### sync.Mutex

我们已经看到信道非常适合在各个 Go 程间进行通信。

但是如果我们并不需要通信呢？比如说，若我们只是想保证每次只有一个 Go 程能够访问一个共享的变量，从而避免冲突？

这里涉及的概念叫做 *互斥（mutual*exclusion）* ，我们通常使用 *互斥锁（Mutex）* 这一数据结构来提供这种机制。

Go 标准库中提供了 `sync.Mutex` 互斥锁类型及其两个方法：

`Lock`
`Unlock`

*我们可以通过在代码前调用 `Lock` 方法，在代码后调用 `Unlock` 方法来保证一段代码的互斥执行。参见 `Inc` 方法。*

*我们也可以用 `defe`r 语句来保证互斥锁一定会被解锁*。参见 `Value` 方法。

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一个时刻只有一个 goroutine能访问c.v
	c.v[key]++
	c.mux.Unlock()
}

// Value 返回给定 key 的计数器的当前值
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一个时刻只有一个 goroutine能访问c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\mutex-counter\mutex-counter.go"
1000

### 练习：Web爬虫

在这个练习中，我们将会使用 Go 的并发特性来并行化一个 Web 爬虫。

修改 `Crawl` 函数来并行地抓取 URL，并且保证不重复。

提示：你可以用一个 map 来缓存已经获取的 URL，但是要注意 map 本身并不是并发安全的！

### *练习原始代码*

```go
package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
        // 下面并没有实现上面两种情况：
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
```

found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
not found: https://golang.org/cmd/
not found: https://golang.org/cmd/
found: https://golang.org/pkg/fmt/ "Package fmt"
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
found: https://golang.org/pkg/os/ "Package os"
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
not found: https://golang.org/cmd/

### *参考答案1* CSDN-BigManing

```go
package main

import (
	"fmt"
	"sync"
	// "error"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (boday string, urls []string, err error)
}

var (
	// map 存放爬取的url
	m = make(map[string]int)
	// 互斥锁
	l sync.Mutex
	// 群组等待 	当添加的任务没有完成时(done()), wait()会一直等待
	// 三个方法 Add() Done() Wait()
	i sync.WaitGroup
)

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL.
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	defer i.Done() // 和 add相对应
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Printf("found: %s %q\n", url, body)
	// for _, u := range urls {
	// 	Crawl(u, depth-1, fetcher)
	// }
	// return

	// 存入数据 需要同步锁 因为这是在子线程中
	l.Lock()
	if m[url] == 0 { // 还未爬取过
		m[url]++ // 存入爬取的url 改变对应的标示
		depth--
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			i.Add(1)
			go Crawl(u, depth, fetcher) // 继续爬取
		}
	}
	l.Unlock()
}

func main() {
	i.Add(1)
	Crawl("https://golang.org/", 4, fetcher)

	i.Wait() // 会一直等待直到子线程任务结束

	for k, _ := range m {
		fmt.Println(k)
	}
	fmt.Println("over")
}

//
// 以下是练习源代码
//

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg",
		},
	},
}
```

PS D:\gocf\src\A-Tour-of-Go> go run "d:\gocf\src\A-Tour-of-Go\exercise-web-crawler\exercise-web-crawler.go"
found: https://golang.org/ "The Go Programming Language"
not found: https://golang.org/cmd
found: https://golang.org/pkg/ "Packages"     
found: https://golang.org/pkg/os/ "Package os"
not found: https://golang.org/pkg/fmt/        
not found: https://golang.org/cmd
not found: https://golang.org/pkg
https://golang.org/
https://golang.org/pkg/
https://golang.org/pkg/os/
over

### *参考答案2* CSDN-帐前卒

```go
package main

import (
  "fmt"
)

type Fetcher interface {
        // Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
  Fetch(url string) (body string, urls []string, err error)
}
var lockx = make(chan int,1)
// 同步通信使用
func LockFun(f func()) {
  lockx<-1
  f()
  <-lockx
}
var visited map[string]bool = make(map[string]bool)
// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, banner chan int) {
 
  if depth <= 0 || visited[url] {
    banner<-1
    return
  }
  body, urls, err := fetcher.Fetch(url)
  LockFun(func(){
    visited[url]=true
  })
  fmt.Printf("found: %s %q\n", url, body)
  if err != nil {
    fmt.Println(err)
    banner<-1
    return
  }
  subBanner := make(chan int, len(urls))
  for _, u := range urls {
     // 并行吧～～ 
      go Crawl(u, depth-1, fetcher, subBanner);
  }
  for i:=0; i < len(urls); i++ {
    // subBanner用来防止退出
    <-subBanner
  }
  // banner用于让父节点退出
  banner<-1
  return
}

func main() {
  mainBanner := make(chan int,1)
  Crawl("http://golang.org/", 4, fetcher, mainBanner)
  <-mainBanner
}


// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
  body string
  urls     []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
  if res, ok := (*f)[url]; ok {
    return res.body, res.urls, nil
  }
  return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = &fakeFetcher{
  "http://golang.org/": &fakeResult{
    "The Go Programming Language",
    []string{
      "http://golang.org/pkg/",
      "http://golang.org/cmd/",
    },
  },
  "http://golang.org/pkg/": &fakeResult{
    "Packages",
    []string{
      "http://golang.org/",
      "http://golang.org/cmd/",
      "http://golang.org/pkg/fmt/",
      "http://golang.org/pkg/os/",
    },
  },
  "http://golang.org/pkg/fmt/": &fakeResult{
    "Package fmt",
    []string{
      "http://golang.org/",
      "http://golang.org/pkg/",
    },
  },
  "http://golang.org/pkg/os/": &fakeResult{
    "Package os",
    []string{
      "http://golang.org/",
      "http://golang.org/pkg/",
    },
  },
}
```

found: http://golang.org/ "The Go Programming Language"
found: http://golang.org/cmd/ ""
not found: http://golang.org/cmd/
found: http://golang.org/pkg/ "Packages"
found: http://golang.org/pkg/os/ "Package os"
found: http://golang.org/pkg/fmt/ "Package fmt"

### *参考答案3* 

```go
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// -------------下面是我进行了修改的部分------------

// 创建map用来存放已抓取的网页, wg组用来管理线程
var m map[string]bool
var wg sync.WaitGroup

// TODO: 并行的抓取 URL。
// TODO: 不重复抓取页面。
// 这里用了辅助函数_crawl来完成上面的任务，并进行主要的抓取工作
func _crawl(url string, depth int, fetcher Fetcher, Results chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth <= 0 {
		return
	}

	if exists := m[url]; exists {
		return
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		Results <- fmt.Sprintf("not found: %s", url)
		return
	}
	m[url] = true

	Results <- fmt.Sprintf("found: %s %q", url, body)

	for _, u := range urls {
		wg.Add(1)
		go _crawl(u, depth-1, fetcher, Results, wg)
	}
	return
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
// 这里Crawl的主要作用是管理线程，并从信道中读取数据
func Crawl(url string, depth int, fetcher Fetcher) {
	Results := make(chan string)
	wg.Add(1)
	go _crawl(url, depth, fetcher, Results, &wg)
	go func() {
		wg.Wait()
		close(Results)
	}()
	for i := range Results {
		fmt.Println(i)
	}
}

func main() {
	m = make(map[string]bool)
	Crawl("https://golang.org/", 4, fetcher)
}

// -------------下面是练习本身的代码--------------

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
```

found: https://golang.org/ "The Go Programming Language"
not found: https://golang.org/cmd/
found: https://golang.org/pkg/ "Packages"
found: https://golang.org/pkg/os/ "Package os"
not found: https://golang.org/cmd/
found: https://golang.org/pkg/fmt/ "Package fmt"

### Go Tour 官方代码

```go
// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"errors"
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// fetched tracks URLs that have been (or are being) fetched.
// The lock must be held while reading from or writing to the map.
// See https://golang.org/ref/spec#Struct_types section on embedded types.
var fetched = struct {
	m map[string]error
	sync.Mutex
}{m: make(map[string]error)}

var loading = errors.New("url load in progress") // sentinel value

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		fmt.Printf("<- Done with %v, depth 0.\n", url)
		return
	}

	fetched.Lock()
	if _, ok := fetched.m[url]; ok {
		fetched.Unlock()
		fmt.Printf("<- Done with %v, already fetched.\n", url)
		return
	}
	// We mark the url to be loading to avoid others reloading it at the same time.
	fetched.m[url] = loading
	fetched.Unlock()

	// We load it concurrently.
	body, urls, err := fetcher.Fetch(url)

	// And update the status in a synced zone.
	fetched.Lock()
	fetched.m[url] = err
	fetched.Unlock()

	if err != nil {
		fmt.Printf("<- Error on %v: %v\n", url, err)
		return
	}
	fmt.Printf("Found: %s %q\n", url, body)
	done := make(chan bool)
	for i, u := range urls {
		fmt.Printf("-> Crawling child %v/%v of %v : %v.\n", i, len(urls), url, u)
		go func(url string) {
			Crawl(url, depth-1, fetcher)
			done <- true
		}(u)
	}
	for i, u := range urls {
		fmt.Printf("<- [%v] %v/%v Waiting for child %v.\n", url, i, len(urls), u)
		<-done
	}
	fmt.Printf("<- Done with %v\n", url)
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)

	fmt.Println("Fetching stats\n--------------")
	for url, err := range fetched.m {
		if err != nil {
			fmt.Printf("%v failed: %v\n", url, err)
		} else {
			fmt.Printf("%v was fetched\n", url)
		}
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
```


Found: https://golang.org/ "The Go Programming Language"
-> Crawling child 0/2 of https://golang.org/ : https://golang.org/pkg/.
-> Crawling child 1/2 of https://golang.org/ : https://golang.org/cmd/.
<- [https://golang.org/] 0/2 Waiting for child https://golang.org/pkg/.
<- Error on https://golang.org/cmd/: not found: https://golang.org/cmd/
<- [https://golang.org/] 1/2 Waiting for child https://golang.org/cmd/.
Found: https://golang.org/pkg/ "Packages"
-> Crawling child 0/4 of https://golang.org/pkg/ : https://golang.org/.
-> Crawling child 1/4 of https://golang.org/pkg/ : https://golang.org/cmd/.
-> Crawling child 2/4 of https://golang.org/pkg/ : https://golang.org/pkg/fmt/.
-> Crawling child 3/4 of https://golang.org/pkg/ : https://golang.org/pkg/os/.
<- [https://golang.org/pkg/] 0/4 Waiting for child https://golang.org/.
Found: https://golang.org/pkg/os/ "Package os"
-> Crawling child 0/2 of https://golang.org/pkg/os/ : https://golang.org/.
-> Crawling child 1/2 of https://golang.org/pkg/os/ : https://golang.org/pkg/.
<- Done with https://golang.org/, already fetched.
Found: https://golang.org/pkg/fmt/ "Package fmt"
-> Crawling child 0/2 of https://golang.org/pkg/fmt/ : https://golang.org/.
-> Crawling child 1/2 of https://golang.org/pkg/fmt/ : https://golang.org/pkg/.
<- [https://golang.org/pkg/fmt/] 0/2 Waiting for child https://golang.org/.
<- Done with https://golang.org/cmd/, already fetched.
<- [https://golang.org/pkg/] 1/4 Waiting for child https://golang.org/cmd/.
<- [https://golang.org/pkg/os/] 0/2 Waiting for child https://golang.org/.
<- [https://golang.org/pkg/os/] 1/2 Waiting for child https://golang.org/pkg/.
<- Done with https://golang.org/pkg/, already fetched.
<- [https://golang.org/pkg/fmt/] 1/2 Waiting for child https://golang.org/pkg/.
<- Done with https://golang.org/, already fetched.
<- Done with https://golang.org/pkg/fmt/
<- [https://golang.org/pkg/] 2/4 Waiting for child https://golang.org/pkg/fmt/.
<- Done with https://golang.org/, already fetched.
<- [https://golang.org/pkg/] 3/4 Waiting for child https://golang.org/pkg/os/.
<- Done with https://golang.org/pkg/, already fetched.
<- Done with https://golang.org/pkg/os/
<- Done with https://golang.org/pkg/
<- Done with https://golang.org/
Fetching stats
--------------
https://golang.org/ was fetched
https://golang.org/pkg/ was fetched
https://golang.org/cmd/ failed: not found: https://golang.org/cmd/
https://golang.org/pkg/os/ was fetched
https://golang.org/pkg/fmt/ was fetched

### 爬虫总结

对比Go Tour官方代码,只有参考答案3对得上

* 参考答案1少爬取了网站: https://golang.org/pkg/fmt/
* 参考答案2对网站https://golang.org/cmd/ 爬取结果不正确
* 参考答案3和标准答案吻合

## 接下来去哪里?

你可以从[安装 Go](https://go-zh.org/doc/install/) 开始。

一旦安装了 Go，Go [文档](https://go-zh.org/doc/)是一个极好的 应当继续阅读的内容。 它包含了参考、指南、视频等等更多资料。

了解如何组织 Go 代码并在其上工作，参阅[此视频](https://www.youtube.com/watch?v=XCsL89YtqCs)，或者阅读[如何编写 Go 代码](https://go-zh.org/doc/code.html)。

如果你需要标准库方面的帮助，请参考[包手册](https://go-zh.org/pkg/)。如果是语言本身的帮助，阅读[语言规范](https://go-zh.org/ref/spec)是件令人愉快的事情。

进一步探索 Go 的并发模型，参阅 [Go 并发模型(幻灯片)](https://www.youtube.com/watch?v=f6kdp27TYZs)以及[深入 Go 并发模型(幻灯片)](https://www.youtube.com/watch?v=QDDwwePbDtw)并阅读[通过通信共享内存](https://go-zh.org/doc/codewalk/sharemem/)的代码之旅。

想要开始编写 Web 应用，请参阅[一个简单的编程环境(幻灯片)](https://vimeo.com/53221558)并阅读[编写 Web 应用](https://go-zh.org/doc/articles/wiki/)的指南。

[函数：Go 中的一等公民](https://go-zh.org/doc/codewalk/functions/)展示了有趣的函数类型。

[Go 博客](https://blog.go-zh.org/)有着众多关于 Go 的文章和信息。

[mikespook 的博客](https://www.mikespook.com/tag/golang/)中有大量中文的关于 Go 的文章和翻译。

开源电子书 [Go Web 编程](https://github.com/astaxie/build-web-application-with-golang)和 [Go 入门指南](https://github.com/Unknwon/the-way-to-go_ZH_CN)能够帮助你更加深入的了解和学习 Go 语言。

访问 [go-zh.org](https://go-zh.org/) 了解更多内容。