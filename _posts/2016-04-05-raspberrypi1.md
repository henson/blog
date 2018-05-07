---
layout: post
title: 树莓派3（Raspberry Pi 3）那点事儿
date: 2016-04-05
categories: 树莓派
cover: 'https://www.raspberrypi.org/app/uploads/2017/05/Raspberry-Pi-3-1-1619x1080.jpg'
tags: 树莓派 Raspberry Raspbian
---

核心内容导读：树莓派版本、USB系统烧录、HDMI分辨率设置、屏幕翻转、超频设置

入手了一个树莓派3（[RASPBERRY PI 3 MODEL B](https://www.raspberrypi.org/products/raspberry-pi-3-model-b/)），捣鼓了一阵子感觉还不错。其实细搜一下，关于树莓派的教程网上是有很多的，当然这其中的坑也有很多。另外，大部分教程是关于树莓派2的，所以特地开个篇把树莓派3（Raspberry Pi 3）那点事儿和踩过的坑统统记下来。

## 关于树莓派的版本

![树莓派的版本](http://www.yupae.cn/images/raspberryforversion.jpg)

树莓派有两个版本，[Element14（易络盟电子）](http://cn.element14.com/)和 [RS（欧时电子）](http://china.rs-online.com/web/)，产地不同而已，Element14是Made in China，RS的是Made in UK，主芯片配置都是完全一样的。某宝上，R版的要E版的贵些，但绝对不是卖家口中的质量差异而只是海外运费的价差而已。

## 关于TF卡

买多大的TF卡才合适呢，这是每一个新上手树莓派的朋友首先考虑的问题。

官方系统[RASPBIAN JESSIE](https://www.raspberrypi.org/downloads/raspbian/)只有1.36G，如果你只是拿来装个系统加上一些常用的软件，理论上4G的TF卡就足够了。当然，如果你打算用来做NAS，那4G显然不够。那是不是容量越大越好呢？经过实测，一些大容量的TF卡会存在一些兼容性问题，同样的，TF卡的速度也并非越快越好，在选择TF卡的时候不要一味追求大和快，在权衡性价比的基础上首要考虑是否兼容。以下给出的是市面上所有TF卡与树莓派兼容性的结果供参考（Reference: [Working / Non-working SD cards](http://elinux.org/RPi_SD_cards)），从中不难发现目前闪迪（Sandisk）各批次的16G TF卡其兼容性都是OK的。

## 系统烧录

树莓派买来了，卡也买来了，开始装系统咯。首先，从官网上下载最新版的[RASPBIAN JESSIE](https://www.raspberrypi.org/downloads/raspbian/)系统，然后把系统烧录到TF卡上，这里就要用到镜像工具，推荐使用USB Image Tool。

![USB Image Tool](http://segmentfault.com/img/bVciD3)

*   官网：[http://www.alexpage.de/usb-image-tool/](http://www.alexpage.de/usb-image-tool/)
*   直接下载：[USB Image Tool.zip](http://pan.baidu.com/share/link?shareid=4232180688&amp;uk=605377859) (216KB)
*   Win7可直接运行；WinXP如果运行出错，请安装[.net Framework 2.0](https://www.microsoft.com/zh-cn/download/details.aspx?id=1639) (22MB)。

对于Raspberry Pi等开发板，几乎所有的新手教程都推荐使用Win32DiskImager作为系统安装工具，其实这个工具问题不少：

*   不支持中文目录名（文件或目录有中文，会出现123错误）；
*   只使用盘符区分不同USB设备，用户需要自行核实，不方便；
*   浏览文件，记不住最后使用的目录，找常用的文件很麻烦；
*   不支持zip或gz格式的压缩档案；
*   必须先插好SD卡，再开软件；
*   其它小问题：界面风格简陋；MD5在XP下乱码等。

而USBIT恰恰不存在这些问题，使用起来非常方便，所以在Windows环境下建议还是使用USBIT作为USB烧录工具。这里再推荐一篇文章详细介绍USB Image Tool的使用，有兴趣的不妨看一下，这里就不赘述了（Reference: [USB Image Tool：Windows下的直接写盘利器](https://segmentfault.com/a/1190000000492510)）。

## 树莓派配置文档 Config.txt 说明

由于树莓派没有传统意义上的BIOS，所以现在各种系统配置参数通常被存在“Config.txt”这个文本文件中，它会在ARM内核初始化之前被GPU读取。

这个文件存在引导分区上的。所以对于Linux，路径通常是/boot/config.txt，如果是Windows（或者OS X）它会被识别为SD卡中可访问部分的一个普通文件。

在Windows环境下进入TF卡，找到文件config.txt，使用文本工具打开，依次将下面项目的“值”修改为等号后的数值，并去掉前面的“#”：

```
hdmi_force_hotplug=1
hdmi_group=2
hdmi_mode=16
hdmi_drive=2
config_hdmi_boost=0
sdtv_mode=2
arm_freq=800  
```


解释：

*   hdmi_force_hotplug=1是支持HDMI设备的热插拔；
*   hdmi_group=2，hdmi_mode=16是显示模式及分辨率修改项；
*   hdmi_drive=2，表示音频从HDMI接口输出；
*   config_hdmi_boost=0，设置HDMI接口的信号强度，默认为0，如果出现HDMI干扰问题可以尝试设为4，最大为7；
*   sdtv_mode=2表示显示制式为标准PAL制；
*   arm_freq=800为调频项，可超频为900、1000，修改前请务必做好散热准备，谨慎超频、后果自负。

关于HDMI分辨率，这里再多说几句，hdmi_group和hdmi_mode一定要找到对应的值（附图是部分分辨率参数值截图）。另外设置好分辨率后，一定要把“hdmi_safe=1”这句用“#”注释掉，要不然HDMI输出的分辨率永远都是640*480。

![部分分辨率参数值](http://www.yupae.cn/images/raspberryhdmi.png)

其它参数意义及其设置，可以参考[http://elinux.org/RPiconfig](http://elinux.org/RPiconfig)上的叙述，也可以参看这篇文章（Reference: [树莓派配置文档 config.txt 说明](http://shumeipai.nxez.com/2015/11/23/raspberry-pi-configuration-file-config-txt-nstructions.html)）。

## 屏幕翻转

相信认真看了上述材料的童鞋已经知道了如何设置屏幕的翻转。对，使用 display_rotate 设置，按顺时针旋转屏幕（默认为0）。


```
display_rotate=0        正常
display_rotate=1        90度
display_rotate=2        180度
display_rotate=3        270度
display_rotate=0x10000  水平翻转
display_rotate=0x20000  垂直翻转
```

## 关于屏幕右上角的彩色块

其实到这里，系统已经算是基本安装配置好了，拔出TF卡插入树莓派背部卡槽，加电，迫不及待的要开机了…

![树莓派供电不足](http://www.yupae.cn/images/raspberryblink.jpg)

等一下，屏幕右上角一个彩色的小方块一闪一闪，难道是屏幕坏了吗？其实是树莓派供电不足造成的，当HDMI转VGA转换器、USB集线器、USB键盘等外设通过树莓派取电时就会造成树莓派的供电不足，这时候请一定使用输出5V  2.5A（树莓派2可以使用2A）的电源。连接树莓派的USB外设最好能独立供电，集线器请勿使用反向供电的，反向供电绕过了树莓派上所有自我保护的电路设计，电压电流过高过大易导致烧毁树莓派。
