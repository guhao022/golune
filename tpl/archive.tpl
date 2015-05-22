<!DOCTYPE HTML>
<html>
<head>
    <meta charset="utf-8">
    <title>文章归档 | The Pursuit of Happiness</title>
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
                <h1 class="ac_title" >文章归档</h1>
                <div class="history" style="padding:10px 80px 30px;">
                    {{range .}}
                    <ul>
                        <h2>{{.Key}}年</h2>
                        {{range .Val}}
                        <li class="posts-item">
                            <span class="datetime">{{.CreatedYear}}-{{.CreatedMonth}}-{{.CreatedDay}}</span>
                            <a href="/blog/html/{{.CreatedYear}}/{{.CreatedMonth}}/{{.Htmltag}}">{{.Title}}</a>
                        </li>
                        {{end}}
                    </ul>
                    {{end}}

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
</body>
</html>


