package main

import (
	"context"
	"net/http"
	"time"
)

var service Service

type Service struct{}

func (s Service) GetUser(ctx context.Context, userId int) error {
	time.Sleep(time.Second)
	return ctx.Err()
}

func healthDeepHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := service.GetUser(ctx, 0); err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}
