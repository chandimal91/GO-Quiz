/*You cannot see the history of search words in the browser you are using. I thought it would be inconvenient not to be able to see the search word history, so I decided to create a function to see the search word history myself.

Search word history is created as follows.

If search word W has been entered before:
delete W in history.
Add a W to the beginning of the history.
If the search term W has not been entered before:
Add W to the beginning of the history.

Given N search words W, write a program that displays the history after N search words have been given.*/

package main

import (
	"fmt"
	"strings"
)
 
func main() {

	var n int
	var words,temp string
    fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		words = temp+" "+words
	}

	arr := strings.Fields(words)

	for i := 0; i < len(arr); i++ {
		if strings.Compare(arr[i],"0") == 0 {
			// To avoid the match same word matched before
			continue 
		} 
		for j:=i+1; j < len(arr); j++{
			if strings.Compare(arr[i],arr[j]) == 0 {
				// To remove the duplicate words
				arr[j] = "0";  
			}  
		}

		fmt.Print(arr[i]+"\n")
	}
}