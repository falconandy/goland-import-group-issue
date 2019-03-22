package main

import (
	"fmt"

	"github.com/falconandy/goland-import-group-issue/utils1"

	uuid "github.com/satori/go.uuid"

	"github.com/falconandy/goland-import-group-issue/utils2"
)

func main() {
	fmt.Println(uuid.NewV4())
	utils2.B()
	utils1.A()
}
