package ddl_reminder

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"net/http"
)

type FriendListReply struct {
	Data	[]Data
}

type Data struct {
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
	UserId   int64 `json:"user_id"`
}

func FillFriendList(o orm.Ormer) {
	//_,err := o.Raw("DELETE FROM users").Exec()
	//_,err = o.Raw("ALTER TABLE users AUTO_INCREMENT=1").Exec()
	//if err != nil {
	//	fmt.Println("row sql error", err.Error())
	//}

	urlGetList := "http://192.168.0.1:5700/get_friend_list"
	reps, err := http.Get(urlGetList)
	if err != nil {
		fmt.Println("http get is err", err.Error())
	}

	list := FriendListReply{}
	defer reps.Body.Close()
	decoder := json.NewDecoder(reps.Body)
	err = decoder.Decode(&list)
	if err != nil {
		fmt.Println("get json is err", err.Error())
	}

	for _, val := range list.Data {
		addUserData(val, o)
	}

}