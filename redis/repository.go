package redis

import (
	"time"

	"github.com/tikivn/b52"

	"github.com/go-redis/redis"
)

type couponRepository struct {
	rc *redis.Client
}

func NewCouponRepository(rc *redis.Client) b52.CacheCouponRepository {
	return &couponRepository{rc}
}

func (r *couponRepository) WriteBatch(
	size int,
	key string,
	val b52.CouponIterator,
	exp time.Duration,
) error {
	pipeline := r.rc.Pipeline()
	i := 0
	for val.Next() {
		var code string
		err := val.Scan(&code)
		if err != nil {
			return err
		}
		pipeline.LPush(key, code)

		i++
		if i%size == 0 {
			_, err := pipeline.Exec()
			if err != nil {
				return err
			}
		}
	}

	pipeline.Expire(key, exp)

	_, err := pipeline.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (r *couponRepository) LPush(key string, val string) error {
	return r.rc.LPush(key, val).Err()
}

func (r *couponRepository) RPop(key string) (string, error) {
	return r.rc.RPop(key).Result()
}
