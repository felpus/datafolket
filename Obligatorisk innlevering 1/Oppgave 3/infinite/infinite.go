package infinite

import "fmt"
import  "os"

func Infinite01 (msg string) {

	var input string

	for i := 0; ; i++ {

		fmt.Scan(&input)
		if input == "quit" {
			fmt.Println("Prosessen er avsluttet.")
			os.Exit(0)


		}
		if input == "term" {
			fmt.Println("Prosessen er terminert.")
			os.Exit(0)

		}

		fmt.Println(msg)

	}

}
