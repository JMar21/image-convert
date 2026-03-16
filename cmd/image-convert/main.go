package main

import (
	"fmt"
	"os"

	conv "github.com/JMar21/image-convert/internal/image-convert"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No arguments provided.")
		return
	}
	if len(args) > 3 {
		fmt.Println("Too many arguments provided.")
		return
	}
	from := args[0]
	to := args[1]
	path := args[2]
	fmt.Println("Arguments provided:")
	for _, arg := range args {
		fmt.Println(arg)
	}
	if from == "jpg" && to == "png" {
		err := conv.ConvertJPGToPNG(path)
		if err != nil {
			fmt.Printf("Error converting image: %v\n", err)
		} else {
			fmt.Println("Image converted successfully.")
		}
	}
}
