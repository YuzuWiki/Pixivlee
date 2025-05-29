package testdata

import "github.com/joho/godotenv"

func init() {
	_ = godotenv.Load("./config.env")
}
