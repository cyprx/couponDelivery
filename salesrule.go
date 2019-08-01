package b52

import (
	"time"
)

type CouponType string

const (
	NoneCouponType    CouponType = "NONE"
	FixedCouponType   CouponType = "FIXED"
	DynamicCouponType CouponType = "DYNAMIC"
)

type Salesrule struct {
	ID                    uint
	Hash                  string
	Name                  string
	FromDate              time.Time
	ToDate                time.Time
	UsesPerCustomer       int
	UsageLimitPerCustomer int
	IsActive              bool
	CouponType            CouponType
}

type SalesruleRepository interface {
	Get(hash string) (*Salesrule, error)
}
