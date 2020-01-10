package env

// IndexYaml -
type IndexYaml struct {
	Mongo []Mongo `yaml:"mongo"`
}

// Mongo -
type Mongo struct {
	Name       string       `yaml:"name"`
	Host       string       `yaml:"host"`
	User       string       `yaml:"user"`
	Password   string       `yaml:"password"`
	Database   string       `yaml:"database"`
	SSL        bool         `yaml:"ssl"`
	Collection []Collection `yaml:"collection"`
}

// Collection -
type Collection struct {
	Name   string  `yaml:"name"`
	Indexs []Index `yaml:"indexs"`
}

// Index -
type Index struct {
	Name        string   `yaml:"name"`
	Keys        []string `yaml:"keys"`
	Unique      bool     `yaml:"unique,omitempty"`
	Backgroud   bool     `yaml:"backgroud,omitempty"`
	Sparse      bool     `yaml:"sparse,omitempty"`
	ExpireAfter int64    `yaml:"expireAfter"`
}
