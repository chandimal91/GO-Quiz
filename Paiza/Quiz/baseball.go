/*Each baseball batter is out when he has three strikes and four balls when he has four balls.
An out or four-ball ends the batter's turn.

You are the referee who judges and calls strikes and balls.
Make appropriate calls according to the situation.

[Call list]
When 1-2 strikes are accumulated → "strike!"
When 3 strikes are accumulated → "out!" When
1-3 balls are accumulated → "ball!"
4 balls are accumulated When → "fourball!"*/

package main

import (
	"fmt"
	"strings"
)
 
func main() {

	var n,ball,strike int
	var words,temp string
    fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		words = words+" "+temp
	}

	arr := strings.Fields(words)

	for i := 0; i < len(arr); i++ {
		if strings.Compare(arr[i],"ball") == 0{
			ball++
			if ball == 4{
				fmt.Print("fourball!\n")
			}else{
				fmt.Print("ball!\n")
			}
		}else if strings.Compare(arr[i],"strike") == 0{
			strike++
			if strike == 3{
				fmt.Print("out!\n")
			}else{
				fmt.Print("strike!\n")
			}
		}
	}
}