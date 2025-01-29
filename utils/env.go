package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadEnvFile(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		LogError(&err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()

		if strings.HasPrefix(str, "#") {
			continue
		}

		parts := strings.SplitN(str, "=", 2)
		if len(parts) == 2 {
			key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
			value = strings.Trim(value, "'\"")
			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func InitEnvs(keys []string, filepath string) (err error) {
	var missingKeys []string
	if filepath != "" {
		LoadEnvFile(filepath)
	}
	for _, k := range keys {
		if v := os.Getenv(k); v == "" {
			missingKeys = append(missingKeys, k)
		}
	}
	if len(missingKeys) > 0 {
		err = errors.New(fmt.Sprintf("Missing envs: %s", strings.Join(missingKeys, ", ")))
	}
	return
}

func GetEnvString(key string, def string) string {
	env := os.Getenv(key)
	if env == "" {
		env = def
	}
	return env
}

func GetEnvUint16(key string, def uint16) uint16 {
	v := os.Getenv(key)
	u16, err := strconv.ParseUint(v, 10, 16)
	if err != nil {
		return def
	}
	return uint16(u16)
}
