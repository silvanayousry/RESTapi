package main

import (
	"fmt"

	"github.com/restapi/routes"
)

//"strconv"

func main() {

	fmt.Println("Running...")
	e := routes.Router()

	e.Logger.Fatal(e.Start(":1221"))

}
