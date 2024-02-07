package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Glowworm represents a single glowworm agent
type Glowworm struct {
	Position      []float64 // Position of the glowworm in the solution space
	Luminosity    float64   // Luminosity of the glowworm
	NeighborCount int       // Number of glowworms within the perception radius
}

// Initialize initializes a glowworm with random position and luminosity
func (g *Glowworm) Initialize(dimensions int, minPosition, maxPosition, minLuminosity, maxLuminosity float64) {
	g.Position = make([]float64, dimensions)
	for i := range g.Position {
		g.Position[i] = rand.Float64()*(maxPosition-minPosition) + minPosition
	}
	g.Luminosity = rand.Float64()*(maxLuminosity-minLuminosity) + minLuminosity
}

// Distance calculates the Euclidean distance between two glowworms
func (g *Glowworm) Distance(other *Glowworm) float64 {
	sum := 0.0
	for i := range g.Position {
		diff := g.Position[i] - other.Position[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

// UpdateNeighborCount updates the number of neighbors within the perception radius
func (g *Glowworm) UpdateNeighborCount(glowworms []*Glowworm, perceptionRadius float64) {
	count := 0
	for _, other := range glowworms {
		if g != other && g.Distance(other) < perceptionRadius {
			count++
		}
	}
	g.NeighborCount = count
}

// GlowwormSwarm represents a swarm of glowworm agents
type GlowwormSwarm struct {
	Glowworms          []*Glowworm
	PerceptionRadius   float64 // Perception radius for glowworms
	AttractionFactor   float64 // Attraction factor for glowworms
	RandomMotionFactor float64 // Random motion factor for glowworms
}

// Initialize initializes the glowworm swarm with random glowworms
func (s *GlowwormSwarm) Initialize(swarmSize, dimensions int, minPosition, maxPosition, minLuminosity, maxLuminosity float64) {
	s.Glowworms = make([]*Glowworm, swarmSize)
	for i := range s.Glowworms {
		s.Glowworms[i] = &Glowworm{}
		s.Glowworms[i].Initialize(dimensions, minPosition, maxPosition, minLuminosity, maxLuminosity)
	}
}

// Update updates the glowworm swarm for one iteration
func (s *GlowwormSwarm) Update() {
	for _, glowworm := range s.Glowworms {
		// Update neighbor count for each glowworm
		glowworm.UpdateNeighborCount(s.Glowworms, s.PerceptionRadius)

		// Update luminosity
		glowworm.Luminosity += s.AttractionFactor * float64(glowworm.NeighborCount)

		// Perform random motion
		for i := range glowworm.Position {
			glowworm.Position[i] += rand.NormFloat64() * s.RandomMotionFactor
		}
	}
}

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Create glowworm swarm
	swarmSize := 10
	dimensions := 2
	minPosition := -10.0
	maxPosition := 10.0
	minLuminosity := 0.0
	maxLuminosity := 1.0
	perceptionRadius := 2.0
	attractionFactor := 0.1
	randomMotionFactor := 0.1

	swarm := &GlowwormSwarm{
		PerceptionRadius:   perceptionRadius,
		AttractionFactor:   attractionFactor,
		RandomMotionFactor: randomMotionFactor,
	}
	swarm.Initialize(swarmSize, dimensions, minPosition, maxPosition, minLuminosity, maxLuminosity)

	// Run glowworm swarm optimization
	iterations := 100
	for i := 0; i < iterations; i++ {
		swarm.Update()
	}

	// Print final positions and luminosities of glowworms
	fmt.Println("Final positions and luminosities:")
	for _, glowworm := range swarm.Glowworms {
		fmt.Printf("Position: %v, Luminosity: %v\n", glowworm.Position, glowworm.Luminosity)
	}
}
