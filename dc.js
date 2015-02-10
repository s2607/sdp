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
	ctx.drawImage(img,0,0);


}
function status (a) {
	console.log(a)
}
var dg=''
function makeArray(n) {
 for (var i=1; i<=n; i++) {
  this[i]=0
 }
 return this
}
function rc4(key, text) {
 var i, x, y, t, x2;
 status("rc4")
 s=makeArray(0);

 for (i=0; i<256; i++) {
  s[i]=i
 }
 y=0
 for (x=0; x<256; x++) {
  y=(key.charCodeAt(x % key.length) + s[x] + y) % 256
  t=s[x]; s[x]=s[y]; s[y]=t
 }
 x=0;  y=0;
 var z=""
 for (x=0; x<text.length; x++) {
  x2=x % 256
  y=( s[x2] + y) % 256
  t=s[x2]; s[x2]=s[y]; s[y]=t
  z+= String.fromCharCode((text.charCodeAt(x) ^ s[(s[x2] + s[y]) % 256]))
 }
 return z
}

