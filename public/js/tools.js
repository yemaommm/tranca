config = {};
config.apihost = 'http://xt.yundapian.com';
config.serverhost = 'http://xt.yundapian.com';
/**
 *create by 2012-08-25 pm 17:48
 *@author hexinglun@gmail.com
 *BASE64 Encode and Decode By UTF-8 unicode
 *可以和java的BASE64编码和解码互相转化
 */
(function(){
    var BASE64_MAPPING = [
        'A','B','C','D','E','F','G','H',
        'I','J','K','L','M','N','O','P',
        'Q','R','S','T','U','V','W','X',
        'Y','Z','a','b','c','d','e','f',
        'g','h','i','j','k','l','m','n',
        'o','p','q','r','s','t','u','v',
        'w','x','y','z','0','1','2','3',
        '4','5','6','7','8','9','+','/'
    ];

    /**
     *ascii convert to binary
     */
    var _toBinary = function(ascii){
        var binary = new Array();
        while(ascii > 0){
            var b = ascii%2;
            ascii = Math.floor(ascii/2);
            binary.push(b);
        }
        /*
        var len = binary.length;
        if(6-len > 0){
            for(var i = 6-len ; i > 0 ; --i){
                binary.push(0);
            }
        }*/
        binary.reverse();
        return binary;
    };

    /**
     *binary convert to decimal
     */
    var _toDecimal  = function(binary){
        var dec = 0;
        var p = 0;
        for(var i = binary.length-1 ; i >= 0 ; --i){
            var b = binary[i];
            if(b == 1){
                dec += Math.pow(2 , p);
            }
            ++p;
        }
        return dec;
    };

    /**
     *unicode convert to utf-8
     */
    var _toUTF8Binary = function(c , binaryArray){
        var mustLen = (8-(c+1)) + ((c-1)*6);
        var fatLen = binaryArray.length;
        var diff = mustLen - fatLen;
        while(--diff >= 0){
            binaryArray.unshift(0);
        }
        var binary = [];
        var _c = c;
        while(--_c >= 0){
            binary.push(1);
        }
        binary.push(0);
        var i = 0 , len = 8 - (c+1);
        for(; i < len ; ++i){
            binary.push(binaryArray[i]);
        }

        for(var j = 0 ; j < c-1 ; ++j){
            binary.push(1);
            binary.push(0);
            var sum = 6;
            while(--sum >= 0){
                binary.push(binaryArray[i++]);
            }
        }
        return binary;
    };

    var __BASE64 = {
            /**
             *BASE64 Encode
             */
            encoder:function(str){
                var base64_Index = [];
                var binaryArray = [];
                for(var i = 0 , len = str.length ; i < len ; ++i){
                    var unicode = str.charCodeAt(i);
                    var _tmpBinary = _toBinary(unicode);
                    if(unicode < 0x80){
                        var _tmpdiff = 8 - _tmpBinary.length;
                        while(--_tmpdiff >= 0){
                            _tmpBinary.unshift(0);
                        }
                        binaryArray = binaryArray.concat(_tmpBinary);
                    }else if(unicode >= 0x80 && unicode <= 0x7FF){
                        binaryArray = binaryArray.concat(_toUTF8Binary(2 , _tmpBinary));
                    }else if(unicode >= 0x800 && unicode <= 0xFFFF){//UTF-8 3byte
                        binaryArray = binaryArray.concat(_toUTF8Binary(3 , _tmpBinary));
                    }else if(unicode >= 0x10000 && unicode <= 0x1FFFFF){//UTF-8 4byte
                        binaryArray = binaryArray.concat(_toUTF8Binary(4 , _tmpBinary));    
                    }else if(unicode >= 0x200000 && unicode <= 0x3FFFFFF){//UTF-8 5byte
                        binaryArray = binaryArray.concat(_toUTF8Binary(5 , _tmpBinary));
                    }else if(unicode >= 4000000 && unicode <= 0x7FFFFFFF){//UTF-8 6byte
                        binaryArray = binaryArray.concat(_toUTF8Binary(6 , _tmpBinary));
                    }
                }

                var extra_Zero_Count = 0;
                for(var i = 0 , len = binaryArray.length ; i < len ; i+=6){
                    var diff = (i+6)-len;
                    if(diff == 2){
                        extra_Zero_Count = 2;
                    }else if(diff == 4){
                        extra_Zero_Count = 4;
                    }
                    //if(extra_Zero_Count > 0){
                    //  len += extra_Zero_Count+1;
                    //}
                    var _tmpExtra_Zero_Count = extra_Zero_Count;
                    while(--_tmpExtra_Zero_Count >= 0){
                        binaryArray.push(0);
                    }
                    base64_Index.push(_toDecimal(binaryArray.slice(i , i+6)));
                }

                var base64 = '';
                for(var i = 0 , len = base64_Index.length ; i < len ; ++i){
                    base64 += BASE64_MAPPING[base64_Index[i]];
                }

                for(var i = 0 , len = extra_Zero_Count/2 ; i < len ; ++i){
                    base64 += '=';
                }
                return base64;
            },
            /**
             *BASE64  Decode for UTF-8 
             */
            decoder : function(_base64Str){
                var _len = _base64Str.length;
                var extra_Zero_Count = 0;
                /**
                 *计算在进行BASE64编码的时候，补了几个0
                 */
                if(_base64Str.charAt(_len-1) == '='){
                    //alert(_base64Str.charAt(_len-1));
                    //alert(_base64Str.charAt(_len-2));
                    if(_base64Str.charAt(_len-2) == '='){//两个等号说明补了4个0
                        extra_Zero_Count = 4;
                        _base64Str = _base64Str.substring(0 , _len-2);
                    }else{//一个等号说明补了2个0
                        extra_Zero_Count = 2;
                        _base64Str = _base64Str.substring(0 , _len - 1);
                    }
                }

                var binaryArray = [];
                for(var i = 0 , len = _base64Str.length; i < len ; ++i){
                    var c = _base64Str.charAt(i);
                    for(var j = 0 , size = BASE64_MAPPING.length ; j < size ; ++j){
                        if(c == BASE64_MAPPING[j]){
                            var _tmp = _toBinary(j);
                            /*不足6位的补0*/
                            var _tmpLen = _tmp.length;
                            if(6-_tmpLen > 0){
                                for(var k = 6-_tmpLen ; k > 0 ; --k){
                                    _tmp.unshift(0);
                                }
                            }
                            binaryArray = binaryArray.concat(_tmp);
                            break;
                        }
                    }
                }

                if(extra_Zero_Count > 0){
                    binaryArray = binaryArray.slice(0 , binaryArray.length - extra_Zero_Count);
                }

                var unicode = [];
                var unicodeBinary = [];
                for(var i = 0 , len = binaryArray.length ; i < len ; ){
                    if(binaryArray[i] == 0){
                        unicode=unicode.concat(_toDecimal(binaryArray.slice(i,i+8)));
                        i += 8;
                    }else{
                        var sum = 0;
                        while(i < len){
                            if(binaryArray[i] == 1){
                                ++sum;
                            }else{
                                break;
                            }
                            ++i;
                        }
                        unicodeBinary = unicodeBinary.concat(binaryArray.slice(i+1 , i+8-sum));
                        i += 8 - sum;
                        while(sum > 1){
                            unicodeBinary = unicodeBinary.concat(binaryArray.slice(i+2 , i+8));
                            i += 8;
                            --sum;
                        }
                        unicode = unicode.concat(_toDecimal(unicodeBinary));
                        unicodeBinary = [];
                    }
                }
                var str = '';
                for(var i = 0 , len =  unicode.length ; i < len ;++i){  
                      str += String.fromCharCode(unicode[i]);  
                }
                return str;
            }
    };

    window.BASE64 = __BASE64;
})();


