<script language="javascript">
  
  function create_table()
  {
   var num = document.getElementById("num").value;
   if(num == null )
   {
    num = 1;
   }
   var s1 = document.getElementById("d1");
   var vTable = document.createElement("Table");
   for(i=0; i<num; i++)
   {
    vTr = vTable.insertRow();
    vTd = vTr.insertCell();
    vTd.innerHTML="文件"+(i+1)+"：<input type=\"file\" name=\"file\" id=\"file\" />";
   }
   s1.appendChild(vTable);
  }
  
  function check_file()
  {
   var n=0;
   var flist = document.all.file;
   if(flist.length != undefined)
   {
    for(i=0;i<flist.length;i++)
    {
     if(flist[i].value != undefined)
     { 
      alert(flist[i].value);
     } 
    }
   }else{
    alert(flist.value);
   }
   
  }
  
 </script>
 </head>
 
<body onload="create_table();">
 <form id="form1" name="form1" onsubmit="check_file();">
  <input type="text" id="num" name="num" style="width:50px;" value="1" /><input type="button" onclick="create_table()"  value="设定" />
  <table cellpadding="0" cellspacing="0" border="0" bgcolor="#CCCCCC" id="t1">
  <tr>
   <td><div id="d1"></div></td>
  </tr>
 </table>
 <input type="submit" value="上传" />
 </form>
