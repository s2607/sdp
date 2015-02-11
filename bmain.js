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

