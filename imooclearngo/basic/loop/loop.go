package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 省略初始条件，相当于while
func convertToBin(n int) string {
	if n == 0 {
		return "0"
	}
	res := ""
	for ; n > 0; n /= 2 {
		res = strconv.Itoa(n%2) + res
	}
	return res
}

// 省略初始条件和递增条件，相当于while
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 省略结束条件，无限循环
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1011 --> 1101
		convertToBin(0),
		convertToBin(4),
	)
	printFile("abc.txt")

	s := `abc"d"
	kkk
	123
	p`

	printFileContents(strings.NewReader(s))
	// forever()
}
