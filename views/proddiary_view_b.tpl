<!DOCTYPE html>
{{template "header"}}
<title>成果分类查看 - 水利设计CMS系统</title>
<style type="text/css">
h3 {line-height: 150%;
  /*color: #DC143C;#000000*/
  color:#FFFFFF;
  background: #4682B4;/*#fff*/
  margin: 10;
  padding: 10;
  font-family: Georgia, Palatino, serif;
  }
</style>
</head>

<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>

<div class="col-lg-12">
  <!-- 这里添加项目目录树显示 color="#A52A2A" -->
  <h3> <font size="5">{{.CategoryProj.Title}} » </font> <font size="4">{{.CategoryPhase.Title}} » </font>
    <font size="4">{{.CategorySpec.Title}} » </font>
    <font size="5">{{.Category.Title}}</font>
  </h3>

  <h2>日记列表
      <a href="/topic/add?id={{.Category.Id}}&mid=3" class="btn btn-primary">写设代日记</a> </h2>
<table class="table table-striped">
  <thead>
    <tr>
      <th style="cursor: pointer">#</th>
      <th style="cursor: pointer">日记编号</th>
      <th style="cursor: pointer">日记名称</th>
      <th style="cursor: pointer">最后更新</th>
      <th style="cursor: pointer">浏览</th>
      <th style="cursor: pointer">回复数</th>
      <th>操作</th>
    </tr>
  </thead>
  <tbody>
    {{range $index,$elem:=.Chengguo}}
    <tr>
      <th>{{$index}}</th>
      <th><a href="/topic/view_b/{{.Id}}">{{.Tnumber}}</a></th>
      <th><a href="/topic/view_b/{{.Id}}" title={{.Title}}>{{substr .Title 0 15}}</a></th>
      <th>{{dateformat .Updated "2006-01-02 T 15:04:05"}}</th>
      <th>{{.Views}}</th>
      <th>{{.ReplyCount}}</th>
      <th><a href="/topic/view_b/{{.Id}}">下载</a>
      <a href="/topic/modify?tid={{.Id}}">修改</a>
      <a href="/topic/delete?tid={{.Id}}">删除</a></th>
      <!--<th>
        <a href="/topic?op=del&id={{.Id}}">删除</a>
      </th>-->
    </tr>
    {{end}}
  </tbody>
  </table>
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