package controllers

import (
	// "bytes"
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	// "image/png"
	"io"
	// "log"
	"net/http"
	"os"
	"path"
	// "quick/models"
	"strings"
)

type UeditorController struct {
	beego.Controller
}

type UploadimageUE struct {
	url      string
	title    string
	original string
	state    string
	// "url": fmt.Sprintf("/static/upload/%s", filename),
	// "title": "demo.jpg",
	// "original": header.Filename,
	// "state": "SUCCESS"
}

// func (c *UeditorController) ControllerUE(w http.ResponseWriter, r *http.Request) {
// 	action := r.URL.Query()["action"][0]
// 	beego.Info(action)
// 	fmt.Println(r.Method, action)
// 	if r.Method == "GET" {
// 		if action == "config" {
// 			Configs(w, r)
// 		}
// 	} else if r.Method == "POST" {
// 		if action == "uploadimage" {
// 			UploadImage(w, r)
// 		}
// 	}
// }

func (c *UeditorController) ControllerUE() {
	// var configJson []byte // 当客户端请求 /ueditor/go/controller?action=config 返回的json内容
	// file, err := os.Open("conf/config.json")
	// if err != nil {
	// 	log.Fatal(err)
	// 	os.Exit(1)
	// }
	// defer file.Close()
	// buf := bytes.NewBuffer(nil)
	// buf.ReadFrom(file)
	// configJson = buf.Bytes()
	// title := c.Input().Get("action")
	op := c.Input().Get("action")
	switch op {
	case "config": //这里还是要优化成conf/config.json
		b := []byte(`{
              "imageActionName": "uploadimage", 
              "imageFieldName": "upfile", 
              "imageMaxSize": 2048000, 
              "imageAllowFiles": [".png", ".jpg", ".jpeg", ".gif", ".bmp"], 
              "imageCompressEnable": true,
              "imageCompressBorder": 1600, 
              "imageInsertAlign": "none", 
              "imageUrlPrefix": "", 
              "imagePathFormat": "/static/upload/{yyyy}{mm}{dd}/{time}{rand:6}"
        }`)
		var r interface{}
		json.Unmarshal(b, &r)
		// c.Data["json"] = map[string]interface{}{
		//    "imageUrl": "http://127.0.0.1/controller",
		//    "imagePath": "/attachment/test",
		//    "imageFieldName": "upfile",
		//    "imageMaxSize": 2048000,
		// "imageAllowFiles": [".png", ".jpg", ".jpeg", ".gif", ".bmp"]                                  //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
		// }
		c.Data["json"] = r
		c.ServeJson()
	case "uploadimage": //为什么这个没用？？？
		file, header, err := c.GetFile("upfile") // r.FormFile("upfile")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)
		err = os.MkdirAll(path.Join("static", "upload"), 0775)
		if err != nil {
			panic(err)
		}
		outFile, err := os.Create(path.Join("static", "upload", filename))
		if err != nil {
			panic(err)
		}
		defer outFile.Close()
		io.Copy(outFile, file)
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": "/static/upload/", "title": "111", "original": "demo.jpg"}
		c.ServeJson()

		// *     "state" => "",          //上传状态，上传成功时必须返回"SUCCESS"
		// *     "url" => "",            //返回的地址
		// *     "title" => "",          //新文件名
		// *     "original" => "",       //原始文件名
		// *     "type" => ""            //文件类型
		// *     "size" => "",           //文件大小
		// f := &UploadimageUE{
		// 	state:    "SUCCESS",
		// 	url:      fmt.Sprintf("/static/upload/%s", filename),
		// 	title:    "demo.jpg",
		// 	original: header.Filename,
		// }
		// c.Data["json"] = f
		// c.ServeJson()
		// 	reply := &Comment{
		// Tid:     tidNum,
		// Name:    nickname,
		// Content: content,

		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println(string(b))
		// w.Write(b)
		// 	c.Data["json"] = map[string]string{
		// 	"url":      fmt.Sprintf("/attachment/test/%s", h.Filename), //保存后的文件路径
		// 	"title":    "",                                             //文件描述，对图片来说在前端会添加到title属性上
		// 	"original": h.Filename,                                     //原始文件名
		// 	"state":    "SUCCESS",                                      //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
		// }
		// var s interface{}
		// json.Unmarshal(f, &s)
		// s, err := json.Marshal(f)
		// if err == nil {
		// c.Ctx.WriteString(string(f))

		// default:
	}

	// c.Write(configJson)
}

// func Configs(w http.ResponseWriter, r *http.Request) {
// 	w.Write(configJson)
// }

// var configJson []byte // 当客户端请求 /ueditor/go/controller?action=config 返回的json内容

// func init() {
// 	file, err := os.Open("conf/config.json")
// 	if err != nil {
// 		log.Fatal(err)
// 		os.Exit(1)
// 	}

// 	defer file.Close()
// 	buf := bytes.NewBuffer(nil)
// 	buf.ReadFrom(file)

// 	configJson = buf.Bytes()
// }

// func (c *UeditorController) UploadImage() {
// 	name := "111"    //c.Input().Get("name")
// 	number := "222"  //c.Input().Get("number")
// 	content := "333" //c.Input().Get("test-editormd-html-code")
// 	path := "c"      //c.Input().Get("tempString")

