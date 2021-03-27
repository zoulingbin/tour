package timer

import "time"

func GetNowTime() time.Time {
	return time.Now()
}

//推算时间
func GetCalculateTime(currentTimer time.Time, s string) (time.Time, error) {
	duration, err := time.ParseDuration(s)
	if err != nil {
		return time.Time{},err
	}
	return currentTimer.Add(duration),nil
}