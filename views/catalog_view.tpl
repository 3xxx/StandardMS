<!-- 显示一个专业下的目录 -->
<!DOCTYPE html>
{{template "header"}}
<title>项目&目录 - 水利设计CMS系统</title>

</head>

<body>
<div class="navbar navba-default navbar-fixed-top">
	<div class="container-fill">{{template "navbar" .}}</div>
</div>

<div class="col-lg-12">
	<h1>目录</h1>
<!-- 	<form method="post" action="/catalog/import_xls_catalog" enctype="multipart/form-data">
		<div class="input-group">
			<label>选择excel：<input type="file" class="form-control" name="excel" id="excel" /></label>
			<br/>
		</div>
		<button type="submit" class="btn btn-default" >提交</button>
	</form> -->

	<table class="table table-striped" id="orderTable" name="orderTable">
		<thead>
			<tr>
				<th>
					<span style="cursor: pointer">#{{.Length}}</span>
				</th>
				<th>
					<span style="cursor: pointer">成果编号</span>
				</th>
				<th>
					<span style="cursor: pointer">成果名称</span>
				</th>
				<th>
					<span style="cursor: pointer">制图</span>
				</th>
				<th>
					<span style="cursor: pointer">设计</span>
				</th>
				<th>
					<span style="cursor: pointer">校核</span>
				</th>
				<th>
					<span style="cursor: pointer">审查</span>
				</th>
				<th>
					<span style="cursor: pointer">核定</span>
				</th>
				<th>
					<span style="cursor: pointer">批准</span>
				</th>
				<th>
					<span style="cursor: pointer">出版日期</span>
				</th>
				<th>
					<span style="cursor: pointer">阶段</span>
				</th>
				<th>
					<span style="cursor: pointer">部分</span>
				</th>
				<th>
					<span style="cursor: pointer">工程</span>
				</th>
				<th>操作</th>
			</tr>
		</thead>
		<tbody>

        {{range $index, $elem :=.Catalogs}}
			<tr id="row{{.Id}}">
				<!--tr表格的行，td定义一个单元格，<th>
				标签定义表格内的表头单元格color="#A52A2A"-->
                <!-- <input type="hidden" id="inputrow{{.Id}}" name="inputrow{{.Id}}" value="{{.Id}}"/> -->
				<td>{{$index}}</td>
				<td>{{.Tnumber}}</td>
				<td >{{if .Exist}}<a href="/topic/view_b/{{.TopicId}}">{{end}}{{.Name}}</a></td>
				<td>{{.Drawn}}</td>
				<td>{{.Designd}}</td>
				<td>{{.Checked}}</td>
				<td>{{.Emamined}}</td>
				<td>{{.Verified}}</td>
				<td>{{.Approved}}</td>
				<td>{{.Data}}</td>
				<td>{{.DesignStage}}</td>
				<td>{{.Section}}</td>
				<td>{{.Projec}}</td>
				<td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow("row{{.Id}}")'/> 
				<input type='button' class='btn btn-default' name='update' value='修改' onclick='updateSelectedRow("row{{.Id}}")' /></td> 
		    </tr>
		{{end}}
	</tbody>
</table>
<input type="hidden" id="CategoryId" name="CategoryId" value="{{.CategoryId}}"/>
	   <tr>    
       <td colspan="4"><input type="button" class="btn btn-default" name="insert" value="增加目录行" onclick="insertNewRow()"/></td>    
       </tr>
</div>

