---
keywords: 树莓派,Raspberry,Raspbian,yupae.cn
title: 树莓派3更改键盘布局
---

# 树莓派3更改键盘布局      

[接上篇《树莓派3更换软件源、设置静态IP、激活蓝牙》](http://www.yupae.cn/iot/raspberrypi3/)

### 导读

很多童鞋发现默认的键盘设置打不出“&#124;\/”等符号，虽然设置了标准的PC104 Keyboard。

![美式标准104键盘](http://img.blog.csdn.net/20140804195858791)

那是因为树莓派源产自英国，英式键盘布局和美式并不相同，而我们国内使用的键盘大多是标准的美式104键盘，下面就一步一步讲讲如何正确配置。

### 树莓派设置

`
sudo raspi-config
`

#### 进入国际化配置选项

![进入国际化配置选项](http://img.blog.csdn.net/20140804195933501?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQveHVrYWk4NzExMDU=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

#### 修改键盘布局

![修改键盘布局](http://img.blog.csdn.net/20140804195737281?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQveHVrYWk4NzExMDU=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

#### 选择PC104标准键盘

![选择PC104标准键盘](http://img.blog.csdn.net/20140804195752484?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQveHVrYWk4NzExMDU=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

#### 选择美国标准

![选择美国标准](http://img.blog.csdn.net/20140804200018258?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQveHVrYWk4NzExMDU=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

#### 选择键盘默认布局

![选择键盘默认布局](http://img.blog.csdn.net/20140804200030738?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQveHVrYWk4NzExMDU=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

#### compose key设置

![compose key设置](http://img.blog.csdn.net/20140804195828906?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQveHVrYWk4NzExMDU=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

#### ctrl+alt+backspace组合键，类似于windows的ctrl+alt+delete。

![ctrl+alt+backspace组合键](http://img.blog.csdn.net/20140804200057133?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQveHVrYWk4NzExMDU=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

#### 完成设置

![完成设置](http://img.blog.csdn.net/20140804200109628?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQveHVrYWk4NzExMDU=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

### 参考资料

*   [树莓派(raspberry pi)学习4: 更改键盘布局](http://blog.csdn.net/c80486/article/details/8460271)——修改为PC101布局，普通键盘同样可以使用。
*   《玩转树莓派 Raspberry Pi》 王江伟等。——书中误写成了PC105键盘，该键盘为欧洲标准和国内使用的键盘不同。
