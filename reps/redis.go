package reps

import (
	"encoding/json"
	"flag"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"strconv"
	"sync"
)

const (
	MAIN     = 0
	RQ       = 1
	EVENTBUS = 15
)

type RedisConnectionInfo struct {
	mu          sync.Mutex
	Host        string
	Port        int
	initialized bool
}

func (r *RedisConnectionInfo) init() {
	if r.initialized {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.initialized {
		return
	}

	host, port := "127.0.0.1", 6379
	h, p := flag.String("host", "", "Redis Host Address"), flag.String("port", "", "Redis Port Number")
	flag.Parse()

	if len(*h) > 0 {
		host = *h
	} else {
		host := os.Getenv("REDIS_HOST")
		if len(host) == 0 {
			host = "127.0.0.1"
		}
	}
	r.Host = host
	log.Println("Redis host: ", host)

	portStr := ""
	if len(*p) > 0 {
		portStr = *p
	} else {
		portStr = os.Getenv("REDIS_PORT")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 6379
		log.Println("An error occurred while converting Redis port value:" + err.Error())
	}
	r.Port = port
	log.Println("Redis port: ", port)
	r.initialized = true
}

var _rci = RedisConnectionInfo{}

func createRedisConnection(db int) *redis.Client {
	_rci.init()

	return redis.NewClient(&redis.Options{
		Addr:     _rci.Host + ":" + strconv.Itoa(_rci.Port),
		Password: "", // no password set
		DB:       db, // use default DB
	})
}

func Map(in interface{}) map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(in)
	json.Unmarshal(inrec, &inInterface)
	return inInterface
}
