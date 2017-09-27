websocket = {}
websocket.sleep = 5000;
websocket.sock = null;
websocket.url = null;
websocket.open = function(url){
    this.url = url;
    try{
        console.log("websocket onload");
        this.sock = new WebSocket(this.url);
        this.sock.onopen = this.onopen;
        this.sock.onmessage = this.onmessage;
        this.sock.onclose = this.onclose;
    }catch (e){
        setTimeout("websocket.open('"+this.url+"')", websocket.sleep);
    }
}
websocket.send = function(msg){
    console.log("send: "+msg);
    this.sock.send(msg);
};
websocket.onmessage = function(e){
    console.log("message received:" + e.data);
};
websocket.onclose = function(e) {
    console.log("connection closed (" + e.code + ")");
    setTimeout("websocket.open('"+this.url+"')", websocket.sleep);
};
websocket.onopen = function(){
    console.log("connected to " + this.url);
};