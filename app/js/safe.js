
window.oncontextmenu = function(e){
    e.preventDefault(); //阻止浏览器自带的右键菜单显示
    var menu = document.getElementById("right-menu");
    menu.style.display = "block"; //将自定义的“右键菜单”显示出来
    menu.style.left = e.clientX + "px";  //设置位置，跟随鼠标
    menu.style.top = e.clientY+"px";
}

window.onclick = function(e){ //点击窗口，右键菜单隐藏
    var menu = document.getElementById("right-menu");
    menu.style.display = "none";
}