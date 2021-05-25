package utils

import "fmt"
import "os"

func GetEnv(key string) (string) {
	result := os.Getenv(key)

	fmt.Printf("Environment variable [%s] resolved to [%s]\n", key, result)

	return result
}
