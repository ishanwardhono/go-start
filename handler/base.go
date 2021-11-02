package handler

import (
	"app/log"
	"context"
	"net/http"

	"app/helper/json"

	"github.com/gorilla/mux"
)

func Handle(f func(ctx context.Context, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := log.ContextStory(r.Context())
		log.Info("Request: ", GetLogRequest(r))
		f(ctx, rw, r)
	}
}

func GetLogRequest(r *http.Request) string {
	return "Host " + r.Host + ", Method:" + r.Method + ", Path:" + r.URL.RawPath
}

//
//Request Parsing function
//

func ParseBody(r *http.Request, v interface{}) error {
	return json.Decode(r.Body, v)
}

func GetQueryParam(r *http.Request, key string) string {
	queries := mux.Vars(r)
	return queries[key]
}

func GetURLParam(r *http.Request, key string) []string {
	return r.URL.Query()[key]
}
