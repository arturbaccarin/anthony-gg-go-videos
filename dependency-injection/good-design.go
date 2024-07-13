// https://www.youtube.com/watch?v=UX4XjxWcDB4
package main

import "fmt"

type SafetyPlacer interface {
	placeSafeties()
}

// Ice
// Sandy rocks
// Concrete
type RockClimber struct {
	rocksClimbed int
	sp           SafetyPlacer // rock climber depends on the behaviour
}

func newRockClimber(sp SafetyPlacer) *RockClimber { // constructor
	return &RockClimber{
		sp: sp,
	}
}

type IceSafetyPlacer struct {
	// db
	// data
	// api
	// ...
}

func (IceSafetyPlacer) placeSafeties() {
	fmt.Println("placing my ICE safeties")
}

type NOPSafetyPlacer struct{}

func (NOPSafetyPlacer) placeSafeties() {
	fmt.Println("placing NO safeties")
}

func (rc *RockClimber) climbRock() {
	rc.rocksClimbed++

	if rc.rocksClimbed == 10 {
		rc.sp.placeSafeties()
	}
}

func main() {
	rc := newRockClimber(IceSafetyPlacer{}) // injecting a dependency

	for i := 0; i < 11; i++ {
		rc.climbRock()
	}
}
