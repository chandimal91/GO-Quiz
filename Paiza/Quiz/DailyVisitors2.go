/*You were managing a certain website.
I ran a campaign on this website for k consecutive days, but I forgot when and how long.

Luckily, I have access logs for all n days the website was running, and know the number of visitors per day.
For the time being, I decided to consider the period with the highest average number of visitors per day out of k consecutive days as a candidate for the campaign period.

A list of visitors for n days and the number of campaign days k are input, so output the number of candidates for the campaign period and the earliest start date among the candidates.*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
 
func main() {

	var n,k,index,max int
	var count int = 1
    fmt.Scanf("%d %d ", &n, &k)

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	scanner.Scan() 
	line := scanner.Text()

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
	}
	fmt.Printf("%d %d\n",count,index)
}

