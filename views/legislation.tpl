<!DOCTYPE html>

<title>首页 - 水利设计CMS系统</title>

<head>
    <link href="/static/youdao/g3.css" rel="stylesheet" type="text/css">
    <link href="/static/youdao/fanyi.css" rel="stylesheet" type="text/css">
    <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
</head>

<body class="open result-default">

    <div id="w" class="cf">
        <div id="transBackground">
            <div id="main" class="cf show-translate">
                <div id="transBtnTip">
                    <div id="transBtnTipInner">
                        点击翻译按钮继续，查看网页翻译结果。
                        <p class="ar">
                            <a href="/#" id="transBtnTipOK">我知道了</a>
                        </p> <b id="transBtnTipArrow"></b>
                    </div>
                </div>
                <div id="inputMod" class="column fl">
                    <div class="wrapper">
                        <form action="" method="post" id="transForm" name="transForm">

                            <div class="clearall">
                                <div id="clearInput" class="clog-js" style="display: block;">
                                    <div class="clearnotice">清空内容</div>
                                </div>
                            </div>
                            <div class="row border content" id="inputContent">
                                <textarea id="inputText" class="text" dir="ltr" tabindex="1" wrap="SOFT" name="name" placeholder="请输入要对标的内容，一行一条，《法规名称》"></textarea>
                                <!-- <div class="typo-suggest" style="display: none;">
                                    您是不是要翻译：
                                    <a class="spell-corrected" href="/#"></a>
                                </div> -->
                            </div>
                            <!-- <div class="row">
                                <a type="submit" id="translateBtn" href="/#" title="Enter自动对标" value="自动对标" name="action" class="button translateBtn"></a>
                            </div> -->

                            <span class="input-group-btn">
                                <button class="button translateBtn" type="button" id="checklist">
                                    <i class="glyphicon glyphicon-search"></i>
                                    Checklist!
                                </button>
                            </span>
                        </form>
                    </div>
                    <!-- end of wrapper --> </div>
                <!-- end of div inputMod -->
                <div id="outputMod" class="column fr">
                    <div class="wrapper">
                        <!-- end of entryList -->
                        <div class="row-hidden" id="outputHidden"></div>
                        <div id="translated" style="display: block;">

                            <div class="row">
                                <div class="row" id="outputText">
                                    <div class="translated_result">
                                        <!-- <p class="tgt">Austin</p> -->
                                    </div>
                                </div>
                            </div>
                            <div id="modeWrapper" class="read-mode" style="display: block;">

                                <a class="open-reading-mode title" href="h/#" title="" hidefocus=""></a>
                                <a class="close-reading-mode title" href="/#" title="" hidefocus=""></a>
                                <div class="opennotice">
                                    <div class="arrow"> <em></em>
                                        <span></span>
                                    </div>
                                    全屏阅读
                                </div>
                                <div class="closenotice">
                                    <div class="arrow">
                                        <em></em>
                                        <span></span>
                                    </div>
                                    关闭全屏阅读
                                </div>
                            </div>

                            <div class="row desc">

                                <div id="selectorSwitcher" class="selector-sprite selector-enable" style="display: block;">
                                    <span id="selectorStatus">划词</span>
                                </div>

                                <div class="read-mode" id="compareMode" style="display: block;">
                                    <label class="compare-mode compare-disable" for="compare">
                                        <input id="compare" name="compare" class="clog-js" type="checkbox" data-clog="COMPARE_CLICK" data-pos="web.o.righttop">
                                        <span class="compare-message">双语对照</span>
                                    </label>

                                </div>

                                <div class="tool">
                                    <a href="/#" id="speech" title=""></a>
                                    <a href="/#" id="resultScore" style="background-position: 0px -27px;"></a>
                                    <div class="speechnotice" style="display: none;">
                                        <div class="arrow">
                                            <em></em>
                                            <span></span>
                                        </div>
                                        朗读
                                    </div>
                                    <div class="copynotice" style="display: none;">
                                        <div class="arrow">
                                            <em></em>
                                            <span></span>
                                        </div>
                                        <span class="copy-notice">复制</span>
                                    </div>
                                    <a class="actions copyIt-js">
                                        <span id="copyOutput" class="copy-init"></span>
                                    </a>
                                </div>
                            </div>

                        </div>
                        <!-- end translated --> </div>
                    <!-- end of wrapper -->
                    <div class="row cf" id="addons" style="display: none;">

                </div>
                <div class="suggest" style="display: none;">
                    <div id="suggestYou"></div>
                </div>
            </div>

        </div>
        <!-- end of main --> </div>
    

</div>
<!-- end of w -->


<script type="text/javascript" src="/static/youdao/openapi.do" charset="utf-8"></script>

