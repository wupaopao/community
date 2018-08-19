package main

import (
	"fmt"

	"github.com/mz-eco/mz/app"
	"github.com/mz-eco/mz/http"
)

type Application struct {
}

func (m *Application) Run(args []string) {

	service := http.NewService()
	service.AddHandlers(impls.Handlers)

	//TODO: before application running
	service.Run()

}

func (m *Application) Flags(flags *app.Flags) {
	//TODO: application flags
}

func (m *Application) GetName() string {
	return "community"
}

func main() {
	err := app.Main(&Application{})

	if err != nil {
		fmt.Println(err)
	}
}
