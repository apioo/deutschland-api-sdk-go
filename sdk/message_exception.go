
// MessageException automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app



import (
    "encoding/json"
    "fmt"
)

type MessageException struct {
    Payload Message
}

func (e *MessageException) Error() string {
    raw, err := json.Marshal(e.Payload)
    if err != nil {
        return "could not marshal provided JSON data"
    }

    return fmt.Sprintf("The server returned an error: %s", raw)
}
