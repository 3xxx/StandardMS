package controllers

import (
	// m "github.com/beego/admin/src/models"
	// "github.com/astaxie/beego/orm"
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
	// "io"
	"fmt"
	"net"
	"os"
	m "quick/models"
	// "regexp"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	beego.Controller
}

// 1 init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
// 2 每个包可以拥有多个init函数
// 3 包的每个源文件也可以拥有多个init函数
// 4 同一个包中多个init函数的执行顺序go语言没有明确的定义(说明)
// 5 不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序
// 6 init函数不能被其他函数调用，而是在main函数执行之前，自动被调用
//读取iprole.txt文件，作为全局变量Iprolemaps，供调用访问者ip的权限用
var (
	Iprolemaps map[string]int
)

func init() {
	Iprolemaps = make(map[string]int)
	f, err := os.OpenFile("./conf/iprole.txt", os.O_RDONLY, 0660)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s err read from %s : %s\n", err)
	}
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		maps := processFlag(args)
		for i, v := range maps {
			Iprolemaps[i] = v
		}
	}
	f.Close()
}

func (c *UserController) Admin() {
	c.TplName = "AdminLTE/index.html"
}

func (c *UserController) Index() {
	var rolename int
	var uname string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, role, _ := checkRolewrite(c.Ctx) //login里的
	if role != "0" {
		rolename, _ = strconv.Atoi(role)
		c.Data["Uname"] = uname
	} else {
		port := strconv.Itoa(c.Ctx.Input.Port())
		route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		return
	}
	// if filetype != "pdf" && filetype != "jpg" && filetype != "diary" {
	if rolename > 1 { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// }

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//2.取得客户端用户名
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// } else {
	// 	c.Data["Uname"] = ck.Value
	// }

	users, count := m.Getuserlist(1, 2000, "Id")
	if c.IsAjax() {
		c.Data["json"] = &map[string]interface{}{"total": count, "rows": &users}
		c.ServeJSON()
		return
	} else {
		// tree := c.GetTree()
		// c.Data["tree"] = &tree
		c.Data["Users"] = &users
		c.TplName = "AdminLTE/users_import.html"
		// c.TplName = "user.tpl"
		// if c.GetTemplatetype() != "easyui" {
		// c.Layout = c.GetTemplatetype() + "/public/layout.tpl"
		// }
		// c.TplName = c.GetTemplatetype() + "/rbac/user.tpl"
	}

}

func (c *UserController) View() {
	var rolename int
	// c.Data["IsCategory"] = true
	// c.TplName = "category.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, role, _ := checkRolewrite(c.Ctx) //login里的
	if role != "0" {
		rolename, _ = strconv.Atoi(role)
		c.Data["Uname"] = uname
	} else {
		port := strconv.Itoa(c.Ctx.Input.Port())
		route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		return
	}
	// if filetype != "pdf" && filetype != "jpg" && filetype != "diary" {
	if rolename > 1 { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}

	userid, _ := strconv.ParseInt(c.Input().Get("useid"), 10, 64)
	// beego.Info(userid)
	user := m.GetUserByUserId(userid)
	// if c.IsAjax() {
	// users, _ := m.Getuserlist(1, 1000, "Id")
	list, _ := m.GetRoleByUserId(userid)
	// if err != nil {
	// 	beego.Error(err)
	// 	c.Redirect("/", 302)
	// 	return
	// }
	// beego.Info(list[1])
	// for i := 0; i < len(users); i++ {
	// 	for x := 0; x < len(list); x++ {
	// 		if users[i]["Id"] == list[x]["Id"] {
	// 			users[i]["checked"] = 1
	// 		}
	// 	}
	// }
	// if len(users) < 1 {
	// 	users = []orm.Params{}
	// }
	// c.Data["json"] = &map[string]interface{}{"total": count, "rows": &users}
	// c.ServeJSON()
	// return
	// } else {
	c.Data["User"] = user
	c.Data["Role"] = list
	// c.Data["Users"] = &users
	c.TplName = "AdminLTE/user_view.html"
	// c.TplName = "admin_user_view.tpl"
	// c.TplName = c.GetTemplatetype() + "/rbac/roletouserlist.tpl"
	// }
}

func (c *UserController) AddUser() {
	u := m.User{}
	if err := c.ParseForm(&u); err != nil {
		//handle error
		// c.Rsp(false, err.Error())
		beego.Error(err.Error)
		return
	}
	id, err := m.AddUser(&u)
	if err == nil && id > 0 {
		// c.Rsp(true, "Success")
		return
	} else {
		// c.Rsp(false, err.Error())
		beego.Error(err.Error)
		return
	}

}

// func (c *UserController) UpdateUser() {
// 	u := m.User{}
// 	if err := c.ParseForm(&u); err != nil {
// 		//handle error
// 		// c.Rsp(false, err.Error())
// 		beego.Error(err.Error)
// 		return
// 	}
// 	id, err := m.UpdateUser(&u)
// 	if err == nil && id > 0 {
// 		// c.Rsp(true, "Success")
// 		return
// 	} else {
// 		// c.Rsp(false, err.Error())
// 		beego.Error(err.Error)
// 		return
// 	}

// }
func (c *UserController) UpdateUser() {

	userid := c.Input().Get("userid")
	// username := c.Input().Get("username")
	nickname := c.Input().Get("nickname")
	email := c.Input().Get("email")
	Pwd1 := c.Input().Get("password")
	if Pwd1 != "" {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(Pwd1))
		cipherStr := md5Ctx.Sum(nil)
		// fmt.Print(cipherStr)
		// fmt.Print("\n")
		// fmt.Print(hex.EncodeToString(cipherStr))

		password := hex.EncodeToString(cipherStr)
		// user.Lastlogintime = time.Now()
		err := m.UpdateUser(userid, nickname, email, password) //这里修改
		if err != nil {
			beego.Error(err)
		}
		//更新role
		roleid := c.Input().Get("roletitle1")
		if roleid != "" {
			roleid1, _ := strconv.ParseInt(roleid, 10, 64)
			roleid2, _ := strconv.ParseInt(c.Input().Get("roletitle2"), 10, 64)
			userid1, _ := strconv.ParseInt(userid, 10, 64)
			_, err = m.UpdateRoleUser(roleid1, roleid2, userid1)
			if err != nil {
				beego.Error(err)
			}
		}
	} else {
		err := m.UpdateUser(userid, nickname, email, "") //这里修改
		if err != nil {
			beego.Error(err)
		}
		//更新role
		roleid := c.Input().Get("roletitle1")
		if roleid != "" {
			roleid1, _ := strconv.ParseInt(roleid, 10, 64)
			roleid2, _ := strconv.ParseInt(c.Input().Get("roletitle2"), 10, 64)
			userid1, _ := strconv.ParseInt(userid, 10, 64)
			_, err = m.UpdateRoleUser(roleid1, roleid2, userid1)
			if err != nil {
				beego.Error(err)
			}
		}
	}
	c.TplName = "user_view.tpl"
}

