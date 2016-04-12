package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `_id`
	Username  string
	Password  string
	Nickname  string
	Diamond   int
	TankHead  int
	TankBody  int
	TankWheel int
	Kill      int
	Dead      int
}

func (this *User) GetId() {
	this.Id = bson.NewObjectId()
}

func (this *User) Insert() (bool, interface{}) {
	err := DB.C("user").Insert(this)
	if err != nil {
		return false, err
	}
	return true, this
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

func (this *User) Authenticate(name string, password string) (bool, interface{}) {
	err := DB.C("user").Find(bson.M{"username": name, "password": password}).One(&this)
	if err != nil {
		return false, "账号或密码错误"
	}
	return true, this
}

func (this *User) Update(data map[string]interface{}) (bool, interface{}) {
	err := DB.C("user").Update(bson.M{"_id": this.Id}, bson.M{"$set": data})
	if err != nil {
		return false, "更新失败"
	}
	return true, "更新成功"
}

func (this *User) IncData(key string, step int) (bool, interface{}) {
	_data := map[string]int{}
	_data[key] = step
	err := DB.C("user").Update(bson.M{"_id": this.Id}, bson.M{"$inc": _data})
	if err != nil {
		return false, "更新失败"
	}
	return true, "更新成功"
}

func (this *User) Save() {
	DB.C("user").Update(bson.M{"_id": this.Id}, this)
}
