~~1. Add another filter to the existing RPC, so we can call `ListRaces` asking for races that are visible only.~~
~~2. We'd like to see the races returned, ordered by their `advertised_start_time`~~
 
3.Our races require a new `status` field that is derived based on their `advertised_start_time`'s. The status is simply, `OPEN` or `CLOSED`. All races that have an `advertised_start_time` in the past should reflect `CLOSED`.
   1. curl -X "POST" "http://localhost:8000/v1/list-races" -H 'Content-Type: application/json' -d $'{"filter": {}}'
      1. standard request as normal, just look for new field.
      2. all are closed due to the dates have passed, please check `races.go` for logic. 

~~4. Introduce a new RPC, that allows us to fetch a single race by its ID.~~
~~5. Create a `sports` service that for sake of simplicity, implements a similar API to racing. This sports API can be called `ListEvents`. We'll leave it up to you to determine what you might think a sports event is made up off, but it should at minimum have an `id`, a `name` and an `advertised_start_time`.~~ 
