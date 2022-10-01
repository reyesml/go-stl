package main

import (
	"fmt"
	"math"
	"os"

	"neilpa.me/go-stl"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s FILE\n", os.Args[0])
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", os.Args[0], err)
		os.Exit(1)
	}
	mesh, err := stl.Decode(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", os.Args[0], err)
		os.Exit(1)
	}

	var (
		minX float32 = math.MaxFloat32
		minY float32 = math.MaxFloat32
		minZ float32 = math.MaxFloat32
		maxX float32 = -math.MaxFloat32
		maxY float32 = -math.MaxFloat32
		maxZ float32 = -math.MaxFloat32
	)
	for _, face := range mesh.Faces {
		v := face.Verts
		minX = min(min(minX, v[0].X), min(v[1].X, v[2].X))
		minY = min(min(minY, v[0].Y), min(v[1].Y, v[2].Y))
		minZ = min(min(minZ, v[0].Z), min(v[1].Z, v[2].Z))
		maxX = max(max(maxX, v[0].X), max(v[1].X, v[2].X))
		maxY = max(max(maxY, v[0].Y), max(v[1].Y, v[2].Y))
		maxZ = max(max(maxZ, v[0].Z), max(v[1].Z, v[2].Z))
	}
	fmt.Printf("min-corner (%f,%f,%f)\n", minX, minY, minZ)
	fmt.Printf("max-corner (%f,%f,%f)\n", maxX, maxY, maxZ)
}

func min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
