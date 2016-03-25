package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"ls-tank-server-go/models"
)

type UserController struct {
	beego.Controller
}

func encrypt(input string) (ouput string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(input))
	ouput = hex.EncodeToString(md5Ctx.Sum(nil))
	return
}

func (this *UserController) Add() {
	ok, data := func() (bool, interface{}) {
		user := &models.User{}
		user.Id = bson.NewObjectId()
		user.Name = this.GetString("name")
		user.Password = encrypt(this.GetString("password"))
		return user.Insert()
	}()

	this.Data["json"] = map[string]interface{}{
		"ok":   ok,
		"data": data,
	}
	this.ServeJSON()
}

func (this *UserController) Get() {
	fmt.Println("get...")

	this.Data["json"] = map[string]interface{}{
		"ok":   "ok",
		"data": "get",
	}
	this.ServeJSON()
}

// todo
// func isExist(name string) {
//
// }
