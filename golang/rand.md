---
keywords: yupae.cn
title: Golang随机字符串生成函数（数字、大小写字母）
---

# Golang随机字符串生成函数（数字、大小写字母）      

```
package main

import(
    "fmt"
    "time"
    "math/rand"
)

const (
    KC_RAND_KIND_NUM   = iota	// 纯数字
    KC_RAND_KIND_LOWER			// 小写字母
    KC_RAND_KIND_UPPER			// 大写字母
    KC_RAND_KIND_ALL			// 数字、大小写字母
)

// 随机字符串
func Krand(size int, kind int) []byte {
    ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
    is_all := kind &gt; 2 || kind &lt; 0
    rand.Seed(time.Now().UnixNano())
    for i :=0; i &lt; size; i++ {
        if is_all { // random ikind
            ikind = rand.Intn(3)
        }
        scope, base := kinds[ikind][0], kinds[ikind][1]
        result[i] = uint8(base+rand.Intn(scope))
    }
    return result
}

func main(){
    fmt.Println("num:   " + string(Krand(16, KC_RAND_KIND_NUM)))
    fmt.Println("lower: " + string(Krand(16, KC_RAND_KIND_LOWER)))
    fmt.Println("upper: " + string(Krand(16, KC_RAND_KIND_UPPER)))
    fmt.Println("all:   " + string(Krand(16, KC_RAND_KIND_ALL)))
}
```

    