function ajax(options) {
    options = options || {};
    options.sync = options.sync == undefined?true:options.sync;
    options.type = (options.type || "GET").toUpperCase();
    options.dataType = options.dataType || "json";
    var params = formatParams(options.data);

    //创建 - 非IE6 - 第一步
    if (window.XMLHttpRequest) {
        var xhr = new XMLHttpRequest();
    } else { //IE6及其以下版本浏览器
        var xhr = new ActiveXObject('Microsoft.XMLHTTP');
    }

    //接收 - 第三步
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            var status = xhr.status;
            if (status >= 200 && status < 300) {
                options.success && options.success(xhr.responseText, xhr.responseXML);
            } else {
                options.fail && options.fail(status);
            }
        }
    }

    //连接 和 发送 - 第二步
    if (options.type == "GET") {
        if(options.url.indexOf("?")==-1){options.url = options.url + "?" + params;}else{options.url = options.url + "&" + params;}
        xhr.open("GET", options.url, options.sync);
        xhr.send(null);
    } else if (options.type == "POST") {
        xhr.open("POST", options.url, options.sync);


        if(options.form != undefined){
            var form = new FormData(options.form);
            xhr.send(form);
        }else if(options.formdata != undefined){
            var form = options.formdata;
            xhr.send(form);
        }else{
            //设置表单提交时的内容类型
            xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
            xhr.send(params);
        }
    }
}
//格式化参数
function formatParams(data) {
    var arr = [];
    for (var name in data) {
        arr.push(encodeURIComponent(name) + "=" + encodeURIComponent(data[name]));
    }
    arr.push(("v=" + Math.random()).replace(".",""));
    return arr.join("&");
}

