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
            <article class="post">
                <div style="margin-left: 30px">
                    <div class="collection-title">
                        <h2>
                            {{.cate | replacecate}}
                            <small>分类</small>
                        </h2>
                    </div>

                    {{range .blog}}
                    <div class="li-type-normal" style="opacity: 1; display: block; transform: translateY(0px);">
                        <div class="post-header">
                            <div class="post-title">
                                <a class="post-title-link" href="/blog/html/{{.CreatedAt | getyear}}/{{.CreatedAt | getmonth}}/{{.Htmltag}}">
                                    {{.Title}}
                                </a>
                            </div>
                            <div class="post-meta">
                                <span class="post-time"> {{.CreatedAt | getmonth}} - {{.CreatedAt | getday}} </span>
                            </div>

                        </div>
                    </div>
                    {{end}}
                    <nav id="pagination">
                        <div class="clearfix"></div>
                    </nav>
                </div>


            </article>


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
<script src="/assets/js/jquery.simplePagination.js"></script>
<link type="text/css" rel="stylesheet" href="/assets/css/simplePagination.css"/>
<script type="text/javascript">
    var page_index;
    var itemsOnPage = 30;
    $(function() {
        $("#pagination").pagination({
            items: {{. | len}},
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
            $(".times ul li").hide();
            for(var i = page_index * itemsOnPage; i < page_index * itemsOnPage + itemsOnPage; i++){
                $(".times ul li:eq(" + i + ")").show();
            }
        }
    });
</script>
</body>
</html>