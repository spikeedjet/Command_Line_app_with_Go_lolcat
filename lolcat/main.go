package main

import (
	"fmt"
	"math"
	// "strings"
	// "syreclabs.com/go/faker"
	"os"
	"bufio"
	"io"
)

// func rgb(i int) (int,int,int){
// 	var f = 0.1
// 	return int(math.Sin(f*float64(i)+0)*127 + 128),
// 		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
// 		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
// }
func rgb(i int) (int, int, int) {
    var f = 0.1
    return int(math.Sin(f*float64(i)+0)*127 + 128),
        int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
        int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

// func main() {

// 	var phrases  []string

// 	for i := 1; i<3; i++{
// 		phrases = append(phrases, faker.Hacker().Phrases()...)
// 	}

// 	output := strings.Join(phrases[:], "; ")	

// 	for j:=0; j<len(output); j++{
// 		// 这行代码使用 ANSI 转义序列来设置文本颜色并打印字符
// 		// \033[38;2;%d;%d;%dm 设置前景色为 RGB 值
// 		// %c 打印 output[j] 中的单个字符
// 		// \033[0m 重置所有文本属性
// 		r,g,b = rgb(j)
// 		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
// 	}
// 	fmt.Println()
// }
func main() {
    info, _ := os.Stdin.Stat()
    var output []rune

    if info.Mode()&os.ModeCharDevice != 0 {
        fmt.Println("The command is intended to work with pipes.")
        fmt.Println("Usage: fortune | gorainbow")
    }

    reader := bufio.NewReader(os.Stdin)
    for {
        input, _, err := reader.ReadRune()
        if err != nil && err == io.EOF {
            break
        }
        output = append(output, input)
    }

    print(output)
}

func print(output []rune){
	for j := 0; j < len(output); j++ {
        r, g, b := rgb(j)
        fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
    }
    fmt.Println()
}