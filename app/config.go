package app

import (
	"os"
)

func Env(value string) (result string) {
	return os.Getenv(value)
}