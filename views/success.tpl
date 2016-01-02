<!DOCTYPE html>
{{template "header"}}
<title>首页 - 水利设计CMS系统</title>
</head>

<body data-spy="scroll" data-target="#myScrollspy">
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>
<div id="content" class="col-md-8 col-md-offset-2">
<div class="col-md-6 auth-page">
  <div class="auth-page">
    <h3 class="title">
      <span class="glyphicon glyphicon-ok"></span>
      注册成功！请登陆。
    </h3>
    <p class="well">注册成功！请登陆。</p>
    <p>
      <a href="./login" class="btn btn-default">
      	立即登陆&nbsp;&nbsp; <span class="glyphicon glyphicon-circle-arrow-right"></span>
      </a>
    </p>
  </div>
</div>
	</div>

 <script type="text/javascript">
  function backToHome(){
  window.location.href="/login";
  return false;
}
</script>  

	</body>
</html>