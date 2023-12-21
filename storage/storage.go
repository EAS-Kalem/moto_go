package storage

import "github.com/EAS-Kalem/moto_go/types"

type Storage interface {
	Get(int) *types.User
}