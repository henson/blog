---
layout: post
title: 树莓派3设置python脚本开机启动
date: 2017-04-25
categories: 树莓派
tags: 树莓派 Raspberry python
---

### 首先写个简单的python脚本

脚本很简单，就是通过控制树莓派上gpio让LED灯闪烁。

文件保存在/home/pi/script/ledblink.py

```
#!/usr/bin/env python

import RPi.GPIO as GPIO
import time
GPIO.setmode(GPIO.BCM)
GPIO.setup(21,GPIO.OUT)
while True:
  try:
    GPIO.output(21,True)
    time.sleep(1)
    GPIO.output(21,False)
    time.sleep(1)
  except (KeyboardInterrupt, SystemExit):
    GPIO.close()
    print "exit"
```


### 开机启动脚本

保存脚本为/etc/init.d/ledblink文件

```
#!/bin/bash
# /etc/init.d/ledblink

### BEGIN INIT INFO
# Provides: embbnux
# Required-Start: $remote_fs $syslog
# Required-Stop: $remote_fs $syslog
# Default-Start: 2 3 4 5
# Default-Stop: 0 1 6
# Short-Description: ledblink initscript
# Description: This service is used to manage a led
### END INIT INFO

case "$1" in
    start)
        echo "Starting LED Blink"
        python /home/pi/script/ledblink.py &amp;
        ;;
    stop)
        echo "Stopping ledblink"
        # killall ledblink.py
        kill $(ps aux | grep -m 1 'python /home/pi/script/ledblink.py' | awk '{ print $2 }')
        ;;
    *)
        echo "Usage: sudo /etc/init.d/ledblink start|stop"
        exit 1
        ;;
esac
exit 0
```


### 设置python脚本开机启动

使用chmod命令让脚本可执行

`
sudo chmod +x /etc/init.d/ledblink
`

这样启动该脚本


```
# 启动
sudo /etc/init.d/ledblink start
# 停止
sudo /etc/init.d/ledblink stop
```

最后设置开机启动就好了

`
sudo update-rc.d ledblink defaults
`

去掉开机启动也很容易

`
sudo update-rc.d -f ledblink remove
`

这样就完工了,重启树莓派就会发现led自己闪烁了。

### 参考资料

[LSB Init Script Notes](http://www.rcramer.com/tech/linux/init_lsb.shtml)

    