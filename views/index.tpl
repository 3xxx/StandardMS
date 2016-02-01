<!DOCTYPE html>
{{template "header"}}

<title>首页 - 水利设计CMS系统</title>
    <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.js"> </script>
    <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
    <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
    <script src="/static/ueditor/ueditor.parse.js"></script>


        <script src="/static/editor.md/lib/marked.min.js"></script>
        <script src="/static/editor.md/lib/prettify.min.js"></script>
      
        <script src="/static/editor.md/lib/raphael.min.js"></script>
        <script src="/static/editor.md/lib/underscore.min.js"></script>
        <script src="/static/editor.md/lib/sequence-diagram.min.js"></script>
        <script src="/static/editor.md/lib/flowchart.min.js"></script>
        <script src="/static/editor.md/lib/jquery.flowchart.min.js"></script>
        <script src="/static/editor.md/editormd.min.js"></script> 
        <link rel="stylesheet" href="/static/editor.md/css/editormd.css" /> 
<style>
.page-header {
    padding-bottom: 9px;
    margin: 40px 0 20px;
    border-bottom: 1px solid #eee;
}
h1[id] {
    padding-top: 80px;
    margin-top: -45px;
}            
/*.editormd-html-preview {
      width: 100%;
      margin: 0 auto;
 }*/
</style>
<style type="text/css">
    /* Custom Styles */
    ul.nav-tabs{
        width: 140px;
        margin-top: 20px;
        border-radius: 4px;
        border: 1px solid #ddd;
        box-shadow: 0 1px 4px rgba(0, 0, 0, 0.067);
    }
    ul.nav-tabs li{
        margin: 0;
        border-top: 1px solid #ddd;
    }
    ul.nav-tabs li:first-child{
        border-top: none;
    }
    ul.nav-tabs li a{
        margin: 0;
        padding: 8px 16px;
        border-radius: 0;
    }
    ul.nav-tabs li.active a, ul.nav-tabs li.active a:hover{
        color: #fff;
        background: #0088cc;
        border: 1px solid #0088cc;
    }
    ul.nav-tabs li:first-child a{
        border-radius: 4px 4px 0 0;
    }
    ul.nav-tabs li:last-child a{
        border-radius: 0 0 4px 4px;
    }
    ul.nav-tabs.affix{
        top: 50px; /* Set the top position of pinned element */
    }
  </style>
</head>

<body data-spy="scroll" data-target="#myScrollspy">
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>

<div class="container-fill">
  <!-- <div class="row">
  -->
  <div class="col-xs-3" id="myScrollspy">
    <ul class="nav nav-tabs nav-stacked" data-spy="affix" data-offset-top="50">
      <li class="active">
        <a href="#section-1"> <i class="glyphicon glyphicon-plane"></i>
          最新项目列表
        </a>
      </li>
      <li>
        <a href="#section-2"> <i class="glyphicon glyphicon-certificate"></i>
          最新成果列表
        </a>
      </li>
      <li>
        <a href="#section-3">
          <i class="glyphicon glyphicon-user"></i>
          成果贡献者
        </a>
      </li>
      <li>
        <a href="#section-4">
          <i class="glyphicon glyphicon-asterisk"></i>
          最多访问成果
        </a>
      </li>
      <li>
        <a href="#section-5">
          <i class="glyphicon glyphicon-download-alt"></i>
          最多下载成果
        </a>
      </li>
    </ul>
  </div>

  <div class="col-lg-9" role="main">
    <div class="bs-docs-section">
      <div class="page-header">
      <!-- <h2>{{.Website}}</h2> -->
      <!-- <h2><a href="mailto:504284@qq.com">{{.Email}}</a></h2> -->
        <h1 id="section-1">
          <i class="glyphicon glyphicon-star-empty"></i>
          最新项目列表
        </h1>
      </div>
      <ol>
{{range $index, $elem := .Categories}}
  {{if lt $index 20}}
        <h3>
          <li>
            <a href="/category?op=view&id={{.Id}}">{{.Title}}</a>
          </li>
        </h3>
        
  <div class="content">
  {{str2html .Content}}
  <!-- 项目简介如何截取html呢？ -->
  </div>
   <!--  <div id="editormd-{{$index}}" name="test-editormd-view2" class="content">class="markdown-body editormd-html-preview"
       <textarea id="{{$index}}" style="display:none;">{{.Content}}</textarea> append-test
    </div> --> 
  <!-- {{substr .Content 0 500}} 因为是存成html格式的内容，所以，剪断后，html就不完整了-->
  <br />
  <!-- <img style="-webkit-user-select: none; cursor: zoom-in;" src="{{.Route}}" width="50%" align="middle"> -->
