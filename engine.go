package ecs

import "time"


const (
	StateEngineContinue = 0
	StateEngineStop     = 1
	Version             = "v0.0.68"
)

// engine is simple a composition of an EntityManager and a SystemManager.
// It handles the stages Setup(), Run() and Teardown() for all the systems.
type Engine struct {
	entityManager *EntityManager
	systemManager *SystemManager
	lastUpdate    time.Time
}

// NewEngine creates a new Engine and returns its address.
func NewEngine(entityManager *EntityManager, systemManager *SystemManager) *Engine {
	return &Engine{
		entityManager: entityManager,
		systemManager: systemManager,
	}
}

// Run calls the Process() method for each System
// until ShouldEngineStop is set to true.
func (e *Engine) Run(tick int) {
	shouldStop := false
	ticker := time.NewTicker(time.Second / time.Duration(tick))

	for range ticker.C {
		if shouldStop {
			break
		}

		timeStamp := time.Now()
		dt := float64(timeStamp.Sub(e.lastUpdate).Milliseconds()) / 1000
		e.lastUpdate = timeStamp
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
func (e *Engine) Setup() {
	for _, sys := range e.systemManager.Systems() {
		sys.Setup(e.entityManager)
	}
}

// Teardown calls the Teardown() method for each System.
func (e *Engine) Teardown() {
	for _, sys := range e.systemManager.Systems() {
		sys.Teardown()
	}
}
