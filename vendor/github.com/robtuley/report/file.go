package report

import (
	"log"
	"os"
	"time"
)

// File write events to the specified log file
func File(filename string) {
	go func() {
		for {
			if err := fileWriter(filename); err != nil {
				log.Println("error:> ", err, filename)
				time.Sleep(time.Second) // prevent error thrashing
				continue
			}
			return
		}
	}()
}

func fileWriter(filename string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	for {
		json, more := <-channel.JSONEncoded
		if !more {
			channel.DrainSignal <- true
			return nil
		}
		if _, err = f.WriteString(json + "\n"); err != nil {
			return err
		}
	}
}