<div id="c_footer">
    <a href="/about.html">关于对标</a>
    <span class="c_fnl">|</span>
    <a href="/openapi">对标API</a>
    <span class="c_fnl">|</span>
    <a href="/copyright">HydroCMS首页</a>
    <span class="c_fnl">|</span>
    <a href="/">CMS智选</a>
    <span class="c_fnl">|</span>
    <a href="/about/index.html">关于HydroCMS</a>
    <span class="c_fnl">|</span>
    <a href="/">官方博客</a>
    <p class="c_fcopyright">© 2016 3xxx QQ504284</p>
</div>

<script type="text/javascript">
    var global = {};
    var abtest = "0";
    global.abTest = "0";
    global.sessionFrom = "https://www.baidu.com/link";
</script>

<script type="text/javascript" src="/static/youdao/dict_req_web_1.1.js"></script>
<script data-main="/fanyi" type="text/javascript" src="/static/youdao/fanyi.js"></script>

<div id="custheme"></div>
<!-- START NetEase Devilfish 2006 -->
<script src="/static/youdao/ntes.js" type="text/javascript"></script>

<script>
// <button class="btn btn-primary" id="export">导出excel</button>
$(document).ready(function(){
$("#checklist").click(function(){//这里应该用button的id来区分按钮的哪一个,因为本页有好几个button
    $("#translated")[0].style.display = 'block';
    $('#translated').css('display','block');
    $("#translated").css({ display: "block"});
                // alert(document.getElementById("inputText").value);
                // alert($("#inputText").val());
                // alert($("#inputText").text());//空
                // alert($("#inputText").html());//空
                $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/legislation/checklist",
                data: {name: $("#inputText").val()},
                success:function(data,status){//数据提交成功时返回数据
                    // alert(data[0].Title);
                    $.each(data,function(i,d){
                        // alert(data[i].Title);
                        $(".translated_result").append('<p>('+data[i].Id+')&nbsp;《'+data[i].LibraryTitle+'》&nbsp;('+data[i].LibraryNumber+')</p>');
                    }); 
                  // alert(data[1].Uname);
                  // $.each(data,function(i,d){
                  //   var tr=$("<tr></tr>");
                  //   var th1=$('<th>' + data[i].Number + '</th>');
                  //   var th2=$('<th>' + data[i].Title + '</th>');
                  //   var th3=$('<th><a href="' + data[i].Route + '"  target="_black"><i                    class="glyphicon glyphicon-download-alt"></i>下载</a></th>');
                  //   var th4=$('<th>' + data[i].Uname + '</th>');
                  //   var th5=$('<th>' + data[i].LiNumber + data[i].LibraryTitle + '</th>');
                  //   tr.append(th1);
                  //   tr.append(th2);
                  //   tr.append(th3);
                  //   tr.append(th4);
                  //   tr.append(th5);
                  //   $("#results").append(tr);
                  //   }); 
                  }       
            });
        });
     });
  </script> 
                  <!--  // function addRow(){
                    //     //声明tr td对象
                    //     var tr=$("<tr></tr>");
                    //     var td1=$("<td></td>");//拼接td，如果有td里有内容拼接时候填充进去
                    //     var td2=$("<td></td>");//拼接td，如果有td里有内容拼接时候填充进去
                    //     var td3=$("<td></td>");//拼接td，如果有td里有内容拼接时候填充进去
                    //     var td4=$("<td></td>");//拼接td，如果有td里有内容拼接时候填充进去
                    //     //将td添加到tr里
                    //     tr.append(td1);
                    //     tr.append(td2);
                    //     tr.append(td3);
                    //     tr.append(td4);
                    //     //或者
                    //     var tr="<tr><td></td> <td></td> <td></td> <td></td></tr>";//                 拼接html，如果有td里有内容拼接时候填充进去
                        
                    //     //再将tr添加到表格内
                    //     $("#表格的id").append(tr);
                    // } -->
                     
    


<!-- <script type="text/javascript">
  _ntes_nacc = "ydfanyiweb";
  neteaseTracker();
</script> -->
<!-- END NetEase Devilfish 2006 -->
<!-- START rlog -->
<!-- <script type="text/javascript">
  __rl_npid = "fanyiweb";
  (function() {
    var _rl = document.createElement('script');
    _rl.type = 'text/javascript'; _rl.async = true;
    _rl.src = 'http://rlogs.youdao.com/rlog.js';
    var s = document.getElementsByTagName('script')[0];
    s.parentNode.insertBefore(_rl, s);
  })();
</script> -->
<!-- END rlog -->


<!-- <div class="cpm-wrap">
    <div class="cpm">
        <div class="dict-cpm-mask"></div>
        <div class="cpm-click dict-inline-block dict-vam">
            <a class="cpm-close clog-js" data-clog="click-mac-close" href="javascript:;"></a>
        </div>
        <div class="cpm-height dict-inline-block dict-vam"></div>
    </div>
</div> -->



</body>
</html>