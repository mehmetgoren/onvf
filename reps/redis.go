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

func getRedisConnInfo() (string, int) {
	var err error
	host, port := "", 0
	h, p := flag.String("redis-host", "", "Redis Host Address"), flag.String("redis-port", "", "Redis Port Number")
	flag.Parse()

	eh := os.Getenv("REDIS_HOST")
	if len(eh) > 0 {
		host = eh
	} else if len(*h) > 0 {
		host = *h
	} else {
		host = "127.0.0.1"
	}

	ep := os.Getenv("REDIS_PORT")
	if len(ep) > 0 {
		port, err = strconv.Atoi(ep)
		if err != nil {
			port = 6379
			log.Println("An error occurred while converting Redis port value from environment variable: " + err.Error())
		}
	} else if len(*p) > 0 {
		port, err = strconv.Atoi(*p)
		if err != nil {
			port = 6379
			log.Println("An error occurred while converting Redis port value from arguments :" + err.Error())
		}
	} else {
		port = 6379
	}

	return host, port
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

	host, port := getRedisConnInfo()
	r.Host = host
	log.Println("Redis host: ", host)
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
