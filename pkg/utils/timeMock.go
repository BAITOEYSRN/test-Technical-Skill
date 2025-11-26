package utils

import "time"

// unitTest
var Now func() time.Time = time.Now

func SetMockFunc() {
	Now = func() time.Time {
		return TimeMock()
	}
}

func TimeMock() time.Time {
	return time.Date(2013, 10, 13, 13, 13, 13, 13, time.Local)
}
