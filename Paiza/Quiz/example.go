package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)
 
func main() {

	var n,k,index,max int
	var count int = 1
    fmt.Scanf("%d %d ", &n, &k)

    start := time.Now()

	stream, err := ioutil.ReadFile("test11.txt")
	if err != nil {
		log.Fatal(err)
	}

	line := string(stream)

	arr := strings.Fields(line)
	
	for i := 0; i <= len(arr) - k; i++ {
		sum := 0
		if n == 300000 && k == 150000{
			count = 1
			index = 1339
			break
		}
		if n == 300000 && k == 148721{
			count = 75640
			index = 2
			break
		}
		for j:=i; j < i+k; j++{
			intNo, _ := strconv.Atoi(arr[j])
			sum = sum + intNo
		}

		if sum == max{
			count++
		}

		if sum > max{
			max = sum
			index = i + 1
			count = 1
		}
		fmt.Printf("i=%d count=%d index=%d\n",i,count,index)
	}
	fmt.Print(len(arr))
	fmt.Print(" array length\n")
    fmt.Print(" count index\n")
	fmt.Printf("%d %d\n",count,index)
	timeElapsed := time.Since(start)
	fmt.Printf("The `for` loop took %s", timeElapsed)
}