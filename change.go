package resource

import (
	"sync"
	"time"
)

const (
	TYPE_NO = "1"
	TYPE_ED = "2"
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
}
