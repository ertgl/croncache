package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Duration struct {
	time.Duration
}

func (d *Duration) MarshalJSON() (b []byte, err error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Duration.String())), nil
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var err error = nil
	if b[0] == '"' {
		inner := string(b[1 : len(b)-1])
		d.Duration, err = time.ParseDuration(inner)
		return err
	}
	var id int64
	id, err = json.Number(string(b)).Int64()
	d.Duration = time.Duration(id)
	return err
}
