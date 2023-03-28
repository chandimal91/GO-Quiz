/*You are given a string of English words separated by spaces.
Please output the number of occurrences of English words contained in the English word string in the order in which they appear.*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
 
func main() {

	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() 
    line := scanner.Text()

	arr := strings.Fields(line)
    for i := 0; i < len(arr); i++ {
		if strings.Compare(arr[i],"0") == 0 {
			// To avoid the count same word counted before
			continue 
		} 
		count := 1
		for j:=i+1; j < len(arr); j++{
			if strings.Compare(arr[i],arr[j]) == 0 {
				count++
				// To avoid the count same word counted before
				arr[j] = "0";  
			}  
		}
		fmt.Print(arr[i]+" "+strconv.Itoa(count)+"\n")
	}
}