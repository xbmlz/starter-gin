package service

import (
	"github.com/xbmlz/starter-gin/api/model"
	"github.com/xbmlz/starter-gin/internal/db"
)

type MenuService struct {
}

func (s *MenuService) CreateMenu(menu *model.Menu) error {
	return db.Get().Create(menu).Error
}

func (s *MenuService) GetMenus() ([]model.Menu, error) {
	var menus []model.Menu
	err := db.Get().Find(&menus).Error
	if err != nil {
		return nil, err
	}

	tree := buildMenuTree(menus, 0)
	return tree, nil
}

func buildMenuTree(menus []model.Menu, pid int) []model.Menu {
	tree := make([]model.Menu, 0)
	for _, menu := range menus {
		if menu.ParentID == pid {
			menu.Children = buildMenuTree(menus, menu.ID)
			tree = append(tree, menu)
		}
	}
	return tree
}
