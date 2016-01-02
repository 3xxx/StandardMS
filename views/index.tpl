<!DOCTYPE html>
{{template "header"}}
<title>首页 - 水利设计CMS系统</title>

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
            .editormd-html-preview {
                width: 100%;
                margin: 0 auto;
            }
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
      <h2>{{.Website}}</h2>
      <h2><a href="mailto:504284@qq.com">{{.Email}}</a></h2>
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
        <!-- <div class="panel panel-default">
        -->
        <!-- <div class="panel-heading">
        -->
        <!--          <h4 class="panel-title">
        <a data-toggle="collapse" data-parent="#accordion" 
               href="#{{$index}}">点击我进行展开，再次点击我进行折叠。第 1 部分--hide 方法</a>
      </h4>
      -->
      <!-- </div>
      -->
      <!--       <div id="{{$index}}" class="panel-collapse collapse in">
      <div class="panel-body">{{.Content}}</div>
    </div>
    -->
    <!-- </div>
    -->
    <!-- <div class="panel panel-default">
    -->
    <!-- <div id="collapseOne" class="panel-collapse collapse in">
    "test-editormd-view2"-->
    <div id="test-editormd-view2-{{.Id}}">
      <textarea id="append-test" style="display:none;">{{substr .Content 0 500}}</textarea>
    </div>



    <!-- <script type="text/javascript">
            $(function() {
                var testEditormdView, testEditormdView2;
                // $.get("test.md", function(markdown) {
                    testEditormdView = editormd.markdownToHTML("test-editormd-view", {
                        // markdown        : markdown ,//+ "\r\n" + $("#append-test").text(),
                        //htmlDecode      : true,       // 开启 HTML 标签解析，为了安全性，默认不开启
                        htmlDecode      : "style,script,iframe",  // you can filter tags decode
                        //toc             : false,
                        tocm            : true,    // Using [TOCM]
                        //tocContainer    : "#custom-toc-container", // 自定义 ToC 容器层
                        //gfm             : false,
                        //tocDropdown     : true,
                        // markdownSourceCode : true, // 是否保留 Markdown 源码，即是否删除保存源码的 Textarea 标签
                        emoji           : true,
                        taskList        : true,
                        tex             : true,  // 默认不解析
                        flowChart       : true,  // 默认不解析
                        sequenceDiagram : true,  // 默认不解析
                    });
                    
                    //console.log("返回一个 jQuery 实例 =>", testEditormdView);
                    
                    // 获取Markdown源码
                    //console.log(testEditormdView.getMarkdown());
                    
                    //alert(testEditormdView.getMarkdown());
                // });
                testEditormdView2 = editormd.markdownToHTML("test-editormd-view2-{{.Id}}", {
                    htmlDecode      : "style,script,iframe",  // you can filter tags decode
                    emoji           : true,
                    taskList        : true,
                    tex             : true,  // 默认不解析
                    flowChart       : true,  // 默认不解析
                    sequenceDiagram : true,  // 默认不解析
                });
            });
    </script> -->



<!-- <script src="js/jquery.min.js"></script> -->
        <script src="/static/editor.md/lib/marked.min.js"></script>
        <script src="/static/editor.md/lib/prettify.min.js"></script>
        
        <script src="/static/editor.md/lib/raphael.min.js"></script>
        <script src="/static/editor.md/lib/underscore.min.js"></script>
        <script src="/static/editor.md/lib/sequence-diagram.min.js"></script>
        <script src="/static/editor.md/lib/flowchart.min.js"></script>
        <script src="/static/editor.md/lib/jquery.flowchart.min.js"></script>

        <!-- <script src="../editormd.js"></script> -->
        <script type="text/javascript">
            $(function() {
                var testEditormdView, testEditormdView2;
                
                $.get("test.md", function(markdown) {
                    
            testEditormdView = editormd.markdownToHTML("test-editormd-view", {
                        markdown        : markdown ,//+ "\r\n" + $("#append-test").text(),
                        //htmlDecode      : true,       // 开启 HTML 标签解析，为了安全性，默认不开启
                        htmlDecode      : "style,script,iframe",  // you can filter tags decode
                        //toc             : false,
                        tocm            : true,    // Using [TOCM]
                        tocContainer    : "#custom-toc-container", // 自定义 ToC 容器层
                        //gfm             : false,
                        //tocDropdown     : true,
                        // markdownSourceCode : true, // 是否保留 Markdown 源码，即是否删除保存源码的 Textarea 标签
                        emoji           : true,
                        taskList        : true,
                        tex             : true,  // 默认不解析
                        flowChart       : true,  // 默认不解析
                        sequenceDiagram : true,  // 默认不解析
                    });
                    
                    //console.log("返回一个 jQuery 实例 =>", testEditormdView);
                    
                    // 获取Markdown源码
                    //console.log(testEditormdView.getMarkdown());
                    
                    //alert(testEditormdView.getMarkdown());
                });
                    
                testEditormdView2 = editormd.markdownToHTML("test-editormd-view2", {
                    htmlDecode      : "style,script,iframe",  // you can filter tags decode
                    emoji           : true,
                    taskList        : true,
                    tex             : true,  // 默认不解析
                    flowChart       : true,  // 默认不解析
                    sequenceDiagram : true,  // 默认不解析
                });
            });
        </script>






    <!-- </div>
    -->
    <!-- </div>
    -->
    <!-- <label>项目简介:</label>
  -->
  <!-- <div id="test-editormd-view2">
  -->
  <!-- <textarea id="append-test" style="display:none;">{{.Content}}</textarea>
</div>
-->
<!-- {{substr .Content 0 500}} -->
<br />

<!-- <a href={{.Route}}></a>
align ="right/middle"> -->
<img style="-webkit-user-select: none; cursor: zoom-in;" src="{{.Route}}" width="50%" align="middle">
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
<!--   <div class="col-md-3">
<h3>项目</h3>
<ul>
<ol>
{{range .Categories}}
<li>
<a href="/?cate={{.Title}}">
<i class="glyphicon glyphicon-th-large"></i>
{{.Title}}
</a>
</li>
{{end}}
</ol>
</ul>
</div>
-->
</div>

</body>
</html>