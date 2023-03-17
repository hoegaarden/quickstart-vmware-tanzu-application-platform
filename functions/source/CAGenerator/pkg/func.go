package pkg

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/cfn"
)

func CAGeneratorFunc(ctx context.Context, event cfn.Event) (string, map[string]interface{}, error) {

	switch event.RequestType {
	case cfn.RequestCreate:
		generator := CA{}

		cn, ok := event.ResourceProperties["CommonName"].(string)
		if !ok {
			return "", nil, fmt.Errorf("invalid CommonName")
		}
		generator.CommonName = cn

		pub, priv, err := generator.Generate()
		if err != nil {
			return "", nil, err
		}

		resData := map[string]any{
			"pub":  pub,
			"priv": priv,
		}

		return "", resData, nil

	case cfn.RequestUpdate, cfn.RequestDelete:
		return "", nil, nil
	}

	return "", nil, fmt.Errorf("unkown request type: '%s'", event.RequestType)
}
