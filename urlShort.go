package main

import "fmt"
import "bytes"

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(s); i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func shortUrl(url string) string {
	hash_map := "abcdefghijklmnopqrstuvwxyz1234567890"
	var buffer bytes.Buffer
	for i := 0; i < len(hash_map); i++ {
		buffer.WriteString(string(hash_map[int(url[i])%32]))
	}
	return buffer.String()
}

func main() {
	fmt.Println(shortUrl("https://www.geeksforgeekzs.org/how-to-design-a-tiny-url-or-url-shortener/"))
	return
}
