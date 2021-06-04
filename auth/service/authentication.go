package service

import (
	"errors"
	"regexp"
	"time"

	"github.com/yijia-cc/grouplive/auth/entity"
	"github.com/yijia-cc/grouplive/auth/idgen"

	"github.com/yijia-cc/grouplive/auth/db/dao"
	"github.com/yijia-cc/grouplive/auth/repo"
	"github.com/yijia-cc/grouplive/auth/security"
	"github.com/yijia-cc/grouplive/auth/tm"
	"github.com/yijia-cc/grouplive/auth/tx"
	"golang.org/x/crypto/bcrypt"
)

var usernameFormat = regexp.MustCompile(`^[0-9a-zA-Z]+$`)

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
	case dao.UserNotFound:
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

func (a Authentication) nextUniqueUserID(transaction tx.Transaction) (entity.ID, error) {
	for {
		id := a.idGenerator.NextID()
		query := repo.FindUserQuery{ID: (*string)(&id)}
		_, err := a.userRepo.FindUser(transaction, query)
		switch err.(type) {
		case nil:
			continue
		case dao.UserNotFound:
			return id, nil
		default:
			return "", err
		}
	}
}

func (a Authentication) SignIn(username string, password string) (string, error) {
	if !validateUsername(username) {
		return "", errors.New("invalid username format")
	}

	if !validatePassword(password) {
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
	case dao.UserNotFound:
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
	panic("Implement me!")
}

func validateUsername(username string) bool {
	return usernameFormat.MatchString(username)
}

func validatePassword(password string) bool {
	return len(password) >= 8
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
