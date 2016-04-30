<!-- 添加项目第二步视图 -->
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
<h2>第二步：添加项目封面图片、封面文字、项目简介（特性）</h2>
  <form method="post" action="/category/post2" enctype="multipart/form-data">
  <input type="hidden" id="number" name="number" value="{{.Category.Number}}"/>
  <input type="hidden" id="name" name="name" value="{{.Category.Title}}"/>
  <input type="hidden" id="categoryid" name="categoryid" value="{{.Category.Id}}"/>
  <input type="hidden" id="route" name="route" value=""/>
    <div class="form-group">
      <label>封面文字：</label>
      <div>
        <script id="editor_cover" name="editor_cover" type="text/plain" style="height:500px;width:1024px;"></script> <!-- 如果不定义那么，则默认为editorValue    --> 
      </div>

      <label>封面图片：</label>
      <!-- <div> -->
        <!-- <script id="editor_photo" name="editor_photo" type="text/plain" style="height:500px;"></script>--> <!-- width:1024px; -->  
      <!-- </div> -->
<!--dom结构部分-->
<div id="uploader-demo">
    <!--用来存放item-->
    <div id="fileList" class="uploader-list"></div>
    <div id="filePicker">选择图片</div>
</div>

    
      <label>项目简介（特性）:</label>
      <div>
        <script id="editor_property" name="editor_property" type="text/plain" style="height:500px;width:1024px;" ></script> <!--   -->
      </div>
   </div> 

<hr>

    <!-- <input type="hidden" id="tempString" name="tempString"/> -->
    <!-- <button type="submit" class="btn btn-default" onclick="return checkInput();">添加</button>
  -->            
  <!-- <input type="submit" class="btn btn-primary" value="回复" id="submit" disabled="disabled"> -->
  <button type="submit" class="btn btn-primary" > 添  加 </button>
  <!--onclick="return checkInput();" 必须加return才能不跳转-->
</form>
<br />
<br />
</div>


<script type="text/javascript">
//实例化编辑器
//议使用工厂方法getEditor创建和引用编辑器实例，如果在某个闭包下引用该编辑器，直接调用UE.ditor('editor')就能拿到相关的实例
var ue1 = UE.getEditor('editor_cover');
// var ue2 = UE.getEditor('editor_photo');没有这个，所以总是出现Cannot read property 'offsetWidth' of null
var ue3 = UE.getEditor('editor_property');
/* 1.传入函数,命令里执行该函数得到参数表,添加到已有参数表里 */
ue1.ready(function () {
// ue1.addListener('focus', function () {
     var name = $('#name').val();
     var number = $('#number').val();
    ue1.execCommand('serverparam', {
        "number":number,
        'name': name,
    });
// });
});
// ue2.ready(function () {
// // ue2.addListener('focus', function () {
//      var name = $('#name').val();
//      var number = $('#number').val();
//     ue2.execCommand('serverparam', {
//         "number":number,
//         'name': name,
//     });
// // });
// });
ue3.ready(function () {
// ue3.addListener('focus', function () {
     var name = $('#name').val();
     var number = $('#number').val();
    ue3.execCommand('serverparam', {
        "number":number,
        'name': name,
    });
// });
});


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

</body>
</html>