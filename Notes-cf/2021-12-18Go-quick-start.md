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

### 接口(interface)

**无缓冲**

**有缓冲**

## 单元测试介绍


## 用ORM连接数据库

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