package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)
 
func main() {

	var n,k,m,out int

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

	fmt.Printf("%d ",n)
	fmt.Printf("%d ",k)
	fmt.Printf("%d \n",m)
	fmt.Print("\n")
	fmt.Print(words)
	fmt.Print("\n")
	fmt.Print(rounds)
	fmt.Print("\n")

	for i := 0; i < len(rounds); i++ {
        // check rule 01
		if rule01(rounds[i],words) == 1{
			rounds[i] = "OUT"
			out++
			continue
		}
        // check rule 02
		if i != 0 && rounds[i-1] != "OUT" && rule02(rounds[i-1],rounds[i]) == 1{
			rounds[i] = "OUT"
			out++
			continue
		}
		// check rule 04
		if rule04(rounds[i]) == 1{
			rounds[i] = "OUT"
			out++
			continue
		}
		// perform rule 03
		rule03(rounds[i],words)

	}
    rem := n-out
    fmt.Printf("%d\n",rem)
	output(n,rounds)
	fmt.Print("\n")
	fmt.Print(words)
	//test(words, rounds)
}

func rule01(str string, s []string) int {
	for _, v := range s {
		if v == str {
			return 0
		}
	}
	return 1
}

func rule02(str1 string, str2 string) int {
	if str1[len(str1)-1] == str2[0]{
		return 0
	}
	return 1
}

func rule04(str string) int {
	if str[len(str)-1] != 'z'{
		return 0
	}
	return 1
}

func rule03(str string, s []string) {
	for i, v := range s {
		if v == str {
			// remove already matched word from the list 
			//fmt.Printf("index %d was %s \n",i,v)
		   s[i] = "0"
		}
	}
}

func output(n int, s []string ){
    var str,str1 string
	var player int

	for i:=1; i < n+1; i++{
		temp := strconv.Itoa(i)
		str = str+" "+temp
	}

	p := strings.Fields(str)

	for i, v := range s {
		if v == "OUT" {
			// remove out players
		   //p[i%n] = "OUT"
		   temp := strconv.Itoa(i)
		   str1 = str1+" "+temp
		}
	}

	
	o := strings.Fields(str1)

	for i, v := range o {
		if i == 0{
			index,_ := strconv.Atoi(v)
			player = index % n
			p[player] = "OUT"
		}else{
			current,_ := strconv.Atoi(v)
			previous,_ := strconv.Atoi(o[i-1])
			diff := current - previous
			start := player
			player = (start + diff - ((diff/n)*i))%n
			p[player] = "OUT"
		}
	}

	for _, v := range p {
		if v != "OUT" {
		   fmt.Printf("%s\n",v)
		}
	}
	fmt.Print("\n")
	fmt.Print(p)
	fmt.Print("\n")
	fmt.Print(s)
	fmt.Print("\n")
	fmt.Print(o)
}

func test(w []string, r []string){
	var count int
	for _, v := range w {
		if v == "0" {
		   count++
		}
	}

	fmt.Printf("count %d\n",count)
	fmt.Printf("word length %d\n",len(w))
	fmt.Printf("round length %d\n",len(r))

}


