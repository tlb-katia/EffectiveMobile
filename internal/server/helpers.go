package server

import (
	"EffectiveMobile_Project/internal/server/lib"
	"net/http"
)

func (s *Server) GetLimitAndOffset(r *http.Request) (uint64, uint64) {
	offsetStr := r.URL.Query().Get("offset")
	offset := lib.ConvertStrIntoInt(offsetStr)

	limitStr := r.URL.Query().Get("limit")
	limit := lib.ConvertStrIntoInt(limitStr)
	if limit == 0 {
		limit = 5
	}

	return offset, limit
}
