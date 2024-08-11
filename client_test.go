package main

import (
	"github.com/apioo/deutschland-api-sdk-go/sdk"
	"github.com/apioo/sdkgen-go"
	"testing"
)

func TestClient(t *testing.T) {

	client, err := sdk.Build("my_token")
	if err != nil {
		t.Error(err)
	}

}
