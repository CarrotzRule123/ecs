package ecs

import "time"

const (
	StateEngineContinue = 0
	StateEngineStop     = 1
	Version             = "v0.0.68"
)

// engine is simple a composition of an EntityManager and a SystemManager.
// It handles the stages Setup(), Run() and Teardown() for all the systems.
type engine struct {
	entityManager *EntityManager
	systemManager *SystemManager
	lastUpdate    time.Time
}

// NewEngine creates a new Engine and returns its address.
func NewEngine(entityManager *EntityManager, systemManager *SystemManager) *engine {
	return &engine{
		entityManager: entityManager,
		systemManager: systemManager,
	}
}

// Run calls the Process() method for each System
// until ShouldEngineStop is set to true.
func (e *engine) Run(tick int) {
	shouldStop := false
	ticker := time.NewTicker(time.Second / time.Duration(tick))

	for range ticker.C {
		if shouldStop { break }

		timeStamp := time.Now()
		dt := float32(timeStamp.Sub(e.lastUpdate) / time.Second)
		for _, system := range e.systemManager.Systems() {
			state := system.Process(e.entityManager, dt)
			if state == StateEngineStop {
				shouldStop = true
				break
			}
		}
	}
}

// Setup calls the Setup() method for each System
// and initializes ShouldEngineStop and ShouldEnginePause with false.
func (e *engine) Setup() {
	for _, sys := range e.systemManager.Systems() {
		sys.Setup()
	}
}

// Teardown calls the Teardown() method for each System.
func (e *engine) Teardown() {
	for _, sys := range e.systemManager.Systems() {
		sys.Teardown()
	}
}
