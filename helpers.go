package telego

import (
	"fmt"
	"net/http"
	"time"
)

const (
	updateChanBuffer = 100

	defaultUpdateInterval = time.Second / 2 // 0.5s
	retryTimeout          = time.Second * 3 // 3s
)

// SetUpdateInterval - Sets interval of calling GetUpdates in GetUpdatesChan method. Ensures that between two calls
// of GetUpdates will be at least specified time, but it could be longer.
func (b *Bot) SetUpdateInterval(interval time.Duration) {
	b.updateInterval = interval
}

// StopGettingUpdates - Stop reviving updates from GetUpdatesChan method
func (b *Bot) StopGettingUpdates() {
	close(b.stopChannel)
}

// GetUpdatesChan - Receive updates in chan
func (b *Bot) GetUpdatesChan(params *GetUpdatesParams) (chan Update, error) {
	b.stopChannel = make(chan struct{})
	updatesChan := make(chan Update, updateChanBuffer)

	go func() {
		for {
			select {
			case <-b.stopChannel:
				close(updatesChan)
				return
			default:
				// Continue getting updates
			}

			updates, err := b.GetUpdates(params)
			if err != nil {
				if b.printErrors {
					fmt.Printf("%s Getting updates: %v\n", logStarting(errorMode), err)
					fmt.Printf("%s Retrying to get updates in %s\n", logStarting(errorMode), retryTimeout.String())
				}

				time.Sleep(retryTimeout)
				continue
			}

			for _, update := range updates {
				if update.UpdateID >= params.Offset {
					params.Offset = update.UpdateID + 1
					updatesChan <- update
				}
			}

			time.Sleep(b.updateInterval)
		}
	}()

	return updatesChan, nil
}

// StartListeningForWebhook - Start server for listening for webhook
func (b *Bot) StartListeningForWebhook(address, certificateFile, keyFile string) {
	go func() {
		err := http.ListenAndServeTLS(address, certificateFile, keyFile, nil)
		if err != nil && b.printErrors {
			fmt.Printf("%s Listening for webhook: %v\n", logStarting(errorMode), err)
		}
	}()
}

// ListenForWebhook - Receive updates in chan from webhook
func (b *Bot) ListenForWebhook(pattern string) (chan Update, error) {
	updatesChan := make(chan Update, updateChanBuffer)

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			err := fmt.Errorf("used invalid HTTP method: %q, required method: %q", r.Method, http.MethodPost)
			respondWithError(w, err)
			return
		}

		var update Update
		err := json.NewDecoder(r.Body).Decode(&update)
		if err != nil {
			respondWithError(w, fmt.Errorf("decoding update: %w", err))
			return
		}

		updatesChan <- update

		w.WriteHeader(http.StatusOK)
	})

	return updatesChan, nil
}

func respondWithError(w http.ResponseWriter, err error) {
	errMsg, _ := json.Marshal(map[string]string{"error": err.Error()})

	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", jsonContentType)

	_, _ = w.Write(errMsg)
}
