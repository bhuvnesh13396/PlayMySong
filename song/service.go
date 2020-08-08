package song

import(
	"context"
)

var (
	errInvalidArgument = err.New(101, "Invalid Arguments.")
)

type Service interface {
	Get()
	Add()
	List()
}