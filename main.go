// Copyright © 2020 Aggelos Kolaitis
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"os"

	"github.com/juju/juju/api"
	"github.com/juju/juju/jujuclient"
	"github.com/juju/names/v4"
)

func handleError(err error) {
	if err != nil {
		log.Fatalf("An error occured: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	// store is used to access configuration from ~/.local/share/juju
	store := jujuclient.NewFileClientStore()

	// Retrieve name of default controller (from ~/.local/share/juju/controllers.yml)
	// Also see `store.AllControllers()`
	controllerName, err := store.CurrentController()
	handleError(err)
	log.Printf("Using default controller: %s\n", controllerName)

	// Retrieve controller configuration (UUID, API endpoints, certs, etc)
	// Also see `store.ControllerByAPIEndpoints()`
	controller, err := store.ControllerByName(controllerName)
	if err != nil {
		panic(err)
	}

	// Retrieve account credentials for controller (from ~/.local/share/juju/accounts.yml)
	account, err := store.AccountDetails(controllerName)
	if err != nil {
		panic(err)
	}
	log.Printf("Using default user: %s\n", account.User)

	// Retrieve the current model (from ~/.local/share/juju/models.yml)
	// Also see `store.AllModels()`.
	modelName, err := store.CurrentModel(controllerName)
	if err != nil {
		panic(err)
	}
	log.Printf("Using default model: %s\n", modelName)

	// Retrieve model information, e.g. UUID
	model, err := store.ModelByName(controllerName, modelName)
	if err != nil {
		panic(err)
	}
	log.Printf("Model UUID is: %s\n", model.ModelUUID)

	// Open a new connection to the API, using the loaded configuration.
	// Connection requires: API endpoints, user credentials and model UUID.
	// Also see `juju.NewAPIConnection()`.
	conn, err := api.Open(&api.Info{
		Addrs:       controller.APIEndpoints,
		CACert:      controller.CACert,
		SNIHostName: controller.PublicDNSName, // optional
		ModelTag:    names.NewModelTag(model.ModelUUID),
		Tag:         names.NewUserTag(account.User),
		Password:    account.Password,
	}, api.DefaultDialOpts())
	handleError(err)
	log.Printf("Connected to %s at %s as %s\n", conn.ControllerTag(), conn.Addr(), conn.ControllerAccess())
	defer conn.Close()

	// client is used to execute Juju commands
	client := conn.Client()
	defer client.Close()

	// Retrieve Juju status for the selected model, optionally with a list of patterns
	status, err := client.Status([]string{""})
	handleError(err)

	// Range over model applications.
	log.Printf("Model has %d machines and %d applications\n", len(status.Applications), len(status.Machines))
	for appName, app := range status.Applications {
		log.Printf("Application %s (%s) has %d units\n", appName, app.Charm, len(app.Units))
	}

}
