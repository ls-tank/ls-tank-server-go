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

func (this *User) GetId() {
	this.Id = bson.NewObjectId()
}

func (this *User) Insert() (bool, interface{}) {
	err := DB.C("user").Insert(this)
	if err != nil {
		return false, err
	}
	return true, this.Id
}

func (this *User) Find(value string) (bool, interface{}) {
	if !bson.IsObjectIdHex(value) {
		return false, "id格式错误"
	}
	err := DB.C("user").Find(bson.M{"_id": bson.ObjectIdHex(value)}).One(&this)
	if err != nil {
		return false, "用户不存在"
	}
	return true, nil
}

func (this *User) Save() {
	DB.C("user").Update(bson.M{"_id": this.Id}, this)
}
