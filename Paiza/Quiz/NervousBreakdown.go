/*Let's play a simulation of a card game called Nervous Breakdown.
This time we only consider numbered playing cards, not jokers.

First, start by arranging the playing cards in a rectangular shape with H length and W width.
One of the numbers 1 to 13 is written on the H x W playing cards.
Also, there are multiple playing cards with the same number written on them.

He has N players, each numbered from 1 to N.
When the game starts, we will play in such a procedure from the first person.

・Choose two cards from the arranged cards and turn them over.
・If the two cards you turn over have different numbers written on them, it will be the next player's turn. If the numbers are the same, do the following:
・First, the two playing cards belong to the flipped player and are removed.
・The game ends when all the playing cards are removed.
・If there are trumps left, the same player returns to the first step and turns over the trumps.

Now let's say that player number N's ​​next player is his number 1 player.

A chronological record is given for the placement of the cards in the initial state of the game and the cards turned over by the end of the game.
Using that record, find the number of cards removed by each player.*/

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


