<!DOCTYPE html>
{{template "header"}}
<title>项目&成果 - 水利设计CMS系统</title>
<style type="text/css">
#button{
  /*height:100px;*/
  width:100px;
  /*line-height: 2*/
  }
/*h3 {line-height: 250%;
  color: #DC143C;#000000
  color:#FFFFFF;
  background: #4682B4;#fff
  margin: 10;
  padding: 10;
  font-family: Georgia, Palatino, serif;
  }*/
</style>


</head>

<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>

<div class="col-lg-12">
<h3>炫耀主机</h3>
{{range $index, $elem :=.Service}}
 <a href="{{.Link}}" class="btn btn-primary" id="button"><i class="glyphicon glyphicon-flash"></i>{{.Name}}<i class="glyphicon glyphicon-screenshot"></i></a> 
{{end}}
</div>

<div class="col-lg-6">
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