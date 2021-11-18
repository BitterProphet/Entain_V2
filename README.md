~~1. Add another filter to the existing RPC, so we can call `ListRaces` asking for races that are visible only.~~
~~2. We'd like to see the races returned, ordered by their `advertised_start_time`~~
~~3. Our races require a new `status` field that is derived based on their `advertised_start_time`'s. The status is simply, `OPEN` or `CLOSED`. All races that have an `advertised_start_time` in the past should reflect `CLOSED`.~~
~~4. Introduce a new RPC, that allows us to fetch a single race by its ID.~~

5. Create a `sports` service that for sake of simplicity, implements a similar API to racing. This sports API can be called `ListEvents`. We'll leave it up to you to determine what you might think a sports event is made up off, but it should at minimum have an `id`, a `name` and an `advertised_start_time`.
   1. Added service with fields as follows
      1. ID 
         1. (just like racing, increments and is the primary key for the table)
      2. Name 
         1. (game: + team 1 vs. team 2 @ advertised start time) eg Tennis: Arizona sorcerors vs. Montana whales @ 2021-11-20 07:15:09 +1000 AEST
      3. Game 
         1. type of sport eg soccer
      4. Team_1
         1. name of first team playing in the sport event
      5. Team_2
         1. name of second team playing in the sport event
      6. Advertised Start Time
         1. just like racing, a time when the event is scheduled to begin.
   2. test yourself with the following post command
      1. **curl -X "POST" "http://localhost:8000/v1/getListOfSports" -H 'Content-Type: application/json'**
