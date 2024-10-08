
// MetaTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app



import (
    
    "encoding/json"
    "errors"
    "github.com/apioo/sdkgen-go"
    "io"
    "net/http"
    "net/url"
    
)

type MetaTag struct {
    internal *sdkgen.TagAbstract
}



// GetAbout 
func (client *MetaTag) GetAbout() (SystemAbout, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/", pathParams))
    if err != nil {
        return SystemAbout{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return SystemAbout{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return SystemAbout{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return SystemAbout{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response SystemAbout
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return SystemAbout{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return SystemAbout{}, errors.New("the server returned an unknown status code")
    }
}



func NewMetaTag(httpClient *http.Client, parser *sdkgen.Parser) *MetaTag {
	return &MetaTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
