package controllers

import (
	// "encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	// m "github.com/beego/admin/src/models"
	m "hydrocms/models"
)

type RoleController struct {
	beego.Controller
}

func (c *RoleController) Index() {
	roles, count := m.GetRolelist(1, 100, "Id")
	if c.IsAjax() {
		// page, _ := c.GetInt64("page")
		// page_size, _ := c.GetInt64("rows")
		// sort := c.GetString("sort")
		// order := c.GetString("order")
		// if len(order) > 0 {
		// 	if order == "desc" {
		// 		sort = "-" + sort
		// 	}
		// } else {
		// 	sort = "Id"
		// }
		if len(roles) < 1 {
			roles = []orm.Params{}
		}
		c.Data["Json"] = &map[string]interface{}{"total": count, "rows": &roles}
		c.ServeJSON()
		return
	} else {
		// c.TplName = c.GetTemplatetype() + "/rbac/role.tpl"
		c.Data["Roles"] = &roles
		c.TplName = "role.tpl"
	}
}

func (c *RoleController) Roleerr() {
	// url := c.Input().Get("url")
	url1 := c.Input().Get("url") //这里不支持这样的url，http://192.168.9.13/login?url=/topic/add?id=955&mid=3
	url2 := c.Input().Get("mid")
	var url string
	if url2 == "" {
		url = url1
	} else {
		url = url1 + "&mid=" + url2
	}
	c.Data["Url"] = url
	c.TplName = "role_err.tpl"
}

func (c *RoleController) AddAndEdit() {
	r := m.Role{}
	if err := c.ParseForm(&r); err != nil {
		//handle error
		// c.Rsp(false, err.Error())
		beego.Error(err.Error)
		return
	}
	var id int64
	var err error
	Rid, _ := c.GetInt64("Id")
	if Rid > 0 {
		id, err = m.UpdateRole(&r)
	} else {
		id, err = m.AddRole(&r)
	}
	if err == nil && id > 0 {
		// c.Rsp(true, "Success")
		return
	} else {
		// c.Rsp(false, err.Error())
		beego.Error(err.Error)
		return
	}

}

func (c *RoleController) DelRole() {
	Id, _ := c.GetInt64("Id")
	status, err := m.DelRoleById(Id)
	if err == nil && status > 0 {
		// c.Rsp(true, "Success")
		return
	} else {
		// c.Rsp(false, err.Error())
		beego.Error(err.Error)
		return
	}
}

func (c *RoleController) Getlist() {
	roles, _ := m.GetRolelist(1, 1000, "Id")
	if len(roles) < 1 {
		roles = []orm.Params{}
	}
	c.Data["json"] = &roles
	c.ServeJSON()
	return
}

// func (c *RoleController) AccessToNode() {
// 	roleid, _ := c.GetInt64("Id")
// 	if c.IsAjax() {
// 		groupid, _ := c.GetInt64("group_id")
// 		nodes, count := m.GetNodelistByGroupid(groupid)
// 		list, _ := m.GetNodelistByRoleId(roleid)
// 		for i := 0; i < len(nodes); i++ {
// 			if nodes[i]["Pid"] != 0 {
// 				nodes[i]["_parentId"] = nodes[i]["Pid"]
// 			} else {
// 				nodes[i]["state"] = "closed"
// 			}
// 			for x := 0; x < len(list); x++ {
// 				if nodes[i]["Id"] == list[x]["Id"] {
// 					nodes[i]["checked"] = 1
// 				}
// 			}
// 		}
// 		if len(nodes) < 1 {
// 			nodes = []orm.Params{}
// 		}
// 		c.Data["json"] = &map[string]interface{}{"total": count, "rows": &nodes}
// 		c.ServeJSON()
// 		return
// 	} else {
// 		grouplist := m.GroupList()
// 		b, _ := json.Marshal(grouplist)
// 		c.Data["grouplist"] = string(b)
// 		c.Data["roleid"] = roleid
// 		// c.TplName = c.GetTemplatetype() + "/rbac/accesstonode.tpl"
// 	}
// }

// func (c *RoleController) AddAccess() {
// 	roleid, _ := c.GetInt64("roleid")
// 	group_id, _ := c.GetInt64("group_id")
// 	err := m.DelGroupNode(roleid, group_id)
// 	if err != nil {
// 		// c.Rsp(false, err.Error())
// 		beego.Error(err.Error)
// 	}
// 	ids := c.GetString("ids")
// 	nodeids := strings.Split(ids, ",")
// 	for _, v := range nodeids {
// 		id, _ := strconv.Atoi(v)
// 		_, err := m.AddRoleNode(roleid, int64(id))
// 		if err != nil {
// 			// c.Rsp(false, err.Error())
// 			beego.Error(err.Error)
// 		}
// 	}
// 	// c.Rsp(true, "success")

// }

func (c *RoleController) RoleToUserList() {
	roleid, _ := c.GetInt64("Id")
	if c.IsAjax() {
		users, count := m.Getuserlist(1, 1000, "Id")
		list, _ := m.GetUserByRoleId(roleid)
		for i := 0; i < len(users); i++ {
			for x := 0; x < len(list); x++ {
				if users[i]["Id"] == list[x]["Id"] {
					users[i]["checked"] = 1
				}
			}
		}
		if len(users) < 1 {
			users = []orm.Params{}
		}
		c.Data["json"] = &map[string]interface{}{"total": count, "rows": &users}
		c.ServeJSON()
		return
	} else {
		c.Data["roleid"] = roleid
		// c.TplName = c.GetTemplatetype() + "/rbac/roletouserlist.tpl"
	}
}

func (c *RoleController) AddRoleToUser() {
	roleid, _ := c.GetInt64("Id")
	ids := c.GetString("ids")
	userids := strings.Split(ids, ",")
	err := m.DelUserRole(roleid)
	if err != nil {
		// c.Rsp(false, err.Error())
		beego.Error(err.Error)
	}
	if len(ids) > 0 {
		for _, v := range userids {
			id, _ := strconv.Atoi(v)
			_, err := m.AddRoleUser(roleid, int64(id))
			if err != nil {
				// c.Rsp(false, err.Error())
				beego.Error(err.Error)
			}
		}
	}
	// c.Rsp(true, "success")
}
