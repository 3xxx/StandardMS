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

<!-- <script type="text/javascript">
// var statuslist = [
//     {statusid:'1',name:'禁用'},
//     {statusid:'2',name:'启用'}
// ];
var URL="/user";
$(function(){
    //用户列表
    $("#datagrid").datagrid({
        title:'用户列表',
        url:URL+'/index',
        method:'POST',
        pagination:true,
        fitColumns:true,
        striped:true,
        rownumbers:true,
        singleSelect:true,
        idField:'Id',
        pagination:true,
        pageSize:20,
        pageList:[10,20,30,50,100],
        columns:[[
            {field:'Id',title:'ID',width:50,sortable:true},
            {field:'Username',title:'用户名',width:100,sortable:true},
            {field:'Nickname',title:'昵称',width:100,align:'center',editor:'text'},
            {field:'Email',title:'Email',width:100,align:'center',editor:'text'},
            {field:'Remark',title:'备注',width:150,align:'center',editor:'text'},
            {field:'Lastlogintime',title:'上次登录时间',width:100,align:'center',
                formatter:function(value,row,index){
                    if(value) return phpjs.date("Y-m-d H:i:s",phpjs.strtotime(value));
                    return value;
                }
            },
            {field:'Createtime',title:'添加时间',width:100,align:'center',
                formatter:function(value,row,index){
                    if(value) return phpjs.date("Y-m-d H:i:s",phpjs.strtotime(value));
                    return value;
                }
            },
            {field:'Status',title:'状态',width:50,align:'center',
                formatter:function(value){
                    for(var i=0; i<statuslist.length; i++){
                        if (statuslist[i].statusid == value) return statuslist[i].name;
                    }
                    return value;
                },
                editor:{
                    type:'combobox',
                    options:{
                        valueField:'statusid',
                        textField:'name',
                        data:statuslist,
                        required:true
                    }
                }
            }
        ]],
        onAfterEdit:function(index, data, changes){
            if(vac.isEmpty(changes)){
                return;
            }
            changes.Id = data.Id;
            vac.ajax(URL+'/UpdateUser', changes, 'POST', function(r){
                if(!r.status){
                    vac.alert(r.info);
                }else{
                    $("#datagrid").datagrid("reload");
                }
            })
        },
        onDblClickRow:function(index,row){
            editrow();
        },
        onRowContextMenu:function(e, index, row){
            e.preventDefault();
            $(this).datagrid("selectRow",index);
            $('#mm').menu('show',{
                left: e.clientX,
                top: e.clientY
            });
        },
        onHeaderContextMenu:function(e, field){
            e.preventDefault();
            $('#mm1').menu('show',{
                left: e.clientX,
                top: e.clientY
            });
        }
    });
    //创建添加用户窗口
    $("#dialog").dialog({
        modal:true,
        resizable:true,
        top:150,
        closed:true,
        buttons:[{
            text:'保存',
            iconCls:'icon-save',
            handler:function(){
                $("#form1").form('submit',{
                    url:URL+'/AddUser',
                    onSubmit:function(){
                        return $("#form1").form('validate');
                    },
                    success:function(r){
                        var r = $.parseJSON( r );
                        if(r.status){
                            $("#dialog").dialog("close");
                            $("#datagrid").datagrid('reload');
                        }else{
                            vac.alert(r.info);
                        }
                    }
                });
            }
        },{
            text:'取消',
            iconCls:'icon-cancel',
            handler:function(){
                $("#dialog").dialog("close");
            }
        }]
    });
    //创建修改密码窗口
    $("#dialog2").dialog({
        modal:true,
        resizable:true,
        top:150,
        closed:true,
        buttons:[{
            text:'保存',
            iconCls:'icon-save',
            handler:function(){
                var selectedRow = $("#datagrid").datagrid('getSelected');
                var password = $('#password').val();
                vac.ajax(URL+'/UpdateUser', {Id:selectedRow.Id,Password:password}, 'post', function(r){
                    if(r.status){
                        $("#dialog2").dialog("close");
                    }else{
                        vac.alert(r.info);
                    }
                })
            }
        },{
            text:'取消',
            iconCls:'icon-cancel',
            handler:function(){
                $("#dialog2").dialog("close");
            }
        }]
    });

})

