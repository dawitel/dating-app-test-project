package repository

import (
	"errors"
	"sort"
	"test-matchmaking-app/internal/domain"

	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// UserRepository represets the storage for the users data.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository returns a pointer to the storage for the users data.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetUserByID retrives the user by the user id from the db.
func (r *UserRepository) GetUserByID(userID string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("user_id = ?", userID).First(&user).Error
	return user, err
}

// GetMatchesForUser retrieves matches based on user preferences, location, mutual interests, and activity status.
func (r *UserRepository) GetMatchesForUser(user domain.User, limit, offset int) ([]domain.User, error) {
	var matches []domain.User

	// Filter based on gender preference, age range, location (distance), and mutual interests
	err := r.db.
		Where("gender = ?", user.Preferences.Gender).
		Where("age BETWEEN ? AND ?", user.Preferences.MinAge, user.Preferences.MaxAge).
		Where("ST_DistanceSphere(ST_MakePoint(longitude, latitude), ST_MakePoint(?, ?)) <= ?",
			user.Location.Longitude, user.Location.Latitude, user.Preferences.MaxDistance).
		Where("array_length(array(select unnest(interests) intersect select unnest(?)), 1) > 0", pq.Array(user.Interests)). // Filter by mutual interests
		Limit(limit).
		Offset(offset).
		Order(clause.OrderByColumn{Column: clause.Column{Name: "last_active"}, Desc: true}). // Prioritize recent activity
		Find(&matches).Error

	if err != nil {
		return nil, err
	}

	return matches, nil
}


// rankByMutualInterests ranks users based on the number of common interests shared with the current user
func (r *UserRepository) rankByMutualInterests(user domain.User, candidates []domain.User) []domain.User {
	ranked := make([]domain.User, len(candidates))

	for i, candidate := range candidates {
		commonInterests := len(intersect(user.Interests, candidate.Interests))
		candidate.Score = commonInterests
		ranked[i] = candidate
	}

	// Sort only by score (higher is better), as users are already sorted by last_active from DB
	sort.SliceStable(ranked, func(i, j int) bool {
		return ranked[i].Score > ranked[j].Score
	})

	return ranked
}


// intersect calculates the common elements between two slices
func intersect(a, b []string) []string {
	m := make(map[string]bool)
	for _, item := range a {
		m[item] = true
	}

	var result []string
	for _, item := range b {
		if m[item] {
			result = append(result, item)
		}
	}
	return result
}

// DeleteUser deletes a user by their ID
func (r *UserRepository) DeleteUser(userID string) error {
	result := r.db.Delete(&domain.User{}, "user_id = ?", userID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

func (repo *UserRepository) DeleteUserByID(userID string) error {
	if err := repo.db.Delete(&domain.User{}, "user_id = ?", userID).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByName retrieves a user by their name.
func (r *UserRepository) GetUserByName(name string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("name = ?", name).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
