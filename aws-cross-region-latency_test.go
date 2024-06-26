package aws_cross_region_latency

import (
	"fmt"
	"testing"
)

// FunctionMap stores a map from src to another map from dst to the generation function

func TestLatency(t *testing.T) {
	path := "data/AWSCrossRegionLatencyMatrixParams_240419.csv"
	LatencySimulator := GetLatencyFunctions(path)
	// PrintLatencyMatrix(LatencySimulator)

	// region0 := GetRegionFromIndex(0)
	// region1 := GetRegionFromIndex(1)
	// region2 := GetRegionFromIndex(2)
	for si, src := range regionsOrdered {
		for di, dst := range regionsOrdered {
			fmt.Printf("Latency from %s[%d] to %s[%d]\t\t%.3fms\n", src, si, dst, di, LatencySimulator[src][dst].Generate())
		}
		// fmt.Printf("index[%d], name[%s], \n", index, region)
	}
	// region0To2 := LatencySimulator[region0][region2]
	// t.Log()
	// fmt.Printf("Simulated Latency from %s\tto\t%s: %.3fms\n", region0, region1, LatencySimulator[region0][region1].Generate())
	// fmt.Printf("Simulated Latency from %s\tto\t%s: %.3fms\n", region2, region1, LatencySimulator[region2][region1].Generate())
	// fmt.Printf("Simulated Latency from Tokyo to Seoul: %.3f", LatencySimulator["Tokyo"]["Seoul"].Generate())
}
