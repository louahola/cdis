package auth
import (
	"github.com/louahola/cdis/repository"
	"io"
	"crypto/rand"
	"encoding/base64"
)

type SessionManager struct {
	SessionRepo repository.Repository
}

func (this *SessionManager) generateToken() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (this *SessionManager) NewSession(user User) (*Session, error) {
	session := &Session{UserId: user.Id, Token: this.generateToken()}
	err := this.SessionRepo.Save(session)
	return session,err
}

func (this *SessionManager) GetUser(token string) (*User, error) {
	session := &Session{Token: token}
	err := this.SessionRepo.Get(session)
	if err != nil {
		return nil, err
	}

	//TODO: if session isn't expired
	user := &User{Id: session.UserId}
	err = this.SessionRepo.Get(user)
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
