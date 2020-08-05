package app

import (
	"net/http"
	"seventestrobot/config"
	"seventestrobot/data"
	"seventestrobot/robotsrv/robot"
)

func Run(w http.ResponseWriter, r *http.Request) {

	RobotConf ,err := config.ReadRobotConfig(r)
	if err != nil {
		return
	}

	err = robot.InitAllRobot(RobotConf)
	if err != nil {
		return
	}

	robotStat := data.GetGlobalStat()
	data.InitStat(robotStat)

}