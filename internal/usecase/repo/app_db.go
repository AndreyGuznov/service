package repo

import (
	"fmt"
	"serv/internal/usecase/entity"
	"serv/pkg/cache"
	"serv/pkg/logger"
	"serv/pkg/reindex"
	"time"

	"github.com/restream/reindexer"
)

var (
	namespaceCity    = "city"
	namespaceCarbase = "carbase"
)

func StartNewspaces(namespace1, namespace2 string) error {

	err := reindex.Conn().Instance.OpenNamespace(namespace1, reindexer.DefaultNamespaceOptions(), entity.City{})
	if err != nil {
		return err
	}

	err = reindex.Conn().Instance.OpenNamespace(namespace2, reindexer.DefaultNamespaceOptions(), entity.Carbase{})
	if err != nil {
		return err
	}

	return nil
}

func CreateCity(model *entity.City) error {

	logger.Info("Trying to create city")

	_, err := reindex.Conn().Instance.
		Insert(namespaceCity,
			model,
			"id=serial()")

	if err != nil {
		return err
	}

	logger.Info("City was created")

	return nil
}

func CreateCarbase(model *entity.Carbase) error {

	logger.Info("Trying to create carbase")

	_, err := reindex.Conn().Instance.
		Insert(namespaceCarbase,
			model,
			"id=serial()")

	if err != nil {
		return err
	}

	logger.Info("Carbase was created")

	return nil
}

// if edit need to delete from cache/ update cache???
func EditCity(model *entity.City) error {

	return nil
}

func EditCarbase(model *entity.Carbase) error {

	return nil
}

func Delete(id int64) error {

	logger.Info("Trying to delete collection")

	_, err := reindex.Conn().Instance.
		Query(namespaceCity).
		WhereInt("id", reindexer.EQ, int(id)).
		Delete()

	if err != nil {
		return err
	}

	logger.Info("Collection was deleted")

	return nil
}

func GetList(lim int64) ([]*entity.City, error) {

	complexItems := make([]*entity.City, 0, lim)

	query := reindex.Conn().Instance.
		Query(namespaceCity).
		Sort("sort", false).
		Limit(int(lim)).
		Offset(3)

	iterator := query.Exec()

	defer iterator.Close()

	for iterator.Next() {
		// complexItems = append(complexItems, iterator.Object().(*entity.ComplexItem))
	}

	if err := iterator.Error(); err != nil {
		return nil, err
	}

	return complexItems, nil
}

// TODO transanction What to do with nill if carbase empty
func GetCity(id int64) (*entity.City, error) {
	var (
		city *entity.City
		// carbases []*entity.Carbase
		found bool
	)

	mod := cache.GetCahe().Get(fmt.Sprintf("%d", id))

	cacheCity, found := mod.(entity.City)
	if found {
		return &cacheCity, nil
	}

	// query := reindex.Conn().Instance.
	// 	Query(namespaceCarbase).
	// 	WhereInt64("cityId", reindexer.EQ, id)

	// iterator := query.Exec()

	// defer iterator.Close()

	// for iterator.Next() {
	// 	carbases = append(carbases, iterator.Object().(*entity.Carbase))
	// }

	// query2, found := reindex.Conn().Instance.
	// 	Query(namespaceCity).
	// 	WhereInt64("id", reindexer.EQ, id).
	// 	Get()

	query2, found := reindex.Conn().Instance.
		Query(namespaceCity).
		WhereInt64("id", reindexer.EQ, id).
		LeftJoin(
			reindex.Conn().Instance.
				Query(namespaceCarbase), "carbases").
		On("cityId", reindexer.SET, "id").
		Get()

	if found {
		city = query2.(*entity.City)
		// if len(carbases) != 0 {
		// 	city.Сarbases = carbases
		// }
		cache.GetCahe().Put(fmt.Sprintf("%d", id), city, 500*time.Second)
	}

	return city, nil
}
