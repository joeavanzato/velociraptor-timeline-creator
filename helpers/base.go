package helpers

import (
	"encoding/csv"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"io"
	"os"
	"path/filepath"
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
	logger = logger.With().Timestamp().Logger()
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

func ListenOnWriteChannel(c chan []string, w *csv.Writer, logger zerolog.Logger, outputF *os.File, bufferSize int, wait *sync.WaitGroup) {
	// TODO - Consider having pool of routines appending records to slice [][]string and a single reader drawing from this to avoid any bottle-necks
	// TODO - Consider sending writer in a goroutine with wait group, refilling buffer, etc.
	defer outputF.Close()
	defer wait.Done()
	tempRecords := make([][]string, 0)
	for {
		record, ok := <-c
		if !ok {
			break
		} else if len(tempRecords) <= bufferSize {
			tempRecords = append(tempRecords, record)
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
