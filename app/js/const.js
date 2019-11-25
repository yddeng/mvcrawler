//const BaseUrl = "http://127.0.0.1:12345"
const BaseUrl = "http://104.168.244.58:12345"

// http请求出错
function httpErr() {
    let list = document.getElementById('item-list')
    list.innerHTML = `<div id="tips"><i class="fa fa-exclamation-circle"></i>网络开小差了,请稍后再试!</div>`
}

// code
function codeErr() {
    let list = document.getElementById('item-list')
    list.innerHTML = `<div id="tips"><i class="fa fa-circle-o-notch"></i>code err!</div>`
}

// 没有资源
function noData() {
    let list = document.getElementById('item-list')
    list.innerHTML = `<div id="tips"><i class="fa fa-circle-o-notch"></i>暂时没有资源呢!</div>`
}


// 字符串格式化
String.format = function(src){
    if (arguments.length == 0) return null;
    var args = Array.prototype.slice.call(arguments, 1);
    return src.replace(/\{(\d+)\}/g, function(m, i){
        return args[i];
    });
};
