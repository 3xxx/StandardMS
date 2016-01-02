<!DOCTYPE html>
{{template "header"}}
<title>{{.Topic.Title}} - 水利设计CMS系统</title>

<!-- <style type="text/css">
h4 {
  color: #DC143C;
  background: #fff;
  margin: 0;
  padding: 0;
  font-family: Georgia, Palatino, serif;
  }
</style> -->
</head>
<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>

<input type="button" id="btn_addtr" value="增行">
<div class="col-lg-12">
<table class="table table-striped">
  <thead>
      <tr>
      <!-- <th><span style="cursor: pointer">Id</span></th> -->
      <th><span style="cursor: pointer">Username</span></th>
      <th><span style="cursor: pointer">Password</span></th>
      <th><span style="cursor: pointer">Nickname</span></th>
      <th><span style="cursor: pointer">Email</span></th>
      <!-- <th><span style="cursor: pointer">Remark</span></th> -->
      <!-- <th><span style="cursor: pointer">Status</span></th> -->
      <th><span style="cursor: pointer">Lastlogintime</span></th>
      <th><span style="cursor: pointer">Createtime</span></th>
      <!-- <th><span style="cursor: pointer">RoleId</span></th> -->
      <th><span style="cursor: pointer">RoleTitle</span></th>
      <th><span style="cursor: pointer">RoleName</span></th>
      <th><span style="cursor: pointer">Remark</span></th>      
      <th>操作</th>
    </tr>
  </thead>


   {{range $index, $elem :=.Role}}
    <tr><!--tr表格的行，td定义一个单元格，<th> 标签定义表格内的表头单元格-->
      <!-- <th>{{$index}}</th> -->
      <!-- <th>{{$.User.Id}}</th> value="{{$.User.Password}}"-->
      <th>{{$.User.Username}}</th>
      <th><input type="password" id="input" name="password"  size='18'/></th>
      <th><input type="text" id="input" name="nickname" value="{{$.User.Nickname}}" size='6'/></th>
      <th><input type="text" id="input" name="email" value="{{$.User.Email}}" size='20'/></th>
      <th>{{dateformat $.User.Lastlogintime "2006-01-02 T 15:04:05"}}</th>
      <th>{{dateformat $.User.Createtime "2006-01-02 T 15:04:05"}}</th>
      <!-- <th>{{.Id}}</th> -->
      <th>{{.Title}}</th>
<!--       <th><input type="hidden" id="input" name="roletitle1" value="{{.Title}}"/></th> -->
      <th>{{.Name}}</th>
      <th>{{.Remark}}</th>
      <td><input type="button" id="btn_deltr" onclick="deltr()" value="删行"></td>
    </tr>
    {{end}}
 </table>
</div>

<script>
$(document).ready(function(){
  // var roletitle1=document.getElementsByName("roletitle");
  // $("#uname").focus(function(){获得焦点
     $("input").blur(function(){//其失去焦点
        var pwd=document.getElementsByName("password");
        var nickname=document.getElementsByName("nickname");
        var email=document.getElementsByName("email");
        // var roletitle2=document.getElementsByName("roletitle");
        // alert(pwd[0].value.length);//什么时候是逗号，什么时候是分号？
        if (pwd[0].value.length<1)
        {
          // alert("请输入密码。")
             $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/user/UpdateUser",
                data: { userid:{{.User.Id}},username:{{.User.Username}},nickname: nickname[0].value,email: email[0].value},
                success:function(data,status){//数据提交成功时返回数据
                  // alert(data,status);
                 }
            });         
        }else{
            $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/user/UpdateUser",
                data: { userid:{{.User.Id}},username:{{.User.Username}},password: pwd[0].value,nickname: nickname[0].value,email: email[0].value},
                success:function(data,status){//数据提交成功时返回数据
                  alert('success modified~');
                  // alert(data,status);
                 }
            });
       }     
 });
});



function checkInput(){
  var uname=document.getElementById("uname");
  if (uname.value.length==0){
    alert("请输入账号");
    return false;
  }
    var pwd=document.getElementById("pwd");
  if (pwd.value.length==0){
    alert("请输入密码");
    return false;
    }
// return true
}




  $(function(){
    var show_count = 20;   //要显示的条数
    var count =1
    // var count = $("input:text").val();    //$(":text") 所有 type="text" 的 <input> 元素     递增的开始值，这里是你的ID
    var fin_count = parseInt(count) + (show_count-1);   //结束递增的条件
    
      $("#btn_addtr").click(function(){
        // if(count < fin_count)    //点击时候，如果当前的数字小于递增结束的条件
        // {
        $("tr:eq(1)").clone().appendTo("table");   //在表格后面添加一行
        $("tr:last td input:first").val(++count);   //改变添加的行的ID值。
        // }
      });
    });
    function deltr(){
    var length=$("tr").length;
      if(length<=2){
      alert("至少保留一行");
      }else{
      $("tr:last").remove();
      }
    }
</script>

<script type="text/javascript">
  $(document).ready(function() {
  $("table").tablesorter();
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
});
</script>
</body>
</html>
<!--     Id       int64  `PK`
    Username string `orm:"unique"`
    Password string
    // Id            int64
    // Username      string    `orm:"unique;size(32)" form:"Username"  valid:"Required;MaxSize(20);MinSize(6)"`
    // Password      string    `orm:"size(32)" form:"Password" valid:"Required;MaxSize(20);MinSize(6)"`
    Repassword    string    `orm:"-" form:"Repassword" valid:"Required"`
    Nickname      string    //`orm:"unique;size(32)" form:"Nickname" valid:"Required;MaxSize(20);MinSize(2)"`
    Email         string    `orm:"size(32)" form:"Email" valid:"Email"`
    Remark        string    `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
    Status        int       `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
    Lastlogintime time.Time `orm:"null;type(datetime)" form:"-"`
    Createtime    time.Time `orm:"type(datetime);auto_now_add" `
    Role          []*Role   `orm:"rel(m2m)"` -->