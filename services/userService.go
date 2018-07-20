package services

import (
	"github.com/XiuCai/XiuCaiAPI/common"
	"github.com/XiuCai/XiuCaiAPI/dao"
	"github.com/XiuCai/XiuCaiAPI/model"
	"math/rand"
	"time"
	"strconv"
)

// ios 显示积分与钱
func GetDisplayFlag(curVersion string) int {
	flag := common.DisplayFlag
	iosVersion := common.IosVersion

	if curVersion > iosVersion && flag == 0 {
		return 0
	} else {
		return 1
	}
}

// 用户等级
func UserLevel(uuid int64) int {
	level := dao.GetUserInstance().UserLevel(uuid)
	if level < 0 {
		return 0
	} else {
		return level
	}
}

func GetUserById(uuid int64) (model.User, error) {
	return dao.GetUserInstance().GetUserById(uuid)
}

func GetUserByTel(tel string) (model.User, error) {
	return dao.GetUserInstance().GetUserByTel(tel)
}

func Register(loginModel model.LoginModel) int64 {
	userName := GenerateUserName()
	uuid := dao.GetUserInstance().Register(loginModel.Tel, userName)

	return uuid
}


func GenerateUserName() string {
	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(88888888)
	return "秀才" + strconv.Itoa(100000 + rand)
}