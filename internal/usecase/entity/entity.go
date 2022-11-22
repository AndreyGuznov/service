package entity

type Actor struct {
	Name string `reindex:"actor_name"`
}

type ComplexItem struct {
	ID   int64  `reindex:"id,,pk"`
	Name string `reindex:"name"`
	Year int64  `reindex:"year"`
	Sort int64  `reindex:"sort,tree"`
}

type Carbase struct {
	Id      int64  `reindex:"id,,pk"`
	CityId  int64  `reindex:"cityId"`
	Name    string `reindex:"name"`
	Vehicle string `reindex:"vehicle"`
}

type City struct {
	Id       int64      `reindex:"id,,pk"`
	Name     string     `reindex:"name"`
	Contry   string     `reindex:"country"`
	Сarbases []*Carbase `reindex:"carbases,,joined"`
	// Sort     int64      `reindex:"sort,tree"`
}
