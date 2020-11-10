package examples

import (
	"fmt"
	"net/url"

	"github.com/pemaxim/go-octopusdeploy/octopusdeploy"
)

func GetProjectByIDExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// project values
		projectID string = "project-id"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, spaceID)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	// get project by its ID
	project, err := client.Projects.GetByID(projectID)
	if err != nil {
		_ = fmt.Errorf("error getting project: %v", err)
		return
	}

	fmt.Printf("project: (%s)\n", project.GetID())
}
