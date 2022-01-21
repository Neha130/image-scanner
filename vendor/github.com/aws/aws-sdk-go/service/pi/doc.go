// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package pi provides the client and types for making API
// requests to AWS Performance Insights.
//
// Amazon RDS Performance Insights enables you to monitor and explore different
// dimensions of database load based on data captured from a running DB instance.
// The guide provides detailed information about Performance Insights data types,
// parameters and errors.
//
// When Performance Insights is enabled, the Amazon RDS Performance Insights
// API provides visibility into the performance of your DB instance. Amazon
// CloudWatch provides the authoritative source for Amazon Web Services service-vended
// monitoring metrics. Performance Insights offers a domain-specific view of
// DB load.
//
// DB load is measured as average active sessions. Performance Insights provides
// the data to API consumers as a two-dimensional time-series dataset. The time
// dimension provides DB load data for each time point in the queried time range.
// Each time point decomposes overall load in relation to the requested dimensions,
// measured at that time point. Examples include SQL, Wait event, User, and
// Host.
//
//    * To learn more about Performance Insights and Amazon Aurora DB instances,
//    go to the Amazon Aurora User Guide (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PerfInsights.html).
//
//    * To learn more about Performance Insights and Amazon RDS DB instances,
//    go to the Amazon RDS User Guide (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html).
//
// See https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27 for more information on this service.
//
// See pi package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/pi/
//
// Using the Client
//
// To contact AWS Performance Insights with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Performance Insights client PI for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/pi/#New
package pi
