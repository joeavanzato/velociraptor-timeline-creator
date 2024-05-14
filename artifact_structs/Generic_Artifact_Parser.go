package artifact_structs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
	"os"
	"slices"
)

func Get_Generic_Headers(fullparse bool, filepath string) ([]string, error) {
	keys := make([]string, 0)
	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		return keys, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tmpMap, tmpErr := parseJSONtoMap(scanner.Text())
		if tmpErr != nil {
			return keys, tmpErr
		}
		fmt.Println(tmpMap)

	}

	return keys, nil
}

func extractAllHeadersFromMap(inputMap map[string]any, currentHeaders []string) []string {
	for k, v := range inputMap {
		keyExists := slices.Contains(currentHeaders, k)
		if keyExists {
			continue
		}

		switch vv := v.(type) {
		case []interface{}:
			new, ierr := v.(map[string]any)
			if !ierr {
				continue
			}
			extractAllHeadersFromMap(new, currentHeaders)
		case map[string]interface{}:
			s := make(map[string]string, len(vv))
			for kk, uu := range vv {
				s[kk] = fmt.Sprint(uu)
			}
			//tempRecord[headerIndex] = fmt.Sprint(s)
		default:
			currentHeaders = append(currentHeaders, k)
		}
	}
	return currentHeaders
}

func parseJSONtoMap(input string) (map[string]any, error) {
	var result map[string]any
	jsonErr := json.Unmarshal([]byte(input), &result)
	if jsonErr != nil {
		return make(map[string]any), jsonErr
	}
	return result, nil
}

func Process_Generic_Artifact(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		var tmp map[string]any
		jsonErr := json.Unmarshal([]byte(line), &tmp)
		if jsonErr != nil {
			logger.Error().Msgf(jsonErr.Error())
			continue
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", []string{}, outputChannel)
			continue
		}
		// We don't do 'super timeline' for 'generic' artifacts since we don't know their time fields, metadata, etc.
		// We assume that every 'row' of a particular file will have the same string of keys across all rows
		// This is not a perfect assumption for deeply nested data but will work 'good enough' for most use cases
	}
}
