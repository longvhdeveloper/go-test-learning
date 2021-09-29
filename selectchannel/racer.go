package selectchannel

import (
	"fmt"
	"net/http"
	"time"
)

var timeoutSecond = 10 * time.Millisecond

func Racer(slowUrl, fastUrl string) (string, error) {
	//aDuration := measureResponseTime(slowUrl)
	//bDuration := measureResponseTime(fastUrl)
	//
	//if bDuration < aDuration {
	//	return fastUrl
	//}
	//
	//return slowUrl
	return ConfigurableRacer(slowUrl, fastUrl, timeoutSecond)

}

func ConfigurableRacer(slowUrl, fastUrl string, timeout time.Duration) (string, error) {
	select {
	case <-ping(fastUrl):
		return fastUrl, nil
	case <-ping(slowUrl):
		return slowUrl, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout waiting for %s and %s", slowUrl, fastUrl)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func(url string) {
		http.Get(url)
		close(ch)
	}(url)

	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
