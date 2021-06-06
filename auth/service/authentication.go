package service

import (
	"errors"
	"time"

	"github.com/yijia-cc/grouplive/auth/validator"

	"github.com/yijia-cc/grouplive/auth/entity"
	"github.com/yijia-cc/grouplive/auth/idgen"

	"github.com/yijia-cc/grouplive/auth/db/dao"
	"github.com/yijia-cc/grouplive/auth/repo"
	"github.com/yijia-cc/grouplive/auth/security"
	"github.com/yijia-cc/grouplive/auth/tm"
	"github.com/yijia-cc/grouplive/auth/tx"
	"golang.org/x/crypto/bcrypt"
)

const hour = 60 * time.Minute
const day = 24 * hour
const week = 7 * day
const authTokenValidDuration = week

type tokenPayload struct {
	EncryptedUserID string    `json:"encrypted_user_id"`
	IssuedAt        time.Time `json:"issued_at"`
}

type Authentication struct {
	timer        tm.Timer
	jwtAuthority security.JWTAuthority
	cipher       security.CaesarCipher
	txFactory    tx.TransactionFactory
	userRepo     repo.User
	idGenerator  idgen.IDGenerator
}

func (a Authentication) SignUp(user entity.User, password string) error {
	err := user.Validate()
	if err != nil {
		return err
	}

	transaction, err := a.txFactory.NewTransaction()
	if err != nil {
		return err
	}

	query := repo.FindUserQuery{Username: &user.Username}

	// no user with username 123
	_, err = a.userRepo.FindUser(transaction, query)
	switch err.(type) {
	case nil:
		return errors.New("user already exists")
	case dao.NotFound:
	default:
		return err
	}

	user.ID, err = a.nextUniqueUserID(transaction)
	if err != nil {
		return err
	}

	encryptedPwdBuf, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.EncryptedPassword = string(encryptedPwdBuf)

	err = a.userRepo.CreateUser(transaction, user)
	if err != nil {
		return err
	}
	return transaction.Commit()
}

func (a Authentication) SignIn(username string, password string) (string, error) {
	if !validator.ValidateUsername(username) {
		return "", errors.New("invalid username format")
	}

	if !validator.ValidatePassword(password) {
		return "", errors.New("invalid password format")
	}

	transaction, err := a.txFactory.NewTransaction()
	if err != nil {
		return "", err
	}

	query := repo.FindUserQuery{Username: &username}
	user, err := a.userRepo.FindUser(transaction, query)
	switch err.(type) {
	case nil:
	case dao.NotFound:
		return "", errors.New("user not found")
	default:
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password))
	if err != nil {
		return "", errors.New("user credential not match")
	}

	payload := tokenPayload{
		EncryptedUserID: a.cipher.Encrypt(username),
		IssuedAt:        a.timer.Now(),
	}
	return a.jwtAuthority.IssueToken(payload)
}

func (a Authentication) VerifyIdentity(authToken string) (string, error) {
	payload := tokenPayload{}
	err := a.jwtAuthority.GetPayload(authToken, &payload)
	if err != nil {
		return "", err
	}

	expiredAt := payload.IssuedAt.Add(authTokenValidDuration)
	if expiredAt.Before(a.timer.Now()) {
		return "", errors.New("auth token expired")
	}

	return a.cipher.Decrypt(payload.EncryptedUserID), nil
}

func (a Authentication) nextUniqueUserID(transaction tx.Transaction) (entity.ID, error) {
	for {
		id := a.idGenerator.NextID()
		query := repo.FindUserQuery{ID: (*string)(&id)}
		_, err := a.userRepo.FindUser(transaction, query)
		switch err.(type) {
		case nil:
			continue
		case dao.NotFound:
			return id, nil
		default:
			return "", err
		}
	}
}

func NewAuthentication(timer tm.Timer, idGenerator idgen.IDGenerator, txFactory tx.TransactionFactory, userDao dao.User, jwtSigningKey string, caesarCipherOffset int) Authentication {
	return Authentication{
		timer:        timer,
		idGenerator:  idGenerator,
		txFactory:    txFactory,
		jwtAuthority: security.NewJWTAuthority(jwtSigningKey),
		cipher:       security.NewCaesarCipher(caesarCipherOffset),
		userRepo:     repo.NewUser(userDao),
	}
}
