package users

import (
	"fmt"

	"github.com/sanikum3/ORP/internal/storage"
)

func List() {

	if len(storage.Users) == 0 {
		fmt.Println()
		fmt.Println("No users found")
		fmt.Println()
		return
	}

	fmt.Println()

	fmt.Printf(
		"%-20s %-40s %-64s\n",
		"USERNAME",
		"ROOM",
		"KEY",
	)

	fmt.Println(
		"----------------------------------------------------------------------------------------------------------------------------",
	)

	for _, user := range storage.Users {
		fmt.Printf(
			"%-20s %-40s %-64s\n",
			user.Username,
			user.Room,
			user.Url,
		)
	}
	fmt.Println()
	fmt.Println(
		"----------------------------------------------------------------------------------------------------------------------------",
	)
}
