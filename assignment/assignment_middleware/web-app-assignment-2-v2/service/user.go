package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserService interface {
	Register(user *model.User) (model.User, error)
	Login(user *model.User) (token *string, err error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepository repo.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Register(user *model.User) (model.User, error) {
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return *user, err
	}

	if dbUser.Email != "" || dbUser.ID != 0 {
		return *user, errors.New("email already exists")
	}

	user.CreatedAt = time.Now()

	newUser, err := s.userRepo.CreateUser(*user)
	if err != nil {
		return *user, err
	}

	return newUser, nil
}

func (s *userService) Login(user *model.User) (token *string, err error) {
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	// Jika pengguna tidak ditemukan dalam basis data
	if dbUser.Email == "" {
		return nil, errors.New("user not found")
	}

	// Memeriksa apakah kata sandi yang diberikan cocok dengan kata sandi yang tersimpan
	if user.Password != dbUser.Password {
		return nil, errors.New("wrong password")
	}

	// Jika autentikasi berhasil, menghasilkan token JWT
	claims := &jwt.StandardClaims{
		Subject:   dbUser.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwtToken.SignedString([]byte("secret_key")) // Ganti dengan kunci rahasia Anda
	if err != nil {
		return nil, err
	}

	return &tokenString, nil // TODO: replace this
}

func (s *userService) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	dbUserTaskCategories, err := s.userRepo.GetUserTaskCategory()
	if err != nil {
		return nil, err
	}
	return dbUserTaskCategories, nil // TODO: replace this
}
