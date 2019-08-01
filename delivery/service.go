package delivery

import (
	"errors"
	"strconv"

	"github.com/tikivn/b52"
)

type Service interface {
	Deliver(*b52.Salesrule) (*b52.Coupon, error)
}

type service struct {
	cache b52.CacheCouponRepository
}

func NewService(c b52.CacheCouponRepository) Service {
	return &service{c}
}

func (s *service) Deliver(sr *b52.Salesrule) (*b52.Coupon, error) {
	code, err := s.cache.RPop(strconv.Itoa(int(sr.ID)))
	if err != nil {
		return nil, err
	}

	if code == "" {
		return nil, errors.New("Coupon might not valid")

	}

	return &b52.Coupon{Salesrule: sr, Code: code}, nil
}
