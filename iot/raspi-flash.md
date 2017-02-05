---
keywords: 树莓派,Raspberry,flash,yupae.cn
title: 树莓派3播放Flash
---

# 树莓派3播放Flash      

### Chromium

首先要安装好Chromium浏览器，不知道怎么安装的请看前篇[《树莓派3安装Chromium浏览器》](http://www.yupae.cn/iot/raspi-chromium/)。

### 删除树莓派自带浏览器

#### 先查新一下系统都有些什么浏览器

`
dpkg -l |grep browser
`

一般情况下会得到如下结果:

```
pi@raspberrypi ~ $ dpkg -l |grep browser
 ii  dillo                   3.0.2-2                 armhf      Small and fast web browser
 ii  epiphany-browser        3.8.2.0-0rpi18rpi1      armhf      Intuitive GNOME web browser
 ii  epiphany-browser-data   3.8.2.0-0rpi18rpi1      all Data   files for the GNOME web browser
 ii  info                    4.13a.dfsg.1-10         armhf      Standalone GNU Info documentation browser
 ii  netsurf-common          2.9-2                   all        Small web browser with CSS support common files
 ii  netsurf-gtk             2.9-2                   armhf      Small web browser with CSS support for GTK
```


#### 删除不打算使用的浏览器和其依赖包并清除配置

删除dillo

`
sudo apt-get autoremove --purge dillo
`

删除epiphany-browser

`
sudo apt-get autoremove --purge epiphany-browser*
`

删除netsurf

`
sudo apt-get autoremove --purge netsurf*
`

### Flash

下载Flash插件，[chromium-pepper-flash-12-12.0.0.77-1-armv7h.pkg.tar.xz](http://odroidxu.leeharris.me.uk/repo/chromium-pepper-flash-12-12.0.0.77-1-armv7h.pkg.tar.xz)

获取flash player for armv7


```
wget http://odroidxu.leeharris.me.uk/repo/chromium-pepper-flash-12-12.0.0.77-1-armv7h.pkg.tar.xz
xz chromium-pepper-flash-12-12.0.0.77-1-armv7h.pkg.tar.xz -d
tar -xvf chromium-pepper-flash-12-12.0.0.77-1-armv7h.pkg.tar
cd ./usr/lib/PepperFlash
chmod +x *
sudo cp * /usr/lib/Chromium-browser/plugins
```


修改配置文件

`
sudo nano /etc/Chromium-browser/default
`

修改最后一句为


`
CHROMIUM_FLAGS="--ppapi-flash-path=/usr/lib/Chromium-browser/plugins/libpepflashplayer.so --ppapi-flash-version=12.0.0.77 -password-store=detect -user-data-dir"
`

进入chromium输入
`
chrome://plugins
`
查看flash是否添加到插件里，是否开启。

    