/*
Hint : Use strconv.Atoi() to convert numbers in string format to number format
*/
package main

import "fmt"

func main() {
	fmt.Println(sum(10, 20))         //=> 30
	fmt.Println(sum(10, "20", 30))   //=> 60
	fmt.Println(sum(10, 20, 30, 40)) //=> 100
	fmt.Println(sum(10, "abc"))      // => 10
	fmt.Println(sum())               // => 0
}

func sum( /* ? */ ) int {
	/* ? */
}
