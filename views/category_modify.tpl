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
    <link rel="stylesheet" type="text/css" href="/static/fex-team-webuploader/css/webuploader.css">
<!--引入JS-->
<script type="text/javascript" src="/static/fex-team-webuploader/dist/webuploader.js"></script>
</head>

<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>

<div class="col-lg-12">
  <form method="post" action="/category/modifycategory" enctype="multipart/form-data">
   <input type="hidden" id="number" name="number" value="{{.Category.Number}}"/>
   <input type="hidden" id="name" name="name" value="{{.Category.Title}}"/>
   <input type="hidden" id="categoryid" name="categoryid" value="{{.Category.Id}}"/>
   <input type="hidden" id="route" name="route" value=""/>
    <div class="form-group">
      <label>项目名称</label>
      <input id="name" class="form-control"  placeholder="Enter ProjectName" name="name" value="{{.Category.Title}}"></div>

    <div class="form-group">
      <label>项目编号</label>
      <input id="number" class="form-control"  placeholder="Enter ProjectNumber" name="number" value="{{.Category.Number}}"></div>

    <label>封面文字</label>
    <!-- <div id="content" class="content" name="content">
    -->
    <!-- 用str2html不转义，就可以用uparse解析了 -->
    <!-- {{str2html .Category.Content}}</div>
    -->
    <div id="editor_cover">
      <script id="editor_cover" type="text/plain" style="height:500px;"></script>
    </div>
    <label>封面图片</label>
      <!-- <div> -->
        <!-- <script id="editor_photo" name="editor_photo" type="text/plain" style="height:500px;"></script>--> <!-- width:1024px; -->  
      <!-- </div> -->
      <!--dom结构部分-->
      <div id="uploader-demo">
          <!--用来存放item-->
          <div id="fileList" class="uploader-list"></div>
          <div id="filePicker">选择图片</div>
       <img src="{{.Category.Route}}" width="800" align="middle">   
      </div>
      
    <label>项目简介:</label>
    <!-- <div id="content" class="content" name="content">
    -->
    <!-- 用str2html不转义，就可以用uparse解析了 -->
    <!-- {{str2html .Category.Content}}</div>
    -->
    <div id="editor_property">
      <script id="editor_property" type="text/plain" style="height:500px;"></script>
    </div>

    <input type="hidden" id="tempString" name="tempString"/>
    <input type="hidden" id="cid" name="cid" value="{{.Category.Id}}" />
    <!-- <button type="submit" class="btn btn-default" onclick="return checkInput();">添加</button>
  -->
  <button type="submit" class="btn btn-default" >修 改</button>
  <!--onclick="return checkInput();"必须加return才能不跳转-->
</form>
</div>

<hr>
        <!-- {{range $k,$v :=.Categoryjieduan}}
          <li>
            <a href="#" aria-expanded="false"><i class="glyphicon glyphicon-stop"></i>&nbsp;{{.Title}} <span class="glyphicon arrow"></span></a>
            <ul aria-expanded="false">
            {{range $k1,$v1 :=$.Categoryzhuanye}}
                {{if eq $v.Id $v1.ParentId}}
              <li>
                <a href="#" aria-expanded="false"><i class="glyphicon glyphicon-th-large"></i>&nbsp;{{.Title}} <span class="fa plus-minus"></span></a>
                <ul aria-expanded="false">
                      {{range $k2,$v2 :=$.Categorychengguo}}
                          {{if eq $v1.Id $v2.ParentId}}
                          <li><a href="/category/view?id={{.Id}}" target='main'>&nbsp;&nbsp;<i class="glyphicon glyphicon-th"></i>&nbsp;{{.Title}}-<span class="badge">{{.TopicCount}}</span></a></li>
                          {{end}}
                      {{end}}
                </ul>
              </li>
                {{end}}
              {{end}}
            </ul>
          </li>
        {{end}} -->


