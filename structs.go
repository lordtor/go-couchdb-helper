package couchdb_helper

type CouchDB struct {
	Schema      string   `json:"schema" yaml:"schema"`
	Server      string   `json:"server" yaml:"server"`
	Port        string   `json:"port" yaml:"port"`
	Username    string   `json:"username" yaml:"username"`
	Password    string   `json:"-" yaml:"password" json:"-"`
	Name        string   `json:"name" yaml:"name"`
	IndexFields []string `json:"indexFields" yaml:"indexFields"`
}
