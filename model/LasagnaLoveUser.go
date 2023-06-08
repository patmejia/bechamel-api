package model

import "encoding/json"

type LasagnaLoveUser struct {
	ID                 int    `json:"id"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	GivenName          string `json:"given_name"`
	MiddleOrMaidenName string `json:"middle_or_maiden_name"`
	FamilyName         string `json:"family_name"`
}

// This overrides the default marshaling of the structure to JSON, removing the password field value.
// It has the unfortunate side effect of leaving an empty string entry in the generated JSON
// which leaks out to users, but does allow the use of the same struct for GETting and POSTing user profiles.
//
// Simply using a marshalling entry of "-" in the struct above causes the marshalling in the POST calls
// to create user Profiles to not unmarshall the supplied password field.
//
// TODO: Better ideas to accomplish allowing and reading the "password" JSON field value when creating
// LasagnaLoveUser data structures, while not emitting these when marshalling, are welcomed.
func (l LasagnaLoveUser) MarshalJSON() ([]byte, error) {
	type lasagnaLoveUser LasagnaLoveUser // prevent recursion
	x := lasagnaLoveUser(l)
	x.Password = ""
	return json.Marshal(x)
}
