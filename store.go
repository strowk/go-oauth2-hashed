package go_oauth2_hashed

import (
	"context"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/models"
	"golang.org/x/crypto/bcrypt"
)

type Hasher interface {
	Hash(password string) (string, error)
	Verify(hashedPassword, password string) error
}

type BcryptHasher struct{}

func (b *BcryptHasher) Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (b *BcryptHasher) Verify(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

type ClientInfoWithHash struct {
	wrapped oauth2.ClientInfo

	hasher Hasher
}

func NewClientInfoWithHash(
	info oauth2.ClientInfo,
	hasher Hasher,
) *ClientInfoWithHash {
	if info == nil {
		return nil
	}
	return &ClientInfoWithHash{
		wrapped: info,
		hasher:  hasher,
	}
}

func (v *ClientInfoWithHash) VerifyPassword(pass string) bool {
	if pass == "" {
		return false
	}
	err := v.hasher.Verify(v.GetSecret(), pass)
	return err == nil
}

func (v *ClientInfoWithHash) GetID() string {
	return v.wrapped.GetID()
}

func (v *ClientInfoWithHash) GetSecret() string {
	return v.wrapped.GetSecret()
}

func (v *ClientInfoWithHash) GetDomain() string {
	return v.wrapped.GetDomain()
}

func (v *ClientInfoWithHash) GetUserID() string {
	return v.wrapped.GetUserID()
}

func (v *ClientInfoWithHash) IsPublic() bool {
	return v.wrapped.IsPublic()
}

type ClientStoreWithHash struct {
	oauth2.ClientStore
	hasher Hasher
}

func NewClientStoreWithBcrypt(store oauth2.ClientStore) *ClientStoreWithHash {
	return NewClientStoreWithHash(store, &BcryptHasher{})
}

func NewClientStoreWithHash(store oauth2.ClientStore, hasher Hasher) *ClientStoreWithHash {
	if hasher == nil {
		hasher = &BcryptHasher{}
	}
	return &ClientStoreWithHash{
		ClientStore: store,
		hasher:      hasher,
	}
}

func (w *ClientStoreWithHash) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	info, err := w.ClientStore.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	rval := NewClientInfoWithHash(info, w.hasher)
	if rval == nil {
		return nil, errors.ErrInvalidClient
	}
	return rval, nil
}

func (w *ClientStoreWithHash) HashAndCreate(
	info oauth2.ClientInfo,
	createFunc func(oauth2.ClientInfo) error,
) error {
	if info == nil {
		return errors.ErrInvalidClient
	}
	if info.GetSecret() == "" {
		return errors.ErrInvalidClient
	}

	hashed, err := w.hasher.Hash(info.GetSecret())
	if err != nil {
		return err
	}
	hashedInfo := models.Client{
		ID:     info.GetID(),
		Secret: string(hashed),
		Domain: info.GetDomain(),
		UserID: info.GetUserID(),
		Public: info.IsPublic(),
	}
	return createFunc(&hashedInfo)
}
