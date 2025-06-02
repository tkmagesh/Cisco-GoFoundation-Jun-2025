package main

import "fmt"

func main() {

	var userName string
	fmt.Println("Enter the name :")
	fmt.Scanln(&userName)
	fmt.Printf("Hi %s, Have a nice day!\n", userName)

	/*
		var firstName, lastName string
		fmt.Println("Enter the fullname (seperated by space)")
		fmt.Scanln(&firstName, &lastName)
		fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
	*/

	// To accept inputs with spaces
	/*
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			txt := scanner.Text()
			fmt.Println(strings.ToUpper(txt))
			if txt == "exit" {
				break
			}
		}
		fmt.Println("Done!")
	*/
}
