package speak

import (
	"fmt"

	"jk404.cn/greetings"
)

func Speak()  {
	message := greetings.Hello("Flora")
	fmt.Println(message)
}