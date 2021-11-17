package service

import (
	"github.com/BitterProphet/Entain_V2/sports/db"
	"github.com/BitterProphet/Entain_V2/sports/proto/sports"
	"golang.org/x/net/context"
)

type Sports interface {
	// ListSports will return a collection of sports.
	ListSports(ctx context.Context, in *sports.ListSportsRequest) (*sports.ListSportsResponse, error)
}

// sportsService implements the Sports interface.
type sportsService struct {
	sportsRepo db.SportsRepo
}

// NewSportsService instantiates and returns a new sportsService.
func NewSportsService(sportsRepo db.SportsRepo) Sports {
	return &sportsService{sportsRepo}
}

func (s *sportsService) ListSports(ctx context.Context, in *sports.ListSportsRequest) (*sports.ListSportsResponse, error) {
	listOfSports, err := s.sportsRepo.List(in.Filter)
	if err != nil {
		return nil, err
	}

	return &sports.ListSportsResponse{Sports: listOfSports}, nil
}
