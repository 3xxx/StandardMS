//2015-10-05这个函数最难，涉及到用goquery获取页面内容，beego的config解析json里的数组
//以及main.go里利用beego的toolbox里的task工具定时执行这个抓取页面的功能。
//用户只要在json.conf里定义好需要抓取的页面特性，ip和端口以及主人即可。
//走的弯路和坑：数据库表名最好不要用驼峰式，因为会自动转成蛇形，操作的时候要用蛇形表名才行。
//开始用go_spider来爬，研究了一天多时间，以为它提供schedule可以用来定时，其实不行。
//开始还用正则表达式来完成页面信息的提取，总是不成功。
//后面转到goquery，也是找了好久的资料，才知道find()对于#id，class，属性/值，src等的操作。
//最难的就是解析json返回数据是map[string]inteface{}类型，怎么也无法转成[]string.
//耗费了一天多时间.
package controllers

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"log"
	"quick/models"
	// "reflect"
	//	"strings"
	// "testing"
)

type Spider struct {
	Number string
	Name   string
	Link   string
}
type Service struct {
	Name string
	Link string
}
type SpiderController struct {
	beego.Controller
}

//显示首页
func (c *SpiderController) GetSpider() {
	// w http.ResponseWriter, r *http.Request
	// sess := globalSessions.SessionStart(w, r)
	// defer sess.SessionRelease()
	// username := sess.Get("username")
	// if r.Method == "GET" {
	// 	t, _ := template.ParseFiles("login.gtpl")
	// 	t.Execute(w, nil)
	// } else {
	// 	sess.Set("username", r.Form["username"])
	// }
	//（4）获取当前的请求会话，并返回当前请求会话的对象
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	//（5）根据当前请求对象，设置一个session
	// sess.Set("mySession", "qin.xc@gpdiwe.com;qq504284")
	// c.Data["Website"] = "广东省水利电力勘测设计研究院■☆●施工预算分院"
	//（6）从session中读取值
	// c.Data["Email"] = sess.Get("mySession")
	// c.Data["Website"] = "127.0.0.1:8080/hello"
	// c.Data["Email"] = "astaxie@gmail.com"
	// beego.Info(c.Ctx.Input.IP())

	c.Data["IsSpider"] = true
	c.TplName = "spider.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx) //大小写害死人！IsLogin
	//2.取得客户端用户名
	// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		c.Data["Uname"] = v.(string)
	}
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// } else {
	// 	c.Data["Uname"] = ck.Value
	// }

	// c.Data["Id"] = c.Ctx.Input.Param(":id")
	spidertopic, err := models.GetSpiderTopic()
	if err != nil {
		beego.Error(err)
	}
	//循环取出机器名称作为按钮展示。没必要，直接从json数据里取，然后判断是否在线。
	jsonconf, err := config.NewConfig("json", "./conf/testjsonWithArray.conf")
	// beego.Info("err!")
	if err != nil {
		beego.Info(err)
		log.Fatal(err)
	}

	rootArray, err := jsonconf.DIY("rootArray")
	if err != nil {
		beego.Error("array does not exist as element")
	}
	rootArrayCasted := rootArray.([]interface{})
	sservice := make([]Service, len(rootArrayCasted))
	if rootArrayCasted == nil {
		beego.Error("array from root is nil")
	} else {
		for index, _ := range rootArrayCasted {
			elem := rootArrayCasted[index].(map[string]interface{})
			sservice[index].Name = elem["username"].(string)
			sservice[index].Link = elem["serviceAPI"].(string) + ":" + elem["port"].(string)
		}
	}
	// for _, v := range spidertopic {
	// 	beego.Info(v.UserName)
	// }
	c.Data["Service"] = sservice
	c.Data["SpiderTopic"] = spidertopic
	spidercategory, err := models.GetSpiderCategory()
	if err != nil {
		beego.Error(err)
	}
	c.Data["SpiderCategory"] = spidercategory
}

//解析jason里的数组……得到ip地址后传给creatspider
func TestJsonStartsWithArray() {
	// beego.Info("err!")
	jsonconf, err := config.NewConfig("json", "./conf/testjsonWithArray.conf")
	if err != nil {
		beego.Info(err)
		log.Fatal(err)
	}

	rootArray, err := jsonconf.DIY("rootArray")
	if err != nil {
		beego.Error("array does not exist as element")
	}
	rootArrayCasted := rootArray.([]interface{})
	if rootArrayCasted == nil {
		beego.Error("array from root is nil")
	} else {
		CreatSpider(rootArrayCasted)
	}

	// elem := rootArrayCasted[0].(map[string]interface{})
	// if elem["url"] != "user" || elem["serviceAPI"] != "http://www.test.com/user" {
	// 	beego.Error("array[0] values are not valid")
	// }

	// elem2 := rootArrayCasted[1].(map[string]interface{})
	// if elem2["url"] != "employee" || elem2["serviceAPI"] != "http://www.test.com/employee" {
	// 	beego.Error("array[1] values are not valid")
	// }
}

