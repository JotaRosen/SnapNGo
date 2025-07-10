package types

type ConnectionParams struct {
	Command  string //should be an ENUM [backup, restore, pingDB]
	Engine   string
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}
