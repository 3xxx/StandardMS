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
  <h3> <font size="5">{{.CategoryProj.Title}}</font> <font size="4"> » {{.CategoryPhase.Title}} » </font>
    <font size="4">{{.CategorySpec.Title}} » </font>
    <font size="5">{{.Category.Title}}</font>
  </h3>

  <h2>成果列表</h2>
  <div class="col-md-6">
  <a href="/topic/add?id={{.Category.Id}}&mid=1" class="btn btn-default">添加成果(一对一模式)</a>
  <a href="/topic/add?id={{.Category.Id}}&mid=2" class="btn btn-default">添加成果(一对多(附件)模式)</a>
    <input type="button" id="print" value="打印" onclick="print()" class="btn btn-default">
    <a href="/catalog/view?id={{.Category.Id}}" class="btn btn-default">查看计划目录</a>
  <!-- </div> -->
<!-- <div class="col-md-7">  -->
     <!-- <input type="hidden" name="tnumber" id="tnumber" value="{{.Category.Id}}"/> -->
    <form style="float:right" method="post" action="/topic/exporttoexcel" > 
     <input type="hidden" name="id" value="{{.Category.Id}}"/><!--这里面原来加入了id=“”不成功，去掉id=后即可-->
     <input type="hidden" name="path" value="{{.CategoryProj.Number}}{{.CategoryProj.Title}}\{{.CategoryPhase.Title}}\{{.CategorySpec.Title}}\{{.Category.Title}}\"/>
     <input type="hidden" name="filename" value="{{.Category.Title}}"/>
     <button type="submit" class="btn btn-primary">导出excel</button>
   </form>
    <form method="post" action="/catalog/import_xls_catalog" enctype="multipart/form-data">
<!--    <div class="form-group">
    <input type="file" class="form-control" placeholder="Search">
  </div> -->
    <div class="input-group">
      <label>选择excel：<input type="file" class="form-control" name="excel" id="excel" /></label>
      <br/>
    </div>
    <input type="hidden" name="id" value="{{.Category.Id}}"/>
    <button type="submit" class="btn btn-default" >提交</button>
  </form>
 </div>
 <div class="col-md-1">
<form style="float:left" method="post" action="/topic/deleteall" enctype="multipart/form-data">
 <input type="hidden" id="tempstring" name="tempstring"/>
 <input type="hidden" id="cid" name="cid" value="{{.Category.Id}}" />
 <button type="submit" class="btn btn-default" onclick="return deleteall();">删除选中</button>
 </form>
 </div>
 <div class="col-md-1">
 <form method="post" action="/topic/downloadall" enctype="multipart/form-data">
 <input type="hidden" id="tempstring1" name="tempstring1"/>
 <input type="hidden" id="cid" name="cid" value="{{.Category.Id}}" />
 <button type="submit" class="btn btn-default" onclick="return downloadall();">下载选中</button>
  </form>
  </div>
  <!-- /topic/add?id={{.Category.Id}}&mid=1 -->
  <!-- <a href="" class="btn btn-default">添加成果(多对一模式)</a>
-->
<!-- /topic/add?id={{.Category.Id}}&mid=3 -->
<!--       <a href="/topic/add?id={{.Category.Id}}&mid=3" class="btn btn-primary">写设代日记</a>
-->
<!-- <input type="button" id="export" value="导出" onclick="exportToExcel()" class="btn btn-default"> -->
<!-- <input type="button" id="print" value="打印" onclick="print()" class="btn btn-default"> -->
<!-- <script>
function validate(){
  if("success"){
    return true;
}else{
   return false;
}
}
</script> -->

