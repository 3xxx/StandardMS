<!DOCTYPE html>
{{template "header"}}
<title>项目&目录 - 水利设计CMS系统</title>
    <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
</head>

<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>

<!-- <div class="col-lg-12">
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
    <div id="editor_cover">
      <script id="editor_cover" type="text/plain" style="height:500px;"></script>
    </div>
    <label>封面图片</label>
      <div id="uploader-demo">
          <div id="fileList" class="uploader-list"></div>
          <div id="filePicker">选择图片</div>
       <img src="{{.Category.Route}}" width="800" align="middle">   
      </div>
      
    <label>项目简介:</label>
    <div id="editor_property">
      <script id="editor_property" type="text/plain" style="height:500px;"></script>
    </div>

    <input type="hidden" id="tempString" name="tempString"/>
    <input type="hidden" id="cid" name="cid" value="{{.Category.Id}}" />
  <button type="submit" class="btn btn-default" >修 改</button>
</form>
</div> -->
<!-- <hr> -->
   <input type="hidden" id="number" name="number" value="{{.Category.Number}}"/>
   <input type="hidden" id="name" name="name" value="{{.Category.Title}}"/>
   <input type="hidden" id="categoryid" name="categoryid" value="{{.Category.Id}}"/>

      <div class="col-lg-3">
        {{range $k,$v :=.Categoryjieduan}}
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
        {{end}}
      </div>

<div class="col-lg-9">
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
        <!-- href="javascript:void(0)" onclick="update(this);" -->
        <th>
          <!-- <input type="hidden" id="{{.Id}}" value="{{$.Category.Id}}"/> -->
          <!-- <button class="btn btn-primary" onclick="prom('{{.Id}}')">添加同级</button> -->
          <a href="" data-toggle="modal" onclick="prom('{{$.Category.Id}}')">添加同级</a><!-- data-target="#myModal" --> 
          <a href="" data-toggle="modal" onclick="prom('{{.Id}}')">添加下级</a>
          <a href="" id="{{.Id}}" value="{{.Title}}" onclick="prom1('{{.Id}}')">修改</a>
          {{if .Isshow}}
          <a href="/category/hidecategory?id={{.Id}}&cid={{$.Category.Id}}">Hide</a>
          {{else}}
          <a href="/category/showcategory?id={{.Id}}&cid={{$.Category.Id}}">Show</a>
          {{end}}
          <a href="/category/deletecategory?id={{.Id}}&cid={{$.Category.Id}}">删除</a>
<!--           <a href="javascript:void(0)">点击</a>点击链接后不会回到网页顶部
          <a href="#">点击</a>  点击后会回到网面顶部 -->
        </th>
                {{range $k1,$v1 :=$.Categoryzhuanye}}
                {{if eq $v.Id $v1.ParentId}}
                 <tr>
                   <th></th>
                 <th>{{.Title}}</th>
                 <th></th>
                 
                 <th>
                  <!-- <input type="hidden" id="{{.Id}}" value="{{$v.Id}}"/> -->
                   <a href="" data-toggle="modal" onclick="prom('{{$v.Id}}')">添加同级</a>
                   <a href="" data-toggle="modal" onclick="prom('{{.Id}}')">添加下级</a>
                  <a href="" id="{{.Id}}" value="{{.Title}}" onclick="prom1('{{.Id}}')">修改</a>
                  {{if .Isshow}}
                  <a href="/category/hidecategory?id={{.Id}}&cid={{$.Category.Id}}">Hide</a>
                  {{else}}
                  <a href="/category/showcategory?id={{.Id}}&cid={{$.Category.Id}}">Show</a>
                  {{end}}
                  <a href="/category/deletecategory?id={{.Id}}&cid={{$.Category.Id}}">删除</a>
                 </th>
                 </tr>
                        {{range $k2,$v2 :=$.Categorychengguo}}
                        {{if eq $v1.Id $v2.ParentId}}
                 <tr>
                   <th></th>
                   <th></th>
                  <th>{{.Title}}</th>                 
                  <th>
                  <!-- <input type="hidden" id="{{.Id}}" value="{{$v1.Id}}"/> -->
                  <a href="" data-toggle="modal" onclick="prom('{{$v1.Id}}')">添加同级</a>
                  <!-- <a href="" onclick="prom1('{{.Id}}')">添加下级</a> -->
                  <a href="" id="{{.Id}}" value="{{.Title}}" onclick="prom1('{{.Id}}')">修改</a>
                  {{if .Isshow}}
                  <a href="/category/hidecategory?id={{.Id}}&cid={{$.Category.Id}}">Hide</a>
                  {{else}}
                  <a href="/category/showcategory?id={{.Id}}&cid={{$.Category.Id}}">Show</a>
                  {{end}}
                  <a href="/category/deletecategory?id={{.Id}}&cid={{$.Category.Id}}">删除</a>
                  <!-- <a href="" onclick="getScrollOffset()">获取位置</a> -->
                  
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