// 	diskdirectory := ".\\attachment\\" + "test" + "\\"
// 	url := "/attachment/" + "test" + "/"
// 	//保存上传的图片
// 	//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
// 	_, h, err := c.GetFile("upfile") //editormd-image-file
// 	beego.Info(h)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	// var attachment string
// 	// var path string
// 	var filesize int64
// 	var route string
// 	if h != nil {
// 		//保存附件
// 		path1 := ".\\attachment\\" + "test" + "\\" + h.Filename
// 		err = c.SaveToFile("upfile", path1) //editormd-image-file  .Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
// 		if err != nil {
// 			beego.Error(err)
// 		}

// 		if strings.Contains(strings.ToLower(h.Filename), ".jpg") { //ToLower转成小写
// 			// 随机名称
// 			// to := path + random_name() + ".jpg"
// 			origin := path1 //path + file.Name()
// 			fmt.Println("正在处理" + origin + ">>>" + origin)
// 			cmd_resize(origin, 2048, 0, origin)
// 			//				defer os.Remove(origin)//删除原文件
// 		}
// 		filesize, _ = FileSize(path1)
// 		filesize = filesize / 1000.0
// 		route = "/attachment/" + "test" + "/" + h.Filename
// 	} else {
// 		img := CreateRandomAvatar([]byte(number + name))
// 		fi, _ := os.Create("./attachment/" + "test" + "/u1.png")
// 		png.Encode(fi, img)
// 		fi.Close()
// 		route = "/attachment/" + "test" + "/u1.png"
// 	}

// 	uname := "4"

// 	//存入数据库
// 	_, err = models.AddCategory(name, number, content, path, route, uname, diskdirectory, url)
// 	if err != nil {
// 		beego.Error(err)
// 	} else {
// 		// f := Uploadimage{
// 		// 	url:     route,
// 		// 	success: 1,
// 		// 	message: "ok",
// 		// }
// 		// beego.Info(f)2016/01/17 01:40:03 [category.go:549] [I] {/attachment/test/u1.png ok 1}

// 		// c.Data["json"] = map[string]interface{}{"success": 1, "message": "111", "url": route}

// 		// c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": "111", "original": "demo.jpg", "url": route}
// 		// c.Data["json"] = f
// 		c.Data["json"] = map[string]string{
// 			"url":      fmt.Sprintf("/attachment/test/%s", h.Filename), //保存后的文件路径
// 			"title":    "",                                             //文件描述，对图片来说在前端会添加到title属性上
// 			"original": h.Filename,                                     //原始文件名
// 			"state":    "SUCCESS",                                      //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
// 		}
// 		c.ServeJson()
// 		// beego.Info(c.Data["json"])
// 		// 2016/01/17 01:42:00 [category.go:554] [I] map[success:1 message:111 url:/attachm
// 		// ent/test/u1.png]
// 		// 		{
// 		//     "state": "SUCCESS",
// 		//     "url": "upload/demo.jpg",
// 		//     "title": "demo.jpg",
// 		//     "original": "demo.jpg"
// 		//      }
// 	}

// 	// c.Data["Uname"] = ck.Value
// 	// id1 := strconv.FormatInt(id, 10)
// 	// c.Redirect("/category?op=view&id="+id1, 301)
// 	return //???
// }

func UploadImage(w http.ResponseWriter, r *http.Request) { //这个没用
	file, header, err := r.FormFile("upfile")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)
	err = os.MkdirAll(path.Join("static", "upload"), 0775)
	if err != nil {
		panic(err)
	}
	outFile, err := os.Create(path.Join("static", "upload", filename))
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	io.Copy(outFile, file)
	b, err := json.Marshal(map[string]string{
		"url":      fmt.Sprintf("/static/upload/%s", filename), //保存后的文件路径
		"title":    "",                                         //文件描述，对图片来说在前端会添加到title属性上
		"original": header.Filename,                            //原始文件名
		"state":    "SUCCESS",                                  //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	w.Write(b)
}

func (c *UeditorController) UploadImage() { //为什么用的是这个方法呢？
	file, header, err := c.GetFile("upfile") // r.FormFile("upfile")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)
	err = os.MkdirAll(path.Join("static", "upload"), 0775)
	if err != nil {
		panic(err)
	}
	outFile, err := os.Create(path.Join("static", "upload", filename))
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	io.Copy(outFile, file)
	c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": "/static/upload/" + filename, "title": "111", "original": "demo.jpg"}
	c.ServeJson()
	// "state": "SUCCESS",
	// "url": "upload/demo.jpg",
	// "title": "demo.jpg",
	// "original": "demo.jpg"
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(b))
	// w.Write(b)
	// 	c.Data["json"] = map[string]string{
	// 	"url":      fmt.Sprintf("/attachment/test/%s", h.Filename), //保存后的文件路径
	// 	"title":    "",                                             //文件描述，对图片来说在前端会添加到title属性上
	// 	"original": h.Filename,                                     //原始文件名
	// 	"state":    "SUCCESS",                                      //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
	// }
	// c.Data["json"] = b
	// c.ServeJson()
}
