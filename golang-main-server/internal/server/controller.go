package server

import "github.com/m4yb3/neural-project-server/internal/database"

type Clip struct {
	Second   int                  `json:"second"`
	Emotions database.ServerResponse `json:"emotions"`
}
