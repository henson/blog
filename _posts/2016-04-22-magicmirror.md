---
layout: post
title: 用树莓派3制作智能镜子（MagicMirror）
date: 2016-04-22
categories: 树莓派
cover: 'http://www.yupae.cn/images/magicmirror.jpeg'
tags: 树莓派 Raspberry MagicMirror
---

### 项目地址

[henson/MagicMirror](https://github.com/henson/MagicMirror)

根据前面的文章，树莓派3已经配置好了，我们拿它来做点什么呢？这里有几篇文章不妨先看下：

### 相关链接

*   [Magic Mirrors around the World!](http://michaelteeuw.nl/post/111886383522/magic-mirrors-around-the-world) via Xonay Labs , [Michael Teeuw](http://michaelteeuw.nl/tagged/magicmirror) 项目地址：[https://github.com/MichMich/MagicMirror](https://github.com/MichMich/MagicMirror)
*   [My Bathroom Mirror Is Smarter Than Yours](https://medium.com/@maxbraun/my-bathroom-mirror-is-smarter-than-yours-94b21c6671ba#.1kttvbngz) （需要翻墙，[国内译文《Google 软件工程师自制智能镜》](http://36kr.com/p/5043096.html)）
*   [Android application powering the mirror in my house](https://github.com/HannahMitt/HomeMirror) 项目地址：[https://github.com/HannahMitt/HomeMirror](https://github.com/HannahMitt/HomeMirror)
*   [Smart Mirror Hardware](http://konkludenz.de/smart-mirror-hardware/) via Konkludenz （需要翻墙，德语）
*   [Magic Mirror](http://www.nils-snake.de/archives/magic-mirror-ein-raspberry-pi-projekt-teil-1) via Nils Hünerfürst （需要翻墙，德语）
*   [Magic Mirror](https://microsoft.hackster.io/en-US/Emmuss/magicmirror-cb222b) 使用Microsoft Windows 10 IoT Core
*   [Smart Mirror](http://adamlaycock.ca/blog/2015/07/01/Smart-Mirror-Intro.html) via Adam Laycock 项目地址：[https://github.com/alaycock/SmartMirror](https://github.com/alaycock/SmartMirror)
*   [Other mirror projects with alternate technologies](https://github.com/HannahMitt/HomeMirror/wiki/Other-mirror-projects-with-alternate-technologies)
*   [forum MirrorMirror](http://mirrormirror.tech/)
*   [来用树莓派打造一个魔镜吧](http://blog.jobbole.com/97180/?utm_source=top.jobbole.com&amp;utm_medium=relatedArticles) 中文
*   [使用树莓派制作一个魔镜](http://www.hellowk.cc/2016/02/18/raspberry-pi-magic-mirror-1/) 中文

是的，我正是看了这些材料才开始动手做智能镜子的。

![树莓派智能镜子](http://www.yupae.cn/images/magicmirror.jpeg)

是不是很心动？跟我学着做吧，Step by step。

### 关于镜子

镜子是制作这个魔镜的关键，要可以单面透光的那种。正面看是反光镜子，同时又可以把背后屏幕上的字清晰的透出来，这种材料的镜子叫**原子镜**，一般的装饰材料城卖玻璃的店都有得卖，当然某宝上也可以买到。原子镜根据透光率的高低分多个型号，8度的、14度的、20度的等等（数值越小透光率越低），最好拿几个试试，看看哪种效果才是你最想要的。我试了8度和14度的，感觉上8度更像镜子而14度的则略透，所以最后选择了8度的。

也可以用亚克力板贴上单透膜当做原子镜来用，当然这种效果是要逊色很多的（我试过…T T）。

### 屏幕设置

要让魔镜变成纵向的肖像模式，那么必须将屏幕顺时针旋转90度，前篇[《树莓派3（Raspberry Pi 3）那点事儿》](http://www.yupae.cn/iot/raspberrypi1)中已经讲过了怎么旋转屏幕，我们再复习一下。

打开树莓派BIOS设置文件 config.txt 文件，在文件内加上以下一行代码：

`
display_rotate=1
`

### 网络服务器

要在树莓派上运行 Web 服务就要安装 Apache 服务器，安装之前先更新一下系统，确定用的是最新系统软件（【重要】很多安装错误都可以通过更新解决）。

`
sudo apt-get update && apt-get upgrade -y
`

安装 Apache：

`
sudo apt-get install apache2 apache2-doc apache2-utils
`

加上 PHP 支持：

`
sudo apt-get install libapache2-mod-php5 php5 php-pear php5-xcache
`

重启之后，网页服务器就挂载上线运行了！在 /var/www/html 文件夹内放置了 index.php 文件，将浏览器首页指向树莓派的 IP 地址，发现成功了。

### 信息模式（kioskmode）

我们需要魔镜在每次开机后自动进入到界面状态，所以我们要安装一个带信息模式（kioskmode）的浏览器，Linux 下 Chromium 浏览器支持 kioskmode。至于如何安装 Chromium 浏览器以及怎样打开信息模式将在下篇中详细说明。