package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	maxNum := 10
	secretNumber := rand.Intn(maxNum)
	//fmt.Println("随机数是:", secretNumber)

	fmt.Println("请输入你的猜测")
	reader := bufio.NewReader(os.Stdin)

	for {
		//
		// 读取一行
		// 后面会跟着一个换行符
		// bufio 比较曲折 但是后面 会用到
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取数据产生错误", err)
			return
		}
		//
		// 用string包去掉换行符
		// 不然Atoi会报错
		input = strings.TrimSuffix(input, "\n")
		//
		// Atoi转成数字
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("无效输入 请输入整数")
		}
		fmt.Println("你猜的数字是:", guess)
		if guess > secretNumber {
			fmt.Println("你猜的数字大了")
		} else if guess < secretNumber {
			fmt.Println("你猜的数字小了")
		} else {
			fmt.Println("猜对啦 !!!")
			break
		}
	}

	// 复习了 很多基础语法

}
