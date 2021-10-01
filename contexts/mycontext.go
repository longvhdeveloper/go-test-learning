package contexts

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//ctx := request.Context()
		//data := make(chan string, 1)
		//
		//go func() {
		//	data <- store.Fetch()
		//}()
		//
		//select {
		//case d := <-data:
		//	fmt.Fprint(writer, d)
		//case <-ctx.Done():
		//	store.Cancel()
		//}

		data, err := store.Fetch(request.Context())

		if err != nil {
			return
		}

		fmt.Fprint(writer, data)
	})
}
