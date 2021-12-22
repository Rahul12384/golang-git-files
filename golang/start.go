package main

import "fmt"


func main()  {
	var abc string;
	fmt.Printf("welcome to golang project\nEnter your name: ");		
	fmt.Scan(&abc);
	fmt.Println("welcome "+abc);	
}