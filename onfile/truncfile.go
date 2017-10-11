package main

import (
		"fmt"
		"os"
)

//var err error
//Difference between = and :=
// = is for assignment only
// := is for declaration plus assignment

func main() {
		//err = os.Truncate("test.txt", 100)
		err := os.Truncate("test.txt", 100)
    if err != nil {
				fmt.Println(err)
		}
}

