package event

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Event struct {
	User_id    uint8  `json:"user_id"`
	Event_name string `json:"event_name"`
	Event      string `json:"event"`
	Time       string `json:"data"`
}

func (ev *Event) Decode(r http.Request) error {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, ev)
	if err != nil {
		return err
	}
	return nil
}

func (ev *Event) Validate() error {

	dateString := ev.Time
	date, err := time.Parse(time.DateOnly, dateString)

	if err != nil {
		return err
	}

	if ok := date.After(time.Now()); !ok {
		return fmt.Errorf("the entered date has already passed")
	}

	return nil
}
