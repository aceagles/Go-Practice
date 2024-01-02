package plantMap

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type plantList []plantInfo

type plantInfo struct {
	prefix, plantType, lookupIP, basePlant string
	plantID, numericID, downstream         int
	upstream                               []int
}

type plantMap struct {
	groupedMap, baseGroupedMap map[string][]plantInfo
	numericMap                 map[int]plantInfo
}

func convertToSlice(input string) ([]int, error) {
	// Split the string using the '#' delimiter
	parts := strings.Split(input, "#")

	// Initialize a slice to store the converted integers
	var intSlice []int

	// Convert each part to an integer and append to the slice
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		intSlice = append(intSlice, num)
	}

	return intSlice, nil
}

func LoadMap(path string) (mp plantMap, rtnError error) {
	//Initialise the maps
	mp.baseGroupedMap = make(map[string][]plantInfo)
	mp.groupedMap = make(map[string][]plantInfo)
	mp.numericMap = make(map[int]plantInfo)

	file, err := os.Open(path)
	if err != nil {
		// TODO: Error return
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Discard headers

	_, err = reader.Read()
	if err != nil {
		return
	}

	// Read the remaining rows
	for {
		// Read the record
		row, err := reader.Read()

		// Break the loop if we reach the end of the file
		if err != nil {
			return
		}

		// convert strings to ints and ignore the errors. Will give 0 which is fine
		plantID, _ := strconv.Atoi(row[2])
		numericID, _ := strconv.Atoi(row[3])
		downstream, _ := strconv.Atoi(row[6])

		// Convert to slice and ignore error, will give empty slice which is fine
		upstream, _ := convertToSlice(row[7])

		// Process the record
		info := plantInfo{
			prefix:     row[0],
			plantType:  row[1],
			plantID:    plantID,
			numericID:  numericID,
			lookupIP:   row[4],
			downstream: downstream,
			upstream:   upstream,
			basePlant:  row[8],
		}
		mp.baseGroupedMap[info.basePlant] = append(mp.baseGroupedMap[info.basePlant], info)
		mp.groupedMap[info.plantType] = append(mp.groupedMap[info.plantType], info)
		mp.numericMap[info.numericID] = info
	}
}

func (m *plantMap) GetPlantByBase(base string) plantList {
	return m.baseGroupedMap[base]
}

func (m *plantMap) GetPlantByType(t string) plantList {
	return m.groupedMap[t]
}

func (m *plantMap) GetPlantByID(t int) plantInfo {
	return m.numericMap[t]
}

func (s plantList) ByPrefix(p string) (plantInfo, error) {
	for _, v := range s {
		if v.prefix == p {
			return v, nil
		}
	}
	return plantInfo{}, fmt.Errorf("no matching prefix")
}
