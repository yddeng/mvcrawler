window.onload = function(){
   getUpdate()
};

function getUpdate() {
    let list = document.getElementById('list');
    list.innerHTML = `<div id="tips"><i class="fa fa-spinner" aria-hidden="true"></i>加载中，请稍候！</div>`;
    $.ajax({
        url:BaseUrl + "/update",
        type: "get",
        async: true,
        dataType: "json",
        //data: {"id":12},
        success: function (res) {
            if(res.code == "OK"){
                showUpdateItems(res.items)
            }else{
                noData()
            }
        },
        error: function (e){
            httpErr();
        }
    });
}

function showUpdateItems(items) {
let tmp = `
    <a href="{0}" target="con">
	    <div class="box">
			<div class="status">{1}</div>
			<div class="box-img"><img src="{2}"  onerror="this.src='./template/images/default_grey_pc.png'"></div>
			<div class="detail">
				<div class="title">{3}</div>
				<div class="from">来源:<text class="from-detail">{4}</text></div>
			</div>
		</div>
	</a>
`;
    let list = document.getElementById('list');
    list.innerHTML = "";
    for(var i=0;i<items.length;i++) {
        var row = items[i];
        for (var j = 0; j < row.length; j++) {
            let oLi = document.createElement("li");
            oLi.innerHTML = String.format(tmp,row[j].url,row[j].status,row[j].img,row[j].title,row[j].from);
            list.appendChild(oLi)
        }
    }
}


/*------------------------------------------------------------------------------------------------------*/

function search(){
    let txt = $("#keyboard").val();
    if (txt != "") {
        getSearch(txt, 0);
    }
}

function keySearch(){
    if (event.keyCode == 13){
        search()
    }
}

function getSearch(txt,page) {
    let list = document.getElementById('list');
    list.innerHTML = `<div id="tips"><i class="fa fa-spinner" aria-hidden="true"></i>搜索中，请稍候！</div>`;
    $.ajax({
        url:BaseUrl + "/search",
        type: "get",
        async: true,
        dataType: "json",
        data: {"txt":txt,"page":page},
        success: function (res) {
            if(res.code == "OK"){
                let text = String.format(`搜索"{0}" 共找到"{1}"个相关资源`,res.txt,res.total_item)
                $('#title-top').text(text);
                showSearchItems(res.items);
                showPage(res.txt,res.page,res.total_page)
            }else{
                noData()
            }
        },
        error: function (e){
            httpErr()
        }
    });
}

function showSearchItems(items) {
    let tmp = `
    <a href="{0}" target="con">
	    <div class="box">
			<div class="status">{1}</div>
			<div class="box-img"><img src="{2}"  onerror="this.src='./template/images/default_grey_pc.png'"></div>
			<div class="detail">
				<div class="title">{3}</div>
				<div class="from">来源:<text class="from-detail">{4}</text></div>
			</div>
		</div>
	</a>
`;
    let list = document.getElementById('list');
    list.innerHTML = "";
    for(var i=0;i<items.length;i++) {
        var row = items[i];
        let oLi = document.createElement("li");
        oLi.innerHTML = String.format(tmp,row.url,row.status,row.img,row.title,row.from);
        list.appendChild(oLi)
    }
}


// page
function showPage(txt,page,total) {
    let box = document.getElementById('page-content');
    box.innerHTML = "";
    let str = ``;

    //显示5个页码，其中一个当前页码
    let cnt = 4;

    // start
    if(page == 0) {
        str += `<a>首页</a>`;
        str += `<a>上一页</a>`;
        str += `<a id="current">第1页</a>`;
    } else {
        str += `<a onclick="getSearch('${txt}',0)">首页</a>`;
        str += `<a onclick="getSearch('${txt}',${page-1})">上一页</a>`;

        let start = page - 2
        if (start < 0 ){
            start = 0
        }
        for (let i = start; i < page; i++){
            str += `<a onclick="getSearch('${txt}',${i})">第${i+1}页</a>`;
            cnt--;
        }
        str += `<a id="current">第${page+1}页</a>`;
    }

    //cnt大于0说明右边还可放置cnt个页码
    for(let i = page + 1; cnt > 0 && i < total; i++) {
        str += `<a onclick="getSearch('${txt}',${i})">第${i+1}页</a>`;
        cnt--;
    }

    // end
    if (page == total-1){
        str += `<a>下一页</a>`;
        str += `<a>尾页</a>`;
    }else {
        str += `<a onclick="getSearch('${txt}',${page+1})">下一页</a>`;
        str += `<a onclick="getSearch('${txt}',${total-1})">尾页</a>`;
    }

    box.innerHTML = str;
}


