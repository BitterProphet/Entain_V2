package db

import (
	"database/sql"
	"github.com/BitterProphet/Entain_V2/sports/proto/sports"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
	"sync"
	"time"
)

// SportsRepo provides repository access to sports.
type SportsRepo interface {
	// Init will initialise our sports repository.
	Init() error

	// List will return a list of sports.
	List() ([]*sports.Sport, error)
}

type sportsRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewSportsRepo creates a new sports repository.
func NewSportsRepo(db *sql.DB) SportsRepo {
	return &sportsRepo{db: db}
}

// Init prepares the sport repository dummy data.
func (r *sportsRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy sports.
		err = r.seed()
	})

	return err
}

func (r *sportsRepo) List() ([]*sports.Sport, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getSportQueries()[sportsList]


	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanSports(rows)
}



func (m *sportsRepo) scanSports(
	rows *sql.Rows,
) ([]*sports.Sport, error) {
	var sportList []*sports.Sport

	for rows.Next() {
		var sport sports.Sport
		var advertisedStart time.Time

		if err := rows.Scan(&sport.Id , &sport.Name , &sport.Game, &sport.Team_1 , &sport.Team_2, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		sport.AdvertisedStartTime = ts

		sportList = append(sportList, &sport)
	}

	return sportList, nil


}
