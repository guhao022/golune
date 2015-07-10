package apis

import (
	"golune/conf"
	"golune/models"
	"golune/utils"

	"fmt"
	"html/template"
	"os"
	"sort"
	"strconv"
	"time"
)

type OutputCacheHandler struct {
	BaseHandler
}

func Unescaped(x string) interface{} {
	return template.HTML(x)
}

func cateIcon(cate string) (icon string) {
	switch cate {
	case "technology":
		icon = "fa-rocket"
	case "life":
		icon = "fa-bus"
	case "talk":
		icon = "fa-leaf"
	default:
		icon = "fa-file-text-o"
	}
	return
}

//缓存博客页面
func (c *OutputCacheHandler) CacheBlogHtml() {

	for _, blog := range GetBlog() { //循环获取的所有blog
		//把tmplate解析成string
		__bloghtml := utils.ParseTmplateToStr("tpl/blog.tpl")
		//设置文章缓存地址   年+月+blogtag
		bloghtmlpath := conf.HtmlPath + strconv.Itoa(blog.CreatedYear) + "/" + strconv.Itoa(blog.CreatedMonth) + "/"
		//创建文件
		path, err := utils.CreateFile(bloghtmlpath)
		if err != nil {
			c.CustomAbort(403, fmt.Sprintln(err))
		}

		htmlname := path + blog.Htmltag      //blog页面保存的路径
		htmlfile, err := os.Create(htmlname) //创建页面
		if err != nil {
			c.CustomAbort(403, fmt.Sprintln(err))
		}
		defer htmlfile.Close()
		t := template.New("bloghtml" + string(utils.RandomCreateBytes(5)))
		t.Funcs(template.FuncMap{"unescaped": Unescaped, "id2str": utils.Id2Str, "unix2str": utils.Unix2Str, "replacecate": utils.ReplaceCate, "markdownCommon": utils.MarkdownCommon, "cateIcon": cateIcon})
		t.Parse(__bloghtml)
		err = t.Execute(htmlfile, blog)
		if err != nil {
			c.CustomAbort(403, fmt.Sprintln(err))
		}
		c.Data["json"] = map[string]int{"Recode": 200}
	}
	c.ServeJson()
}

//根据ID缓存博客
func (c *OutputCacheHandler) CacheBlogHtmlById() {
	mgo, err := models.NewDB()
	if err != nil {
		c.Ctx.Redirect(500, fmt.Sprintln(err))
	}
	defer mgo.Close()
	blogid := c.GetString("Id")
	blog, _ := mgo.FindBlogById(blogid)

	createTime := time.Unix(blog.CreatedAt, 0)
	createYear := createTime.Year()
	createMonth := utils.Month2Int(createTime.Month())
	//createDay := createTime.Day()

	__bloghtml := utils.ParseTmplateToStr("tpl/blog.tpl")
	//设置文章缓存地址   年+月+blogtag
	bloghtmlpath := conf.HtmlPath + strconv.Itoa(createYear) + "/" + strconv.Itoa(createMonth) + "/"
	//创建文件
	path, err := utils.CreateFile(bloghtmlpath)
	if err != nil {
		c.CustomAbort(403, fmt.Sprintln(err))
	}

	htmlname := path + blog.Htmltag      //blog页面保存的路径
	htmlfile, err := os.Create(htmlname) //创建页面
	if err != nil {
		c.CustomAbort(403, fmt.Sprintln(err))
	}
	defer htmlfile.Close()
	t := template.New("bloghtml" + string(utils.RandomCreateBytes(5)))
	t.Funcs(template.FuncMap{"unescaped": Unescaped, "id2str": utils.Id2Str, "unix2str": utils.Unix2Str, "replacecate": utils.ReplaceCate, "markdownCommon": utils.MarkdownCommon, "cateIcon": cateIcon})
	t.Parse(__bloghtml)
	err = t.Execute(htmlfile, blog)
	if err != nil {
		c.CustomAbort(403, fmt.Sprintln(err))
	}
	c.Data["json"] = map[string]int{"Recode": 200}
	c.ServeJson()
}

