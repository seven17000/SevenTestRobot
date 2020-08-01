package config

import (
	"SevenTestRobort/common"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RobotConfig struct {
	TestID    		int64
	RoboNum   		int
	Qps       		int
	RobotSence   	string
	TestTime  		int64
	SuccessRate    	float32
	TestAddr        string
}

func LoadRobotConfig(path string)(error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, &common.RobotConf)
	if err != nil {
		return err
	}
	return err
}

func ReadRobotConfig(r *http.Request) (config RobotConfig, err error)  {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(buf, &config)
	if err != nil {
		return
	}
	return
}