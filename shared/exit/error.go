package exit

import (
	"fmt"
	"os"
)

func If(err error) {
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