//写cookies 
function setCookie(name,value) 
{ 
    var Days = 30; 
    var exp = new Date(); 
    exp.setTime(exp.getTime() + Days*24*60*60*1000); 
    document.cookie = name + "="+ escape (value) + ";expires=" + exp.toGMTString(); 
}
//读取cookies
function getCookie(name) 
{ 
    var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)"); 
    if(arr=document.cookie.match(reg)) 
        return unescape(arr[2]); 
    else
        return ''; 
}

function param_url(url, param){
    if(url.split('?').length <2){
        url += '?';
    }
    if(url[url.length - 1] != '?' && url[url.length - 1] != '&'){
        url += '&';
    }
    for(i in param){
        url += i + "=" + param[i].toString() + "&"
    }
    return url.substring(0, url.length - 1)
}

//字符串格式化
printf = function() {
    var num = arguments.length;
    var oStr = arguments[0];
    for (var i = 1; i < num; i++) {
        var pattern = "\\{" + (i-1) + "\\}";
        var re = new RegExp(pattern, "g");
        oStr = oStr.replace(re, arguments[i]);
    }
    return oStr;
}

// 向标签内添加html
function appendElement(h, str){
    var el = document.createElement("div");
    el.innerHTML = str;
    for(i = 0;i < el.childNodes.length;i++){
        // console.log(el.childNodes[i])
        h.appendChild(el.childNodes[i]);
    }
}

//get参数获取
function request(paras) {
    var url = location.href;
    var paraString = url.substring(url.indexOf("?") + 1, url.length).split("&");
    var paraObj = {}
    for (i = 0; j = paraString[i]; i++) {
        paraObj[j.substring(0, j.indexOf("=")).toLowerCase()] = j.substring(j.indexOf("=") + 1, j.length);
    }
    var returnValue = paraObj[paras.toLowerCase()];
    if (typeof (returnValue) == "undefined") {
        return "";
    } else {
        //t=new String(returnValue.getBytes("ISO8859_1"),"UTF-8");
        //return t;
        return returnValue;
    }
}

//图片上传预览    IE是用了滤镜。
function previewImage(file, mid, defaultimg){
    if(file.value == ""){
        document.getElementById(mid).src = defaultimg;
    }else if (file.files && file.files[0]){
        var img = document.getElementById(mid);
        var reader = new FileReader();
        reader.onload = function(evt){img.src = evt.target.result;}
        reader.readAsDataURL(file.files[0]);
    }else{ //兼容IE
        var sFilter='filter:progid:DXImageTransform.Microsoft.AlphaImageLoader(sizingMethod=scale,src="';
        file.select();
        var src = document.selection.createRange().text;
        div.innerHTML = '<img id=imghead>';
        var img = document.getElementById('imghead');
        img.filters.item('DXImageTransform.Microsoft.AlphaImageLoader').src = src;
        var rect = clacImgZoomParam(MAXWIDTH, MAXHEIGHT, img.offsetWidth, img.offsetHeight);
        status =('rect:'+rect.top+','+rect.left+','+rect.width+','+rect.height);
        div.innerHTML = "<div id=divhead style='width:"+rect.width+"px;height:"+rect.height+"px;margin-top:"+rect.top+"px;"+sFilter+src+"\"'></div>";
    }
}
function clacImgZoomParam( maxWidth, maxHeight, width, height ){
    var param = {top:0, left:0, width:width, height:height};
    if( width>maxWidth || height>maxHeight )
    {
        rateWidth = width / maxWidth;
        rateHeight = height / maxHeight;
         
        if( rateWidth > rateHeight )
        {
            param.width =  maxWidth;
            param.height = Math.round(height / rateWidth);
        }else
        {
            param.width = Math.round(width / rateHeight);
            param.height = maxHeight;
        }
    }
     
    param.left = Math.round((maxWidth - param.width) / 2);
    param.top = Math.round((maxHeight - param.height) / 2);
    return param;
}

