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
		minX = minimum(minX, v[0].X, v[1].X, v[2].X)
		minY = minimum(minY, v[0].Y, v[1].Y, v[2].Y)
		minZ = minimum(minZ, v[0].Z, v[1].Z, v[2].Z)
		maxX = maximum(maxX, v[0].X, v[1].X, v[2].X)
		maxY = maximum(maxY, v[0].Y, v[1].Y, v[2].Y)
		maxZ = maximum(maxZ, v[0].Z, v[1].Z, v[2].Z)
	}
	fmt.Printf("min-corner (%f,%f,%f)\n", minX, minY, minZ)
	fmt.Printf("max-corner (%f,%f,%f)\n", maxX, maxY, maxZ)
}

func minimum(nums ...float32) float32 {
	count := len(nums)
	if count == 0 {
		return 0
	}
	m := nums[0]
	for i, n := range nums {
		if i == count-1 {
			break
		}
		if n < m {
			m = n
		}
	}
	return m
}

func maximum(nums ...float32) float32 {
	count := len(nums)
	if count == 0 {
		return 0
	}
	m := nums[0]
	for i, n := range nums {
		if i == count-1 {
			break
		}
		if n > m {
			m = n
		}
	}
	return m
}