<!-- Modal aria-hidden="true"aria-hidden="true"-->
<div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" >
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" >&times;</button>
        <h4 class="modal-title" id="myModalLabel">Modal title</h4>
      </div>
      <div class="modal-body">

      <!-- <ul> -->
        <span>输入名称</span>
        <input id="cname" type="text" name="categoryName"/>
        <br>
        
        <span>选择类别</span>
        <input  type="radio" name="Sex" checked="true" value="Attachment"/>
        <label for="p_man">附件模式</label>
        <input type="radio" name="Sex" value="Fdiary"/>
        <label for="p_woman">图文模式</label>
      <!-- </ul> -->

      </div>
      <div class="modal-footer">
        <!-- <button type="button" class="btn btn-default" data-dismiss="modal">确定</button> -->
        <!-- <button type="button" class="btn btn-primary">Save changes</button> -->
        <!-- <a class="button" data-dismiss="modal" aria-hidden="true">取消</a> -->
                <a class="button" onclick="update()">确定</a>
      </div>
    </div><!-- /.modal-content onclick="prom()"-->
  </div><!-- /.modal-dialog -->
</div><!-- /.modal -->


<script type="text/javascript">

//添加同级 
    function prom(id) {
      // $('#myModal').modal('show');
      $('#myModal').modal({
        show:true,
        backdrop:'static'
        });
      $('#myModal').on('hide.bs.modal', function () { 
        var radio =$("input[type='radio']:checked").val();  
        alert("您选择的是："+ radio + "。抱歉！添加功能暂时不提供。");
        $('#myModal').removeData("bs.modal"); // 移除数据 
         }); 
        // if (cname)
        // {  
        //   // var pid = $('#'+id).val();//父级id
        //     $.ajax({
        //         type:"post",
        //         url:"/category/post",
        //         data: {pid:id,title:cname},//父级id
        //         success:function(data,status){
        //           alert("添加“"+data+"”成功！(status:"+status+".)");
        //          }
        //     });  
        // } 
      // $('#myModal').on('hidden.bs.modal', function () { 
       // var radio =$("input[type='radio']:checked").val();  
        // alert("您选择的是：" + radio + "。抱歉！添加功能暂时不提供。");
       // alert("hidden");
       // type = undefined;
      // $(this).removeData("bs.modal"); //移除数据 
      // });  
    } 
  function update(){

alert("欢迎您：");

   } 

   //修改名称
    function prom1(id) {
      // var my=document.getElementById(id);
      // mytitle=my.title
      // $("#aa").attr("value")
      var title = $('#'+id).attr("value");
    // var title = $('#'+id).attr("href");
      var categoryid = $('#categoryid').val();
    // var a = document.getElementsByName(id);
      // a=a[0].value;
    // alert("欢迎您：" + a);
        var name = prompt("旧名称为：" + title + "。请输入新名称", ""); //将输入的内容赋给变量 name ，  
        //这里需要注意的是，prompt有两个参数，前面是提示的话，后面是当对话框出来后，在对话框里的默认值  
        if (name)//如果返回的有内容  
        {  
          // var pid = $('#'+id).val();
            // alert("欢迎您：" + name) 
            $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/category/modifycategorytitle",
                data: {pid:id,title:name,cid:categoryid},
                success:function(data,status){//数据提交成功时返回数据
                  alert("添加“"+data+"”成功！(status:"+status+".)");
                 }
            });  
        }  
    } 


 //弹出一个输入框，输入一段文字，可以提交
 //添加下级  
    // function prom1(id) {  
        // var name = prompt("请输入名称", ""); //将输入的内容赋给变量 name ，  
        //这里需要注意的是，prompt有两个参数，前面是提示的话，后面是当对话框出来后，在对话框里的默认值  
        // if (name)//如果返回的有内容  
        // {  
          // var pid = $('#'+id).val();
            // alert("欢迎您：" + name) 
            // $.ajax({
                // type:"post",//这里是否一定要用post？？？
                // url:"/category/post",
                // data: {pid:id,title:name},
                // success:function(data,status){//数据提交成功时返回数据
                  // alert("添加“"+data+"”成功！(status:"+status+".)");
                 // }
            // });  
        // }
    // }

//下面这个靠谱啊。页面刷新，保持滚动条位置。
window.onbeforeunload = function () { 
var scrollPos; 
if (typeof window.pageYOffset != 'undefined') { 
scrollPos = window.pageYOffset; 
} 
else if (typeof document.compatMode != 'undefined' && 
document.compatMode != 'BackCompat') { 
scrollPos = document.documentElement.scrollTop; 
} 
else if (typeof document.body != 'undefined') { 
scrollPos = document.body.scrollTop; 
} 
document.cookie = "scrollTop=" + scrollPos; //存储滚动条位置到cookies中 
} 

window.onload = function () { 
if (document.cookie.match(/scrollTop=([^;]+)(;|$)/) != null) { 
var arr = document.cookie.match(/scrollTop=([^;]+)(;|$)/); //cookies中不为空，则读取滚动条位置 
document.documentElement.scrollTop = parseInt(arr[1]); 
document.body.scrollTop = parseInt(arr[1]); 
} 
}


/**
 *  获取窗口滚动条的位置——这个没有用到。
 */
function getScrollOffset(){
    // 除IE8及更早版本
    if( window.pageXOffset != null ){
        return {
            x : window.pageXOffset,
            y : window.pageYOffset
        }
    }
    // 标准模式下的IE
    if( document.compatMode == "css1Compat" ){
        return {
            x : document.documentElement.scrollLeft,
            y : document.documentElement.scrollTop
        }
    }
    // 怪异模式下的浏览器
    return {
        x : document.body.scrollLeft,
        y : document.body.scrollTop,
    }   
}
</script>


</body>
</html>