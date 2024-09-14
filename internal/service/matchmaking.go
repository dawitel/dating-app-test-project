package service

import (
	"test-matchmaking-app/internal/domain"
	"test-matchmaking-app/internal/repository"
)

type MatchmakingService struct {
	repo *repository.UserRepository
}

func NewMatchmakingService(repo *repository.UserRepository) *MatchmakingService {
	return &MatchmakingService{repo: repo}
}

// GetMatchesForUser retrieves matches for a user based on preferences filtering,
// mutual interests, and activity status. Results are paginated.
func (s *MatchmakingService) GetMatchesForUser(user domain.User, limit, offset int) ([]domain.User, int, error) {
    // Fetch matches using the repository method and return total count
    matches, totalMatches, err := s.repo.GetMatchesForUser(user, limit, offset)
    if err != nil {
        return nil, 0, err
    }

    return matches, totalMatches, nil
}


