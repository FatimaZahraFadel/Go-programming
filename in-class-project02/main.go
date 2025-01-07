package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Salary    int    `json:"salary"`
	Education string `json:"education"`
}

type Statistics struct {
	AverageAge           float64
	AverageSalary        float64
	YoungestPersons      []string
	OldestPersons        []string
	HighestSalaryPersons []string
	LowestSalaryPersons  []string
	EducationCounts      map[string]int
}

func main() {
	var people []Person

	data, err := ioutil.ReadFile("./peaple.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(data, &people)
	if err != nil {
		fmt.Println(err)
		return
	}

	stats := calculateStats(people)
	writeStatsToFile(stats)

}

func calculateStats(people []Person) Statistics {
	var totalAge, totalSalary, min, max, minSalary, maxSalary int
	min, max = 200, 0
	minSalary, maxSalary = 1000000, 0
	educationCount := make(map[string]int)

	var youngest, oldest, richest, poorest []string

	for _, p := range people {
		totalAge += p.Age
		totalSalary += p.Salary
		educationCount[p.Education]++

		if p.Age < min {
			min = p.Age
			youngest = []string{p.Name}
		} else if p.Age == min {
			youngest = append(youngest, p.Name)
		}

		if p.Age > max {
			max = p.Age
			oldest = []string{p.Name}
		} else if p.Age == max {
			oldest = append(oldest, p.Name)
		}

		if p.Salary > maxSalary {
			maxSalary = p.Salary
			richest = []string{p.Name}
		} else if p.Salary == maxSalary {
			richest = append(richest, p.Name)
		}

		if p.Salary < minSalary {
			minSalary = p.Salary
			poorest = []string{p.Name}
		} else if p.Salary == minSalary {
			poorest = append(poorest, p.Name)
		}
	}

	return Statistics{
		AverageAge:           float64(totalAge) / float64(len(people)),
		AverageSalary:        float64(totalSalary) / float64(len(people)),
		YoungestPersons:      youngest,
		OldestPersons:        oldest,
		HighestSalaryPersons: richest,
		LowestSalaryPersons:  poorest,
		EducationCounts:      educationCount,
	}

}

func writeStatsToFile(stats Statistics) {
	jsonData, err := json.MarshalIndent(stats, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := ioutil.WriteFile("statistics.json", jsonData, 0644); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Statistics written to 'statistics.json'")
	}
}
