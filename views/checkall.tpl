<!-- 给每个复选框赋值： -->
 
<!-- <input name="checkbox2" type="checkbox" class="input01" id="checkbox2" value="<%#DataBinder.Eval(Container.DataItem,"messageid")%>"> -->
 <div class="checkbox">
   <label><input name="checkbox2" type="checkbox" id="1" value=1>选项 11</label>
</div>
<div class="checkbox">
   <label><input name="checkbox2" type="checkbox" id="2" value=2>选项 21</label>
</div>
<!-- 点选“全选“复选框，实现对一组复选框操作 -->
<input id="all" type="checkbox" onclick="CheckAll()" /> 全选

<div class="checkbox">
  <label><input name="checkbox2" type="checkbox"  id="all" value=0 onclick="CheckAll()">全选</label></div><!--这2个全选跟先后有关系，无论哪个，只有放到后面的有效-->



<script type="text/javascript">
function CheckAll()
{
var a = document.getElementsByTagName('input');
var n = a.length;  
var allchecked = true;
for(var i=0;i<n;i++)
{
if(a[i].checked)
allchecked = true;
else
allchecked = false;
}
for(var i=0;i<n;i++)
{
if(allchecked)
a[i].checked = true;
else
a[i].checked = false;
}
}
</script>


<!-- <label for="name">默认的复选框和单选按钮的实例</label>
<div class="checkbox">
<label>
<input type="checkbox" value="">选项 1</label>
</div>
<div class="checkbox">
<label>
<input type="checkbox" value="">选项 2</label>
</div>

<div class="radio">
<label>
<input type="radio" name="optionsRadios" id="optionsRadios1" 
         value="option1" checked>选项 1</label>
</div>
<div class="radio">
<label>
<input type="radio" name="optionsRadios" id="optionsRadios2" 
         value="option2">选项 2 - 选择它将会取消选择选项 1</label>
</div>
-->
<!-- <label for="name">内联的复选框和单选按钮的实例</label>
<div>
<label class="checkbox-inline">
<input type="checkbox" id="inlineCheckbox1" value="option1">选项 1</label>
<label class="checkbox-inline">
<input type="checkbox" id="inlineCheckbox2" value="option2">选项 2</label>
<label class="checkbox-inline">
<input type="checkbox" id="inlineCheckbox3" value="option3">选项 3</label>
<label class="checkbox-inline">
<input type="radio" name="optionsRadiosinline" id="optionsRadios3" 
         value="option1" checked>选项 1</label>
<label class="checkbox-inline">
<input type="radio" name="optionsRadiosinline" id="optionsRadios4" 
         value="option2">选项 2</label>
</div>
-->
<!-- 给每个复选框赋值： -->

<!-- <input name="checkbox2" type="checkbox" class="input01" id="checkbox2" value="<%#DataBinder.Eval(Container.DataItem,"messageid")%>
"> -->
<!--  <div class="checkbox">
<label>
<input name="checkbox2" type="checkbox" id="1" value=1>选项 11</label>
</div>
<div class="checkbox">
<label>
<input name="checkbox2" type="checkbox" id="2" value=2>选项 21</label>
</div>
-->
<!-- <! 点选“全选“复选框，实现对一组复选框操作 -->
<!-- <input id="all" type="checkbox" onclick="CheckAll()" />
全选 -->
<!-- <div class="checkbox">
<label>
<input name="checkbox2" type="checkbox"  id="all" value=0 onclick="CheckAll()">全选</label>
</div>
-->
<!--这2个全选跟先后有关系，无论哪个，只有放到后面的有效-->
<!-- /*全选*/ -->
<!-- <script type="text/javascript">
function checkall (s,k){
var a = document.getElementsByTagName('input');
var n = a.length;
for (var i=0; i<n; i++)
{
if((a[i].type == "checkbox") && ( a[i].name.substr(0,k-1)==s )){
a[i].checked = true;
}
}
}
</script>
-->
<!-- /*反选*/ -->
<!-- <script type="text/javascript">
function uncheck (s,k){
var a = document.getElementsByTagName('input');
var n = a.length;
for (var i=0; i
<n; i++){
if((a[i].type == "checkbox") &&  (a[i].name.substr(0,k-1)==s) ){
if(a[i].checked == true){
a[i].checked = false;
}
else{ a[i].checked = true; }
}
}
}
</script>
-->
<!-- <script type="text/javascript">
function doall (s,k,n ){
if( k ){  checkall (s,n ) }
else{ uncheck (s,n ) }
}
</script>
-->
<!-- 而在视图中首先是设计每一行的checkbox的id，然后在最后一个代表“全选”的checkbox中调用上面代码的doall 方法，由参数的不同实现，全选和反选。代码片段如下：每一行的checkbox： -->
<!-- <input id="all" type="checkbox" onclick="doall('depart',true,7);" />
全选 -->


