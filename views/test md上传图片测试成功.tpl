<!DOCTYPE html>
{{template "header"}}
<title>项目&目录 - 水利设计CMS系统</title>

</head>

<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>

<div class="col-lg-12">

  <form method="post" action="/category" enctype="multipart/form-data">
 <!--  <form method="post" action="/topic/addtopic1" enctype="multipart/form-data"> -->
    <div class="form-group">
      <label>项目名称</label>
      <input id="name" class="form-control"  placeholder="Enter ProjectName" name="name"></div>

    <div class="form-group">
      <label>项目编号</label>
      <input id="number" class="form-control"  placeholder="Enter ProjectNumber" name="number"></div>

    <label>项目简介:</label>

    <div class="form-group" id="test-editormd">
      <textarea style="display:none;" name="test-editormd-html-code">这个可以插入图片了，我了个去……</textarea>
    </div>

  <button type="submit" class="btn btn-default" onclick="return checkInput();">添加</button>

  <!--必须加return才能不跳转-->
</form>
</div>

<script type="text/javascript">

   function checkInput(){
    //是下面这段代码出了问题，等下修改
      var name=document.getElementById("name");
      if (name.value.length==0){
        alert("请输入项目名称");
        return false;
      }
      var name=document.getElementById("number");
      if (name.value.length==0){
        alert("请输入项目编号");
        return false;
      }


     var bb = "";  
     var temp = ""; 
     var a = document.getElementsByName("checkbox");//方法可返回带有指定名称的对象的集合
     for ( var i = 0; i<a.length; i++) {  
     if (a[i].checked) {  
     temp = a[i].value;  
     bb = bb + "," +temp;  
     }  
     } 
     document.getElementById("tempString").value = bb
      .substring(1, bb.length); 
    return true;  //这个return必须放最后，前面的值才能传到后台
   }

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

            var testEditor;
            $(function() {
                // $.get("./test.md", function(md) {
                    testEditor = editormd("test-editormd", {
                        width  : "100%",
                        height : 640,
                        path   : "/static/editor.md/lib/",
                        imageUpload : true,
                        imageFormats : ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
                        imageUploadURL : "/category/uploadimages",//这里需要好好写一个上传的控制器。
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


<!-- <!DOCTYPE html>
<html lang="zh">
    <head>
        <meta charset="utf-8" />
        <title>图片上传示例 - Editor.md examples</title>
        <link rel="stylesheet" href="css/style.css" />
        <link rel="stylesheet" href="../css/editormd.css" />
        <link rel="shortcut icon" href="https://pandao.github.io/editor.md/favicon.ico" type="image/x-icon" />
    </head>
    <body>
        <div id="layout" style="height: 2000px;background: #f6f6f6;">
            <header>
                <h1>图片上传示例</h1>
                <p>Image upload example</p>
            </header>
            <div id="test-editormd">
                <textarea style="display:none;">#### Settings

```javascript
{
    imageUpload    : false,
    imageFormats   : ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
    imageUploadURL : "./php/upload.php",
}
```

#### JSON data

```json
{
    success : 0 | 1,           // 0 表示上传失败，1 表示上传成功
    message : "提示的信息，上传成功或上传失败及错误信息等。",
    url     : "图片地址"        // 上传成功时才返回
}
```</textarea>
        </div>
        </div>        
        <script src="js/jquery.min.js"></script>
        <script src="../editormd.js"></script>
        <script type="text/javascript">
            $(function() {                
                var testEditor = editormd("test-editormd", {
                    width: "90%",
                    height: 640,
                    markdown : "",
                    path : '../lib/',
                    //dialogLockScreen : false,   // 设置弹出层对话框不锁屏，全局通用，默认为 true
                    //dialogShowMask : false,     // 设置弹出层对话框显示透明遮罩层，全局通用，默认为 true
                    //dialogDraggable : false,    // 设置弹出层对话框不可拖动，全局通用，默认为 true
                    //dialogMaskOpacity : 0.4,    // 设置透明遮罩层的透明度，全局通用，默认值为 0.1
                    //dialogMaskBgColor : "#000", // 设置透明遮罩层的背景颜色，全局通用，默认为 #fff
                    imageUpload : true,
                    imageFormats : ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
                    imageUploadURL : "./php/upload.php?test=dfdf",
                    
                    /*
                     上传的后台只需要返回一个 JSON 数据，结构如下：
                     {
                        success : 0 | 1,           // 0 表示上传失败，1 表示上传成功
                        message : "提示的信息，上传成功或上传失败及错误信息等。",
                        url     : "图片地址"        // 上传成功时才返回
                     }
                     */
                });
            });
        </script>
    </body>
</html> -->