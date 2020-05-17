package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"time"
)

const appName = "schoolScheduler"

var (
	w *log.Logger // waring
	i *log.Logger // info
	e *log.Logger // error
	f *log.Logger // fatal
)

var logFile *os.File
var err error

func Init(logDir string) {
	// generate a new file name
	t := time.Now()
	logDir = fmt.Sprintf(
		"%s/%s-%s.log",
		logDir,
		appName,
		t.Format("20060102150405"),
	)
	// If the file doesn't exist, create it or append to the file
	logFile, err = os.OpenFile(logDir, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// print the log file
	fmt.Println("log > ", logDir)

	// write to console and standard out
	b := io.MultiWriter(logFile, os.Stderr)

	// create new type of logger for each type of log
	i = log.New(b, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	w = log.New(b, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	e = log.New(b, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	f = log.New(b, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// close the file when the program exits
	c := make(chan os.Signal) // channel to send signal data
	// this function takes a signal channel and sends data through it
	// when the given given singal occurs
	signal.Notify(c, os.Interrupt)

	// goroutine that waits for the os signal thow channel
	// and closes the file when the program is interrupted
	go func() {
		select {
		case <-c:
			logFile.Close()
			os.Exit(1)
		}
	}()

}

// Info write logs with info tag
//  ip - ip address of the client
//  clientID - ID of the client
//  serverID - ID of the server that sent the data
//  lSeq - length of the sequence
//  gSize - size of gene
//  nNType - no. of nucleotide type
//  msg - log messeage
func Info(clientID, serverID string, lSeq, gSize, nNType int, msg string) {
	i.Println(
		clientID,
		serverID,
		lSeq,
		gSize,
		nNType,
		msg,
	)
}

// Warning write logs with info tag
//  ip - ip address of the client
//  clientID - ID of the client
//  serverID - ID of the server that sent the data
//  lSeq - length of the sequence
//  gSize - size of gene
//  nNType - no. of nucleotide type
//  msg - log messeage
func Warning(clientID, serverID string, lSeq, gSize, nNType int, msg string) {
	i.Println(
		clientID,
		serverID,
		lSeq,
		gSize,
		nNType,
		msg,
	)
}

// Error write logs with info tag
//  ip - ip address of the client
//  clientID - ID of the client
//  serverID - ID of the server that sent the data
//  lSeq - length of the sequence
//  gSize - size of gene
//  nNType - no. of nucleotide type
//  msg - log messeage
func Error(clientID, serverID string, lSeq, gSize, nNType int, msg string) {
	e.Println(
		clientID,
		serverID,
		lSeq,
		gSize,
		nNType,
		msg,
	)
}

// Warning write logs with info tag
//  ip - ip address of the client
//  clientID - ID of the client
//  serverID - ID of the server that sent the data
//  lSeq - length of the sequence
//  gSize - size of gene
//  nNType - no. of nucleotide type
//  msg - log messeage
func Fatal(clientID, serverID string, lSeq, gSize, nNType int, msg string) {
	f.Println(
		clientID,
		serverID,
		lSeq,
		gSize,
		nNType,
		msg,
	)
	logFile.Close()

	os.Exit(1)
}
