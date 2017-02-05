---
keywords: yupae.cn
title: Golang随机字符串生成函数（数字、大小写字母）
---

# Golang随机字符串生成函数（数字、大小写字母）      

    `package main

import(
    &#34;fmt&#34;
    &#34;time&#34;
    &#34;math/rand&#34;
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
    fmt.Println(&#34;num:   &#34; + string(Krand(16, KC_RAND_KIND_NUM)))
    fmt.Println(&#34;lower: &#34; + string(Krand(16, KC_RAND_KIND_LOWER)))
    fmt.Println(&#34;upper: &#34; + string(Krand(16, KC_RAND_KIND_UPPER)))
    fmt.Println(&#34;all:   &#34; + string(Krand(16, KC_RAND_KIND_ALL)))
}
`

    