package api

import (
	"net/http"

	"github.com/abihf/entitle/webhook"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	webhook.Handle(w, r)
}