//缓存首页
func (c *OutputCacheHandler) CacheIndexHtml() {
	//获取blog
	outBlog := GetBlog()

	//获取tag
	outTag := GetTags()

	//获取文章归档
	//outArchive := Archiving()
	__indexhtml := utils.ParseTmplateToStr("tpl/index.tpl")
	htmlfile, err := os.Create(conf.HtmlPath + "index.html")
	if err != nil {
		c.CustomAbort(403, fmt.Sprintln(err))
	}
	defer htmlfile.Close()

	out := map[string]interface{}{"blog": outBlog, "tag": outTag}

	t := template.New("bloghtml" + string(utils.RandomCreateBytes(5)))
	t.Funcs(template.FuncMap{"unescaped": Unescaped, "id2str": utils.Id2Str, "unix2str": utils.Unix2Str, "replacecate": utils.ReplaceCate, "markdownCommon": utils.MarkdownCommon, "cateIcon": cateIcon})
	t.Parse(__indexhtml)
	err = t.Execute(htmlfile, out)
	if err != nil {
		c.CustomAbort(403, fmt.Sprintln(err))
	}

}

//缓存分类文章页面
func (c *OutputCacheHandler) CacheCateHtml() {
	mgo, err := models.NewDB()
	if err != nil {
		c.Ctx.Redirect(500, fmt.Sprintln(err))
	}
	defer mgo.Close()
	cate := conf.Cate
	for k, _ := range cate {
		blogs, _ := mgo.FindBlogByCate(k)
		//把tmplate解析成string
		__catehtml := utils.ParseTmplateToStr("tpl/cate.tpl")
		//设置cate缓存地址
		catehtmlpath := conf.HtmlPath + "cate/"
		//创建文件
		path, err := utils.CreateFile(catehtmlpath)
		if err != nil {
			c.CustomAbort(403, fmt.Sprintln(err))
		}

		htmlname := path + k                 //cate页面保存的路径
		htmlfile, err := os.Create(htmlname) //创建页面
		if err != nil {
			c.CustomAbort(403, fmt.Sprintln(err))
		}
		defer htmlfile.Close()
		t := template.New("catehtml" + string(utils.RandomCreateBytes(5)))
		t.Funcs(template.FuncMap{"id2str": utils.Id2Str, "replacecate": utils.ReplaceCate, "getyear": utils.GetYear, "getmonth": utils.GetMonth, "getday": utils.GetDay})
		t.Parse(__catehtml)

		out := map[string]interface{}{"blog": blogs, "cate": k}

		err = t.Execute(htmlfile, out)
		if err != nil {
			c.CustomAbort(403, fmt.Sprintln(err))
		}
		c.Data["json"] = map[string]int{"Recode": 200}
	}
	c.ServeJson()
}

//缓存标签云页面
func (c *OutputCacheHandler) CacheTagsHtml() {
	mgo, err := models.NewDB()
	if err != nil {
		c.Ctx.Redirect(500, fmt.Sprintln(err))
	}
	defer mgo.Close()

	tags, err := mgo.GetTags()
	if err != nil {
		c.Ctx.Redirect(500, fmt.Sprintln(err))
	}

	//设置cate缓存地址
	tagshtmlpath := conf.HtmlPath + "tags/"
	//创建文件
	path, err := utils.CreateFile(tagshtmlpath)
	if err != nil {
		c.CustomAbort(403, fmt.Sprintln(err))
	}

	__taghtml := utils.ParseTmplateToStr("tpl/tags.tpl")
	htmlfile, err := os.Create(path + "index.html")
	if err != nil {
		c.CustomAbort(403, fmt.Sprintln(err))
	}
	defer htmlfile.Close()

	t := template.New("tagshtml" + string(utils.RandomCreateBytes(5)))
	t.Funcs(template.FuncMap{"unescaped": Unescaped, "id2str": utils.Id2Str, "unix2str": utils.Unix2Str, "replacecate": utils.ReplaceCate, "markdownCommon": utils.MarkdownCommon, "cateIcon": cateIcon})
	t.Parse(__taghtml)
	err = t.Execute(htmlfile, tags)
	if err != nil {
		c.CustomAbort(403, fmt.Sprintln(err))
	}
}

