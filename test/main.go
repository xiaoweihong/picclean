package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func urlTest() {
	url := "http://192.168.100.118:9333/2/1561002f31d6"
	// http://192.168.100.118:9301/2,1561002f31d6
	//index := strings.LastIndex(url, "/")
	//replace := strings.Replace(url, "/*/", "/*,", -1)
	//newUrl:=url[:strings.LastIndex(url, "/")]+url[strings.LastIndex(url, "/")+2:]
	//fmt.Println(newUrl)
	s1 := url[:strings.LastIndex(url, "/")]
	s2 := url[strings.LastIndex(url, "/")+1:]
	fmt.Println(s1)
	fmt.Println(s2)
	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString(",")
	buffer.WriteString(s2)
	fmt.Println(buffer.String())
	t := time.Now().UnixNano()
	fmt.Println(t / 1e6)
}

func timeParse() {
	strT1 := "2020-03-12 08:00:00"
	parse, err := time.ParseInLocation("2006-01-02 15:04:05", strT1,time.Local)
	if err != nil {

	}
	fmt.Println(parse.Unix())
}

func main() {
	timeParse()
}
