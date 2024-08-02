package model

import "github.com/xbmlz/starter-gin/internal/db"

type Menu struct {
	ID       int    `json:"id"`
	Type     int    `json:"type"`
	Title    string `json:"title"`
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

func GetMenuTree() (menus []Menu, err error) {
	err = db.Get().Find(&menus).Error
	if err != nil {
		return nil, err
	}

	tree := buildMenuTree(menus, 0)
	return tree, nil
}

func CreateMenu(menu *Menu) error {
	return db.Get().Create(menu).Error
}

func buildMenuTree(menus []Menu, pid int) []Menu {
	tree := make([]Menu, 0)
	for _, menu := range menus {
		if menu.ParentID == pid {
			menu.Children = buildMenuTree(menus, menu.ID)
			tree = append(tree, menu)
		}
	}
	return tree
}
