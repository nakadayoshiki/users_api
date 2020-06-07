package main

import "github.com/nakadayoshiki/users_api/types"

func main() {
	db := types.Init()
	defer db.Close()
}
