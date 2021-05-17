package bscscan

import "fmt"
import "os"

func getEnv(key string) (string) {
	result := os.Getenv(key)

	fmt.Printf("Environment variable [%s] resolved to [%s]\n", key, result)

	return result
}
