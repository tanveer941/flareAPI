package database

import "flareAPI/types"

type Storer interface {
	GetAdmin(id int) *types.User
}
