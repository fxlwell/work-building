package resource

import (
	"sync"
	"time"
)

type Config struct {
	// Ensures atomic writes; protects the following fields
	mu *sync.Mutex
	// The content of config
	Content map[string]interface{}
	// Functions to read and update configuration
	FRead func() (interface{}, error)
	// Update Content Period
	// The cycle of calling the function FRead
	Period time.Duration
	// Change the configured logging pipeline
	revise chan bool
	// Is it updated in real time. IsRealtime=true or IsRealtime=false
	IsRealtime bool
	//Support realtime to update config
	FReadRealtime <-chan bool
}

//Return to a basic configuration
func NewConfig() *Config {
	return &Resource{
		mu:      new(sync.Mutex),
		Content: make(map[string]interface{}),
		FRead:   nil,
		Period:  0,
	}
}

// Configuration upgrade
// It will only be executed if the period is set and the read configuration function is not nil
func (c *Config) Upgrade() <-chan bool {
	if c.IsRealtime && c.FReadRealtime != nil {
		return c.FReadRealtime
	}

	if c.Period > 0 && c.FRead != nil {
		r.ch <- false
		return r.ch
	} else {
		go func() {
			time.Sleep(c.Period)
			if context, err := c.FRead; err != nil {
				log.Prinf("")
			} else {
				c.content = context
				r.ch <- true
			}
		}()
	}

	return c.ch
}

func (c *Config) GetContext() map[string]interface{} {
	return c.Content
}
