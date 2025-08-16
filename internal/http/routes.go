package http

import (
	"log/slog"
	"net/http"

	"github.com/Polo123456789/control-visitas-garitas/internal/http/guard"
	"github.com/Polo123456789/control-visitas-garitas/internal/http/neighbor"
	"github.com/Polo123456789/control-visitas-garitas/internal/templates"
)

func setupRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content := templates.Index()
		content.Render(r.Context(), w)
	})

	mux.Handle("/neighbor/", neighbor.Handle(logger))
	mux.Handle("/guard/", guard.Handle(logger))
}
