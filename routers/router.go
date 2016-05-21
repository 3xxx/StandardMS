package routers

import (
	"github.com/astaxie/beego"
	// "github.com/beego/admin" //admin 包
	"quick/controllers"
)

func init() {
	beego.Router("/admin", &controllers.UserController{}, "*:Admin")

	beego.Router("/category/uploadimages", &controllers.CategoryController{}, "*:Uploadimagesct")
	beego.Router("/controller", &controllers.UeditorController{}, "*:ControllerUE")
	// beego.Router("/controller", &controllers.UeditorController{}, "post:UploadImage")
	//除了<form里用了post，其他都要默认用get或*代替！！！！！
	beego.Router("/todo", &controllers.TaskController{}) //MainController{}
	beego.Router("/todo/task", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	beego.Router("/todo/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")
	beego.Router("/todo/addtask", &controllers.TaskController{}, "get:AddTask") //这里post是因为form里用了这个方法
	beego.Router("/todo/delete", &controllers.TaskController{}, "get:Delete")
	beego.Router("/todo/update", &controllers.TaskController{}, "get:Update")
	beego.Router("/todo/update1", &controllers.TaskController{}, "get:Update1")
	beego.Router("/todo/showdetails", &controllers.TaskController{}, "*:ShowDetails")

	// beego.Router("/login1", &controllers.IndexController{})//这个作废了
	beego.Router("/regist", &controllers.RegistController{})
	// beego.Router("/registerr", &controllers.RegistController{}, "get:RegistErr")
	beego.Router("/regist/checkuname", &controllers.RegistController{}, "post:CheckUname")
	beego.Router("/regist/getuname", &controllers.RegistController{}, "post:GetUname")

	// admin.Run()
	// beego.Router("/", &controllers.MainController{}, "*:Get")
	//	beego.Router("/hello", &controllers.MainController{}, "get:Hello")
	// beego.Router("/hello/:id([0-9]+)", &controllers.MainController{}, "get:Hello")
	// beego.Router("/hello/:id([0-9]+)", &controllers.MainController{}, "get,post:Get")
	// beego.Router("/manage/add", &controllers.ManageController{}, "get,post:Add")
	// beego.Router("/manage/view", &controllers.ManageController{}, "get:View")
	//	beego.Router("/manage/login", &controllers.ManageController{}, "*:home")
	// beego.Router("/manage/home", &controllers.ManageController{}, "*:Home")
	// beego.Router("/manage/show/:id([0-9]+)", &controllers.ManageController{}, "get:Show")
	// beego.Router("/manage/delete/:id([0-9]+)", &controllers.ManageController{}, "*:Delete")
	// beego.Router("/manage/update/:id([0-9]+)", &controllers.ManageController{}, "*:Update")
	// beego.Router("/images/", &controllers.ManageController{}, "*:Update")
	//注册beego路由
	beego.Router("/getspider", &controllers.SpiderController{}, "get:GetSpider") //路由也分大小写

	beego.Router("/", &controllers.MainController{})
	beego.Router("/help", &controllers.MainController{}, "*:Help")
	beego.Router("/test", &controllers.MainController{}, "*:Test")
	beego.Router("/test1", &controllers.MainController{}, "*:Test1")
	beego.Router("/test2", &controllers.MainController{}, "*:Test2")

	beego.Router("/topic/uploadimages", &controllers.TopicController{}, "*:Uploadimagesmd")

	// beego.Router("/", &controllers.MainController{})
	// beego.SetStaticPath("/ueditor", "ueditor")
	// beego.Router("/ue", &controllers.UEController{})
	// beego.AutoRouter(&controllers.UEController{})

	beego.Router("/catalog/import_xls_catalog", &controllers.CatalogController{}, "post:Import_Xls_Catalog")
	beego.Router("/catalog/add", &controllers.CatalogController{}, "get:Get")
	beego.Router("/catalog/view", &controllers.CatalogController{}, "get:View")
	//添加一行目录
	beego.Router("/catalog/post", &controllers.CatalogController{}, "post:Post")
	//修改一条目录
	beego.Router("/catalog/modifycatalog", &controllers.CatalogController{}, "post:ModifyCatalog")
	//删除一条目录
	beego.Router("/catalog/delete", &controllers.CatalogController{}, "post:Delete")

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/loginerr", &controllers.LoginController{}, "get:Loginerr")
	//导航栏本地搜索
	beego.Router("/search", &controllers.SearchController{})
	//水利院本地搜索
	beego.Router("/searchlocal", &controllers.SearchController{}, "post:Searchlocal")

	beego.Router("/category", &controllers.CategoryController{}) //相当于get
	//删除一个项目
	beego.Router("/category/delete", &controllers.CategoryController{}, "post:Delete")
	//添加项目视图第一步
	beego.Router("/category/add", &controllers.CategoryController{}, "get:Add")
	//添加项目视图第二步
	beego.Router("/category/add2", &controllers.CategoryController{}, "get:Add2")
	//添加项目第二步方法提交
	beego.Router("/category/post2", &controllers.CategoryController{}, "post:Post2")
	//添加项目第二步上传封面图片
	beego.Router("/category/addcoverphoto", &controllers.CategoryController{}, "post:AddCoverPhoto")
	//删除部分目录
	beego.Router("/category/deletecategory", &controllers.CategoryController{}, "get:DeleteCategory")
	//显示隐藏category
	beego.Router("/category/hidecategory", &controllers.CategoryController{}, "get:HideCategory")
	beego.Router("/category/showcategory", &controllers.CategoryController{}, "get:ShowCategory")
	//修改目录名称
	beego.Router("/category/modifytitle", &controllers.CategoryController{}, "post:ModifyCategoryTitle")

	//添加自定义目录视图第一步
	beego.Router("/category/add_b", &controllers.CategoryController{}, "get:Add_b")
	//添加自定义目录第一步提交方法
	beego.Router("/category/userdefinedpost", &controllers.CategoryController{}, "post:UserdefinedPost")
	//自定义项目（B）显示路由
	beego.Router("/category_b", &controllers.CategoryController{}, "get:Get_b")

	//添加项目第一步方法提交post 等等
	beego.AutoRouter(&controllers.CategoryController{}) //这句代替上句也行

	beego.Router("/category/view", &controllers.CategoryController{}, "get:View")
	beego.Router("/category/viewbrief", &controllers.CategoryController{}, "get:ViewBrief")
	//项目B模式显示指定成果类型里的成果
	beego.Router("/category/view_b", &controllers.CategoryController{}, "get:View_b")
	//iframe中的默认显示
	beego.Router("/category/category_prod_view", &controllers.CategoryController{}, "post:Category_prod_view")
	beego.Router("/category/modifycategory", &controllers.CategoryController{}, "post:ModifyCategory")

	//删除文章中的附件delete必须用get，为什么？
	beego.Router("/attachment/delete", &controllers.TopicController{}, "get:DeleteAttachment")

	//wiki页面
	beego.Router("/wiki", &controllers.WikiController{}) //get
	//发表文章界面
	beego.Router("/wiki/add", &controllers.WikiController{}, "get:Add")
	//发表文章提交
	beego.Router("/wiki/addwiki", &controllers.WikiController{}, "post:AddWiki")
	//查看一个文章
	beego.Router("/wiki/view/", &controllers.WikiController{}, "get:View")
	beego.AutoRouter(&controllers.WikiController{})

	//分页后修改成以下形式
	beego.Router("/topic", &controllers.TopicController{}, "*:ListAllPosts")
	// beego.Router("/topic/download", &controllers.TopicController{}, "get:Download")
	beego.Router("/topic/exporttoexcel", &controllers.TopicController{}, "post:ExportToExcel")
	//批量删除和批量下载
	beego.Router("/topic/downloadall", &controllers.TopicController{}, "post:DownloadAll")
	beego.Router("/topic/deleteall", &controllers.TopicController{}, "post:DeleteAll")
	//上面用post是因为<form style="float:left" method="post" action="/topic/deleteall"
	//一对多模式添加文章
	beego.Router("/topic/topic_many_add", &controllers.TopicController{}, "post:Topic_many_add")
	//一对一模式添加文章
	beego.Router("/topic/topic_one_add", &controllers.TopicController{}, "post:Topic_one_add")
	beego.Router("/topic/topic_one_addbaidu", &controllers.TopicController{}, "*:Topic_one_addbaidu")
	beego.Router("/topic/Topic_many_addbaidu", &controllers.TopicController{}, "*:Topic_many_addbaidu")
	//快捷上传里的文件处理
	beego.Router("/topic/topic_one_addstandard", &controllers.TopicController{}, "*:Topic_one_addstandard")
	//http://localhost:8081/topic/add?id=717&mid=3
	// case "3"://添加设代日记
	// c.TplName = "topic_add3.html"
	//'uploader':'/topic/addtopic1', 添加日记采用一对多模式
	//第二次添加日记的图片说明
	beego.Router("/topic/viewdiary", &controllers.TopicController{}, "get:ViewDiary") //显示附件上传后中间结果
	//存设代日记图片说明<form name="form" method="post" action="/topic/addtopic3">
	beego.Router("/topic/diary_add", &controllers.TopicController{}, "post:Diary_add")
	beego.Router("/topic/diary_second_add", &controllers.TopicController{}, "post:Diary_second_add")

	//第二次添加日记上图片的描述后显示的版面
	beego.Router("/topic/viewdiary1", &controllers.TopicController{}, "get:ViewDiary1") //显示最终日记效果
	//beego.Router("/topic/post", &controllers.TopicController{}, "post:Post")下面这句代替
	beego.AutoRouter(&controllers.TopicController{})
	//与beego.Router("/topic", &controllers.TopicController{})的区别：自动匹配和固定路由

	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
	//添加删除wiki的评论
	beego.Router("/reply/addwiki", &controllers.ReplyController{}, "post:AddWiki")
	beego.Router("/reply/deletewiki", &controllers.ReplyController{}, "get:DeleteWiki")

	beego.Router("/user/AddUser", &controllers.UserController{}, "*:AddUser")
	beego.Router("/user/UpdateUser", &controllers.UserController{}, "*:UpdateUser")
	beego.Router("/user/deluser", &controllers.UserController{}, "*:DelUser")
	beego.Router("/user/index", &controllers.UserController{}, "*:Index")
	beego.Router("/user/view", &controllers.UserController{}, "get:View")
	beego.Router("/user/view/*", &controllers.UserController{}, "get:View")
	beego.Router("/user/importexcel", &controllers.UserController{}, "post:ImportExcel")
	beego.Router("/user/getuserbyusername", &controllers.UserController{}, "get:GetUserByUsername")

	beego.Router("/role/AddAndEdit", &controllers.RoleController{}, "*:AddAndEdit")
	beego.Router("/role/DelRole", &controllers.RoleController{}, "*:DelRole")
	beego.Router("/role/AccessToNode", &controllers.RoleController{}, "*:AccessToNode")
	beego.Router("/role/AddAccess", &controllers.RoleController{}, "*:AddAccess")
	beego.Router("/role/RoleToUserList", &controllers.RoleController{}, "*:RoleToUserList")
	beego.Router("/role/AddRoleToUser", &controllers.RoleController{}, "*:AddRoleToUser")
	beego.Router("/role/Getlist", &controllers.RoleController{}, "*:Getlist")
	beego.Router("/role/index", &controllers.RoleController{}, "*:Index")
	beego.Router("/roleerr", &controllers.RoleController{}, "*:Roleerr") //显示权限不够

	beego.Router("/standard", &controllers.StandardController{}, "*:Index")
	beego.Router("/standard/search", &controllers.StandardController{}, "*:Search")
	beego.Router("/standard/importexcel", &controllers.StandardController{}, "*:ImportExcel")
	beego.Router("/standard/standard_one_addbaidu", &controllers.StandardController{}, "post:Standard_one_addbaidu")
	beego.Router("/standard/importlibrary", &controllers.StandardController{}, "post:ImportLibrary")

	//app.conf中要设置DirectoryIndex = true是否开启静态目录的列表显示，默认不显示目录，返回 403 错误
	beego.SetStaticPath("/m", "models")
	beego.SetStaticPath("/v", "views")
	beego.SetStaticPath("/c", "controllers")
	beego.SetStaticPath("/s", "static")
	beego.SetStaticPath("/t", "tests")
	beego.SetStaticPath("/a", "attachment")
	beego.SetStaticPath("/d", "database")
	// 作为静态文件：beego.SetStaticPath("/attachment", "attachment")
	//beego.Router("/attachment/:all", &controllers.AttachController{})
	beego.Router("/attachment/*", &controllers.AttachController{})
	// *全匹配方式 //匹配 /download/ceshi/file/api.json :splat=file/api.json

}
