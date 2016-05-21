<!DOCTYPE html>
<!-- 显示指定项目中的所有成果 -->
<head>
<!-- 对jquery的引用必须放在head中-->
<!-- <meta http-equiv="Content-Type" content="text/html;charset=utf-8"> -->
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.js"></script>
  <link type="text/css" href="/static/css/bootstrap.min.css" rel="stylesheet" />
<style>
i#delete
{
color:#DC143C;
}
</style>
</head>
<body>
<!-- <div class="navbar navba-default navbar-fixed-top">
</div> -->

<div class="container-fill">

<h3>成果列表</h3>

<table class="table table-striped">
  <thead>
    <tr>
      <th style="cursor: pointer">#{{.Length}}</th>
      <th style="cursor: pointer">成果编号</th>
      <th style="cursor: pointer">成果名称</th>
      <th style="cursor: pointer">成果类型</th>
      <th style="cursor: pointer">作者</th>
      <th style="cursor: pointer">最后更新</th>
      <!-- <th style="cursor: pointer">浏览</th> -->
      <!-- <th style="cursor: pointer">最后回复</th> -->
      <th>操作</th>
    </tr>
  </thead>
  <tbody>
    {{range $index, $elem :=.Chengguo}}
    <tr>
      <th>{{$index}}</th>
      <th><a href="/topic/view_b/{{.Id}}" title={{.Title}} target="_blank">{{substr .Tnumber  0 15}}</a></th>
      <th><a href="/topic/view_b/{{.Id}}" title={{.Title}} target="_blank">{{substr .Title 0 15}}</a></th>
      <th>{{.Category}}</th>
      <th>{{.Author}}</th>
      <th>{{dateformat .Updated "2006-01-02"}}</th>
      <!-- <th>{{.Views}}</th> -->
      <!-- <th>{{.ReplyCount}}</th> -->
<!--       <th><a href="/topic/view/{{.Id}}">下载</a>
      <a href="/topic/modify?tid={{.Id}}">修改</a>
      <a href="/topic/delete?tid={{.Id}}">删除</a></th> -->
        <th>
          <a href="/topic/view_b/{{.Id}}" title={{.Title}} target="_blank"><i class="glyphicon glyphicon-download-alt"></i>下载</a>
          <a href="/topic/modify?tid={{.Id}}"><i class="glyphicon glyphicon-edit" target="_blank"></i>修改</a>
          <a href="/topic/delete?tid={{.Id}}"><i id="delete" class="glyphicon glyphicon-remove-sign"></i>删除</a>
        </th>
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
  //     // set sorting column列号 and direction排序方式, this will sort on the first第一列降序 and third column第三列升序
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
<!--     $(document).ready(function(){
        $("#mytable").tablesorter({ //mytable是报表table的id
            1: {sorter:"integer"}, //按照integer数据类型进行排序
            2: {sorter:"integer"}, 
            3: {sorter:"double"}, 
            sortList: [[1,1],[2,0],[3,1]] //[列号，排序标志]初始化表格时第二列按照降序排序，第三列按照升序排列，第四列按照降序排列
        });
     } ); -->