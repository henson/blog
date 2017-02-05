---
keywords: yupae.cn
title: Go写的文件缓存系统
---

# Go写的文件缓存系统      

大概的看了一下，感觉有点squid的感觉，缓存文件。

初始化开启一个缓存

cache := filecache.NewDefaultCache()

cache.MaxSize = 128 * filecache.Megabyte

cache.Start()

调用中使用

    `func FileServer(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if len(path) &gt; 1 {
		path = path[1:len(path)]
	} else {
		path = &#34;.&#34;
	}
	err := cache.WriteFile(w, path)
	if err == nil {
		ServerError(w, r)
	} else if err == filecache.ItemIsDirectory {
		DirServer(w, r)
	}
}
`

[http://gokyle.github.com/filecache/](http://gokyle.github.com/filecache/)

    