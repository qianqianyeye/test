package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/itsjamie/gin-cors"
	"fmt"
	"git.jiaxianghudong.com/go/monitor/mysql"
	"git.jiaxianghudong.com/go/utils"
	"net/http"
)
var (
	userName="root"
	passWord="root"
	addr="mysql:3306"
	db="test"
	maxOpen=10
	maxIdle=5
)

func main()  {
	driver := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", userName, passWord, addr, db)
	err := mysql.Init(driver, maxOpen, maxIdle)
	if err!=nil {
		panic(err)
	}
	router := gin.Default()
	//网页跨域问题
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET,PUT,POST,DELETE,OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	registerRouter(router)
	router.Run(":8011")
}

func registerRouter(router *gin.Engine)  {
	r := router.Group("serviceTest/")
	r.POST("add",Add)
	r.POST("delete/:id",Delete)
	r.POST("update/:id",Update)
	r.GET("query",Query)
	r.POST("login",Login)
}

type User struct {
	ID int64 `sql:"id"" json:"id" form:"id"`
	UserName string `sql:"user_name" json:"user_name" form:"user_name"`
	Status int `sql:"status" json:"status" form:"status"`
	PassWord string `sql:"pass_word" json:"pass_word" form:"pass_word"`
	CreateTime time.Time `sql:"create_time" json:"-" form:"create_time"`
	UpdateTime time.Time `sql:"update_time" json:"-" form:"update_time"`
	DeleteTime time.Time `sql:"delete_time" json:"-" form:"delete_time"`
}

func Add(ctx *gin.Context)  {
	user :=User{}
	user.CreateTime=time.Now()
	user.Status=0
	if err:=ctx.Bind(&user);err==nil{
		query:=fmt.Sprintf("insert into user (user_name,pass_word,status) values ('%s','%s',%v)",user.UserName,user.PassWord,user.Status)
		fmt.Println(query)
		mysql.Insert(query)
		ctx.JSON(http.StatusOK,gin.H{"data":"success"})
	}

}

func Delete(ctx *gin.Context)  {
	id:=utils.Atoi(ctx.Param("id"))
	query :=fmt.Sprintf("update user set status=2 where id=%v",id)
	mysql.Insert(query)
	ctx.JSON(http.StatusOK,gin.H{"data":"success"})
}

func Update(ctx *gin.Context)  {
	user :=User{}
	user.ID=int64(utils.Atoi(ctx.Param("id")))
	if err:=ctx.Bind(&user);err==nil{
		query :=fmt.Sprintf("update user set user_name ='%s' ,pass_word='%s',status=%v where id=%v",user.UserName,user.PassWord,user.Status,user.ID)
		mysql.Insert(query)
		ctx.JSON(http.StatusOK,gin.H{"data":"success"})
	}
}
func Query(ctx *gin.Context)  {
	var users []User
	query:=fmt.Sprintf("select * from user")
	mysql.Query(query,&users)
	ctx.JSON(http.StatusOK,gin.H{"data":users})
}

func Login(ctx *gin.Context)  {
	user :=User{}
	if err:=ctx.Bind(&user);err==nil{
		var result []User
		query:=fmt.Sprintf("select * from user where user_name='%s' and pass_word='%s' and status=0",user.UserName,user.PassWord)
		//fmt.Println(query)
		err :=mysql.Query(query,&result)
		if err!=nil {
			fmt.Println(err)
		}
		if len(result)>0 {
			ctx.JSON(http.StatusOK,gin.H{"data":"success"})
		}else {
			ctx.JSON(http.StatusOK,gin.H{"data":"username or password err"})
		}

	}
}