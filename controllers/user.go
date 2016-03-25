package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	// "gopkg.in/mgo.v2/bson"
	// "ls-tank-server-go/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Add() {
	fmt.Println("add...")

	this.Data["json"] = map[string]interface{}{
		"ok":   "ok",
		"data": "add",
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
