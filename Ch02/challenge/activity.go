package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"
)

/*Challenge:
- Limit size of incoming data (User, String?)
- Validate the incoming data (e.g. StartTime > EndTime)
- activity-1.json: valid data
- activity-2.json: invalid data
	* user is undefined
	* start_time > end_time
	* description: ...long
*/

type Activity struct {
	User        string    `json:"user"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Description string    `json:"description"`
}

func processActivity(r io.Reader) error {
	var act Activity

	dec := json.NewDecoder(r)
	if err := dec.Decode(&act); err != nil {
		return err
	}

	log.Printf("activity: %#v", act)
	// TODO: Store in database

	return nil
}

func main() {
	if err := processActivity(os.Stdin); err != nil {
		log.Fatal(err)
	}
}
