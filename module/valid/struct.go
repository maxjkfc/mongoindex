package valid

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database -
type Database struct {
	Name       string
	Database   string
	Options    *options.ClientOptions
	Collection map[string][]*options.IndexOptions
}
