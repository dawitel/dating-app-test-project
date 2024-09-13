package repository

import (
	"errors"
	"test-matchmaking-app/internal/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(userID string) (domain.User, error) {
    var user domain.User
    err := r.db.Where("user_id = ?", userID).First(&user).Error
    return user, err
}

func (r *UserRepository) GetMatchesForUser(user domain.User) ([]domain.User, error) {
    var matches []domain.User

    // Query users based on preferences filtering (age, gender, distance)
    // and mutual interests. This is simplified; for real-world, geolocation
    // queries would require PostGIS.
    err := r.db.
        Where("gender = ?", user.Preferences.Gender).
        Where("age >= ? AND age <= ?", user.Preferences.MinAge, user.Preferences.MaxAge).
        Find(&matches).Error

    return matches, err
}

func (r *UserRepository) CreateUser(user *domain.User) error {
    return r.db.Create(user).Error
}

// method to delete a user
func (r *UserRepository) DeleteUser(userID string) error {
    result := r.db.Delete(&domain.User{}, "user_id = ?", userID)
    if result.Error != nil {
        return result.Error
    }

    // Check if any row was affected (i.e., user existed)
    if result.RowsAffected == 0 {
        return errors.New("user not found")
    }

    return nil
}
func (repo *UserRepository) DeleteUserByID(userID string) error {
    // Perform the delete operation
    if err := repo.db.Delete(&domain.User{}, "user_id = ?", userID).Error; err != nil {
        return err
    }
    return nil
}