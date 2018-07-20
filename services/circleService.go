package services

import "github.com/XiuCai/XiuCaiAPI/dao"

func HasJoinCircle(uuid, circleId int64) bool {
	return dao.GetCircleInstance().HasJoinCircle(uuid, circleId)
}

func JoinToCircle(uuid, circleId int64) error {
	return dao.GetCircleInstance().JoinToCircle(uuid, circleId)
}
