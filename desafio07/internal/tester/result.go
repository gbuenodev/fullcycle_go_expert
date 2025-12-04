package tester

import "time"

type RequestResult struct {
	StatusCode int
	Duration   time.Duration
	Error      error
}

type Statistics struct {
	TotalRequests  int
	TotalDuration  time.Duration
	StatusCodeDist map[int]int
	SuccessCount   int
	ErrorCount     int
}

func NewStatistics() *Statistics {
	return &Statistics{
		StatusCodeDist: make(map[int]int),
	}
}

func (s *Statistics) AddRequestResult(result RequestResult) {
	s.TotalRequests++

	if result.Error != nil {
		s.ErrorCount++
		return
	}

	s.StatusCodeDist[result.StatusCode]++
	if result.StatusCode == 200 {
		s.SuccessCount++
	}
}
