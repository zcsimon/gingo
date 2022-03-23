package utils

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Env struct{}

func NewConfig() *Env {
	return &Env{}
}

var env map[string]string

func InitConfig() map[string]string {
	env = make(map[string]string)

	f, err := os.Open(".env")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		env[key] = value
	}
	return env

}

func (e *Env) GetEnv(item string) string {
	if data, ok := env[item]; ok {
		return data
	}
	return ""
}
