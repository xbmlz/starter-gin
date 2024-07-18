package model

type Menu struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID int    `json:"parent_id"`
	Path     string `json:"path"`
	Icon     string `json:"icon"`
	Sort     int    `json:"sort"`
	Enabled  int    `json:"enabled"`
	Children []Menu `json:"children" gorm:"-"`
}

func (m *Menu) TableName() string {
	return "sys_menu"
}
