<!DOCTYPE html>
{{template "header"}}
<title>自定义项目&目录 - 水利设计CMS系统</title>
</head>

<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>

<div class="col-lg-12">

  <form method="post" action="/category/userdefinedpost" enctype="multipart/form-data">
 <!--  <form method="post" action="/topic/addtopic1" enctype="multipart/form-data"> -->
    <div class="form-group">
      <label>项目名称</label>
      <input id="name" class="form-control"  placeholder="Enter ProjectName" name="name"></div>

    <div class="form-group">
      <label>项目编号</label>
      <input id="number" class="form-control"  placeholder="Enter ProjectNumber" name="number"></div>

    <!-- <div class="form-group"> -->
      <label>项目简介:</label>
    <div class="form-group" id="test-editormd">
      <textarea style="display:none;" name="test-editormd-html-code"></textarea>
    </div>
      
      <div class="input-group">
      <label>选择效果图：<input type="file" name="image" id="image" /></label>
      </div>
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
  // function check_file(){
  //    var n=0;
  //  var flist = document.all.file;
  //  if(flist.length != undefined){
  //   for(i=0;i<flist.length;i++){
  //    if(flist[i].value != undefined){
  //     alert(flist[i].value);
  //    } 
  //   }
  //  }else{
  //   alert(flist.value);
  //  }
  // }
  var testEditor;
            $(function() {
                // $.get("./test.md", function(md) {
                    testEditor = editormd("test-editormd", {
                        width  : "100%",
                        height : 640,
                        path   : "/static/editor.md/lib/",
                        imageUpload : true,
                        imageFormats : ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
                        imageUploadURL : "/category",//这里需要好好写一个上传的控制器。
                        // appendMarkdown : md,
                        saveHTMLToTextarea : true
                     /*
                     上传的后台只需要返回一个 JSON 数据，结构如下：
                     {
                        success : 0 | 1,           // 0 表示上传失败，1 表示上传成功
                        message : "提示的信息，上传成功或上传失败及错误信息等。",
                        url     : "图片地址"        // 上传成功时才返回
                     }
                     */
                    });
                // });

                //testEditor.getMarkdown();       // 获取 Markdown 源码
                //testEditor.getHTML();           // 获取 Textarea 保存的 HTML 源码
                //testEditor.getPreviewedHTML();  // 获取预览窗口里的 HTML，在开启 watch 且没有开启 saveHTMLToTextarea 时使用
            });

                $("#goto-line-btn").bind("click", function(){
                    testEditor.gotoLine(90);
                });
                
                $("#show-btn").bind('click', function(){
                    testEditor.show();
                });
                
                $("#hide-btn").bind('click', function(){
                    testEditor.hide();
                });
                
                $("#get-md-btn").bind('click', function(){
                    alert(testEditor.getMarkdown());
                });
                
                $("#get-html-btn").bind('click', function() {
                    alert(testEditor.getHTML());
                });                
                
                $("#watch-btn").bind('click', function() {
                    testEditor.watch();
                });                 
                
                $("#unwatch-btn").bind('click', function() {
                    testEditor.unwatch();
                });              
                
                $("#preview-btn").bind('click', function() {
                    testEditor.previewing();
                });
                
                $("#fullscreen-btn").bind('click', function() {
                    testEditor.fullscreen();
                });
                
                $("#show-toolbar-btn").bind('click', function() {
                    testEditor.showToolbar();
                });
                
                $("#close-toolbar-btn").bind('click', function() {
                    testEditor.hideToolbar();
                });
                
                $("#toc-menu-btn").click(function(){
                    testEditor.config({
                        tocDropdown   : true,
                        tocTitle      : "目录 Table of Contents",
                    });
                });
                
                $("#toc-default-btn").click(function() {
                    testEditor.config("tocDropdown", false);
                });
 </script>
</body>
</html>