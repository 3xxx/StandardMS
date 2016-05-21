 <!DOCTYPE html>
{{template "header"}}
<title>查阅规范和图集 - HydroCMS</title>
<link rel="stylesheet" type="text/css" href="/static/fex-team-webuploader/css/webuploader.css">
<script type="text/javascript" src="/static/fex-team-webuploader/dist/webuploader.js"></script>
<style type="text/css">
#footer{clear:both;height:150px;margin:0;padding:0;position:absolute;bottom:30px;width:100%;}
#footerbox{margin:0;padding:0;clear:both;position:absolute;bottom:0px;width:100%;}
.footer{width:980px;line-height:30px !important;margin:0 auto;padding:0;}
</style>
</head>

<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>

<!-- <div style="margin:0 auto;width:300px;"> 
<input type="text" class="form-control" placeholder="请输入关键字或编号进行搜索" name="standard" autocomplete="off" size="50"> 
</div> -->

<!--  <div class="col-lg-12">
  <div class="form-group">
  <form method="post" action="/standard/importexcel" enctype="multipart/form-data">
      <label>
        选择数据文件：
        <input type="file" name="excel" id="excel" />
      </label>
    <button type="submit" class="btn btn-default" >导入数据</button>
  </form>
</div>
</div> -->

<div class="text-center">
  <h1 > <i class="glyphicon glyphicon-chevron-right"></i> <i class="glyphicon glyphicon-minus"></i>
  </h1>
  <h1 >搜索{{.Length}}个 文件</h1>
  <p class="large">
    HydroCMS 是一个微服务系统，你可以上传你的资料，也可以在自己电脑上运行HydroCMS，像本站这样发布资料，还可以将资料打包共享。
  </p>
  <p class="large">HydroCMS 与院《规范目录版本管理》数据进行比对，标记库中录入的规范是否有效。</p>

  <div class="col-lg-4">
</div>
  
  <div class="col-lg-4">
  <!-- <form >   form支持回车，但是不支持json，如何做到支持json？-->
    <div class="input-group">
      <input type="text" class="form-control" placeholder="请输入关键字或编号进行搜索" name="name" autocomplete="off" size="30" id="name" onkeypress="getKey();">
      <span class="input-group-btn">
        <button class="btn btn-default" type="button" id="search"><!-- type="submit" -->
          <i class="glyphicon glyphicon-search"></i>
          Search!
        </button>
      </span>
    </div>
   <!-- </form>  -->
  </div>

    <!-- <div class="results" id="results"></div> -->

  <div class="col-lg-12">
        <h1 class="ui icon header">
          <i class="browser icon"></i>
          查询结果
        </h1>
        <table class="table table-striped"><!-- class="ui basic table" -->
          <thead>
            <tr>
              <th >编号</th>
              <th >名称</th>
              <th>链接</th>
              <th>上传者</th>
              <th>有效版本库</th>
            </tr>
          </thead>
          <tbody id="results">

          </tbody>
        </table>
  </div>

 <br>
 <div class="col-lg-12">
 <!-- <div class="col-lg-6"> -->
<!--SWF在初始化的时候指定，在后面将展示-->
<div id="uploader" class="wu-example">
    <!--用来存放文件信息-->
    <div id="thelist" class="uploader-list"></div>
    <!-- <div class="btns"> -->
        <div id="picker"><i class="glyphicon glyphicon-plus-sign"></i></div>
        <button id="ctlBtn" class="btn btn-default"><i class="glyphicon glyphicon-upload"></i></button>
    <!-- </div> -->
</div>
<!-- </div> -->
<br>
 <!-- <div class="col-lg-6"> -->
  <div class="form-group">
  <form method="post" action="/standard/importexcel" enctype="multipart/form-data">
      <label>
        <input type="file" name="excel" id="excel" class="btn btn-default"/>
      </label>
    <button type="submit" class="btn btn-default" >导入数据</button>
  </form>
  </div>
  <div class="form-group">
  <form method="post" action="/standard/importlibrary" enctype="multipart/form-data">
      <label>
        <input type="file" name="excel" id="excel" class="btn btn-default"/>
      </label>
    <button type="submit" class="btn btn-default" >导入有效库</button>
  </form>
  </div>
