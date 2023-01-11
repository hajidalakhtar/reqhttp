package exception

import (
	"fmt"
	"go-http-cli/color"
)

func ConnectionFailed(err error) {
	fmt.Println(color.Red + "Failed connect to URL\n" + err.Error())
}
