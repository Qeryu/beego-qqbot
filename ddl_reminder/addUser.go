package ddl_reminder

import (
	"demo1/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

func addUserData(val Data, o orm.Ormer) {
	user := new(models.Users)
	user.Nickname = val.Nickname
	user.UserId = val.UserId
	created, id, er := o.ReadOrCreate(user, "UserId")
	if er == nil {
		if created {
			o := orm.NewOrm()
			qs := o.QueryTable("users")
			var user models.Users
			qs.Filter("user_id", val.UserId).One(&user)
			o.Raw("UPDATE users SET use_ddl_reminder='N' WHERE user_id=?",
				strconv.FormatInt(user.UserId, 10)).Exec()
			qs.Filter("user_id", val.UserId).One(&user)
			fmt.Println("add users", user)
			fmt.Println("New insert id", id)
		} else {
			fmt.Println("Get an object", id)
		}
	} else {
		fmt.Println("some error", er.Error())
	}
}

