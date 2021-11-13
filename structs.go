package couchdb_helper

type CouchDB struct {
	Schema      string   `yaml:"schema"`
	Server      string   `yaml:"server"`
	Port        string   `yaml:"port"`
	Username    string   `yaml:"username"`
	Password    string   `yaml:"password"`
	Name        string   `yaml:"name"`
	IndexFields []string `yaml:"indexFields"`
}
