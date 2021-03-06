// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package route53

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol/restxml"
	"github.com/aws/aws-sdk-go/private/signer/v4"
)

// Route53 is a client for Route 53.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type Route53 struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "route53"

// New creates a new instance of the Route53 client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a Route53 client from just a session.
//     svc := route53.New(mySession)
//
//     // Create a Route53 client with additional configuration
//     svc := route53.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *Route53 {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *Route53 {
	svc := &Route53{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2013-04-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(restxml.Build)
	svc.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a Route53 operation and runs any
// custom request initialization.
func (c *Route53) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
