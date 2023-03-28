/*You and your friends decide to do Shiritori with N people.
1st person, 2nd person,..., Nth person, 1st person, 2nd person,...

Here, each person must adhere to her four Shiritori rules when speaking.

1. An utterance must be one of the K words in the word list.
2. The first letter of any other person's statement must be the same as the last letter of the previous person's statement.
3. Do not say words that have been said before.
4. Don't say words ending in z.

Here, if the above rule is broken in the middle of speaking, the person who broke the rule will be removed from the Shiritori.
And we will continue to continue shiritori by pulling out the person. At this time, subsequent people do not have to follow rule 2.

She will be given M rows of logs that N people have performed Shiritori.
At this time, write a program that displays a list of people who have not dropped out from the Shiritori after M utterances.*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)
 
func main() {

	var n,k,m,out int
	var w,s,temp string
    fmt.Scanf("%d %d %d ", &n, &k, &m)

	for i := 0; i < k; i++ {
		fmt.Scan(&temp)
		w = w+" "+temp
	}

	for i := 0; i < m; i++ {
		fmt.Scan(&temp)
		s = s+" "+temp
	}

	words := strings.Fields(w)
	rounds := strings.Fields(s)

	for i := 0; i < len(rounds); i++ {
        // check rule 01
		if rule01(rounds[i],words) == 1{
			rounds[i] = "OUT"
			out++
			continue
		}

		// perform rule 03
		rule03(rounds[i],words)

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
	}
    rem := n-out
    fmt.Printf("%d\n",rem)
	// perform output 
	output(n,rounds)
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
}


