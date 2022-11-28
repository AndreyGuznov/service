package entity

type Automobile struct {
	Model string
	Brand string
}

type Carbase struct {
	Location    string
	Sort        int64
	Automobiles []Automobile
}

type City struct {
	Id       int64     `reindex:"id,,pk"`
	Name     string    `reindex:"name"`
	Сarbases []Carbase `reindex:"carbases"`
}
