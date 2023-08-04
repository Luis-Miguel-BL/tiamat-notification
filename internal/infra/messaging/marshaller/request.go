package marshaller

import (
	"encoding/json"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/api"
)

func UnmarshalRequestBody[T api.RequestBody](req api.Request) (parsedBody T, err error) {
	err = json.Unmarshal([]byte(req.Body), parsedBody)
	if err != nil {
		return parsedBody, err
	}

	err = parsedBody.Validate()
	if err != nil {
		return parsedBody, err
	}

	return parsedBody, nil
}
