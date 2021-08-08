package ecs_test

import (
	"testing"

	"github.com/CarrotzRule123/ecs"
	"github.com/andygeiss/assert"
)

func TestRun(t *testing.T) {
	em := ecs.NewEntityManager()
	component := &mockComponent{name: "foo", mask: 1}
	entity := ecs.NewEntity("foo", []ecs.Component{component})
	em.Add(entity)
	sm := ecs.NewSystemManager()
	sm.Add(&mockupChangeOneEntitySystem{})
	ecs.Run(em, sm, 30)
	assert.That("run should change name to bar", t, component.name, "bar")
}

func TestRunAsMain(t *testing.T) {
	em := ecs.NewEntityManager()
	component := &mockComponent{name: "foo", mask: 1}
	entity := ecs.NewEntity("foo", []ecs.Component{component})
	em.Add(entity)
	sm := ecs.NewSystemManager()
	sm.Add(&mockupChangeOneEntitySystem{})
	ecs.Main(func() {
		ecs.Do(func() {
			ecs.Run(em, sm, 30)
		})
	})
	assert.That("run should change name to bar", t, component.name, "bar")
}

type mockupChangeOneEntitySystem struct{}

func (s *mockupChangeOneEntitySystem) Process(em *ecs.EntityManager, dt float64) (state int) {
	e := em.Get("foo")
	e.Get(1).(*mockComponent).name = "bar"
	return ecs.StateEngineStop
}
func (s *mockupChangeOneEntitySystem) Setup(entityManager *ecs.EntityManager) {}
func (s *mockupChangeOneEntitySystem) Teardown()                              {}
