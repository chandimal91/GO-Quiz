package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)
 
func main() {

	var n,k,m int
    //fmt.Scanf("%d %d %d ", &n, &k, &m)

	stream, err := ioutil.ReadFile("6_6.txt")
	if err != nil {
		log.Fatal(err)
	}

	line := string(stream)
	arr := strings.Fields(line)
	n,_ = strconv.Atoi(arr[0])
	k,_ = strconv.Atoi(arr[1])
	m,_ = strconv.Atoi(arr[2])
	words := arr[3:k+3]
	rounds := arr[k+3:]
	fmt.Printf("%d\n",n)
	fmt.Printf("%d\n",k)
	fmt.Printf("%d\n",m)
	fmt.Print(words)
	fmt.Print(rounds)
}