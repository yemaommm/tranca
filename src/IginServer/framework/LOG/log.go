package LOG

import (
	// "fmt"
	"IginServer/framework/API/R"
	// "net/http"
	"strings"
	// "github.com/go-martini/martini"
	"IginServer/conf"
	"IginServer/lib/Imartini"
	// "IginServer/lib/redis"
	"encoding/json"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

var LOG_ID map[string]int

type Ixml struct {
	Id    string `xml:"id"`
	Title string `xml:"title"`
	Msg   string `xml:"msg"`
}

func send(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var v Ixml
	xml.Unmarshal(body, &v)
	// log.Printf("%s\n%s", body, v)
	if LOG_ID[v.Id] != 1 {
		R.Write(c.Writer, map[string]interface{}{
			"status":  "500",
			"message": "id不存在",
			"data":    nil,
		})
	} else if v.Id == "" || v.Title == "" || v.Msg == "" {
		R.Write(c.Writer, map[string]interface{}{
			"status":  "300",
			"message": "参数不正确",
			"data":    nil,
		})
	} else if strings.Index(v.Title, "/") != -1 || strings.Index(v.Title, ".") != -1 || strings.Index(v.Title, "~") != -1 {
		R.Write(c.Writer, map[string]interface{}{
			"status":  "305",
			"message": "id,title参数不能带有'/'，'.'，'~'等字符",
			"data":    nil,
		})
	}

	Imartini.MyLog.Other(v.Id+"/"+v.Title+"/", "%v:%v\n%v", v.Id, v.Title, v.Msg)
	R.Write(c.Writer, map[string]interface{}{
		"status":  "200",
		"message": "",
		"data":    "SUCCESS",
	})
}

func sendjson(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var v Ixml
	json.Unmarshal(body, &v)
	// log.Printf("%s\n%s", body, v)
	if LOG_ID[v.Id] != 1 {
		R.Write(c.Writer, map[string]interface{}{
			"status":  "500",
			"message": "id不存在",
			"data":    nil,
		})
	} else if v.Id == "" || v.Title == "" || v.Msg == "" {
		R.Write(c.Writer, map[string]interface{}{
			"status":  "300",
			"message": "参数不正确",
			"data":    nil,
		})
	} else if strings.Index(v.Title, "/") != -1 || strings.Index(v.Title, ".") != -1 || strings.Index(v.Title, "~") != -1 {
		R.Write(c.Writer, map[string]interface{}{
			"status":  "305",
			"message": "id,title参数不能带有'/'，'.'，'~'等字符",
			"data":    nil,
		})
	} else {
		Imartini.MyLog.Other(v.Id+"/"+v.Title+"/", "%v:%v\n%v", v.Id, v.Title, v.Msg)
		R.Write(c.Writer, map[string]interface{}{
			"status":  "200",
			"message": "",
			"data":    "SUCCESS",
		})
	}

}

func test(c *gin.Context) {
	c.Data(200, "text/html", []byte(`
    <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd"> 
	<html xmlns="http://www.w3.org/1999/xhtml" > 
		<head> 
			<title>js操作Xml(向服务器发送Xml,处理服务器返回的Xml)(IE下有效)</title> 
			<script type="text/javascript"><!-- 
				function html_encode(str)   
				{   
				  var s = "";   
				  if (str.length == 0) return "";   
				  s = str.replace(/</g, "&lt;");   
				  s = s.replace(/>/g, "&gt;");   
				  s = s.replace(/ /g, "&nbsp;");   
				  s = s.replace(/\'/g, "&#39;");   
				  s = s.replace(/\"/g, "&quot;");   
				  s = s.replace(/\n/g, "<br>");   
				  s = s.replace(/&/g, "&amp;");   
				  return s;   
				}   
				 
				function html_decode(str)   
				{   
				  var s = "";   
				  if (str.length == 0) return "";   
				  s = str.replace(/&lt;/g, "<");   
				  s = s.replace(/&gt;/g, ">");   
				  s = s.replace(/&nbsp;/g, " ");   
				  s = s.replace(/&#39;/g, "\'");   
				  s = s.replace(/&quot;/g, "\"");   
				  s = s.replace(/<br>/g, "\n");   
				  s = s.replace(/&amp;/g, "&");   
				  return s;   
				}
				var xmlHttp = null;//XmlHttp对象,Ajax核心 
				//创建一个Xml文档,向服务器发送. 
				function f(data){ 
					// var xmlDoc = new ActiveXObject("Msxml2.DOMDocument.3.0");//1创建xml对象,Active控件. 
					// xmlDoc.async = false;//设置异步还是非异步 
					// xmlDoc.loadXML(data); 

					sendXml( data,"/M/log/send"); 
				} 
				//向服务器发送Xml文档 
				function sendXml(xmlDoc,serverURL){ 
					// if (window.XMLHttpRequest) {
					  xmlHttp = new XMLHttpRequest();
					// } else {
					//   var MSXML = new Array('MSXML2.XMLHTTP.5.0', 'MSXML2.XMLHTTP.4.0', 'MSXML2.XMLHTTP.3.0', 'MSXML2.XMLHTTP', 'Microsoft.XMLHTTP');
					//   for(var n = 0; n < MSXML.length; n ++) {
					//     try {
					//         xmlHttp = new ActiveXObject(MSXML[n]);
					//         break;
					//     } catch(e) {
					//     }
					//   }
					// }

					xmlHttp.open ("POST",serverURL ,true);//第三个参数如果为真,则调用onreadystatechange属性指定的回调函数。 
					xmlHttp.onreadystatechange=getData; 
					xmlHttp.send(xmlDoc);//向服务器发传的数据. 
				} 
				function getData(){ 
					if (xmlHttp.readyState==4) //状态为4表示完成. 
					{ 
						var strxml=xmlHttp.responseText;//取得返回的Xml 
						alert(strxml); 
					} 
				} 

			// --></script> 
		</head> 
	<body> 
	<textarea style="width:300px;height:300px;" id="dd"><xml>
<id><![CDATA[id]]></id>
<title><![CDATA[title]]></title>
<msg><![CDATA[msg]]></msg>
</xml></textarea>
	<input type="button" onclick="f(document.getElementById('dd').value);" value="request" /> 
	<input type="button" onclick="f('<xml><id><![CDATA[android]]></id><title><![CDATA[title]]></title><msg><![CDATA[msg]]></msg></xml>');" value="request" /> 
	</body> 
	</html>
`))
}

func init() {
	LOG_ID = make(map[string]int)
	var istmp []string
	stmp := conf.GET["api"]["LOG_ID"]
	json.Unmarshal([]byte(stmp), &istmp)
	for _, i := range istmp {
		LOG_ID[i] = 1
	}
}