//判断obj是否为json对象  
function isJson(obj){  
    var isjson = typeof(obj) == "object" && Object.prototype.toString.call(obj).toLowerCase() == "[object object]" && !obj.length;   
    return isjson;  
}

tools = {}
tools.indexhtml = 'index.html';
tools.loginhtml = 'login.html';
tools.registerhtml = 'sign.html';
tools.host = config.apihost;//'http://121.41.116.104:8003';

tools.authinit = function(b){  // 参数，是否要登录验证进行登录跳转
    if(b == undefined) b = false;
    if(b){
        tools.auth()
    }
    userinfo = tools.getmsg();
    token = userinfo.token;
    username = userinfo.username;

    if (token == undefined || token == ''){
        return // 未登录
    }
    // alert('已登录，更改相关信息') // 已登录
    pc = '<span class="font_18 font_999999"><a href="javascript:tools.logout();">退出</a></span><span class="font_18 font_333333">{0}</span>';
    phone = '<span class="font_18 font_999999"><a href="javascript:tools.logout();">退出登录</a></span>';
    document.getElementsByClassName("titleDiv_nav")[0].innerHTML = printf(pc, username);
    document.getElementsByClassName("titleDiv_nav_phone")[0].innerHTML = phone;
}

tools.auth = function(){
    userinfo = tools.getmsg();
    token = userinfo.token;
    username = userinfo.username;

    if (token == undefined || token == ''){
        window.location.href = param_url(tools.loginhtml, {'re_url':encodeURIComponent(window.location.href)})
        return ''
    }
    tools.authpost("/api/auth/", {})

    return username, token
}

tools.login = function(username, password){
    tools.post("/api/auth/login", {'username':username, 'password':password}, function(data){
        tools.setmsg(data);
        tools.alert('登录成功', function(){
            re_url = request('re_url');
            if(re_url == ''){
                window.location.href = tools.indexhtml;
            }else{
                window.location.href = decodeURIComponent(re_url);
            }
        });
    });
}

tools.logout = function(){
    tools.authpost("/api/auth/logout", {}, function(data){
        tools.setmsg({});
        window.location.reload()
        // tools.alert('退出登录成功', function(){
        //     window.location.reload()
        // });
    });
}

tools.register = function(username, password, code){
    tools.post("/api/auth/register", {'username':username, 'password':password, 'code':code}, function(data){
        tools.alert("注册成功，请返回登录页进行登录", function(){
            window.location.href = tools.loginhtml;
        });
    });
}

tools.getCode = function(type, phone, callback){
    tools.post("/api/sms/send", {'phone':phone, 'mode':type}, callback);
}

tools.setmsg = function(jsondata){ // errorDiv
    str = JSON.stringify(jsondata)
    setCookie("userinfo", BASE64.encoder(str))
    // console.log(str)
    // console.log(BASE64.encoder(str))
    // console.log(BASE64.decoder(''))
}

tools.getmsg = function(){ // errorDiv
    str = getCookie("userinfo");
    try{
        return eval("(" + BASE64.decoder(str) + ")");
    }catch(e){
        return {}
    }
    // console.log(str)
    // console.log(BASE64.encoder(str))
    // console.log(BASE64.decoder(''))
}

tools.alert = function(msg, callback){
    if(tools.tools_alert == null){
        appendElement(document.getElementsByTagName('body')[0], "<div id='tools_alert'></div>");
        tools.tools_alert = document.getElementById('tools_alert')
    }
    msg_div = "<div class='errorWDiv'><div class='errorDiv'><span>{0}</span></div></div>"
    appendElement(tools.tools_alert, printf(msg_div, msg.toString()));

    fn = function(callback){
        return function(){
            tools_alert = tools.tools_alert.getElementsByClassName("errorWDiv");
            if(tools_alert.length > 0){
                if(tools_alert[0].remove != undefined){
                    tools_alert[0].remove()
                }else{
                    tools.tools_alert.removeChild(tools_alert[0]);
                }
            }
            if(typeof callback == 'function') callback();
        }
    }
    setTimeout(fn(callback), 3000);
    
}

