package main

import "fmt"
import "net/url"
import "strings"
import "io/ioutil"

//import "io"
import "encoding/base64"

import "crypto/rc4"

func eprint(lvl int, a ...interface{}) {
	if lvl > 0 && a != nil && a[0] != nil {
		fmt.Println(a)
	}
}

func proto() string {
	return "data:text/html,"
}
func defhtml() string {
	return "<title>self decrypting image</title><body onload=\"main();\"><canvas id=a>"
}
func js() string {
	s, e := ioutil.ReadFile("./dc.js")
	if e != nil {
		eprint(1, e)
		return ""
	}
	return "<script>" + string(s)
}
func escape(in string) string {
	return strings.Replace(url.QueryEscape(in), "+", "%20", -1)
}
func enc(src []byte) []byte {
	key := []byte("abc123")
	c, err := rc4.NewCipher(key)
	if err != nil {
		fmt.Println("fuck")
	}
	dst := make([]byte, len(src))
	c.XORKeyStream(dst, src)

	return dst
}
func data() string {
	/*s, e := ioutil.ReadFile("./hellokitty2.jpg")
	if e != nil {
		eprint(1, e)
		return ""
	}*/
	s := []byte("hello world from crypto land")
	out := base64.StdEncoding.EncodeToString(enc(s))

	return " \n da=\"" + out + "\"; \n</script>"
}
func main() {
	fmt.Println(proto() + escape(defhtml()+js()+data()))
}
