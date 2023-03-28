/*The season for the Paiza lottery has come again this year. Paisa's lottery ticket has his 6 digit number between 100000 and 199999. One winning number (100000 or more and 199999 or less) is announced every year, and the winning lottery number is determined as follows.

1st prize: A number that matches the
winning number Sequential prize: A number within ±1 of the winning number (If the winning number is 100000 or 199999, there is only one Sequential prize)
2nd prize: A number that matches the winning number in the last four digits (Excluding 1st prize numbers)
3rd prize: Numbers with the same last three digits as the winning number (excluding 1st and 2nd prize numbers)

For example, if the winning number is 142358, the winning numbers are: It looks like.

1st prize:
around 142358 Prizes: 142357 and 142359
2nd prize: 102358, 112358, 122358, …, 192358 (9 total)
3rd prize: 100358, 101358, 102358, …, 199358 (90 total)

n pieces you purchased For each lottery ticket number, write a program that outputs what kind of lottery you won.*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)
 
func main() {

	var n int
	var b,lot,words,temp string
	fmt.Scanln(&b)
    fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		words = words+" "+temp
	}

	arr := strings.Fields(words)

	for i := 0; i < len(arr); i++ {
		lot = arr[i]
		if first(b,lot) == 0{
			fmt.Print("first\n")
			continue
		}else if around(b,lot) == 0{
			fmt.Print("adjacent\n")
			continue
		}else if second(b,lot) == 0{
			fmt.Print("second\n")
			continue
		}else if third(b,lot) == 0{
			fmt.Print("third\n")
			continue
		}else{
			fmt.Print("blank\n")
		}
	}
}

func first(b,no string) int {
	if strings.Compare(b,no) == 0{
		return 0
	}
	return 1
}

func around(b,no string) int {

    intNo, _ := strconv.Atoi(no)
	intB, _ := strconv.Atoi(b)

	if strings.Compare(b,"100000") == 0 && strings.Compare(no,"100001") == 0{
		return 0
	}else if strings.Compare(b,"199999") == 0 && strings.Compare(no,"199998") == 0{
		return 0
	} else if intNo == intB + 1 || intNo == intB - 1{
		return 0
	}else{
		return 1
	}
}

func second(b,no string) int {
	strB := b[2:]
	strNo := no[2:]

	if strings.Compare(strB,strNo) == 0{
		return 0
	}
	return 1
}

func third(b,no string) int {
	strB := b[3:]
	strNo := no[3:]

	if strings.Compare(strB,strNo) == 0{
		return 0
	}
	return 1
}