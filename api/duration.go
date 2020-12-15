package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gravitational/teleport"
	"github.com/gravitational/trace"
)

// Duration is a wrapper around duration to set up custom marshal/unmarshal
type Duration time.Duration

// Duration returns time.Duration from Duration typex
func (d Duration) Duration() time.Duration {
	return time.Duration(d)
}

// Value returns time.Duration value of this wrapper
func (d Duration) Value() time.Duration {
	return time.Duration(d)
}

// MarshalJSON marshals Duration to string
func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%v", d.Duration()))
}

// UnmarshalJSON marshals Duration to string
func (d *Duration) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	var stringVar string
	if err := json.Unmarshal(data, &stringVar); err != nil {
		return trace.Wrap(err)
	}
	if stringVar == teleport.DurationNever {
		*d = Duration(0)
	} else {
		out, err := time.ParseDuration(stringVar)
		if err != nil {
			return trace.BadParameter(err.Error())
		}
		*d = Duration(out)
	}
	return nil
}

// MarshalYAML marshals duration into YAML value,
// encodes it as a string in format "1m"
func (d Duration) MarshalYAML() (interface{}, error) {
	return fmt.Sprintf("%v", d.Duration()), nil
}

func (d *Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var stringVar string
	if err := unmarshal(&stringVar); err != nil {
		return trace.Wrap(err)
	}
	if stringVar == teleport.DurationNever {
		*d = Duration(0)
	} else {
		out, err := time.ParseDuration(stringVar)
		if err != nil {
			return trace.BadParameter(err.Error())
		}
		*d = Duration(out)
	}
	return nil
}

// MaxDuration returns maximum duration that is possible
func MaxDuration() Duration {
	return NewDuration(1<<63 - 1)
}

// NewDuration returns Duration struct based on time.Duration
func NewDuration(d time.Duration) Duration {
	return Duration(d)
}
