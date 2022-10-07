package user

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	CreateUser("test", "pass")
}
