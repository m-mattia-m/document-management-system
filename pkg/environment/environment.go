package environment

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string) (string, error) {
	env, err := os.LookupEnv(key)
	if !err || env == "" {
		return env, errors.New("environment variable '" + key + "' empty or not found")
	}
	return env, nil
}

func init() {
	log.Println("[INIT MAIN]: Status -> Main-init is loaded")
	godotenv.Load()
}
