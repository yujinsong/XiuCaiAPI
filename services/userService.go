package services

import (
	"github.com/XiuCai/XiuCaiAPI/common"
	"github.com/XiuCai/XiuCaiAPI/dao"
	"github.com/XiuCai/XiuCaiAPI/model"
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