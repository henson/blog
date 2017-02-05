---
keywords: 树莓派,Raspberry,Raspbian,yupae.cn
title: 树莓派3更换软件源、设置静态IP、激活蓝牙
---

# 树莓派3更换软件源、设置静态IP、激活蓝牙      

[接上篇《树莓派3 raspi-config 设置》](http://www.yupae.cn/post/raspberrypi2/)

## 更换软件源（apt-get sources）

树莓派的服务器超级慢！导致安装一个几M的东西都要等上大半天！好在树莓派官方有提供一个镜像列表：[http://www.raspbian.org/RaspbianMirrors](http://www.raspbian.org/RaspbianMirrors) 在里面找了几个国内的镜像，经过几番尝试，觉得还是阿里云的速度不错，镜像主页：[http://mirrors.aliyun.com/raspbian/raspbian/](http://mirrors.aliyun.com/raspbian/raspbian/)

编辑 /etc/apt/sources.list 文件，这里推荐就用系统自带的 nano 命令编辑，命令如下：

`
sudo nano /etc/apt/sources.list
`

进入编辑界面，删除原有的内容，粘贴镜像主页提供的内容，结果如下：

```
deb http://mirrors.aliyun.com/raspbian/raspbian/ wheezy main non-free contrib rpi
deb-src http://mirrors.aliyun.com/raspbian/raspbian/ wheezy main non-free contrib rpi
```


其中wheezy是版本号，早期的版本是wheezy，如果你下载的是最新的，那么版本号应该改成jessie。然后使用 Ctrl+O 保存文件，Ctrl+X 退出编辑器。

```
// 更新软件索引清单
sudo apt-get update

// 比较索引清单更新依赖关系
sudo apt-get upgrade -y
```


## 设置静态 IP 地址

树莓派3默认是启用DHCP自动获得IP的，当然你也可以进到桌面模式设置静态IP，这里要讲的是如何在命令行模式下给树莓派设置一个静态IP，还是用 nano 来编辑网络接口文件：

`
sudo nano /etc/network/interfaces
`

如果你要设置的是有线网卡的 IP 地址，那么把 eth0 的 dhcp 改成 static 然后在下一行追加 IP 信息，结果大概如下：

```
iface eth0 inet static
address 192.168.1.200   # 设定的静态IP地址
netmask 255.255.255.0   # 网络掩码
gateway 192.168.1.1     # 网关
```


如果你要设置的是无线网卡的，那么除了把 wlan0 的 dhcp 改成 static 之外，还需要填写无线网的名称和密码，编辑后的结果大概如下：

```
iface wlan0 inet static
    wpa-ssid Your_Wifi_SSID
    wpa-psk Your_Wifi_Password
address 192.168.1.200 # 设定的静态IP地址
netmask 255.255.255.0 # 网络掩码
gateway 192.168.1.1   # 网关
network 192.168.1.1   # 网络地址
#wpa-roam /etc/wpa_supplicant/wpa_supplicant.conf
```


▲ 注意：注释掉最后一行

搞定之后，咱们用 poweroff 命令关掉树莓派，等到机器上的绿灯不闪了，把电源拔掉，再把网线拔掉，重新连接电源，稍等一会，看看是不是就通过无线网络的 IP 地址可以访问了。

## 激活树莓派3自带的蓝牙功能

树莓派3是自带WIFI和蓝牙的，以下是专门针对树莓派3的蓝牙功能编写的方法。Pi2和Pi1由于主板无内置蓝牙模块，所以不适用于本教程。

![树莓派3蓝牙设置](http://www.yupae.cn/images/raspiblue.png)


```
// 更新软件源及升级软件（如果前面做过了可以忽略）
sudo apt-get update
sudo apt-get updgrade -y
sudo apt-get dist-upgrade -y

// 安装蓝牙软件包
sudo apt-get install pi-bluetooth bluez bluez-firmware blueman
```


最关键一点，


```
// 添加pi用户到蓝牙组
sudo usermod -G bluetooth -a pi

// 重启
sudo reboot
```


控制台下蓝牙的连接和使用，用bluetoothctl就可以了：

![树莓派3蓝牙设置](http://www.yupae.cn/images/raspibluetooth.png)

大家看到可以使用命令


```
agent on
default-agent

// 用scan on之后，就可以看到扫描得到的蓝牙设备的物理地址了，格式XX:XX:XX:XX:XX:XX
scan on

// 配对使用以下命令
pair XX:XX:XX:XX:XX:XX

// 配对之后把设备添加到信任列表
trust XX:XX:XX:XX:XX:XX

// 之后连接设备
connect XX:XX:XX:XX:XX:XX

// 最后quit退出就可以了
quit
```

以后蓝牙设备开机后，树莓派会自动进行连接。

## 常见问题

基本上，现在树莓派系统已经可以正常工作了，但是还是有几个问题：

1、明明刚才已经设置了中文，为什么系统没有显示中文？

那是因为Raspbian JESSIE系统本身是没带中文字库的，需要我们手动安装：

`
sudo apt-get install ttf-wqy-microhei ttf-wqy-zenhei xfonts-wqy
`

2、字体安装

Raspbian JESSIE系统自带的字体比较少尤其是中文字体，可以直接把字体文件拷贝进系统font文件夹下：

`
sudo cp simsun.ttc usr/share/fonts/truetype/
`

3、安装拼音输入法

`
sudo apt-get install scim-pinyin
`

4、进入桌面系统的命令

`
startx
`

    