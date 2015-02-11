function
main() {
	alert("testing 123\n Hello World");
	var canvas = document.getElementById("a");
	var ctx = canvas.getContext("2d");
	ctx.fillStyle =  "rgb(0,0,0)";
	ctx.fillRect(0,0,ctx.canvas.width ,ctx.canvas.height);
	var img = new Image;
	img.src = "data:image/jpeg;base64,"+window.btoa(rc4("abc123",window.atob(da)));
	console.log(img.src);
	console.log(da)
	console.log(rc4(window.atob(da),"abc123"))
	ctx.drawImage(img,0,0);


}
func src4(src , key ) {
	//s := make([]byte, 256)
	var s
	//out := make([]byte, len(src))
	var out
	var i=0
	for (i = 0; i < 255; i += 1) {
		s[i] = i
	}
	var j := 0
	for (i = 0; i < 255; i += 1) {
		j = (byte(j) + s[i] + key[i%len(key)]) % 255
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
	}
	j = 0
	i := byte(0)
	//key initialised
	for x := 0; x < len(src); x += 1 {
		i = (i + 1) % 255
		j = (j + s[i]) % 255
		{
			temp := s[i]
			s[i] = s[j]
			s[j] = temp
		}
		out[x] = s[(s[i]+s[j])%255]
	}
	return out
}