//把ip传给ccrape，返回页面内容，存入数据库
func CreatSpider(rootArrayCasted []interface{}) {
	// url := []string{
	// 	"http://127.0.0.1:8081/topic/get",
	// 	"http://127.0.0.1:8081/category",
	// }
	for index, _ := range rootArrayCasted {
		elem := rootArrayCasted[index].(map[string]interface{})
		// fmt.Printf("\n", elem["port"].(string), elem["username"].(string), elem["serviceAPI"].(string))
		// strArray := make([]string, len(elem))
		// for _, arg := range elem {
		// 	strArray[1] = arg.(string)
		// 	fmt.Printf("\n", strArray[1])
		// }
		m := Scrape(elem["serviceAPI"].(string)+":"+elem["port"].(string), "/topic/get")
		user := elem["username"].(string)
		for _, v := range m {
			// models.AddSpider(v.Number, v.Name, v.Link, user, elem["serviceAPI"].(string)+":"+elem["port"].(string))
			// beego.Info(v.Name, v.Number)
			models.AddSpiderTopic(v.Number, v.Name, v.Link, user, elem["serviceAPI"].(string)+":"+elem["port"].(string))
		}
		m1 := Scrape(elem["serviceAPI"].(string)+":"+elem["port"].(string), "/category")
		for _, v := range m1 {
			// models.AddSpider(v.Number, v.Name, v.Link, user, elem["serviceAPI"].(string)+":"+elem["port"].(string))
			models.AddSpiderCategory(v.Number, v.Name, v.Link, user, elem["serviceAPI"].(string)+":"+elem["port"].(string))
		}
	}
}

//由地址返回页面内容
func Scrape(url, key string) (topic map[int]Spider) {
	doc, err := goquery.NewDocument(url + key)
	if err != nil {
		fmt.Println(err)
		return
	}
	m := map[int]Spider{}
	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		number := s.Find("#number").Text()
		name := s.Find("#name").Text() //s.Find("[title]").Text()
		link, _ := s.Find("a").Attr("href")
		f := Spider{
			Number: number,
			Name:   name,
			Link:   url + link,
		}
		//这里不用map，用结构体也行吗？如下
		//f:=make([]*Spider)
		//f[i].Number=
		//f[i].Name=
		//		fmt.Printf("%s\n", f)
		m[i] = f //2015-10-18之前这里的i用Number，导致编号相同的无法存入
		//		fmt.Printf("%s\n", m)
	})
	// beego.Info(m)
	return m
}

//这个map保留作为参考。
// func CreatSpider(url []string) {
// 	// url := []string{
// 	// 	"http://127.0.0.1:8081/topic/get",
// 	// 	"http://127.0.0.1:8081/category",
// 	// }
// 	for _, u := range url {
// 		//	m := Scrape("http://127.0.0.1:8081/topic", "")
// 		m := Scrape(u, "")
// 		for _, v := range m {
// 			fmt.Printf("编号: %s -名称： %s-地址：http://127.0.0.1:8081%s\n", v.Number, v.Name, v.Link)
// 			//编号: 788 -名称： 8-地址：http://127.0.0.1:8081/topic/view_b/186
// 			//编号: 455 -名称： 655-地址：http://127.0.0.1:8081/topic/view_b/183
// 			models.AddSpider(v.Number, v.Name, v.Link, "qingo", "127.0.0.1:8081")
// 		}
// 	}
// }

