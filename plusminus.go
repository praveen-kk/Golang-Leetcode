package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here

	postive := 0
	negative := 0
	zeroes := 0
	l := len(arr)
	for _, v := range arr {
		if v > 0 {
			postive++
			fmt.Println(postive)
		} else if v < 0 {
			negative++
		} else if v == 0 {
			zeroes++
		}
	}
	fmt.Println(l)
	postiveprop float64() := postive / l
	negativeprop := negative / l
	Zeroesprop := zeroes / l
	fmt.Printf("%d\n%d\n%d\n", postiveprop, negativeprop, Zeroesprop)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
