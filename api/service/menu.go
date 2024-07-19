package service

import (
	"errors"

	"github.com/xbmlz/starter-gin/api/model"
	"github.com/xbmlz/starter-gin/internal/db"
)

type MenuService struct {
}

func (s *MenuService) Create(menu *model.Menu) error {
	return db.Get().Create(menu).Error
}

func (s *MenuService) Update(menu *model.Menu) error {
	if menu.ID == menu.ParentID {
		return errors.New("parent id cannot be the same as menu id")
	}
	return db.Get().Save(menu).Error
}

func (s *MenuService) List() ([]model.Menu, error) {
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
