package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/lentscode/booking-server/internals/models"
	"github.com/lentscode/booking-server/internals/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) SignUp(ctx context.Context, user *models.User) (string, error) {
	userExistsErr := s.checkIfUserExists(ctx, user.Email)
	if userExistsErr != nil {
		return "", userExistsErr
	}

	hashedPassword, err := s.hashPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword

	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	sessionId := s.generateSessionId()
	if sessionId == "" {
		return "", errors.New("failed to generate session id")
	}

	session := &models.UserSession{
		SessionID: sessionId,
		UserID:    user.ID,
	}

	if err = s.userRepo.CreateSession(ctx, session); err != nil {
		return "", err
	}

	return sessionId, nil
}

func (s *UserService) Login(ctx context.Context, user *models.User) (string, error) {
	dbUser, err := s.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}

	if err := s.validatePassword(user.Password, dbUser.Password); err != nil {
		return "", err
	}

	sessionId := s.generateSessionId()
	if sessionId == "" {
		return "", errors.New("failed to generate session id")
	}

	session := &models.UserSession{
		SessionID: sessionId,
		UserID:    dbUser.ID,
	}

	if err = s.userRepo.CreateSession(ctx, session); err != nil {
		return "", err
	}
	

	return sessionId, nil
}

func (s *UserService) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (s *UserService) validatePassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *UserService) checkIfUserExists(ctx context.Context, email string) error {
	_, err := s.userRepo.GetUserByEmail(ctx, email)
	if err == nil {
		return errors.New("user already exists")
	}

	return nil
}

func (s *UserService) generateSessionId() string {
	return uuid.New().String()
}
