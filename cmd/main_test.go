package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/tikivn/b52/redis"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func TestMain(t *testing.T) {
	redis.Init()
	rc := redis.GetClient()
	repo := redis.NewCouponRepository(rc)

	for i := 0; i < 1000; i++ {
		repo.LPush("123", StringWithCharset(10, charset))
	}
}
