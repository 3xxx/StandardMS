<!DOCTYPE html>
{{template "header"}}
<title>{{.Topic.Title}} - 水利设计CMS系统</title>
</head>


<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>
<div class="col-lg-12">
<table id="datagrid" toolbar="#tb"></table>

<form method="post" action="/user/importexcel" enctype="multipart/form-data"> 
   <div class="input-group">
      <label>选择excel：<input type="file" name="excel" id="excel" /></label>
      <br/>
   </div>
   <button type="submit" class="btn btn-default" >提交</button>
</form>

<table class="table table-striped">
  <thead>
      <tr>
      <th>#{{.Users|len}}</th>
      <th><span style="cursor: pointer">Id</span></th>
      <th><span style="cursor: pointer">Username</span></th>
      <th><span style="cursor: pointer">Password</span></th>
      <th><span style="cursor: pointer">Nickname</span></th>
      <th><span style="cursor: pointer">Email</span></th>
        <th><span style="cursor: pointer">分院</span></th>
        <th><span style="cursor: pointer">科室</span></th>
      <th><span style="cursor: pointer">Remark</span></th>
      <th><span style="cursor: pointer">Status</span></th>
      <th><span style="cursor: pointer">Lastlogintime</span></th>
      <th><span style="cursor: pointer">Createtime</span></th>
      <th><span style="cursor: pointer">Role</span></th>
      <th>操作</th>
    </tr>
  </thead>
   {{range $index, $elem :=.Users}}
    <tr><!--tr表格的行，td定义一个单元格，<th> 标签定义表格内的表头单元格-->
      <th>{{$index}}</th>
      <th><a href="/user/view?useid={{.Id}}">{{.Id}}</a></th>
      <th><a href="/user/view?useid={{.Id}}"><i class="glyphicon glyphicon-user"></i>{{.Username}}</a></th>
      <th><a href="/user/view?useid={{.Id}}" title={{.Title}}>{{substr .Password 0 15}}</a></th>
      <th>{{.Nickname}}</th>
      <th>{{.Email}}</th>
      <th>{{.Department}}</th>
      <th>{{.Secoffice}}</th>
      <th>{{.Remark}}</th>
      <th>{{.Status}}</th>
      <th>{{dateformat .Lastlogintime "2006-01-02"}}</th>
      <th>{{dateformat .Createtime "2006-01-02"}}</th>
      <th>{{.Role}}</th>

      <th>
      <a href="/user/deluser?userid={{.Id}}">删除</a></th>
    </tr>
    {{end}}
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