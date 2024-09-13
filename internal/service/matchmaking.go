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

func (s *MatchmakingService) FindMatchesForUser(userID string) ([]domain.User, error) {
    user, err := s.repo.GetUserByID(userID)
    if err != nil {
        return nil, err
    }

    matches, err := s.repo.GetMatchesForUser(user)
    if err != nil {
        return nil, err
    }

    // Rank by mutual interests
    rankedMatches := s.rankMatchesByInterests(user, matches)

    return rankedMatches, nil
}

func (s *MatchmakingService) rankMatchesByInterests(user domain.User, matches []domain.User) []domain.User {
    // Basic mutual interest ranking
    for i, match := range matches {
        matchInterestCount := countMutualInterests(user.Interests, match.Interests)
        matches[i].Score = matchInterestCount // Set score based on mutual interests
    }

    // Sort matches by score
    sort.SliceStable(matches, func(i, j int) bool {
        return matches[i].Score > matches[j].Score
    })

    return matches
}

func countMutualInterests(a, b []string) int {
    set := make(map[string]bool)
    for _, interest := range a {
        set[interest] = true
    }

    mutualCount := 0
    for _, interest := range b {
        if set[interest] {
            mutualCount++
        }
    }

    return mutualCount
}
