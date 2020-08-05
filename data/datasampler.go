package data

import "time"

type Sampler struct {
	Key          string
	Qps          int
	StartTime    time.Time
	EndTime      time.Time
	responseTime int
	Result       bool
}

func (s *Sampler)StartSampler()  {
	s.StartTime = time.Now()
}

func (s *Sampler)EndSampler()  {
	s.EndTime = time.Now()
}

func (s *Sampler)ResultSampler(isFailed bool) {
	s.Result = isFailed
}