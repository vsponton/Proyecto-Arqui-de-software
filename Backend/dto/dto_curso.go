package dto

type CursoDetailDto struct {
	Id          int    `json:"id"`
	Descripcion string `json:"curso"`

	// SensoresMinDto SensoresMinDto `json:"sensores,omitempty"`
}

