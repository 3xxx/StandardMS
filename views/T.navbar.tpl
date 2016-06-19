{{define "navbar"}}
<nav class="navbar navbar-default">
  <!-- /.navbar-collapse 前后必须用这个nav包裹起来！否则没有高亮显示-->
  <!-- <div class="container-fluid"> -->
    <!-- /.container-fluid -->
    <!-- Brand and toggle get grouped for better mobile display -->

    <!-- <a class="navbar-brand" href="/spider">水利设计院</a> -->

    <ul class="nav navbar-nav">
      <li {{if .IsSpider}}class="active"{{end}}>
        <a href="/getspider">水利设计院</a>
      </li>

      <li {{if .IsHome}}class="active"{{end}}>
        <a href="/">
          首页
          <span class="sr-only">(current)</span>
        </a>
      </li>
      <li {{if .IsCategory}}class="active"{{end}}>
        <a href="/category">项目</a>
      </li>
<!--       <li {{if .IsCategoryb}}class="active"{{end}}>
        <a href="/category_b">项目B</a>
      </li> -->
      <li {{if .IsTopic}}class="active"{{end}}>
        <a href="/topic">成果</a>
      </li>
      <li {{if .IsStandard}}class="active"{{end}}>
        <a href="/standard">规范</a>
      </li>
    </ul>

    <form class="navbar-form navbar-left" role="search" method="get" action="/search">
      <div class="form-group">
        <input id="tuming" type="text" class="form-control"  class="search-query span2" placeholder="Search local" name="tuming"></div>
      <button type="submit" class="btn btn-default">Submit</button>
    </form>

    <div class="pull-right">
      <ul class="nav navbar-nav">
        {{if .IsLogin}}
<!--         <li>
          <a href="/login?exit=true">{{.Uname}}</a>
        </li> -->
        <li class="dropdown">
        <a href="#" class="dropdown-toggle" data-toggle="dropdown">{{.Uname}} <b class="caret"></b></a>
        <ul class="dropdown-menu">
          <li><a href="/user/getuserbyusername?username={{.Uname}}">用户资料</a></li>
          <li><a href="/category/viewbyuname?uname={{.Uname}}">项目列表</a></li>
          <li><a href="/topic/viewbyuname?uname={{.Uname}}">成果列表</a></li>
          <li class="divider"></li>
          <li><a href="/login?exit=true">退出</a></li>
        </ul>
      </li>
        {{else}}
        <li>
          <a href="/login">登陆</a>
        </li>
        {{end}}
        <li {{if .IsWiki}}class="active"{{end}}>
          <a href="/wiki">Wiki</a>
        </li>
        <!-- <li {{if .IsHelp}}class="active"{{end}}>
          <a href="/help">自述</a>
        </li> -->

        <li {{if .IsTask}}class="active"{{end}}>
          <a href="/todo">Todo</a>
        </li>        
      </ul>
    </div>

  <!-- </div> -->
</nav>
{{end}}