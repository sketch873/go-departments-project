package environment

import "github.com/gofor-little/env"

func LoadEnv() {
	// Load an .env file and set the key-value pairs as environment variables.
	if err := env.Load("environment/.env.local"); err != nil {
		panic(err)
	}
}

func GetEnv(key string) string {

	// Get an environment variable's value with a default backup value.
	value := env.Get(key, "")

	return value
}