function editrow(){
    if(!$("#datagrid").datagrid("getSelected")){
        vac.alert("请选择要编辑的行");
        return;
    }
    $('#datagrid').datagrid('beginEdit', vac.getindex("datagrid"));
}
function saverow(index){
    if(!$("#datagrid").datagrid("getSelected")){
        vac.alert("请选择要保存的行");
        return;
    }
    $('#datagrid').datagrid('endEdit', vac.getindex("datagrid"));
}
//取消
function cancelrow(){
    if(! $("#datagrid").datagrid("getSelected")){
        vac.alert("请选择要取消的行");
        return;
    }
    $("#datagrid").datagrid("cancelEdit",vac.getindex("datagrid"));
}
//刷新
function reloadrow(){
    $("#datagrid").datagrid("reload");
}

//添加用户弹窗
function addrow(){
    $("#dialog").dialog('open');
    $("#form1").form('clear');
}

//编辑用户密码
function updateuserpassword(){
    var dg = $("#datagrid")
    var selectedRow = dg.datagrid('getSelected');
    if(selectedRow == null){
        vac.alert("请选择用户");
        return;
    }
    $("#dialog2").dialog('open');
    $("form2").form('load',{password:''});
}

//删除
function delrow(){
    $.messager.confirm('Confirm','你确定要删除?',function(r){
        if (r){
            var row = $("#datagrid").datagrid("getSelected");
            if(!row){
                vac.alert("请选择要删除的行");
                return;
            }
            vac.ajax(URL+'/DelUser', {Id:row.Id}, 'POST', function(r){
                if(r.status){
                    $("#datagrid").datagrid('reload');
                }else{
                    vac.alert(r.info);
                }
            })
        }
    });
}
</script> -->
<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>
<div class="col-lg-12">
<table id="datagrid" toolbar="#tb"></table>
<!-- <div id="tb" style="padding:5px;height:auto">
    <a href="#" icon='icon-add' plain="true" onclick="addrow()" class="easyui-linkbutton" >新增</a>
    <a href="#" icon='icon-edit' plain="true" onclick="editrow()" class="easyui-linkbutton" >编辑</a>
    <a href="#" icon='icon-save' plain="true" onclick="saverow()" class="easyui-linkbutton" >保存</a>
    <a href="#" icon='icon-cancel' plain="true" onclick="delrow()" class="easyui-linkbutton" >删除</a>
    <a href="#" icon='icon-reload' plain="true" onclick="reloadrow()" class="easyui-linkbutton" >刷新</a>
    <a href="#" icon='icon-edit' plain="true" onclick="updateuserpassword()" class="easyui-linkbutton" >修改用户密码</a>
</div> -->
<!--表格内的右键菜单-->
<!-- <div id="mm" class="easyui-menu" style="width:120px;display: none" >
    <div iconCls='icon-add' onclick="addrow()">新增</div>
    <div iconCls="icon-edit" onclick="editrow()">编辑</div>
    <div iconCls='icon-save' onclick="saverow()">保存</div>
    <div iconCls='icon-cancel' onclick="cancelrow()">取消</div>
    <div iconCls='icon-edit' onclick="updateuserpassword()">修改密码</div>
    <div class="menu-sep"></div>
    <div iconCls='icon-cancel' onclick="delrow()">删除</div>
    <div iconCls='icon-reload' onclick="reloadrow()">刷新</div>
    <div class="menu-sep"></div>
    <div>Exit</div>
</div> -->
<!--表头的右键菜单-->
<!-- <div id="mm1" class="easyui-menu" style="width:120px;display: none"  >
    <div icon='icon-add' onclick="addrow()">新增</div>
