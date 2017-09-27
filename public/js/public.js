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
//get参数获取
function GetQueryString(name) {
     var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)");
     var r = window.location.search.substr(1).match(reg);
     if(r!=null)return  unescape(r[2]); return null;
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
//删除cookies
function delCookie(name) 
{ 
    var exp = new Date(); 
    exp.setTime(exp.getTime() - 1); 
    var cval=getCookie(name); 
    if(cval!=null) 
        document.cookie= name + "="+cval+";expires="+exp.toGMTString(); 
} 

//去掉空格
function Trim(str){ 
    return str.replace(/(^\s*)|(\s*$)/g, ""); 
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

// 向标签内添加html
function appendElement(h, str){
    var el = document.createElement("div");
    el.innerHTML = str;
    for(i = 0;i < el.childNodes.length;i++){
        // console.log(el.childNodes[i])
        h.appendChild(el.childNodes[i]);
    }
}

//判断obj是否为json对象  
function isJson(obj){  
    var isjson = typeof(obj) == "object" && Object.prototype.toString.call(obj).toLowerCase() == "[object object]" && !obj.length;   
    return isjson;  
}

// url加参数
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

//JSON数据格式化显示
function jsformat(txt,compress/*是否为压缩模式*/){/* 格式化JSON源码(对象转换为JSON文本) */  
    var indentChar = '  ';   
    if(/^\s*$/.test(txt)){   
        alert('数据为空,无法格式化! ');   
        return;   
    }   
    try{var data=eval('('+txt+')');}   
    catch(e){   
        alert('数据源语法错误,格式化失败! 错误信息: '+e.description,'err');   
        return;   
    };   
    var draw=[],last=false,This=this,line=compress?'':'\n',nodeCount=0,maxDepth=0;   
       
    var notify=function(name,value,isLast,indent/*缩进*/,formObj){   
        nodeCount++;/*节点计数*/  
        for (var i=0,tab='';i<indent;i++ )tab+=indentChar;/* 缩进HTML */  
        tab=compress?'':tab;/*压缩模式忽略缩进*/  
        maxDepth=++indent;/*缩进递增并记录*/  
        if(value&&value.constructor==Array){/*处理数组*/  
            draw.push(tab+(formObj?('"'+name+'":'):'')+'['+line);/*缩进'[' 然后换行*/  
            for (var i=0;i<value.length;i++)   
                notify(i,value[i],i==value.length-1,indent,false);   
            draw.push(tab+']'+(isLast?line:(','+line)));/*缩进']'换行,若非尾元素则添加逗号*/  
        }else   if(value&&typeof value=='object'){/*处理对象*/  
                draw.push(tab+(formObj?('"'+name+'":'):'')+'{'+line);/*缩进'{' 然后换行*/  
                var len=0,i=0;   
                for(var key in value)len++;   
                for(var key in value)notify(key,value[key],++i==len,indent,true);   
                draw.push(tab+'}'+(isLast?line:(','+line)));/*缩进'}'换行,若非尾元素则添加逗号*/  
            }else{   
                    if(typeof value=='string')value='"'+value+'"';   
                    draw.push(tab+(formObj?('"'+name+'":'):'')+value+(isLast?'':',')+line);   
            };   
    };   
    var isLast=true,indent=0;   
    notify('',data,isLast,indent,false);   
    return draw.join('');   
}  

// 对Date的扩展，将 Date 转化为指定格式的String
// 月(M)、日(d)、小时(h)、分(m)、秒(s)、季度(q) 可以用 1-2 个占位符， 
// 年(y)可以用 1-4 个占位符，毫秒(S)只能用 1 个占位符(是 1-3 位的数字) 
// 例子： 
// (new Date()).Format("yyyy-MM-dd hh:mm:ss.S") ==> 2006-07-02 08:09:04.423 
// (new Date()).Format("yyyy-M-d h:m:s.S")      ==> 2006-7-2 8:9:4.18 
Date.prototype.Format = function (fmt) { //author: meizz 
    var o = {
        "M+": this.getMonth() + 1, //月份 
        "d+": this.getDate(), //日 
        "h+": this.getHours(), //小时 
        "m+": this.getMinutes(), //分 
        "s+": this.getSeconds(), //秒 
        "q+": Math.floor((this.getMonth() + 3) / 3), //季度 
        "S": this.getMilliseconds() //毫秒 
    };
    if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
    for (var k in o)
    if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
    return fmt;
}

//时间戳转日期
function getLocalTime(nS) {     
   return new Date(parseInt(nS) * 1000).Format("yyyy-MM-dd hh:mm");     
}

//页码生成
ipage = {}
ipage.create = function(div, total, pp, now){
    if(now==undefined){
        if(GetQueryString("page")==null){
            now = 1;
        }else{
            now = parseInt(GetQueryString("page"));
        }
    }
    if(pp==undefined){
        pp = 10
    }
    total = total%pp>0?parseInt(total/pp)+2:parseInt(total/pp)+1;
    var search = window.location.search;
    search = search.replace(new RegExp(/page=([^&]*)(&|$)/g),'');
    if(search==""){
        search = "?page={0}";
    }else{
        if(search.substr(-1)=="&"||search.substr(-1)=="?"){
            search += "page={0}";
        }else{
            search += "&page={0}";
        }
    }
    var p = '<ul style="float:right;" class="pagination pagination-sm">{0}</ul>';
    var pi = '';
    var tp = now;
    tp = parseInt(tp/pp)*pp;
    if(tp>1){
        pi += printf('<li><a href="{0}">&laquo;</a></li>', printf(search, tp-1));
    }
    for (var i = tp;i<(tp+pp)&&i<total;i++){
        if(i!=0){
            if(i==now){
                pi += printf('<li class="active"><a href="{0}">'+(i).toString()+'</a></li>', printf(search, i));
            }else{
                pi += printf('<li><a href="{0}">'+(i).toString()+'</a></li>', printf(search, i));
            }
        }
    }
    if(tp+pp<total){
        pi += printf('<li><a href="{0}">&raquo;</a></li>', printf(search, (tp+pp)));
    }
    $(div).append(printf(p, pi));
}

ialert = {}
ialert.init = function(body){
    ialert.body = body
}
ialert.body = '<div class="ialert"><div style="width:100%;height:100%;background:rgba(0,0,0,0.5);position: fixed;top:0px;left:0px;z-index:9999;"></div><div style="width:90%;position:fixed;left:5%;top:200px;background:#fff;z-index:10000;box-shadow:0px 0px 6px #e0e0e0;"><div style="width:95%;margin-left:2.5%;font-size:3.5rem;font-weight:bold;line-height:5rem;margin-top:20px;text-align:center;border-bottom:1px solid #e0e0e0;padding-bottom:20px;">{0}</div><div style="width:95%;margin-left:2.5%;font-size:3.5rem;line-height:7rem;text-align:center;margin-top:20px;">{1}</div><div style="margin-bottom:30px;margin-top:30px;"></div><div style="width:100%;"><div style="width:100%;font-size:4rem;line-height:9rem;text-align:center;color:#fff;background:#f2425b;float:left;" class="ialert-click">知道了&nbsp;&nbsp;>&nbsp;></div></div></div></div>';
ialert.open = function(title, body){
    var body = printf(ialert.body, title, body);
    $('body').append(body);
    $('.ialert-click').click(function(){
        $('.ialert').remove();
    });
}
iconfirm = {}
iconfirm.init = function(body){
    iconfirm.body = body
}
iconfirm.body = '<div class="iconfirm"><div style="width:100%;height:100%;background:rgba(0,0,0,0.5);position: fixed;top:0px;left:0px;z-index:9999;"></div><div style="width:90%;position:fixed;left:5%;top:200px;background:#fff;z-index:10000;box-shadow:0px 0px 6px #e0e0e0;"><div style="width:95%;margin-left:2.5%;font-size:3.5rem;font-weight:bold;line-height:5rem;margin-top:20px;text-align:center;border-bottom:1px solid #e0e0e0;padding-bottom:20px;">{0}</div><div style="width:95%;margin-left:2.5%;font-size:3.5rem;line-height:15rem;text-align:center;">{1}</div><div style="margin-bottom:30px;margin-top:30px;"></div><div style="width:100%;"><div style="width:50%;font-size:4rem;line-height:9rem;text-align:center;color:#fff;background:#3c466a;float:left;" class="iconfirm-off">取消</div><div style="width:50%;font-size:4rem;line-height:9rem;text-align:center;color:#fff;background:#f2425b;float:left;" class="iconfirm-ok">确定</div></div></div></div>'
iconfirm.open = function(title, body, func){
    var body = printf(iconfirm.body, title, body);
    $('body').append(body);
    $('.iconfirm-off').click(function(){
        $('.iconfirm').remove();
    });
    $('.iconfirm-ok').click(function(){
        func();
        $('.iconfirm').remove();
    });
}
// function alert(){
//     if(arguments.length==1){
//         ialert.open("提示",arguments[0]);
//     }else{
//         ialert.open(arguments[0],arguments[1]);
//     }
// }
//调用ajax之后填充标签
function plink(t){
    if($(t).attr('plink') != '' && $(t).attr('plink') != undefined){
        aAlert = '<div class="palert" id="aAlert"><img src="/admin/image/loading.gif">加载中……</div>'
        $('body').append(aAlert);
        palert.open("#aAlert");
        mthis = $(t);
        $.get($(t).attr('href'), function(data, status, xhr){ 
            $(mthis.attr('plink')).html(data);
            init();
            palert.close("#aAlert");
            $("#aAlert").remove();
        });
        return true;
    }
}
//数字验证
function checkNum(obj) {  
    if(isNaN(obj.value)){
        obj.value=obj.value.replace(/[^\d.]/g,"")
    }
}
//调用ajax之后执行方法
function alink(t, f){
    if($(t).attr('link') != '' && $(t).attr('link') != undefined){
        aAlert = '<div class="palert" id="aAlert"><img src="/admin/image/loading.gif">加载中……</div>'
        $('body').append(aAlert);
        palert.open("#aAlert");
        mthis = $(t);
        $.get($(t).attr('link'), function(data, status, xhr){ 
            f(data)
            palert.close("#aAlert");
            $("#aAlert").remove();
        });
        return true;
    }
}
//输出所有方法
function displayProp(obj){    
    var names="";       
    for(var name in obj){       
       names+=name+": "+obj[name]+", ";  
    }  
    document.write(names)
}
//图片上传
UploadImg = {"param":{}}
UploadImg.isstart = 0
UploadImg.type = ["image/png", "image/jpg", "image/jpeg","image/gif"]
UploadImg.getData = function(mid){
    var data = document.getElementsByName(mid+'-'+'fupfile')
    var ret = [];
    for(var i=0;i<data.length;i++){
        var img = data[i].src;
        if(img.substr(0, 5)!="data:"){
            ret.push(data[i].src)
        }
    }
    return ret
}
UploadImg.addfile = function(t, files, mid, cat){
    for(var i=0;i<files.length;i++){
        var style, msg = "";
        var html = '<div class="col-sm-6 col-md-3 '+mid+'-filediv" style="margin-bottom:10px;"> \
          <div class="thumbnail"> \
             <img name="'+mid+'-'+'fupfile" style="height:100px;width:100px;" > \
             <span class="success" {0}>{1}</span>\
          </div> \
          <div class="caption"> \
            <a href="#" class="btn btn-default '+mid+'-'+'del" role="button" style="width: 100%;"> \
               删除 \
            </a> \
          </div> \
        </div>';
        if(UploadImg.type.indexOf(files[files.length-1-i].type)==-1){
            style = 'style="background: #f43838;display: block;position: absolute;right: 15px;bottom: 55px;height: 30px;width: 79%;text-align:center;line-height:30px;color:#fff"';
            msg = "图片格式错误";
        }else{
            style = "";
            msg = "";
        }
        html = printf(html, style, msg);
        cat.append(html);
        cat.find('.'+mid+'-'+'del').unbind().bind('click', function(){
            var img = $(this).parent().prev().find("img");
            var success = $(this).parent().prev().find(".success");
            var top = $(this).parent().parent();
            if(success.attr('style')!=undefined){
                htmlobj=$.ajax({url:UploadImg.param[mid+"url"]+"controller.php?action=uploaddel&url="+$(img).attr('src')+"&path="+UploadImg.param[mid+"path"],async:false});
            }
            top.remove();
        });
        if (files && files[i] ){
            var img = document.getElementsByName(""+mid+"-"+"fupfile")
            var stmp = function(limg, x){
                return function(evt){
                    // alert(limg[limg.length - 1].id)
                    if($(limg[limg.length - 1 - x]).next(".success").attr("style")==undefined)
                        limg[limg.length - 1 - x].src = evt.target.result;
                }
            }
            var reader = new FileReader();
            reader.onload = stmp(img, i);
            reader.readAsDataURL(files[i]);
        }
    }
    t.value = "";
    // alert(this.files.length)
}
UploadImg.clear = function(mid){
    $("."+mid+'-'+'filediv').remove();
}
UploadImg.addimg = function(mid, list){
    var idata = $("#"+mid).find("."+mid+"-"+"data");
    var button = $("#"+mid).find("."+mid+"-"+"fbutton");
    var file = $("#"+mid).find("."+mid+"-"+"ffile");
    var cat = $("#"+mid).find("."+mid+"-"+"fcatimg");
    var start = $("#"+mid).find("."+mid+"-"+"mstart");
    var html = '<div class="col-sm-6 col-md-3 '+mid+'-filediv" style="margin-bottom:10px;"> \
      <div class="thumbnail"> \
         <img name="'+mid+'-'+'fupfile" style="height:100px;width:100px;" src="{0}" > \
         <span class="success" style="background: url(/ueditor/dialogs/image/images/success.png) no-repeat right bottom;display: block;position: absolute;right: 15px;bottom: 55px;height: 40px;width: 40px;"></span>\
      </div> \
      <div class="caption"> \
        <a href="#" class="btn btn-default '+mid+'-'+'del" role="button" style="width: 100%;"> \
           删除 \
        </a> \
      </div> \
    </div>';
    for(i=0;i<list.length;i++){
        if(list[i] != ""){
            cat.append(printf(html, list[i]));
            cat.find('.'+mid+'-'+'del').unbind().bind('click', function(){
                var img = $(this).parent().prev().find("img");
                var success = $(this).parent().prev().find(".success");
                var top = $(this).parent().parent();
                if(success.attr('style')!=undefined){
                    
                    htmlobj=$.ajax({url:UploadImg.param[mid+"url"]+"controller.php?action=uploaddel&url="+$(img).attr('src')+"&path="+UploadImg.param[mid+"path"],async:false});
                }
                top.remove();
            });
        }
        
    }
}
UploadImg.create = function(mid, url, func, path){
    this.param[mid+"url"] = url;
    this.param[mid+"path"] = path;
    // var mid = '#bpimg';
    $("#"+mid).append('<div style="width: 100%;" class="'+mid+'-'+'fcatimg"></div>\
    <div>\
        <input type="hidden" class="'+mid+'-'+'data" value="" />\
        <div style="width: 100%;">\
            <div class="col-sm-6 col-md-3 ">\
              <div class="thumbnail">\
                 <img src="/ueditor/dialogs/image/images/image.png" style="height:100px;width:100px;" \
                 alt="上传图片">\
              </div>\
              <div class="caption">\
                <a href="#" class="btn btn-primary '+mid+'-'+'fbutton" role="button" style="width: 100%;">\
                   选择图片\
                </a>\
                <a href="#" class="btn '+mid+'-'+'mstart" role="button" style="width: 100%;">\
                   开始上传\
                </a>\
                <div id="ddfile">\
                    <input type="file" style="display: none;" class="'+mid+'-'+'ffile" multiple="true" accept="image/*" >\
                </div>\
              </div>\
            </div>\
        </div>\
    </div>')
    var idata = $("#"+mid).find("."+mid+"-"+"data");
    var button = $("#"+mid).find("."+mid+"-"+"fbutton");
    var file = $("#"+mid).find("."+mid+"-"+"ffile");
    var cat = $("#"+mid).find("."+mid+"-"+"fcatimg");
    var start = $("#"+mid).find("."+mid+"-"+"mstart");
    button.click(function(){
        file.click();
    });
    start.click(function(){
        if(start.attr('readonly')=='readonly'){
            return
        }
        var img = document.getElementsByName(""+mid+"-"+"fupfile")
        for(var i=0;i<img.length;i++){
            var success = $(img[i]).next(".success");
            if(success.attr('style')==undefined){
                start.attr('readonly', 'readonly');
                success.after('<div class="mprogress"><div style="position: absolute;left: 13px;right: 15px;bottom: 35px;height: 20px;width: calc(100%-17px);" class="progress progress-striped active"><div class="progress-bar progress-bar-success" role="progressbar" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100" style="width: 100%;"></div></div><div style="text-align:center"></div></div>')
                var run = function(mimg, msuccess){
                    UploadImg.isstart += 1;
                    $.ajax({
                        type: "POST",
                        url:url+"controller.php?action=uploadscrawl&path="+path,
                        async:true,
                        dataType: "json",
                        contentType: 'multipart/form-data',
                        data: mimg.src,
                        success: function (data) {
                            UploadImg.isstart -= 1;
                            msuccess.next(".mprogress").remove();
                            if(data['state']=='SUCCESS'){
                                mimg.src = data['url'];
                                msuccess.attr('style', "background: url(/ueditor/dialogs/image/images/success.png) no-repeat right bottom;display: block;position: absolute;right: 15px;bottom: 55px;height: 40px;width: 40px;");
                                if(idata.val()==""){
                                    sidata = [];
                                }else{
                                    sidata = idata.val().split(",");
                                }
                                sidata.push(data['url'])
                                idata.val(sidata.toString())
                            }else{
                                msuccess.attr('style', "background: #f43838;display: block;position: absolute;right: 15px;bottom: 55px;height: 30px;width: 79%;text-align:center;line-height:30px;color:#fff;");
                                if(data['state']=='-1'){
                                    msuccess.html('图片格式不对');
                                }else{
                                    msuccess.html(data['state']);
                                }
                            }
                            if(UploadImg.isstart == 0){
                                start.removeAttr('readonly');
                            }
                            func(data);
                        },
                        error: function(data){
                            UploadImg.isstart -= 1;
                            msuccess.next(".mprogress").remove();
                            start.removeAttr('readonly');
                            msuccess.attr('style', "background: #f43838;display: block;position: absolute;right: 15px;bottom: 55px;height: 30px;width: 79%;text-align:center;line-height:30px;color:#fff;");
                            if(data['state']=='-1'){
                                msuccess.html('图片格式错误');
                            }else{
                                msuccess.html(data['state']);
                            }
                            alert(data)
                            if(UploadImg.isstart == 0){
                                start.removeAttr('readonly');
                            }
                        }
                    });
                }
                run(img[i], success)
            }
        }
    });
    file.change(function(){
        UploadImg.addfile(this, this.files, mid, cat);
    });
    document.getElementById(mid).ondragover = function(event){
        event.preventDefault();
    }
    document.getElementById(mid).ondrop = function(event){
        event.preventDefault();

        UploadImg.addfile(file[0], event.dataTransfer.files, mid, cat);
    }
}
palert = {
    link : function(url, id){
        aAlert = '<div class="palert" id="aAlert"><img src="/admin/image/loading.gif">加载中……</div>'
        $('body').append(aAlert);
        palert.open("#aAlert");
        $.get(url, function(data, status, xhr){ 
            $(id).html(data);
            init();
            palert.close("#aAlert");
            $("#aAlert").remove();
        });
    },
    open : function(newalert){
        var bodyheight = document.body.scrollHeight;
        var bodywidth = document.body.scrollWidth;
        $(newalert).css('position', 'absolute');
        $(newalert).css('z-index', '999999');
        $(newalert).css('display', 'block');
        var palertheight = $(newalert).height();
        var palertwidth = $(newalert).width();
        var top = bodyheight/2 - $(newalert).height()/2;
        var left = bodywidth/2 - $(newalert).width()/2;
        $(newalert).css('top', "200px");
        $(newalert).css('left', left + 'px');
        html = "<div class='nalert' style='position: absolute;top: 0px;left: 0px;height: "+bodyheight+"px;width: "+bodywidth+"px'></div>";
        $('body').append(html);
    },
    close: function(newalert){
        $(newalert).css('display', 'none');
        $(".nalert").remove();
    },
}

//图片上传预览    IE是用了滤镜。
function previewImage(file, mid, w, h)
{
  var MAXWIDTH  = w; 
  var MAXHEIGHT = h;
  var div = document.getElementById(mid);
  if (file.files && file.files[0])
  {
      div.innerHTML ='<img id="imghead_'+mid+'">';
      var img = document.getElementById('imghead_'+mid);
      img.onload = function(){
        var rect = clacImgZoomParam(MAXWIDTH, MAXHEIGHT, img.offsetWidth, img.offsetHeight);
        // img.style.width = '100%'
        if(w == 0){
            img.style.width = '100%';
        }else{
            img.style.width  =  w+'px';
        }
        if(h != 0){
            img.style.height =  h+'px';
        }
//                 img.style.marginLeft = rect.left+'px';
        // img.style.marginTop = rect.top+'px';
      }
      var reader = new FileReader();
      reader.onload = function(evt){img.src = evt.target.result;}
      reader.readAsDataURL(file.files[0]);
  }
  else //兼容IE
  {
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

//判断图片大小是否正确
function getimageWH(file, MAXWIDTH, MAXHEIGHT, fn)
{
  if (file.files && file.files[0])
  {
      var img = new Image();
      img.onload = function(){
        if(this.width != MAXWIDTH || this.height != MAXHEIGHT){
            fn();
            file.value = "";
        }
      }
      var reader = new FileReader();
      reader.onload = function(evt){img.src = evt.target.result;}
      reader.readAsDataURL(file.files[0]);
  }
}

//桌面提示
function notify(title, content, imgurl) {
    if(!title && !content){
        title = "桌面提醒";
        content = "您看到此条信息桌面提醒设置成功";
    }
    var iconUrl = "/images/send_ok.png";
    if(imgurl){
        iconUrl = imgurl;
    }
    if (window.webkitNotifications) {
        //chrome老版本
        if (window.webkitNotifications.checkPermission() == 0) {
            var notif = window.webkitNotifications.createNotification(iconUrl, title, content);
            notif.display = function() {}
            notif.onerror = function() {}
            notif.onclose = function() {}
            notif.onclick = function() {this.cancel();}
            notif.replaceId = 'Meteoric';
            notif.show();
        } else {
            window.webkitNotifications.requestPermission($jy.notify);
        }
    }
    else if("Notification" in window){
        // 判断是否有权限
        if (Notification.permission === "granted") {
            var notification = new Notification(title, {
                "icon": iconUrl,
                "body": content,
            });
        }
        //如果没权限，则请求权限
        else if (Notification.permission !== 'denied') {
            Notification.requestPermission(function(permission) {
                // Whatever the user answers, we make sure we store the
                // information
                if (!('permission' in Notification)) {
                    Notification.permission = permission;
                }
                //如果接受请求
                if (permission === "granted") {
                    var notification = new Notification(title, {
                        "icon": iconUrl,
                        "body": content,
                    });
                }
            });
        }
    }
} 

//滚到顶部
$Scroll = {};
//滚动条在Y轴上的滚动距离
$Scroll.getScrollTop = function(){
　　var scrollTop = 0, bodyScrollTop = 0, documentScrollTop = 0;
　　if(document.body){
　　　　bodyScrollTop = document.body.scrollTop;
　　}
　　if(document.documentElement){
　　　　documentScrollTop = document.documentElement.scrollTop;
　　}
　　scrollTop = (bodyScrollTop - documentScrollTop > 0) ? bodyScrollTop : documentScrollTop;
　　return scrollTop;
}
//文档的总高度
$Scroll.getScrollHeight = function(){
　　var scrollHeight = 0, bodyScrollHeight = 0, documentScrollHeight = 0;
　　if(document.body){
　　　　bodyScrollHeight = document.body.scrollHeight;
　　}
　　if(document.documentElement){
　　　　documentScrollHeight = document.documentElement.scrollHeight;
　　}
　　scrollHeight = (bodyScrollHeight - documentScrollHeight > 0) ? bodyScrollHeight : documentScrollHeight;
　　return scrollHeight;
}
//浏览器视口的高度
$Scroll.getWindowHeight = function(){
　　var windowHeight = 0;
　　if(document.compatMode == "CSS1Compat"){
　　　　windowHeight = document.documentElement.clientHeight;
　　}else{
　　　　windowHeight = document.body.clientHeight;
　　}
　　return windowHeight;
}
$Scroll._timer = "";
$Scroll.top = function(target,time){
    var start = document.documentElement.scrollTop || document.body.scrollTop;
    var dis = target - start;
    var count = Math.floor(time/10);
    var n=0;
    
    clearInterval($Scroll._timer);
    $Scroll._timer = setInterval(function(){
        n++;
        bFlag=false;
        document.documentElement.scrollTop = start + dis*n/count;
        document.body.scrollTop = start+dis*n/count;
        if(n==count){
            clearInterval($Scroll._timer);    
        }

    }, 10);
}
$Scroll.Create = function(id, h){
    if(h == undefined){h = 30};
    document.getElementById(id).style.opacity = '0';
    document.getElementById(id).onclick = function(){
        $Scroll.top(0, 500);
    };
    //document.all.push('12')
    //document.all[0].innerHTML += '<style>@-webkit-keyframes fade {from { opacity: 0; } to { opacity: 1; }}@keyframes fade {from {opacity: 0; -webkit-opacity:0; } to {opacity: 1;-webkit-opacity: 1; }}</style>';
    window.onscroll = function(){
    　　if($Scroll.getScrollTop() + $Scroll.getWindowHeight() >= $Scroll.getScrollHeight()-h){
            //document.getElementById(id).style.cssText += '-webkit-animation: fade 2s ease both;animation: fade 2s ease both;';
            document.getElementById(id).style.opacity = '1';
        }
    };
}

function init(){
    $('a').unbind("click").click(function(){
        // alert($(this).attr('plink'))
        if($(this).attr('plink') != '' && $(this).attr('plink') != undefined){
            aAlert = '<div class="palert" id="aAlert"><img src="/admin/image/loading.gif">加载中……</div>'
            $($(this).attr('iframe')).append(aAlert);
            palert.open("#aAlert");
            mthis = $(this);
            $.get($(this).attr('href'), function(data, status, xhr){ 
                $(mthis.attr('plink')).html(data);
                init();
                palert.close("#aAlert");
                $("#aAlert").remove();
            });
            return false;
        }else if ($(this).attr('iframe') != '' && $(this).attr('iframe') != undefined){
            var height = $('#navbar').height();
            var iframe = '<iframe style="width:100%;height:{1}px;" src="{0}"></iframe>'
            $($(this).attr('iframe')).html(printf(iframe, $(this).attr('href'), height));
            return false;
        }
        return true;
    });
}
init();