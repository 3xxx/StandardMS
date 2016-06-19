 <!DOCTYPE html>
{{template "header"}}
<title>成果 - 水利设计CMS系统</title>
<style>
i#delete
{
color:#DC143C;
}
</style>
</head>

<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>

<div class="col-lg-12">
  <h1>成果列表</h1>
  <table class="table table-striped">
    <thead>
      <tr>
        <th>
          <span style="cursor: pointer">#{{.Length}}</span>
        </th>
        <th>
          <span style="cursor: pointer">成果编号</span>
        </th>
        <th>
          <span style="cursor: pointer">成果名称</span>
        </th>
        <th>
          <span style="cursor: pointer">成果类型</span>
        </th>
        <th>
          <span style="cursor: pointer">作者</span>
        </th>
        <th>
          <span style="cursor: pointer">最后更新</span>
        </th>
<!--         <th>
          <span style="cursor: pointer">浏览</span>
        </th> -->
<!--         <th>
          <span style="cursor: pointer">评论数</span>
        </th> -->
        <th>
          <span style="cursor: pointer">最后评论</span>
        </th>
        <th>操作</th>
      </tr>
    </thead>
    <tbody>
      <!-- <ol>
      -->
   {{range $index, $elem :=.Topics}}
      <tr>
        <!--tr表格的行，td定义一个单元格，<th>
        标签定义表格内的表头单元格-->
        <th>{{$index}}</th>
        <th>
          <a href="/topic/view_b/{{.Id}}" id="number">{{substr .Tnumber 0 15}}</a>
        </th>
        <th>
          <a href="/topic/view_b/{{.Id}}" id="name" title={{.Title}}> <i class="glyphicon glyphicon-fire"></i>
            {{substr .Title 0 15}}
          </a>
        </th>
        <th>{{.Category}}</th>
        <th>{{.Author}}</th>
        <!-- {{.Attachment}} -->
        <th>{{dateformat .Updated "2006-01-02"}}</th>
        <!-- <th>{{.Views}}</th> -->
        <!-- <th>{{.ReplyCount}}</th> -->
        <th>{{dateformat .ReplyTime "2006-01-02"}}</th>
        <th>
          <a href="/topic/view_b/{{.Id}}"><i class="glyphicon glyphicon-download-alt"></i>下载</a>
          <a href="/topic/modify?tid={{.Id}}"><i class="glyphicon glyphicon-edit"></i>修改</a>
          <!-- <a href="/topic/delete?tid={{.Id}}"><i id="delete" class="glyphicon glyphicon-remove-sign"></i>删除</a> -->
          <a href="" id="{{.Id}}" onclick="deletetopic('{{.Id}}')"><i id="delete" class="glyphicon glyphicon-remove-sign"></i>删除</a>
        </th>
      </tr>
      {{end}}
      <!-- </ol>--></tbody>
  </table>

  <!-- float: right;调整位置 -->
  <div style="text-align:center;padding-left: 100px;margin-top: -24px;float: right;" class="pagination">
    {{if .paginator}}
        {{if gt .paginator.PageNums 1}}
    <ul class="pagination pagination-sm">
      {{if .paginator.HasPrev}}
      <li>
        <a href="{{.paginator.PageLinkFirst}}">首页</a>
      </li>
      <li>
        <a href="{{.paginator.PageLinkPrev}}">上一页</a>
      </li>
      {{else}}
      <li class="disabled">
        <a>首页</a>
      </li>
      <li class="disabled">
        <a>上一页</a>
      </li>
      {{end}}
            {{range $index, $page := .paginator.Pages}}
      <li{{if $.paginator.IsActive .}} class="active"{{end}}>
        <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
      </li>
      {{end}}
            {{if .paginator.HasNext}}
      <li>
        <a href="{{.paginator.PageLinkNext}}">下一页</a>
      </li>
      <li>
        <a href="{{.paginator.PageLinkLast}}">末页</a>
      </li>
      {{else}}
      <li class="disabled">
        <a>下一页</a>
      </li>
      <li class="disabled">
        <a>末页</a>
      </li>
      {{end}}
      <li class="disabled">
        <a>
          共{{.paginator.Nums }}条数据 每页{{.paginator.PerPageNums}}条 当前{{.paginator.Page}}/{{.paginator.PageNums}}页
        </a>
      </li>
      <li>
        <input type="text" type="submit" id="p" name="p" placeholder="跳转页" style="width: 47px;height: 30px;border: 1px solid #dddddd;border-left: 0px;border-radius: 0px 4px 4px 0px;text-align: center;"/>
      </li>
    </ul>
    {{end}} 
  {{end}}
  </div>

  <!--  <input type="hidden" id="p" name="p" value="2" />
  <button type="submit" class="btn btn-default" >第2页</button>
  -->
  <!-- </form>--></div>

<script type="text/javascript">
  $(document).ready(function() {
  $("table").tablesorter();
  // $("#ajax-append").click(function() {
  //    $.get("assets/ajax-content.html", function(html) {
  //     // append the "ajax'd" data to the table body
  //     $("table tbody").append(html);
  //     // let the plugin know that we made a update
  //     $("table").trigger("update");
  //     // set sorting column and direction, this will sort on the first and third column
  //     var sorting = [[2,1],[0,0]];
  //     // sort on the first column
  //     $("table").trigger("sorton",[sorting]);
  //   });
  //   return false;
  // });
});

  function deletetopic(id) {
    if(confirm("确定删除吗？")){
 $.ajax({
                type:"post",
                url:"/topic/delete",
                data: {tid:id,url:window.location.href},//父级id
                success:function(data,status){
                  alert("删除“"+data+"”成功！(status:"+status+".)");
                  // window.location=window.location
                 }
            });
 // window.location.reload();这句可有可无？
// window.location.href='findAllFoods.action';

}else{
return false;
}
}
// alert(window.location.href);

//保持页面滚动条位置
window.onbeforeunload = function () { 
var scrollPos; 
if (typeof window.pageYOffset != 'undefined') { 
scrollPos = window.pageYOffset; 
} 
else if (typeof document.compatMode != 'undefined' && 
document.compatMode != 'BackCompat') { 
scrollPos = document.documentElement.scrollTop; 
} 
else if (typeof document.body != 'undefined') { 
scrollPos = document.body.scrollTop; 
} 
document.cookie = "scrollTop=" + scrollPos; //存储滚动条位置到cookies中 
} 

window.onload = function () { 
if (document.cookie.match(/scrollTop=([^;]+)(;|$)/) != null) { 
var arr = document.cookie.match(/scrollTop=([^;]+)(;|$)/); //cookies中不为空，则读取滚动条位置 
document.documentElement.scrollTop = parseInt(arr[1]); 
document.body.scrollTop = parseInt(arr[1]); 
} 
}

</script>

</body>
</html> 