<div class="col-lg-12">
  <table class="table table-striped">
    <thead>
      <tr>
        <th>设计阶段</th>
        <th>文件类型</th>
        <th>专业分类</th>
        <th>操作</th>
      </tr>
    </thead>

    <tbody>
      {{range $k,$v :=.Categoryjieduan}}
      <tr>
        <th>{{.Title}}</th>
        <th></th>
        <th></th>
        
        <th>
          <input type="hidden" id="{{.Id}}" value="{{$v.Id}}"/>
          <a href="" onclick="prom('{{.Id}}')">添加同级</a>
          <a href="" onclick="prom1('{{.Id}}')">添加下级</a>
          <a href="/modify?id={{.Id}}">修改</a>
          <a href="/delete?id={{.Id}}">删除</a>
        </th>
                {{range $k1,$v1 :=$.Categoryzhuanye}}
                {{if eq $v.Id $v1.ParentId}}

                 <tr>
                   <th></th>
                 <th>{{.Title}}</th>
                 <th></th>
                 
                 <th>
                  <input type="hidden" id="{{.Id}}" value="{{$v1.Id}}"/>
                   <a href="" onclick="prom('{{.Id}}')">添加同级</a>
                   <a href="" onclick="prom1('{{.Id}}')">添加下级</a>
                  <a href="/modify?id={{.Id}}">修改</a>
                  <a href="/delete?id={{.Id}}">删除</a>
                 </th>
                 </tr>
                        {{range $k2,$v2 :=$.Categorychengguo}}
                        {{if eq $v1.Id $v2.ParentId}}


                 <tr>
                   <th></th>
                   <th></th>
                  <th>{{.Title}}</th>
                                  
                  <th>
                  <input type="hidden" id="{{.Id}}" value="{{$v2.Id}}"/>
                  <a href="" onclick="prom('{{.Id}}')">添加同级</a>
                  <a href="" onclick="prom1('{{.Id}}')">添加下级</a>
                  <a href="/modify?id={{.Id}}">修改</a>
                  <a href="/delete?id={{.Id}}">删除</a>
                  </th>
                 </tr>

                 {{end}} 
                 {{end}}

                 {{end}} 
                  {{end}}
      </tr>
      {{end}}

    </tbody>
  </table>
</div>




<script type="text/javascript">

//实例化编辑器
//议使用工厂方法getEditor创建和引用编辑器实例，如果在某个闭包下引用该编辑器，直接调用UE.ditor('editor')就能拿到相关的实例
var ue1 = UE.getEditor('editor_cover');
//  setTimeout(function(){ uParse('#content', {
//  'highlightJsUrl':'{/static/ueditor/third-party/SyntaxHighlighter/shCore.js',
//  'highlightCssUrl':'/static/ueditor/third-party/SyntaxHighlighter/shCoreDefault.css'})
// }, 300);   
ue1.addListener("ready", function () {
uParse('.editor_cover', {
    rootPath: '/static/ueditor/'
});
});

$(function(){
        var content =$('#editor_cover').val();
        //判断ueditor 编辑器是否创建成功
        ue1.addListener("ready", function () {
        // editor准备好之后才可以使用
        ue1.setContent({{str2html .Category.Cover}});
        });
});

var ue2 = UE.getEditor('editor_property');  
ue2.addListener("ready", function () {
uParse('.editor_property', {
    rootPath: '/static/ueditor/'
});
});
$(function(){
        var editor_property =$('#content_property').val();
        //判断ueditor 编辑器是否创建成功
        ue2.addListener("ready", function () {
        // editor准备好之后才可以使用
        ue2.setContent({{str2html .Category.Content}});
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
//弹出一个输入框，输入一段文字，可以提交 
//添加同级 
    function prom(id) {  
        var name = prompt("请输入名称", ""); //将输入的内容赋给变量 name ，  
        //这里需要注意的是，prompt有两个参数，前面是提示的话，后面是当对话框出来后，在对话框里的默认值  
        if (name)//如果返回的有内容  
        {  
          var pid = $('#'+id).val();
            // alert("欢迎您：" + name) 
            $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/add",
                data: {pid:pid,title:name},
                success:function(data,status){//数据提交成功时返回数据
                  alert("添加“"+data+"”成功！(status:"+status+".)");
                 }
            });  
        }  
  
    } 
 //弹出一个输入框，输入一段文字，可以提交
 //添加下级  
    function prom1(id) {  
        var name = prompt("请输入名称", ""); //将输入的内容赋给变量 name ，  
        //这里需要注意的是，prompt有两个参数，前面是提示的话，后面是当对话框出来后，在对话框里的默认值  
        if (name)//如果返回的有内容  
        {  
          // var pid = $('#'+id).val();
            // alert("欢迎您：" + name) 
            $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/addj",
                data: {pid:id,title:name},
                success:function(data,status){//数据提交成功时返回数据
                  alert("添加“"+data+"”成功！(status:"+status+".)");
                 }
            });  
        }  
  
    } 

