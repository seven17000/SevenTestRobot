package app

import (
	"SevenTestRobort/config"
	"SevenTestRobort/robotsrv/robot"
	"net/http"
)

var (
	RobotConf config.RobotConfig
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
}