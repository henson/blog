---
keywords: yupae.cn
title: Golang 实现的线程安全的队列
---

# Golang 实现的线程安全的队列      

Golang 的 channel 可以作为线程安全的队列来使用，但是只能是固定大小的，如果填满的 channel，就只能阻塞的等待。

goqueue 就是来解决这个问题的，参考了 Python 里面的 Queue(Py3 queue)模块。

详见 [https://github.com/Damnever/goqueue](https://github.com/Damnever/goqueue)

    `package main

import (
&#34;github.com/Damnever/goqueue&#34;
&#34;fmt&#34;
)

func main() {
    queue := goqueue.NewQueue(0)

    worker := func(queue *goqueue.Queue) {
        for {
            e, err := queue.Get(true, 0)
            if err != nil {
                fmt.Println(&#34;Unexpect Error: %v\n&#34;, err)
            }
            num := e.Value.(int)
            fmt.Printf(&#34;-&gt; %v\n&#34;, num)
            queue.TaskDone()
            if num % 3 == 0 {
                for i := num + 1; i &lt; num + 3; i ++ {
                     queue.PutNoWait(i)
                }
            }
        }
    }

    for i := 0; i &lt;= 27; i += 3 {
        queue.PutNoWait(i)
    }

    for i := 0; i &lt; 5; i++ {
        go worker(queue)
    }

    queue.WaitAllComplete()
    fmt.Println(&#34;All task done!!!&#34;)
}
`

    