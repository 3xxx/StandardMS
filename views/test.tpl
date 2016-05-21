	<!DOCTYPE html>
	{{template "header"}}
	<html>
	<head>
	<meta charset="utf-8">
	
		<link rel="stylesheet" href="/static/css/bootstrap.min.css">
		<link rel="stylesheet" href="/static/css/bootstrap-theme.min.css">
	
	</head>

<!-- <button class="btn btn-primary btn-lg" data-toggle="modal" data-target="#myModal">
  Launch demo modal
</button> -->
<button class="btn btn-primary" onclick="update()">添加同级</button>
<a href="#myModal" >添加下级</a>
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
				<span>姓名:</span>
				<input id="name" type="text" name="personName"/>
				<br>
				<span>性别：</span>
				<input id="p_man" type="radio" name="personSex" value="男"/>
				<label for="p_man">男</label>
				<input id="p_woman" type="radio" name="personSex" value="女"/>
				<label for="p_woman">女</label>
				<br>
				<span>手机号码：</span>
				<input id="phone" type="text" maxlength="11" name="personPhone"	/>
				<span class="errorMeg" id="errorPhone"></span>
				<br>
				<span>邮箱：</span>
				<input type="text" name="personMail" id="email" onblur="checkEmail()"/>
				<span class="errorMeg" id="errorEmail"></span>
				<br>
				<span>地址：</span>
				<input type="text" name="personAddress" />
			<!-- </ul> -->

      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">确定</button>
        <!-- <button type="button" class="btn btn-primary" data-dismiss="modal">确定</button> -->
      </div>
    </div><!-- /.modal-content -->
  </div><!-- /.modal-dialog -->
</div><!-- /.modal -->

<!-- <script src="/static/js/bootstrap.min.js"></script> -->

	<script>
    function prom(id) { 	
    var pid = $('#name').val();
    var pid = $(':radio').val();//选取所有 type="button" 的 <input> 元素 和 <button> 元素
    alert("欢迎您：" + pid) 
     } 





	function update(){
	// var tds=$(obj).parent().parent().find('td');
	// $('#name').val(tds.eq(0).text());
	$('#myModal').modal('show');

   // $('#myModal').on('show.bs.modal', function () {  
   //     alert("show");  
   // });     
// $('#myModal').on('shown.bs.modal', function () {  
//        alert("shown");  
//    });     
$('#myModal').on('hide.bs.modal', function () { 
var radio =$("input[type='radio']:checked").val();	
alert("欢迎您：" + radio);
       // alert("hide");  
   });     
// $('#myModal').on('hidden.bs.modal', function () {  
//        alert("hidden");
// $(this).removeData("bs.modal");  
//    }); 


	
	}

function showCommentBoxHistory(eventid) {  
    // var id_ = eventid;  
    // if(!eventid){  
    //     var arr = new Array();  
    //     if ($("input[name^='selectIDs_history[]']:checked").length == 0) {  
    //         alert('请选择告警项!');  
    //         return;  
    //     }  
    //     $("input[name^='selectIDs_history[]']:checked").each(function(index, element) {  
    //         arr.push($(this).val());  
    //     });  
    //     id_ = arr.join(',');  
    // }  
    // $("#message").val("");  
    // $("#ack_eventid").val(id_);  
    $('#myModal').modal('show'); 
    return true;
} 

	</script>




<!-- <div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
        <h4 class="modal-title">Modal title</h4>
      </div>
      <div class="modal-body">
        <p>One fine body&hellip;</p>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
        <button type="button" class="btn btn-primary">Save changes</button>
      </div>
    </div>
  </div>
</div> -->


	<!-- <body>
	<div class="d_table" id="d_table">
	
		<table class="table table-hover table-striped table-bordered">
	
	<thead>
	<tr>
	<th>姓名</th>
	<th>性别</th>
	<th>手机号码</th>
	<th>邮箱</th>
	<th>操作</th>
	</tr>
	</thead>
	<tbody id="tbody">

	</tbody>
	</table>
	</div>
	
		<div class="modal fade" id="update" aria-hidden="true">
	
	<div class="modal-dialog">
	<div class="modal-content">
	
		<form id="saveDeviceForm" action="saveDevice" method="post">
	
	<div class="modal-header">
	
		<button type="button" class="close" data-dismiss="modal">×</button>
	
	<h3>修改</h3>
	</div>
	<div class="modal-body">
	<ul>
	
		<span>姓名:</span><input id="name" type="text" name="personName"/>
	
	<span>性别：</span>
	
		<input id="p_man" type="radio" name="personSex" value="男"/><label for="p_man">男</label>
	
	
		<input id="p_woman" type="radio" name="personSex" value="女"/><label for="p_woman">女</label>
	
	
	
		<span>手机号码：</span><input id="phone" type="text" maxlength="11" name="personPhone"/><span class="errorMeg" id="errorPhone"></span>
	
	
		<span>邮箱：</span><input type="text" name="personMail" id="email" onblur="checkEmail()"/><span class="errorMeg" id="errorEmail"></span>
	
	
		<span>地址：</span><input type="text" name="personAddress" />
	
	</ul>
	</div>
	<div class="modal-footer">
	
		<a class="button" data-dismiss="modal" aria-hidden="true">取消</a>
	
	<a class="button" onclick="submitOK()">确定</a>
	</div>
	</form>
	</div>
	</div>
	</div>

	</body> -->
	</html>
	
		
