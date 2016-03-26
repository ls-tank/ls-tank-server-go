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

func get(id string) (bool, interface{}) {
	user := &models.User{}
	if _ok, _data := user.Find(id); !_ok {
		return false, _data
	}
	return true, user
}

func (this *UserController) Get() {
	ok, data := get(this.GetString("id"))

	this.Data["json"] = map[string]interface{}{
		"ok":   ok,
		"data": data,
	}
	this.ServeJSON()
}

func (this *UserController) Update() {
	// todo
	this.Data["json"] = map[string]interface{}{
		"ok":   "ok",
		"data": "data",
	}
	this.ServeJSON()
}
