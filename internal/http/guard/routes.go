package guard

import (
	"log/slog"
	"net/http"

	templates "github.com/Polo123456789/control-visitas-garitas/internal/templates/guard"
)

func Handle(
	logger *slog.Logger,
) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /guard/{$}", func(w http.ResponseWriter, r *http.Request) {
		_ = templates.Dashboard().Render(r.Context(), w)
	})

	return mux
}