func (c *OutputCacheHandler) CacheArchiveHtml() {
	archiveBlog := Archiving()
	//设置cate缓存地址
	archivehtmlpath := conf.HtmlPath + "archive/"
	//创建文件
	path, err := utils.CreateFile(archivehtmlpath)
	if err != nil {
		c.CustomAbort(403, fmt.Sprintln(err))
	}
	__archivehtml := utils.ParseTmplateToStr("tpl/archive.tpl")
	htmlfile, err := os.Create(path + "index.html")
	if err != nil {
		c.CustomAbort(403, fmt.Sprintln(err))
	}
	defer htmlfile.Close()

	t := template.New("bloghtml" + string(utils.RandomCreateBytes(5)))
	t.Funcs(template.FuncMap{"unescaped": Unescaped, "id2str": utils.Id2Str, "unix2str": utils.Unix2Str})
	t.Parse(__archivehtml)
	err = t.Execute(htmlfile, archiveBlog)
	if err != nil {
		c.Data["json"] = map[string]string{"Recode": "400", "Error": fmt.Sprintln(err)}
	} else {
		c.Data["json"] = map[string]string{"Recode": "200"}
	}
	c.ServeJson()
}

//文章归档
func Archiving() interface{} {
	var blogyear = make(map[string]interface{})
	var endYear int = time.Now().Year()
	for i := endYear; i >= conf.ArchiveStartYear; i-- {
		blogyear[strconv.Itoa(i)] = GetBlogByYear(i)
	}
	ms := utils.NewMapSorter(blogyear)
	sort.Sort(ms)
	return ms
}

//根据年份获取blog
func GetBlogByYear(year int) interface{} {
	blogs := GetBlog()
	var blogmap []models.Blog
	for _, blog := range blogs { //循环所有的博客  获取单个博客信息
		if blog.CreatedYear == year {
			blogmap = append(blogmap, blog)
		}
	}
	return blogmap
}

//获取所有的tag
func GetTags() (outTag []models.Tag) {
	var c *BaseHandler
	var err error
	mgo, err := models.NewDB()
	if err != nil {
		c.Ctx.Redirect(500, fmt.Sprintln(err))
	}
	defer mgo.Close()

	if outTag, err = mgo.GetTags(); err != nil {
		c.Ctx.Redirect(500, fmt.Sprintln(err))
	}
	return
}

//获取所有的blog
func GetBlog() (outBlog []models.Blog) {
	var c *BaseHandler
	mgo, err := models.NewDB()
	if err != nil {
		c.Ctx.Redirect(500, fmt.Sprintln(err))
	}
	defer mgo.Close()
	if blogs, err := mgo.FindAllBlog(); err == nil {
		for _, blog := range blogs {
			//解析分类
			for key, val := range conf.Cate {
				if blog.Category == key {
					blog.CateBox = val
				}
			}

			//查找tag
			tags, err := mgo.FindBlogTagByBlog(blog.Id)
			if err != nil {
				c.CustomAbort(403, fmt.Sprintln(err))
			}
			blog.TagBox = tags

			//解析时间
			createTime := time.Unix(blog.CreatedAt, 0)
			createYear := createTime.Year()
			createMonth := utils.Month2Int(createTime.Month())
			createDay := createTime.Day()
			blog.CreatedYear = createYear
			blog.CreatedMonth = createMonth
			blog.CreatedDay = createDay
			outBlog = append(outBlog, blog)
		}
		return
	}

	return nil
}
