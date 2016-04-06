package controllers

import (
	"crypto/md5"
	"encoding/hex"
	// "fmt"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"ls-tank-server-go/models"
	"strconv"
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
		user.Username = this.GetString("username")
		user.Password = encrypt(this.GetString("password"))
		user.Nickname = this.GetString("nickname")
		return user.Insert()
	}()

	this.Data["json"] = map[string]interface{}{
		"ok":   ok,
		"data": data,
	}
	this.ServeJSON()
}

func (this *UserController) Login() {
	user := &models.User{}
	username := this.GetString("username")
	password := encrypt(this.GetString("password"))

	ok, data := user.Authenticate(username, password)
	if ok {
		this.SetSession("uid", user.Id.Hex())
	}

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
	user := &models.User{}
	user.Find(this.GetSession("uid").(string))

	_data := make(map[string]interface{})
	for key, value := range this.Input() {
		_data[key], _ = strconv.Atoi(value[0])
	}

	ok, data := user.Update(_data)

	this.Data["json"] = map[string]interface{}{
		"ok":   ok,
		"data": data,
	}
	this.ServeJSON()
}

func (this *UserController) AddCoins() {
	user := &models.User{}
	uid := this.GetString("uid")
	coins, ok := this.GetInt("coins")

	user.Find(uid)
	user.Diamond += coins
	user.Save()

	this.Data["json"] = map[string]interface{}{
		"ok":   ok,
		"data": user,
	}
	this.ServeJSON()

}
