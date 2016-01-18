<!DOCTYPE html>
{{template "header"}}
<title>项目&目录 - 水利设计CMS系统</title>

</head>

<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>

<div class="col-lg-12">
<!--             <div class="btns">
                <button id="goto-line-btn">Goto line 90</button>
                <button id="show-btn">Show editor</button>
                <button id="hide-btn">Hide editor</button>
                <button id="get-md-btn">Get Markdown</button>
                <button id="get-html-btn">Get HTML</button>
                <button id="watch-btn">Watch</button>
                <button id="unwatch-btn">Unwatch</button>
                <button id="preview-btn">Preview HTML (Press Shift + ESC cancel)</button>
                <button id="fullscreen-btn">Fullscreen (Press ESC cancel)</button>
                <button id="show-toolbar-btn">Show toolbar</button>
                <button id="close-toolbar-btn">Hide toolbar</button>
                <button id="toc-menu-btn">ToC Dropdown menu</button>
                <button id="toc-default-btn">ToC default</button>
            </div> -->
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
      <textarea style="display:none;" name="test-editormd-html-code"></textarea>
    </div>

      <div class="input-group">
      <label>选择效果图：<input type="file" name="image" id="image" /></label>
      <p> <font size="4" color="#A52A2A">全部选中将向数据库中写入7*7*6=294行记录，耗时较长，请选中1~2个体验即可。</font>
      </p>
    </div>
<hr>
    <!-- <input type="hidden" name="op" value="add"> 这句应该没什么用了，因为controller里用post函数了-->

    <div class="col-md-4">
      <label>阶段</label>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="jd" name="checkbox" value="ghj">规划阶段</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="jd" name="checkbox" value="xj">项目建议书阶段</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="jd" name="checkbox" value="ky">可行性研究阶段</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="jd" name="checkbox" value="cs">初步设计阶段</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="jd" name="checkbox" value="zb">招标阶段</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="jd" name="checkbox" value="sgt">施工图阶段</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="jd" name="checkbox" value="jgt">竣工图阶段</label>
      </div>

    </div>

    <div class="col-md-4">
      <label>专业</label>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="zy" name="checkbox" value="gh">规划专业</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="zy" name="checkbox" value="sg">水工专业</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="zy" name="checkbox" value="jd">机电专业</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="zy" name="checkbox" value="shg">施工专业</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="zy" name="checkbox" value="dz">地质专业</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="zy" name="checkbox" value="ys">预算专业</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="zy" name="checkbox" value="zh">资环专业</label>
      </div>

    </div>

    <div class="col-md-4">
      <label>成果类型</label>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="lx" name="checkbox" value="dwg">DWG</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="lx" name="checkbox" value="doc">doc/docx</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="lx" name="checkbox" value="xls">xls/xlsx</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="lx" name="checkbox" value="pdf">pdf</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="lx" name="checkbox" value="jpg">jpg</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="lx" name="checkbox" value="tif">tif</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="lx" name="checkbox" value="diary">设代日记</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox"  id="lx" name="checkall" onclick="CheckAll();">全选</label>
      </div>

    </div>
    <input type="hidden" id="tempString" name="tempString"/>
    <!-- <button type="submit" class="btn btn-default" onclick="return checkInput();">添加</button>
  -->            
  <!-- <input type="submit" class="btn btn-primary" value="回复" id="submit" disabled="disabled"> -->
  <button type="submit" class="btn btn-default" onclick="return checkInput();">添加</button>

  <!--必须加return才能不跳转-->
</form>
</div>

<script type="text/javascript">
// document.getElementById()返回对拥有指定 id 的第一个对象的引用。
// document.getElementsByName()返回带有指定名称的对象集合。
// document.getElementsByTagName()返回带有指定标签名的对象集合。
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
//下面这3个无效
     //  var name=document.getElementById("jd");
     //  alert(name.value[1]);
     //  var j=0;
     //       for ( var i = 0; i<name.value.length; i++) {  
     //       if (name[i].checked==false) {  
     //         j=j+1;
     //        } 
     //      }  
     //    if (j>1){
     // // if (name.value.length==0){   
     //    alert("请至少选择一个阶段");
     //   return false;
     // } 
     //  var name=document.getElementById("zy");//返回对拥有指定 ID 的第一个对象的引用
     //  if (name.value.length==0){
     //    alert("请至少选择一个专业");
     //    return false;
     //  }
     //  var name=document.getElementById("lx");
     //  if (name.value.length==0){
     //    alert("请至少选择一个成果类型");
     //    return false;
     //  }

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
   
//秦改进了原代码，原代码循环所有的checkbox，实际有用的还是最后一个
// for(var i=0;i<n;i++)
// {
// if(a[i].checked)
// allchecked = true;
// else
// allchecked = false;
// }

    // var editor, editor_edit;
    // $(function() {
    //   $('#submit').attr('disabled', true);

    //   editor = createEditorMd("editormd", "#submit");

    //   $('.edit-comment').click(function () {
    //     var commentId = $(this).attr('comment-id');
    //     $.getJSON("/comment/" + commentId + ".json", function (data) {
    //       $('#edit-comment-form').attr('action', '/comment/' + commentId + '/edit');
    //       $('#myModal').on('shown.bs.modal', function (e) {
    //         if (editor_edit) {
    //           editor_edit.setMarkdown(data.markdown);
    //         } else {
    //           editor_edit = createEditorMd("editormd-edit", "#edit-submit", data.markdown);
    //         }
    //       });

    //       $('#myModal').modal({});
    //     });
    //   });
    // });
                    // markdown : "",
                    // path : '../lib/',
                    //dialogLockScreen : false,   // 设置弹出层对话框不锁屏，全局通用，默认为 true
                    //dialogShowMask : false,     // 设置弹出层对话框显示透明遮罩层，全局通用，默认为 true
                    //dialogDraggable : false,    // 设置弹出层对话框不可拖动，全局通用，默认为 true
                    //dialogMaskOpacity : 0.4,    // 设置透明遮罩层的透明度，全局通用，默认值为 0.1
                    //dialogMaskBgColor : "#000", // 设置透明遮罩层的背景颜色，全局通用，默认为 #fff
                    // imageUpload : true,
                    // imageFormats : ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
                    // imageUploadURL : "./php/upload.php?test=dfdf",  
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
// document.getElementById("tempString").value = bb;
</script>



<!--         <script type="text/javascript">
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
 -->
<!--  // var valus = document.getElementsByName("aaa");
  // if(a.length!=0)
  // {
  // var str = ""；
  // for(var i=0;i<a.length;i++)  
  // {
  // if(a[i].checked)
  // {
  // str+= a[i].value;
// }
// }
// } -->

</body>
</html>