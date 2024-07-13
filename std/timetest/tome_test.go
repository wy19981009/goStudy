package timetest

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	//02d输出的整数不足两位 用0补足
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func TestTimeUnix(t *testing.T) {
	now := time.Now()
	unixSec := now.Unix()
	fmt.Printf("unixSec:%d\n", unixSec)
	milliSec := now.UnixMilli()
	fmt.Printf("millisec:%d\n", milliSec)
	nanoSec := now.UnixNano()
	fmt.Printf("nanoSec:%d\n", nanoSec)
}

func TestParseTime(t *testing.T) {
	//t, err := time.Parse("2006-01-02 15:04:05", "2022-07-28 18:06:00")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(t)
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	//按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", now.Format("2006/01/02 15:04:05"), loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
}

func TestFormatTime(t *testing.T) {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分05秒
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

func TestDuration(t *testing.T) {
	five := 5 * time.Minute
	now := time.Now()
	fmt.Printf("time now:%s \n", now)
	add := now.Add(five)
	fmt.Printf("time add:%s \n", add)
	sub := add.Sub(now)
	fmt.Printf("time sub:%v \n", sub)
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	inLocation, err := time.ParseInLocation("2006-01-02 15:04:05", "2024-07-14 00:00:00", location)
	if err != nil {
		panic(err)
	}
	sub2 := inLocation.Sub(now)
	fmt.Printf("time sub2:%v \n", sub2)

}

func TestComputer(t *testing.T) {
	//now := time.Now()
	//now1 := time.Now()
	//fmt.Println(now == now1)
	location, _ := time.LoadLocation("Asia/Shanghai")
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2024-07-14 00:00:00", location)
	t2, _ := time.Parse("2006-01-02 15:04:05", "2024-07-14 00:00:00")
	fmt.Println(t1, t2)
	fmt.Println(t1.Equal(t2))
	fmt.Println(t1.Before(t2))
	fmt.Println(t1.After(t2))
}

func TestTicker(t *testing.T) {
	//tick := time.Tick(5 * time.Second)
	//for t2 := range tick {
	//	fmt.Println(t2)
	//}
	fmt.Println(time.Now())
	//time.AfterFunc(time.Second, func() {
	//	fmt.Println("5s之后执行")
	//})
	time.Sleep(5 * time.Second)
	fmt.Println(time.Now())
}
