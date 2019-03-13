package models

import "sync"

type TimeLimit interface {
	Count(nowTime int64) (curCount int)
	Check(nowTime int64) int
}

// min implement TimeLimit interface
type MinLimit struct {
	count   int
	curTime int64
}

func (m *MinLimit) Count(nowTime int64) (curCount int) {
	if nowTime-m.curTime > 60 {
		m.count = 1
		m.curTime = nowTime
		curCount = m.count
	}
	m.count++
	curCount = m.count
	return
}

func (m *MinLimit) Check(nowTime int64) int {
	if nowTime-m.curTime > 60 {
		return 0
	}
	return m.count
}

//SecLimit implement TimeLimit interface
type SecLimit struct {
	count   int
	curTime int64
}

func (s *SecLimit) Count(nowTime int64) (curCount int) {
	if s.curTime != nowTime {
		s.count = 1
		s.curTime = nowTime
		curCount = s.count
		return
	}

	s.count++
	curCount = s.count
	return
}

func (s *SecLimit) Check(nowTime int64) int {
	if s.curTime != nowTime {
		return 0
	}
	return s.count
}

//manage Limit
type Limit struct {
	secLimit TimeLimit
	minLimit TimeLimit
}

type SecLimitMgr struct {
	UserLimitMap map[int]*Limit
	IpLimitMap   map[string]*Limit
	Lock         sync.Mutex
}
