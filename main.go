package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"
)

var tempDate = time.Now().Format("2006-01-02")

func main() {
	var preface, content string
	preface = `目录
===

`
	content = preface + catalog("_posts") + "\n"
	content += "\n最后更新时间: " + time.Now().Format("2006-01-02 15:04:05") + "\n"

	writeMarkDown("README", content)
	println("READMD.md is rewrited.")

	//gitPull()
	gitAddAll()
	gitCommit()
	gitPush()
}

func catalog(docName string) (result string) {
	docPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	mdFiles, err := listDir(docPath+string(os.PathSeparator)+docName, ".md")
	if err != nil {
		log.Fatal(err)
	}
	fileLists := make(map[string][3]string)
	var array [3]string
	for _, v := range mdFiles {
		title := strings.Replace(readByLine(docName+string(os.PathSeparator)+v, 2), "title: ", "", -1)
		date := strings.Replace(readByLine(docName+string(os.PathSeparator)+v, 3), "date: ", "", -1)
		d, _ := time.Parse("2006-01-02", date)
		datring := strconv.FormatInt(d.Unix(), 10)
		array[0], array[1], array[2] = v, title, date
		fileLists[datring] = array
	}

	var keys []string
	for k := range fileLists {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	for i, k := range keys {
		pathname := strings.SplitN(fileLists[k][0], "-", 4)
		result = result + strconv.Itoa(i+1) + ". **[" + fileLists[k][1] + "](http://yupae.cn/" + strings.Replace(strings.Replace(strings.Trim(fmt.Sprint(pathname), "[]"), " ", "/", -1), ".md", ".html", 1) + ")** " + fileLists[k][2] + "\n"
	}
	return
}

//listDir
func listDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, fi.Name())
		}
	}
	return files, nil
}

func readByLine(path string, line int) (result string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	// 缺省的分隔函数是bufio.ScanLines
	// 也可以使用ScanWords
	// scanner.Split(bufio.ScanWords)
	// 或者定制一个SplitFunc类型的分隔函数
	i := 0
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
		if i == line {
			// 得到数据，Bytes() 或者 Text()
			result = scanner.Text()
		} else if i > line {
			break
		}
		i++
	}
	return
}

func writeMarkDown(fileName, content string) {
	// open output file
	fo, err := os.Create(fileName + ".md")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)
	w.WriteString(content)
	w.Flush()
}

func gitPull() {
	app := "git"
	arg0 := "pull"
	arg1 := "origin"
	arg2 := "master"
	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}

func gitAddAll() {
	app := "git"
	arg0 := "add"
	arg1 := "."
	cmd := exec.Command(app, arg0, arg1)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}

func gitCommit() {
	app := "git"
	arg0 := "commit"
	arg1 := "-m"
	arg2 := tempDate
	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}

func gitPush() {
	app := "git"
	arg0 := "push"
	arg1 := "origin"
	arg2 := "master"
	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}
