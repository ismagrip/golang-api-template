package service

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

func getQueries(ctx context.Context) url.Values {
	queries := ctx.Value("queries").(queryParams)
	return queries.queries
}
func writeJSONResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Connection", "close")
	w.WriteHeader(status)
	w.Write(data)
}
