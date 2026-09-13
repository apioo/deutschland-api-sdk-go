
# deutschland-api-sdk-go

This [SDK](https://github.com/apioo/deutschland-api-sdk-go) is managed by the [SDK Fabric](https://sdk-fabric.org/) project, a global infrastructure to
automatically generate SDKs for every API.

You can find more information about this SDK at [TypeHub](https://typehub.cloud/):
https://app.typehub.cloud/d/deutschland-api/sdk

## Usage

```go
import (
	"github.com/apioo/deutschland-api-sdk-go/sdk"
)

var client, _ = sdk.Build("[access_token]");

// Returns all available warnings from the modular warning system (MoWaS).
response, err := client.Warning().getAll()

// Returns a specific city.
response, err := client.City().get("city_id")

// Returns all available cities.
response, err := client.City().getAll(1, "state", "district", "name", "zipCode")

// Returns a specific district.
response, err := client.District().get("district_id")

// Returns all available districts.
response, err := client.District().getAll(1, "state", "name")

// Returns a specific state.
response, err := client.State().get("state_id")

// Returns all available states.
response, err := client.State().getAll(1, "name")

// Returns specific member of the Bundestag.
response, err := client.Bundestag().Member().get("member_id")

// Returns all current members of the Bundestag.
response, err := client.Bundestag().Member().getAll()

// Returns all current members of the Bundesrat.
response, err := client.Bundesrat().Member().getAll()

// Returns available warnings for a specific autobahn.
response, err := client.Autobahn().Warning().getAll("autobahn_id")

// Returns available parking lorries for a specific autobahn.
response, err := client.Autobahn().Parking_lorry().getAll("autobahn_id")

// Returns available closures for a specific autobahn.
response, err := client.Autobahn().Closure().getAll("autobahn_id")

// Returns available charging stations for a specific autobahn.
response, err := client.Autobahn().Charging_station().getAll("autobahn_id")

// Returns all available autobahns.
response, err := client.Autobahn().getAll()

response, err := client.Authorization().getWhoami()

response, err := client.Authorization().revoke()

response, err := client.Meta().getAbout()
```
