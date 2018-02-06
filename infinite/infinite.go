package infinite

import "fmt"
import "time"

func Infinite01 (msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}

}