//func main() {
//	//	doc, err := goquery.NewDocument("http://127.0.0.1:8081/topic") //<span id="14_nwp" style="width: auto; height: auto; float: none;"><a id="14_nwl" href="http://cpro.baidu.com/cpro/ui/uijs.php?adclass=0&app_id=0&c=news&cf=1001&ch=0&di=128&fv=18&is_app=0&jk=380be00b331ad1a5&k=topic&k0=topic&kdi0=0&luki=2&n=10&p=baidu&q=74042097_cpr&rb=0&rs=1&seller_id=1&sid=a5d11a33be00b38&ssp2=1&stid=0&t=tpclicked3_hc&td=1989498&tu=u1989498&u=http%3A%2F%2Fblog%2Estudygolang%2Ecom%2Fcategory%2Fgo%5Fthird%5Fpkg%2F&urlid=0" target="_blank" mpid="14" style="text-decoration: none;"><span style="color:#0000ff;font-size:13.92px;width:auto;height:auto;float:none;">topic</span></a></span>s")
//	//	fmt.Println(doc)
//	//	if err != nil {
//	//		fmt.Println(err)
//	//		return
//	//		//		log.Fatal(err)
//	//	}
//	//	doc.Find(".topics .topic").Each(func(i int, contentSelection *goquery.Selection) {
//	//		title := contentSelection.Find(".title a").Text()
//	//		//		log.Println("第", i+1, "个帖子的标题：", title)
//	//		fmt.Println("第", i+1, "个帖子的标题：", title)
//	//	})
//	//	img1 := `<img src="assets/images/gallery/thumb-1.jpg" alt="150x150" />`
//	//	img2 := `<img alt="150x150" src="assets/images/gallery/thumb-1.jpg" />`
//	//	cases := []struct {
//	//		s string
//	//		l int
//	//	}{
//	//		{s: img1 + img2, l: 2},
//	//		{s: img1, l: 1},
//	//		{s: img2, l: 1},
//	//	}
//	//	for _, c := range cases {
//	//		doc, err := NewDocumentFromReader(strings.NewReader(c.s))
//	//		if err != nil {
//	//			t.Fatal(err)
//	//		}
//	//		sel := doc.Find("img[src]")
//	//		assertLength(t, sel.Nodes, c.l)
//	//	}
//	m := map[string][]string{}
//	numbers, names, titles := Scrape("http://127.0.0.1:8081/topic", "/get")
//	m["number"] = numbers
//	m["name"] = names //map[name:[原创 8 SL888-500-001
//	m["title"] = titles
//	fmt.Print(m)
//	//	for _, v := range numbers {
//	//		fmt.Printf("%s\n", v)
//	//	}
//	//	for _, v := range names {
//	//		fmt.Printf("%s\n", v)
//	//	}
//	//	for _, v := range titles {
//	//		fmt.Printf("%s\n", v)
//	//	}
//}

//func Scrape(url, key string) (numbers, names, titles []string) {
//	//1:对于category
//	//	doc, err := goquery.NewDocument("http://127.0.0.1:8081/category")
//	//	if err != nil {
//	//		fmt.Println(err)
//	//		return
//	//	}
//	//	//sel := Doc().Find(".selector")
//	//	//for i := range sel.Nodes {
//	//	//	single := sel.Eq(i)
//	//	//    // use `single` as a selection of 1 node
//	//	//}
//	//	var names []string
//	//	var titles []string
//	//	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
//	//		name := s.Find("#category").Text()
//	//		names = append(names, name)
//	//		title, _ := s.Find("a").Attr("href")
//	//		titles = append(titles, title)
//	//		fmt.Printf("Review %d: %s - %s\n", i, name, title)
//	//		//"WeName:%v link:http://127.0.0.1:8081%v \r\n"
//	//		//%v-值的默认格式表示;%s-直接输出字符串或者[]byte   fmt.Printf(title)
//	//	})
//	//	fmt.Printf("Review %d: %s - %s\n", names, titles)
//	//	for i, v := range names {
//	//		fmt.Printf("%d", i, v)
//	//	}
//	//2：对于topic
//	doc, err := goquery.NewDocument(url + key)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	//	var names []string
//	//	var titles []string
//	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
//		number := s.Find("#number").Text()
//		numbers = append(numbers, number)
//		name := s.Find("[title]").Text()
//		names = append(names, name)
//		title, _ := s.Find("a").Attr("href")
//		titles = append(titles, title)
//		//		fmt.Printf("Review %d: %s - %s\n", i, name, title)
//		//"WeName:%v link:http://127.0.0.1:8081%v \r\n"
//		//%v-值的默认格式表示;%s-直接输出字符串或者[]byte   fmt.Printf(title)
//	})
//	//	fmt.Printf("Review %d: %s - %s\n", names, titles)
//	//	for _, v := range names {
//	//		fmt.Printf(v)
//	//	}
//	return numbers, names, titles
//}

//package main

//import (
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"regexp"
//)

//var (
//	ptnIndexItem    = regexp.MustCompile(`<h4>\\s.*?<a.*?>(.*?)</a>\\s.*?</h4>`) //(`<a target="_blank" href="(.+\.html)" title=".+" >(.+)</a>`)
//	ptnContentRough = regexp.MustCompile(`(?s).*<div class="artcontent">(.*)<div id="zhanwei">.*`)
//	ptnBrTag        = regexp.MustCompile(`<br>`)
//	ptnHTMLTag      = regexp.MustCompile(`(?s)</?.*?>`)
//	ptnSpace        = regexp.MustCompile(`(^\s+)|( )`)
//)

