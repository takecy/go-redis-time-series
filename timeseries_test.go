package timeseries

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis"
)

func TestNew(t *testing.T) {
	var rcli Redis

	rcli = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", "localhost", 6379),
		DB:   0,
	})
}

type RedisTest struct {
	*redis.Client
}

func newRedis() *RedisTest {
	rcli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", "localhost", 6379),
		DB:   0,
	})

	return &RedisTest{rcli}
}
