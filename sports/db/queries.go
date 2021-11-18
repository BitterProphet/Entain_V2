package db

const (
	sportsList = "list"
)

func getSportQueries() map[string]string {
	return map[string]string{
		sportsList: `
			SELECT 
				id,
				name,
				game,
				team_1,
				team_2,
				advertised_start_time 
			FROM sports
		`,
	}
}