//func Get(url string) (content string, statusCode int) {
//	resp, err1 := http.Get(url)
//	if err1 != nil {
//		statusCode = -100
//		return
//	}
//	defer resp.Body.Close()
//	data, err2 := ioutil.ReadAll(resp.Body)
//	if err2 != nil {
//		statusCode = -200
//		return
//	}
//	statusCode = resp.StatusCode
//	content = string(data)
//	return
//}

//type IndexItem struct {
//	url   string
//	title string
//}

//func findIndex(content string) (index []IndexItem, err error) {
//	matches := ptnIndexItem.FindAllStringSubmatch(content, 10000)
//	index = make([]IndexItem, len(matches))
//	for i, item := range matches {
//		index[i] = IndexItem{"http://127.0.0.1:8081/category_b" + item[1], item[2]}
//	}
//	return
//}

//func readContent(url string) (content string) {
//	raw, statusCode := Get(url)
//	if statusCode != 200 {
//		fmt.Print("Fail to get the raw data from", url, "\n")
//		return
//	}
//	match := ptnContentRough.FindStringSubmatch(raw)
//	if match != nil {
//		content = match[1]
//	} else {
//		return
//	}
//	content = ptnBrTag.ReplaceAllString(content, "\r\n")
//	content = ptnHTMLTag.ReplaceAllString(content, "")
//	content = ptnSpace.ReplaceAllString(content, "")
//	return
//}

//func main() {
//	fmt.Println(`Get index ...`)
//	s, statusCode := Get("http://127.0.0.1:8081/category_b")
//	if statusCode != 200 {
//		return
//	}

//	fmt.Println(s)

//	index, _ := findIndex(s)
//	fmt.Println(`Get contents and write to file ...`)
//	for _, item := range index {
//		fmt.Printf("Get content %s from %s and write to file.\n", item.title, item.url)
//		fileName := fmt.Sprintf("%s.txt", item.title)
//		content := readContent(item.url)
//		ioutil.WriteFile(fileName, []byte(content), 0644)
//		fmt.Printf("Finish writing to %s.\n", fileName)
//	}
//}

//这个也可以。只是正则表达式问题了。
//package main

//import (
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"regexp"
//)

////定义新的数据类型
//type Spider struct {
//	url    string
//	header map[string]string
//}

////定义 Spider的方法
//func (keyword Spider) get_html_header() string {
//	client := &http.Client{}
//	req, err := http.NewRequest("GET", keyword.url, nil)
//	if err != nil {
//	}
//	for key, value := range keyword.header {
//		req.Header.Add(key, value)
//	}
//	resp, err := client.Do(req)
//	if err != nil {
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//	}
//	return string(body)
//}

//func main() {
//	header := map[string]string{
//		"Host":       "http://127.0.0.1:8081/",
//		"Referer":    "http://127.0.0.1:8081/",
//		"DNT":        "1",
//		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36",
//		"Cookie":     "__huid=104rl%2B0HjG2YltBarbPPIz2w7HTbLrv43gETLeVtBdIEI%3D",
//	}
//	//	keyword := "187"
//	url := "http://127.0.0.1:8081/topic" //+ keyword
//	spider := &Spider{url, header}
//	html := spider.get_html_header()
//	fmt.Println(html)
//	rp1 := regexp.MustCompile("<tr>\\s.*?<th>\\s.*?<a.*?>(.*?)</a>\\s.*?</tr>") //<h4>\\s.*?<a.*?>(?P<name>.*?)</a>\\s.*?</h4>
//	find_txt := rp1.FindAllString(html, -1)
//	//	fmt.Println(find_txt)
//	for _, item := range find_txt {
//		fmt.Printf(item)
//		//			fmt.Printf("Get content %s from %s and write to file.\n", item.title, item.url)
//		//			fileName := fmt.Sprintf("%s.txt", item.title)
//		//			content := readContent(item.url)
//		//			ioutil.WriteFile(fileName, []byte(content), 0644)
//		//			fmt.Printf("Finish writing to %s.\n", fileName)
//	}
//	//	fmt.Println(html)
//}

//package main

//import (
//	"fmt"
//)

//type Fetcher interface {
//	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
//	Fetch(url string) (body string, urls []string, err error)
//}

//// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
//func Crawl(url string, depth int, fetcher Fetcher) {
//	// TODO: 并行的抓取 URL。
//	// TODO: 不重复抓取页面。
//	// 下面并没有实现上面两种情况：
//	if depth <= 0 {
//		return
//	}
//	body, urls, err := fetcher.Fetch(url)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Printf("found: %s %q\n", url, body)
//	for _, u := range urls {
//		Crawl(u, depth-1, fetcher)
//	}
//	return
//}

//func main() {
//	Crawl("http://127.0.0.1:8081/", 4, fetcher)
//}

//// fakeFetcher 是返回若干结果的 Fetcher。
//type fakeFetcher map[string]*fakeResult

//type fakeResult struct {
//	body string
//	urls []string
//}

//func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
//	if res, ok := (*f)[url]; ok {
//		return res.body, res.urls, nil
//	}
//	return "", nil, fmt.Errorf("not found: %s", url)
//}

//// fetcher 是填充后的 fakeFetcher。
//var fetcher = &fakeFetcher{
//	"http://127.0.0.1:8081/": &fakeResult{
//		"The Go Programming Language",
//		[]string{
//			"http://127.0.0.1:8081/topic/",
//			"http://127.0.0.1:8081/category/",
//		},
//	},
//	"http://127.0.0.1:8081/topic/": &fakeResult{
//		"Packages",
//		[]string{
//			"http://127.0.0.1:8081/topic",
//			"http://golang.org/cmd/",
//			"http://golang.org/pkg/fmt/",
//			"http://golang.org/pkg/os/",
//		},
//	},
//	"http://golang.org/pkg/fmt/": &fakeResult{
//		"Package fmt",
//		[]string{
//			"http://golang.org/",
//			"http://golang.org/pkg/",
//		},
//	},
//	"http://golang.org/pkg/os/": &fakeResult{
//		"Package os",
//		[]string{
//			"http://golang.org/",
//			"http://golang.org/pkg/",
//		},
//	},
//}

//改进
//package main

//import (
//  "fmt"
//)

//type Fetcher interface {
//        // Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
//  Fetch(url string) (body string, urls []string, err error)
//}
//var lockx = make(chan int,1)
//// 同步通信使用
//func LockFun(f func()) {
//  lockx<-1
//  f()
//  <-lockx
//}
//var visited map[string]bool = make(map[string]bool)
//// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
//func Crawl(url string, depth int, fetcher Fetcher, banner chan int) {

//  if depth <= 0 || visited[url] {
//    banner<-1
//    return
//  }
//  body, urls, err := fetcher.Fetch(url)
//  LockFun(func(){
//    visited[url]=true
//  })
//  fmt.Printf("found: %s %q\n", url, body)
//  if err != nil {
//    fmt.Println(err)
//    banner<-1
//    return
//  }
//  subBanner := make(chan int, len(urls))
//  for _, u := range urls {
//     // 并行吧～～
//      go Crawl(u, depth-1, fetcher, subBanner);
//  }
//  for i:=0; i < len(urls); i++ {
//    // subBanner用来防止退出
//    <-subBanner
//  }
//  // banner用于让父节点退出
//  banner<-1
//  return
//}

//func main() {
//  mainBanner := make(chan int,1)
//  Crawl("http://golang.org/", 4, fetcher, mainBanner)
//  <-mainBanner
//}

//// fakeFetcher 是返回若干结果的 Fetcher。
//type fakeFetcher map[string]*fakeResult

//type fakeResult struct {
//  body string
//  urls     []string
//}

//func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
//  if res, ok := (*f)[url]; ok {
//    return res.body, res.urls, nil
//  }
//  return "", nil, fmt.Errorf("not found: %s", url)
//}

//// fetcher 是填充后的 fakeFetcher。
//var fetcher = &fakeFetcher{
//  "http://golang.org/": &fakeResult{
//    "The Go Programming Language",
//    []string{
//      "http://golang.org/pkg/",
//      "http://golang.org/cmd/",
//    },
//  },
//  "http://golang.org/pkg/": &fakeResult{
//    "Packages",
//    []string{
//      "http://golang.org/",
//      "http://golang.org/cmd/",
//      "http://golang.org/pkg/fmt/",
//      "http://golang.org/pkg/os/",
//    },
//  },
//  "http://golang.org/pkg/fmt/": &fakeResult{
//    "Package fmt",
//    []string{
//      "http://golang.org/",
//      "http://golang.org/pkg/",
//    },
//  },
//  "http://golang.org/pkg/os/": &fakeResult{
//    "Package os",
//    []string{
//      "http://golang.org/",
//      "http://golang.org/pkg/",
//    },
//  },
//}
