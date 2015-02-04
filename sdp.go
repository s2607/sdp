package main

import "fmt"
import "net/url"
import "strings"
import "io/ioutil"
import "encoding/base64"

func eprint(lvl int, a ...interface{}) {
	if lvl > 0 && a != nil && a[0] != nil {
		fmt.Println(a)
	}
}

func proto() string {
	return "data:text/html,"
}
func defhtml() string {
	return "<title>self decrypting image</title><body onload=\"main();\"><canvas id=a><script>"
}
func js() string {
	s, e := ioutil.ReadFile("./dc.js")
	if e != nil {
		eprint(1, e)
		return ""
	}
	return string(s)
}
func posthtml() string {
	return "</script>"
}
func escape(in string) string {
	return strings.Replace(url.QueryEscape(in), "+", "%20", -1)
}
func data() string {
	s, e := ioutil.ReadFile("./in.jpg")
	if e != nil {
		eprint(1, e)
		return ""
	}
	out := EncodeToString(s)

	return "a=\"" + out + "\";"
}
func main() {
	fmt.Println(proto() + escape(defhtml()+js()+posthtml()))
}
