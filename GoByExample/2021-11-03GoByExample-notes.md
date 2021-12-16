# 2021-11-03GoByExample-notes

[toc]

## 更新记录

- [x] 2021-11-03 1-Hello World
- [x] 2021-11-04 中午2+3
- [x] 2021-11-06~23 4~20
- [x] 2021-11-24 21
- [x] 2021-11-25 22~26
- [x] 2021-11-26 27~40 有一些没有完全消化，需要对照相应资料深入理解和思考
- [x] 2021-12-13 41~53
- [x] 2021-12-14 54~60 

## Readme

Go by Example 是一个通过带注释的示例程序学习 Go 语言的网站。网站包含了从简单的 Hello World 到高级特性 Goroutine、Channel 等一系列示例程序，并附带了注释说明，非常适合 Go 语言初学者。

如果您想学习 Go 语言基础知识，不要犹豫，请直接前往 Go by Example 开始学习！

## notes

1. 指针

\* 解引用指针，从对应地址获取值；

& 取得对象的内存地址，即指向对象的指针。

2. Go 计算程序运行时间

**计算代码块的运行时间**

> 其中time.Since()函数返回字符串类型，例如1h2m3s等，可能还有us等.

```go
start := time.Now()
//some func or operation
cost := time.Since(start)
fmt.Printf("cost=[%s]",cost) 
```

3. Go select使用

Go中的select和channel配合使用，通过select可以监听多个channel的I/O读写事件，当 IO操作发生时，触发相应的动作。

**基本用法**

```go
//select基本用法
select {
case <- chan1:
// 如果chan1成功读到数据，则进行该case处理语句
case chan2 <- 1:
// 如果成功向chan2写入数据，则进行该case处理语句
default:
// 如果上面都没有成功，则进入default处理流程
```

**使用规则**

1. 如果没有default分支,select会阻塞在多个channel上，对多个channel的读/写事件进行监控。
2. 如果有一个或多个IO操作可以完成，则Go运行时系统会随机的选择一个执行，否则的话，如果有default分支，则执行default分支语句，如果连default都没有，则select语句会一直阻塞，直到至少有一个IO操作可以进行。　　　

**计算函数的运行时间**

> 利用defer的作用，可以在函数开始的时候获取一个时间，使用time.Now()在函数结束的时候，获取程序从标记开始的时间段，可以得到函数运行的时间。

```go
func compute() {
    start := time.Now()
    defer func() {
        cost := time.Since(start)
        fmt.Println("cost=", cost)
    }()
    // some computation
}
```

## 1-Hello World

我们的第一个程序将打印传说中的“hello world”， 右边是完整的程序代码。

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

要运行这个程序，先将将代码放到名为 hello-world.go 的文件中，然后执行 go run。

```shell
PS C:\Users\chenfengyuan\Coding-cf\Go\GoByExample\hello-world> go run hello-world.go
hello world
```

如果我们想将程序编译成二进制文件（Windows 平台是 .exe 可执行文件）， 可以通过 go build 来达到目的。

然后我们可以直接运行这个二进制文件。

现在我们可以运行和编译基础的 Go 程序了， 让我们开始学习更多关于这门语言的知识吧。

```shell
PS C:\Users\chenfengyuan\Coding-cf\Go\GoByExample\hello-world> go build hello-world.go
PS C:\Users\chenfengyuan\Coding-cf\Go\GoByExample\hello-world> ls
    目录: C:\Users\chenfengyuan\Coding-cf\Go\GoByExample\hello-world

Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
-a----         2021/11/4      0:12        1924608 hello-world.exe
-a----         2021/11/4      0:09             77 hello-world.go
PS C:\Users\chenfengyuan\Coding-cf\Go\GoByExample\hello-world> ./hello-world
hello world
```

## 2-值

Go 拥有多种值类型，包括字符串、整型、浮点型、布尔型等。 下面是一些基础的例子。

字符串可以通过 + 连接。

整数和浮点数

布尔型，以及常见的布尔操作。

```go
package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
```

PS C:\Us> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\values\values.go"
golang
1+1 =  2
7.0/3.0 = 2.3333333333333335
false
true
false

## 3-变量

在 Go 中，*变量* 需要显式声明，并且在函数调用等情况下， 编译器会检查其类型的正确性。

```go
package main

import "fmt"

func main() {
	// var 声明 1 个或者多个变量。
	var a = "initial"
	fmt.Println(a)

	// 你可以一次性声明多个变量。
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go 会自动推断已经有初始值的变量的类型。
	var d = true
	fmt.Println(d)

	// 声明后却没有给出对应的初始值时，变量将会初始化为 零值 。
	// 例如，int 的零值是 0。
	var e int
	fmt.Println(e)

	// := 语法是声明并初始化变量的简写， 
	// 例如 var f string = "short" 
	// 可以简写为右边这样。
	f := "short"
	fmt.Println(f)
}
```

PS C:\Users\chenfengyuan\Coding-cf\Go\GoByExample\hello-world> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\variables\variables.go"
initial
1 2
true
0
short


## 4-常量

Go 支持字符、字符串、布尔和数值 **常量** 。

```go
package main

import (
	"fmt"
	"math"
)

// `const`用于声明一个常量。
const s string = "constant"

func main() {
	fmt.Println(s)

	// `const`语句可以出现在任何var语句可以出现的地方
	const n = 500000000

	// 常数表达式可以执行任意精度的运算
	const d = 3e20 / n
	fmt.Println(d)

	// 数值型常量没有确定的类型，直到被给定某个类型，比如显式类型转化。
	fmt.Println(int64(d))

	// 一个数字可以根据上下文的需要（比如变量赋值、函数调用）自动确定类型。 
	// 举个例子，这里的 math.Sin 函数需要一个 float64 的参数，n 会自动确定类型。
	fmt.Println(math.Sin(n))
}
```

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\constant\constant.go"        
constant
6e+11
600000000000
-0.28470407323754404

## 5-For循环

for 是 Go 中唯一的循环结构。这里会展示 for 循环的三种基本使用方式。

```go
package main

import "fmt"

func main() {
	
	// 最基础的方式，单个循环条件
	i := 1
	for i <= 3 {
		fmt.Println(i)
		 i = i + 1
	}

	// 经典的初始/条件/后续 for循环
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 不带条件的for循环将一直重复执行，直到在循环体内
	// 使用了break或者return跳出循环
	for {
		fmt.Println("loop")
		break
	}

	// 你也可以使用continue直接进入下一次循环
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
```

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\for\for.go"    
1
2
3
7
8
9
loop
1
3
5

我们在后续教程中学习 range 语句，channels 以及其他数据结构时， 还会看到一些 for 的其它用法。

## 6-If/Else分支

if 和 else 分支结构在 Go 中非常直接。

```go
package main

import "fmt"

func main() {
	
	// 这里是一个基本的例子
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 你可以不要else只用if语句
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 在条件语句之前可以有一个声明语句；在这里声明的变量可以在
	// 这个语句所有的条件分支中使用
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
```

> 注意，在 Go 中，条件语句的圆括号不是必需的，但是花括号是必需的。

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\if-else\if-else.go"
7 is odd
8 is divisible by 4
9 has 1 digit

*Go 没有三目运算符， 即使是基本的条件判断，依然需要使用完整的 if 语句。*

## 7-Switch分支结构

switch 是多分支情况时快捷的条件语句。

> 突然发现一个Bug，更新Vscode后，好像不能正常检查Go代码了，这就很无语呀。
> 但是终究不是影响运行的大问题，忍忍吧。😂

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("write ", i, " as ")
	// 一个基本的switch
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	// 在同一个 case 语句中，你可以使用逗号来分隔多个表达式。 
	// 在这个例子中，我们还使用了可选的 default 分支。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// 带表达式的 switch 是实现 if/else 逻辑的另一种方式。 
	// 这里还展示了 case 表达式也可以不使用常量。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// 类型开关 (type switch) 比较类型而非值。可以用来发现一个接口值的类型。 
	// 在这个例子中，变量 t 在每个分支中会有相应的类型。
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
```

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\switch\switch.go"
write 2 as two
It's a weekday
It's after noon
I'm a bool     
I'm an int
Don't know type string

## 8-数组

在 Go 中，**数组** 是一个具***有编号且长度固定***的元素序列。

```go
package main

import "fmt"

