package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"recommend/redis"
	"time"
)

type Article struct {
	Score  int64
	Member string
}

func main() {
	redis.InitRedis()
	http.HandleFunc("/mock", MockData)
	http.HandleFunc("/post", PostArticle)
	http.HandleFunc("/get", GetArticle)
	fmt.Println("listening...")
	http.ListenAndServe(":9000", nil)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func PostArticle(w http.ResponseWriter, r *http.Request) {
	redis.Conn.Do("set", "a", 1)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	data, err := redis.Conn.Do("get", "a")
	if err != nil {
		log.Print("err----", err)
	}
	w.Write(data.([]uint8))
}

func MockData(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 1000; i++ {
		article := &Article{
			Score:  time.Now().UnixNano(),
			Member: ID(10),
		}
		_, err := redis.Conn.Do("zadd", "article", article.Score, article.Member)
		if err != nil {
			log.Println("---------", err)
		}
	}
}

func ID(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
