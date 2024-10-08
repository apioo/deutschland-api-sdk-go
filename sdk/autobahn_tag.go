
// AutobahnTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app



import (
    
    "encoding/json"
    "errors"
    "github.com/apioo/sdkgen-go"
    "io"
    "net/http"
    "net/url"
    
)

type AutobahnTag struct {
    internal *sdkgen.TagAbstract
}

func (client *AutobahnTag) Warning() *AutobahnWarningTag {
    return NewAutobahnWarningTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *AutobahnTag) ParkingLorry() *AutobahnParkingLorryTag {
    return NewAutobahnParkingLorryTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *AutobahnTag) Closure() *AutobahnClosureTag {
    return NewAutobahnClosureTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *AutobahnTag) ChargingStation() *AutobahnChargingStationTag {
    return NewAutobahnChargingStationTag(client.internal.HttpClient, client.internal.Parser)
}



// GetAll Returns all available autobahns
func (client *AutobahnTag) GetAll() (AutobahnCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/autobahn", pathParams))
    if err != nil {
        return AutobahnCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return AutobahnCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return AutobahnCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return AutobahnCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response AutobahnCollection
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return AutobahnCollection{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 400:
            var response Message
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return AutobahnCollection{}, err
            }

            return AutobahnCollection{}, &MessageException{
                Payload: response,
            }
        case 404:
            var response Message
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return AutobahnCollection{}, err
            }

            return AutobahnCollection{}, &MessageException{
                Payload: response,
            }
        case 500:
            var response Message
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return AutobahnCollection{}, err
            }

            return AutobahnCollection{}, &MessageException{
                Payload: response,
            }
        default:
            return AutobahnCollection{}, errors.New("the server returned an unknown status code")
    }
}



func NewAutobahnTag(httpClient *http.Client, parser *sdkgen.Parser) *AutobahnTag {
	return &AutobahnTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
