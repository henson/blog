---
layout: post
title: 树莓派3使用RTC实时时钟
date: 2017-04-25
categories: 树莓派
tags: 树莓派 Raspberry RTC
---

在树莓派使用上RTC实时时钟（DS3231，I2C接口），不用担心断电后时间归零的问题，开机后自动同步RTC时钟！！！

### 购买基于DS3231的RTC时钟模块，并且支持3.3V的那种

![DS3231的RTC时钟模块1](http://www.yupae.cn/images/raspberry-pi-real-time-clock-rtc-ds3231-1.jpg)

### 配置树莓派

![DS3231的RTC时钟模块2](http://www.yupae.cn/images/raspberry-pi-real-time-clock-rtc-ds3231-2.jpg)

#### 打开树莓派的i2c接口

`
sudo raspi-config --&gt;Advanced Option--&gt;I2C，全部选择yes
`

#### 添加i2c模块

`
sudo nano /etc/modules
`

然后添加以下两行内容：

```
i2c-bcm2708
i2c-dev
```


#### 安装i2c工具，查看i2c设备b

`
sudo apt-get install i2c-tools
`

执行命令查看i2c设备

`
sudo i2cdetect -y -a 1
`

```
     0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f
00:          -- -- -- -- -- -- -- -- -- -- -- -- --
10: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
20: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
30: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
40: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
50: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
60: -- -- -- -- -- -- -- -- 68 -- -- -- -- -- -- --
70: -- -- -- -- -- -- -- --
```


显示的地址就是我们需要使用的，树莓派一般是68

#### 编辑启动文件 /etc/rc.local

`
sudo nano /etc/rc.local
`


```
// 将以下内容加入“exit 0”行之前
sudo modprobe i2c-dev
sudo modprobe i2c-bcm2708
echo ds3231 0x68 | sudo tee  /sys/class/i2c-adapter/i2c-1/new_device
sudo hwclock -s
```


### hwclock命令介绍


```
sudo hwclock -r     #读时钟
sudo hwclock -s     #写时钟
```

### Remove fake-hwclock

`
sudo apt-get --purge remove fake-hwclock
`

`
sudo systemctl disable hwclock-stop
`

### 参考材料

*   [Raspberry Pi 2 Model B——添加硬件时钟](http://overcosine.tk/2016/01/21/raspberry-pi-2-model-b-8%E6%B7%BB%E5%8A%A0%E7%A1%AC%E4%BB%B6%E6%97%B6%E9%92%9F/)
*   [Raspberry Pi B+ 使用 DS3231 实时时钟(RTC) 模块](http://0oor.com/archives/114)
*   [树莓派(Raspberry Pi)日期时间不准的修正方法](http://www.cnblogs.com/infopi/p/3947652.html)
*   [使用ntpdate更新系统时间](http://blog.csdn.net/suer0101/article/details/7868813)
*   [Tutorial: Add a Real Time Clock to the Raspberry Pi](https://www.raspberrypi.org/forums/viewtopic.php?f=44&amp;t=16218&amp;start=25)
*   [给Raspberry Pi添加RTC模块](https://www.ijser.cn/add-rtc-mode-to-resparrypi/)

    