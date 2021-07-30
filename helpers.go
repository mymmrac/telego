package telego

import (
	"fmt"
	"net/http"
	"time"
)

const (
	updateChanBuffer = 100

	defaultUpdateInterval = time.Second
	retryTimeout          = time.Second * 3
)

func (b *Bot) SetUpdateInterval(interval time.Duration) {
	b.updateInterval = interval
}

func (b *Bot) StopGettingUpdates() {
	close(b.stopChannel)
}

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
				b.log.Error(err)
				b.log.Infof("Retrying to get updates in %s", retryTimeout.String())

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

func (b *Bot) StartListeningForWebhook(address, certificateFile, keyFile string) {
	go func() {
		err := http.ListenAndServeTLS(address, certificateFile, keyFile, nil)
		if err != nil {
			b.log.Errorf("listening for webhook: %v", err)
		}
	}()
}

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
