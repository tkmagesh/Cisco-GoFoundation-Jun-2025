/* Modify the below to accept the "start" & "end" range from the user*/

package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the range :")
	fmt.Scanln(&start, &end)
LOOP:
	for no := start; no <= end; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		fmt.Println("Prime No :", no)
	}
}
