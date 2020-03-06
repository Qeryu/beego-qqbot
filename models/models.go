package models

import (
	"github.com/astaxie/beego/orm"
)

type Users struct {
	 Id           	int
	 UserId         int64
	 Nickname       string
	 UseDdlReminder string
}

type Ddls struct {
	Id		int
	Thing	string
	Time	string
}

func init()  {
	orm.RegisterModel(new(Users))
	orm.RegisterModel(new(Ddls))
}

