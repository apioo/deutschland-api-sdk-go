
# DeutschlandAPI SDK

This SDK helps to access the [DeutschlandAPI](https://deutschland-api.dev)

## Usage

The following example shows how you initialize the client:

```go
import (
	"github.com/apioo/deutschland-api-sdk-go/sdk"
)

var client, _ = sdk.BuildAnonymous();

var collection = client.state().getAll()
for _, state := range collection.Entries {
    fmt.Println(state.Name)
}

```

More information about the complete API at:
https://app.typehub.cloud/d/deutschland-api/sdk
