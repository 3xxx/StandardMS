<!DOCTYPE html>
<head>
    <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.js"> </script>
    <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
    <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
    <script src="/static/ueditor/ueditor.parse.js"></script>

  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.js"></script>
  <link type="text/css" href="/static/css/bootstrap.min.css" rel="stylesheet" />
<style type="text/css">
h3 {line-height: 150%;
  /*color: #DC143C;#000000*/
  color:#FFFFFF;
  background: #4682B4;/*#fff*/
  margin: 10;
  padding: 10;
  font-family: Georgia, Palatino, serif;
  }
</style>
<style>
i#delete
{
color:#DC143C;
}
</style>
</head>
<body>
<!-- <div class="navbar navba-default navbar-fixed-top">
</div> -->

<div class="container-fill">
  <!-- 这里添加项目目录树显示 -->
<!--   <h3> <font size="5">{{.CategoryPhase.Title}}</font>
    <font size="4"> » {{.CategorySpec.Title}} » </font>
    <font size="5">{{.Category.Title}}</font>
  </h3> -->
  <h2>项目简介</h2>
<div class="panel panel-default">
      <div class="panel-heading">
         <h4 class="panel-title">
            <a data-toggle="collapse" data-parent="#accordion" 
               href="#collapseOne">
               隐藏/展开项目简介
            </a>
         </h4>
      </div>
      <div id="collapseOne" class="panel-collapse collapse in">
  <div class="content">
  {{str2html .Category.Content}}
  <!-- 项目简介如何截取html呢？ -->
  </div>
      </div>
   </div> 



  </div>

<script>
        // 语法
        // uParse(selector,[option])
        /*
         selector支持
         id,class,tagName
         */
        /*
         目前支持的参数
         option:
         highlightJsUrl 代码高亮相关js的路径 如果展示有代码高亮，必须给定该属性
         highlightCssUrl 代码高亮相关css的路径 如果展示有代码高亮，必须给定该属性
         liiconpath 自定义列表样式的图标路径，可以不给定，默认'http://bs.baidu.com/listicon/',
         listDefaultPaddingLeft : 自定义列表样式的左边宽度 默认'20',
         customRule 可以传入你自己的处理规则函数，函数第一个参数是容器节点
         */

        uParse('.content',{
            rootPath : '/static/ueditor/'
        })
    </script>

</body>
</html>