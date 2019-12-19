package config

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/viper"
)

// NewViper is convenience way to init viper.
// It will traverse up to 10 level up to find an application.yml
func NewViper() *viper.Viper {
	v := viper.New()
	v.AutomaticEnv() // we are letting viper use env var if no application.yml found
	v.SetConfigType("yaml")

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := findFile(cwd, "application.yml", 10)
	if err != nil {
		log.Println(err.Error())
	} else {
		if err2 := v.ReadConfig(file); err2 != nil {
			log.Println(err2.Error())
		}
	}

	return v
}

func findFile(cwd string, name string, depth int) (io.Reader, error) {
	for i := 0; i < depth; i++ {
		file, err := os.Open(cwd + "/" + name)
		if err == nil {
			return file, nil
		}
		cwd = getParentDir(cwd)
	}

	return nil, fmt.Errorf("%s is not found after traversing %d directories from %s", name, depth, cwd)
}

func getParentDir(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if string(path[i]) == "/" {
			return path[0:i]
		}
	}
	return path
}
