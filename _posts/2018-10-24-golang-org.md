---
layout: post
title: 解决 golang.org/x/ 包下载问题
date: 2018-10-24
categories: Golang
tags: hosts golang
---

### 使用方法

从 github.com 上先clone下来，再做软链接

```bash
mkdir -p $GOPATH/src/github.com/golang/
git clone https://github.com/golang/sys.git $GOPATH/src/github.com/golang/sys
git clone https://github.com/golang/net.git $GOPATH/src/github.com/golang/net
git clone https://github.com/golang/text.git $GOPATH/src/github.com/golang/text
git clone https://github.com/golang/lint.git $GOPATH/src/github.com/golang/lint
git clone https://github.com/golang/tools.git $GOPATH/src/github.com/golang/tools
git clone https://github.com/golang/crypto.git $GOPATH/src/github.com/golang/crypto
git clone https://github.com/golang/sync.git $GOPATH/src/github.com/golang/sync
git clone https://github.com/golang/time.git $GOPATH/src/github.com/golang/time
git clone https://github.com/golang/image.git $GOPATH/src/github.com/golang/image


ln -s $GOPATH/src/github.com/golang/ $GOPATH/src/golang.org/x
```
