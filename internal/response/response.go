package response

import (
	"encoding/json"
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

func (r CustomResponse) Success(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(ResponseData{
		Meta: Meta{
			Code:    http.StatusOK,
			Message: "ok",
		},
		Data: data,
	})
	if err != nil {
		http.Error(w, "could not encode success response", http.StatusInternalServerError)
	}
}

func (r CustomResponse) BadRequest(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	err = json.NewEncoder(w).Encode(ResponseData{
		Meta: Meta{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		},
	})
	if err != nil {
		http.Error(w, "could not encode success response", http.StatusInternalServerError)
	}
}

func (r CustomResponse) InternalServerError(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	err = json.NewEncoder(w).Encode(ResponseData{
		Meta: Meta{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		},
	})
	if err != nil {
		http.Error(w, "could not encode success response", http.StatusInternalServerError)
	}
}

func (r CustomResponse) Unauthorized(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	err = json.NewEncoder(w).Encode(ResponseData{
		Meta: Meta{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		},
	})
	if err != nil {
		http.Error(w, "could not encode success response", http.StatusInternalServerError)
	}
}
