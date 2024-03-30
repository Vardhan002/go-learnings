package main

import (
	"fmt"
	"math"
)

type Body struct {
	x, y, z    float64
	vx, vy, vz float64
	mass       float64
}

func (b *Body) distanceTo(other *Body) float64 {
	dx := b.x - other.x
	dy := b.y - other.y
	dz := b.z - other.z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (b *Body) updateVelocity(dt float64, others []*Body) {
	for _, other := range others {
		if b != other {
			distance := b.distanceTo(other)
			force := other.mass / math.Pow(distance, 2)
			dx := other.x - b.x
			dy := other.y - b.y
			dz := other.z - b.z
			b.vx += force * dx / distance * dt
			b.vy += force * dy / distance * dt
			b.vz += force * dz / distance * dt
		}
	}
}

func (b *Body) updatePosition(dt float64) {
	b.x += b.vx * dt
	b.y += b.vy * dt
	b.z += b.vz * dt
}

func simulate(bodies []*Body, steps int, dt float64) {
	for i := 0; i < steps; i++ {
		for _, body := range bodies {
			body.updateVelocity(dt, bodies)
		}
		for _, body := range bodies {
			body.updatePosition(dt)
		}
	}
}

func main() {
	bodies := []*Body{
		{0, 0, 0, 0, 0, 0, 100},
		{0, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{0, 0, 1, 0, 0, 0, 1},
	}
	simulate(bodies, 1000, 0.01)
	for _, body := range bodies {
		fmt.Println(body.x, body.y, body.z)
	}
}
