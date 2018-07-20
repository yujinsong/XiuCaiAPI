package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/XiuCai/XiuCaiAPI/utils/mysql-utils"
	"fmt"
	"github.com/XiuCai/XiuCaiAPI/model"
	"github.com/XiuCai/XiuCaiAPI/utils"
	"github.com/XiuCai/XiuCaiAPI/common"
)

var userInstance *UserInstance = nil

type UserInstance struct {
	tableName       string
	databaseEngineer *xorm.Engine
}

func GetUserInstance() *UserInstance {
	if userInstance == nil {
		userInstance = &UserInstance{
			tableName:       "userinfo",
			databaseEngineer: mysql_utils.GetInstance(),
		}
	}
	return userInstance
}

// 用户等级,1:超级用户，可以删帖、推荐到首页, 2: 管理员，可以删帖, 3:只能推荐帖子到首页的超级用户
func (user *UserInstance) UserLevel(uuid int64) int {
	sql := fmt.Sprintf("select kind from superuser where uuid = %d", uuid)
	var superUser model.SuperUser
	_, err := user.databaseEngineer.SQL(sql).Get(&superUser)
	if err != nil {
		utils.Logger.Error("------ Get User Level error, uuid:", uuid, " err:", err)
		return -1
	}

	return superUser.Kind
}


func (user *UserInstance) GetUserById(uuid int64) (model.User, error) {
	sql := fmt.Sprintf("select * from userinfo where id = %d and flag in(0, 1)", uuid)
	var userInfo model.User
	_, err := user.databaseEngineer.SQL(sql).Get(&userInfo)
	if err != nil {
		utils.Logger.Error("------ Get user by id error, uuid:", uuid, " err:", err)
		return model.User{}, err
	}

	return userInfo, nil
}


func (user *UserInstance) GetUserByTel(tel string) (model.User, error) {
	sql := fmt.Sprintf("select * from userinfo where tel = '%s' and flag in(0, 1)", tel)
	var userInfo model.User
	_, err := user.databaseEngineer.SQL(sql).Get(&userInfo)
	if err != nil {
		utils.Logger.Error("------ Get user by id error, tel:", tel, " err:", err)
		return model.User{}, err
	}

	return userInfo, nil
}

// 用户注册
func (user *UserInstance) Register(tel, username string) int64 {
	sql := fmt.Sprintf("insert into userinfo(`tel`, `username`, `regtime`) values ('%s', '%s', '%s')",
		tel, username, common.CurrentTime())
	result, err := user.databaseEngineer.Exec(sql)
	if err != nil {
		utils.Logger.Error("------ Register error, tel:", tel, " username:",
			username, "regtime:", common.CurrentTime(), "err:", err)
		return -1
	}

	uuid, _ := result.LastInsertId()
	return uuid
}