<!-- </div> -->
</div>

<div id="footer">
  <div class="col-lg-12">
    <br>
    <hr/>
  </div>

  <div class="col-lg-6">
    <h4>Copyright © 2016 HydroCMS</h4>
    <p>
      网站由 <i class="user icon"></i>
      <a target="_blank" href="https://github.com/3xxx">@3xxx</a>
      建设，并由
      <a target="_blank" href="http://golang.org">golang</a>
      和
      <a target="_blank" href="http://beego.me">beego</a>
      提供动力。
    </p>

    <p>
      请给 <i class="glyphicon glyphicon-envelope"></i>
      <a class="email" href="mailto:qin.xc@gpdiwe.com">我们</a>
      发送反馈信息或提交
      <i class="tasks icon"></i>
      <a target="_blank" href="https://github.com/3xxx/hydrocms/issues">网站问题</a>
      。
    </p>
  </div>
  <div class="col-lg-6">
    <h4 >更多项目</h4>
    <div >
      <p>
        <a href="https://github.com/3xxx/pipeline">管线分段设计工具</a>
      </p>
      <p>
        <a href="https://github.com/3xxx/merit">技术人员价值管理系统</a>
      </p>
    </div>
  </div>
</div>
</div>

<script>
// <button class="btn btn-primary" id="export">导出excel</button>
$(document).ready(function(){
$("#search").click(function(){//这里应该用button的id来区分按钮的哪一个,因为本页有好几个button
                $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/standard/search",
                data: {name: $("#name").val()},
                success:function(data,status){//数据提交成功时返回数据
                  // alert(data);
                  // alert(data[1].Uname);
                  $.each(data,function(i,d){
                    var tr=$("<tr></tr>");
                    var th1=$('<th>' + data[i].Number + '</th>');
                    var th2=$('<th>' + data[i].Title + '</th>');
                    var th3=$('<th><a href="' + data[i].Route + '"  target="_black"><i                    class="glyphicon glyphicon-download-alt"></i>下载</a></th>');
                    var th4=$('<th>' + data[i].Uname + '</th>');
                    var th5=$('<th>' + data[i].LiNumber + data[i].LibraryTitle + '</th>');
                    tr.append(th1);
                    tr.append(th2);
                    tr.append(th3);
                    tr.append(th4);
                    tr.append(th5);
                    $("#results").append(tr);

                  // $("#results").append('<tr>'); 
                  // $("#results").append('<th>' + data[i].Number + '</th>');
                  // $("#results").append('<th>' + data[i].Title + '</th>');
                  // $("#results").append('<th><a href="' + data[i].Route + '"  target="_black"><i class="glyphicon glyphicon-download-alt"></i>下载</a></th>');
                  // $("#results").append('<th>' + data[i].Uname + '</th>');
                  // $("#results").append('<th>' + data[i].LiNumber + data[i].LibraryTitle + '</th>');
                  // $("#results").append('</tr>');
                    }); 
                    // alert("成功！"+data[0].Title); 
                  }       
            });

// function addRow(){
//     //声明tr td对象
//     var tr=$("<tr></tr>");
//     var td1=$("<td></td>");//拼接td，如果有td里有内容拼接时候填充进去
//     var td2=$("<td></td>");//拼接td，如果有td里有内容拼接时候填充进去
//     var td3=$("<td></td>");//拼接td，如果有td里有内容拼接时候填充进去
//     var td4=$("<td></td>");//拼接td，如果有td里有内容拼接时候填充进去
//     //将td添加到tr里
//     tr.append(td1);
//     tr.append(td2);
//     tr.append(td3);
//     tr.append(td4);
//     //或者
//     var tr="<tr><td></td> <td></td> <td></td> <td></td></tr>";//拼接html，如果有td里有内容拼接时候填充进去
    
//     //再将tr添加到表格内
//     $("#表格的id").append(tr);
// }
 });
});