<!-- <script >
 //导出到Excel
 function exportToExcel() {
  if(document.getElementById("title")) {
   try {
    var oRangeRef = document.body.createTextRange();
    oRangeRef.execCommand("Copy");
    var appExcel = new ActiveXObject("Excel.Application");
    appExcel.visible = true;
    appExcel.Workbooks.Add().WorkSheets.Item(1).Paste();
   } catch(e) {
    alert("出错啦！可能是浏览器或者是数据量太大咯哦！");
    return;
   }
   appExcel = null;
   oRangeRef = null;
  }
 }
 </script> -->

     <!-- <button type="submit" class="btn btn-default">提交日记</button>  -->
     <!-- 上面这行代码用html自己的方式提交 -->
     <!-- 下面这行代码用jquery的方式提交data。当按钮被点击时执行Javascript代码：: -->
     <!-- <input required="required" id="tnumber" type="text" class="form-control" name="tnumber" placeholder="Enter Data"></div> -->
     <!-- <input type="hidden" name="tnumber" id="tnumber" value="{{.Category.Id}}"/> -->
     <!-- <button class="btn btn-primary" id="export">导出excel</button> -->
  <!-- <div class="form-group"> -->
   <!-- <input type="button" value="导出" class="btn btn-primary" onclick="function()" /> -->
   <!-- <label for="exampleInputFile">File input</label> -->
  <!-- <input type="file" id="exampleInputFile" value="导出" onclick="function()"> -->
  <!-- </div> -->


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
 // });
// $(function(){
//     $('#send').click(function(){
//          $.ajax({
//              type: "GET",
//              url: "test.json",
//              data: {username:$("#username").val(), content:$("#content").val()},
//              dataType: "json",
//              success: function(data){
//                          $('#resText').empty();   //清空resText里面的所有内容
//                          var html = ''; 
//                          $.each(data, function(commentIndex, comment){
//                              html += '<div class="comment"><h6>' + comment['username']+ ':</h6><p class="para"' + comment['content']+ '</p></div>';
//                          });
//                          $('#resText').html(html);
//                       }
//          });
//     });
// });

// $.ajax({
//     url: "/topic/addtopic3", 
//     type: "POST",
//     data: { id: 2, name: "aaa", tel: "~!@#$%^&*()_+-=<>?|", xxxx: "要多少还可以写多少", encoding: "见鬼去吧。?& :)" },
//     success: function(responseText) {
//         $("#divResult").html(responseText);
//     }
// });
// $("button").click(function(){
//   $.post("demo_test_post.asp",
//   {
//     name:"Donald Duck",
//     city:"Duckburg"
//   },
//   function(data,status){
//     alert("Data: " + data + "\nStatus: " + status);
//   });
// });
</script>

<table class="table table-striped">
<thead>
  <tr>
    
    <th style="cursor: pointer"><input type="checkbox" name="checkall" onclick="CheckAll();" />全选#{{.Length}}</th>
    <th style="cursor: pointer">成果编号</th>
    <th style="cursor: pointer">成果名称</th>
    <th style="cursor: pointer">最后更新</th>
    <th style="cursor: pointer">浏览</th>
    <th style="cursor: pointer">回复数</th>
    <th>操作</th>
  </tr>
</thead>
<tbody>
  {{range $index,$elem:=.Chengguo}}
  <tr>
    
    <th><input type="checkbox" id="jd" name="checkbox" value="{{.Id}}"/>{{$index}}</th>
    <th>
      <a href="/topic/view_b/{{.Id}}">{{substr .Tnumber 0 15}}</a>
    </th>
    <th>
      <a href="/topic/view_b/{{.Id}}" title="{{.Title}}">{{substr .Title 0 15}}</a>
    </th>
    <th>{{dateformat .Updated "2006-01-02"}}</th>
    <th>{{.Views}}</th>
    <th>{{.ReplyCount}}</th>
    <th>
      <a href="/topic/view_b/{{.Id}}">下载</a>
      <a href="/topic/modify?tid={{.Id}}">修改</a>
      <a href="/topic/delete?tid={{.Id}}">删除</a>
    </th>
    <!--<th>
    <a href="/topic?op=del&id={{.Id}}">删除</a>
  </th>-->
  
</tr>
{{end}}
</tbody>
</table>

</div>

<!-- <script>
function toggle_checkall(field_name, state) {
  var checkboxes = document.getElementsByTagName('input');
  var count = checkboxes.length;
  for (var i = 0; i < count; i++) {
    if (checkboxes[i].type == "checkbox"
        && checkboxes[i].name == field_name + "_ids[]") {
      checkboxes[i].checked = state;
    }
  }
}
</script> -->
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
             .substring(1, bb.length); //bb=",102,101,100,99"所以要截掉第一个","号。
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
    return true;  //这个return必须放最后，前面的值才能传到后台    
   }
</script>
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