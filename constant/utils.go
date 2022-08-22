package constant

import "log"

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