func (c *UserController) DelUser() {
	Id, _ := c.GetInt64("userid")
	status, err := m.DelUserById(Id)
	if err == nil && status > 0 {
		// c.Rsp(true, "Success")
		c.Redirect("/user/index", 302)
		return
	} else {
		// c.Rsp(false, err.Error())
		beego.Error(err.Error)
		return
	}
}

func (c *UserController) GetUserByUsername() {
	var rolename int
	// 	c.Data["IsCategory"] = true
	// c.TplName = "category.tpl"
	username := c.Input().Get("username")
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, role, _ := checkRolewrite(c.Ctx) //login里的
	if role != "0" {
		rolename, _ = strconv.Atoi(role)
		c.Data["Uname"] = uname
	} else {
		port := strconv.Itoa(c.Ctx.Input.Port())
		route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		return
	}
	// if filetype != "pdf" && filetype != "jpg" && filetype != "diary" {
	if rolename > 1 && uname != username { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}

	user := m.GetUserByUsername(username)
	list, _, _ := m.GetRoleByUsername(username)
	c.Data["User"] = user
	c.Data["Role"] = list
	c.TplName = "user_view.tpl"
}

//上传excel文件，导入到数据库
//引用来自category的查看成果类型里的成果
func (c *UserController) ImportExcel() {
	//解析表单
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	// id := c.Input().Get("id")
	// path := c.Input().Get("path")
	// filename := c.Input().Get("filename")

	//获取上传的文件
	_, h, err := c.GetFile("excel")
	if err != nil {
		beego.Error(err)
	}
	// var attachment string
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		// attachment = h.Filename
		// beego.Info(attachment)
		path = ".\\attachment\\" + h.Filename

		// path := c.Input().Get("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("excel", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
	}
	// if title == "" || tnumber == "" {
	// 	//将附件的编号和名称写入数据库
	// 	filename1, filename2 := SubStrings(attachment)
	// 	if filename1 == "" {
	// 		filename1 = filename2 //如果编号为空，则用文件名代替，否则多个编号为空导致存入数据库唯一性检查错误
	// 	}
	// 	title = filename2
	// 	tnumber = filename1
	// }
	// var err error
	// var tid string //这里是增加的，不知为何教程没有
	// path := ".\\attachment\\" + categoryproj.Number + " " + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename
	// ck, err := c.Ctx.Request.Cookie("uname")
	if err != nil {
		beego.Error(err)
	}
	// uname := ck.Value

	// route := "/attachment/" + categoryproj.Number + " " + categoryproj.Title + "/" + categoryphase.Title + "/" + categoryspec.Title + "/" + category + "/" + h.Filename
	//Catalogid := c.Input().Get("Catalogid")
	var user m.User

	//读出excel内容写入数据库
	// excelFileName := path                    //"/home/tealeg/foo.xlsx"
	xlFile, err := xlsx.OpenFile(path) //excelFileName
	if err != nil {
		beego.Error(err)
	}
	j := 0 //表格第一列从0开始
	var i int
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			// for i, cell := range row.Cells { //第一行从0开始，过滤掉第一行表头
			if i != 0 { //忽略第一行
				// len(row.Cells) //总列数是从1开始
				user.Username, _ = row.Cells[j+1].String() //第一列从0开始，第二列是用户名qin.xc
				Pwd1, _ := row.Cells[j+2].String()
				md5Ctx := md5.New()
				md5Ctx.Write([]byte(Pwd1))
				cipherStr := md5Ctx.Sum(nil)
				user.Password = hex.EncodeToString(cipherStr)
				user.Email, _ = row.Cells[j+3].String()
				user.Nickname, _ = row.Cells[j+4].String()
				user.Department, _ = row.Cells[j+6].String()
				user.Secoffice, _ = row.Cells[j+7].String()
				user.Lastlogintime = time.Now()
				uid, err := m.SaveUser(user)
				role, _ := row.Cells[j+5].String() //这里应该是由rolename查询roleid
				roleid, _ := strconv.ParseInt(role, 10, 64)
				_, err = m.AddRoleUser(roleid, uid)
				if err != nil {
					beego.Error(err)
				}
			}
			// }
			// for _, cell := range row.Cells {
			// 	fmt.Printf("%s\n", cell.String())

			// }
		}
	}
}

