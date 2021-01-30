package time_test

import (
	"fmt"
	"testing"
	"time"
)

/* Sleep()
const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
	)
*/

func TestTime(t *testing.T) {
	// 获取当前时间
	time.Now()

	// 获取指定时间
	time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	// 日期转为文本
	time.Now().Format("2021/1/27 15:32:00")

	// 文本转为日期：
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	p1, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(p1)
	// output : 2013-02-03 19:54:00 -0800 PST

	// 时间戳
	// 指定的日期(当前日期，指定的某个日期)，距离1970年1月1日0点0时0分0秒的时间差值
	// 有秒，纳秒两种函数。
	// 获取单位为秒的时间戳
	// time.Unix()
}
