package tool

import (
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type notifier struct {
	listener *pq.Listener
	failed   chan error
}

var Config ConfigObject

func (n *notifier) Fetch(data chan []byte) error {
	//var fetchCounter uint64
	log.Println("start fetch funcion ")

	for {
		select {
		case e := <-n.listener.Notify:
			if e == nil {
				continue
			}
			log.Println("listener.Notify ", e.Extra)
			// fetchCounter++
			// data <- []byte(e.Extra)
			log.Println("FETCHED DAta", []byte(e.Extra))
			CallPusherClient(e.Extra, Config.PusherKey)
		case err := <-n.failed:
			return err
		case <-time.After(time.Minute):
			go n.listener.Ping()
		}
	}
}

func NewNotifier(dsn, channelName string, config ConfigObject) (*notifier, error) {
	Config = config
	n := &notifier{failed: make(chan error, 2)}

	listener := pq.NewListener(
		dsn,
		3*time.Second, time.Minute,
		n.logListener)

	if err := listener.Listen(channelName); err != nil {
		listener.Close()
		log.Println("ERROR!:", err)
		return nil, err
	}
	fmt.Println("Successfully connected ja !")

	n.listener = listener
	return n, nil
}

// logListener is the state change callback for the listener.
func (n *notifier) logListener(event pq.ListenerEventType, err error) {
	fmt.Println("in logListener")

	if err != nil {
		log.Printf("listener error: %s\n", err)
	}
	if event == pq.ListenerEventConnectionAttemptFailed {
		n.failed <- err
	}
}
