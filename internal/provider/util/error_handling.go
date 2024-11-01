package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	config_api_client "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
)

func GenerateErrorSummary(actionName string, entityName string) string {
	return fmt.Sprintf("Error performing %s on %s", actionName, entityName)
}

func RaiseForStatus(response *http.Response, err error) (bool, string) {
	if err != nil {
		var detail string
		var data map[string]interface{}

		switch e := err.(type) {
		case *url.Error:
			detail = handleURLError(e)
		case *config_api_client.GenericOpenAPIError:
			if jsonDecodeErr := json.NewDecoder(response.Body).Decode(&data); jsonDecodeErr != nil {
				detail = "Unexpected error: " + jsonDecodeErr.Error()
			} else if message, ok := data["message"]; ok {
				detail = message.(string)
				if debugId, ok := data["debugId"]; ok {
					detail += "\nDebugID: " + debugId.(string)
				}
			} else {
				detail = "Unexpected error: " + e.Error()
			}
		default:
			detail = "Unexpected error: " + e.Error()
		}

		return true, detail
	}
	return false, ""
}

func handleURLError(uErr *url.Error) string {
	if uErr.Timeout() {
		return "Error: Request timed out. Please check your network."
	} else if uErr.Temporary() {
		return "Error: Temporary network error. Please try again later."
	} else {
		return fmt.Sprintf("URL Error: %v\n", uErr)
	}
}