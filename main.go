package main

import (
	"gin_Ranking/router"
)

func main() {

	r := router.Router()
	err := r.Run(":8888")
	if err != nil {
		return
	}
}