// 图片上传demo
jQuery(function() {
      var a = document.getElementsByName("categoryid");
      a=a[0].value;
    var $ = jQuery,
        $list = $('#fileList'),
        // 优化retina, 在retina下这个值是2
        ratio = window.devicePixelRatio || 1,
        // 缩略图大小
        thumbnailWidth = 100 * ratio,
        thumbnailHeight = 100 * ratio,
        // Web Uploader实例
        uploader;
    // 初始化Web Uploader
    uploader = WebUploader.create({
        // 自动上传。
        auto: true,
        // swf文件路径
        swf: '/static/fex-team-webuploader/dist/Uploader.swf',
        // 文件接收服务端。
        server: '/category/addcoverphoto',
        // 选择文件的按钮。可选。
        // 内部根据当前运行是创建，可能是input元素，也可能是flash.
        pick: '#filePicker',
        // 只允许选择文件，可选。
        accept: {
            title: 'Images',
            extensions: 'gif,jpg,jpeg,bmp,png',
            mimeTypes: 'image/*'
        }
    });

    // 当有文件添加进来的时候
    uploader.on( 'fileQueued', function( file ) {
    // var c=$('#container').val();
      uploader.option('formData', {
        // uid: 123
        // 'tnumber':a,
        // 'title':b,
        'categoryid':a,
        // {'tnumber':a,'title':b,'categoryid':c,'category':d,'content':e}
      });
        var $li = $(
                '<div id="' + file.id + '" class="file-item thumbnail">' +
                    '<img>' +
                    '<div class="info">' + file.name + '</div>' +
                '</div>'
                ),
            $img = $li.find('img');
        $list.append( $li );
        // 创建缩略图
        uploader.makeThumb( file, function( error, src ) {
            if ( error ) {
                $img.replaceWith('<span>不能预览</span>');
                return;
            }
            $img.attr( 'src', src );
        }, thumbnailWidth, thumbnailHeight );
    });

    // 文件上传过程中创建进度条实时显示。
    uploader.on( 'uploadProgress', function( file, percentage ) {
        var $li = $( '#'+file.id ),
            $percent = $li.find('.progress span');
        // 避免重复创建
        if ( !$percent.length ) {
            $percent = $('<p class="progress"><span></span></p>')
                    .appendTo( $li )
                    .find('span');
        }
        $percent.css( 'width', percentage * 100 + '%' );
    });

    // 文件上传成功，给item添加成功class, 用样式标记上传成功。
    uploader.on( 'uploadSuccess', function( file,response ) {
        $( '#'+file.id ).addClass('upload-state-done');
        var json = eval(response);
        alert(json.url);
        // document.getElementBy("inputId").value = json.url;//原生
        $("#route").val(json.url);

        // $.ajax({
        //         success:function(data,status){//数据提交成功时返回数据
        //           alert(data);
        //         }
        //       });
    });
    // 文件上传失败，显示上传出错。
    uploader.on( 'uploadError', function( file ) {
        var $li = $( '#'+file.id ),
            $error = $li.find('div.error');
        // 避免重复创建
        if ( !$error.length ) {
            $error = $('<div class="error"></div>').appendTo( $li );
        }
        $error.text('上传失败');
    });
    // 完成上传完了，成功或者失败，先删除进度条。
    uploader.on( 'uploadComplete', function( file ) {
        $( '#'+file.id ).find('.progress').remove();
    });
});

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