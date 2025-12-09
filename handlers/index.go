package handlers

import (
	"net/http"

	"github.com/ggof/template/views"
)

func GetIndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		views.Index().Render(r.Context(), w)
	}
}
