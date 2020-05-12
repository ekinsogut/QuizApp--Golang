package models

type Users struct {
	Id      int
	Name    string
	Surname string
}

type Conf struct {
	ID           int
	Question     string
	TrueAnswer   string
	FalseAnswer1 string
	FalseAnswer2 string
}

type LoadConf struct {
	ID           int
	Question     string
	TrueAnswer   string
	FalseAnswer1 string
	FalseAnswer2 string
}
