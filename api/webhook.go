package api

import (
	"net/http"

	"github.com/abihf/entitle"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	entitle.HandleWebhook(w, r)
}
