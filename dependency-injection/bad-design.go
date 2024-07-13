// https://www.youtube.com/watch?v=UX4XjxWcDB4
package main

// Ice
// Sandy rocks
// Concrete
type RockClimber struct {
	rocksClimbed int
	sp           SafetyPlacer // rock climber depends on implementation of safetyplacer
}

type SafetyPlacer struct {
	kind int
}

func (sp SafetyPlacer) placeSafeties() {
	switch sp.kind {
	case 1:
		// ICE
	case 2:
		// sand
	case 3:
		// concrete
	}
}

func (rc *RockClimber) climbRock() {
	rc.rocksClimbed++

	if rc.rocksClimbed == 10 {
		rc.sp.placeSafeties()
	}
}

func main() {
	rc := &RockClimber{}

	for i := 0; i < 11; i++ {
		rc.climbRock()
	}
}
