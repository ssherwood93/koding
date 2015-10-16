// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package rds

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/aws/service/serviceinfo"
	"github.com/aws/aws-sdk-go/internal/protocol/query"
	"github.com/aws/aws-sdk-go/internal/signer/v4"
)

// Amazon Relational Database Service (Amazon RDS) is a web service that makes
// it easier to set up, operate, and scale a relational database in the cloud.
// It provides cost-efficient, resizeable capacity for an industry-standard
// relational database and manages common database administration tasks, freeing
// up developers to focus on what makes their applications and businesses unique.
//
//  Amazon RDS gives you access to the capabilities of a MySQL, PostgreSQL,
// Microsoft SQL Server, Oracle, or Aurora database server. This means the code,
// applications, and tools you already use today with your existing databases
// work with Amazon RDS without modification. Amazon RDS automatically backs
// up your database and maintains the database software that powers your DB
// instance. Amazon RDS is flexible: you can scale your database instance's
// compute resources and storage capacity to meet your application's demand.
// As with all Amazon Web Services, there are no up-front investments, and you
// pay only for the resources you use.
//
//  This is an interface reference for Amazon RDS. It contains documentation
// for a programming or command line interface you can use to manage Amazon
// RDS. Note that Amazon RDS is asynchronous, which means that some interfaces
// might require techniques such as polling or callback functions to determine
// when a command has been applied. In this reference, the parameter descriptions
// indicate whether a command is applied immediately, on the next instance reboot,
// or during the maintenance window. For a summary of the Amazon RDS interfaces,
// go to Available RDS Interfaces (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Welcome.html#Welcome.Interfaces).
type RDS struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new RDS client.
func New(config *aws.Config) *RDS {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "rds",
			APIVersion:  "2014-10-31",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &RDS{service}
}

// newRequest creates a new request for a RDS operation and runs any
// custom request initialization.
func (c *RDS) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
