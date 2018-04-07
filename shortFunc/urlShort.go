package shortFunc

import (
	"bytes"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	pool *redis.Pool
)

func init() {
	pool = &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(s); i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ShortUrl(url string, length int) string {
	c := pool.Get()
	defer c.Close()
	exists, shortUrl := checkIfExists(c, url)
	if exists {
		return shortUrl
	}
	shortUrl = createShortUrl(url, length)
	c.Do("SET", url, shortUrl)
	return shortUrl
}
func createShortUrl(url string, length int) string {
	hash_map := "abcdefghijklmnopqrstuvwxyz1234567890"
	var buffer bytes.Buffer
	string_length := len(url)
	var step int = 1
	if string_length > length {
		step = string_length / length
	} else {
		step = 1
	}

	for i := 0; i < string_length; i = i + step {
		buffer.WriteString(string(hash_map[int(int(url[i]%32)+i)%32]))
	}

	return buffer.String()
}
func checkIfExists(c redis.Conn, url string) (bool, string) {
	reply, err := redis.Bytes(c.Do("GET", url))
	if err != nil {
		panic(err)
	}
	if string(reply) == "" {
		return false, ""
	}
	return true, string(reply)
}
