package neighbor

import (
	"log/slog"
	"net/http"

	templates "github.com/Polo123456789/control-visitas-garitas/internal/templates/neighbor"
)

func Handle(
	logger *slog.Logger,
) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /neighbor/{$}", func(w http.ResponseWriter, r *http.Request) {
		_ = templates.Dashboard().Render(r.Context(), w)
	})

	mux.HandleFunc("/neighbor/create-visit", func(w http.ResponseWriter, r *http.Request) {
		_ = templates.ShareCode().Render(r.Context(), w)
	})

	return mux
}
