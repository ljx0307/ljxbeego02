package controllers

import (
	"BeegoProject0603/db_mysql"
	models2 "BeegoProject0603/models"
	"bee04/models"

	"beego"
	"encoding/json"
	"fmt"
	"io/ioutil"

)

type RegisterController struct {
	beego.Controller
}

/**
 * 该方法用于处理post类型请求
 */
func (r *RegisterController) Post(){
	fmt.Println(r == nil)
	fmt.Println(r.Ctx == nil)
	fmt.Println(r.Ctx.Request == nil)
	//1、接收前端传递的数据
	//body,err := r.Ctx.Request.GetBody()
	//if err != nil {
	//	r.Ctx.WriteString("数据接收错误")
	//	return
	//}
	bodyBytes,err :=ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil {
		r.Ctx.WriteString("数据接收错误,请重试")
		return
	}
	var user models.User
	err = json.Unmarshal(bodyBytes,&user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("数据解析错误")
		return
	}

	//3、将解析到的用户数据，保存到数据
	var id, _ = db_mysql.InsertUser(models2.User(user))
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("用户保存失败.")
		return
	}
	fmt.Println(id)

	result := models.ResponseResult{
		Code:    0,
		Message: "保存成功",
		Data:    nil,
	}
	r.Data["json"] = &result
	r.ServeJSON()
	//r.Ctx.WriteString("恭喜，用户保存成功")

	//var user models.User
	//err := r.ParseForm(&user)
	//if err != nil {
	//	r.Ctx.WriteString("数据解析错误")
	//}

	//fmt.Println(user.User)
	//fmt.Println(user.Nick)
	//r.Ctx.WriteString("数据接收成功，并成功解析。")
}

