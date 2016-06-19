<!DOCTYPE html>
{{template "header"}}
<title>项目&成果 - 水利设计CMS系统</title>
<style type="text/css">
/*body{
    margin:0;
    padding:0;
    font:12px/1.5 tahoma,arial,'Hiragino Sans GB',\5b8b\4f53,sans-serif;
    position:absolute;width:100%;min-height:100%;   
}
 .content{
    padding-bottom: 54px;
 }
 
a:link,a:visited,a:active{color:#00749E;text-decoration:none;}
a:hover{color:#000;text-decoration:underline;}
 
#header{width:980px;height:69px;margin:0 auto;padding:0;}
#header h1.blogtitle{font-family:Arial, Helvetica, sans-serif;margin:0;padding:10px 0 0;}
#header p.desc{font-family:Verdana, Arial, Helvetica, sans-serif;font-size:11px;margin:0;padding:0;}
#top{background-color:#fff;background-image:url(images/topbck.gif);background-repeat:repeat-x;height:34px;}
 
#ddnav{background-color:#312e2e;height:35px;margin:0;padding:0;}
#nav{height:25px;width:980px;font-weight:700;margin:0 auto;padding:10px 0 0;}
ul.nav{height:25px;line-height:25px;float:left;list-style:none;font-size:11px;text-transform:inherit;margin:0;padding:0;}
ul.nav li{float:left !important;list-style-type:none;border-right:2px solid #312e2e;margin:0;padding:0;}
ul.nav li a,ul.nav li a:link,ul.nav li a:visited{background:url(images/navsilver.gif) repeat-x;color:#818181;float:left;display:block;text-decoration:none;padding:0 15px;}
ul.nav li a:hover,ul.nav li a:active {background:url(images/navblue.gif) repeat-x;color:#205387;text-decoration:none;}
ul.nav li.current_page_item a{text-decoration:none;background:url(images/navblue.gif) repeat-x;color:#fff;}
ul.nav li ul{float:left;margin:0;padding:0;}*/

/*#footer{background-image:url(images/footerbck.gif);background-repeat:repeat-x;clear:both;height:24px;margin:0;padding:0;position:absolute;bottom:30px;width:100%;}
#footerbox{color:#fff;background-color:#312e2e;height:30px;line-height:30px !important;margin:0;padding:0;clear:both;position:absolute;bottom:0px;width:100%;}
.footer{width:980px;color:#fff;height:30px;line-height:30px !important;margin:0 auto;padding:0;}*/
/*.footer a:hover{text-decoration:underline;color:#fff;border:none;}*/
/*.footer a:link,.footer a:active,.footer a:visited{text-decoration:underline;color:#25aacd;border:none;}*/

/*#top{background-color:#fff;background-image:url(images/topbck.gif);background-repeat:repeat-x;height:34px;}*/

/*让footer固定在页面底部
#footer{clear:both;height:150px;margin:0;padding:0;position:absolute;bottom:30px;width:100%;}
#footerbox{margin:0;padding:0;clear:both;position:absolute;bottom:0px;width:100%;}
.footer{width:980px;line-height:30px !important;margin:0 auto;padding:0;}*/

#button{
  width:100px;
  }
</style>
</head>

<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>

<div class="text-center">
  <h1 > <i class="glyphicon glyphicon-chevron-right"></i> <i class="glyphicon glyphicon-minus"></i>
  </h1>
  <h1 >搜索888个 文件</h1>
  <p class="large">
    HydroCMS 是一个微服务系统，您可以在自己电脑上运行HydroCMS，像本站这样发布资料，还可以将资料打包共享。
  </p>
  <p class="large">在任何一个 HydroCMS 上可搜索到局域网上所有资源，最终组成一个交织的文档体系。</p>

  <div class="col-lg-4">
</div>
  
  <div class="col-lg-4">
  <!-- <form >   form支持回车，但是不支持json，如何做到支持json？-->
    <div class="input-group">
      <input type="text" class="form-control" placeholder="请输入关键字进行搜索" name="name" autocomplete="off" size="30" id="name" onkeypress="getKey();">
      <span class="input-group-btn">
        <button class="btn btn-default" type="button" id="search"><!-- type="submit" -->
          <i class="glyphicon glyphicon-search"></i>
          Search!
        </button>
      </span>

    </div>
         <span>选择搜索范围：</span>
        <input  type="radio" name="range" checked="true" value="local"/>
        <label >本机</label>
        <input type="radio" name="range" value="global"/>
        <label >全局</label> 
   <!-- </form>  -->
  </div>

  <div class="col-lg-12">
        <h1 class="ui icon header">
          <i class="browser icon"></i>
          查询结果
        </h1>
        <table class="table table-striped">
          <thead>
            <tr>
              <th >编号</th>
              <th >名称</th>
              <th>链接</th>
              <th>分类</th>
              <th>上传者</th>
            </tr>
          </thead>
          <tbody id="results">
          </tbody>
        </table>
  </div>


