package authorization

import "time"

type CookiePayload struct {
	Name     string        `json:"name"`
	Value    string        `json:"value"`
	HTTPOnly bool          `json:"http_only" default:"true"`
	Expires  time.Duration `json:"expires"`
}
