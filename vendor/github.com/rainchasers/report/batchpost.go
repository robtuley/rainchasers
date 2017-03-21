package report

import (
	"bytes"
	"log"
	"net/http"
	"sync"
	"time"
)

// BatchPostToURL HTTP POSTs events in batches to a URL
func BatchPostToURL(url string) {
	log.Println("reporting:> batch POST to " + url)

	go batchPoster(url)
}

func batchPoster(url string) {
	var wg sync.WaitGroup
	stopping := false

	client := &http.Client{}
	ticker := time.NewTicker(2 * time.Second)

	for {
		buffer := ""

	buffering:
		for {
			select {
			case json, more := <-channel.JSONEncoded:
				if !more {
					stopping = true
					break buffering
				}
				if len(buffer) > 0 {
					buffer += "\n"
				}
				buffer += json
			case <-ticker.C:
				if len(buffer) > 0 {
					break buffering
				}
			}
		}

		if len(buffer) > 0 {
			wg.Add(1)
			go func(data string) {
				defer wg.Done()

				req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
				if err != nil {
					log.Println("error:> ", err)
					return
				}
				req.Header.Set("Content-Type", "text/plain")

				resp, err := client.Do(req)
				if err != nil {
					log.Println("error:> ", err)
					return
				}
				defer resp.Body.Close()

				if resp.StatusCode != http.StatusOK {
					log.Println("error:> status code ", resp.StatusCode)
					return
				}

			}(buffer)
		}

		if stopping {
			ticker.Stop()
			wg.Wait()
			channel.DrainSignal <- true
			return
		}
	}
}
