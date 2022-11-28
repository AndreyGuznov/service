package service

import (
	"fmt"
	"serv/internal/repo"
	"serv/pkg/cache"
)

func RefresfData() error {

	err := cache.Instance.ClearAll()

	if err != nil {
		return err
	}

	cities, err := repo.CacheCities()

	if err != nil {
		return err
	}

	for i, v := range cities {
		err := cache.Instance.Put(fmt.Sprintf("%d", i+1), v, 0)
		if err != nil {
			return err
		}
	}

	return nil
}
