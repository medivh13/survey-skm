package postgres

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : auth-skm
 */

import (
	"context"

	dtoLogin "auth-skm/src/app/dtos/login"
	dto "auth-skm/src/app/dtos/users"
	repositories "auth-skm/src/domain/repositories"
	models "auth-skm/src/infra/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type usersRepository struct {
	connection *gorm.DB
}

func NewUsersRepository(db *gorm.DB) repositories.UsersRepository {
	return &usersRepository{
		connection: db,
	}
}

func (repo *usersRepository) Register(ctx context.Context, data *dto.UserReqDTO) (*models.Users, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data.PassWord), 14)
	if err != nil {
		return nil, err
	}
	password := string(bytes)
	userModel := models.Users{
		Name:  data.Name,
		Email: data.Email,
	}

	q := repo.connection.WithContext(ctx)
	tx := q.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Raw("INSERT INTO master.users (display_name, email_address, password) VALUES (?,?,?)", data.Name, data.Email, password).Scan(&userModel).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &userModel, tx.Commit().Error
}

func (repo *usersRepository) GetUser(ctx context.Context, data *dtoLogin.LoginReqDTO) (*models.Users, error) {
	var userModel models.Users
	q := repo.connection.WithContext(ctx)

	if err := q.Raw("SELECT id, password from master.users WHERE email_address = ?", data.Email).Scan(&userModel).Error; err != nil {
		return nil, err
	}

	return &userModel, nil
}

func (repo *usersRepository) UpdateToken(ctx context.Context, id int64, token string) error {
	q := repo.connection.WithContext(ctx)

	if err := q.Raw("UPDATE master.users SET token = ? WHERE id = ? RETURNING token", token, id).Scan(&token).Error; err != nil {
		return err
	}

	return nil
}
