---
layout: post
title: 树莓派3 raspi-config 设置
date: 2016-04-06
categories: 树莓派
cover: 'http://www.yupae.cn/images/raspiconfig.jpg'
tags: 树莓派 Raspberry Raspbian
---

[接上篇《树莓派3（Raspberry Pi 3）那点事儿》](http://www.yupae.cn/2016-04-05-raspberrypi1)

## 树莓派 raspi-config 工具

初装系统首次启动过程中，会出现raspi-config界面。

![raspi-config界面](http://www.yupae.cn/images/raspiconfig.jpg)

用以下命令亦可调出上述界面。

`
 sudo raspi-config
`

进入系统设置界面进行设置，第1步必须做，其它根据自己需要设置。

```
1、Expand Filesystem：将系统扩展到整个TF卡，必须执行，一路敲回车即可；
2、Change User Password：修改密码及账号，若要修改，请牢记；
3、Boot Options：启动选项，选“desktop”，下次启动将直接进入桌面系统；
4、Wait for Network at Boot：等待网络唤醒，服务器模式；
5、Internationalisation Options：地区和语言设置，选择zh_CN UTF-8，即可切换到中文模式；
6、Enable Camera：CSI摄像头开启或关闭，选择enable则开启；
7、Add to Rastrack：将树莓派加入Rastrack社区；
8、Overclock：超频，一般不超频，若要超频请谨慎操作，并做好散热；
9、Advanced Options：高级选项，包括Overscan、Hostname、Memory Split（内存分配）、SSH等；
0、About raspi-config：本机相关信息。
```


设置完成后，树莓派自动重启。

## ROOT 账号设置

如果你安装的是官方的 Raspbian 系统，那么默认的登录帐号为 pi 密码是 raspberry 。

为了方便折腾，建议第一时间启用 ROOT 账号。这个也很简单的，只需要执行一下两句命令即可：


```
// 设置 root 账号的密码，会让你输入两次新密码
sudo passwd root

// 启用 root 账号登录
sudo passwd --unlock root
```

执行完之后，用 reboot 命令重启就可以用 root 登录啦。
