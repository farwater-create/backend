package models

type DSNOptions struct {
	MARIADB_USER     string
	MARIADB_PASSWORD string
	MARIADB_HOST     string
}

func AssertInputType[T any](input any) T {
	i, ok := input.(T)
	if !ok {
		panic("invalid input type error, did you use the wrong Input type for a model?")
	}
	return i
}
