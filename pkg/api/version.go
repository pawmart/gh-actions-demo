package api

import (
	"net/http"

	"github.com/stefanprodan/gh-actions-demo/pkg/version"
)

func (s *Server) versionHandler(w http.ResponseWriter, r *http.Request) {
	result := map[string]string{
		"version": version.VERSION,
		"commit":  version.REVISION,
	}
	s.JSONResponse(w, r, result)
}