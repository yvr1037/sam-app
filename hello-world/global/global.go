package global

import "github.com/guregu/dynamo"

var (
	DB                         *dynamo.DB
	UserTable                  dynamo.Table
	AuthTable                  dynamo.Table
	CharacterRelationshipTable dynamo.Table
	CharacterFetchTable        dynamo.Table
)