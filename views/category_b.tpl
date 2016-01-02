<!DOCTYPE html>
{{template "header"}}
<title>项目&目录 - 水利设计CMS系统</title>
</head>

<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>

<div class="col-lg-12">
  <h1>项目列表</h1>
    <a href="/category/add" class="btn btn-default">添加项目&目录</a>
      <table class="table table-striped">
       <thead>
         <tr>
           <th style="cursor: pointer">#{{.Length}}</th>
           <th style="cursor: pointer">项目编号</th>
           <th style="cursor: pointer">项目名称</th>
           <th style="cursor: pointer">建立时间</th>
           <th style="cursor: pointer">修改时间</th>
           <th style="cursor: pointer">浏览</th>
           <th style="cursor: pointer">成果数量</th>
           <th>操作</th>
         </tr>
       </thead>

       <tbody>
         {{range $index, $elem :=.Category}}
         <tr>
          <th>{{$index}}</th>
         <th>
           {{.Number}}
         </th>
         <th><a href="/category?op=view_b&id={{.Id}}"><i class="glyphicon glyphicon-plane"></i>{{.Title}}</a></th>
         
         <th>{{dateformat .Created "2006-01-02 "}}</th>
         <th>{{dateformat .Updated "2006-01-02"}}</th>
         <th>{{.Views}}</th>
         <th>{{.TopicCount}}</th>
       <th>
         <a href="/category?op=view_b&id={{.Id}}">显示</a>
         <a href="/category?op=modify&id={{.Id}}">修改</a>
         <a href="/category?op=del&id={{.Id}}">删除</a>
       </th>
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
 <!-- .Created "2006-01-02 T 15:04:05" -->
</body>
</html>