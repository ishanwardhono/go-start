package handler

import (
	"app/errors"
	"app/log"
	"context"
	"net/http"

	"app/helper/json"

	"github.com/gorilla/mux"
)

type HttpHandler interface {
	RegisterHandlers(router *mux.Router)
}

type response struct {
	StatusCode   int
	Data         interface{}
	ErrorMessage string
}

func Handle(f func(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := log.ContextStory(r.Context())
		log.Info(ctx, "Request: ", GetLogRequest(r))
		result, err := f(ctx, w, r)
		WriteResponse(ctx, w, r, result, err)
	}
}

func WriteResponse(ctx context.Context, w http.ResponseWriter, r *http.Request, data interface{}, err error) {
	w.Header().Add("Content-Type", "Application/json")
	res := response{
		Data:       data,
		StatusCode: http.StatusOK,
	}

	if err != nil {
		if errs, ok := err.(*errors.Errs); ok {
			res.StatusCode = errs.Code
			res.ErrorMessage = errs.Error()
		} else {
			res.StatusCode = http.StatusInternalServerError
			res.ErrorMessage = err.Error()
		}
	}

	byteData, err := json.Marshal(res)
	if err != nil {
		log.Error(ctx, "Error marshalling response")
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
