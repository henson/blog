---
keywords: yupae.cn
title: 深入学习Go包
---

# 深入学习Go包      

学习一门语言，熟悉语言语法、规范等之后，应该学习语言的标准库。在Python中，会有一些函数来探究包的内容。在Go中，更多的是通过查看Go标准库文档来学习。

实际写代码中，肯定需要用到很多标准库中的包，在学习阶段，可以在需要用某个包时，彻底学习这个包，掌握它。标准库中每个包的文档是学习包最好的资料，一定要仔细看明白。

Go包具体该怎么学了？以下是我自己的学习方法，仅供参考（以time包为例）

### 1、看文档中的Overview，整体上对该包有一个了解

从这知道，该包用于处理和显示日期和时间。日期都是公历。

### 2、看Index，关注函数和类型（函数是指不属于某种类型的func），不需要关注类型的方法（Variable和Constants需要大概关注）

*   定义了一些常量，表示时间格式
*   定义了Duration、Location、Month、ParseError、Ticker、Time、Timer和Weekday等类型，从名字可以很容易的知道这些类型所代表的含义
*   有After、Sleep和Tick三个全局函数

其实全局函数还有其他的，但是其他全局函数在文档中放在了某种类型的下面，表示该全局函数会生成一个该类型的实例，这个可以当做类型的方法来研究

这一步，我们可以不关心这些类型该怎么使用，我们关心全局函数怎么用

*   func After(d Duration) 在【Index】视图 点击该函数，跳转到该函数的说明处。函数注释已经说的很清楚：指定的时间过去之后，将当前发送到chanel中返回，和NewTimer(d).C功能相同。另外，还提供了使用示例。
*   func Sleep(d Duration)很明显，暂停当前goroutine一段时间
*   func Tick(d Duration) 这是包设计中 设计函数 常用的方式：方便包的使用者，直接使用Tiker类型，而不需要另外实例化再调用其方法。Tick方法适用于那些需要使用Tiker，但是永远不需要停止的场景。

以上三个函数文档都提供了对应的例子

### 3、关注类型及其方法 （以Duration为例）

1）看类型的定义（结构）

`
type Duration int64
`

类型注释说，Duration表示两个时间的间隔，单位是纳秒（nanosecond）

该类型预定义了一些常量，这样方便类型转换、计算以及阅读。另外还提供了一些示例。

2）实例化方式

由于Duration是int64的别名，整型字面值常量可以直接赋值给Duration。对于整型类型的变量，则需要类型转换

比如，Sleep函数接收Duration类型的参数。如果想Sleep 2秒，这样写：

`
Sleep(2 * time.Second)  //或者Sleep(2e9)。
`

或者

`
i := 2； Sleep(time.Duration(i) * time.Second)
`

另外，ParseDuration可以将字符串转为Duration；Since()也可以获取Duration的实例

3）提供的方法

分别转为时、分、纳秒、秒。除了纳秒，其他都是float64。

这里说一下String()这个方法：

△如果某种类型定义了String()方法，那么，fmt包格式化时，直接传入类型，内部会调用类型的String()方法进行格式化，类似于Java中toString方法

### 4、自己写几个例子练练手

比如，想要格式化输出当前时间，格式为：2012-11-27 15:10:00

`
time.Now().Format(“2006-01-02 15:04:05″)
`

注意，2006这样的必须固定，不能为2007什么的。（不知道为什么这么设计）
设计了一个包实现类似PHP中date()和strtotime()的功能

[https://github.com/polaris1119/times](https://github.com/polaris1119/times)

### 5、学习第三方包（几种生成文档的方式）

1.  将包放入$GOROOT/src/pkg/ 目录下，godoc 时会生成该包的文档（当做标准库包了）
2.  设定GOPATH，将第三方包放在GOPATH中
3.  使用godoc命令时，通过path指定第三方包的路径
4.  直接通过godoc命令在命令行看起文档（不直观）
5.  [http://go.pkgdoc.org/](http://go.pkgdoc.org/)在这个网站搜索包，有的话，会提供文档（这个网站会搜索github上的资源）

    