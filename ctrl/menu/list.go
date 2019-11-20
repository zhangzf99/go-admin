package menu

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-admin/conf"
	"go-admin/models"
	"go-admin/modules/response"
)

func List(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get(conf.Cfg.Token)
	if v == nil {
		response.ShowError(c, "fail")
		return
	}
	uid := session.Get(v)
	user := models.SystemUser{Id: uid.(int)}
	has := user.GetRow()
	if !has {
		response.ShowError(c, "fail")
		return
	}

	menu := models.SystemMenu{}
	if user.Nickname == "admin" {
		menuArr, err := menu.GetAll()
		if err != nil {
			response.ShowError(c, "fail")
			return
		}
		jsonArr :=tree(menuArr)
		response.ShowData(c,jsonArr)
		return
	} else {
		//var constant []models.SystemMenu
		//menu := models.SystemMenu{Type:1}
		//constant,err:=menu.GetAll()




	}


}

func tree(menuArr []models.SystemMenu) ([]interface{}) {
	role := models.SystemRole{}
	mrArr := role.GetRowMenu()
	fmt.Println(mrArr)
	var menuMap = make(map[int][]models.SystemMenu, 0)
	for _, value := range menuArr {
		menuMap[value.Pid] = append(menuMap[value.Pid], value)
	}
	var jsonArr []interface{}

	mainMenu, ok := menuMap[0]
	if !ok {
		return nil
	}
	for _, value := range mainMenu {
		var item = make(map[string]interface{})
		item["path"] = value.Path
		item["component"] = value.Component
		if value.Redirect != "" {
			item["redirect"] = value.Redirect
		}
		if value.Alwaysshow ==1 {
			item["alwaysShow"] = true
		}
		if value.Hidden == 1 {
			item["hidden"] = true
		}
		var meta=make(map[string]interface{})
		if value.MetaTitle!=""{
			meta["title"]=value.MetaTitle
		}
		if value.MetaIcon!="" {
			meta["icon"]=value.MetaIcon
		}
		if value.MetaAffix==1 {
			meta["affix"] = true
		}
		if value.MetaNocache==1 {
			meta["noCache"] = true
		}
		if len(meta)>0 {
			item["meta"]=meta
		}
		if _,ok:=menuMap[value.Id] ;ok{
			item["children"]=treeChilden(menuMap[value.Id])
		}
		jsonArr = append(jsonArr,item)
	}
	return jsonArr

}
func treeChilden(menuArr []models.SystemMenu )[]interface{} {
	var jsonArr []interface{}
	for _,value:=range menuArr  {
		var item = make(map[string]interface{})
		item["path"] = value.Path
		item["component"] = value.Component
		if value.Redirect != "" {
			item["redirect"] = value.Redirect
		}
		if value.Alwaysshow ==1 {
			item["alwaysShow"] = true
		}
		if value.Hidden == 1 {
			item["hidden"] = true
		}
		var meta=make(map[string]interface{})
		if value.MetaTitle!=""{
			meta["title"]=value.MetaTitle
		}
		if value.MetaIcon!="" {
			meta["icon"]=value.MetaIcon
		}
		if value.MetaAffix==1 {
			meta["affix"] = true
		}
		if value.MetaNocache==1 {
			meta["noCache"] = true
		}
		if len(meta)>0 {
			item["meta"]=meta
		}
		jsonArr = append(jsonArr,item)
	}
	return jsonArr
}
