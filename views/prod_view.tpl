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
  <!-- 这里添加项目目录树显示 -->
  <h3> <font size="5">{{.CategoryPhase.Title}}</font>
    <font size="4"> » {{.CategorySpec.Title}} » </font>
    <font size="5">{{.Category.Title}}</font>
  </h3>
  <h2>成果列表</h2>

<!--    <a href="/topic/add?id={{.Category.Id}}&mid=2" class="btn btn-primary">添加成果(一对一模式)</a> 
  <a href="/topic/add?id={{.Category.Id}}&mid=1" class="btn btn-default">添加成果(一对多(附件)模式)</a> -->

<!-- ！！！！！！！保留！！！！！！！！ -->
  <!-- <button id="directNextpage" onclick="window.location.reload('/topic/add?id={{.Category.Id}}&mid=1')">Direct Next Page</button> 这个方法跳不出去iframe
  onclick="window.open('/topic/add?id={{.Category.Id}}&mid=1')"新标签页中打开
  onclick="parent.location.href='/topic/add?id={{.Category.Id}}&mid=2'跳出iframe重新打开
  <button id="directNextpage" class="btn btn-default" onclick="window.open('/topic/add?id={{.Category.Id}}&mid=1')">A+(1to1)</button> -->


<!-- <div class="row"> -->
<!-- <div class="col-sm-5"> -->
<!--   <button id="directNextpage" class="btn btn-default" onclick="window.open('/topic/add?id={{.Category.Id}}&mid=1')">A+(1to1)</button> -->
<button id="directNextpage" class="btn btn-default" onclick="parent.location.href='/topic/add?id={{.Category.Id}}&mid=6'">A+(1to1)</button>
  <button id="directNextpage" class="btn btn-default" onclick="parent.location.href='/topic/add?id={{.Category.Id}}&mid=8'">B+(1toM)</button>
  
  <button id="directNextpage" class="btn btn-default" onclick="window.open('/catalog/view?id={{.Category.Id}}')">查看计划目录</button>
  <input type="button" id="print" value="打印" onclick="print()" class="btn btn-default">
  <!-- <a href="/catalog/view?id={{.Category.Id}}" class="btn btn-default">Plan Catalog</a> -->
<!-- </div> -->
<button type="submit" class="btn btn-default" onclick="return deleteall();">删除选中</button>
<button type="submit" class="btn btn-default" onclick="return downloadall();">下载选中</button>
<button type="submit" class="btn btn-default" onclick="ExportToExcel();">导出excel</button>
<!-- <div class="col-sm-1"> -->
  <form id="form1" method="post" action="/topic/deleteall" enctype="multipart/form-data">
    <input type="hidden" id="tempstring" name="tempstring"/>
    <input type="hidden" id="cid" name="cid" value="{{.Category.Id}}" />
<!--     <button type="submit" type="button" class="btn btn-default" onclick="return deleteall();">删除选中</button> -->
  </form>
   <!-- </div> -->

  <!-- <div class="col-sm-1"> -->
    <form id="form2" method="post" action="/topic/downloadall" enctype="multipart/form-data">
      <input type="hidden" id="tempstring1" name="tempstring1"/>
      <input type="hidden" id="cid" name="cid" value="{{.Category.Id}}" />
<!--       <button type="submit" type="button" class="btn btn-default" onclick="return downloadall();">下载选中</button> -->
    </form>
  <!-- </div> -->


    <form id="form3" method="post" action="/topic/exporttoexcel" >
      <input type="hidden" name="id" value="{{.Category.Id}}"/>
      <input type="hidden" name="path" value="{{.Category.DiskDirectory}}"/>
      <input type="hidden" name="filename" value="{{.Category.Title}}"/>
      <!-- {{.CategoryProj.Number}}{{.CategoryProj.Title}}\{{.CategoryPhase.Title}}\{{.CategorySpec.Title}}\{{.Category.Title}}\<button type="submit" class="btn btn-default">导出excel</button> -->
    </form>



  <!-- <div class="row">
<div class="col-sm-1"> -->
  <hr />
      <form class="form-inline" method="post" action="/catalog/import_xls_catalog" enctype="multipart/form-data">
        <div class="form-group">
          <label>选择excel</label>
          <input type="file" class="form-control" name="excel" id="excel"></div>
        <input type="hidden" name="id" value="{{.Category.Id}}"/>
        <button type="submit" class="btn btn-default">提交</button>
      </form>
        <!--        <div class="input-group">
        <label>选择excel：</label>
        <input type="file" class="form-control" name="excel" id="excel"/>
      </div>
      <input type="hidden" name="id" value="{{.Category.Id}}"/>
      <button type="submit" class="btn btn-default">提交</button>
      -->
    
    

