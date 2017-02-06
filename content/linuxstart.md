---
keywords: linux命令、开机脚本
title: 开机自动执行脚本 与 update-rc.d
---

# 开机自动执行脚本 与 update-rc.d

所有的机器都有可能重启，很多应用程序、服务需要在机器启动的时候自动执行，这里记录两种开机执行脚本的方法及相关的一个命令。

## 修改/etc/rc.local

在`/etc/rc.local`的`exit 0`语句之间添加启动脚本。脚本必须具有可执行权限。

## 用update-rc.d命令添加开机执行脚本

创建要开机自动执行的脚本：`/home/test/blog/startBlog.sh`，并给予可执行权限：`chmod +x /home/test/blog/startBlog.sh`。

在`/etc/init.d`目录下创建链接文件到前面的脚本： `ln -s  /home/test/blog/startBlog.sh  /etc/init.d/startBlog`。

进入`/etc/init.d`目录，用 `update-rc.d` 命令将连接文件 startBlog 添加到启动脚本中去：`update-rc.d  startBlog  defaults  99`。  

其中的99表示启动顺序，取值范围是`0-99`。序号越大的越晚执行。

移除启动的脚本：`update-rc.d  -f  startBlog  remove`。  

`-f`选项表示强制执行。

## update-rc.d命令

此命令用于安装或移除System-V风格的初始化脚本连接。脚本是存放在 `/etc/init.d/`目录下的，当然可以在此目录创建连接文件连接到存放在其他地方的脚本文件。

此命令可以指定脚本的执行序号，序号的取值范围是 `0-99`，序号越大，越迟执行。

### 用法

`update-rc.d  [-n]  [-f]  name  remove`   用于移除脚本。  

`update-rc.d  [-n]  name  default  [NN  |  SS  KK]`，NN表示执行序号（0-99），SS表示启动时的执行序号，KK表示关机时的执行序号，SS、KK主要用于在脚本直接的执行顺序上有依赖关系的情况下。

### 选项

*   `-n`：不做任何事情，只显示将要做的。（预览、做测试）
*   `-f`：强制移除符号连接，即使 `/etc/init.d/script-name` 仍然存在。

### 举例

`update-rc.d  startBlog  defaults  99`：添加一个启动连接，执行序号是99。

如果执行脚本B需要先执行脚本A，如下设置（A的启动顺序比B的小，结束顺序比B的大）：  

`update-rc.d  script_for_A  defaults  80  20`  

`update-rc.d  script_for_B  defaults  90  10`

添加一个不被其他任何服务需要的服务：`update-rc.d  script_name  defaults   98  02`，

添加一个需要 开始/结束 序号在20的服务的服务：`update-rc.d  script_depends_on_service_20  default  21  19`。

移除一个脚本，假定`/etc/init.d/`目录下的脚本文件已先被删除：`update-rc.d  script_name  remove`。

移除一个脚本，不管`/etc/init.d/`目录下的脚本文件是否已删除：`update-rc.d  -f  script_name  remove`。
