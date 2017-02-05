---
keywords: 树莓派,Raspberry,chromium,yupae.cn
title: 树莓派3安装Chromium浏览器
---

# 树莓派3安装Chromium浏览器      

### 安装Chromium浏览器

首先，在这里下载最新版的chromium:
[http://ports.ubuntu.com/pool/universe/c/chromium-browser/](http://ports.ubuntu.com/pool/universe/c/chromium-browser/)

```
chromium-codecs-ffmpeg-extra_49.0.2623.108-0ubuntu1.1233_armhf.deb
chromium-browser-l10n_49.0.2623.108-0ubuntu1.1233_all.deb
chromium-browser_49.0.2623.108-0ubuntu1.1233_armhf.deb
```


这三个文件应该只是中间的版本号不同，其他的是一样的。然后把这三个文件传入你的树莓派中，比如 ~/deb/chromium，进入这个文件夹，然后运行下面两条命令，恩，换成你自己下载的那个版本的：


```
sudo dpkg -i chromium-codecs-ffmpeg-extra_49.0.2623.108-0ubuntu1.1233_armhf.deb

sudo dpkg -i chromium-browser-l10n_49.0.2623.108-0ubuntu1.1233_all.deb chromium-browser_49.0.2623.108-0ubuntu1.1233_armhf.deb
```


这样 chromium 就安装好了。然后安装屏蔽鼠标的软件：

`
sudo apt-get install x11-xserver-utils unclutter
`

### 安装 Apache

`
sudo apt-get install apache2 apache2-doc apache2-utils
`

### 加上 PHP 支持

`
sudo apt-get install libapache2-mod-php5 php5 php-pear php5-xcache
`

重启之后，网页服务器就挂载上线运行了！在 /var/www/html 文件夹内放置了 index.php 文件，将浏览器首页指向树莓派的 IP 地址，发现成功了。

### 禁用屏幕保护和自动重启

启用信息模式（kioskmode）

修改这个文件：

`
sudo nano ~/.config/lxsession/LXDE-pi/autostart
`

在
`
@xscreensaver -no-splash
`
前面加上了 # 号。

还加入了以下代码：


```
@xset s off
@xset -dpms
@xset s noblank
@chromium-browser --kiosk --incognito http://localhost
```

这样就能完全禁用所有屏保功能，Chromium 浏览器也将在开机后自动启动，开启全屏模式并导向本地主页。

    