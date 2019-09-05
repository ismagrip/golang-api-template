package service

import (
	"context"
	"net/http"
	"net/url"
)

type queryParams struct {
	queries url.Values
}

func setQueriesOnCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := queryParams{}
		q.queries = r.URL.Query()

		ctx := context.WithValue(r.Context(), "queries", q)
		r.WithContext(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (q *queryParams) queriesExist(required ...string) bool {
	for _, query := range required {
		if _, ok := q.queries[query]; !ok {
			return false
		}
	}
	return true
}