func main() {
	var a[5]int
	fmt.Println("emp:", a)
	// 这里我们创建了一个刚好可以存放 5 个 int 元素的数组 a。 
	// 元素的类型和长度都是数组类型的一部分。 
	// 数组默认值是零值，对于 int 数组来说，元素的零值是 0。

	// 我们可以使用 array[index] = value 语法来设置数组指定位置的值， 
	// 或者用 array[index] 得到值。
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 内置函数 len 可以返回数组的长度。
	fmt.Println("len:", len(a))

	// 使用这个语法在一行内声明并初始化一个数组。
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 数组类型是一维的，但是你可以组合构造多维的数据结构。
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
```

> 注意，使用 fmt.Println 打印数组时，会按照 [v1 v2 v3 ...] 的格式打印。

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\arrays\arrays.go"
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d: [[0 1 2] [1 2 3]]

## 9-切片

在 Go 程序中，相较于数组，用得更多的是 _切片(slice)_。我们接着来看切片。

Slice 是 Go 中一个重要的数据类型，它提供了比数组更强大的序列交互方式。

```go
package main

import "fmt"

func main() {

	// 与数组不同，slice 的类型仅由它所包含的元素的类型决定（与元素个数无关）。 
	// 要创建一个长度不为 0 的空 slice，需要使用内建函数 make。 
	// 这里我们创建了一个长度为 3 的 string 类型的 slice（初始值为零值）。
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 我们可以和数组一样设置和得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len 返回 slice 的长度
	fmt.Println("len:", len(s))

	// 除了基本操作外，slice 支持比数组更丰富的操作。
	// 比如 slice 支持内建函数 append， 该函数会返回一个包含了一个或者多个新值的 slice。 
	// 注意由于 append 可能返回一个新的 slice，我们需要接收其返回值。
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// slice 还可以 copy。
	// 这里我们创建一个空的和 s 有相同长度的 slice——c， 然后将 s 复制给 c。
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice 支持通过 slice[low:high] 语法进行“切片”操作。 
	// 例如，右边的操作可以得到一个包含元素 s[2]、s[3] 和 s[4] 的 slice。
	l1 := s[2:5] // 含左不含右
	fmt.Println("sl1:", l1)

	// 这个 slice 包含从 s[0] 到 s[5]（不包含 5）的元素。
	l2 := s[:5]
	fmt.Println("sl2:", l2)

	// 这个 slice 包含从 s[2]（包含 2）之后的元素。
	l3 := s[2:]
	fmt.Println("sl3:", l3)

	// 我们可以在一行代码中声明并初始化一个 slice 变量。
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slice 可以组成多维数据结构。内部的 slice 长度可以不一致，这一点和多维数组不同。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
```

> 注意，slice 和数组是不同的类型，但它们通过 fmt.Println 打印的输出结果是类似的。

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\slices\slices.go"
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d: [[0] [1 2] [2 3 4]]

看看这个由 Go 团队撰写的一篇很棒的博文http://blog.golang.org/2011/01/go-slices-usage-and-internals.html ，了解更多关于 Go 中 slice 的设计和实现细节。

## 10-Map

现在，我们已经学习了数组和 slice，接下来我们将学习 Go 中的另一个重要的内建数据类型：map。

**map** 是 Go 内建的关联数据类型 （在一些其他的语言中也被称为 **哈希(hash)** 或者 **字典(dict)** ）。

要创建一个空 map，需要使用内建函数 make：`make(map[key-type]val-type)`。

使用典型的 `name[key] = val` 语法来设置键值对。

打印 map。例如，使用 `fmt.Println` 打印一个 map，会输出它所有的键值对。

使用 `name[key]` 来获取一个键的值。

内建函数 `len` 可以返回一个 map 的键值对数量。

内建函数 `delete` 可以从一个 map 中移除键值对。

当从一个 map 中取值时，还有可以选择是否接收的第二个返回值，该值表明了 map 中是否存在这个键。 这可以用来消除 *键不存在* 和 *键的值为零值* 产生的歧义， 例如 0 和 ""。这里我们不需要值，所以用 *空白标识符(blank identifier) `_`*将其忽略

你也可以通过右边的语法在一行代码中声明并初始化一个新的 map。

注意，使用 `fmt.Println` 打印一个 map 的时候， 是以 `map[k:v k:v]` 的格式输出的。

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n:=map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
```

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\maps\maps.go"
map: map[k1:7 k2:13]
v1: 7
len: 2
map: map[k1:7]
prs: false
map: map[bar:2 foo:1]

## 11-Range遍历

`range` 用于迭代各种各样的数据结构。 让我们来看看如何在我们已经学过的数据结构上使用 `range`。

这里我们使用 `range` 来对 `slice` 中的元素求和。 *数组也可以用这种方法初始化并赋初值*。

`range` 在数组和 `slice` 中提供对每项的索引和值的访问。 上面我们不需要索引，所以我们使用 空白标识符 _ 将其忽略。 实际上，我们有时候是需要这个索引的。

range 在 map 中迭代键值对。

range 也可以只遍历 map 的键。

range 在字符串中迭代 unicode 码点(code point)。 第一个返回值是字符的起始字节位置，然后第二个是字符本身。

```go
package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
```

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\range\range.go"
sum: 9
index: 1
a -> apple
b -> banana
key: a
key: b
0 103
1 111

## 12-函数

**函数** 是 Go 的核心。我们将通过一些不同的例子来进行学习它。

这里是一个函数，接受两个 int 并且以 int 返回它们的和

Go 需要明确的 return，也就是说，它不会自动 return 最后一个表达式的值

当多个连续的参数为同样类型时， 可以仅声明最后一个参数的类型，忽略之前相同类型参数的类型声明。

通过 *函数名(参数列表)* 来调用函数，

Go 函数还有很多其他的特性。 其中一个就是多值返回，它也是我们接下来要接触的。

```go
package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
```

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\functions\functions.go"    
1+2 =  3
1+2+3 = 6

## 13-多返回值

Go 原生支持 _多返回值_。 这个特性在 Go 语言中经常用到，*例如用来同时返回一个函数的结果和错误信息。*

(int, int) 在这个函数中标志着这个函数返回 2 个 int。

这里我们通过 **多赋值** 操作来使用这两个不同的返回值。

如果你仅仅需要返回值的一部分的话，你可以使用空白标识符 _。

```go
package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
```

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\multiple-return-values\multiple-return-values.go"
3
7
7

## 14-变参函数

我们接下来要学习的是 Go 函数另一个很好的特性：变参函数。

**可变参数函数**。 在调用时可以传递任意数量的参数。 例如，`fmt.Println` 就是一个常见的变参函数。

```go
package main

import "fmt"

//这个函数接受任意数量的 `int` 作为参数。
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
    // 变参函数使用常规的调用方式，传入独立的参数。
	sum(1, 2)
	sum(1, 2, 3)

    //如果你有一个含有多个值的 `slice`，想把它们作为参数使用， 你需要这样调用 `func(slice...)`。
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
```

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\variadic-functions\variadic-functions.go"
[1 2] 3
[1 2 3] 6
[1 2 3 4] 10

## 15-闭包

接下来我们要看的是 Go 函数的另一个关键特性：闭包。

Go 支持**匿名函数**， 并能用其构造 **闭包**。 *匿名函数在你想定义一个不需要命名的内联函数时是很实用的。*

`intSeq` 函数返回一个在其函数体内定义的匿名函数。 返回的函数使用闭包的方式 **隐藏**变量 `i`。 返回的函数 **隐藏** 变量 `i` 以形成闭包。

我们调用 `intSeq` 函数，将返回值（一个函数）赋给 `nextInt`。 这个函数的值包含了自己的值 `i`，这样在每次调用 `nextInt` 时，都会更新 `i` 的值。

通过多次调用 nextInt 来看看闭包的效果。

为了确认这个状态对于这个特定的函数是唯一的，我们重新创建并测试一下。

```go
package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
```

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\closures\closures.go"      
1
2
3
1

## 16-递归

我们马上要学习关于函数的最后一个特性：递归。

Go 支持 **递归**。 这里是一个经典的阶乘示例。

`fact` 函数在到达 `fact(0)` 前一直调用自身。

```go
package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
```

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\recursion\recursion.go"    
5040

## 17-指针

Go 支持 **指针**， 允许在程序中通过 `引用传递` 来传递值和数据结构。

我们将通过两个函数：`zeroval` 和 `zeroptr` 来比较 `指针` 和 `值`。 `zeroval` 有一个 `int` 型参数，所以使用值传递。 `zeroval` 将从调用它的那个函数中得到一个**实参的拷贝**：`ival`。

`zeroptr` 有一个和上面不同的参数：`*int`，这意味着它使用了 `int` 指针。 紧接着，函数体内的 `*iptr` 会 **解引用** 这个指针，*从它的内存地址得到这个地址当前对应的值*。 **对解引用的指针赋值，会改变这个指针引用的真实地址的值。**

通过 `&i` 语法来取得 `i` 的内存地址，即指向 `i` 的指针。

指针也是可以被打印的。

`zeroval` 在 `main` 函数中不能改变 `i` 的值， 但是 `zeroptr` 可以，*因为它有这个变量的内存地址的引用*。

```go
package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
```

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\pointers\pointers.go"      
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0xc000012078

## 18-结构体

Go 的**结构体(struct)** 是带类型的字段(fields)集合。 这在组织数据时非常有用。

这里的 `person` 结构体包含了 `name` 和 `age` 两个字段。

使用这个语法创建新的结构体元素。

你可以在初始化一个结构体元素时指定字段名字。

省略的字段将被初始化为零值。

```go
package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	// `&` 前缀生成一个结构体指针。
	fmt.Println(&person{name: "Ann", age: 40})

	// 使用`.`来访问结构体字段。
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 也可以对结构体指针使用`.` - 指针会被自动解引用。
	sp := &s
	fmt.Println(sp.age)

	// 结构体是可变(mutable)的
	sp.age = 51
	fmt.Println(sp.age)
}
```

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\structs\structs.go"        
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
Sean
50
51

## 19-方法

Go 支持为结构体类型定义**方法(methods)** 。

这里的 `area` 是一个拥有 **`*rect` 类型接收器(receiver)**的方法。

可以为值类型或者指针类型的接收者定义方法。 这是一个值类型接收者的例子。

这里我们调用上面为结构体定义的两个方法。

调用方法时，Go 会自动处理值和指针之间的转换。 *想要避免在调用方法时产生一个拷贝，或者想让方法可以修改接受结构体的值， 你都可以使用指针来调用方法。*

```go
package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r // & 取得 r 的地址
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
```

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\methods\methods.go"        
area:  50
perim:  30
area:  50
perim:  30

## 20-接口

接下来，我们将学习 Go 对方法集进行命名和分组的另一机制：接口。

**方法签名的集合**叫做：_接口(Interfaces)_。

```go
package main

import (
	"fmt"
	"math"
)

// 这是一个几何体的基本接口。
type geometry interface {
	area() float64 // 接口中的方法1
	perim() float64 // 接口中的方法2
}

// 在这个例子中，我们将为 `rect` 和 `circle` 实现该接口。
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// **要在 Go 中实现一个接口，我们只需要实现接口中的所有方法。** 
// 这里我们为 `rect` 实现了 `geometry` 接口。
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// `circle` 的实现。
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量实现了某个接口，我们就可以调用指定接口中的方法。 
// 这儿有一个通用的 `measure` 函数，我们可以通过它来使用所有的 `geometry`。
func measure(g geometry) {
    // measure函数中使用geometry的接口
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

    // 结构体类型 circle 和 rect 都实现了 geometry 接口，
    // 所以我们可以将其实例作为 measure 的参数。
	measure(r)
	measure(c)
}
```

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\interfaces\interfaces.go"  
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793

要学习更多关于 Go 接口的知识， 可以看看这篇[很棒的博文](http://jordanorelli.tumblr.com/post/32665860244/how-to-use-interfaces-in-go)。

## 21-错误处理 反复思考🤔

*符合 Go 语言习惯的做法是使用一个独立、明确的返回值来传递错误信息*。 这与 Java、Ruby 使用的异常（exception） 以及在 C 语言中有时用到的重载 (overloaded) 的单返回/错误值有着明显的不同。 Go 语言的处理方式能清楚的知道哪个函数返回了错误，并使用跟其他（无异常处理的）语言类似的方式来处理错误。

按照惯例，错误通常是最后一个返回值并且是 `error` 类型，它是一个内建的接口。

`errors.New` 使用给定的错误信息构造一个基本的 `error` 值。

返回错误值为 `nil` 代表没有错误。

你还可以通过实现 `Error()` 方法来自定义 error 类型。 这里使用自定义错误类型来表示上面例子中的参数错误。

在这个例子中，我们使用 `&argError` 语法来建立一个新的结构体， 并提供了 `arg` 和 `prob` 两个字段的值。

下面的两个循环测试了每一个会返回错误的函数。 注意，在 if 的同一行进行错误检查，是 Go 代码中的一种常见用法。

如果你想在程序中使用自定义错误类型的数据， 你需要通过类型断言来得到这个自定义错误类型的实例。

到 Go 官方博客去看这篇[很棒的文章](http://blog.golang.org/2011/07/error-handling-and-go.html)， 以获取更多关于错误处理的信息。

```go
package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New 使用给定的错误信息构造一个基本的 error 值。
		return -1, errors.New("can't work with 42")
	}
	
	return arg + 3, nil // 返回错误值为 nil 代表没有错误。
}

type argError struct {
	arg int
	prob string
}

// 你还可以通过实现 Error() 方法来自定义 error 类型。 
// 这里使用自定义错误类型来表示上面例子中的参数错误。
func (e *argError) Error() string { // * 解引用指针，从对应地址获取值
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// 在这个例子中，我们使用 &argError 语法来建立一个新的结构体， 
// 并提供了 arg 和 prob 两个字段的值。
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// 下面的两个循环测试了每一个会返回错误的函数。 
	// 注意，在 if 的同一行进行错误检查，是 Go 代码中的一种常见用法。
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

// 如果你想在程序中使用自定义错误类型的数据， 
// 你需要通过类型断言来得到这个自定义错误类型的实例。
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
		fmt.Println("***Test***") // test
		fmt.Println(ok) // test
		fmt.Println(e.(*argError))
		fmt.Println(e)
	}
}
```

PS D:\Coding-cf> go run "d:\Coding-cf\Go\GoByExample\errors\errors.go"
f1 worked: 10
f1 failed: can't work with 42
f2 worked: 10
f2 failed: 42 - can't work with it
42
can't work with it
***Test***
true
42 - can't work with it
42 - can't work with it

符合 Go 语言习惯的做法是使用一个独立、明确的返回值来传递错误信息。 这与 Java、Ruby 使用的异常（exception） 以及在 C 语言中有时用到的重载 (overloaded) 的单返回/错误值有着明显的不同。 Go 语言的处理方式能清楚的知道哪个函数返回了错误，并使用跟其他（无异常处理的）语言类似的方式来处理错误。

按照惯例，错误通常是最后一个返回值并且是 `error` 类型，它是一个内建的接口。

`errors.New` 使用给定的错误信息构造一个基本的 `error` 值。

返回错误值为 `nil` 代表没有错误。

你还可以通过实现 `Error()` 方法来自定义 `error` 类型。 这里使用自定义错误类型来表示上面例子中的参数错误。

在这个例子中，我们使用 `&argError` 语法来建立一个新的结构体， 并提供了 `arg` 和 `prob` 两个字段的值。

下面的两个循环测试了每一个会返回错误的函数。 *注意，在 if 的同一行进行错误检查，是 Go 代码中的一种常见用法。*

如果你想在程序中使用自定义错误类型的数据， 你需要*通过类型断言来得到这个自定义错误类型的实例*。

到 Go 官方博客去看这篇[很棒的文章](http://blog.golang.org/2011/07/error-handling-and-go.html)， 以获取更多关于错误处理的信息。

```go
package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
```

PS C:\Users\chenfengyuan\Coding-cf> go run "c:\Users\chenfengyuan\Coding-cf\Go\GoByExample\errors\errors.go"
f1 worked: 10
f1 failed: can't work with 42
f2 worked: 10
f2 failed: 42 - can't work with it
42
can't work with it

## 22-协程

**协程(goroutine)**是轻量级的执行线程。

假设我们有一个函数叫做 f(s)。 我们一般会这样 *同步地* 调用它

使用 go f(s) 在一个协程中调用这个函数。 这个新的 Go 协程将会 *并发地* 执行这个函数。

你也可以为匿名函数启动一个协程。

现在两个协程在独立的协程中 *异步地* 运行， 然后等待两个协程完成（更好的方法是使用 [WaitGroup](https://gobyexample-cn.github.io/waitgroups)）。

```go
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// 假设我们有一个函数叫做 f(s)。 我们一般会这样 同步地 调用它
	f("direct")

	// 使用 go f(s) 在一个协程中调用这个函数。 这个新的 Go 协程将会 并发地 执行这个函数。
	go f("goroutine")

	// 你也可以为匿名函数启动一个协程。
	go func(msg string) {
		fmt.Println(msg)
	} ("going")

	// 现在两个协程在独立的协程中 异步地 运行， 然后等待两个协程完成（更好的方法是使用 WaitGroup）。
	time.Sleep(time.Second)
	fmt.Println("done")
}
```

PS D:\Coding-cf\C\Hello> go run "d:\Coding-cf\Go\GoByExample\goroutines\goroutines.go"
direct : 0
direct : 1
direct : 2
going
goroutine : 0
goroutine : 1
goroutine : 2
done

当我们运行这个程序时，首先会看到阻塞式调用的输出，然后是两个协程的交替输出。 这种*交替的情况表示 Go runtime 是以并发的方式运行协程的*。

## 23-通道

下来我们将学习协程的辅助特性：通道（channels）。

**通道(channels)** 是**连接多个协程**的管道。 你可以*从一个协程将值发送到通道，然后在另一个协程中接收*。

使用 `make(chan val-type)` 创建一个新的通道。 通道类型就是他们需要传递值的类型。

使用 `channel <-` 语法 *发送* 一个新的值到通道中。 这里我们在一个新的协程中发送 `"ping"` 到上面创建的 `messages` 通道中。

使用 `<-channe`l 语法从通道中 *接收* 一个值。 这里我们会收到在上面发送的 `"ping"` 消息并将其打印出来。

```go
package main

import "fmt"

func main() {
	// 使用 make(chan val-type) 创建一个新的通道。 
	// 通道类型就是他们需要传递值的类型。
	messages := make(chan string)

	go func() {
		// 使用 channel <- 语法 发送 一个新的值到通道中。 
		// 这里我们在一个新的协程中发送 "ping" 到上面创建的 messages 通道中。
		messages <- "ping"
	}()

	// 使用 <-channel 语法从通道中 接收 一个值。 
	// 这里我们会收到在上面发送的 "ping" 消息并将其打印出来。

	msg := <- messages
	fmt.Println(msg)
}
```

PS D:\Coding-cf\C\Hello> go run "d:\Coding-cf\Go\GoByExample\channels\channels.go"
ping

我们运行程序时，通过通道， 成功的将消息 "ping" 从一个协程传送到了另一个协程中。

*默认发送和接收操作是阻塞的，直到发送方和接收方都就绪。* 这个特性允许我们，不使用任何其它的同步操作， 就可以在程序结尾处等待消息 "ping"。

## 24-通道缓冲

默认情况下，通道是 *无缓冲* 的，这意味着只有对应的接收（`<- chan`） 通道准备好接收时，才允许进行发送（`chan <-`）。 *有缓冲通道* 允许在没有对应接收者的情况下，缓存一定数量的值。

```go
package main

import "fmt"

func main() {
	// 这里我们 make 了一个字符串通道，最多允许缓存 2 个值。
	messages := make(chan string, 2)

	// 由于此通道是有缓冲的， 因此我们可以将这些值发送到通道中，而无需并发的接收。
	messages <- "buffered"
	messages <- "channel"

	// 然后我们可以正常接收这两个值。
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
```

PS D:\Coding-cf\C\Hello> go run "d:\Coding-cf\Go\GoByExample\channel-buffering\channel-buffering.go"
buffered
channel

## 25-通道同步 思考🤔

我们可以**使用通道来同步协程之间的执行状态**。 这儿有一个例子，*使用阻塞接收的方式，实现了等待另一个协程完成*。 **如果需要等待多个协程，WaitGroup 是一个更好的选择。**

```go
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	// 我们将要在协程中运行这个函数。 
	// done 通道将被用于通知其他协程这个函数已经完成工作。
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值来通知我们已经完工啦。
	done <- true // 从一个协程将值发送到通道
}

func main() {

	// 运行一个 worker 协程，并给予用于通知的通道。
	done := make(chan bool, 1)
	go worker(done)

	// 程序将一直阻塞，直至收到 worker 使用通道发送的通知。
	<-done // 从通道接受
	// 如果你把 <- done 这行代码从程序中移除， 程序甚至可能在 worker 开始运行前就结束了。
}
```

PS D:\Coding-cf\C\Hello> go run "d:\Coding-cf\Go\GoByExample\channel-synchronization\channel-synchronization.go"
working...done

## 26-通道方向

*当使用通道作为函数的参数时，你可以指定这个通道是否为只读或只写。 该特性可以提升程序的类型安全。*

```go
package main

import "fmt"

// ping 函数定义了一个只能发送数据的（只写）通道。 
// 尝试从这个通道接收数据会是一个编译时错误。
func ping(pings chan<- string, msg string) {
	pings <- msg // 发送数据到管道
}

// pong 函数接收两个通道，pings 仅用于接收数据（只读），
// pongs 仅用于发送数据（只写）。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <- pings // 从管道接收数据
	pongs <- msg // 发送数据到管道
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs) // 从通道接收
}
```

PS D:\Coding-cf\C\Hello> go run "d:\Coding-cf\Go\GoByExample\channel-directions\channel-directions.go"
passed message

> 直接打印通道，会输出通道对应的内存地址

```go
PS D:\Coding-cf\C\Hello> go run "d:\Coding-cf\Go\GoByExample\channel-directions\channel-directions.go"
0xc00006c0c0
```

## 27-通道选择器

Go 的 **选择器（select）** 让你可以同时等待多个通道操作。 **将协程、通道和选择器结合，是 Go 的一个强大特性。**

```GO
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// 在这个例子中，我们将从两个通道中选择。
	c1 := make(chan string)
	c2 := make(chan string)

	// 各个通道将在一定时间后接收一个值， 
	// 通过这种方式来模拟并行的协程执行（例如，RPC 操作）时造成的阻塞（耗时）。
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	} ()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	} ()

	// 我们使用 select 关键字来同时等待这两个值， 
	// 并打印各自接收到的值。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- c1:
			fmt.Println("received", msg1)
		case msg2 := <- c2:
			fmt.Println("received", msg2)
		}
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)
}
```

PS D:\Coding-cf\C\Hello> go run "d:\Coding-cf\Go\GoByExample\select\select.go"
received one
received two
cost=[2.0040901s]

跟预期的一样，我们首先接收到值 "one"，然后是 "two"。

注意，程序总共仅运行了两秒左右。因为 1 秒 和 2 秒的 Sleeps 是并发执行的，

## 28-超时处理

**超时** 对于一个需要连接外部资源， 或者有耗时较长的操作的程序而言是很重要的。 得益于通道和 select，在 Go 中实现超时操作是简洁而优雅的。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 在这个例子中，假如我们执行一个外部调用， 
	// 并在 2 秒后使用通道 c1 返回它的执行结果。
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	} ()

	// 这里是使用 select 实现一个超时操作。 
	// res := <- c1 等待结果，<-Time.After 等待超时（1秒钟）以后发送的值。 
	// 由于 select 默认处理第一个已准备好的接收操作，
	// 因此如果操作耗时超过了允许的 1 秒的话，将会执行超时 case。
	select {
	case res := <- c1:
		fmt.Println(res)
	case <- time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	// 如果我们允许一个长一点的超时时间：3 秒， 
	// 就可以成功的从 c2 接收到值，并且打印出结果。
	select {
	case res := <- c2:
		fmt.Println(res)
	case <- time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
```

PS D:\Coding-cf\Go\GoByExample\select> go run "d:\Coding-cf\Go\GoByExample\timeouts\timeouts.go"
timeout 1
result 2

运行这个程序，首先显示运行超时的操作，然后是成功接收的。

## 29-非阻塞通道操作

常规的通过通道发送和接收数据是阻塞的。 然而，我们可以*使用带一个 `default` 子句的 `select` 来实现 **非阻塞** 的发送、接收，甚至是非阻塞的多路 `select`*。

```go
package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// 这是一个非阻塞接收的例子。 
	// 如果在 messages 中存在，然后 select 将这个值带入 <-messages case 中。 
	// 否则，就直接到 default 分支中。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 一个非阻塞发送的例子，代码结构和上面接收的类似。 
	// msg 不能被发送到 message 通道，
	// 因为这是 个无缓冲区通道，并且也没有接收者，
	// 因此， default 会执行。
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 我们可以在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器。 
	// 这里我们试图在 messages 和 signals 上同时使用非阻塞的接收操作。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
```

PS D:\Coding-cf\Go\GoByExample\select> go run "d:\Coding-cf\Go\GoByExample\non-blocking-channel-operations\non-blocking-channel-operations.go"
no message received
no message sent
no activity

## 30-通道的关闭 Question

> 此部分运行结果和教程中的给出结果不一致，应该是Go版本迭代出现的问题，根据最新Go特性来解释。

**关闭** 一个通道意味着不能再向这个通道发送值了。 该特性可以向通道的接收方传达工作已经完成的信息。

```go
package main

import "fmt"

// 在这个例子中，我们将使用一个 jobs 通道，将工作内容， 
// 从 main() 协程传递到一个工作协程中。 
// 当我们没有更多的任务传递给工作协程时，我们将 close 这个 jobs 通道。

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 这是工作协程。使用 j, more := <- jobs 循环的从 jobs 接收数据。 
	// 根据接收的第二个值，如果 jobs 已经关闭了， 并且通道中所有的值都已经接收完毕，
	// 那么 more 的值将是 false。 当我们完成所有的任务时，
	// 会使用这个特性通过 done 通道通知 main 协程。
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 使用 jobs 发送 3 个任务到工作协程中，然后关闭 jobs。
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	fmt.Println("Test1")

	// 使用前面学到的通道同步方法等待任务结束。
	<-done // 从通道接受
}
```

PS D:\Coding-cf\Go\GoByExample\select> go run "d:\Coding-cf\Go\GoByExample\closing-channels\closing-channels.go"
sent job 1
sent job 2
sent job 3
sent all jobs
Test1
received job 1
received job 2
received job 3
received all jobs

给出的参考结果：
```shell
$ go run closing-channels.go
sent job 1
received job 1
sent job 2
received job 2
sent job 3
received job 3
sent all jobs
received all jobs
```

根据**关闭通道**的思想，可以引出我们的下一个示例：遍历通道。

## 31-通道遍历

在前面的例子中， 我们讲过 `for` 和 `range` 为基本的数据结构提供了迭代的功能。 我们也可以使用这个语法来遍历的从通道中取值。

```go
package main

import "fmt"

func main() {

	// 我们将遍历在 queue 通道中的两个值。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// range 迭代从 queue 中得到每个值。 
	// 因为我们在前面 close 了这个通道，
	// 所以，这个迭代会在接收完 2 个值之后结束。
	for elem := range queue {
		fmt.Println(elem)
	}
}
```

PS D:\Coding-cf\Go\GoByExample\select> go run "d:\Coding-cf\Go\GoByExample\range-over-channels\tempCodeRunnerFile.go"
one
two

*这个例子也让我们看到，一个非空的通道也是可以关闭的， 并且，通道中剩下的值仍然可以被接收到。*

## 32-Timer

我们经常需要在未来的某个时间点运行 Go 代码，或者每隔一定时间重复运行代码。 Go 内置的 *定时器* 和 *打点器* 特性让这些变得很简单。 我们会先学习定时器，然后再学习打点器。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// 定时器表示在未来某一时刻的独立事件。
	// 你告诉定时器需要等待的时间，
	// 然后它将提供一个用于通知的通道。
	// 这里的定时器将等待 2 秒。
	timer1 := time.NewTimer(2 * time.Second)

	// <-timer1.C 会一直阻塞，
	// 直到定时器的通道 C 明确的发送了定时器失效的值。
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// 如果你需要的仅仅是单纯的等待，使用 time.Sleep 就够了。
	// 使用定时器的原因之一就是，你可以在定时器触发之前将其取消。例如这样。
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	// 给 timer2 足够的时间来触发它，以证明它实际上已经停止了。
	time.Sleep(2 * time.Second)

	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
```

第一个定时器将在程序开始后大约 2s 触发， 但是第二个定时器还未触发就停止了。

PS D:\Coding-cf\Go\GoByExample\select> go run "d:\Coding-cf\Go\GoByExample\timers\timers.go"
Timer 1 fired
Timer 2 stopped
cost=[4.0115602s]

## 33-Ticker

- **定时器** 是当你*想要在未来某一刻执行一次时使用的*
- **打点器** 则是为你*想要以固定的时间间隔重复执行而准备的。* 这里是一个打点器的例子，它将定时的执行，直到我们将它停止。

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	// 打点器和定时器的机制有点相似：使用一个通道来发送数据。
	// 这里我们使用通道内建的 select，等待每 500ms 到达一次的值。
	fmt.Println("Test1") // Test
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// 打点器可以和定时器一样被停止。
	// 打点器一旦停止，将不能再从它的通道中接收到值。
	// 我们将在运行 1600ms 后停止这个打点器。
	fmt.Println("Test2") // Test
	time.Sleep(1600 * time.Millisecond)
	fmt.Println("Test3") // Test
	ticker.Stop()
	fmt.Println("Test4") // Test
	done <- true
	fmt.Println("Ticker stopped")

	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}

```

我们运行这个程序时，打点器会在我们停止它前打点 3 次。

PS D:\Coding-cf\Go\GoByExample\select> go run "d:\Coding-cf\Go\GoByExample\tickers\tickers.go"
Test1
Test2
Tick at 2021-11-26 14:54:56.7337073 +0800 CST m=+0.514811001
Tick at 2021-11-26 14:54:57.2339533 +0800 CST m=+1.015057001
Tick at 2021-11-26 14:54:57.7348781 +0800 CST m=+1.515981801
Test3
Test4
Ticker stopped
cost=[1.6045433s]

## 34-工作池

在这个例子中，我们将看到如何使用*协程与通道*实现一个**工作池**。

```go
package main

import (
	"fmt"
	"time"
)

// 这是 worker 程序，我们会并发的运行多个 worker。
// worker 将在 jobs 频道上接收工作，并在 results 上发送相应的结果。
// 每个 worker 我们都会 sleep 一秒钟，
// 以模拟一项昂贵的（耗时一秒钟的）任务。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	start := time.Now()

	// 为了使用 worker 工作池并且收集其的结果，我们需要 2 个通道。
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 这里启动了 3 个 worker， 初始是阻塞的，因为还没有传递任务。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	fmt.Println("Test1") // Test
	// 这里我们发送 5 个 jobs，
	// 然后 close 这些通道，表示这些就是所有的任务了。
	for j := 1; j <= numJobs; j++ {
		jobs <- j // 协程输出的地方
	}

	close(jobs)
	fmt.Println("Test2") // Test
	// 最后，我们收集所有这些任务的返回值。
	// 这也确保了所有的 worker 协程都已完成。
	// 另一个等待多个协程的方法是使用WaitGroup。
	for a := 1; a <= numJobs; a++ {
		<-results
	}
	fmt.Println("Test3") // Test

	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
```

运行程序，显示 5 个任务被多个 worker 执行。 尽管所有的工作总共要花费 5 秒钟，但该程序只花了 2 秒钟， 因为 3 个 worker 是并行的。

PS D:\Coding-cf\Go\GoByExample\select> go run "d:\Coding-cf\Go\GoByExample\worker-pools\worker-pools.go"
Test1
Test2
worker 3 started job 3
worker 1 started job 1
worker 2 started job 2
worker 2 finished job 2
worker 2 started job 4
worker 1 finished job 1
worker 1 started job 5
worker 3 finished job 3
worker 1 finished job 5
worker 2 finished job 4
Test3
cost=[2.0199268s]

## 35-WaitGroup

想要等待多个协程完成，我们可以使用 *wait group*。

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// 每个协程都会运行该函数。 注意，WaitGroup 必须通过指针传递给函数。
func worker(id int, wg *sync.WaitGroup) {

	// return 时，通知 WaitGroup，当前协程的工作已经完成。
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	// 睡眠一秒钟，以此来模拟耗时的任务。
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	start := time.Now()

	// WaitGroup 用于等待该函数启动的所有协程。
	var wg sync.WaitGroup

	// 启动几个协程，并为其递增 WaitGroup 的计数器
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	fmt.Println("Test")

	// 阻塞，直到 WaitGroup 计数器恢复为 0； 即所有协程的工作都已经完成。
	wg.Wait()

	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
```

每次运行，各个协程开启和完成的时间可能是不同的。

PS D:\Coding-cf\Go\GoByExample\select> go run "d:\Coding-cf\Go\GoByExample\waitgroups\waitgroups.go"
Worker 2 starting
Worker 3 starting
Test
Worker 5 starting
Worker 4 starting
Worker 1 starting
Worker 1 done
Worker 3 done
Worker 4 done
Worker 2 done
Worker 5 done
cost=[1.009809s]

## 36-速率限制

**速率限制** 是控制服务资源利用和质量的重要机制。 *基于协程、通道和打点器，Go 优雅的支持速率限制*。

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	// 首先，我们将看一个基本的速率限制。
	// 假设我们想限制对收到请求的处理，我们可以通过一个通道处理这些请求
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// limiter 通道每 200ms 接收一个值。 这是我们任务速率限制的调度器。
	limiter := time.Tick(200 * time.Millisecond)

	// 通过在每次请求前阻塞 limiter 通道的一个接收，
	// 可以将频率限制为，每 200ms 执行一次请求。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 有时候我们可能希望在速率限制方案中允许短暂的并发请求，
	// 并同时保留总体速率限制。 我们可以通过缓冲通道来完成此任务。
	// burstyLimiter 通道允许最多 3 个爆发（bursts）事件。
	burstLimiter := make(chan time.Time, 3)

	// 填充通道，表示允许的爆发（bursts）。
	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	// 每 200ms 我们将尝试添加一个新的值到 burstyLimiter中， 直到达到 3 个的限制。
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstLimiter <- t
		}
	}()

	// 现在，模拟另外 5 个传入请求。 受益于 burstyLimiter 的爆发（bursts）能力，
	// 前 3 个请求可以快速完成。
	burstRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstRequests <- i
	}
	close(burstRequests)
	for req := range burstRequests {
		<-burstLimiter
		fmt.Println("request", req, time.Now())
	}

	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
```

运行程序，我们看到第一批请求意料之中的大约每 200ms 处理一次。

PS D:\Coding-cf\Go\GoByExample\select> go run "d:\Coding-cf\Go\GoByExample\rate-limiting\rate-limiting.go"
request 1 2021-11-26 15:54:52.7125512 +0800 CST m=+0.218740401
request 2 2021-11-26 15:54:52.9143874 +0800 CST m=+0.420576601
request 3 2021-11-26 15:54:53.116389 +0800 CST m=+0.622578201
request 4 2021-11-26 15:54:53.3061452 +0800 CST m=+0.812334401
request 5 2021-11-26 15:54:53.5101447 +0800 CST m=+1.016333901
request 1 2021-11-26 15:54:53.5101447 +0800 CST m=+1.016333901
request 2 2021-11-26 15:54:53.5101447 +0800 CST m=+1.016333901 
request 3 2021-11-26 15:54:53.5101447 +0800 CST m=+1.016333901 
request 4 2021-11-26 15:54:53.7157193 +0800 CST m=+1.221908501
request 5 2021-11-26 15:54:53.9223151 +0800 CST m=+1.428504301
cost=[1.4199199s]

第二批请求，由于爆发（burstable）速率控制，我们直接连续处理了 3 个请求， 然后以大约每 200ms 一次的速度，处理了剩余的 2 个请求。

## 37-原子计数器

Go 中最主要的状态管理机制是依靠通道间的通信来完成的。 我们已经在工作池的例子中看到过，并且，还有一些其他的方法来管理状态。 这里，我们来看看如何使用 `sync/atomic` 包在多个协程中进行 _原子计数_。

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	start := time.Now()

	// 我们将使用一个无符号整型（永远是正整数）变量来表示这个计数器。
	var ops uint64

	// WaitGroup 帮助我们等待所有协程完成它们的工作。
	var wg sync.WaitGroup

	// 我们会启动 50 个协程，并且每个协程会将计数器递增 1000 次。
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				// 使用 AddUint64 来让计数器自动增加，
				// 使用 & 语法给定 ops 的内存地址。
				atomic.AddUint64(&ops, 1)

				// 由于多个协程会互相干扰，运行时值会改变，可能会导致我们得到一个不同的数字
				// ops++
			}
			wg.Done()
		}()
	}

	// 等待，直到所有协程完成。
	wg.Wait()

	// 现在可以安全的访问 ops，因为我们知道，此时没有协程写入 ops， 此外，
	// 还可以使用 atomic.LoadUint64 之类的函数，在原子更新的同时安全地读取它们。
	fmt.Println("ops:", ops)

	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
```

PS D:\Coding-cf\Go\GoByExample\atomic-counters> go run "d:\Coding-cf\Go\GoByExample\atomic-counters\atomic-counters.go"
ops: 50000
cost=[1.0355ms]

预计会进行 `50,000` 次操作。如果我们使用非原子的 `ops++` 来增加计数器， 由于多个协程会互相干扰，运行时值会改变，可能会导致我们得到一个不同的数字。 此外，运行程序时带上 `-race` 标志，我们可以获取数据竞争失败的详情

PS D:\Coding-cf\Go\GoByExample\atomic-counters> go run atomic-counters.go -race      
ops: 49157
cost=[0s]

接下来，我们看一下管理状态的另一个工具——互斥锁。

## 38-互斥锁

在前面的例子中，我们看到了如何使用原子操作来管理简单的计数器。 对于更加复杂的情况，我们可以使用一个**互斥锁** 来在 Go 协程间安全的访问数据。

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	start := time.Now()

	// 在这个例子中，state 是一个 map。
	var state = make(map[int]int)

	// mutex 将同步对 state 的访问。
	var mutex = &sync.Mutex{}

	// 我们会持续追踪读写操作的数量。
	var readOps uint64
	var writeOps uint64

	// 这里我们启动 100 个协程， 每个协程以每 1ms 一次的频率来重复读取 state。
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				// 每次循环读取，我们使用一个键来进行访问，
				// Lock() 这个 mutex 来确保对 state 的独占访问，
				// 读取选定的键的值，Unlock() 这个 mutex，
				// 并将 ops 值加 1。
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				// 在下次读取前等待片刻。
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 同样的，我们启动 10 个协程来模拟写入操作， 使用和读取相同的模式。

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 让这 10 个协程对 state 和 mutex 的操作持续 1 s。
	time.Sleep(time.Second)

	// 获取并输出最终的操作计数。
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// 对 state 使用一个最终的锁，展示它是如何结束的。
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()

	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
```

**Question**：*导致教程、PC、The Go Playground运行结果不一致的原因？*

> GoByExample 教程中的运行结果

运行这个程序，显示我们进行了大约 90,000 次 mutex 同步的 state 操作。

```shell
$ go run mutexes.go
readOps: 83285
writeOps: 8320
state: map[1:97 4:53 0:33 2:15 3:2]
```

> PC运行结果

PS D:\Coding-cf\Go\GoByExample\atomic-counters> go run "d:\Coding-cf\Go\GoByExample\mutexes\mutexes.go"
readOps: 6557
writeOps: 660
state: map[0:28 1:8 2:50 3:78 4:9]
cost=[1.0066632s]

> The Go Playground运行结果

readOps: 100001
writeOps: 10000
state: map[0:94 1:87 2:53 3:90 4:8]
cost=[1s]

接下来我们将看一下，*只使用协程和通道， 如何实现相同的任务状态管理*。

## 39-状态协程 思考🤔

在前面的例子中，我们用 互斥锁 进行了明确的锁定， 来让共享的 state 跨多个 Go 协程同步访问。 另一个选择是，*使用内建协程和通道的同步特性来达到同样的效果*。 **Go 共享内存的思想是**，*通过通信使每个数据仅被单个协程所拥有，即**通过通信实现共享内存***。 基于通道的方法与该思想完全一致！

```go
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 在这个例子中，state 将被一个单独的协程拥有。
// 这能保证数据在并行读取时不会混乱。 为了对 state 进行读取或者写入，
// 其它的协程将发送一条数据到目前拥有数据的协程中，
// 然后等待接收对应的回复。 结构体 readOp 和 writeOp 封装了这些请求，
// 并提供了响应协程的方法。

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	start := time.Now()

	// 和前面的例子一样，我们会计算操作执行的次数。
	var readOps uint64
	var writeOps uint64

	// 其他协程将通过 reads 和 writes 通道来发布 读 和 写 请求。
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// 这就是拥有 state 的那个协程， 和前面例子中的 map 一样，
	// 不过这里的 state 是被这个状态协程私有的。
	// 这个协程不断地在 reads 和 writes 通道上进行选择，
	// 并在请求到达时做出响应。 首先，执行请求的操作；然后，执行响应，
	// 在响应通道 resp 上发送一个值，
	// 表明请求成功（reads 的值则为 state 对应的值）。
	go func() { // 此协程的运行机制是什么？
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 启动 100 个协程通过 reads 通道向拥有 state 的协程发起读取请求。
	// 每个读取请求需要构造一个 readOp，发送它到 reads 通道中，
	// 并通过给定的 resp 通道接收结果。

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read // 发送数据
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 用相同的方法启动 10 个写操作。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 让协程们跑 1s。
	time.Sleep(time.Second)

	// 最后，获取并报告 ops 值。
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
```

运行这个程序显示这个基于协程的状态管理的例子 达到了每秒大约 80,000 次操作。

> GoByExample 教程中的运行结果

```shell
$ go run stateful-goroutines.go
readOps: 71708
writeOps: 7177
```

> PC运行结果

PS D:\Coding-cf\Go\GoByExample\atomic-counters> go run "d:\Coding-cf\Go\GoByExample\stateful-goroutines\stateful-goroutines.go"
readOps: 6589
writeOps: 660

> The Go Playground运行结果

readOps: 100099
writeOps: 10010
cost=[1s]

通过这个例子我们可以看到，**基于协程的方法比基于互斥锁的方法要复杂得多**。 但是，在某些情况下它可能很有用， 例如，当你涉及其他通道，或者管理多个同类互斥锁时，会很容易出错。 *您应该使用最自然的方法，尤其是在理解程序正确性方面。*

## 40-排序

Go 的 sort 包实现了内建及用户自定义数据类型的排序功能。 我们先来看看内建数据类型的排序。

```go
package main

import (
	"fmt"
	"sort"
)

func main() {

	// 排序方法是针对内置数据类型的； 这是一个字符串排序的例子。
	// 注意，它是原地排序的，所以他会直接改变给定的切片，
	// 而不是返回一个新切片。
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// 一个 int 排序的例子。
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	// 我们也可以使用 sort 来检查一个切片是否为有序的。
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
```

运行程序，打印排序好的字符串和整型切片， 以及数组是否有序的检查结果：true。

PS D:\Coding-cf\Go\GoByExample\atomic-counters> go run "d:\Coding-cf\Go\GoByExample\sorting\sorting.go"
Strings: [a b c]
Ints:  [2 4 7]
Sorted:  true

## 41-使用函数自定义排序

有时候，我们可能想根据自然顺序以外的方式来对集合进行排序。 例如，假设我们要按字符串的长度而不是按字母顺序对它们进行排序。 这儿有一个在 Go 中自定义排序的示例。

```go
package main

import (
	"fmt"
	"sort"
)

// 为了在 Go 中使用自定义函数进行排序，我们需要一个对应的类型。 
// 我们在这里创建了一个 byLength 类型，它只是内建类型 []string 的别名。
type byLength []string

// 我们为该类型实现了 sort.Interface 接口的 Len、Less 和 Swap 方法， 
// 这样我们就可以使用 sort 包的通用 Sort 方法了， Len 和 Swap 在各个类型中的实现都差不多， 
// Less 将控制实际的自定义排序逻辑。 
// 在这个的例子中，我们想按字符串长度递增的顺序来排序， 
// 所以这里使用了 len(s[i]) 和 len(s[j]) 来实现 Less。
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	// 一切准备就绪后，我们就可以通过将切片 fruits 强转为 byLength 类型的切片， 
	// 然后对该切片使用 sort.Sort 来实现自定义排序。
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
```

运行这个程序，和预期的一样， 显示了一个按照字符串长度排序的列表。

PS D:\> go run "d:\Coding-cf\Go\GoByExample\sorting-by-functions\sorting-by-functions.go"
[kiwi peach banana]

类似的，参照这个例子，*创建一个自定义类型， 为它实现 Interface 接口的三个方法， 然后对这个自定义类型的切片调用 sort.Sort 方法， 我们就可以通过任意函数对 Go 切片进行排序了。*

## 42-Panic

`panic` 意味着有些出乎意料的错误发生。 通常我们用它来表示程序正常运行中不应该出现的错误， 或者我们不准备优雅处理的错误。

```go
package main

import "os"

func main() {

	// 我们将使用 panic 来检查这个站点上预期之外的错误。 
	// 而该站点上只有一个程序：触发 panic。
	panic("a problem")

	// panic 的一种常见用法是：
	// 当函数返回我们不知道如何处理（或不想处理）的错误值时，中止操作。 
	// 如果创建新文件时遇到意外错误该如何处理？这里有一个很好的 panic 示例。

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
```

运行程序将会导致 `panic`： 输出一个错误消息和协程追踪信息，并以非零的状态退出程序。

PS D:\Coding-cf\Go\Go-algorithm\数据结构与算法之美\06_linkedlist> go run "d:\Coding-cf\Go\GoByExample\panic\panic.go"
panic: a problem

goroutine 1 [running]:
main.main()
        d:/Coding-cf/Go/GoByExample/panic/panic.go:9 +0x27
exit status 2

> *注意，与某些使用 exception 处理错误的语言不同， 在 Go 中，通常会尽可能的使用返回值来标示错误。*

## 43-Defer

`Defer` *用于确保程序在执行完成后，会调用某个函数，一般是执行清理工作。* Defer 的用途跟其他语言的 `ensure` 或 `finally` 类似。

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	// 假设我们想要创建一个文件、然后写入数据、
	// 最后在程序结束时关闭该文件。 
	// 这里展示了如何通过 defer 来做到这一切。

	// 在 createFile 后立即得到一个文件对象， 
	// 我们使用 defer 通过 closeFile 来关闭这个文件。 
	// 这会在封闭函数（main）结束时执行，即 writeFile 完成以后。
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

// 关闭文件时，进行错误检查是非常重要的， 即使在 defer 函数中也是如此。
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// 写入文件的地址：D:\tmp
```

执行程序，确认写入数据后，文件已关闭。

PS D:\Coding-cf\Go\Go-algorithm\数据结构与算法之美\06_linkedlist> go run "d:\Coding-cf\Go\GoByExample\defer\defer.go"
creating
writing
closing

## 44-组合函数

我们经常需要程序对数据集合执行操作， 例如选择满足给定条件的全部 item， 或通过自定义函数将全部 item 映射到一个新的集合。

在其它语言中，通常会使用*泛型*数据结构和算法。 但 Go 不支持泛型，如果你的程序或者数据类型有需要，通常的做法是*提供函数集*。

这是一些 `strings` 切片的组合函数示例。 你可以使用这些例子来构建自己的函数。 注意，在某些情况下，最简单明了的方法是： **直接内联操作方法集，而不是创建并调用帮助函数。**

```go
package main

import (
	"fmt"
	"strings"
)

// Index 返回目标字符串 t 在 vs 中第一次出现位置的索引， 
// 或者在没有匹配值时返回 -1。
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include 如果目标字符串 t 存在于切片 vs 中，则返回 true。
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any 如果切片 vs 中的任意一个字符串满足条件 f，则返回 true。
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All 如果切片 vs 中的所有字符串都满足条件 f，则返回 true。
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter 返回一个新的切片，新切片由原切片 vs 中满足条件 f 的字符串构成。
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map 返回一个新的切片，新切片的字符串由原切片 vs 中的字符串经过函数 f 映射后得到。
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	// 试试各种组合函数。
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p") 
		// HasPrefix tests whether the string s begins with prefix.
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
	// 上面的例子都是用的匿名函数，当前，你也可以使用正确类型的命名函数
}
```

PS D:\Coding-cf\Go\Go-algorithm\数据结构与算法之美\06_linkedlist> go run "d:\Coding-cf\Go\GoByExample\collection-functions\collection-functions.go"
2
false
true
false
[peach apple pear]
[PEACH APPLE PEAR PLUM]

## 45-字符串函数

标准库的 `strings` 包提供了很多有用的字符串相关的函数。 这儿有一些用来让你对 `strings` 包有一个初步了解的例子。

```go
package main

import (
	"fmt"
	s "strings"
)

// 我们给 fmt.Println 一个较短的别名， 因为我们随后会大量的使用它。
var p = fmt.Println

// 这是一些 strings 中有用的函数例子。 由于它们都是包的函数，
// 而不是字符串对象自身的方法， 这意味着我们需要在调用函数时，
// 将字符串作为第一个参数进行传递。 
// 你可以在 strings 包文档中找到更多的函数。
func main() {
	p("Contains: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index: ", s.Index("test", "e"))
	p("Join: ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat: ", s.Repeat("a", 5))
	p("Replace: ", s.Replace("foo", "o", "0", -1))
	p("Replace: ", s.Replace("foo", "o", "0", 1))
	p("Split: ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:", s.ToLower("TEST"))
	p("ToUpper: ", s.ToUpper("test"))
	p()

	// 虽然不是 strings 的函数，但仍然值得一提的是， 
	// 获取字符串长度（以字节为单位）以及通过索引获取一个字节的机制。
	p("Len: ", len("hello"))
	p("Char: ", "ahello"[0])
	// 注意，上面的 len 以及索引工作在字节级别上。 
	// Go 使用 UTF-8 编码字符串，因此通常按原样使用。 
	// 如果您可能使用多字节的字符，则需要使用可识别编码的操作。 
	// 详情请参考 strings, bytes, runes and characters in Go。
}
```

PS D:\Coding-cf\Go\Go-algorithm\数据结构与算法之美\06_linkedlist> go run "d:\Coding-cf\Go\GoByExample\string-functions\string-functions.go"
Contains:  true
Count:  2
HasPrefix:  true
HasSuffix:  true
Index:  1
Join:  a-b
Repeat:  aaaaa
Replace:  f00
Replace:  f0o
Split:  [a b c d e]
ToLower: test
ToUpper:  TEST

Len:  5
Char:  97

## 46-字符串格式化

Go 在传统的 printf 中对字符串格式化提供了优异的支持。 这儿有一些基本的*字符串格式化的任务*的例子。

```go
package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	// Go 提供了一些用于格式化常规值的打印“动词”。 
	// 例如，这样打印 point 结构体的实例。
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// 如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
	fmt.Printf("%+v\n", p)

	// %#v 根据 Go 语法输出值，即会产生该值的源码片段。
	fmt.Printf("%#v\n", p)

	// 需要打印值的类型，使用 %T。
	fmt.Printf("%T\n", p)

	// 格式化布尔值很简单。
	fmt.Printf("%t\n", true)

	// 格式化整型数有多种方式，使用 %d 进行标准的十进制格式化。
	fmt.Printf("%d\n", 123)

	// 这个输出二进制表示形式。
	fmt.Printf("%b\n", 14)

	// 输出给定整数的对应字符。
	fmt.Printf("%c\n", 33)

	// %x 提供了十六进制编码。
	fmt.Printf("%x\n", 456)

	// 同样的，也为浮点型提供了多种格式化选项。 
	// 使用 %f 进行最基本的十进制格式化。
	fmt.Printf("%f\n", 78.9)

	// %e 和 %E 将浮点型格式化为（稍微有一点不同的）科学记数法表示形式。
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// 使用 %s 进行基本的字符串输出。
	fmt.Printf("%s\n", "\"string\"")

	// 像 Go 源代码中那样带有双引号的输出，使用 %q。
	fmt.Printf("%q\n", "\"string\"")

	// 和上面的整型数一样，%x 输出使用 base-16 编码的字符串， 
	// 每个字节使用 2 个字符表示。
	fmt.Printf("%x\n", "hex this")

	// 要输出一个指针的值，使用 %p。
	fmt.Printf("%p\n", &p)

	// 格式化数字时，您经常会希望控制输出结果的宽度和精度。 
	// 要指定整数的宽度，请在动词 “%” 之后使用数字。 
	// 默认情况下，结果会右对齐并用空格填充。
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// 你也可以指定浮点型的输出宽度，
	// 同时也可以通过 宽度.精度 的语法来指定输出的精度。
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// 要左对齐，使用 - 标志。
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// 你也许也想控制字符串输出时的宽度，
	// 特别是要确保他们在类表格输出时的对齐。 这是基本的宽度右对齐方法。
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// 要左对齐，和数字一样，使用 - 标志。
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// 到目前为止，我们已经看过 Printf 了， 
	// 它通过 os.Stdout 输出格式化的字符串。
	//  Sprintf 则格式化并返回一个字符串而没有任何输出。
	s := fmt.Sprintf("a %s", "string")

	fmt.Println(s)

	// 你可以使用 Fprintf 来格式化并输出到 io.Writers 而不是 os.Stdout。
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
```

PS D:\Coding-cf\Go\Go-algorithm\数据结构与算法之美\06_linkedlist> go run "d:\Coding-cf\Go\GoByExample\string-formatting\string-formatting.go"
{1 2}
{x:1 y:2}
main.point{x:1, y:2}
main.point
true
123
1110
!
1c8
78.900000
1.234000e+08
1.234000E+08
"string"
"\"string\""
6865782074686973
0xc0000140a0
|    12|   345|
|  1.20|  3.45|
|1.20  |3.45  |
|   foo|     b|
|foo   |b     |
a string
an error

## 47-正则表达式

Go 提供了内建的[正则表达式](http://zh.wikipedia.org/wiki/%E6%AD%A3%E5%88%99%E8%A1%A8%E8%BE%BE%E5%BC%8F)支持。 这儿有一些在 Go 中与 regexp 相关的常见用法示例。

```go
package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// 测试一个字符串是否符合一个表达式。
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 上面我们是直接使用字符串，但是对于一些其他的正则任务， 
	// 你需要通过 Compile 得到一个优化过的 Regexp 结构体。
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 该结构体有很多方法。这是一个类似于我们前面看到的匹配测试。
	fmt.Println(r.MatchString("peach"))

	// 查找匹配的字符串。
	fmt.Println(r.FindString("peach punch"))

	// 这个也是查找首次匹配的字符串， 但是它的返回值是，
	// 匹配开始和结束位置的索引，而不是匹配的内容。
	fmt.Println(r.FindStringIndex("peach punch")) // [0 5]

	// Submatch 返回完全匹配和局部匹配的字符串。 
	// 例如，这里会返回 p([a-z]+)ch 和 ([a-z]+) 的信息。
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// 类似的，这个会返回完全匹配和局部匹配位置的索引。
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// 带 All 的这些函数返回全部的匹配项， 而不仅仅是首次匹配项。
	// 例如查找匹配表达式全部的项。
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// All 同样可以对应到上面的所有函数。
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	
	// 这些函数接收一个正整数作为第二个参数，来限制匹配次数。
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 上面的例子中，我们使用了字符串作为参数， 并使用了 MatchString 之类的方法。 
	// 我们也可以将 String 从函数命中去掉，并提供 []byte 的参数。
	fmt.Println(r.Match([]byte("peach")))

	// 创建正则表达式的全局变量时，可以使用 Compile 的变体 MustCompile 。 
	// MustCompile 用 panic 代替返回一个错误 ，这样使用全局变量更加安全。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// regexp 包也可以用来将子字符串替换为为其它值。
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Func 变体允许您使用给定的函数来转换匹配的文本。
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
```

PS D:\Coding-cf\Go\Go-algorithm\数据结构与算法之美\06_linkedlist> go run "d:\Coding-cf\Go\GoByExample\regular-expressions\regular-expressions.go"
true
true
peach
[0 5]
[peach ea]
[0 5 1 3]
[peach punch pinch]
[[0 5 1 3] [6 11 7 9] [12 17 13 15]]
[peach punch]
true
p([a-z]+)ch
a <fruit>
a PEACH

有关 Go 正则表达式的说明，请参考 [regexp 包文档](http://golang.org/pkg/regexp/)。

## 48-JSON

Go 提供内建的 **JSON 编码解码（序列化反序列化）支持**， 包括内建及自定义类型与 JSON 数据之间的转化。

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 下面我们将使用这两个结构体来演示自定义类型的编码和解码。
type response1 struct {
	Page   int
	Fruits []string
}

// 只有 可导出 的字段才会被 JSON 编码/解码。
// 必须以大写字母开头的字段才是可导出的。
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

// 首先我们来看一下基本数据类型到 JSON 字符串的编码过程。
// 这是一些原子值的例子。
func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// 这是一些切片和 map 编码成 JSON 数组和对象的例子。
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// JSON 包可以自动的编码你的自定义类型。
	// 编码的输出只包含可导出的字段，
	// 并且默认使用字段名作为 JSON 数据的键名。
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 你可以给结构字段声明标签来自定义编码的 JSON 数据的键名。
	// 上面 Response2 的定义，就是这种标签的一个例子。
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
	// fmt.Println(res2B)
	// fmt.Printf("%+v", res2D)

	// 现在来看看将 JSON 数据解码为 Go 值的过程。
	// 这是一个普通数据结构的解码例子。
	byt := []byte(`{"num": 6.13, "strs":["a", "b"]}`)

	// 我们需要提供一个 JSON 包可以存放解码数据的变量。
	// 这里的 map[string]interface{} 是一个键为 string，
	// 值为任意值的 map。
	var dat map[string]interface{}

	// 这是实际的解码，以及相关错误的检查。
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// 为了使用解码 map 中的值，我们需要将他们进行适当的类型转换。
	// 例如，这里我们将 num 的值转换成 float64 类型。
	num := dat["num"].(float64)
	fmt.Println(num)

	// 访问嵌套的值需要一系列的转化。
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// 我们还可以将 JSON 解码为自定义数据类型。
	// 这样做的好处是，可以为我们的程序增加附加的类型安全性，
	// 并在访问解码后的数据时不需要类型断言。
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 在上面例子的标准输出上，
	// 我们总是使用 byte和 string 作为数据和 JSON 表示形式之间的中介。
	// 当然，我们也可以像 os.Stdout 一样直接将 JSON 编码流
	// 传输到 os.Writer 甚至 HTTP 响应体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
```

PS D:\Coding-cf\Go\Go-algorithm\数据结构与算法之美\06_linkedlist> go run "d:\Coding-cf\Go\GoByExample\json\json.go"
true
1
2.34
"gopher"
["apple","peach","pear"]
{"apple":5,"lettuce":7}
{"Page":1,"Fruits":["apple","peach","pear"]}
{"page":1,"fruits":["apple","peach","pear"]}
map[num:6.13 strs:[a b]]
6.13
a
{1 [apple peach]}
apple
{"apple":5,"lettuce":7}

至此，我们已经学习了基本的 Go JSON 知识，如果想要获取更多的信息， 请查阅 [JSON and Go](http://blog.golang.org/2011/01/json-and-go.html) 博文 和 [JSON package docs](http://golang.org/pkg/encoding/json/)。

## 49-XML

Go 通过 `encoding.xml` 包为 XML 和 类-XML 格式提供了内建支持。

```go
package main

import (
	"encoding/xml"
	"fmt"
)

// Plant 结构将被映射到 XML 。 与 JSON 示例类似，
// 字段标签包含用于编码器和解码器的指令。
// 这里我们使用了 XML 包的一些特性：
// XMLName 字段名称规定了该结构的 XML 元素名称；
// id,attrr 表示 Id 字段是一个 XML 属性，而不是嵌套元素。
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id.attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// 传入我们声明了 XML 的 Plant 类型。 
	// 使用 MarshalIndent 生成可读性更好的输出结果。
	out, _ := xml.MarshalIndent(coffee, " ", " ")
	fmt.Println(string(out))

	// 明确的为输出结果添加一个通用的 XML 头部信息。
	fmt.Println(xml.Header + string(out))

	// 使用 Unmarshal 将 XML 格式字节流解析到 Plant 结构。 
	// 如果 XML 格式错误或无法映射到 Plant 结构，将返回一个描述性错误。
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"mexico", "California"}

	// parent>child>plant 字段标签告诉编码器，
	// 将 Plants 中的元素嵌套在 <parent><child> 里面。
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", " ")
	fmt.Println(string(out))

}
```

PS D:\Coding-cf\Go\Go-algorithm\数据结构与算法之美\06_linkedlist> go run "d:\Coding-cf\Go\GoByExample\xml\xml.go"
 <plant>
  <id.attr>27</id.attr>
  <name>coffee</name>
  <origin>Ethiopia</origin>
  <origin>Brazil</origin>
 </plant>
<?xml version="1.0" encoding="UTF-8"?>
 <plant>
  <id.attr>27</id.attr>
  <name>coffee</name>
  <origin>Ethiopia</origin>
  <origin>Brazil</origin>
 </plant>
Plant id=27, name=coffee, origin=[Ethiopia Brazil]
 <nesting>
  <parent>
   <child>
    <plant>
     <id.attr>27</id.attr>
     <name>coffee</name>
     <origin>Ethiopia</origin>
     <origin>Brazil</origin>
    </plant>
    <plant>
     <id.attr>81</id.attr>
     <name>Tomato</name>
     <origin>mexico</origin>
     <origin>California</origin>
    </plant>
   </child>
  </parent>
 </nesting>

## 50-时间

Go 为时间（time）和时间段（duration）提供了大量的支持；这儿有一些例子。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// 从获取当前时间时间开始。
	now := time.Now()
	p(now)

	// 通过提供年月日等信息，你可以构建一个 time。 
	// 时间总是与 Location 有关，也就是时区。
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// 你可以提取出时间的各个组成部分。
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 支持通过 Weekday 输出星期一到星期日。
	p(then.Weekday())

	// 这些方法用来比较两个时间，
	// 分别测试一下是否为之前、之后或者是同一时刻，精确到秒。
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// 方法 Sub 返回一个 Duration 来表示两个时间点的间隔时间。
	diff := now.Sub(then)
	p(diff)

	// 我们可以用各种单位来表示时间段的长度。
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 你可以使用 Add 将时间后移一个时间段，
	// 或者使用一个 - 来将时间前移一个时间段。
	p(then.Add(diff))
	p(then.Add(-diff))
}
```

PS D:\Coding-cf\Go\Go-algorithm\数据结构与算法之美\06_linkedlist> go run "d:\Coding-cf\Go\GoByExample\time\time.go"
2021-12-13 16:04:13.1625301 +0800 CST m=+0.008733101
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
true
false
false
105803h29m14.511142863s
105803.48736420635
6.348209241852381e+06
3.8089255451114285e+08
380892554511142863
2021-12-13 08:04:13.1625301 +0000 UTC
1997-10-23 09:05:44.140244374 +0000 UTC

## 51-时间戳

接下来，我们将研究与 Unix 纪元相关的概念。

一般程序会有获取 [Unix 时间](http://zh.wikipedia.org/wiki/UNIX%E6%97%B6%E9%97%B4) 的秒数，毫秒数，或者微秒数的需求。来看看如何用 Go 来实现。

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// 分别使用 time.Now 的 Unix 和 UnixNano， 
	// 来获取从 Unix 纪元起，到现在经过的秒数和纳秒数。
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// 注意 UnixMillis 是不存在的，
	// 所以要得到毫秒数的话， 你需要手动的从纳秒转化一下。
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// 你也可以将 Unix 纪元起的整数秒或者纳秒转化到相应的时间。
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
```

PS D:\Coding-cf\Go\Go-algorithm\数据结构与算法之美\06_linkedlist> go run "d:\Coding-cf\Go\GoByExample\epoch\epoch.go"
2021-12-13 16:10:53.8933231 +0800 CST m=+0.007848001
1639383053
1639383053893
1639383053893323100
2021-12-13 16:10:53 +0800 CST
2021-12-13 16:10:53.8933231 +0800 CST

## 52-时间的格式化和解析

下面我们将看看另一个时间相关的任务：时间解析与格式化。

Go 支持通过**基于描述模板的时间**格式化与解析。

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	// 这是一个遵循 RFC3339，
	// 并使用对应的 布局（layout）常量进行格式化的基本例子。
	t := time.Now()
	p(t.Format(time.RFC3339))

	// 时间解析使用与 Format 相同的布局值。
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)
	p(e)

	// Format 和 Parse 使用基于例子的布局来决定日期格式，
	// 一般你只要使用 time 包中提供的布局常量就行了，
	// 但是你也可以实现自定义布局。
	// 布局时间必须使用 Mon Jan 2 15:04:05 MST 2006 的格式，
	// 来指定 格式化/解析给定时间/字符串 的布局。
	// 时间一定要遵循：2006 为年，15 为小时，Monday 代表星期几等规则。
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// 对于纯数字表示的时间（时间戳），
	// 您还可以将标准字符串格式与提取时间值的一部分一起使用。
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	
	// 当输入的时间格式不正确时，Parse 会返回一个解析错误。
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
```

PS D:\Coding-cf\Go\Go-algorithm\数据结构与算法之美\06_linkedlist> go run "d:\Coding-cf\Go\GoByExample\time-formatting-parsing\time-formatting-parsing.go"
2021-12-13T16:26:19+08:00
2012-11-01 22:08:41 +0000 +0000
<nil>
4:26PM
Mon Dec 13 16:26:19 2021
2021-12-13T16:26:19.58813+08:00
0000-01-01 20:41:00 +0000 UTC
2021-12-13T16:26:19-00:00
parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"


## 53-随机数

Go 的 math/rand 包提供了[伪随机数](http://en.wikipedia.org/wiki/Pseudorandom_number_generator)生成器。

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 例如，rand.Intn 返回一个随机的整数 n，且 0 <= n < 100。
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// rand.Float64 返回一个64位浮点数 f，且 0.0 <= f < 1.0。
	fmt.Println(rand.Float64())

	// 这个技巧可以用来生成其他范围的随机浮点数， 例如，5.0 <= f < 10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。 
	// 要产生不同的数字序列，需要给定一个不同的种子。 
	// 注意，对于想要加密的随机数，使用此方法并不安全， 
	// 应该使用 crypto/rand。
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// 调用上面返回的 rand.Rand，就像调用 rand 包中函数一样。
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// 如果使用相同种子生成的随机数生成器，会生成相同的随机数序列。
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
```

PS D:\Coding-cf> go run "d:\Coding-cf\Go\GoByExample\random-numbers\random-numbers.go" 
81,87
0.6645600532184904
7.1885709359349015,7.123187485356329
56,60
5,87
5,87

有关 Go 提供的其他随机数的信息， 请参阅 [math/rand ](http://golang.org/pkg/math/rand/)包文档。

## 54-数字解析

从字符串中解析数字在很多程序中是一个基础常见的任务， 而在 Go 中，是这样处理的。

```go
package main

// 内建的 strconv 包提供了数字解析能力。
import(
	"fmt"
	"strconv"
)

func main() {

	// 使用 ParseFloat，这里的 64 表示解析的数的位数。
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// 在使用 ParseInt 解析整型数时， 
	// 例子中的参数 0 表示自动推断字符串所表示的数字的进制。 
	// 64 表示返回的整型数是以 64 位存储的。
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt 会自动识别出字符串是十六进制数。
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// ParseUint 也是可用的。
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi 是一个基础的 10 进制整型数转换函数。
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// 在输入错误时，解析函数会返回一个错误。
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
```

PS D:\Coding-cf> go run "d:\Coding-cf\Go\GoByExample\number-parsing\number-parsing.go" 
1.234
123
456
789
135
strconv.Atoi: parsing "wat": invalid syntax

## 55-URL解析

下面我们将了解一下另一个常见的解析任务：URL 解析。

URL 提供了[统一资源定位方式](http://adam.heroku.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/)。 这里展示了在 Go 中是如何解析 URL 的。

```go
package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// 我们将解析这个 URL 示例，它包含了一个 scheme、 
	// 认证信息、主机名、端口、路径、查询参数以及查询片段。
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// 解析这个 URL 并确保解析没有出错。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// 直接访问 scheme。
	fmt.Println(u.Scheme)

	// User 包含了所有的认证信息， 
	// 这里调用 Username 和 Password 来获取单独的值。
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// Host 包含主机名和端口号（如果存在）。
	// 使用 SplitHostPort 提取它们。
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// 这里我们提取路径和 # 后面的查询片段信息。
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// 要得到字符串中的 k=v 这种格式的查询参数，可以使用 RawQuery 函数。 
	// 你也可以将查询参数解析为一个 map。
	// 已解析的查询参数 map 以查询字符串为键， 
	// 已解析的查询参数会从字符串映射到到字符串的切片， 
	// 因此如果您只想要第一个值，则索引为 [0]。
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
```

运行我们的 URL 解析程序， 显示全部我们提取的 URL 的不同数据块。

PS D:\Coding-cf> go run "d:\Coding-cf\Go\GoByExample\url-parsing\url-parsing.go"       
postgres
user:pass
user
pass
host.com:5432
host.com
5432
/path
f
k=v
map[k:[v]]
v

## 56-SHA1哈希

[**SHA1 散列（hash）**](http://en.wikipedia.org/wiki/SHA-1) 经常用于生成二进制文件或者文本块的短标识。 例如，[git 版本控制系统](http://git-scm.com/) 大量的使用了 SHA1 来标识受版本控制的文件和目录。 这是 Go 中如何进行 SHA1 散列计算的例子。

```go
package main

// Go 在多个 crypto/* 包中实现了一系列散列函数。
import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	// 产生一个散列值的方式是 sha1.New()，sha1.Write(bytes)， 
	// 然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。
	h := sha1.New()

	// 写入要处理的字节。
	// 如果是一个字符串， 需要使用 []byte(s) 将其强制转换成字节数组。
	h.Write([]byte(s))

	// Sum 得到最终的散列值的字符切片。
	// Sum 接收一个参数， 可以用来给现有的字符切片追加额外的字节切片：
	// 但是一般都不需要这样做。
	bs := h.Sum(nil)

	// SHA1 值经常以 16 进制输出，例如在 git commit 中。 
	// 我们这里也使用 %x 来将散列结果格式化为 16 进制字符串。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
```

运行程序计算散列值，并以可读的 16 进制格式输出。

PS D:\Coding-cf> go run "d:\Coding-cf\Go\GoByExample\sha1-hashes\sha1-hashes.go"       
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94

你可以使用和上面相似的方式来计算其他形式的散列值。 *例如，计算 MD5 散列，引入 crypto/md5 并使用 md5.New() 方法。*

注意，如果你需要密码学上的安全散列，你需要仔细的研究一下 [加密散列函数](http://en.wikipedia.org/wiki/Cryptographic_hash_function)。

## 57-Base64编码

Go 提供了对 [base64 编解码](http://zh.wikipedia.org/wiki/Base64)的内建支持。

```go
package main

// 这个语法引入了 encoding/base64 包， 
// 并使用别名 b64 代替默认的 base64。这样可以节省点空间。
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// 这是要编解码的字符串。
	data := "abc123!?$*&()'-=@~"

	// Go 同时支持标准 base64 以及 URL 兼容 base64。 
	// 这是使用标准编码器进行编码的方法。 
	// 编码器需要一个 []byte，因此我们将 string 转换为该类型。
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 解码可能会返回错误，如果不确定输入信息格式是否正确， 
	// 那么，你就需要进行错误检查了。
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// 使用 URL base64 格式进行编解码。
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
```

标准 base64 编码和 URL base64 编码的 编码字符串存在稍许不同（后缀为 + 和 -）， 但是两者都可以正确解码为原始字符串。

PS D:\Coding-cf> go run "d:\Coding-cf\Go\GoByExample\base64-encoding\base64-encoding.go"
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~


## 58-读文件

读写文件在很多程序中都是必须的基本任务。 首先我们来看一些读文件的例子。

```go
package main

import(
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读取文件需要经常进行错误检查， 这个帮助方法可以精简下面的错误检查过程。
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 最基本的文件读取任务或许就是将文件内容读取到内存中。
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// 您通常会希望对文件的读取方式和内容进行更多控制。 
	// 对于这个任务，首先使用 Open 打开一个文件，以获取一个 os.File 值。
	f, err := os.Open("/tmp/dat")
	check(err)

	// 从文件的开始位置读取一些字节。 
	// 最多允许读取 5 个字节，但还要注意实际读取了多少个。
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// 你也可以 Seek 到一个文件中已知的位置，并从这个位置开始读取。
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// 例如，io 包提供了一个更健壮的实现 ReadAtLeast，用于读取上面那种文件。
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 没有内建的倒带，但是 Seek(0, 0) 实现了这一功能。
	_, err = f.Seek(0, 0)
	check(err)

	// bufio 包实现了一个缓冲读取器，
	// 这可能有助于提高许多小读操作的效率，以及它提供了很多附加的读取函数。
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// 任务结束后要关闭这个文件 
	// （通常这个操作应该在 Open 操作后立即使用 defer 来完成）。
	f.Close()
}
```
PS D:\Coding-cf\Go\GoByExample\reading-files> echo "hello" > /tmp/dat
PS D:\Coding-cf\Go\GoByExample\reading-files> echo "go" >> /tmp/dat
PS D:\Coding-cf\Go\GoByExample\reading-files> go run "d:\Coding-cf\Go\GoByExample\reading-files\reading-files.go"
��hello
go
5 bytes: ��he
2 bytes @ 6: l
2 bytes @ 6: l
5 bytes: ��he

## 59-写文件

下面我们来看一下写入文件。
在 Go 中，写文件与我们前面看过的读文件方法类似。

```go
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 开始！这里展示了如何写入一个字符串（或者只是一些字节）到一个文件。
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// 对于更细粒度的写入，先打开一个文件。
	f, err := os.Create("/tmp/dat2")
	check(err)

	// 打开文件后，一个习惯性的操作是：立即使用 defer 调用文件的 Close。
	defer f.Close()

	// 您可以按期望的那样 Write 字节切片。
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// WriteString 也是可用的。
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// 调用 Sync 将缓冲区的数据写入硬盘。
	f.Sync()

	// 与我们前面看到的带缓冲的 Reader 一样，
	// bufio 还提供了的带缓冲的 Writer。
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// 使用 Flush 来确保，已将所有的缓冲操作应用于底层 writer。
	w.Flush()
}
```

运行这段文件写入代码。

PS D:\Coding-cf\Go\GoByExample\reading-files> go run "d:\Coding-cf\Go\GoByExample\writing-files\writing-files.go"
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

然后检查写入文件的内容。

PS D:\Coding-cf\Go\GoByExample\reading-files> cat /tmp/dat1
hello
go
PS D:\Coding-cf\Go\GoByExample\reading-files> cat /tmp/dat2
some
writes
buffered

## 60-行过滤器

我们刚刚看到了文件 I/O 思想， 接下来，我们看看它在 stdin 和 stdout 流中的应用。

**行过滤器（line filter）** 是一种常见的程序类型， 它读取 stdin 上的输入，对其进行处理，然后将处理结果打印到 stdout。 grep 和 sed 就是常见的行过滤器。

这里是一个使用 Go 编写的行过滤器示例，*它将所有的输入文字转化为大写的版本。* 你可以使用这个模式来写一个你自己的 Go 行过滤器。

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// 用带缓冲的 scanner 包装无缓冲的 os.Stdin， 
	// 这为我们提供了一种方便的 Scan 方法， 
	// 将 scanner 前进到下一个 令牌（默认为：下一行）。
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		// Text 返回当前的 token，这里指的是输入的下一行。
		ucl := strings.ToUpper((scanner.Text()))

		// 输出转换为大写后的行。
		fmt.Println(ucl)
	}

	// 检查 Scan 的错误。 
	// 文件结束符（EOF）是可以接受的，它不会被 Scan 当作一个错误。
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
```

试一下我们的行过滤器，首先创建一个有小写行的文件。

PS D:\Coding-cf\Go\GoByExample\reading-files> echo 'hello' > /tmp/lines
PS D:\Coding-cf\Go\GoByExample\reading-files> echo 'filter' >> /tmp/lines

然后使用行过滤器来得到大写的行。

PS D:\Coding-cf\Go\GoByExample\line-filters> cat /tmp/lines | go run line-filters.go   
HELLO
FILTER

## 61-文件路径

## 62-目录

## 63-临时文件和目录

## 64-单元测试

## 65-命令行参数

## 66-命令行标志

## 67-命令子命令

## 68-环境变量

## 69-HTTP客户端

## 70-Context

## 71-生成进程

## 72-执行进程

## 73-信号

## 74-退出

