
<!DOCTYPE HTML>
<html>
<head>
    <meta charset="utf-8">
    <title>Karl | The Pursuit of Happiness</title>
    <meta name="author" content="SkyArrow">
    <meta name="viewport" content="width=device-width; initial-scale=1; maximum-scale=1">
    <meta property="og:site_name" content="Light"/>
    <link rel="stylesheet" href="/assets/css/font-awesome.min.css" media="screen" type="text/css">
    <link rel="stylesheet" href="/assets/css/style.css" media="screen" type="text/css">
</head>
<body>
<header id="header" class="inner"><div class="alignleft">
        <h1><a href="/">Golune</a></h1>
    </div>
    <nav id="main-nav" class="alignright">
        <ul>
            <li>
                <a href="/">
                    <i class="fa fa-home"></i>
                    首页
                </a>
            </li>
            <li>
                <a href="/blog/archive">
                    <i class="fa fa-archive"></i>
                    归档
                </a>
            </li>
            <li>
                <a href="/blog/tags">
                    <i class="fa fa-tags"></i>
                    标签
                </a>
            </li>

        </ul>
        <div class="clearfix"></div>
    </nav>
    <div class="clearfix"></div></header>
<div id="content" class="inner">
    <div id="main-col" class="alignleft">
        <div id="wrapper">
            {*<div class="text-editor-wrap">
                <div class="title-bar"><span class="title">bash</span></div>
                <div class="text-body">
                    $ <span id="typed"></span>
                </div>
            </div>
            <br>*}
            {{range .blog}}
            <article class="post">
                <div class="post-content">
                    <header>
                        <div class="icon">
                            <i class="fa fa-2x {{.Category | cateIcon}}"></i> </div>
                        <h1 class="title"><a href="/blog/html/{{.CreatedYear}}/{{.CreatedMonth}}/{{.Htmltag}}">{{.Title}}</a></h1>
                        <div class="mate">
                            <span class="cate"><i class="fa fa-folder-open"></i> <a href="/blog/cate/{{.Category}}">{{.Category | replacecate}}</a></span>
                            <span class="cate"><i class="fa fa-eye"></i> {{.Hot}}</span>
                            <span class="created_at"><i class="fa fa-clock-o"></i> {{.CreatedAt | unix2str}}</span>
                            <span class="comment"><i class="fa fa-comments-o"></i>
                                <span id="sourceId::c5ee1e7c382433411387ea2ca52cc8833e1f1" class="cy_cmt_count">0</span>
                            </span>
                            <span class="tags"><i class="fa fa-tags"></i>
                                {{range .TagBox}}
                                <a href="/blog/tags/{{.Id | id2str}}">{{.Name}}</a><e></e>
                                {{end}}
                            </span>
                        </div>
                    </header>
                    <div class="entry">
                        {{.Content | markdownCommon | unescaped}}
                    </div>
                    <div class="footer">
                    </div>
                </div>

            </article>
            {{end}}

            <nav id="pagination">
                <div class="clearfix"></div>
            </nav>
        </div>
    </div>
    <aside id="sidebar" class="alignright">


        <div class="search">
            <form action="http://google.com/search" method="get" accept-charset="utf-8">
                <input type="text" name="q" results="0" placeholder="Search">
                <input type="hidden" name="q" value="site:zespia.tw/hexo-theme-light">
            </form>
        </div>


        <div class="widget">
            <h3 class="title">分类</h3>
            <ul class="entry cate">
                <li>
                    <i class="nfa fa fa-rocket"></i>
                    <a href="/blog/cate/technology">
                        技术
                    </a>
                </li>
                <li>
                    <i class="nfa fa fa-bus"></i>
                    <a href="/blog/cate/life">
                        生活
                    </a>
                </li>
                <li>
                    <i class="nfa fa fa-leaf"></i>
                    <a href="/blog/cate/talk">
                        杂谈
                    </a>
                </li>
            </ul>
        </div>

    </aside>
    <div class="clearfix"></div>
</div>
<footer>
    <div class="footavatar">
        <!--<img src="http://guimeng.sinaapp.com/wp-content/themes/forget/images/logo.jpg" class="footphoto">-->
        <a class="footright" target="_blank" href="https://github.com/guhao022"><i class="fa fa-github fa-2x foottag"></i></a>
        <ul class="footinfo">

            <li class="fname">寥落风</li>
            <li class="finfo">放眼难觅旧衣冠,疑真疑幻,如梦如烟...</li>
        </ul>
    </div>
    <div class="Copyright">
        <p>© 2014 Powered by Golune</p>
    </div>
</footer>
<script src="http://libs.baidu.com/jquery/1.9.0/jquery.js"></script>
<script src="/assets/js/highlight.js"></script>
<script src="/assets/js/jquery.simplePagination.js"></script>
<script src="/assets/js/typed.min.js"></script>
<link type="text/css" rel="stylesheet" href="/assets/css/simplePagination.css"/>
<link type="text/css" rel="stylesheet" href="/assets/css/monokai.css"/>
<script type="text/javascript">
    $(document).ready(function() {
        $('pre code').each(function(i, block) {
            hljs.highlightBlock(block);
        });
    });


    var page_index;
    var itemsOnPage = 10;
    $(function() {
        $("#typed").typed({
            strings: ["There is my blog.", "Welcome!",  "Have a great day!"],
            typeSpeed: 50,
            backDelay: 500,
            loop: false,
            contentType: 'html', // or text
            // defaults to false for infinite loop
            loopCount: false,
            callback: function(){ foo(); },
            resetCallback: function() { newTyped(); }
        });

        $(".reset").click(function(){
            $("#typed").typed('reset');
        });
        $("#pagination").pagination({
            items: {{.blog | len}},
            itemsOnPage: itemsOnPage,
            onInit: changePage,
            onPageClick: changePage,
            prevText: '上一页',
            nextText: '下一页',
            cssStyle: 'light-theme'
        });
        function changePage(){
            console.log("changePage");
            page_index = $("#pagination").pagination('getCurrentPage') -1;
            $("#wrapper .post").hide();
            for(var i = page_index * itemsOnPage; i < page_index * itemsOnPage + itemsOnPage; i++){
                $("#wrapper .post:eq(" + i + ")").show();
            }
        }
    });
</script>
</body>
</html>