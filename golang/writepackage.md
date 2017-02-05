---
keywords: yupae.cn
title: 设计Go包
---

# 设计Go包      

### 这里以第三方包[goptions](https://github.com/polaris1119/goptions)为例说明：（只说核心数据结构）

1、定义数据结构（struct等），如FlagSet、Flag，以及数据结构对应的方法。（这里一般会提供实例化数据结构的方法，比如：NewFlagSet())
一般地，依赖这些就可以供外部使用（一般会就该包的功能提供一个外部可用的入口）
比如，外部可以这么使用goptions这个包：

    `options := struct{}{}
flagSet := goptions.NewFlagSet(filepath.Base(os.Args[0]), &amp;options)
flagSet.Parse(os.Args[1:])
…
`

相当于自己实现了goptions包中的Parse()或ParseAndFail()

这样，每个调用者都得这么实现一次。

2、定义函数，提供给外部方便调用的接口；
即：将上面的封装成函数供使用者直接使用。
goptions提供了两种封装：

*   1）普通封装-&gt;Parse()，错误需要自己处理；
*   2）完全封装-&gt;ParseAndFail()，错误都处理好，比显示出友好的错误提示

当然，这里之所以处理错误，正如文档中说的，出错应该是编译期错误，而不是在运行时报出来

3、另外，为了区分不同的错误提示不一样，包中经常会自定义自己的错误类型，比如：

`
var ErrHelpRequest = errors.New(“Request for Help”)
`

### 包的其他方面

1、测试用例
以_test.go为后缀，测试方法包含唯一参数：

`
t *testing.T
`

2、跨平台
以 _$GOOS.go为后缀。如：

*   windows平台以 _windows.go为后缀
*   Linux平台以 _linux.go为后缀
*   Mac平台以_darwin.go

3、为包提供例子
以example_test.go为文件名，包名为：要包名加上_test，函数名为Example加上实际的函数。如：
为time包中的After()函数提供例子
文件名：example_test.go，包名：time_test，函数：ExampleAfter()

上面是建议的命名规则，实际上，两处是必须的，其他的地方无所谓，但建议标准库中的方式写。必须的两处是：

*   ①文件名必须以_test.go为后缀（这和测试用例一样）
*   ②为哪个函数提供例子，必须以Example开头，后接函数名

另外，如果函数名最后加上_xxx，则文档中，xxx会变为(xxx)
在 Example 函数尾部用”// Output: ” 来标注正确输出结果

    