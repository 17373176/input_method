// Package main is special.  It defines a
// standalone executable program, not a library.
// Within package main the function main is also
// special—it’s where execution of the program begins.
// Whatever main does is what the program does.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"input_method/controller"
	"input_method/library"
)

// main the function where execution of the program begins
func main() {
	// 注册，启动日志
	library.LogService = library.NewLog()

	im := controller.NewInputMethod(os.Args[1:])
	loop(im)

	// 关闭日志文件
	library.LogService.CloseLog()
}

// loop loop input method
func loop(im *controller.InputMethod) {
	stdin := bufio.NewReader(os.Stdin)
	for {
		spell, err := stdin.ReadString('\n')
		if err != nil {
			break
		}
		// 去掉所有换行
		spell = strings.TrimRight(spell, "\n")
		words := im.FindWords(spell)
		fmt.Println(strings.Join(words, ", "))
	}
}
