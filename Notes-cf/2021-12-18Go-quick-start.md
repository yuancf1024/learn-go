# 2021-12-18Go-quick-start

[toc]

## Readme

参考文章: [【一文Go起来】快速上手篇](https://mp.weixin.qq.com/s/rEBfg_RIJtxfJkyE30_D7g) 
作者: 牛牛码特 (Wechat-official)

## 概要

Golang是云原生时代的宠儿，它最大的优点在于*简单有效*，简单在于*上手迅速、代码规范、部署方便*；有效在于它能*很容易写出高并发的代码，处理能力强*。

Golang能适用于**web后台、数据库、云原生、区块链**等大多数场景，大厂与其相关的招聘岗位也在逐年增加，因此，学习Golang这样相对较新、发展前景很好的语言，我们是可以*实现弯道超车的*。

牛牛也秉承Golang简单、有效的理念推出一份golang学习套餐，本文是其中的快速上手篇，每个可执行代码也都附上了运行结果，希望小伙伴们读完此文，自己动手试一试，*实现快速入门，用Golang开启新的旅程*。

下面我们就从最基础的环境部署开始，开启我们的Golang之旅吧~

## 环境准备

由官网的安装介绍，我们可以了解到各个系统的安装流程，对Linux来说：

**Linux 安装方式**

1. 下载安装包

下载安装包到当前目录

2. 解压到指定目录

rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.2.linux-amd64.tar.gz

3. 设置环境变量PATH

export PATH=$PATH:/usr/local/go/bin

4. 检查Go版本

go version

可以看到，Linux安装只用下载安装包，并解压到特定目录，设置PATH环境变量之后即完成安装。

**Mac 安装**

Mac更加简单粗暴，直接下载安装包，点击安装。

**Windows安装**

Windows和Mac一样，直接点击安装包进入安装界面即可。

Golang包官方的资源地址是在：https://go.dev/dl/，小伙伴们可以上去选择自己需要的版本，通常来说，建议是下载最新版本。

**环境变量设置**

Golang有一个环境变量`GOPATH`，这个变量*表示第三方依赖包在本地的位置*，大家指定一个方便访问的路径即可。

这样*第三方依赖包都可以下载到GOPATH下面，项目也可以自动从GOPATH加载第三方依赖包。*

**IDE推荐**

推荐GoLand，功能强大，开箱即用，还有完善的插件生态。习惯用vim在linux下编程的同学也请留步，GoLand可以非常方便的安装vim插件，可以同时享受IDE的强大功能和vim高效的编辑能力。

## 语法介绍

### 包的概念

```go
package main

import "fmt"

func main() {
    fmt.Println("niuniumart")
}
```

**Output:**

PS D:\gocf\src\github.com> go run "d:\gocf\src\github.com\yuancf1024\learn-go\Notes-cf\Go-quick-start\main.go"
niuniumart

以上代码是组成一个可执行代码最基础的三部分，换言之，*每个可执行代码都必须包含Package、import以及function这三个要素。*

Golang**以包来管理代码**，一个目录承载一个包的内容，代码文件必须在一个包下面，比如这里我们在`code`目录下建了一个`main.go`文件，`package`指示代码是属于`main`这个包的。main函数必须要在main包下面。import用来引用外部的包，如上面示例中import引用了fmt包，就可以直接使用其方法`fmt.Println`。

包管理工具有三种：

1. `GOPATH`：把依赖包通过go get命令拉到本地GOPATH目录下，缺点是没法实现依赖包多版本管理。

2. DEP：将依赖包通过DEP命令打包到工程下的vendor目录。Shopee金融团队、字节跳动教育团队用的就是DEP；

3. GoMod：将依赖包拉取到统一的pkg目录下，分版本存储。腾讯云用GoMod的团队会比较多。

针对包管理，本文我们就不做过多扩展，后续有文章会进行专门的讲解。

回到我们的例子，针对这个main.go文件，进行如下操作，即可运行程序：

```shell
go build main.go //得到二进制文件main
./main //执行代码
```

### 变量定义及初始化

```go
package main

import "fmt"

var globalStr string
var globalInt int

func main() {
	var localStr string
	var localInt int
	localStr = "first local"
	localInt = 2021
	globalInt = 1024
	globalStr = "first global"
	fmt.Printf("globalStr is %s\n", globalStr) // globalStr is first global
	fmt.Printf("globalInt is %d\n", globalInt) // globalInt is 1024
	fmt.Printf("localStr is %s\n", localStr) // localStr is first local
	fmt.Printf("localInt is %d\n", localInt) // localInt is 2021

}
```

**Output:**

PS D:\gocf\src\github.com> go run "d:\gocf\src\github.com\yuancf1024\learn-go\Notes-cf\Go-quick-start\variable.go"
globalStr is first global
globalInt is 1024      
localStr is first local
localInt is 2021

上面的代码定义了以下四个变量：

一个名字叫globalStr的*全局字符串变量*; 

一个名字叫globalInt的*全局整型变量*；

一个名字叫localStr的*局部字符串变量*; 

一个名字叫localInt的*局部整型变量*；

> 注意，这里的全局变量如果要在包外访问，首字母需要大写，对，你没有看错，golang是以首字母大小写来区分对包外是否可见。

### 数组&切片&字典

```go
package main

import "fmt"

func main() {

	// 数组初始化
	var strAry = [10]string{"aa", "bb", "cc", "dd", "ee"}
	// 切片初始化
	var sliceAry = make([]string, 0)
	sliceAry = strAry[1:3]
	// 字典初始化
	var dic = map[string]int {
		"apple": 1,
		"watermelon": 2,
	}

	fmt.Printf("strAry %+v\n", strAry)
	fmt.Printf("sliceAry %+v\n", sliceAry)
	fmt.Printf("dic %+v\n", dic)
}
```

**Output:**

PS D:\gocf\src\github.com> go run "d:\gocf\src\github.com\yuancf1024\learn-go\Notes-cf\Go-quick-start\array-slice-dic.go"
strAry [aa bb cc dd ee     ]
sliceAry [bb cc]
dic map[apple:1 watermelon:2]

以上代码演示了数组、切片、字典的定义及初始化。可以看到*切片通过索引的方式指向了数组***。切片是可以更改某个元素内容的，数组则不能，在开发中，主要都是使用切片来进行逻辑处理。**

### 条件选择语法

```go
package main

import "fmt"

func main() {

	localStr := "case3" // 是的,还可以通过 := 这种方式直接初始化基础变量
	if localStr == "case3" {
		fmt.Printf("into true logic\n")
	} else {
		fmt.Printf("into false logic\n")
	}

	// 字典初始化
	var dic = map[string]int {
		"apple": 1,
		"watermelon": 2,
	}

	if num, ok := dic["orange"]; ok {
		fmt.Printf("orange num %d\n", num)
	}

	if num, ok := dic["watermelon"]; ok {
		fmt.Printf("watermelon num %d\n", num)
	}

	switch localStr {
	case "case1":
		fmt.Println("case1")
	case "case2":
		fmt.Println("case2")
	case "case3":
		fmt.Println("case3")
	default:
		fmt.Println("default")
	}
}
```

**Output:**

PS D:\gocf\src\github.com> go run "d:\gocf\src\github.com\yuancf1024\learn-go\Notes-cf\Go-quick-start\condition.go"
into true logic
watermelon num 2
case3

if语句在Golang和其他语言中的表现形式一样，没啥区别。上面的例子同时也展示了用if判断某个key在map是否为空的写法。

*switch中，每个case都默认break。*即如果是case1，那么执行完之后，就会跳出switch条件选择。*如果希望从某个case顺序往下执行，可以使用fallthrough关键字。*

### 循环写法

```go
package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Printf("current i %d\n", i)
	}

	j := 0
	for {
		if j == 5 {
			break
		}
		fmt.Printf("current j %d\n", j)
		j++
	}

	var strAry = []string{"aa", "bb", "cc", "dd", "ee"} // 是的,不指定初始个数也ok

	// 切片初始化
	var sliceAry = make([]string, 0)
	sliceAry = strAry[1:3]
	for i, str := range sliceAry {
		fmt.Printf("slice i %d, str %s\n", i, str)
	}

	// 字典初始化
	var dic = map[string]int {
		"apple": 1,
		"google": 2,
		"facebook": 3,
		"amazon": 4,
	}

	for k, v := range dic {
		fmt.Printf("key %s, value %d\n", k, v)
	}
}
```

**Output:**

PS D:\gocf\src\github.com> go run "d:\gocf\src\github.com\yuancf1024\learn-go\Notes-cf\Go-quick-start\circulation.go"
current i 0
current i 1
current i 2
current i 3
current i 4
current j 0
current j 1
current j 2
current j 3
current j 4
slice i 0, str bb
slice i 1, str cc
key apple, value 1
key google, value 2
key facebook, value 3
key amazon, value 4

## 语言特性

### 协程（goroutine）

**协程是Golang最重要的一个特性。**

在协程出现之前，线程被作为调度的最小单位。*协程可以理解是一种用户态，逻辑层面的线程。*

通过协程，我们将很容易地实现高并发：假如你要做三件事，假设要执行a，b，c三个方法。代码该怎么写？平常我们的写法就是：

```go
package main

import (
	"fmt"
	"time"
)

func a() {
	time.Sleep(3 * time.Second)
	fmt.Println("it's a")
}

func b() {
	time.Sleep(2 * time.Second)
	fmt.Println("it's b")
}

func c() {
	time.Sleep(1 * time.Second)
	fmt.Println("it's c")
}

func main() {

	start := time.Now()

	fmt.Println("程序开始:")
	a()
	b()
	c()
	time.Sleep(1 * time.Second)
	fmt.Println("执行完毕!")
	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
```

以上的代码只有a做完了，才能做b，b做完了，才能做c。

但Golang语言层面支持协程，通过关键字go，后面跟一个方法，就生成了一个协程：

```go
package main

import (
	"fmt"
	"time"
)

func a() {
	time.Sleep(3 * time.Second)
	fmt.Println("it's a")
}

func b() {
	time.Sleep(2 * time.Second)
	fmt.Println("it's b")
}

func c() {
	time.Sleep(1 * time.Second)
	fmt.Println("it's c")
}

func main() {

	start := time.Now()

	fmt.Println("程序开始:")
	go a()
	go b()
	go c()
	time.Sleep(1 * time.Second)
	fmt.Println("执行完毕!")
	time.Sleep(2 * time.Second) // 为了等待3秒,以免a() b() 还没运行就结束程序了
	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
```

**Output:**

PS D:\gocf\src\github.com> go run "d:\gocf\src\github.com\yuancf1024\learn-go\Notes-cf\Go-quick-start\concurrency-goroutine.go"
程序开始:
执行完毕!
it's c
it's b
it's a
cost=[3.0133822s]

在协程里，三个方法就可以并发进行，可以看到，由于方法a执行时间最久，所以最后才打印。协程Golang运行时调度，是充分利用了Golang多核的性能。后续文章牛牛会专门深入讲解协程的原理，我们现在作为入门者，只需要会使用它即可。

小伙伴们也可以想想，牛牛为何要在a，b，c三个方法之后还要sleep5秒，这里先留个悬念。

> 实际上程序本身运行耗时1s,再留2s睡眠就足够啦.

### 通道（channel）

通道的要点：

1. 类似unix中管道（pipe），*先进先出*；

2. *线程安全，多个goroutine同时访问，不需要加锁*；

3. *channel是有类型的，一个整数的channel只能存放整数*。

通道的定义：

```go
var ch0 chan int
var ch1 chan string
var ch2 chan map[string]string

type stu struct{}

var ch3 chan stu
var ch4 chan *stu
```

通道可以用于协程之间数据的传递，一般分为**有缓冲通道**和**无缓冲通道**。

两个协程间如果有数据交流怎么办？这时候就可以用通道来传递。**Golang的设计思想就是用通信代替共享内存。**

```go
package main

import (
	"fmt"
	"time"
)

var ch chan int

func a() {
	time.Sleep(3 * time.Second)
	a := 5
	ch <- a
	fmt.Println("out of a")
}

func b() {
	time.Sleep(1 * time.Second)
	fromA := <- ch
	b := fromA + 3
	fmt.Println("b is ", b)
}

func main() {
	ch = make(chan int, 1)
	go a()
	go b()
	time.Sleep(5 * time.Second)
	fmt.Println("out of main")
}
```

**Output:**

PS D:\gocf\src\github.com\yuancf1024\web> go run "d:\gocf\src\github.com\yuancf1024\learn-go\Notes-cf\Go-quick-start\channel.go"
out of a
b is  8
out of main

可以看到，更慢一些的b，是等管道有数据才继续运行，并成功拿到了a往管道里放入的数字5！这就完成了协程间的通信。


另外，这里也涉及到一个**面试高频问题**：有缓冲和无缓冲通道的区别？

*通道可以带缓冲，就是说可以往通道里放多个数据，放满了，才会阻塞。*

有一段时间，牛牛一直误以为无缓冲通道就是容量为1的有缓冲通道，于是就以此为例来进行讲解：

```go
chSync := make(chan int) // 无缓冲
chAsyn := make(chan int, 1) // 有缓冲
```

同样是向通道里塞一个数据：chSync <-1

**无缓冲场景**：*一直要等有别的协程通过<-chSync接手了这个参数，那么chSync<-1才会继续下去，要不然就一直阻塞着。*

**有缓冲场景**：*chAsyn<-1则不会阻塞，因为缓冲大小是1，只有当放第二个值的时候，第一个还没被人拿走，这时候才会阻塞。*

仔细理解下，实际这就是同步和异步的区别，**无缓冲一定是同步等待**，**有缓冲只有在缓冲满了，异步又处理不过来的时候，才会阻塞**。

**无缓冲**

```go
package main

import (
	"fmt"
	"time"
)

var ch chan int

func a() {
	time.Sleep(3 * time.Second)
	a := 5
	ch <- a
	fmt.Println("out of a")
}

func b() {
	time.Sleep(1 * time.Second)
}

func main() {
	ch = make(chan int) // 无缓冲管道
	go a()
	go b()
	time.Sleep(5 * time.Second)
	fmt.Println("out of main")
}
```

**Output:**

PS D:\gocf\src\github.com\yuancf1024\web> go run "d:\gocf\src\github.com\yuancf1024\learn-go\Notes-cf\Go-quick-start\channel-no-buffer.go"
out of main

可以看到，在没有接盘侠的情况下，a在写管道时被阻塞了。

**有缓冲**

```go
package main

import (
	"fmt"
	"time"
)

var ch chan int

func a() {
	time.Sleep(3 * time.Second)
	a := 5
	ch <- a
	fmt.Println("out of a")
}

func b() {
	time.Sleep(1 * time.Second)
}

func main() {
	ch = make(chan int, 1) // 有缓冲管道
	go a()
	go b()
	time.Sleep(5 * time.Second)
	fmt.Println("out of main")
}
```

**Output:**



可以看到，函数a往管道写入一个数据，即使没有消费者，也并未阻塞。

PS D:\gocf\src\github.com\yuancf1024\web> go run "d:\gocf\src\github.com\yuancf1024\learn-go\Notes-cf\Go-quick-start\channel-buffer.go"
out of a
out of main

### 接口(interface)

Go 语言提供了一种特别的数据类型——**接口**，*它把所有具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。*

话不多说，看看🌰：

```go
package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

// type rect
type Rect struct {
	height float64
	weight float64
}

func (p *Rect) Area() float64 {
	return p.height * p.weight
}

func (p *Rect) Perimeter() float64 {
	return 2 * (p.height + p.weight)
}

func main() {
	var s Shape = &Rect{height: 10, weight: 8}
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}
```

**Output:**

PS D:\gocf\src\github.com\yuancf1024\web> go run "d:\gocf\src\github.com\yuancf1024\learn-go\Notes-cf\Go-quick-start\interface.go"
80
36

代码中Shape就是一个接口，声明了两个方法：面积（Area）和周长（Perimeter）。

咱们定义了一个具体结构Rect，实现这个接口。可以看到，用基础的Shape接口，可以一个指向Rect对象，并调用其方法。

接口提供了面向对象编程的能力，如果你掌握多种语言，比如Golang、C++、Java等等，那么一定会问**Golang的多态和C++的多态有什么区别**（**使用相同类型的引用，指向不同类型对象，即多态**）。

答案就是*C++或者Java是需要主动声明基础类*，而*Golang，只需要实现某个interface的全部方法，那么就是实现了该类型。*所以，Golang的继承关系是**非侵入式的**，这也是Golang的特色与优点。

## 单元测试介绍

为了保证代码的质量，很多公司都会要求写单元测试。这里介绍单元测试的两个常用指标：

1. **函数覆盖率**：被调用到的函数个数/总函数个数，通常要求100%；

2. 行覆盖率：被调用到的行数/总行数，通常要求>60%。

通过单元测试，我们可以针对不同场景测试代码，是研发自己对质量的把控。

牛牛之前在**字节跳动SaaS**化部门，没有专门的测试人员，对单元测试的要求就非常高，行覆盖率需要达到80%。


**go test**

- go的test一般以xxx_test.go为文件名，xxx并没有特别要求必须是要实测的文件名；

- TestMain作为初始化test；

- Testxxx(t* testing.T)；

- go test即可运行单元测试；

- go test --v fileName --test.run funcName可以指定单测某个方法。

我们来创建一个main_test.go文件进行示例，main.go文件就使用上面的interface例子，包结构如下：

```shell
├── main.go
├── main_test.go
```

`main_test.go`

```go
package main

import "testing"

func TestRect(t *testing.T) {
	var s Shape = &Rect{height: 10, weight: 8}
	if s.Area() != 80 {
		t.Errorf("area %f\n", s.Area())
	}

	if s.Perimeter() != 30 {
		t.Errorf("perimeter %f\n", s.Perimeter())
	}
}
```

使用`go test-v`

由于周长Perimeter不符合预期，则会有如下提示：

```shell
PS D:\gocf\src\github.com\yuancf1024\learn-go\Notes-cf\Go-quick-start\test> go test -v
=== RUN   TestRect
    main_test.go:12: perimeter 36.000000
--- FAIL: TestRect (0.00s)
FAIL
exit status 1
FAIL    Notes-cf/Go-quick-start/test    0.370s
```

**go convey**

go convey可以很好的支持setup和teardown，它可以在运行单个测试用例前都进行一次状态初始化，在结束后再进行销毁。这样如果有多个子用例，可以复用同一套初始化环境。

go convey还有很多已经定义好，能够直接使用的assert函数，并且还可以自定义assert函数。

常用的assert如下：

```go

var (
    ShouldEqual          = assertions.ShouldEqual
    ShouldNotEqual       = assertions.ShouldNotEqual
    ShouldBeGreaterThan          = assertions.ShouldBeGreaterThan
    ShouldBeGreaterThanOrEqualTo = assertions.ShouldBeGreaterThanOrEqualTo
    ShouldBeLessThan             = assertions.ShouldBeLessThan
    ShouldBeLessThanOrEqualTo    = assertions.ShouldBeLessThanOrEqualTo
    ShouldBeBetween              = assertions.ShouldBeBetween
    ShouldNotBeBetween           = assertions.ShouldNotBeBetween
    ShouldBeBetweenOrEqual       = assertions.ShouldBeBetweenOrEqual
    ShouldNotBeBetweenOrEqual    = assertions.ShouldNotBeBetweenOrEqual
    ShouldContainSubstring    = assertions.ShouldContainSubstring
    ShouldNotContainSubstring = assertions.ShouldNotContainSubstring
    ShouldPanic        = assertions.ShouldPanic
    ShouldBeError      = assertions.ShouldBeError
)
```

使用举例：

```go

package main

import (
 "testing"

 "github.com/smartystreets/goconvey/convey"
)

func TestRect(t *testing.T) {
    convey.Convey("TestRect", t, func() {
        var s Shape = &Rect{height: 10, weight: 8}
        convey.So(s.Area(), convey.ShouldEqual, 80)
        convey.So(s.Perimeter(), convey.ShouldEqual, 30)
    })
}
```

由于Perimeter不符合预期，会出现如下提示：

用convey做断言，是不是更清晰明了了。

## 用ORM连接数据库

**什么是ORM?**

ORM的全称是：Object Relational Mapping(对象关系映射)，其主要作用是在编程中，把面向对象的概念跟数据库中表的概念对应起来。

举例来说就是我们*定义一个对象，那就对应着一张表，这个对象的实例，就对应着表中的一条记录。*

GORM使用示例：

```go

package main

import (
"fmt"
   "github.com/jinzhu/gorm" 
   _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
    Name string
    Age int
}

func main() {

    username := ""
    pwd := ""
    addr := "" //ip:port
    database := ""
    args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, pwd, addr, database)

    // step1 : 连接数据库
    db, err := gorm.Open("mysql", args) 
    if err != nil {
        fmt.Println(err)
        //do something
        return
    }
    defer db.Close()

    // step2 : 插入一行记录
    user := User{Name: "niuniu", Age: 18}
    err = db.Create(&user)
    if err != nil {
        fmt.Println(err)
        return
    } 

    // step3 ：查询记录
    var tmpUser User
    err = db.Where("name = ?", "niuniu").First(&tmpUser).Error //查询User并将信息保存到tmpUser
    if err != nil {
        fmt.Println(err)
        return
    } 
    fmt.Println(tmpUser)
}
```

最终结果:

{niuniu 18}

## 以一个web server结束

**最简化样例**

Golang http server有几种写法，这里介绍最简单一种，让我们看看到底有多简单：这里我们实现一个SayHello接口，访问该接口，会以“hello"字符串回包。

```go
package main

import (
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello")) // 以字符串"hello"作为返回包
}

func main() {
	http.HandleFunc("/say_hello", sayHello)
	err := http.ListenAndServe(":8080", nil) // 开启一个http服务
	if err != nil {
		log.Print("ListenAndServe:", err)
		return
	}
}
```

**用框架来一发**

在实际开发中，很少会直接用http裸写sever，因为如果进行功能的完善，比如可插拔中间件实现，最终就是自己实现了框架，而实际开发中，我们会选择久经考验的完善框架，比如gin：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SayHello(c *gin.Context) {
	c.String(http.StatusOK, "hello") // 以字符串"hello"作为返回包
}

func main() {
	engine := gin.Default() // 生成一个默认的gin引擎
	engine.Handle(http.MethodGet, "/say_hello", SayHello)
	err := engine.Run(":8080") // 使用8080端口号,开启一个web服务
	if err != nil {
		log.Print("engine run err: ", err.Error())
	}
	return
}
```

## 小结

至此，Golang的基本玩法，大家有所了解了吗？

希望Go起来这个系列的文章可以帮助大家快速入门，尽快投入开发，但如果要成为资深的Golang开发者，还需要针对细节，做深入研究。