<!-- <script type="text/javascript">
            $(function() {
                var testEditormdView2;//testEditormdView, 
            //     $.get("test.md", function(markdown) {
            // testEditormdView = editormd.markdownToHTML("test-editormd-view", {
            //             markdown        : markdown ,//+ "\r\n" + $("#append-test").text(),
            //             //htmlDecode      : true,       // 开启 HTML 标签解析，为了安全性，默认不开启
            //             htmlDecode      : "style,script,iframe",  // you can filter tags decode
            //             //toc             : false,
            //             tocm            : true,    // Using [TOCM]
            //             tocContainer    : "#custom-toc-container", // 自定义 ToC 容器层
            //             //gfm             : false,
            //             //tocDropdown     : true,
            //             // markdownSourceCode : true, // 是否保留 Markdown 源码，即是否删除保存源码的 Textarea 标签
            //             emoji           : true,
            //             taskList        : true,
            //             tex             : true,  // 默认不解析
            //             flowChart       : true,  // 默认不解析
            //             sequenceDiagram : true,  // 默认不解析
            //         });
            //     });
                testEditormdView2 = editormd.markdownToHTML("editormd-{{$index}}", {
                    htmlDecode      : "style,script,iframe",  // you can filter tags decode
                    emoji           : true,
                    taskList        : true,
                    tex             : true,  // 默认不解析
                    flowChart       : true,  // 默认不解析
                    sequenceDiagram : true,  // 默认不解析
                });
            });
          </script> -->

  {{end}}
{{end}}
</ol>
</div>
<hr>

<div class="bs-docs-section">
      <div class="page-header">
      <h1 id="section-2"> <i class="glyphicon glyphicon-star-empty"></i>
        最新成果列表
      </h1>
      </div>

<ol>
{{range $index, $elem := .Topics}}
     {{if lt $index 20}}
<li>
<h4>
  <a href="/topic/view_b/{{.Id}}">{{.Title}}</a>
</h4>
</li>
<h6 class="text-muted">
成果由{{.Author}}上传于{{dateformat .Created "2006-01-02 T 15:04:05"}}，共有{{.Views}}次浏览，{{.ReplyCount}}个评论
</h6>
<p>{{.Content}}</p>
{{end}}
    {{end}}
</ol>
</div>
  </div>

</div>

<script type="text/javascript">
            // $(function() {
            //     var testEditormdView, testEditormdView2;
                
            //     $.get("test.md", function(markdown) {
                    
            // testEditormdView = editormd.markdownToHTML("test-editormd-view", {
            //             markdown        : markdown ,//+ "\r\n" + $("#append-test").text(),
            //             //htmlDecode      : true,       // 开启 HTML 标签解析，为了安全性，默认不开启
            //             htmlDecode      : "style,script,iframe",  // you can filter tags decode
            //             //toc             : false,
            //             tocm            : true,    // Using [TOCM]
            //             tocContainer    : "#custom-toc-container", // 自定义 ToC 容器层
            //             //gfm             : false,
            //             //tocDropdown     : true,
            //             // markdownSourceCode : true, // 是否保留 Markdown 源码，即是否删除保存源码的 Textarea 标签
            //             emoji           : true,
            //             taskList        : true,
            //             tex             : true,  // 默认不解析
            //             flowChart       : true,  // 默认不解析
            //             sequenceDiagram : true,  // 默认不解析
            //         });
                    
            //         //console.log("返回一个 jQuery 实例 =>", testEditormdView);
                    
            //         // 获取Markdown源码
            //         //console.log(testEditormdView.getMarkdown());
                    
            //         //alert(testEditormdView.getMarkdown());
            //     });
                    
            //     testEditormdView2 = editormd.markdownToHTML("test-editormd-view2", {
            //         htmlDecode      : "style,script,iframe",  // you can filter tags decode
            //         emoji           : true,
            //         taskList        : true,
            //         tex             : true,  // 默认不解析
            //         flowChart       : true,  // 默认不解析
            //         sequenceDiagram : true,  // 默认不解析
            //     });
            // });




    //实例化编辑器
    //议使用工厂方法getEditor创建和引用编辑器实例，如果在某个闭包下引用该编辑器，直接调用UE.getEditor('editor')就能拿到相关的实例
    // var ue = UE.getEditor('editor');

//  setTimeout(function(){ uParse('#content', {
//  'highlightJsUrl':'{/static/ueditor/third-party/SyntaxHighlighter/shCore.js',
//  'highlightCssUrl':'/static/ueditor/third-party/SyntaxHighlighter/shCoreDefault.css'})
// }, 300);   
// ue.addListener("ready", function () {
// uParse('.content', {
//     rootPath: '/static/ueditor/'
// });
// });

// $(function(){
//         var content =$('#content').val();
//         //判断ueditor 编辑器是否创建成功
//         ue.addListener("ready", function () {
//         // editor准备好之后才可以使用
//         ue.setContent({{.Category.Content}});
//         });
//     });
</script>
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