package db

import (
	"time"

	"syreclabs.com/go/faker"
)

func (r *sportsRepo) seed() error {
	statement, err := r.db.Prepare(
		`CREATE TABLE IF NOT EXISTS SPORTS (id INTEGER PRIMARY KEY, 
													name text,
													game TEXT,
													team_1 TEXT,
													team_2 TEXT, 
													advertised_start_time DATETIME)`)
	if err == nil {
		_, err = statement.Exec()
	}

	var sportArr [6]string
	var team1, team2, sport string
	var eventTime time.Time

	sportArr = [6]string{"Basketball","Baseball","Netball",	"Tennis","Soccer","Volleyball"}

	for i := 1; i <= 100; i++ {
		statement, err = r.db.Prepare(
			`INSERT OR IGNORE INTO SPORTS(id, 
												name,
												game, 
												team_1, 
												team_2,
												advertised_start_time) VALUES (?,?,?,?,?,?)`)
		team1 = faker.Team().Name()
		team2 = faker.Team().Name()
		sport = sportArr[faker.RandomInt(0,5)]
		eventTime = faker.Time().Between(time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 2))

		if err == nil {
			_, err = statement.Exec(
				i,
				sport + ": " + team1 + " vs. " + team2 + " @ " + eventTime.String(),
				sport,
				team1,
				team2,
				eventTime.Format(time.RFC3339),
			)
		}
	}

	return err
}
