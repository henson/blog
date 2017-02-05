---
keywords: Golang,加密解密,yupae.cn
title: Base64加密解密
---

# Base64加密解密      

    `package main

import (
    &#34;encoding/base64&#34;
    &#34;fmt&#34;
)

const (
    base64Table = &#34;123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912&#34;
)

var coder = base64.NewEncoding(base64Table)

func base64Encode(src []byte) []byte {
    return []byte(coder.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
    return coder.DecodeString(string(src))
}

func main() {
    // encode   
    hello := &#34;hello world&#34;
    debyte := base64Encode([]byte(hello))

    // decode   
    enbyte, err := base64Decode(debyte)
    if err != nil {
        fmt.Println(err.Error())
    }

    if hello != string(enbyte) {
        fmt.Println(&#34;hello is not equal to enbyte&#34;)
    }

    fmt.Println(string(enbyte))
}
`

    