</div>
<div id="dialog" title="添加用户" style="width:400px;height:400px;">
    <div style="padding:20px 20px 40px 80px;" >
        <form id="form1" method="post">
            <table>
                <tr>
                    <td>用户名：</td>
                    <td><input name="Username" class="easyui-validatebox" required="true"/></td>
                </tr>
                <tr>
                    <td>昵称：</td>
                    <td><input name="Nickname" class="easyui-validatebox" required="true"  /></td>
                </tr>
                <tr>
                    <td>密码：</td>
                    <td><input name="Password" type="password" class="easyui-validatebox" required="true"   validType="password[4,20]" /></td>
                </tr>
                <tr>
                    <td>重复密码：</td>
                    <td><input name="Repassword" type="password" class="easyui-validatebox" required="true"   validType="password[4,20]" /></td>
                </tr>
                <tr>
                    <td>Email：</td>
                    <td><input name="Email" class="easyui-validatebox" validType="email"  /></td>
                </tr>
                <tr>
                    <td>状态：</td>
                    <td>
                        <select name="Status"  style="width:153px;" class="easyui-combobox " editable="false" required="true"  >
                            <option value="2" >启用</option>
                            <option value="1">禁用</option>
                        </select>
                    </td>
                </tr>
                <tr>
                    <td>备注：</td>
                    <td><textarea name="Remark" class="easyui-validatebox" validType="length[0,200]"></textarea></td>
                </tr>
            </table>
        </form>
    </div>
</div> -->
<!-- <div id="dialog2" title="修改用户密码" style="width:400px;height:200px;">
    <div style="padding:20px 20px 40px 80px;" >
        <table>
            <tr>
                <td>密码：</td>
                <td><input name="Password" type="password" id="password" class="easyui-validatebox" required="true"   validType="password[4,20]" /></td>
            </tr>
        </table>
    </div>
</div> -->

<form method="post" action="/user/importexcel" enctype="multipart/form-data"> 
   <div class="input-group">
      <label>选择excel：<input type="file" name="excel" id="excel" /></label>
      <br/>
   </div>
   <button type="submit" class="btn btn-default" >提交</button>
</form>

<table class="table table-striped">
  <thead>
      <tr>
      <th>#{{.Users|len}}</th>
      <th><span style="cursor: pointer">Id</span></th>
      <th><span style="cursor: pointer">Username</span></th>
      <th><span style="cursor: pointer">Password</span></th>
      <th><span style="cursor: pointer">Nickname</span></th>
      <th><span style="cursor: pointer">Email</span></th>
      <th><span style="cursor: pointer">Remark</span></th>
      <th><span style="cursor: pointer">Status</span></th>
      <th><span style="cursor: pointer">Lastlogintime</span></th>
      <th><span style="cursor: pointer">Createtime</span></th>
      <th><span style="cursor: pointer">Role</span></th>
      <th>操作</th>
    </tr>
  </thead>
   {{range $index, $elem :=.Users}}
    <tr><!--tr表格的行，td定义一个单元格，<th> 标签定义表格内的表头单元格-->
      <th>{{$index}}</th>
      <th><a href="/user/view?useid={{.Id}}">{{.Id}}</a></th>
      <th><a href="/user/view?useid={{.Id}}"><i class="glyphicon glyphicon-user"></i>{{.Username}}</a></th>
      <th><a href="/user/view?useid={{.Id}}" title={{.Title}}>{{substr .Password 0 15}}</a></th>
      <th>{{.Nickname}}</th>
      <th>{{.Email}}</th>
      <th>{{.Remark}}</th>
      <th>{{.Status}}</th>
      <th>{{dateformat .Lastlogintime "2006-01-02"}}</th>
      <th>{{dateformat .Createtime "2006-01-02"}}</th>
      <th>{{.Role}}</th>

      <th>
      <a href="/user/deluser?userid={{.Id}}">删除</a></th>
    </tr>
    {{end}}
 </table>
</div>

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