package repo

import (
	"fmt"
	"serv/internal/entity"
	"serv/pkg/cache"
	"serv/pkg/logger"
	"serv/pkg/reindex"
	"sort"
	"sync"

	"github.com/restream/reindexer"
)

const namespaceCity = "city"

func StartNewspaces(namespace string) error {

	if err := reindex.Conn().Instance.OpenNamespace(namespace, reindexer.DefaultNamespaceOptions(), entity.City{}); err != nil {
		return err
	}

	return nil
}

func GetList(lim, offset int) ([]*entity.City, error) {

	logger.Info("Trying to get All cities")

	wg := sync.WaitGroup{}

	query := reindex.Conn().Instance.
		Query(namespaceCity).
		Sort("id", false).
		Offset(offset).
		Limit(lim)

	iterator := query.Exec()

	defer iterator.Close()

	models, sl, err := iterator.FetchAllWithRank()
	if err != nil {
		return nil, err
	}

	l := len(sl)

	cities := make([]*entity.City, l)

	wg.Add(l)

	for it := range models {
		cities[it] = models[it].(*entity.City)

		go func(ind int, wgr *sync.WaitGroup) {
			defer wg.Done()
			sort.SliceStable(cities[ind].Сarbases, func(i, j int) bool {
				return cities[ind].Сarbases[i].Sort > cities[ind].Сarbases[j].Sort
			})
		}(it, &wg)

	}

	wg.Wait()

	if err := iterator.Error(); err != nil {
		return nil, err
	}

	logger.Info("Success of geting all cities")

	return cities, nil
}

func GetOne(id int64) (*entity.City, error) {
	var city *entity.City

	logger.Info("Trying to get city")

	mod := cache.Instance.Get(fmt.Sprintf("%d", id))
	if mod == nil {
		return nil, nil
	}
	city = mod.(*entity.City)
	return city, nil
}

func Create(city *entity.City) error {

	logger.Info("Trying to create city")

	_, err := reindex.Conn().Instance.
		Insert(namespaceCity,
			city,
			"id=serial()")

	if err != nil {
		return err
	}

	logger.Info("Success of creating city")

	return nil
}

func Edit(city entity.City) (int64, error) {

	logger.Info("Trying to edit city")

	res, err := reindex.Conn().Instance.Update(namespaceCity, &city)
	if err != nil {
		return 0, err
	}

	if res != 0 {
		logger.Info("Success of edit city")

	} else {
		logger.Info("No such city to edit")
	}

	return int64(res), nil
}

func Delete(id int64) (int, error) {

	logger.Info("Trying to delete city")

	num, err := reindex.Conn().Instance.
		Query(namespaceCity).
		WhereInt("id", reindexer.EQ, int(id)).
		Delete()

	if err != nil {
		return 0, err
	}

	if num != 0 {
		logger.Info("Success of delete city")
	}

	return num, nil
}

func CacheCities() ([]*entity.City, error) {

	query := reindex.Conn().Instance.
		Query(namespaceCity).
		ReqTotal()

	iterator := query.Exec()

	defer iterator.Close()

	cities := make([]*entity.City, iterator.Count())

	inc := 0
	for iterator.Next() {
		cities[inc] = iterator.Object().(*entity.City)
		inc++
	}

	if err := iterator.Error(); err != nil {
		return nil, err
	}

	return cities, nil
}
