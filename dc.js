function
src4(ssrc , skey ) {
	var key = []
	for (var i=0 ; i<skey.length; i++)
		key.push(skey.charCodeAt(i))
	var src = []
	for (var i=0 ; i<ssrc.length; i++)
		src.push(ssrc.charCodeAt(i))

	//s := make([]byte, 256)
	var s = []

	//out := make([]byte, len(src))
	var out = ""
	var i=0
	for (i = 0; i < 255; i += 1) {
		s[i] = i
	}
	var j = 0
	for (i = 0; i < 255; i += 1) {
		j = (j + s[i] + key[i%key.length]) % 255
		temp = s[i]
		s[i] = s[j]
		s[j] = temp
	}
	j = 0
	i = 0
	//key initialised
	for (x = 0; x < src.length; x += 1) {
		i = (i + 1) % 255
		j = (j + s[i]) % 255
		{
			temp = s[i]
			s[i] = s[j]
			s[j] = temp
		}
		//out = out + String.fromCharCode(s[(s[i]+s[j])%255]^src[x])
		out = out +" " + s[(s[i]+s[j])%255]
	}
	return out
}

