package es

import (
	"context"
)

func Del() error {
	_, err := client.DeleteIndex("message").Do(context.Background())
	return err
}