//取得访问者的权限
func Getiprole(ip string) (role int) {
	role = Iprolemaps[ip]
	return role
}

//获取下一个IP
func nextIp(ip string) string {
	ips := strings.Split(ip, ".")
	var i int
	for i = len(ips) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(ips[i])
		if n >= 255 {
			//进位
			ips[i] = "1"
		} else {
			//+1
			n++
			ips[i] = strconv.Itoa(n)
			break
		}
	}
	if i == -1 {
		//全部IP段都进行了进位,说明此IP本身已超出范围
		return ""
	}
	ip = ""
	leng := len(ips)
	for i := 0; i < leng; i++ {
		if i == leng-1 {
			ip += ips[i]
		} else {
			ip += ips[i] + "."
		}
	}
	return ip
}

//生成IP地址列表
func processIp(startIp, endIp string) []string {
	var ips = make([]string, 0)
	for ; startIp != endIp; startIp = nextIp(startIp) {
		if startIp != "" {
			ips = append(ips, startIp)
		}
	}
	ips = append(ips, startIp)
	return ips
}

func processFlag(arg []string) (maps map[string]int) {
	//开始IP,结束IP
	var startIp, endIp string
	//端口
	var ports []int = make([]int, 0)
	index := 0
	startIp = arg[index]
	si := net.ParseIP(startIp)
	if si == nil {
		//开始IP不合法
		fmt.Println("'startIp' Setting error")
		return nil
	}
	index++
	endIp = arg[index]
	ei := net.ParseIP(endIp)
	if ei == nil {
		//未指定结束IP,即只扫描一个IP
		endIp = startIp
	} else {
		index++
	}

	tmpPort := arg[index]
	if strings.Index(tmpPort, "-") != -1 {
		//连续端口
		tmpPorts := strings.Split(tmpPort, "-")
		var startPort, endPort int
		var err error
		startPort, err = strconv.Atoi(tmpPorts[0])
		if err != nil || startPort < 1 || startPort > 65535 {
			//开始端口不合法
			return nil
		}
		if len(tmpPorts) >= 2 {
			//指定结束端口
			endPort, err = strconv.Atoi(tmpPorts[1])
			if err != nil || endPort < 1 || endPort > 65535 || endPort < startPort {
				//结束端口不合法
				fmt.Println("'endPort' Setting error")
				return nil
			}
		} else {
			//未指定结束端口
			endPort = 65535
		}
		for i := 0; startPort+i <= endPort; i++ {
			ports = append(ports, startPort+i)
		}
	} else {
		//一个或多个端口
		ps := strings.Split(tmpPort, ",")
		for i := 0; i < len(ps); i++ {
			p, err := strconv.Atoi(ps[i])
			if err != nil {
				//端口不合法
				fmt.Println("'port' Setting error")
				return nil
			}
			ports = append(ports, p)
		}
	}

	//生成扫描地址列表
	ips := processIp(startIp, endIp)
	il := len(ips)
	m1 := make(map[string]int)
	for i := 0; i < il; i++ {
		pl := len(ports)
		for j := 0; j < pl; j++ {
			//			ipAddrs <- ips[i] + ":" + strconv.Itoa(ports[j])
			//			ipAddrs := ips[i] + ":" + strconv.Itoa(ports[j])
			m1[ips[i]] = ports[j]
		}
	}
	//	fmt.Print(slice1)
	return m1
	//	close(ipAddrs)
}
