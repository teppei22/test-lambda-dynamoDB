package lambdaFunc

import "github.com/google/uuid"

func UUID() string {

	u, _ := uuid.NewRandom()
	return u.String()
}
