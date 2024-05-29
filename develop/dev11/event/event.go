package event

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Event struct {
	user_id    uint8
	event_name string
	data       time.Time
}

type Data struct {
	data map[int]Event
}

func NewData() *Data {
	return &Data{data: make(map[int]Event)}
}

func (d *Data) Decode(r *http.Request, evn *Event) error {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &evn)

	if err != nil {
		return err
	}
	return nil
}
