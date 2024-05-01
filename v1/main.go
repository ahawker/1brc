package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type sample struct {
	count uint16
	min   float64
	max   float64
	sum   float64
}

func main() {
	var (
		mean    float64
		station string
		temp    float64
		err     error
		parts   []string
	)

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = file.Close()
	}()

	results := make(map[string]sample, 10)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts = strings.Split(scanner.Text(), ";")

		station = parts[0]

		temp, err = strconv.ParseFloat(parts[1], 32)
		if err != nil {
			panic(err)
		}

		if curr, ok := results[parts[0]]; ok {
			curr.count += 1
			curr.min = math.Min(curr.min, temp)
			curr.min = math.Max(curr.max, temp)
			curr.sum += temp
			results[station] = curr
		} else {
			results[station] = sample{
				count: 1,
				min:   temp,
				max:   temp,
				sum:   temp,
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	keys := make([]string, 0, len(results))

	for k := range results {
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys))

	length := len(results)

	var sb strings.Builder

	_, _ = sb.WriteString("{")

	for i, k := range keys {
		s := results[k]
		mean = s.sum / float64(s.count)
		_, _ = sb.WriteString(fmt.Sprintf("%s=%0.1f/%0.1f/%0.1f", k, s.min, mean, s.max))
		if i < length-1 {
			_, _ = sb.WriteString(", ")
		}
	}

	_, _ = sb.WriteString("}")

	fmt.Print(sb.String())
}
