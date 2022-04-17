package http

import (
	"app/core/errors"
	"app/core/log"
	"context"
	"net/http"
	"time"

	"app/core/helper/json"

	"github.com/gorilla/mux"
)

type HttpHandler interface {
	RegisterHandlers(router *mux.Router)
}

type Response struct {
	StatusCode   int         `json:"status"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage string      `json:"errorMessage,omitempty"`
}

func Handle(f func(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := log.ContextStory(r.Context())
		log.Info(ctx, "HTTP Request: ", GetLogRequest(r))
		startTime := time.Now()
		result, err := f(ctx, w, r)
		WriteResponse(ctx, w, r, result, err, startTime)
	}
}

func WriteResponse(ctx context.Context, w http.ResponseWriter, r *http.Request, data interface{}, err error, startTime time.Time) {
	w.Header().Add("Content-Type", "Application/json")
	var res Response

	if customResponse, ok := data.(Response); ok {
		res = customResponse
	} else {
		res.Data = data
		res.StatusCode = http.StatusOK
	}

	if err != nil {
		if errs, ok := err.(*errors.Errs); ok {
			w.WriteHeader(errs.Code)
			res.StatusCode = errs.Code
			res.ErrorMessage = errs.Error()
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			res.StatusCode = http.StatusInternalServerError
			res.ErrorMessage = err.Error()
		}
	}

	byteData, err := json.Marshal(res)
	if err != nil {
		log.Error(ctx, "Error marshalling response")
	}
	log.Infof(ctx, "HTTP Response Code: %d, execution time: %dms", res.StatusCode, time.Since(startTime).Milliseconds())
	w.Write(byteData)
}

func GetLogRequest(r *http.Request) string {
	return "Host " + r.Host + ", Method:" + r.Method + ", Path:" + r.URL.Path
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
