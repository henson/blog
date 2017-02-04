package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("README.md")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	// 缺省的分隔函数是bufio.ScanLines,我们这里使用ScanWords。
	// 也可以定制一个SplitFunc类型的分隔函数
	//scanner.Split(bufio.ScanWords)
	var temp string
	for {
		// scan下一个token.
		success := scanner.Scan()
		if success == false {
			// 出现错误或者EOF是返回Error
			err = scanner.Err()
			if err == nil {
				//log.Println("Scan completed and reached EOF")
				break
			} else {
				log.Fatal(err)
			}
		}
		// 得到数据，Bytes() 或者 Text()
		//fmt.Println("Line found:", scanner.Text())
		temp = temp + scanner.Text() + "\n"
		// 再次调用scanner.Scan()发现下一个token
	}
	fmt.Println(temp)
}