<div class="col-lg-12"> 
<br>
<hr/>
</div> 

<div class="col-lg-12">
<h3>炫耀主机</h3>
{{range $index, $elem :=.Service}}
 <a href="{{.Link}}" class="btn btn-primary" id="button"><i class="glyphicon glyphicon-flash"></i>{{.Name}}<i class="glyphicon glyphicon-screenshot"></i></a> 
{{end}}
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
<!-- <div id="footerbox">这个也行。
    <div class="footer">
        <a href="http://www.cnblogs.com/chenyuming507950417/">助你软件工作室</a>&nbsp;&nbsp;&nbsp; 二十八年，我用青春修行&nbsp;&nbsp;&nbsp; &copy; Copyright
        2014
    </div>
</div> -->


<!-- <div class="col-lg-6">
  <h3>项目列表</h3>
      <table class="table table-striped">
       <thead>
         <tr>
           <th style="cursor: pointer">#{{(.SpiderCategory|len)}}</th>
           <th style="cursor: pointer">项目编号</th>
           <th style="cursor: pointer">项目名称</th>
           <th style="cursor: pointer">主机名称</th>
         </tr>
       </thead>

       <tbody>
        {{range $index, $elem :=.SpiderCategory}}
         <tr>
          <th>{{$index}}</th>
          <th><a href="{{.Link}}" id="number">{{.Number}}</a></th>
          <th><a href="{{.Link}}" id="name"><i class="glyphicon glyphicon-plane"></i>{{.Name}}</a></th>
          <th><a href="{{.UserIp}}" id="usename">{{.UserName}}</a></th>
         </tr>
       {{end}}
     </tbody>
     </table>
  </div>

<div class="col-lg-6">
  <h3>成果列表</h3>
      <table class="table table-striped">
       <thead>
         <tr>
           <th style="cursor: pointer">#{{(.SpiderTopic|len)}}</th>
           <th style="cursor: pointer">成果编号</th>
           <th style="cursor: pointer">成果名称</th>
           <th style="cursor: pointer">主机名称</th>
         </tr>
       </thead>

       <tbody>
        {{range $index, $elem :=.SpiderTopic}}
         <tr>
          <th>{{$index}}</th>
          <th><a href="{{.Link}}" id="number">{{.Number}}</a></th>
          <th><a href="{{.Link}}" id="name"><i class="glyphicon glyphicon-fire"></i>{{.Name}}</a></th>
          <th><a href="{{.UserIp}}" id="usename">{{.UserName}}</a></th>
         </tr>
       {{end}}
     </tbody>
     </table>
  </div> -->

<script type="text/javascript">
    // <form class="navbar-form navbar-left" role="search" method="get" action="/search">
    //   <div class="form-group">
    //     <input id="tuming" type="text" class="form-control"  class="search-query span2" placeholder="Search local" name="tuming"></div>
    //   <button type="submit" class="btn btn-default">Submit</button>
    // </form>

$(document).ready(function(){
$("#search").click(function(){//这里应该用button的id来区分按钮的哪一个,因为本页有好几个button
                $.ajax({
                type:"post",//这里是否一定要用post，是的，因为get会缓存？？
                url:"/searchlocal",
                data: {name: $("#name").val()},
                success:function(data,status){//数据提交成功时返回数据
                  // alert(data);
                  // alert(data[1].Uname);
                  $.each(data,function(i,d){
                    var tr=$("<tr></tr>");
                    var th1=$('<th>' + data[i].Tnumber + '</th>');
                    var th2=$('<th>' + data[i].Title + '</th>');
                    var th3=$('<th><a href=/topic/view_b/' + data[i].Id + ' target="_black"><i                    class="glyphicon glyphicon-download-alt"></i>下载</a></th>');
                    var th4=$('<th>' + data[i].Category + '</th>');
                    var th5=$('<th>' + data[i].Author + '</th>');
                    tr.append(th1);
                    tr.append(th2);
                    tr.append(th3);
                    tr.append(th4);
                    tr.append(th5);
                    $("#results").append(tr);
                    }); 
                  }       
            });
 });
});
function getKey()  
{  
    if(event.keyCode==13){  
     // alert('click enter'); 
      $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/searchlocal",
                data: {name: $("#name").val()},
                success:function(data,status){//数据提交成功时返回数据
                  // alert(data);
                  $.each(data,function(i,d){
                    var tr=$("<tr></tr>");
                    var th1=$('<th>' + data[i].Tnumber + '</th>');
                    var th2=$('<th>' + data[i].Title + '</th>');
                    var th3=$('<th><a href=/topic/view_b/' + data[i].Id + ' target="_black"><i                    class="glyphicon glyphicon-download-alt"></i>下载</a></th>');
                    var th4=$('<th>' + data[i].Category + '</th>');
                    var th5=$('<th>' + data[i].Author + '</th>');
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

  // $(document).ready(function() {
  // $("table").tablesorter();
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
// });
</script>
</body>
</html>