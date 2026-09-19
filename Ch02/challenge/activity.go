package main

import (
	"encoding/json"
	"fmt"
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

Solution:
- Instructor used a LimitReader to limit the data read before continuing (which he mentioned previously, and I like the idea)
- He also included an empty Description check
- No enforced max user field size
*/

const maxLenUser = 16
const maxSize = 10 * 1024

type Activity struct {
	User        string    `json:"user"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Description string    `json:"description"`
}

func (a *Activity) Validate() error {
	//We could check that all fields are present, and populated.
	if l := len(a.User); l <= 0 || l > maxLenUser {
		return fmt.Errorf("invalid user: field must contain between 0 to %d characters, got %d", maxLenUser, l)
	}

	//The decoder appears to do _some_ validation when it can't map a string to the target type
	//	For instance, it attempts to map the string "not a real time" to "2006-01-02T15:04:05Z07:00", but errors out with a "cannot parse ..." message
	if a.EndTime.Before(a.StartTime) {
		return fmt.Errorf("invalid end_time: end_time must occur after start_time")
	}

	if l := len(a.Description); l == 0 {
		return fmt.Errorf("invalid description: field is mandatory")
	}

	//All checks passed
	return nil
}

func processActivity(r io.Reader) error {
	var act Activity

	r = io.LimitReader(r, maxSize)
	dec := json.NewDecoder(r)
	if err := dec.Decode(&act); err != nil {
		return err
	}

	err := act.Validate()

	if err != nil {
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
