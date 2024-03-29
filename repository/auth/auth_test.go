package auth

import (
	config "Back-End-Ecommers/configs"
	"Back-End-Ecommers/entities"
	"Back-End-Ecommers/repository/user"
	"Back-End-Ecommers/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	config := config.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run login", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{})
		db.AutoMigrate(&entities.User{})

		mockUser := entities.User{Name: "test", Email: "test", Password: "test", Status: "admin"}
		if _, err := user.New(db).Register(mockUser); err != nil {
			t.Fail()
		}

		mockLogin := entities.User{Email: "test", Password: "test"}
		res, err := repo.Login(mockLogin.Email, mockLogin.Password)

		assert.Nil(t, err)
		assert.Equal(t, "test", res.Email)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("fail run login", func(t *testing.T) {
		mockLogin := entities.User{Email: "anonim@123", Password: "anonim123"}
		_, err := repo.Login(mockLogin.Email, mockLogin.Password)
		assert.NotNil(t, err)
	})

}