// 文件上传
jQuery(function() {
    var $ = jQuery,
        $list = $('#thelist'),
        $btn = $('#ctlBtn'),
        state = 'pending',
        uploader;
        uploader = WebUploader.create({
        // 不压缩image
        resize: false,
        // swf文件路径
        swf: '/static/fex-team-webuploader/dist/Uploader.swf',
        // 文件接收服务端。
        server: '/standard/standard_one_addbaidu',
        // 选择文件的按钮。可选。
        // 内部根据当前运行是创建，可能是input元素，也可能是flash.
        pick: '#picker'
    });

    // 当有文件添加进来的时候
    uploader.on( 'fileQueued', function( file ) {
        $list.append( '<div id="' + file.id + '" class="item">' +
            '<h4 class="info">' + file.name + '</h4>' +
            '<p class="state">等待上传...</p>' +
        '</div>' );
    });
//当某个文件的分块在发送前触发，主要用来询问是否要添加附带参数，
//大文件在开起分片上传的前提下此事件可能会触发多次。
// uploader.on( 'fileQueued', function( file ) {
    // do some things.
// });

 // uploader.on( 'startUpload', function() {
 //    var tnumber = $('#tnumber').val();
 //    var title = $('#title').val();
 //    var categoryid = $('#categoryid').val();
 //    var category = $('#category').val();
 //    var html = ue.getContent();
 //      uploader.option('formData', {
 //        "tnumber":tnumber,
 //        "title":title,
 //        "categoryid":categoryid,
 //        "category":category,
 //        'content':html,
 //      });        
 //    });

    // 文件上传过程中创建进度条实时显示。
    uploader.on( 'uploadProgress', function( file, percentage ) {
        var $li = $( '#'+file.id ),
            $percent = $li.find('.progress .progress-bar');
        // 避免重复创建
        if ( !$percent.length ) {
            $percent = $('<div class="progress progress-striped active">' +
              '<div class="progress-bar" role="progressbar" style="width: 0%">' +
              '</div>' +
            '</div>').appendTo( $li ).find('.progress-bar');
        }

        $li.find('p.state').text('上传中');

        $percent.css( 'width', percentage * 100 + '%' );
    });

    uploader.on( 'uploadSuccess', function( file ) {
        $( '#'+file.id ).find('p.state').text('已上传');
    });

    uploader.on( 'uploadError', function( file ) {
        $( '#'+file.id ).find('p.state').text('上传出错');
    });

    uploader.on( 'uploadComplete', function( file ) {
        $( '#'+file.id ).find('.progress').fadeOut();
    });

    uploader.on( 'all', function( type ) {
        if ( type === 'startUpload' ) {
            state = 'uploading';
        } else if ( type === 'stopUpload' ) {
            state = 'paused';
        } else if ( type === 'uploadFinished' ) {
            state = 'done';
        }

        if ( state === 'uploading' ) {
            $btn.text('暂停上传');
        } else {
            $btn.text('开始上传');
        }
    });

    $btn.on( 'click', function() {
        if ( state === 'uploading' ) {
            uploader.stop();
        } else {
            uploader.upload();
        }
    });
});



function getKey()  
{  
    if(event.keyCode==13){  
     // alert('click enter'); 
      $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/standard/search",
                data: {name: $("#name").val()},
                success:function(data,status){//数据提交成功时返回数据
                  // alert(data);
                  $.each(data,function(i,d){
                    var tr=$("<tr></tr>");
                    var th1=$('<th>' + data[i].Number + '</th>');
                    var th2=$('<th>' + data[i].Title + '</th>');
                    var th3=$('<th><a href="' + data[i].Route + '"  target="_black"><i                    class="glyphicon glyphicon-download-alt"></i>下载</a></th>');
                    var th4=$('<th>' + data[i].Uname + '</th>');
                    var th5=$('<th>' + data[i].LiNumber + data[i].LibraryTitle + '</th>');
                    tr.append(th1);
                    tr.append(th2);
                    tr.append(th3);
                    tr.append(th4);
                    tr.append(th5);
                    $("#results").append(tr);
                  // <a href="/topic/view_b/{{.Id}}"><i class="glyphicon glyphicon-download-alt"></i>下载</a>
                  // $("#results").append("<li>"+data[i].Title+"</li>");
                            }); 
                    // alert("成功！"+data[0].Title); 
                  }       
            });
    }     
} 
  </script>



</body>
</html> 