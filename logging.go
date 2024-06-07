package main

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/lipgloss"
	clog "github.com/charmbracelet/log"
)

type LogResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lmw *LogResponseWriter) WriteHeader(code int) {
	lmw.statusCode = code
	lmw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	style := lipgloss.NewStyle().Bold(true)
	return func(w http.ResponseWriter, r *http.Request) {
		lrw := LogResponseWriter{w, 0}
		next.ServeHTTP(&lrw, r)
		color := "4"
		switch group := lrw.statusCode - (lrw.statusCode % 100); group {
		case 200: // Success - green
			color = "2"
		case 300: // Redirect - pink
			color = "5"
		case 400: // Bad - yellow
			color = "3"
		case 500: // Server issue - red
			color = "1"
		default:
			color = "4"
		}
		style = style.Foreground(lipgloss.Color(color))
		clog.Info(style.Render(fmt.Sprintf("%d", lrw.statusCode)), "path", r.URL.Path, "method", r.Method)
	}
}
