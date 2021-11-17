package db

import (
	"database/sql"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/BitterProphet/Entain_V2/racing/proto/racing"
)

// RacesRepo provides repository access to races.
type RacesRepo interface {
	// Init will initialise our races repository.
	Init() error

	// List will return a list of races.
	List(filter *racing.ListRacesRequestFilter, sort *racing.ListRacesRequestSorting) ([]*racing.Race, error)
	// SinglRace will return a single race, dependant on the request ID given.
	SingleRace(id int32)([]*racing.Race, error)
}

type racesRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewRacesRepo creates a new races repository.
func NewRacesRepo(db *sql.DB) RacesRepo {
	return &racesRepo{db: db}
}

// Init prepares the race repository dummy data.
func (r *racesRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy races.
		err = r.seed()
	})

	return err
}

func (r *racesRepo) List(filter *racing.ListRacesRequestFilter, sort *racing.ListRacesRequestSorting) ([]*racing.Race, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getRaceQueries()[racesList]

	query, args = r.applyFilterAndSort(query, filter, sort)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanRaces(rows)
}

func (r *racesRepo) SingleRace(id int32) ([]*racing.Race, error) {
	var (
		err   error
		query string
		args  []interface{}
	)
	//retrieve base query
	query = getRaceQueries()[racesList]

	// append WHERE clause to filter based on request.
	query = query + " WHERE ID = " + strconv.Itoa(int(id))

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanRaces(rows)
}

func (r *racesRepo) applyFilterAndSort(query string, filter *racing.ListRacesRequestFilter, sort *racing.ListRacesRequestSorting) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)


	if (filter  == nil) && (sort == nil) {
		return query, args
	}

	if filter != nil {
		if len(filter.MeetingIds) > 0 {
			clauses = append(clauses, "meeting_id IN ("+strings.Repeat("?,", len(filter.MeetingIds)-1)+"?)")

			for _, meetingID := range filter.MeetingIds {
				args = append(args, meetingID)
			}
		}

		//detect 'visible' boolean from request and filters accordingly
		if filter.Visible != nil {
			if filter.GetVisible() == true {
				clauses = append(clauses, "visible = true")
			}
			if filter.GetVisible() == false {
				clauses = append(clauses, "visible = false")
			}
		}
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	if sort != nil { //if 'sort' was found in request
		if sort.GetField() == "advertisedStartTime" {
			query = query + "ORDER BY " + "advertised_start_time"
			query = query + " " + sort.GetOrder()
		}
	}

	return query, args
}

func (m *racesRepo) scanRaces(
	rows *sql.Rows,
) ([]*racing.Race, error) {
	var races []*racing.Race

	for rows.Next() {
		var race racing.Race
		var advertisedStart time.Time

		if err := rows.Scan(&race.Id, &race.MeetingId, &race.Name, &race.Number, &race.Visible, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		race.AdvertisedStartTime = ts

		//set status of race to 'closed' when 'advertised_starting_time' is in th past
		// otherwise, set to OPEN.
		if advertisedStart.Before(time.Now()) {
			race.Status = "CLOSED"
		}else{
			race.Status = "OPEN"
		}
		
		races = append(races, &race)
	}

	return races, nil
}
