package b52

import (
	"time"
)

type Coupon struct {
	Salesrule *Salesrule
	Code      string
}

type CouponIterator interface {
	Next() bool
	Scan(...interface{}) error
}

type DBCouponRepository interface {
	WriteBatch(int, string, CouponIterator, time.Duration) error
	GetAll(salesruleID uint) CouponIterator
}

type CacheCouponRepository interface {
	WriteBatch(int, string, CouponIterator, time.Duration) error
	LPush(string, string) error
	RPop(string) (string, error)
}
