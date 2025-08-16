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

	mux.HandleFunc("POST /guard/validate", func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")

		if code == "11cff407-ef97-4631-ab40-e063127226db" {
			_ = templates.ValidCodeDialog().Render(r.Context(), w)
			return
		}

		_ = templates.InvalidCodeDialog("Motivo XD").Render(r.Context(), w)
		return
	})

	return mux
}
