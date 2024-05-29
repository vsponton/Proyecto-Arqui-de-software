package dto

type UsuarioDetailDto struct {
	Id          int    `json:"id"`
	Descripcion string `json:"usuario"`

	// SensoresMinDto SensoresMinDto `json:"sensores,omitempty"`
}

// type BarriosDetailDto []BarrioDetailDto