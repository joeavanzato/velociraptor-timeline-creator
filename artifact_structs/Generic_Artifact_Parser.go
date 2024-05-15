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
		keys = extractAllHeadersFromMap(tmpMap, keys, "")
		break

	}
	return keys, nil
}

// TODO - Implement checking inside []interface{] to see if the objects are of type object and iterate down recursively through those
func extractAllHeadersFromMap(inputMap map[string]any, currentHeaders []string, baseKey string) []string {
	for k, v := range inputMap {
		if baseKey != "" {
			k = baseKey + "_" + k
		}
		keyExists := slices.Contains(currentHeaders, k)
		if keyExists {
			continue
		}

		// JSON unmarshaller uses the following rules when parsing a JSON strirng
		/*		bool, for JSON booleans
				float64, for JSON numbers
				string, for JSON strings
				[]interface{}, for JSON arrays
				map[string]interface{}, for JSON objects
				nil for JSON null*/

		// For all of these values except for map[string]interface{}, we can just store the key as normal - when it's a map[string]interface, we just recursively handle and build keys as required

		switch v.(type) {
		case map[string]interface{}:
			nv := v.(map[string]any)
			newheaders := extractAllHeadersFromMap(nv, currentHeaders, k)
			for _, v := range newheaders {
				currentHeaders = helpers.AddToSliceIfNotPresent(v, currentHeaders)
			}
		default:
			// If it's a 'normal' field, then we just add the header and parse as normal
			currentHeaders = append(currentHeaders, k)
		}
	}
	return currentHeaders
}

func buildRecordFromMap(inputMap map[string]any, keys []string, baseKey string, record []string) []string {

	for k, v := range inputMap {
		if baseKey != "" {
			k = baseKey + "_" + k
		}
		// JSON unmarshaller uses the following rules when parsing a JSON strirng
		/*		bool, for JSON booleans
				float64, for JSON numbers
				string, for JSON strings
				[]interface{}, for JSON arrays
				map[string]interface{}, for JSON objects
				nil for JSON null*/
		// For all of these values except for map[string]interface{}, we can just store the key as normal - when it's a map[string]interface, we just recursively handle and build keys as required

		switch v.(type) {
		case map[string]interface{}:
			nv := v.(map[string]any)
			record = buildRecordFromMap(nv, keys, k, record)
		default:
			// If it's a 'normal' field, then we just check key index and add to record
			valueIndex := helpers.FindIndexInSlice(k, keys)
			record[valueIndex] = fmt.Sprint(v)
		}
	}
	return record
}

func parseJSONtoMap(input string) (map[string]any, error) {
	var result map[string]any
	jsonErr := json.Unmarshal([]byte(input), &result)
	if jsonErr != nil {
		return make(map[string]any), jsonErr
	}
	return result, nil
}

func Process_Generic_Artifact(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger, keys []string) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	// For generic artifact, we build a slice of len(keys), then we walk the full record - if the current key exists, we find it's index and insert into the current record
	// -3 to account for the initial 3 headers we add to all artifact dumps
	keys = keys[3:]
	for _, line := range inputLines {
		jsonMap, jsonParseErr := parseJSONtoMap(line)
		if jsonParseErr != nil {
			logger.Error().Err(jsonParseErr)
			continue
		}
		record := make([]string, len(keys))
		record = buildRecordFromMap(jsonMap, keys, "", record)

		// We won't ever know which key represents timestamp or hostname for 'generic' parsed artifacts
		helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", record, outputChannel)
		// We don't do 'super timeline' for 'generic' artifacts since we don't know their time fields, metadata, etc.
		// We assume that every 'row' of a particular file will have the same string of keys across all rows
		// This is not a perfect assumption for deeply nested data but will work 'good enough' for most use cases
	}
}
