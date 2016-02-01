<!DOCTYPE html>
{{template "header"}}
<title>项目&目录 - 水利设计CMS系统</title>
    <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.js"> </script>
    <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
    <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
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
      <label>项目编号：</label>
      <input id="number" class="form-control"  placeholder="Enter ProjectNumber:上传图片之前必须先填写项目编号和项目名称" name="number"></div> 
    <div class="form-group">
      <label>项目名称：</label>
      <input id="name" class="form-control"  placeholder="Enter ProjectName:上传图片之前必须先填写项目编号和项目名称" name="name"></div>
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
      <p> <font size="4" color="#A52A2A">全部选中将向数据库中写入7*7*6=294行记录，耗时较长，请选中1~2个体验即可。</font>
      </p>
    </div> -->
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
   

    //实例化编辑器
    //议使用工厂方法getEditor创建和引用编辑器实例，如果在某个闭包下引用该编辑器，直接调用UE.getEditor('editor')就能拿到相关的实例
    var ue = UE.getEditor('editor');

    /* 1.传入函数,命令里执行该函数得到参数表,添加到已有参数表里 */
 
ue.ready(function () {
// // 文件上传  UEditor中调用webupload事件不行。
// jQuery(function() {
//     // 当有文件添加进来的时候
//     uploader.on( 'fileQueued', function( file ) {
//     var name = $('#name').val();
//       uploader.option('formData', {
//         'categoryid':name,
//       });
// });

// });

ue.addListener('focus', function () {//startUpload start-upload startUpload beforeExecCommand是在插入图片之前触发
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

// ue.ready(function () {

// ue.addListener('focus', function () {//startUpload   beforeExecCommand是在插入图片之前触发

     // var name = $('#name').val();
      // alert(name)
    // ue.execCommand('serverparam', {
    //     'key': "1"
    // });


// });
// });

//     ue.addListener('beforeinsertimage',function() {
      

//     });
// });
    // function isFocus(e){
    //     alert(UE.getEditor('editor').isFocus());
    //     UE.dom.domUtils.preventDefault(e)
    // }
    // function setblur(e){
    //     UE.getEditor('editor').blur();
    //     UE.dom.domUtils.preventDefault(e)
    // }
    // function insertHtml() {
    //     var value = prompt('插入html代码', '');
    //     UE.getEditor('editor').execCommand('insertHtml', value)
    // }
    // function createEditor() {
    //     enableBtn();
    //     UE.getEditor('editor');
    // }
    // function getAllHtml() {
    //     alert(UE.getEditor('editor').getAllHtml())
    // }
    // function getContent() {
    //     var arr = [];
    //     arr.push("使用editor.getContent()方法可以获得编辑器的内容");
    //     arr.push("内容为：");
    //     arr.push(UE.getEditor('editor').getContent());
    //     alert(arr.join("\n"));
    // }
    // function getPlainTxt() {
    //     var arr = [];
    //     arr.push("使用editor.getPlainTxt()方法可以获得编辑器的带格式的纯文本内容");
    //     arr.push("内容为：");
    //     arr.push(UE.getEditor('editor').getPlainTxt());
    //     alert(arr.join('\n'))
    // }
    // function setContent(isAppendTo) {
    //     var arr = [];
    //     arr.push("使用editor.setContent('欢迎使用ueditor')方法可以设置编辑器的内容");
    //     UE.getEditor('editor').setContent('欢迎使用ueditor', isAppendTo);
    //     alert(arr.join("\n"));
    // }
    // function setDisabled() {
    //     UE.getEditor('editor').setDisabled('fullscreen');
    //     disableBtn("enable");
    // }

    // function setEnabled() {
    //     UE.getEditor('editor').setEnabled();
    //     enableBtn();
    // }

    // function getText() {
    //     //当你点击按钮时编辑区域已经失去了焦点，如果直接用getText将不会得到内容，所以要在选回来，然后取得内容
    //     var range = UE.getEditor('editor').selection.getRange();
    //     range.select();
    //     var txt = UE.getEditor('editor').selection.getText();
    //     alert(txt)
    // }

    // function getContentTxt() {
    //     var arr = [];
    //     arr.push("使用editor.getContentTxt()方法可以获得编辑器的纯文本内容");
    //     arr.push("编辑器的纯文本内容为：");
    //     arr.push(UE.getEditor('editor').getContentTxt());
    //     alert(arr.join("\n"));
    // }
    // function hasContent() {
    //     var arr = [];
    //     arr.push("使用editor.hasContents()方法判断编辑器里是否有内容");
    //     arr.push("判断结果为：");
    //     arr.push(UE.getEditor('editor').hasContents());
    //     alert(arr.join("\n"));
    // }
    // function setFocus() {
    //     UE.getEditor('editor').focus();
    // }
    // function deleteEditor() {
    //     disableBtn();
    //     UE.getEditor('editor').destroy();
    // }
    // function disableBtn(str) {
    //     var div = document.getElementById('btns');
    //     var btns = UE.dom.domUtils.getElementsByTagName(div, "button");
    //     for (var i = 0, btn; btn = btns[i++];) {
    //         if (btn.id == str) {
    //             UE.dom.domUtils.removeAttributes(btn, ["disabled"]);
    //         } else {
    //             btn.setAttribute("disabled", "true");
    //         }
    //     }
    // }
    // function enableBtn() {
    //     var div = document.getElementById('btns');
    //     var btns = UE.dom.domUtils.getElementsByTagName(div, "button");
    //     for (var i = 0, btn; btn = btns[i++];) {
    //         UE.dom.domUtils.removeAttributes(btn, ["disabled"]);
    //     }
    // }

    // function getLocalData () {
    //     alert(UE.getEditor('editor').execCommand( "getlocaldata" ));
    // }

    // function clearLocalData () {
    //     UE.getEditor('editor').execCommand( "clearlocaldata" );
    //     alert("已清空草稿箱")
    // }

</script>


</body>
</html>