package resource

import "fmt"

func HandleErr(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
