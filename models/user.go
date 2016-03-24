package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `_id`
	Name      string
	Password  string
	Diamond   uint
	TankHead  uint
	TankBody  uint
	TankWheel uint
	Kill      uint
	Dead      uint
}

// 获取 id
func (this *User) GetId() {
	this.Id = bson.NewObjectId()
}

// 创建新用户
func (this *User) Insert() (bool, interface{}) {
	err := DB.C("user").Insert(this)
	if err != nil {
		return false, err
	}
	return true, this.Id
}

// 根据 ObjectId 获取用户信息
func (this *User) Find(value string) bool {
	err := DB.C("user").Find(bson.M{"_id": bson.ObjectIdHex(value)}).One(&this)
	if err != nil {
		return false
	}
	return true
}
