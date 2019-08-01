package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/tikivn/b52"
	"github.com/tikivn/b52/delivery"
)

type CouponHandler struct {
	DeliveryService delivery.Service
}

func (h *CouponHandler) getCoupon(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	sID, err := strconv.Atoi(q["salesruleID"][0])
	log.Println(q)
	sr := &b52.Salesrule{ID: uint(sID)}
	coupon, err := h.DeliveryService.Deliver(sr)
	log.Println(coupon)
	log.Println(err)
	if err != nil || coupon == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	type response struct {
		Code string `json:"code"`
	}
	js, err := json.Marshal(response{Code: coupon.Code})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
