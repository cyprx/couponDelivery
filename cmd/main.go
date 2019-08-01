package main

import (
	"github.com/tikivn/b52/delivery"
	"github.com/tikivn/b52/http"
	"github.com/tikivn/b52/redis"
)

func main() {

	redis.Init()
	rc := redis.GetClient()

	cache := redis.NewCouponRepository(rc)
	ds := delivery.NewService(cache)
	ch := &http.CouponHandler{ds}

	s := http.NewServer(ch)
	s.Serve()

}