tools.loading = function(){
    stmp = '<div id="loading" class="loadNow" background><img src="img/loading/loading.png" class="loadNow_img"></div>';
    appendElement(document.getElementsByTagName("body")[0], stmp);
}

tools.closeloading = function(){
    if(document.getElementById("loading") != null){
        document.getElementById("loading").classList.add('danC');
        setTimeout(function(){
            if(document.getElementById("loading") != null){
                if(document.getElementById("loading").remove != undefined){
                    document.getElementById("loading").remove()
                }else{
                    document.getElementsByTagName('body')[0].removeChild(document.getElementById("loading"));
                }
                tools.closeloading();
            }
        }, 800);
    }
}

tools.get = function(url, param, callback, errorback){
    ajax({
        url: tools.host+url,              //请求地址
        type: "get",                       //请求方式
        data: param,        //请求参数
        dataType: "json",
        success: function (response, xml) {
            json = eval("(" + response + ")");
            console.log(json)

            if (json.status == 200) {
                if(typeof callback == 'function') callback(json.data);
            }else{
                if(typeof errorback == 'function') errorback(json.status, json.error); 
                else if(isJson(errorback) && errorback[json.status] != undefined) errorback[json.status](json.error);
                else if(isJson(errorback) && errorback['other'] != undefined) errorback['other'](json.error);
                else tools.alert(json.error)
            }
        },fail: function (status) {
            // 此处放失败后执行的代码
            console.log(status);
        }
    });
}

tools.post = function(url, param, callback, errorback){
    ajax({
        url: tools.host+url,              //请求地址
        type: "post",                       //请求方式
        data: param,        //请求参数
        dataType: "json",
        success: function (response, xml) {
            json = eval("(" + response + ")");
            console.log(json)

            if (json.status == 200) {
                if(typeof callback == 'function') callback(json.data);
            }else{
                if(typeof errorback == 'function') errorback(json.status, json.error); 
                else if(isJson(errorback) && errorback[json.status] != undefined) errorback[json.status](json.error);
                else if(isJson(errorback) && errorback['other'] != undefined) errorback['other'](json.error);
                else tools.alert(json.error)
            }
            
        },fail: function (status) {
            // 此处放失败后执行的代码
            console.log(status);
        }
    });
}

tools.postfile = function(url, param, callback, errorback){
    ajax({
        url: tools.host+url,              //请求地址
        type: "post",                       //请求方式
        formdata: param,        //请求参数
        // dataType: "json",
        success: function (response, xml) {
            json = eval("(" + response + ")");
            console.log(json)

            if (json.status == 200) {
                if(typeof callback == 'function') callback(json.data);
            }else{
                if(typeof errorback == 'function') errorback(json.status, json.error); 
                else if(isJson(errorback) && errorback[json.status] != undefined) errorback[json.status](json.error);
                else if(isJson(errorback) && errorback['other'] != undefined) errorback['other'](json.error);
                else tools.alert(json.error);
            }
            
        },fail: function (status) {
            // 此处放失败后执行的代码
            console.log(status);
            alert(status)
        }
    });
}

tools.authget = function(url, param, callback, errorback){
    userinfo = tools.getmsg();
    token = userinfo.token;
    username = userinfo.username;

    param.token = token;
    param.username = username;

    eback = {};
    eback[400] = function(){
        // alert('验证失败，进行跳转登录')
        window.location.href = param_url(tools.loginhtml, {'re_url':encodeURIComponent(window.location.href)})
    }
    if (isJson(errorback)){
        for(i in errorback){
            if (typeof errorback[i] == 'function'){
                eback[i] = errorback[i];
            }
        }
    }else{
        eback['other'] = errorback;   
    }
    tools.get(url, param, callback, eback);
}

tools.authpost = function(url, param, callback, errorback){
    userinfo = tools.getmsg();
    token = userinfo.token;
    username = userinfo.username;

    param.token = token;
    param.username = username;

    eback = {};
    eback[400] = function(){
        // alert('验证失败，进行跳转登录')
        window.location.href = param_url(tools.loginhtml, {'re_url':encodeURIComponent(window.location.href)})
    }
    if (isJson(errorback)){
        for(i in errorback){
            if (typeof errorback[i] == 'function'){
                eback[i] = errorback[i];
            }
        }
    }else{
        eback['other'] = errorback;   
    }
    tools.post(url, param, callback, eback);
}

