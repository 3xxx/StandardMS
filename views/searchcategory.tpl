<!DOCTYPE html>
{{template "header"}}
<title>检索 - 水利设计CMS系统</title>
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
  <h1>检索项目列表</h1>
<table class="table table-striped">
  <thead>
    <tr>
      <th><span style="cursor: pointer">#{{.Length}}</span></th>
      <th><span style="cursor: pointer">项目编号</span></th>
      <th><span style="cursor: pointer">项目名称</span></th>
      <th><span style="cursor: pointer">项目类型</span></th>
      <th><span style="cursor: pointer">最后更新</span></th>
      <th><span style="cursor: pointer">浏览</span></th>
      <th><span style="cursor: pointer">成果数</span></th>
      <th><span style="cursor: pointer">作者</span></th>
      <th>操作</th>
    </tr>
  </thead>

  <tbody>
  <!-- <ol> -->
   {{range $index, $elem :=.Searchs}}
    <tr><!--tr表格的行，td定义一个单元格，<th> 标签定义表格内的表头单元格-->
      <th>{{$index}}</th>
      <th><a href="/category?op=view&id={{.Id}}">{{substr .Number 0 15}}</a></th>
      <th><a href="/category?op=view&id={{.Id}}" title={{.Title}}><i class="glyphicon glyphicon-fire"></i>{{substr .Title 0 15}}</a></th>
      <th>
          {{range $k1,$v1 :=$.Label}}
          {{if eq $elem.Id $v1.Category.Id}}
          <a href="/category?op=viewlabel&label={{.Title}}" ><span class="label label-info">{{.Title}}</span></a>
          {{end}}
          {{end}}
      </th>

      <th>{{dateformat .Updated "2006-01-02"}}</th>
      <th>{{.Views}}</th>
      <th>{{.TopicCount}}</th>
      <th>{{.Author}}</th>
      <th>
      <a href="/category/modify?cid={{.Id}}"><i class="glyphicon glyphicon-edit"></i>修改</a>
      <a href="" id="{{.Id}}" onclick="deletecategory('{{.Id}}')"><i id="delete" class="glyphicon glyphicon-remove-sign"></i>删除</a></th>
    </tr>
    {{end}}

  </tbody>
 </table>

  </div>
<script type="text/javascript">
  $(document).ready(function() {
  $("table").tablesorter();
  $("#ajax-append").click(function() {
     $.get("assets/ajax-content.html", function(html) {
      // append the "ajax'd" data to the table body
      $("table tbody").append(html);
      // let the plugin know that we made a update
      $("table").trigger("update");
      // set sorting column and direction, this will sort on the first and third column
      var sorting = [[2,1],[0,0]];
      // sort on the first column
      $("table").trigger("sorton",[sorting]);
    });
    return false;
  });
});

 function deletecategory(id) {
    if(confirm("确定删除吗？")){
 $.ajax({
                type:"post",
                url:"/category/delete",
                data: {cid:id,url:window.location.href},//父级id
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

</script>

</body>
</html>
