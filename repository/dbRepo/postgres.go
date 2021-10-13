package dbrepo

import "fmt"

func (m *postgresDbRepo) Test() bool {
	fmt.Println("postgresDbRepo.Test()")
	return true
}
