package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	ch *CouponHandler
}

func NewServer(ch *CouponHandler) *Server {
	return &Server{ch}
}

func (s *Server) Route() chi.Router {
	r := chi.NewRouter()

	r.Get("/heartbeat", healthcheck)
	r.Get("/coupon", s.ch.getCoupon)

	return r
}

func (s *Server) Serve() {

	http.ListenAndServe(":9000", s.Route())
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	type (
		response struct {
			Message string `json:"message"`
		}
	)
	js, err := json.Marshal(response{"ok"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
