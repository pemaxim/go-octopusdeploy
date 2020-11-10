package examples

import (
	"fmt"
	"net/url"

	"github.com/pemaxim/go-octopusdeploy/octopusdeploy"
)

func GetFeedByIDExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// feed values
		feedID string = "feed-id"
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

	// get feed by its ID
	feed, err := client.Feeds.GetByID(feedID)
	if err != nil {
		_ = fmt.Errorf("error getting feed: %v", err)
		return
	}

	fmt.Printf("feed: (%s)\n", feed.GetID())
}