tools.authpostfile = function(url, param, callback, errorback){
    userinfo = tools.getmsg();
    token = userinfo.token;
    username = userinfo.username;

    param.append('token', token);
    param.append('username', username);

    eback = {};
    eback[400] = function(){
        // alert('验证失败，进行跳转登录')
        window.location.href = param_url(tools.loginhtml, {'re_url':encodeURIComponent(window.location.href)})
    }
    if (isJson(errorback)){
        for(i in errorback){
            if (typeof errorback[i] == 'function'){
                eback[i] = errorback[i];
            }
        }
    }else{
        eback['other'] = errorback;   
    }
    tools.postfile(url, param, callback, eback);
}

tools.isPC = function() {  
    var userAgentInfo = window.navigator.userAgent.toLowerCase();
    var Agents = new Array("Android", "iPhone", "SymbianOS", "Windows Phone", "iPad", "iPod");
    var flag = true;
    for (var v = 0; v < Agents.length; v++) {
        if (userAgentInfo.indexOf(Agents[v].toLowerCase()) > 0) { flag = false; break; }
    }
    return flag;
}

tools.isWeiXin = function(){
    var ua = window.navigator.userAgent.toLowerCase();
    if(ua.match(/MicroMessenger/i) == 'micromessenger'){
        return true;
    }else{
        return false;
    }
}

tools.getProductList = function(callback){
    tools.get("/api/product/list", {}, callback);
}

tools.getProductQRCode = function(order_info_id, callback){
    tools.get("/api/weixin/GetQRcodeUrl", {"order_info_id":order_info_id}, callback);
}

tools.saveProductOrder = function(form, callback, errorback){
    tools.authpostfile("/api/product/createinfo", form, callback, errorback);
}

tools.weixinLogin = function(){
    if(tools.isWeiXin()){
        userinfo = tools.getmsg();
        token = userinfo.token;
        tools.get("/weixin/GetWeixinLoginInfo", {"token":token}, function(data){
            if(data == null){
                userinfo = tools.getmsg();
                token = userinfo.token;
                var url = tools.host+'/weixin/ApiBaseWeixinLogin?token={0}&re_url={1}';
                url = printf(url, token, encodeURIComponent(window.location.href));

                window.location.href = url;
            }
        });
    }
}

tools.weixinJsPay = function(order_info_id, callback, errorback){
    userinfo = tools.getmsg();
    token = userinfo.token;
    var url = "/api/weixin/GetJsConfig";
    tools.get(url, {"token":token, "order_info_id":order_info_id}, function(data){
        WeixinJSBridge.invoke(
            'getBrandWCPayRequest', {
                "appId" : data.appId,     //公众号名称，由商户传入     
                "timeStamp":data.timeStamp,         //时间戳，自1970年以来的秒数     
                "nonceStr" : data.nonceStr, //随机串     
                "package" : data.package,     
                "signType" : data.signType,         //微信签名方式:     
                "paySign" : data.paySign //微信签名 
            },
            function(res){     
                if(res.err_msg == "get_brand_wcpay_request:ok" ) {
                    alert('支付完成')
                    tools.alert('支付完成');
                    callback();
                }else{     // 使用以上方式判断前端返回,微信团队郑重提示:res.err_msg将在用户支付成功后返回    ok，但并不保证它绝对可靠。 
                    errorback();
                }
            }
        ); 
    });
}

