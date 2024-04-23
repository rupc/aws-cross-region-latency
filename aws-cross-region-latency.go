package aws_cross_region_latency

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"text/tabwriter"
)

var (
	// regionsOrdered is a list of AWS regions in the order
	regionsOrdered = []string{"Seoul", "Tokyo", "Hong Kong", "Osaka", "Singapore", "Sydney", "Frankfurt", "London", "N. California", "Ireland", "Mumbai", "N. Virginia", "Ohio", "Oregon", "Stockholm", "Paris", "Central", "São Paulo", "Bahrain", "Milan", "Cape Town"}
)

func GetRegionFromIndex(index int) string {
	return regionsOrdered[index]
}

type LatencySimulator struct {
	mean float64
	std  float64
}

func (l *LatencySimulator) Generate() float64 {
	return l.mean + rand.NormFloat64()*l.std
}

// GetLatencyFunctions inputs filepath aws cross region latency matrixes (e.g., AWSCrossRegionLatencyMatrixParams_240419.csv)
// returns a latency map accessible by map[src][dst]
func GetLatencyFunctions(filepath string) map[string]map[string]*LatencySimulator {
	FunctionMap, err := loadFunctions(filepath)
	if err != nil {
		panic(err)
	}
	return FunctionMap
}

// loadFunctions reads latency parameters from a CSV file and creates functions for each src to dst
func loadFunctions(filePath string) (map[string]map[string]*LatencySimulator, error) {
	FunctionMap := make(map[string]map[string]*LatencySimulator)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Unable to read input file %s: %v\n", filePath, err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Unable to parse file as CSV: %v\n", err)
		return nil, err
	}

	for i, record := range records {
		if i == 0 {
			continue // Skip header
		}
		src := record[0]
		dst := record[1]
		mean, _ := strconv.ParseFloat(record[2], 64)
		std, _ := strconv.ParseFloat(record[3], 64)

		// Ensure the src map exists
		if _, exists := FunctionMap[src]; !exists {
			FunctionMap[src] = make(map[string]*LatencySimulator)
		}

		// Create a function for each src to dst pair
		FunctionMap[src][dst] = makeLatencyFunc(mean, std)
	}
	return FunctionMap, nil
}

// makeLatencyFunc creates a function to generate latencies for specific mean and std
func makeLatencyFunc(mean, std float64) *LatencySimulator {
	return &LatencySimulator{
		mean: mean,
		std:  std,
	}
}

func PrintLatencyMatrix(FunctionMap map[string]map[string]*LatencySimulator) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(writer, "Source\tDestination\tMean(ms)\tStd\t") // Header

	for src, destinations := range FunctionMap {
		for dst, params := range destinations {
			fmt.Fprintf(writer, "%s\t%s\t%.3f\t%.3f\t\n", src, dst, params.mean, params.std)
		}
	}
	writer.Flush() // Send output to standard output
}
