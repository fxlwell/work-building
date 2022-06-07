package resource

import "sync"

type Resource struct {
	// Ensures atomic writes; protects the following fields
	mu *sync.Mutex
	// The Config
	Config Config

	//
	Initialization func() (interface{}, error)

	//Nodes
	Nodes map[string]interface{}
}

func NewResource() *Resource {
	return &Resource{}
}

func (r *Resource) Start() {

	IsRevise := <-r.Config.Upgrade()
	if !IsRevise {
		return
	}

	context := r.Config.GetContext()

	if w, err := r.Initialization(); err != nil {

	} else {
		r.mu.Lock()
		r.Warpper = w
		r.mu.Unlock()
	}
	r.Start()
}

func (r *Resource) Get(sn string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.Nodes[sn]
}
