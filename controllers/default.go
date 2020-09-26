package controllers

import (
	"bee04/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}


func (c *MainController) Get() {
	c.GetString("user")
	c.GetInt("age")
	userName := c.Ctx.Input.Query("user")
	password :=  c.Ctx.Input.Query("pwd")
	if userName != "admin" || password != "123456"{
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据错误"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("恭喜，数据正确"))

	//c.Data["Website"] = "www.badu.com"
	//c.Data["Email"] = "2632507285@qq.com"
	//c.TplName = "index.tpl"
}

//func (c *MainController) Post(){
//	for i :=0;i<10;i++{
//		fmt.Printf("第%d次打印\n",i)
//	}
//}
//func (c *MainController) Post(){
//	name :=c.Ctx.Request.FormValue("name")
//	age := c.Ctx.Request.FormValue("age")
//	sex :=c.Ctx.Request.FormValue("sex")
//	fmt.Println(name,age,sex)
//	if name !="admin"&& age != "18"{
//		c.Ctx.WriteString("数据校验失败")
//		return
//	}
//	c.Ctx.WriteString("数据校验成功")
//}
//
// */
//func(c *MainController)Post(){
//	var person models.Person
//	dataBytes,err :=ioutil.ReadAll(c.Ctx.Request.Body)
//	if err !=nil{
//		c.Ctx.WriteString("数据解析失败，请重试")
//	}
//	err = json.Unmarshal(dataBytes,&person)
//	if err !=nil {
//		c.Ctx.WriteString("数据解析失败")
//		return
//	}
//	fmt.Println("姓名",person.Name)
//	fmt.Println("年龄",person.Age)
//	fmt.Println("性别",person.Sex)
//	c.Ctx.WriteString("数据解析成功")
//}

func(c *MainController)Post(){
	var person models.Personer
	dataBytes,err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err !=nil{
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	err = json.Unmarshal(dataBytes,&person)
	if err !=nil {
		c.Ctx.WriteString("数据解析失败")
		return
	}
	fmt.Println("姓名",person.Name)
	fmt.Println("生日",person.Birthday)
	fmt.Println("地址",person.Address)
	fmt.Println("昵称",person.Nick)
	c.Ctx.WriteString("数据解析成功")
}
