package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== 欢迎来到猜数字游戏！ ===")
	fmt.Println("我已经想好了一个1到100之间的数字")
	fmt.Println("请你来猜猜看！")

	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成1到100之间的随机数
	target := rand.Intn(100) + 1

	var guess int
	attempts := 0

	// 游戏主循环
	for {
		fmt.Print("\n请输入你的猜测 (1-100): ")

		// 读取用户输入
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("输入无效，请输入一个整数！")
			// 清理输入缓冲区
			fmt.Scanln()
			continue
		}

		attempts++

		// 检查输入范围
		if guess < 1 || guess > 100 {
			fmt.Println("请输入1到100之间的数字！")
			continue
		}

		// 判断猜测结果
		if guess < target {
			fmt.Printf("太小了！再试一次 (第%d次尝试)\n", attempts)
		} else if guess > target {
			fmt.Printf("太大了！再试一次 (第%d次尝试)\n", attempts)
		} else {
			fmt.Printf("🎉 恭喜你！猜对了！\n")
			fmt.Printf("答案是: %d\n", target)
			fmt.Printf("你总共尝试了 %d 次\n", attempts)

			// 评价表现
			switch {
			case attempts <= 3:
				fmt.Println("太厉害了！你真是个天才！")
			case attempts <= 7:
				fmt.Println("不错哦！你的运气很好！")
			case attempts <= 10:
				fmt.Println("还可以，继续努力！")
			default:
				fmt.Println("多练习一下，你会更好的！")
			}

			break
		}

		// 给一些提示
		if attempts >= 5 {
			if target <= 25 {
				fmt.Println("💡 小提示：数字比较小哦")
			} else if target <= 50 {
				fmt.Println("💡 小提示：数字在中等范围")
			} else if target <= 75 {
				fmt.Println("💡 小提示：数字有点大")
			} else {
				fmt.Println("💡 小提示：数字很大哦")
			}
		}
	}

	fmt.Println("\n感谢游戏！再见！")
}
