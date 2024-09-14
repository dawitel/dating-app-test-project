package service

import (
	"sort"
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
func (s *MatchmakingService) GetMatchesForUser(user domain.User, limit, offset int) ([]domain.User, error) {
	// Fetch matches using the repository method
	matches, err := s.repo.GetMatchesForUser(user, limit, offset)
	if err != nil {
		return nil, err
	}

	// Rank matches based on mutual interests and activity status
	rankedMatches := s.rankByMutualInterests(user, matches)

	return rankedMatches, nil
}

// rankByMutualInterests ranks users based on mutual interests and last activity
func (s *MatchmakingService) rankByMutualInterests(user domain.User, candidates []domain.User) []domain.User {
	ranked := make([]domain.User, len(candidates))

	for i, candidate := range candidates {
		commonInterests := len(intersect(user.Interests, candidate.Interests))
		candidate.Score = commonInterests // Set score based on mutual interests
		ranked[i] = candidate
	}

	// Sort candidates by score (higher is better) and last activity date
	sort.SliceStable(ranked, func(i, j int) bool {
		if ranked[i].Score != ranked[j].Score {
			return ranked[i].Score > ranked[j].Score
		}
		// Use LastActive as time.Time
		return ranked[i].LastActive.After(ranked[j].LastActive)
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
