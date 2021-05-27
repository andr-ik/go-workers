package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
)

type Config struct {
	numWorker     int
	numWorkerLock sync.Mutex
}

func NewConfig() *Config {
	return &Config{
		numWorker:     0,
		numWorkerLock: sync.Mutex{},
	}
}

func (c *Config) GetNumWorker() int {
	return c.numWorker
}

func (c *Config) Update() error {
	pwd, _ := os.Getwd()
	data, err := ioutil.ReadFile(pwd + "/config/config.txt")
	if err != nil {
		fmt.Println("Error. Can't read config file.")

		return err
	}

	num, err := strconv.Atoi(string(data))
	value := c.validate(num)

	c.numWorkerLock.Lock()
	c.numWorker = value
	c.numWorkerLock.Unlock()

	return nil
}

func (c *Config) validate(num int) int {
	if num < 0 {
		return 0
	}

	if num > 100 {
		return 100
	}

	return num
}
