
// autobahn_parking_lorry automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


type AutobahnParkingLorry struct {
    Id string `json:"id"`
    Title string `json:"title"`
    Subtitle string `json:"subtitle"`
    Description []string `json:"description"`
    Coordinate *AutobahnCoordinate `json:"coordinate"`
    Features []string `json:"features"`
}
