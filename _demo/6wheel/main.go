package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// 奖品中奖概率
type Prate struct {
	Rate  int    // 万分之 n 的中奖概率
	Total int    // 总数量限制，0 表示无限数量
	CodeA int    // 中奖概率起始编码（包含）
	CodeB int    // 中奖概率终止编码（包含）
	Left  *int32 // 剩余数
}

// 奖品列表
var prizeList []string = []string{
	"一等奖，火星一日游",
	"二等奖，地球半日游",
	"三等奖，月球一日游",
	"", //没有中奖
}

// 奖品的中奖概率设置，与上面的prizeList对应的设置
var left = int32(1000)
var rateList []Prate = []Prate{
	{100, 1000, 0, 9999, &left},
	//Prate{1, 1, 0, 0, 1},
	//Prate{2, 2, 1, 2, 2},
	//Prate{5, 10, 3, 5, 10},
	//Prate{100, 0, 0, 9999, 0},
}

var mu sync.Mutex = sync.Mutex{}

type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	app := newApp()
	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

// http://localhost:8080
func (c *lotteryController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return fmt.Sprintf("大转盘奖品列表：<br/>%s",
		strings.Join(prizeList, "<br/>\n"))
}
func (c *lotteryController) GetDebug() string {
	return fmt.Sprintf("获奖概率：%v\n", rateList)
}

func (c *lotteryController) GetPrize() string {
	//第一步，抽奖，需要根据随机数匹配奖品
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	code := r.Intn(10000)

	var myprize string
	var prizeRate *Prate
	for i, prize := range prizeList {
		rate := &rateList[i]
		if code >= rate.CodeA && code <= rate.CodeB {
			//满足中奖条件
			myprize = prize
			prizeRate = rate
			break
		}
	}
	if myprize == "" {
		myprize = "很遗憾，再来一次吧"
		return myprize
	}
	//第二部，中奖了，开始要发奖了
	if prizeRate.Total == 0 {
		//无限制奖品
		return myprize
	} else if *prizeRate.Left > 0 {
		left := atomic.AddInt32(prizeRate.Left, -1)
		if left >= 0 {
			log.Println("奖品：", myprize)
			return myprize
		}

	}
	myprize = "很遗憾，再来一次吧"
	return myprize

}
