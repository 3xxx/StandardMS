<!DOCTYPE html>
{{template "header"}}
<title>自定义项目&目录 - 水利设计CMS系统</title>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.min.js"> </script>
    <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
    <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
</head>

<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>

<div class="col-lg-12">

  <form method="post" action="/category/userdefinedpost" enctype="multipart/form-data">
 <!--  <form method="post" action="/topic/addtopic1" enctype="multipart/form-data"> -->
    <div class="form-group">
      <label>项目编号：</label>
      <input id="number" class="form-control"  placeholder="Enter ProjectNumber" name="number"></div>    

    <div class="form-group">
      <label>项目名称：</label>
      <input id="name" class="form-control"  placeholder="Enter ProjectName" name="name"></div>
    <!-- <div class="form-group"> -->
      <label>项目简介:</label>
      <div>
    <!-- <h1>项目简介:</h1> -->
       <script id="editor" type="text/plain" style="height:500px;"></script><!-- width:1024px; -->
      </div>
<!--     <div class="form-group" id="test-editormd">
      <textarea style="display:none;" name="test-editormd-html-code"></textarea>
    </div> -->
      
<!--       <div class="input-group">
      <label>选择效果图：<input type="file" name="image" id="image" /></label>
      </div> -->
<hr>
    

    <div class="col-md-4">
    <label>二级目录</label>
       <!-- <form id="form1" name="form1" onsubmit="check_file();"> -->
        <input type="text" id="num" name="num" style="width:50px;" value="1" />
        <input type="button" onclick="create_table()"  value="设定" />
        <table cellpadding="0" cellspacing="0" border="0" bgcolor="#CCCCCC" id="t1">
          <tr> <td><div id="d1"></div></td> </tr>
        </table>
       <!-- </form> -->
    </div>

    <div class="col-md-4">
      <label>三级目录</label>
       <!-- <form id="form2" name="form2" onsubmit="check_file();"> -->
        <input type="text" id="num2" name="num2" style="width:50px;" value="1" />
        <input type="button" onclick="create_table2()"  value="设定" />
        <table cellpadding="0" cellspacing="0" border="0" bgcolor="#CCCCCC" id="t2">
          <tr> <td><div id="d2"></div></td> </tr>
        </table>
       <!-- </form> -->
      
    </div>

    <div class="col-md-4">
      <label>四级目录</label>
       <!-- <form id="form3" name="form3" onsubmit="check_file();"> -->
        <input type="text" id="num3" name="num3" style="width:50px;" value="1" />
        <input type="button" onclick="create_table3()"  value="设定" />
        <table cellpadding="0" cellspacing="0" border="0" bgcolor="#CCCCCC" id="t3">
          <tr> <td><div id="d3"></div></td> </tr>
        </table>
       <!-- </form> -->
    </div>
<br />
<br />
<br />
<!-- <hr /> -->
<br />
<!-- <hr /> -->
<br />
<hr />
<!--      <tr>    
       <td colspan="4"><input type="button" class="btn btn-default" name="insert" value="增加目录行" onclick="insertNewRow()"/></td>    
       </tr> -->
  <!-- <input type="hidden" name="op" value="add"> -->
  <!-- <input type="hidden" id="tempString" name="tempString"/> -->
  <button style="float:left" type="submit" class="btn btn-default" >添加自定义项目&目录</button>
  <!--必须加return才能不跳转onclick="return Input();"-->
</form>
</div>

<script>
// function Input(){
//      var bb = "";  
//      var temp = ""; 
//      var a = document.getElementsByName("checkbox");
//      for ( var i = 0; i<a.length; i++) {  
//      if (a[i].checked) {  
//      temp = a[i].value;  
//      bb = bb + "," +temp;  
//      }  
//      } 
//      document.getElementById("tempString").value = bb
//       .substring(1, bb.length); 
//     return true;  //这个return必须放最后，前面的值才能传到后台
// }


  function create_table(){
   var num = document.getElementById("num").value;
   if(num == null ){
    num = 1;
   }
   var s1 = document.getElementById("d1");
   var vTable = document.createElement("Table");
   for(i=0; i<num; i++){
    vTr = vTable.insertRow();
    vTd = vTr.insertCell();
    vTd.innerHTML="目录"+(i+1)+"：<input type=\"text\" name=\"category2\" id=\"category2\" />";
    // vTd.innerHTML="目录"+(i+1)+'：<input type="'+text+'" name="'category2'" id='"category2" />';
   }
   s1.appendChild(vTable);
  }

  function create_table2(){
   var num = document.getElementById("num2").value;
   if(num == null ){
    num = 1;
   }
   var s1 = document.getElementById("d2");
   var vTable = document.createElement("Table");
   for(i=0; i<num; i++){
    vTr = vTable.insertRow();
    vTd = vTr.insertCell();
    vTd.innerHTML="目录"+(i+1)+"：<input type=\"text\" name=\"category3\" id=\"category3\" />";
   }//type=\"file\"
   s1.appendChild(vTable);
  } 

   function create_table3(){
   var num = document.getElementById("num3").value;
   if(num == null ){
    num = 1;
   }
   var s1 = document.getElementById("d3");
   var vTable = document.createElement("Table");
   for(i=0; i<num; i++){
    vTr = vTable.insertRow();
    vTd = vTr.insertCell();
    vTd.innerHTML="目录"+(i+1)+"：<input type=\"text\" name=\"category4\" id=\"category4\" />";
   }//type=\"file\"
   s1.appendChild(vTable);
  }  
//实例化编辑器
    //议使用工厂方法getEditor创建和引用编辑器实例，如果在某个闭包下引用该编辑器，直接调用UE.getEditor('editor')就能拿到相关的实例
    var ue = UE.getEditor('editor');

    /* 1.传入函数,命令里执行该函数得到参数表,添加到已有参数表里 */
 
ue.ready(function () {
ue.addListener('focus', function () {//startUpload   beforeExecCommand是在插入图片之前触发
     var name = $('#name').val();
      //  if (name.length==0){
      //   alert("请输入项目名称");
      //   return false;
      // }    
     var number = $('#number').val();
      //  if (name==""){
      //   alert("请输入项目编号");
      //   return false;
      // }    
      // alert(name)
    ue.execCommand('serverparam', {
        "number":number,
        'name': name,
    });
});
});
    
 </script>
</body>
</html>