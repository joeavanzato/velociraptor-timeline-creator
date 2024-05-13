package helpers

import "C"
import (
	"crypto/md5"
	"encoding/csv"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"time"
)

func SetupLogger() zerolog.Logger {
	logFileName := vars.LogFile
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Couldn't Initialize Log File: %s", err)
		if err != nil {
			panic(nil)
		}
		panic(err)
	}
	cw := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
	}
	cw.NoColor = true
	mw := io.MultiWriter(cw, logFile)
	logger := zerolog.New(mw).Level(zerolog.TraceLevel)
	logger = logger.With().Timestamp().Caller().Logger()
	return logger
}

func DoesFileOrDirExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CloseChannelWhenDone(c chan []string, wg *sync.WaitGroup) {
	// Waits on a WaitGroup to conclude and closes the associated channel when done - used to synchronize waitgroups sending to a channel
	wg.Wait()
	close(c)
}

func BuildAndSendArtifactRecord(timestamp string, clientID string, hostname string, record []string, channel chan<- []string) {
	r := append([]string{timestamp, clientID, hostname}, record...)
	channel <- r
}

func GetStructHeadersAsStringSlice(s interface{}) []string {
	t := reflect.TypeOf(s)
	names := make([]string, t.NumField())
	for i := range names {
		names[i] = t.Field(i).Name
	}
	return names
}

func getAttr(obj interface{}, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		panic("not struct")
	}
	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		panic("not found:" + fieldName)
	}
	return curField
}

func GetStructValuesAsStringSlice(s interface{}) []string {
	// Useful for simple structs that are 1-Dimensional with 'basic' fields
	t := reflect.TypeOf(s)
	rv := reflect.ValueOf(&s).Elem().Elem()
	structValues := make([]string, rv.NumField())
	for i := range structValues {
		tmp := rv.FieldByName(t.Field(i).Name)
		structValues[i] = fmt.Sprint(tmp)
	}
	return structValues
}

func ListenOnWriteChannel(c chan []string, w *csv.Writer, logger zerolog.Logger, outputF *os.File, bufferSize int, wait *sync.WaitGroup) {
	// TODO - Consider having pool of routines appending records to slice [][]string and a single reader drawing from this to avoid any bottle-necks
	// TODO - Consider sending writer in a goroutine with wait group, refilling buffer, etc.
	// TODO - Hash incoming records and throw out duplicates
	// Receives handle to a file, csv writer and input channel (among other things)
	// Listens for records on the channel and writes them to CSV when the buffer fills up or until channel is closed
	// Channel close causes a buffer flush to disk and return
	defer outputF.Close()
	defer wait.Done()
	hashTrack := make(map[[16]byte]struct{})

	tempRecords := make([][]string, 0)
	for {
		record, ok := <-c
		if !ok {
			break
		} else if len(tempRecords) <= bufferSize {
			// Track hashes of incoming records and throw out duplicates that may come in - this is useful since we may
			// encounter multiples of the same collectin when parsing velociraptor data store and we aren't currently pulling 'latest' versions of file
			// TODO: When 'latest' version of same file per-client is established, we can remove this safely.
			hash := md5.Sum([]byte(fmt.Sprint(record)))
			_, recordAlreadyProcessed := hashTrack[hash]
			if !recordAlreadyProcessed {
				hashTrack[hash] = struct{}{}
				tempRecords = append(tempRecords, record)
			}
		} else {
			err := w.WriteAll(tempRecords)
			if err != nil {
				logger.Error().Msg(err.Error())
			}
			tempRecords = nil
		}
	}
	err := w.WriteAll(tempRecords)
	if err != nil {
		logger.Error().Msg(err.Error())
	}
	w.Flush()
	err = w.Error()
	if err != nil {
		logger.Error().Msg(err.Error())
	}
}

func setupReadWrite(inputF *os.File, outputF *os.File) (*csv.Reader, *csv.Writer, error) {
	writer := csv.NewWriter(outputF)
	parser := csv.NewReader(inputF)
	parser.LazyQuotes = true
	return parser, writer, nil
}

func GetAllJSONFromDirectory(path string) []string {
	paths := make([]string, 0)
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, ".json") {
				paths = append(paths, path)
			}
			return nil
		})
	if err != nil {
		fmt.Println(err)
	}
	return paths
}

func CreateOutput(outputFile string) (*os.File, error) {
	outputF, err := os.Create(outputFile)
	return outputF, err
}

func OpenInput(inputFile string) (*os.File, error) {
	inputF, err := os.Open(inputFile)
	return inputF, err
}

func GetNewPW(logger zerolog.Logger, inputFile string, outputFile string) (*csv.Reader, *csv.Writer, *os.File, *os.File, error) {
	inputF, err := OpenInput(inputFile)
	if err != nil {
		logger.Error().Msg(err.Error())
	}
	outputF, err := CreateOutput(outputFile)
	if err != nil {
		logger.Error().Msg(err.Error())
	}
	parser, writer, err := setupReadWrite(inputF, outputF)
	if err != nil {
		logger.Error().Msg(err.Error())
	}
	return parser, writer, inputF, outputF, err
}
