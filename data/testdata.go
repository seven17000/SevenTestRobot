package data

import (
	"seventestrobot/common"
	"time"
)

var baseTime = time.Now()
var globalStat RobotStat

//测试机器人状态
type RobotStat struct {
	StatMap       map[string]*TestData //按场景储存TestData
	TotalTask     int                  //所有任务数
	robotNum      int                  //机器人数
	interval      time.Duration        //统计周期
	intervalCount int                  //周期内采样数

	robotNumRun  int //在线机器人数
	robotNumStop int //离线机器人数

	firstTranTime time.Time //周期内的测试开始时间
	lastTranTime  time.Time //周期内的测试结束时间

	StatTimePrecision time.Duration //90%事务时间块大小
	StatTimeMax       time.Duration //90%事务执行时间得区间大小

	showStatCount    int //统计的次数
	timeSectionCount int //分段数

	isStop       bool
	AllRobotDone bool
	TestResult   string //当前测试结果
}

//测试数据
type TestData struct {
	Qps         float64       //Qps
	TotalTask   int           //总任务数
	ErrTask     int           //错误任务数
	TimeoutNum  int           //超时连接数
	SuccessRate float32       //成功率
	AvgMillisec time.Duration //平均响应时间
	MinMillisec time.Duration //最小响应时间
	MaxMillisec time.Duration //最大响应时间

	Per90MaxTime time.Duration //百分之九十响应时间
	Per90AvgTime time.Duration //百分之九十最大响应时间

	UploadCount int     //上传次数
	TestTime    float64 //测试时间

	totalMillsec       time.Duration //总响应时间
	section            []int         //周期内的时间块列表
	statTimeMaxSection time.Duration //90%事务执行时间得区间大小
	statTimePreSection time.Duration //90%事务执行时间时得时间块大小,也就是时间精度
	firstTranTime      time.Time     //第一次执行时间
	lastTranTime       time.Time     //最后一次执行时间
	timeSectionCount   int           //时间分段计数器

	//周期内数据
	intervalTotalTask     int           //间隔内任务总数
	intervalErrTask       int           //间隔内错误数
	intervalFisrtTranTime time.Time     //周期内第一次执行时间
	intervalLastTranTime  time.Time     //周期内最后一次执行时间
	intervalTotal         time.Duration //间隔内总的响应时间
}

func InitStat(s *RobotStat) {
	s.robotNumRun = 0
	s.robotNumStop = 0
	s.TotalTask = 0
	s.AllRobotDone = false
	s.isStop = false
	s.firstTranTime = time.Now()

	s.lastTranTime = baseTime
	s.StatMap = make(map[string]*TestData)
	s.interval = time.Duration(common.Interval) * time.Second
}

func GetGlobalStat()(*RobotStat){
	return &globalStat
}

func SetGlobalStat(stat RobotStat) {
	globalStat = stat
}

