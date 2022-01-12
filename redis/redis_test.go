package redis_test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func Test1(t *testing.T) {
	script := "if (redis.call('llen', KEYS[1]) > 5) then " +
		"redis.call('rpop', KEYS[1]) " +
		"end " +
		"redis.call('lpush', KEYS[1], ARGV[1])"
	rdb.Eval(ctx, script, []string{"list"}, time.Now().Unix())
}
