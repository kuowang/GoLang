package model

import (
	"fmt"
	"testing"
)

func TestGetUserByName(t *testing.T) {
	fmt.Println(getUserByName("shineyork"))
}

func TestGetUserById(t *testing.T) {
	fmt.Println(getUserById(9))
}
