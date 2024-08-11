
// AutobahnChargingStationTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app



import (
    
    "encoding/json"
    "errors"
    "github.com/apioo/sdkgen-go"
    "io"
    "net/http"
    "net/url"
    
)

type AutobahnChargingStationTag struct {
    internal *sdkgen.TagAbstract
}



// GetAll Returns available charging stations for a specific autobahn
func (client *AutobahnChargingStationTag) GetAll(autobahnId string) (AutobahnChargingStationCollection, error) {
    pathParams := make(map[string]interface{})
    pathParams["autobahn_id"] = autobahnId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/autobahn/:autobahn_id/charging_station", pathParams))
    if err != nil {
        return AutobahnChargingStationCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return AutobahnChargingStationCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return AutobahnChargingStationCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return AutobahnChargingStationCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response AutobahnChargingStationCollection
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return AutobahnChargingStationCollection{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 400:
            var response Message
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return AutobahnChargingStationCollection{}, err
            }

            return AutobahnChargingStationCollection{}, &MessageException{
                Payload: response,
            }
        case 404:
            var response Message
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return AutobahnChargingStationCollection{}, err
            }

            return AutobahnChargingStationCollection{}, &MessageException{
                Payload: response,
            }
        case 500:
            var response Message
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return AutobahnChargingStationCollection{}, err
            }

            return AutobahnChargingStationCollection{}, &MessageException{
                Payload: response,
            }
        default:
            return AutobahnChargingStationCollection{}, errors.New("the server returned an unknown status code")
    }
}



func NewAutobahnChargingStationTag(httpClient *http.Client, parser *sdkgen.Parser) *AutobahnChargingStationTag {
	return &AutobahnChargingStationTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
