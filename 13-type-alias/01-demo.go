package main

import "fmt"

// type alias
type MyString string

func (s MyString) Length() int {
	return len(s)
}

func main() {
	/*
		var str MyString
		str = "Velit nulla dolore incididunt est ea culpa minim aliqua aliquip est mollit reprehenderit dolor aliquip."
	*/

	str := MyString("Velit nulla dolore incididunt est ea culpa minim aliqua aliquip est mollit reprehenderit dolor aliquip.")
	fmt.Println(str.Length())
}
