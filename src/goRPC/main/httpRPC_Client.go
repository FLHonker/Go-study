package main

/*
Usage: httpRPC_Client localhost
 */

import (
	"os"
	"fmt"
	"net/rpc"
	"log"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], " server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]
	client, err := rpc.DialHTTP("tcp", serverAddress + ":8811")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//Synchronous
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d ... %d\n", args.A, args.B, quot.Quo, quot.Rem)

}

/*
初次相识，
我像个手忙脚乱的孩子
看着漂亮的雪
喜爱，又不敢触碰
每天都奢求和你一起

我努力成长
希望雪能看透我的心思
捧在手里，很快化了
飘落发间，无影无踪
抬头望去，
她又像是在飞舞逗我开心

渐渐
我学会了去欣赏
才懂得她真正的纯洁与可爱
越发努力期待能与她共舞

殊不知
我的心
早已被她偷走
不是上天注定
是心甘情愿

暴风雨来临时
我不顾一切想啊哟保护她
有时
我的自私
想把他捧在手里
就会伤到她

或许
我应该以更加宽广的胸怀
在风中与她共舞
深拥
以更加有力的臂膀
保护她
以更加细腻的心思
守护她
以更加独特的眼光
欣赏她

相爱不易
手忙家乱的男孩
经历了这么多风雨洗礼
终究还是会手忙焦
他害怕
他珍惜
他不舍
他小心翼翼地呵护着

他忘不掉
风雨中雪对他的鼓励
烈日下雪耐心的给他降暑
他跌倒了
雪就是他爬起来的动力
她在空中招手
像灯塔
像触手可见的黎明

他努力把自己变得更好
身上的责任也越来越重
因为前方还有更烈的太阳
更高的山，更坎坷的路

雪，
依旧美丽动人。
但，
男孩有一丝愧疚
雪陪伴了2年风风雨雨
他还没学会与她共舞

他心存感激
难忘
也不会忘记最初
雪没有嫌弃这个手忙脚乱的笨蛋

*/