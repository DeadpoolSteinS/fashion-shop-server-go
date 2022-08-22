package initializers

import (
	"fashion-shop/constant"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	constant.CheckError(err)
}
