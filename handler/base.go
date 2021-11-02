package handler

import (
	"app/log"
	"context"
	"net/http"

	"app/helper/json"

	"github.com/gorilla/mux"
)

type response struct {
	StatusCode   int
	Data         interface{}
	ErrorMessage string
}

func Handle(f func(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := log.ContextStory(r.Context())
		//TODO: add context to logger
		log.Info("Request: ", GetLogRequest(r))
		result, err := f(ctx, w, r)
		WriteResponse(w, r, result, err)
	}
}

func WriteResponse(w http.ResponseWriter, r *http.Request, data interface{}, err error) {
	w.Header().Add("Content-Type", "Application/json")
	res := response{
		Data:       data,
		StatusCode: 200,
	}

	if err != nil {
		//TODO: define error code
		res.StatusCode = 400
		res.ErrorMessage = err.Error()
	}

	byteData, err := json.Marshal(res)
	if err != nil {
		log.Error("Error marshalling response")
	}
	w.Write(byteData)
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
