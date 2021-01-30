package webserver

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"testing"
)

/*
这里的程序为一种数据形式提供了一个更好的接口：
给定一小段文本，它会在图表服务器上调用以产生QR码，
即编码文本的盒子矩阵。
该图像可以用手机的摄像头捕获，并解释为例如URL，
从而省去了在手机的小键盘上键入URL的麻烦。
*/
// 访问连接  http://localhost:1718/
var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func TestServer(t *testing.T) {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
    <input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
    <input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`

/*
最主要的部分应该易于理解。一个标志为我们的服务器设置默认的HTTP端口。
模板变量templ是有趣的地方。它构建了一个HTML模板，该模板将由服务器执行以显示页面。

主要功能解析标志，并使用我们上面讨论的机制将功能QR绑定到服务器的根路径。
然后调用http.ListenAndServe启动服务器；服务器运行时会阻塞。

QR只是接收包含表单数据的请求，并以名为s的表单值对数据执行模板。

模板包html / template功能强大；该程序仅涉及其功能。
本质上，它通过替换从传递给templ.Execute的数据项派生的元素来即时重写HTML文本，
在这种情况下为表单值。在模板文本（templateStr）中，双括号分隔的部分表示模板动作。
从{{if。}}到{{end}}的那一部分仅在当前数据项的值称为时执行。
（点），非空。即，当字符串为空时，该模板部分被抑制。

这两个摘要{{。}}表示要在网页上显示提供给模板的数据（查询字符串）。
HTML模板包会自动提供适当的转义符，因此可以安全地显示文本。

模板字符串的其余部分只是页面加载时显示的HTML。
如果解释太快，请参阅模板包的文档以进行更全面的讨论。

在那里，您可以找到：一个有用的Web服务器，其中包含几行代码以及一些数据驱动的HTML文本。
Go的功能强大到足以在几行中完成很多事情。
*/
