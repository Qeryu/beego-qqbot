package ddl_reminder

import (
	"demo1/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

func addUserData(val Data, o orm.Ormer) {
	user := new(models.Users)
	user.Nickname = val.Nickname
	user.UserId = val.UserId
	created, id, er := o.ReadOrCreate(user, "UserId")
	if er == nil {
		if created {
			fmt.Println("New insert id", id)
		} else {
			fmt.Println("Get an object", id)
		}
	} else {
		fmt.Println("some error", er.Error())
	}
}

