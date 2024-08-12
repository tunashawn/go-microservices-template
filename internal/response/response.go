package response

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type CustomResponse struct {
}

type ResponseData struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}
type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *CustomResponse) Success(data interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := ResponseData{
		Meta: Meta{
			Code:    http.StatusOK,
			Message: "ok",
		},
		Data: data,
	}

	_ = json.NewEncoder(w).Encode(res)

	c.logging(res, slog.LevelInfo, r)
}

func (c *CustomResponse) BadRequest(err error, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	res := ResponseData{
		Meta: Meta{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		},
	}

	_ = json.NewEncoder(w).Encode(res)

	c.logging(res, slog.LevelError, r)
}

func (c *CustomResponse) InternalServerError(err error, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	res := ResponseData{
		Meta: Meta{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		},
	}

	_ = json.NewEncoder(w).Encode(res)

	c.logging(res, slog.LevelError, r)
}

func (c *CustomResponse) Unauthorized(err error, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	res := ResponseData{
		Meta: Meta{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		},
	}

	_ = json.NewEncoder(w).Encode(res)

	c.logging(res, slog.LevelError, r)
}

func (c *CustomResponse) logging(res ResponseData, logLevel slog.Level, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	slog.LogAttrs(
		context.Background(),
		logLevel,
		res.Meta.Message,
		slog.Group("request",
			"method", r.Method,
			"uri", r.RequestURI,
			"request_body", body),
		slog.Group("response",
			"code", res.Meta.Code,
			"message", res.Meta.Message,
			"data", res.Data),
	)
}
