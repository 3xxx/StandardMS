<!DOCTYPE html>
{{template "header"}}
<title>项目&目录 - 水利设计CMS系统</title>
    <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.js"> </script>
    <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
    <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
    <script src="/static/ueditor/ueditor.parse.js"></script>
</head>

<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>

<div class="col-lg-12">
  <form method="post" action="/category/modifycategory" enctype="multipart/form-data">
    <div class="form-group">
      <label>项目名称</label>
      <input id="name" class="form-control"  placeholder="Enter ProjectName" name="name" value="{{.Category.Title}}"></div>

    <div class="form-group">
      <label>项目编号</label>
      <input id="number" class="form-control"  placeholder="Enter ProjectNumber" name="number" value="{{.Category.Number}}"></div>

    <label>项目简介:</label>
     <!-- <div id="content" class="content" name="content"> -->
     <!-- 用str2html不转义，就可以用uparse解析了 -->
    <!-- {{str2html .Category.Content}}</div> -->
<div id="content">
    <script id="editor" type="text/plain" style="height:500px;"></script>
</div>
<!--     <div class="form-group" id="test-editormd">
      <textarea style="display:none;" name="test-editormd-html-code">{{.Category.Content}}</textarea>
    </div> -->
<!--       <div class="input-group">
      <label>选择效果图：<input type="file" name="image" id="image" />{{.Filename}}</label><br/>
      <img style="-webkit-user-select: none; cursor: zoom-in;" src="{{.Category.Route}}" width="200">
    </div> -->
<hr>
    <!-- <input type="hidden" name="op" value="modify"> -->
    <div class="col-md-4">
      <label>阶段</label>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="jd" name="checkbox" value="ghj" {{if .Ghj}}checked {{end}}>规划阶段</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="jd" name="checkbox" value="xj" {{if .Xj}}checked{{end}}>项目建议书阶段</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="jd" name="checkbox" value="ky" {{if .Ky}}checked{{end}}>可行性研究阶段</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="jd" name="checkbox" value="cs" {{if .Cs}}checked{{end}}>初步设计阶段</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="jd" name="checkbox" value="zb" {{if.Zb}}checked{{end}}>招标阶段</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="jd" name="checkbox" value="sgt" {{if.Sgt}}checked{{end}}>施工图阶段</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="jd" name="checkbox" value="jgt" {{if .Jgt}}checked{{end}}>竣工图阶段</label>
      </div>

    </div>

    <div class="col-md-4">
      <label>专业</label>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="zy" name="checkbox" value="gh" {{if .Gh}}checked{{end}}>规划专业</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="zy" name="checkbox" value="sg" {{if .Sg}}checked{{end}}>水工专业</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="zy" name="checkbox" value="jd" {{if .Jd}}checked{{end}}>机电专业</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="zy" name="checkbox" value="shg" {{if .Shg}}checked{{end}}>施工专业</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="zy" name="checkbox" value="dz" {{if .Dz}}checked{{end}}>地质专业</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="zy" name="checkbox" value="ys" {{if .Ys}}checked{{end}}>预算专业</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="zy" name="checkbox" value="zh" {{if .Zh}}checked{{end}}>资环专业</label>
      </div>

    </div>

    <div class="col-md-4">
      <label>成果类型</label>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="lx" name="checkbox" value="dwg" {{if .Dwg}}checked{{end}}>DWG</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="lx" name="checkbox" value="doc" {{if .Doc}}checked{{end}}>doc/docx</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="lx" name="checkbox" value="xls" {{if .Xls}}checked{{end}}>xls/xlsx</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="lx" name="checkbox" value="pdf" {{if .Pdf}}checked{{end}}>pdf</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="lx" name="checkbox" value="jpg" {{if .Jpg}}checked{{end}}>jpg</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="lx" name="checkbox" value="tif" {{if .Tif}}checked{{end}}>tif</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" id="diary" name="checkbox" value="diary" {{if .Diary}}checked{{end}}>设代日记</label>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox"  id="lx" name="checkall" onclick="CheckAll();">全选</label>
      </div>

    </div>
    <input type="hidden" id="tempString" name="tempString"/>
    <input type="hidden" id="cid" name="cid" value="{{.Category.Id}}" />
    <!-- <button type="submit" class="btn btn-default" onclick="return checkInput();">添加</button>
  -->
  <button type="submit" class="btn btn-default" onclick="return checkInput();">修改</button>
  <!--必须加return才能不跳转-->
</form>
</div>

<script type="text/javascript">

    //实例化编辑器
    //议使用工厂方法getEditor创建和引用编辑器实例，如果在某个闭包下引用该编辑器，直接调用UE.getEditor('editor')就能拿到相关的实例
    var ue = UE.getEditor('editor');

//  setTimeout(function(){ uParse('#content', {
//  'highlightJsUrl':'{/static/ueditor/third-party/SyntaxHighlighter/shCore.js',
//  'highlightCssUrl':'/static/ueditor/third-party/SyntaxHighlighter/shCoreDefault.css'})
// }, 300);   
ue.addListener("ready", function () {
uParse('.content', {
    rootPath: '/static/ueditor/'
});
});

$(function(){
        var content =$('#content').val();
        //判断ueditor 编辑器是否创建成功
        ue.addListener("ready", function () {
        // editor准备好之后才可以使用
        ue.setContent({{str2html .Category.Content}});
        });
    });
// ue.ready(function () {
//     function setContent(isAppendTo) {
//         var arr = [];
//         // arr.push("使用editor.setContent('欢迎使用ueditor')方法可以设置编辑器的内容");
//         UE.getEditor('editor').setContent({{.Category.Content}} , isAppendTo);
//         // alert(arr.join("\n"));
//     };
// });
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
      var name=document.getElementById("jd");
      if (name.length==0){
        alert("请至少选择一个阶段");
        return false;
      }

      var name=document.getElementById("zy");
      if (name.length==0){
        alert("请至少选择一个专业");
        return false;
      }

      var name=document.getElementById("lx");
      if (name.length==0){
        alert("请至少选择一个成果类型");
        return false;
      }

     var bb = "";  
     var temp = ""; 
     var a = document.getElementsByName("checkbox");
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

</script>

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