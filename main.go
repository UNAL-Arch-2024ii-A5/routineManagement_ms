package main

import (
	"context"
	"fmt"

	"github.com/hectorhernandezalfonso/exercise_ms.git/service"
)

func main() {
	service := service.New()
	err := service.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}
}