<!--     <script type="text/javascript">
  function fun1(){
    var form1 = window.document.getElementById("form1");//获取form1对象
    form1.submit(); 
  };
  </script>


    <body>
    <form id="form3" action="/catalog/import_xls_catalog" method="get" >
      <input type="button" onclick="fun1();" value="sss" />    
    </form>
    </body> -->


  <!-- </div>   -->
      <!-- <button class="btn btn-primary" id="export">导出excel</button>  -->
<!-- 上面这个因为用ajax传值，所以不能成功下载。ajax不能重定向？
<a href="/topic/ExportToExcel?id={{.Category.Id}}" class="btn btn-primary">导出下载</a> -->


<table class="table table-striped">
  <thead>
    <tr>
      <th style="cursor: pointer"><input type="checkbox" name="checkall" onclick="CheckAll();" />全选#{{.Length}}</th>
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
    {{range $index,$elem:=.Chengguo}}
    <tr>
      <th><input type="checkbox" name="checkbox" value="{{.Id}}"/>{{$index}}</th>
      <th><a href="/topic/view/{{.Id}}">{{substr .Tnumber 0 15}}</a></th>
      <th><a href="/topic/view_b/{{.Id}}" title={{.Title}} target="_blank">{{substr .Title 0 15}}</a></th>
      <th>{{.Category}}</th>
      <th>{{.Author}}</th>
      <th>{{dateformat .Updated "2006-01-02"}}</th>
      <!-- <th>{{dateformat .ReplyTime "2006-01-02"}}</th> -->
      <!-- <th>{{.Views}}</th> -->
      <!-- <th>{{.ReplyCount}}</th> -->
<!--       <th><a href="/topic/view/{{.Id}}">下载</a>
      <a href="/topic/modify?tid={{.Id}}">修改</a>
      <a href="/topic/delete?tid={{.Id}}">删除</a></th> -->
        <th>
          <a href="/topic/view/{{.Id}}"><i class="glyphicon glyphicon-download-alt"></i>下载</a>
          <a href="/topic/modify?tid={{.Id}}"><i class="glyphicon glyphicon-edit"></i>修改</a>
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

<script>
    function CheckAll()
    {
    var a = document.getElementsByTagName('input');
    var b = document.getElementsByName('checkall');
    var n = a.length;  
    var allchecked = true;
         var bb = "";  
         var temp = "";
    if(b[0].checked)
    allchecked = true;
    else
    allchecked = false;
    for(var i=0;i<n;i++)
    {
    if(allchecked)
    a[i].checked = true;
    else
    a[i].checked = false;
    }
   }

   function deleteall()
   {
     var bb = "";  
     var temp = ""; 
     var a = document.getElementsByName("checkbox");
     for ( var i = 0; i<a.length; i++) {  
     if (a[i].checked) {  
     temp = a[i].value;  
     bb = bb + "," +temp;  
     }  
     } 
     document.getElementById("tempstring").value = bb
      .substring(1, bb.length); 

    var form1 = window.document.getElementById("form1");//获取form1对象
    form1.submit();  

    return true;  //这个return必须放最后，前面的值才能传到后台    
   }

   function downloadall()
   {
     var bb = "";  
     var temp = ""; 
     var a = document.getElementsByName("checkbox");
     for ( var i = 0; i<a.length; i++) {  
     if (a[i].checked) {  
     temp = a[i].value;  
     bb = bb + "," +temp;  
     }  
     } 
     document.getElementById("tempstring1").value = bb
             .substring(1, bb.length);

     var form2 = window.document.getElementById("form2");//获取form1对象
    form2.submit(); 

    return true;  //这个return必须放最后，前面的值才能传到后台    
   }

   function ExportToExcel()
   {
    var form3 = window.document.getElementById("form3");//获取form1对象
    form3.submit(); 

    // return true;  //这个return必须放最后，前面的值才能传到后台    
   }
</script>
<script type="text/javascript">
  // $(document).ready(function() {
  // $("table").tablesorter();
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
// });
</script>

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
                data: { 
                  id: "{{.Category.Id}}",
                  tel:$("#tnumber").val(),
                  path:"{{.CategoryProj.Number}}{{.CategoryProj.Title}}\\{{.CategoryPhase.Title}}\\{{.CategorySpec.Title}}\\{{.Category.Title}}\\",
                  filename:"{{.Category.Title}}"
                },
                // success: function(responseText) {
                // $("#divResult").html(responseText);
                // }
                success:function(data){//数据提交成功时返回数据
                    // alert(data);
                    alert("导出成功！")
                }
            });
            return true;//这里true和false结果都一样。不刷新页面的意思？
 });
});
  </script>
</body>
</html>