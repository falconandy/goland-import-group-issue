package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"

	"github.com/falconandy/goland-import-group-issue/utils2"
)

func main() {
	fmt.Println(uuid.NewV4())
	utils2.B()
	// Try to write and confirm a suggested import: utils1.A()
}
