package main

import "fmt"
import "net/url"
import "strings"
import "io/ioutil"

//import "io"
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
	return "<title>self decrypting image</title><body onload=\"main();\"><canvas id=a>" + "<script>"
}
func js(mfile string) string {
	dc, e := ioutil.ReadFile("./dc.js")
	out := ""
	if e != nil {
		eprint(1, e)
		return ""
	}
	out = out + string(dc)
	if len(mfile) > 0 {
		jsmain, e := ioutil.ReadFile(mfile)
		if e != nil {
			eprint(1, e)
			return ""
		}
		out = out + string(jsmain)
	}
	return out + data()
}
func escape(in string) string {
	return strings.Replace(url.QueryEscape(in), "+", "%20", -1)
}

func src4(src []byte, key []byte) []byte {
	s := make([]byte, 256)
	out := make([]byte, len(src))
	for i := byte(0); i < 255; i += 1 {
		s[i] = i
	}
	fmt.Print("//")
	fmt.Println(s)
	fmt.Print("//")
	j := byte(0)
	for i := 0; i < 255; i += 1 {
		j = (byte(j) + s[i] + key[i%len(key)]) % 255
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
		fmt.Print(s[i])
		fmt.Print(" ")
	}
	fmt.Println("//--")
	fmt.Print("//")
	fmt.Println(s)
	j = 0
	i := byte(0)
	//key initialised
	fmt.Print("//")
	for x := 0; x < len(src); x += 1 {
		i = (i + 1) % 255
		j = (j + s[i]) % 255
		{
			temp := s[i]
			s[i] = s[j]
			s[j] = temp
		}
		out[x] = s[(s[i]+s[j])%255] ^ src[x]
		fmt.Print(s[(s[i]+s[j])%255])
		fmt.Print(" ")
		//	out[x] = s[(s[i]+s[j])%255]
	}
	fmt.Println("")
	return out
}
func data() string {
	/*s, e := ioutil.ReadFile("./hellokitty2.jpg")
	if e != nil {
		eprint(1, e)
		return ""
	}*/
	s := []byte("hello world from crypto land")
	out := base64.StdEncoding.EncodeToString(src4(s, []byte("abc123")))

	return " \n da=\"" + string(out) + "\";"
}
func ehtml() string {
	return "\n</script>"
}
func genbrowser() string {
	return proto() + escape(defhtml()+js("./bmain.js")+ehtml())
}
func main() {
	fmt.Println(js(""))
	//fmt.Println(src4([]byte("umpalumpaumpalumpa"), []byte("abc123")))
}
