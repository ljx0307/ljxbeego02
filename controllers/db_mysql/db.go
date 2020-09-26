package db_mysql

import (
	"bee04/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
)

var DB *sql.DB

func init()  {
	fmt.Println("连接mysql数据库")
	config := beego.AppConfig

	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp :=config.String("db_ip")
	dbName :=config.String("db_name")
	connUrl :=dbUser +":" + dbPassword +"@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	db,err := sql.Open(dbDriver,connUrl)
	if err !=nil {
		panic("数据库连接错误，请检查配置")
	}
	DB = db
	fmt.Println(db)
}
func InserUser(user models.User)(int64,error){
	//将用户密码进行hash脱敏，使用md5计算密码hash值，并储存hash值
	hashMd5 := md5.New()
	hashMd5.Write([]byte(user.Password))
	bytes := hashMd5.Sum([]byte("nil"))
	user.Password = hex.EncodeToString(bytes)
	fmt.Println("将要保存的用户名:",user.Nick,"密码:",user.Password)
	result,err :=DB.Exec("insert into user (nick,password)values (?,?)",user.Nick,user.Password)
	if err != nil {
		return -1,err
	}
	id, err := result.RowsAffected()
	if err != nil{
		return  -1,err
	}

	return id,nil
}
func QueryUser(){
	DB.QueryRow("select*from")
}