<!-- var  values=new array();
$("input[name='checkbox']").each(function(){
var  check=$(this).attr("checked");//判断是否选中
if(check){ 
  var  value=$(this).val();
  values.push(value);//values  里面存放的就是选中的值
}
}
)

//在servlet里面取就是：String[] str = getParameterValues（"aaa");
 //如果是用js传就是：
 var valus = document.getElementsByName("aaa");
  if(valus.length!=0)
  {
  var str = ""；
  for(var i=0;i<valus.length;i++)  
  {
  if(valus[i].checked)
  {
  str+= valus[i].value;
  //如果被选择就形成一个字符串格式是:123456,124568,456788这种形式的，传到后台再分割出来就可以了，试试吧........
  }
  }
  }

</script> -->

<!-- a.jsp -->
<!-- while(...){
<td>
  <input type="checkbox" name="item_no" value="<%=rs.getString(1)%>"><%=rs.getString(1)%>
  </td>
} -->

<!-- b.jsp -->
<!-- String item_no[] = request.getParameterValues("item_no");
for(int i = 0; i<item_no.length; i++){
item_no[i] .... //选中的值。
} -->

<script type="text/javascript">
  function modifyUser() {
   var count = 0;
   var j = 0;
   for (var i = 0; i<document.getElementsByName("selectFlag").length; i++) {
   if (document.getElementsByName("selectFlag")[i].checked) {
   j = i;
   count++;
   }
     
  if (count == 1) {
   window.self.location = "usermaint.do?userId=" +  
  document.getElementsByName("selectFlag")[j].value;
   }
}
</script>

<!-- selectFlag是复选框的名字.  
window.self.location = "usermaint.do?userId=" +  
  document.getElementsByName("selectFlag")[j].value;  
这个是你转向的*.do,当然你也可以换成自己的JSP(*.jsp)页面+传递过去的ID值  
当你再提交时 onClick="modifyUser()" 按钮的type="button"; -->


<!-- <input type="checkbox" name="aaa" value="<%=id%>" onclick="check(this)">
 你需要用javascript获得checkbox选中的值，然后再把获得的值放到一个hidden域里
javascript获取了值也是要传放到后台的，你只接用servlet来获得
 楼主还是去看一下getParameterValues方法
 String[] s=request.getParameterValues("checkboxname");
 它是返回所有name为checkboxname的value值
 比如你单选了id=a,id=b,id=c的复选框，则String[] s={a,b,c}
 在Servlet中得到了S，然后
 for(int i=0;i<s.length;i++)
 {
     //查询数据库，得到相应的信息，select * from tablename where id=(String)s[i]
      将信息返回加到一个数组中，list
 }
 再就是动态生成一个网页，将list的信息取出，加到表格 里-->
<!--   <script>
function aa() 
{  
     var bb = "";  
     var temp = ""; 
     var a = document.getElementsByTagName("input");
     for ( var i = 0; i<a.length; i++) {  
     if (a[i].checked) {  
     temp = a[i].value;  
     bb = bb + "," +temp;  
     }  
     } 
      // return bb 
     document.getElementById("tempString").value = bb
      .substring(1, bb.length);  
}
       
</script>