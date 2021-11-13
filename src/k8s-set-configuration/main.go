package main

import (
	"fmt"
	"os"
)

func main()  {
	var homeDir, _ = os.UserHomeDir()
	fmt.Println("Home:" +homeDir)

}