tools.weixinShare = function(title, desc, link, imgUrl){
    var stmp = document.getElementById('weixinscript');
    if(stmp == null){
        var run = 'tools.weixinShare("'+title+'", "'+desc+'", "'+link+'", "'+imgUrl+'")';
        var head= document.getElementsByTagName('head')[0]; 
        var script= document.createElement('script'); 
        script.id = 'weixinscript';
        script.type= 'text/javascript'; 
        script.src= 'http://res.wx.qq.com/open/js/jweixin-1.0.0.js'; 
        script.onload = function(){tools.weixinShare(title, desc, link, imgUrl);};
        head.appendChild(script); 

        return;
    }
    var url = "/weixin/info";
    tools.get(url, {"url":encodeURIComponent(window.location.href)}, function(data){
        wx.checkJsApi({
            jsApiList: [
                "onMenuShareTimeline",
                "onMenuShareAppMessage",
                "onMenuShareQQ",
                "onMenuShareWeibo",
                "onMenuShareQZone"
            ],
            success: function(res) {
                console.log(JSON.stringify(res));
            }
        });
        wx.config({
            debug: false, // 开启调试模式,调用的所有api的返回值会在客户端alert出来，若要查看传入的参数，可以在pc端打开，参数信息会通过log打出，仅在pc端时才会打印。
            appId: data.appId, // 必填，公众号的唯一标识
            timestamp: data.timestamp, // 必填，生成签名的时间戳
            nonceStr: data.noncestr, // 必填，生成签名的随机串
            signature: data.signature, // 必填，签名，见附录1
            jsApiList: ['onMenuShareAppMessage', 'onMenuShareTimeline', 'onMenuShareQQ', 'onMenuShareWeibo', 'onMenuShareQZone'] // 必填，需要使用的JS接口列表，所有JS接口列表见附录2
        });

        wx.error(function(res){
            // config信息验证失败会执行error函数，如签名过期导致验证失败，具体错误信息可以打开config的debug模式查看，也可以在返回的res参数中查看，对于SPA可以在这里更新签名。
            alert(JSON.stringify(res))
        });

        wx.ready(function() {
            wx.onMenuShareAppMessage({
                title: title, // 分享标题
                desc: desc, // 分享描述
                link: link, // 分享链接
                imgUrl: imgUrl, // 分享图标
                type: 'link', // 分享类型,music、video或link，不填默认为link
                dataUrl: '', // 如果type是music或video，则要提供数据链接，默认为空
                success: function () { 
                    // 用户确认分享后执行的回调函数
                },
                cancel: function () { 
                    // 用户取消分享后执行的回调函数
                }
            });
            wx.onMenuShareTimeline({
                title: title, // 分享标题
                desc: desc, // 分享描述
                link: link, // 分享链接
                imgUrl: imgUrl, // 分享图标
                // type: '', // 分享类型,music、video或link，不填默认为link
                // dataUrl: '', // 如果type是music或video，则要提供数据链接，默认为空
                success: function () { 
                    // 用户确认分享后执行的回调函数
                },
                cancel: function () { 
                    // 用户取消分享后执行的回调函数
                }
            });
            wx.onMenuShareQQ({
                title: title, // 分享标题
                desc: desc, // 分享描述
                link: link, // 分享链接
                imgUrl: imgUrl, // 分享图标
                // type: '', // 分享类型,music、video或link，不填默认为link
                // dataUrl: '', // 如果type是music或video，则要提供数据链接，默认为空
                success: function () { 
                    // 用户确认分享后执行的回调函数
                },
                cancel: function () { 
                    // 用户取消分享后执行的回调函数
                }
            });
            wx.onMenuShareWeibo({
                title: title, // 分享标题
                desc: desc, // 分享描述
                link: link, // 分享链接
                imgUrl: imgUrl, // 分享图标
                // type: '', // 分享类型,music、video或link，不填默认为link
                // dataUrl: '', // 如果type是music或video，则要提供数据链接，默认为空
                success: function () { 
                    // 用户确认分享后执行的回调函数
                },
                cancel: function () { 
                    // 用户取消分享后执行的回调函数
                }
            });
            wx.onMenuShareQZone({
                title: title, // 分享标题
                desc: desc, // 分享描述
                link: link, // 分享链接
                imgUrl: imgUrl, // 分享图标
                // type: '', // 分享类型,music、video或link，不填默认为link
                // dataUrl: '', // 如果type是music或video，则要提供数据链接，默认为空
                success: function () { 
                    // 用户确认分享后执行的回调函数
                },
                cancel: function () { 
                    // 用户取消分享后执行的回调函数
                }
            });
        });
    });
}
// tools.weixinShare("8999直播销售产品", "8999元可以请20位美女网红主播，三天销售您的产品，让您的产品在几百万粉丝里面卖起来！", config.serverhost+"/tranca/", "http://sc.shanghaitong.biz//upload/3/2017/March/7/7a8a814567c2b3589a6c9237f99ddd52.jpg");