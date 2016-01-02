<!DOCTYPE html>
<head>
<!-- 对jquery的引用必须放在head中-->
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.js"></script>
  <link type="text/css" href="/static/css/bootstrap.min.css" rel="stylesheet" />
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
<!-- <div class="navbar navba-default navbar-fixed-top">
</div> -->

<div class="container-fill">
  <!-- 这里添加项目目录树显示 -->
  <h3> <font size="4">{{.CategoryPhase.Title}} » </font>
    <font size="4">{{.CategorySpec.Title}} » </font>
    <font size="5">{{.Category.Title}}</font>
  </h3>

  <h2>日记列表
      <!-- <a href="/topic/add?id={{.Category.Id}}&mid=3" class="btn btn-primary">写设代日记</a> onclick="parent.location.href='/topic/add?id={{.Category.Id}}&mid=3'" -->
          <button id="directNextpage" class="btn btn-default" onclick="parent.location.href='/topic/add?id={{.Category.Id}}&mid=3'">写设代日记</button>
</h2>
<script>
$(document).ready(function(){
$("#export").click(function(){//这里应该用button的id来区分按钮的哪一个,因为本页有好几个button
  // $(function(){
            $.ajax({
                type:"POST",
                url:"/topic/exporttoexcel",
                // data:$('#form').serialize(),
                //格式化表单参数或者也可以使用data:{'folder':$('input[name=folder]').val(),'page':$('input[name=page]').val()},
                // data:{'aid':54,'content':55,'aid':56,'content':57},
                // data:o,
                // datetype:'text',
                data: { id: "{{.Category.Id}}",tel:$("#tnumber").val()},
                // success: function(responseText) {
                // $("#divResult").html(responseText);
                // }
                success:function(data){//数据提交成功时返回数据
                    // alert(data);
                    alert("导出成功！")
                }
            });
            return false;//这里应该是不刷新页面的意思
 });
});
  </script>



<table class="table table-striped">
  <thead>
    <tr>
      <th style="cursor: pointer">#{{.Length}}</th>
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
      <th><a href="/topic/view/{{.Id}}" target="_blank">{{.Tnumber}}</a></th>
      <th><a href="/topic/view/{{.Id}}" title={{.Title}}>{{substr .Title 0 15}}</a></th>
      <th>{{dateformat .Updated "2006-01-02"}}</th>
      <th>{{.Views}}</th>
      <th>{{.ReplyCount}}</th>
      <th><a href="/topic/view/{{.Id}}">下载</a>
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