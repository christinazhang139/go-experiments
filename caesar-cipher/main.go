package main // 声明这个文件属于 main 包，Go 程序入口必须叫 main

import (
	"fmt"     // 导入 fmt 包，用来输出文字和格式化
	"strings" // 导入 strings 包，用来处理字符串
	"unicode" // 导入 unicode 包，用来判断字符类型
)

// caesarCipher 函数实现凯撒加密算法
// text: 要加密的文本
// shift: 位移量（正数加密，负数解密）
// 返回: 加密后的文本
func caesarCipher(text string, shift int) string {
	// 创建一个字符串构建器，用来高效地构建结果字符串
	var result strings.Builder

	// 遍历文本中的每个字符（rune 是 Go 中的 Unicode 字符类型）
	for _, char := range text {
		// 如果是字母，则进行位移加密
		if unicode.IsLetter(char) {
			// 判断是大写还是小写字母
			if unicode.IsUpper(char) {
				// 大写字母处理：A=65, B=66, ..., Z=90
				// 将字符转换为 0-25 的数字，加上位移量，再取模，最后转回字符
				shifted := (int(char-'A') + shift) % 26
				if shifted < 0 {
					shifted += 26 // 处理负数情况
				}
				result.WriteRune(rune(shifted + 'A'))
			} else {
				// 小写字母处理：a=97, b=98, ..., z=122
				shifted := (int(char-'a') + shift) % 26
				if shifted < 0 {
					shifted += 26 // 处理负数情况
				}
				result.WriteRune(rune(shifted + 'a'))
			}
		} else {
			// 如果不是字母（如空格、标点符号等），直接保持原样
			result.WriteRune(char)
		}
	}

	return result.String()
}

// encrypt 函数：加密文本
func encrypt(text string, shift int) string {
	return caesarCipher(text, shift)
}

// decrypt 函数：解密文本
func decrypt(text string, shift int) string {
	// 解密就是用负的位移量加密
	return caesarCipher(text, -shift)
}

// printMenu 函数：显示菜单
func printMenu() {
	fmt.Println("\n=== 凯撒加密程序 ===")
	fmt.Println("1. 加密文本")
	fmt.Println("2. 解密文本")
	fmt.Println("3. 演示示例")
	fmt.Println("4. 退出程序")
	fmt.Print("请选择操作 (1-4): ")
}

// demonstrateExample 函数：演示凯撒加密的例子
func demonstrateExample() {
	fmt.Println("\n=== 演示示例 ===")

	// 定义一些示例
	examples := []struct {
		text  string
		shift int
	}{
		{"Hello World!", 3},
		{"Go Programming", 5},
		{"Caesar Cipher", 13},
		{"ABC xyz", 1},
	}

	// 遍历示例并演示
	for i, example := range examples {
		fmt.Printf("\n示例 %d:\n", i+1)
		fmt.Printf("原文: %s\n", example.text)
		fmt.Printf("位移: %d\n", example.shift)

		encrypted := encrypt(example.text, example.shift)
		fmt.Printf("加密后: %s\n", encrypted)

		decrypted := decrypt(encrypted, example.shift)
		fmt.Printf("解密后: %s\n", decrypted)
	}
}

func main() { // main 函数是程序的入口
	fmt.Println("欢迎使用凯撒加密程序！")
	fmt.Println("这是一个用Go语言编写的简单加密解密工具。")

	// 主程序循环
	for {
		printMenu()

		// 声明变量来存储用户选择
		var choice int

		// 读取用户输入
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("输入错误，请输入数字 1-4")
			// 清理输入缓冲区
			fmt.Scanln()
			continue
		}

		// 根据用户选择执行相应操作
		switch choice {
		case 1:
			// 加密操作
			fmt.Print("请输入要加密的文本: ")
			var text string
			fmt.Scanln(&text)

			fmt.Print("请输入位移量 (1-25): ")
			var shift int
			fmt.Scanf("%d", &shift)

			// 确保位移量在有效范围内
			if shift < 1 || shift > 25 {
				fmt.Println("位移量必须在 1-25 之间")
				continue
			}

			encrypted := encrypt(text, shift)
			fmt.Printf("加密结果: %s\n", encrypted)

		case 2:
			// 解密操作
			fmt.Print("请输入要解密的文本: ")
			var text string
			fmt.Scanln(&text)

			fmt.Print("请输入位移量 (1-25): ")
			var shift int
			fmt.Scanf("%d", &shift)

			// 确保位移量在有效范围内
			if shift < 1 || shift > 25 {
				fmt.Println("位移量必须在 1-25 之间")
				continue
			}

			decrypted := decrypt(text, shift)
			fmt.Printf("解密结果: %s\n", decrypted)

		case 3:
			// 演示示例
			demonstrateExample()

		case 4:
			// 退出程序
			fmt.Println("感谢使用凯撒加密程序！再见！")
			return

		default:
			fmt.Println("无效选择，请输入 1-4")
		}
	}
}