<script type="text/javascript">
//*********这个是编辑表格
var flag = 0;  //标志位，标志第几行  
        /*    
         *添加一个新行    
         */    
        function insertNewRow(){    
            //获得表格有多少行    
            var rowLength = $("#orderTable tr").length;  
            //这里的rowId 就是row加上标志位组合的，为了方便起见所以分开好一点。  
            var rowId = "row" + flag;  
            //每次都往低flag+1的下标出添加tr，因为append是往标签内追加，所以用after
            //"<td>￥<input type='text' id='txtDrawn"+flag+"' value='' size='10'/></td>"  
            var insertStr = "<tr id="+rowId+">" 
                         +      "<td><input type='text' placeholder='序号' id='txtIndex"+flag+"' value='' size='10'/></td>" 
                         +      "<td><input type='text' placeholder='成果编号' id='txtTnumber"+flag+"' value='' size='10'/></td>"  
                         +      "<td><input type='text' placeholder='成果名称' id='txtName"+flag+"' value='' size='10'/></td>"  
                         +      "<td><input type='text' placeholder='制图' id='txtDrawn"+flag+"' value='' size='10'/></td>"
                         +      "<td><input type='text' placeholder='设计' id='txtDesignd"+flag+"' value='' size='10'/></td>" 
                         +      "<td><input type='text' placeholder='校核' id='txtChecked"+flag+"' value='' size='10'/></td>" 
                         +      "<td><input type='text' placeholder='审查' id='txtEmamined"+flag+"' value='' size='10'/></td>" 
                         +      "<td><input type='text' placeholder='核定' id='txtVerified"+flag+"' value='' size='10'/></td>" 
                         +      "<td><input type='text' placeholder='批准' id='txtApproved"+flag+"' value='' size='10'/></td>" 
                         +      "<td><input type='text' placeholder='出版日期' id='txtData"+flag+"' value='' size='10'/></td>" 
                         +      "<td><input type='text' placeholder='阶段' id='txtDesignStage"+flag+"' value='' size='10'/></td>"
                         +      "<td><input type='text' placeholder='专业' id='txtSection"+flag+"' value='' size='10'/></td>"
                         +      "<td><input type='text' placeholder='项目名称' id='txtProjec"+flag+"' value='' size='10'/></td>"
                         +      "<td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow(\""+rowId+"\")'/> <input type='button' class='btn btn-default' name='update' value='确定' onclick='saveAddRow(\""+rowId+"\",\""+flag+"\")' /></td>"                   
                         + "</tr>";  
            $("#orderTable tr:eq("+(rowLength-1)+")").after(insertStr);  //这里之所以减2 ，是因为减去底部的一行和顶部一行，剩下的为开始插入的索引。  
            flag++;  
        }    
  
        /*    
         *删除选中的行    
         */    
         function deleteSelectedRow(rowId){    
            //根据rowId查询出该行所在的行索引    
            if(confirm("确定删除该行吗？")){    
                $("#"+rowId).remove();    //这里需要注意删除一行之后 我的标志位没有-1，因为如果减一，那么我再增加一行的话，可能会导致我的tr的id重复，不好维护。
                // 提交到后台进行删除数据库
                    // alert("欢迎您：" + name) 
                    $.ajax({
                    type:"post",//这里是否一定要用post？？？
                    url:"/catalog/delete",
                    data: {CatalogId:rowId},
                        success:function(data,status){//数据提交成功时返回数据
                        alert("删除“"+data+"”成功！(status:"+status+".)");
                        }
                    });  
            }       
         }    
          
         /*    
          *修改选中的行    
          */    
         function updateSelectedRow(rowId){
         	var oldIndex = $("#"+rowId+" td:eq(0)").html();
            var oldTnumber = $("#"+rowId+" td:eq(1)").html();  
            var oldName = $("#"+rowId+" td:eq(2)").html();  
            var oldDrawn = $("#"+rowId+" td:eq(3)").html();
            var oldDesignd = $("#"+rowId+" td:eq(4)").html();
            var oldChecked = $("#"+rowId+" td:eq(5)").html();
            var oldEmamined = $("#"+rowId+" td:eq(6)").html();
            var oldVerified = $("#"+rowId+" td:eq(7)").html();
            var oldApproved = $("#"+rowId+" td:eq(8)").html();
            var oldData = $("#"+rowId+" td:eq(9)").html();
            var oldDesignStage = $("#"+rowId+" td:eq(10)").html();
            var oldSection = $("#"+rowId+" td:eq(11)").html();
            var oldProjec = $("#"+rowId+" td:eq(12)").html(); 
            // if(oldPrice != ""){//去掉第一个人民币符号  
            //     oldPrice = oldPrice.substring(1);  
            // }  
            var uploadStr = "<td><input type='text' id='txtIndex"+flag+"' value='"+oldIndex+"' size='10'/></td>"
                        + "<td><input type='text' id='txtTnumber"+flag+"' value='"+oldTnumber+"' size='10'/></td>"  
                        + "<td><input type='text' id='txtName"+flag+"' value='"+oldName+"' size='10'/></td>"  
                        + "<td><input type='text' id='txtDrawn"+flag+"' value='"+oldDrawn+"' size='10'/></td>"
                        + "<td><input type='text' id='txtDesignd"+flag+"' value='"+oldDesignd+"' size='10'/></td>"
                        + "<td><input type='text' id='txtChecked"+flag+"' value='"+oldChecked+"' size='10'/></td>"
                        + "<td><input type='text' id='txtEmamined"+flag+"' value='"+oldEmamined+"' size='10'/></td>"
                        + "<td><input type='text' id='txtVerified"+flag+"' value='"+oldVerified+"' size='10'/></td>"
                        + "<td><input type='text' id='txtApproved"+flag+"' value='"+oldApproved+"' size='10'/></td>"
                        + "<td><input type='text' id='txtData"+flag+"' value='"+oldData+"' size='10'/></td>"
                        + "<td><input type='text' id='txtDesignStage"+flag+"' value='"+oldDesignStage+"' size='10'/></td>"
                        + "<td><input type='text' id='txtSection"+flag+"' value='"+oldSection+"' size='10'/></td>"
                        + "<td><input type='text' id='txtProjec"+flag+"' value='"+oldProjec+"' size='10'/></td>"
                        + "<td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow(\""+rowId+"\")'/> <input type='button' class='btn btn-default' name='update' value='确定' onclick='saveUpdateRow(\""+rowId+"\",\""+flag+"\")' /></td>";  
            $("#"+rowId).html(uploadStr);  
         }    
  
         /*    
          *保存添加    
          */    
          function saveAddRow(rowId,flag){ 
            var newIndex = $("#txtIndex"+flag).val();
            var newTnumber = $("#txtTnumber"+flag).val();    
            var newName = $("#txtName"+flag).val();    
            var newDrawn = $("#txtDrawn"+flag).val();
            var newDesignd = $("#txtDesignd"+flag).val();
            var newChecked = $("#txtChecked"+flag).val();
            var newEmamined = $("#txtEmamined"+flag).val();
            var newVerified = $("#txtVerified"+flag).val();
            var newApproved = $("#txtApproved"+flag).val();
            var newData = $("#txtData"+flag).val();
            var newDesignStage = $("#txtDesignStage"+flag).val();
            var newSection = $("#txtSection"+flag).val();
            var newProjec = $("#txtProjec"+flag).val();
 
            var saveStr = "<td>" + newIndex + "</td>"
                        + "<td>" + newTnumber + "</td>"  
                        + "<td>" + newName + "</td>"  
                        + "<td>" + newDrawn + "</td>"
                        + "<td>" + newDesignd + "</td>"
                        + "<td>" + newChecked + "</td>"
                        + "<td>" + newEmamined + "</td>"
                        + "<td>" + newVerified + "</td>"
                        + "<td>" + newApproved + "</td>"
                        + "<td>" + newData + "</td>"
                        + "<td>" + newDesignStage + "</td>"
                        + "<td>" + newSection + "</td>"
                        + "<td>" + newProjec + "</td>"
                        + "<td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow(\""+rowId+"\")'/> <input type='button' class='btn btn-default' name='update' value='修改' onclick='updateSelectedRow(\""+rowId+"\")' /></td>";  
            $("#"+rowId).html(saveStr);//因为替换的时候只替换tr标签内的html 所以不用加上tr 
            // 这里再提交到后台保存起来update 
            if (newName)//如果返回的有内容  
                {  
                 var pid = $('#CategoryId').val();
                    // alert("欢迎您：" + name) 
                    $.ajax({
                    type:"post",//这里是否一定要用post？？？
                    url:"/catalog/post",
                    data: {Tnumber:newTnumber,Name:newName,Drawn:newDrawn,Designd:newDesignd,Checked:newChecked,Emamined:newEmamined,Verified:newVerified,Approved:newApproved,Data:newData,DesignStage:newDesignStage,Section:newSection,Projec:newProjec,ParentId:pid},
                        success:function(data,status){//数据提交成功时返回数据
                        alert("添加“"+data+"”成功！(status:"+status+".)");
                        }
                    });  
                }
          }

          /*    
          *保存修改    
          */
        function saveUpdateRow(rowId,flag){ 
            var newIndex = $("#txtIndex"+flag).val();
            var newTnumber = $("#txtTnumber"+flag).val();    
            var newName = $("#txtName"+flag).val();    
            var newDrawn = $("#txtDrawn"+flag).val();
            var newDesignd = $("#txtDesignd"+flag).val();
            var newChecked = $("#txtChecked"+flag).val();
            var newEmamined = $("#txtEmamined"+flag).val();
            var newVerified = $("#txtVerified"+flag).val();
            var newApproved = $("#txtApproved"+flag).val();
            var newData = $("#txtData"+flag).val();
            var newDesignStage = $("#txtDesignStage"+flag).val();
            var newSection = $("#txtSection"+flag).val();
            var newProjec = $("#txtProjec"+flag).val(); 
            var saveStr = "<td>" + newIndex + "</td>"
                        + "<td>" + newTnumber + "</td>"  
                        + "<td>" + newName + "</td>"  
                        + "<td>" + newDrawn + "</td>"
                        + "<td>" + newDesignd + "</td>"
                        + "<td>" + newChecked + "</td>"
                        + "<td>" + newEmamined + "</td>"
                        + "<td>" + newVerified + "</td>"
                        + "<td>" + newApproved + "</td>"
                        + "<td>" + newData + "</td>"
                        + "<td>" + newDesignStage + "</td>"
                        + "<td>" + newSection + "</td>"
                        + "<td>" + newProjec + "</td>"
                        + "<td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow(\""+rowId+"\")'/> <input type='button' class='btn btn-default' name='update' value='修改' onclick='updateSelectedRow(\""+rowId+"\")' /></td>";  
            $("#"+rowId).html(saveStr);//因为替换的时候只替换tr标签内的html 所以不用加上tr 
            // 这里再提交到后台保存起来update 
            if (newName)//如果返回的有内容  
                {  
                    $.ajax({
                    type:"post",//这里是否一定要用post？？？
                    url:"/catalog/modifycatalog",
                    data: {Tnumber:newTnumber,Name:newName,Drawn:newDrawn,Designd:newDesignd,Checked:newChecked,Emamined:newEmamined,Verified:newVerified,Approved:newApproved,Data:newData,DesignStage:newDesignStage,Section:newSection,Projec:newProjec,CatalogId:rowId},
                        success:function(data,status){//数据提交成功时返回数据
                        alert("修改“"+data+"”成功！(status:"+status+".)");
                        }
                    });  
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
           <!--  //使用正则表达式验证一下  
            // if(!/.+/.test(newTnumber)){  
            //     alert("输入的名称为空！");  
            //     return false;  
            // }  
            // if(!/^[0-9]+$/.test(newName)){  
            //     alert("输入的数目错误！");  
            //     return false;  
            // }  
            // if(!/^[0-9]+.?[0-9]*$/.test(newDrawn)){  
            //     alert("输入的价格错误！");  
            //     return false;  
            // }  -->