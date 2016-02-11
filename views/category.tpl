<!DOCTYPE html>
{{template "header"}}
<title>项目&目录 - 水利设计CMS系统</title>
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
  <h1>项目列表</h1>
    <a href="/category/add" class="btn btn-default">添加项目&目录</a>
    <a href="/category/add_b" class="btn btn-default">添加自定义目录</a>
    <a href="/category/add_b" class="btn btn-default">添加标准成果</a>
    <button class="btn btn-default" onclick="parent.location.href='/topic/add?mid=7'">快捷上传</button>
      <table class="table table-striped">
       <thead>
         <tr>
           <th style="cursor: pointer">#{{.Length}}</th>
           <th style="cursor: pointer">项目编号</th>
           <th style="cursor: pointer">项目名称</th>
           <th style="cursor: pointer">项目类型</th>
           <th style="cursor: pointer">负责人</th>
           <th style="cursor: pointer">成果数量</th>
           <th style="cursor: pointer">建立时间</th>
           <!-- <th style="cursor: pointer">修改时间</th> -->
           <th>操作</th>
         </tr>
       </thead>

       <tbody>
   
         {{range $index, $elem :=.Category}}
         <tr>
          <th>{{$index}}</th>
          <th><a href="/category?op=view&id={{.Id}}" id="number">{{.Number}}</a></th>
         <th><a href="/category?op=view&id={{.Id}}" id="name"><i class="glyphicon glyphicon-plane"></i>{{.Title}}</a></th>
         <th> </th>
         <th>{{.Author}}</th>
         <th>{{.TopicCount}}</th>
         <th>{{dateformat .Created "2006-01-02 "}}</th>
         <!-- <th>{{dateformat .Updated "2006-01-02 "}}</th>          -->
       <th>
         <!-- <a href="/category?op=view&id={{.Id}}">显示</a> -->
         <a href="/category/modify?cid={{.Id}}"><i class="glyphicon glyphicon-edit"></i>修改</a>
         <a href="/category?op=del&id={{.Id}}"><i id="delete" class="glyphicon glyphicon-remove-sign"></i>删除</a>
       </th>
     </tr>
     {{end}}
   
     </tbody>
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

     
  </div>

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
</script>
</body>
</html>