package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var (
	ErrCacheMiss = errors.New("cache miss")
	sf           = singleflight.Group{}
)

func getData(key string) (string, error) {
	data, err := getFromCache("key")
	if err != nil {
		if errors.Is(err, ErrCacheMiss) {
			// cache miss get from db
			res, err, _ := sf.Do(key, func() (interface{}, error) {
				dbData, err := getFromDB("key")
				if err != nil {
					return "", err
				}
				// set cache
				setCache(key, dbData)
				return dbData, nil
			})
			if err != nil {
				return "", err
			}

			return res.(string), nil
		}
		return "", err
	}

	return data, nil
}

func setCache(key, data string) {
}

func getFromCache(key string) (string, error) {
	return "", ErrCacheMiss
}

func getFromDB(key string) (string, error) {
	fmt.Println("query from db")
	res := strconv.Itoa(int(time.Now().UnixNano()))
	return res, nil
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data, err := getData("key")
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(data)
		}()
	}

	wg.Wait()
}
