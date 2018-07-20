package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/XiuCai/XiuCaiAPI/utils/mysql-utils"
	"fmt"
	"github.com/XiuCai/XiuCaiAPI/utils"
)

var circleInstance *CircleInstance = nil

type CircleInstance struct {
	tableName       string
	databaseEngineer *xorm.Engine
}

func GetCircleInstance() *CircleInstance {
	if circleInstance == nil {
		circleInstance = &CircleInstance{
			tableName:       "usercircle",
			databaseEngineer: mysql_utils.GetInstance(),
		}
	}
	return circleInstance
}


func (circle *CircleInstance) HasJoinCircle(uuid, circleId int64) bool {
	cnt, err := circle.databaseEngineer.Table(circle.tableName).Where("uuid = ? and circleid = ?",
		uuid, circleId).Count()

	if err != nil {
		utils.Logger.Error("------ check has join circle error, uuid:", uuid, "circleId:", circleId, " error:", err)
		return true
	}

	if cnt > 0 {
		return true
	} else {
		return false
	}
}


func (circle *CircleInstance) JoinToCircle(uuid, circleId int64) error {
	sql := fmt.Sprintf("insert into usercircle(`uuid`, `circleid`) values(%d, %d)", uuid, circleId)
	_, err := circle.databaseEngineer.Exec(sql)
	if err != nil {
		utils.Logger.Error("------ join to circle error, uuid:", uuid, " circleId:", circleId, " err:", err)
		return err
	}

	return nil
}