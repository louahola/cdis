package auth
import (
	"github.com/louahola/cdis/repository"
	"fmt"
)

type UserManager struct {
	repo repository.Repository
}

func (this *UserManager) UserExists(user *User) bool {
	err := this.repo.Get(user)
	if err != nil {
		return false
	}
	return true
}

func (this *UserManager) CreateUser(user *User) error {
	if (this.UserExists(user)) {
		return fmt.Errorf("User already exists!")
	}

	err := this.repo.Save(user)
	return err
}
