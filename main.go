package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type RedisStatus struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func main() {
	redisConnection := flag.String("redis", "localhost:6379", "redis connection string")
	port := flag.Int("port", 3000, "application port")
	flag.Parse()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     *redisConnection,
		Password: "",
		DB:       0,
	})

	r := http.NewServeMux()
	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		_, err := redisClient.Ping().Result()
		w.Header().Add("Content-Type", "application/json")
		if err != nil {
			json.NewEncoder(w).Encode(RedisStatus{
				Status: "could not connect to redis server",
				Error:  err.Error(),
			})
			return
		}

		json.NewEncoder(w).Encode(RedisStatus{
			Status: "connection to redis is healthy",
		})

	})

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: r,
	}

	logrus.Infof("server starting on port %d", *port)
	if err := s.ListenAndServe(); err != nil {
		logrus.Fatalf("error stopping server: %s", err.Error())
	}
	logrus.Info("server stopped gracefully")
}
