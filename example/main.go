package main

import (
	"fmt"

	".."
)

func main() {
	f, _ := faker.New()
	fmt.Println(f.Name().FirstName())
	fmt.Println(f.Name().JobTitle())
	fmt.Println(f.Name().FullName())
	fmt.Println(f.Lang("de").Name().Suffix())
	fmt.Println(f.Name().FullName())

}
