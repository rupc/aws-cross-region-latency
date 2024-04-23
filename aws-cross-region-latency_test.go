package aws_cross_region_latency

import (
	"testing"
)

// FunctionMap stores a map from src to another map from dst to the generation function

func TestLatency(t *testing.T) {
	path := "data/AWSCrossRegionLatencyMatrixParams_240419.csv"
	LatencySimulator := GetLatencyFunctions(path)
	// PrintLatencyMatrix(LatencySimulator)

	region0 := GetRegionFromIndex(0)
	region1 := GetRegionFromIndex(1)
	region2 := GetRegionFromIndex(2)

	// region0To2 := LatencySimulator[region0][region2]
	// t.Log()
	t.Logf("Simulated Latency from %s\tto\t%s: %.3f\n", region0, region1, LatencySimulator[region0][region1].Generate())
	t.Logf("Simulated Latency from %s\tto\t%s: %.3f\n", region2, region1, LatencySimulator[region2][region1].Generate())
	// fmt.Printf("Simulated Latency from Tokyo to Seoul: %.3f", LatencySimulator["Tokyo"]["Seoul"].Generate())
}
