// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ecriface provides an interface to enable mocking the Amazon EC2 Container Registry service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ecriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ecr"
)

// ECRAPI provides an interface to enable mocking the
// ecr.ECR service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon EC2 Container Registry.
//    func myFunc(svc ecriface.ECRAPI) bool {
//        // Make svc.BatchCheckLayerAvailability request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := ecr.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockECRClient struct {
//        ecriface.ECRAPI
//    }
//    func (m *mockECRClient) BatchCheckLayerAvailability(input *ecr.BatchCheckLayerAvailabilityInput) (*ecr.BatchCheckLayerAvailabilityOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockECRClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ECRAPI interface {
	BatchCheckLayerAvailability(*ecr.BatchCheckLayerAvailabilityInput) (*ecr.BatchCheckLayerAvailabilityOutput, error)
	BatchCheckLayerAvailabilityWithContext(aws.Context, *ecr.BatchCheckLayerAvailabilityInput, ...request.Option) (*ecr.BatchCheckLayerAvailabilityOutput, error)
	BatchCheckLayerAvailabilityRequest(*ecr.BatchCheckLayerAvailabilityInput) (*request.Request, *ecr.BatchCheckLayerAvailabilityOutput)

	BatchDeleteImage(*ecr.BatchDeleteImageInput) (*ecr.BatchDeleteImageOutput, error)
	BatchDeleteImageWithContext(aws.Context, *ecr.BatchDeleteImageInput, ...request.Option) (*ecr.BatchDeleteImageOutput, error)
	BatchDeleteImageRequest(*ecr.BatchDeleteImageInput) (*request.Request, *ecr.BatchDeleteImageOutput)

	BatchGetImage(*ecr.BatchGetImageInput) (*ecr.BatchGetImageOutput, error)
	BatchGetImageWithContext(aws.Context, *ecr.BatchGetImageInput, ...request.Option) (*ecr.BatchGetImageOutput, error)
	BatchGetImageRequest(*ecr.BatchGetImageInput) (*request.Request, *ecr.BatchGetImageOutput)

	BatchGetRepositoryScanningConfiguration(*ecr.BatchGetRepositoryScanningConfigurationInput) (*ecr.BatchGetRepositoryScanningConfigurationOutput, error)
	BatchGetRepositoryScanningConfigurationWithContext(aws.Context, *ecr.BatchGetRepositoryScanningConfigurationInput, ...request.Option) (*ecr.BatchGetRepositoryScanningConfigurationOutput, error)
	BatchGetRepositoryScanningConfigurationRequest(*ecr.BatchGetRepositoryScanningConfigurationInput) (*request.Request, *ecr.BatchGetRepositoryScanningConfigurationOutput)

	CompleteLayerUpload(*ecr.CompleteLayerUploadInput) (*ecr.CompleteLayerUploadOutput, error)
	CompleteLayerUploadWithContext(aws.Context, *ecr.CompleteLayerUploadInput, ...request.Option) (*ecr.CompleteLayerUploadOutput, error)
	CompleteLayerUploadRequest(*ecr.CompleteLayerUploadInput) (*request.Request, *ecr.CompleteLayerUploadOutput)

	CreatePullThroughCacheRule(*ecr.CreatePullThroughCacheRuleInput) (*ecr.CreatePullThroughCacheRuleOutput, error)
	CreatePullThroughCacheRuleWithContext(aws.Context, *ecr.CreatePullThroughCacheRuleInput, ...request.Option) (*ecr.CreatePullThroughCacheRuleOutput, error)
	CreatePullThroughCacheRuleRequest(*ecr.CreatePullThroughCacheRuleInput) (*request.Request, *ecr.CreatePullThroughCacheRuleOutput)

	CreateRepository(*ecr.CreateRepositoryInput) (*ecr.CreateRepositoryOutput, error)
	CreateRepositoryWithContext(aws.Context, *ecr.CreateRepositoryInput, ...request.Option) (*ecr.CreateRepositoryOutput, error)
	CreateRepositoryRequest(*ecr.CreateRepositoryInput) (*request.Request, *ecr.CreateRepositoryOutput)

	DeleteLifecyclePolicy(*ecr.DeleteLifecyclePolicyInput) (*ecr.DeleteLifecyclePolicyOutput, error)
	DeleteLifecyclePolicyWithContext(aws.Context, *ecr.DeleteLifecyclePolicyInput, ...request.Option) (*ecr.DeleteLifecyclePolicyOutput, error)
	DeleteLifecyclePolicyRequest(*ecr.DeleteLifecyclePolicyInput) (*request.Request, *ecr.DeleteLifecyclePolicyOutput)

	DeletePullThroughCacheRule(*ecr.DeletePullThroughCacheRuleInput) (*ecr.DeletePullThroughCacheRuleOutput, error)
	DeletePullThroughCacheRuleWithContext(aws.Context, *ecr.DeletePullThroughCacheRuleInput, ...request.Option) (*ecr.DeletePullThroughCacheRuleOutput, error)
	DeletePullThroughCacheRuleRequest(*ecr.DeletePullThroughCacheRuleInput) (*request.Request, *ecr.DeletePullThroughCacheRuleOutput)

	DeleteRegistryPolicy(*ecr.DeleteRegistryPolicyInput) (*ecr.DeleteRegistryPolicyOutput, error)
	DeleteRegistryPolicyWithContext(aws.Context, *ecr.DeleteRegistryPolicyInput, ...request.Option) (*ecr.DeleteRegistryPolicyOutput, error)
	DeleteRegistryPolicyRequest(*ecr.DeleteRegistryPolicyInput) (*request.Request, *ecr.DeleteRegistryPolicyOutput)

	DeleteRepository(*ecr.DeleteRepositoryInput) (*ecr.DeleteRepositoryOutput, error)
	DeleteRepositoryWithContext(aws.Context, *ecr.DeleteRepositoryInput, ...request.Option) (*ecr.DeleteRepositoryOutput, error)
	DeleteRepositoryRequest(*ecr.DeleteRepositoryInput) (*request.Request, *ecr.DeleteRepositoryOutput)

	DeleteRepositoryPolicy(*ecr.DeleteRepositoryPolicyInput) (*ecr.DeleteRepositoryPolicyOutput, error)
	DeleteRepositoryPolicyWithContext(aws.Context, *ecr.DeleteRepositoryPolicyInput, ...request.Option) (*ecr.DeleteRepositoryPolicyOutput, error)
	DeleteRepositoryPolicyRequest(*ecr.DeleteRepositoryPolicyInput) (*request.Request, *ecr.DeleteRepositoryPolicyOutput)

	DescribeImageReplicationStatus(*ecr.DescribeImageReplicationStatusInput) (*ecr.DescribeImageReplicationStatusOutput, error)
	DescribeImageReplicationStatusWithContext(aws.Context, *ecr.DescribeImageReplicationStatusInput, ...request.Option) (*ecr.DescribeImageReplicationStatusOutput, error)
	DescribeImageReplicationStatusRequest(*ecr.DescribeImageReplicationStatusInput) (*request.Request, *ecr.DescribeImageReplicationStatusOutput)

	DescribeImageScanFindings(*ecr.DescribeImageScanFindingsInput) (*ecr.DescribeImageScanFindingsOutput, error)
	DescribeImageScanFindingsWithContext(aws.Context, *ecr.DescribeImageScanFindingsInput, ...request.Option) (*ecr.DescribeImageScanFindingsOutput, error)
	DescribeImageScanFindingsRequest(*ecr.DescribeImageScanFindingsInput) (*request.Request, *ecr.DescribeImageScanFindingsOutput)

	DescribeImageScanFindingsPages(*ecr.DescribeImageScanFindingsInput, func(*ecr.DescribeImageScanFindingsOutput, bool) bool) error
	DescribeImageScanFindingsPagesWithContext(aws.Context, *ecr.DescribeImageScanFindingsInput, func(*ecr.DescribeImageScanFindingsOutput, bool) bool, ...request.Option) error

	DescribeImages(*ecr.DescribeImagesInput) (*ecr.DescribeImagesOutput, error)
	DescribeImagesWithContext(aws.Context, *ecr.DescribeImagesInput, ...request.Option) (*ecr.DescribeImagesOutput, error)
	DescribeImagesRequest(*ecr.DescribeImagesInput) (*request.Request, *ecr.DescribeImagesOutput)

	DescribeImagesPages(*ecr.DescribeImagesInput, func(*ecr.DescribeImagesOutput, bool) bool) error
	DescribeImagesPagesWithContext(aws.Context, *ecr.DescribeImagesInput, func(*ecr.DescribeImagesOutput, bool) bool, ...request.Option) error

	DescribePullThroughCacheRules(*ecr.DescribePullThroughCacheRulesInput) (*ecr.DescribePullThroughCacheRulesOutput, error)
	DescribePullThroughCacheRulesWithContext(aws.Context, *ecr.DescribePullThroughCacheRulesInput, ...request.Option) (*ecr.DescribePullThroughCacheRulesOutput, error)
	DescribePullThroughCacheRulesRequest(*ecr.DescribePullThroughCacheRulesInput) (*request.Request, *ecr.DescribePullThroughCacheRulesOutput)

	DescribePullThroughCacheRulesPages(*ecr.DescribePullThroughCacheRulesInput, func(*ecr.DescribePullThroughCacheRulesOutput, bool) bool) error
	DescribePullThroughCacheRulesPagesWithContext(aws.Context, *ecr.DescribePullThroughCacheRulesInput, func(*ecr.DescribePullThroughCacheRulesOutput, bool) bool, ...request.Option) error

	DescribeRegistry(*ecr.DescribeRegistryInput) (*ecr.DescribeRegistryOutput, error)
	DescribeRegistryWithContext(aws.Context, *ecr.DescribeRegistryInput, ...request.Option) (*ecr.DescribeRegistryOutput, error)
	DescribeRegistryRequest(*ecr.DescribeRegistryInput) (*request.Request, *ecr.DescribeRegistryOutput)

	DescribeRepositories(*ecr.DescribeRepositoriesInput) (*ecr.DescribeRepositoriesOutput, error)
	DescribeRepositoriesWithContext(aws.Context, *ecr.DescribeRepositoriesInput, ...request.Option) (*ecr.DescribeRepositoriesOutput, error)
	DescribeRepositoriesRequest(*ecr.DescribeRepositoriesInput) (*request.Request, *ecr.DescribeRepositoriesOutput)

	DescribeRepositoriesPages(*ecr.DescribeRepositoriesInput, func(*ecr.DescribeRepositoriesOutput, bool) bool) error
	DescribeRepositoriesPagesWithContext(aws.Context, *ecr.DescribeRepositoriesInput, func(*ecr.DescribeRepositoriesOutput, bool) bool, ...request.Option) error

	GetAuthorizationToken(*ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error)
	GetAuthorizationTokenWithContext(aws.Context, *ecr.GetAuthorizationTokenInput, ...request.Option) (*ecr.GetAuthorizationTokenOutput, error)
	GetAuthorizationTokenRequest(*ecr.GetAuthorizationTokenInput) (*request.Request, *ecr.GetAuthorizationTokenOutput)

	GetDownloadUrlForLayer(*ecr.GetDownloadUrlForLayerInput) (*ecr.GetDownloadUrlForLayerOutput, error)
	GetDownloadUrlForLayerWithContext(aws.Context, *ecr.GetDownloadUrlForLayerInput, ...request.Option) (*ecr.GetDownloadUrlForLayerOutput, error)
	GetDownloadUrlForLayerRequest(*ecr.GetDownloadUrlForLayerInput) (*request.Request, *ecr.GetDownloadUrlForLayerOutput)

	GetLifecyclePolicy(*ecr.GetLifecyclePolicyInput) (*ecr.GetLifecyclePolicyOutput, error)
	GetLifecyclePolicyWithContext(aws.Context, *ecr.GetLifecyclePolicyInput, ...request.Option) (*ecr.GetLifecyclePolicyOutput, error)
	GetLifecyclePolicyRequest(*ecr.GetLifecyclePolicyInput) (*request.Request, *ecr.GetLifecyclePolicyOutput)

	GetLifecyclePolicyPreview(*ecr.GetLifecyclePolicyPreviewInput) (*ecr.GetLifecyclePolicyPreviewOutput, error)
	GetLifecyclePolicyPreviewWithContext(aws.Context, *ecr.GetLifecyclePolicyPreviewInput, ...request.Option) (*ecr.GetLifecyclePolicyPreviewOutput, error)
	GetLifecyclePolicyPreviewRequest(*ecr.GetLifecyclePolicyPreviewInput) (*request.Request, *ecr.GetLifecyclePolicyPreviewOutput)

	GetLifecyclePolicyPreviewPages(*ecr.GetLifecyclePolicyPreviewInput, func(*ecr.GetLifecyclePolicyPreviewOutput, bool) bool) error
	GetLifecyclePolicyPreviewPagesWithContext(aws.Context, *ecr.GetLifecyclePolicyPreviewInput, func(*ecr.GetLifecyclePolicyPreviewOutput, bool) bool, ...request.Option) error

	GetRegistryPolicy(*ecr.GetRegistryPolicyInput) (*ecr.GetRegistryPolicyOutput, error)
	GetRegistryPolicyWithContext(aws.Context, *ecr.GetRegistryPolicyInput, ...request.Option) (*ecr.GetRegistryPolicyOutput, error)
	GetRegistryPolicyRequest(*ecr.GetRegistryPolicyInput) (*request.Request, *ecr.GetRegistryPolicyOutput)

	GetRegistryScanningConfiguration(*ecr.GetRegistryScanningConfigurationInput) (*ecr.GetRegistryScanningConfigurationOutput, error)
	GetRegistryScanningConfigurationWithContext(aws.Context, *ecr.GetRegistryScanningConfigurationInput, ...request.Option) (*ecr.GetRegistryScanningConfigurationOutput, error)
	GetRegistryScanningConfigurationRequest(*ecr.GetRegistryScanningConfigurationInput) (*request.Request, *ecr.GetRegistryScanningConfigurationOutput)

	GetRepositoryPolicy(*ecr.GetRepositoryPolicyInput) (*ecr.GetRepositoryPolicyOutput, error)
	GetRepositoryPolicyWithContext(aws.Context, *ecr.GetRepositoryPolicyInput, ...request.Option) (*ecr.GetRepositoryPolicyOutput, error)
	GetRepositoryPolicyRequest(*ecr.GetRepositoryPolicyInput) (*request.Request, *ecr.GetRepositoryPolicyOutput)

	InitiateLayerUpload(*ecr.InitiateLayerUploadInput) (*ecr.InitiateLayerUploadOutput, error)
	InitiateLayerUploadWithContext(aws.Context, *ecr.InitiateLayerUploadInput, ...request.Option) (*ecr.InitiateLayerUploadOutput, error)
	InitiateLayerUploadRequest(*ecr.InitiateLayerUploadInput) (*request.Request, *ecr.InitiateLayerUploadOutput)

	ListImages(*ecr.ListImagesInput) (*ecr.ListImagesOutput, error)
	ListImagesWithContext(aws.Context, *ecr.ListImagesInput, ...request.Option) (*ecr.ListImagesOutput, error)
	ListImagesRequest(*ecr.ListImagesInput) (*request.Request, *ecr.ListImagesOutput)

	ListImagesPages(*ecr.ListImagesInput, func(*ecr.ListImagesOutput, bool) bool) error
	ListImagesPagesWithContext(aws.Context, *ecr.ListImagesInput, func(*ecr.ListImagesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*ecr.ListTagsForResourceInput) (*ecr.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *ecr.ListTagsForResourceInput, ...request.Option) (*ecr.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*ecr.ListTagsForResourceInput) (*request.Request, *ecr.ListTagsForResourceOutput)

	PutImage(*ecr.PutImageInput) (*ecr.PutImageOutput, error)
	PutImageWithContext(aws.Context, *ecr.PutImageInput, ...request.Option) (*ecr.PutImageOutput, error)
	PutImageRequest(*ecr.PutImageInput) (*request.Request, *ecr.PutImageOutput)

	PutImageScanningConfiguration(*ecr.PutImageScanningConfigurationInput) (*ecr.PutImageScanningConfigurationOutput, error)
	PutImageScanningConfigurationWithContext(aws.Context, *ecr.PutImageScanningConfigurationInput, ...request.Option) (*ecr.PutImageScanningConfigurationOutput, error)
	PutImageScanningConfigurationRequest(*ecr.PutImageScanningConfigurationInput) (*request.Request, *ecr.PutImageScanningConfigurationOutput)

	PutImageTagMutability(*ecr.PutImageTagMutabilityInput) (*ecr.PutImageTagMutabilityOutput, error)
	PutImageTagMutabilityWithContext(aws.Context, *ecr.PutImageTagMutabilityInput, ...request.Option) (*ecr.PutImageTagMutabilityOutput, error)
	PutImageTagMutabilityRequest(*ecr.PutImageTagMutabilityInput) (*request.Request, *ecr.PutImageTagMutabilityOutput)

	PutLifecyclePolicy(*ecr.PutLifecyclePolicyInput) (*ecr.PutLifecyclePolicyOutput, error)
	PutLifecyclePolicyWithContext(aws.Context, *ecr.PutLifecyclePolicyInput, ...request.Option) (*ecr.PutLifecyclePolicyOutput, error)
	PutLifecyclePolicyRequest(*ecr.PutLifecyclePolicyInput) (*request.Request, *ecr.PutLifecyclePolicyOutput)

	PutRegistryPolicy(*ecr.PutRegistryPolicyInput) (*ecr.PutRegistryPolicyOutput, error)
	PutRegistryPolicyWithContext(aws.Context, *ecr.PutRegistryPolicyInput, ...request.Option) (*ecr.PutRegistryPolicyOutput, error)
	PutRegistryPolicyRequest(*ecr.PutRegistryPolicyInput) (*request.Request, *ecr.PutRegistryPolicyOutput)

	PutRegistryScanningConfiguration(*ecr.PutRegistryScanningConfigurationInput) (*ecr.PutRegistryScanningConfigurationOutput, error)
	PutRegistryScanningConfigurationWithContext(aws.Context, *ecr.PutRegistryScanningConfigurationInput, ...request.Option) (*ecr.PutRegistryScanningConfigurationOutput, error)
	PutRegistryScanningConfigurationRequest(*ecr.PutRegistryScanningConfigurationInput) (*request.Request, *ecr.PutRegistryScanningConfigurationOutput)

	PutReplicationConfiguration(*ecr.PutReplicationConfigurationInput) (*ecr.PutReplicationConfigurationOutput, error)
	PutReplicationConfigurationWithContext(aws.Context, *ecr.PutReplicationConfigurationInput, ...request.Option) (*ecr.PutReplicationConfigurationOutput, error)
	PutReplicationConfigurationRequest(*ecr.PutReplicationConfigurationInput) (*request.Request, *ecr.PutReplicationConfigurationOutput)

	SetRepositoryPolicy(*ecr.SetRepositoryPolicyInput) (*ecr.SetRepositoryPolicyOutput, error)
	SetRepositoryPolicyWithContext(aws.Context, *ecr.SetRepositoryPolicyInput, ...request.Option) (*ecr.SetRepositoryPolicyOutput, error)
	SetRepositoryPolicyRequest(*ecr.SetRepositoryPolicyInput) (*request.Request, *ecr.SetRepositoryPolicyOutput)

	StartImageScan(*ecr.StartImageScanInput) (*ecr.StartImageScanOutput, error)
	StartImageScanWithContext(aws.Context, *ecr.StartImageScanInput, ...request.Option) (*ecr.StartImageScanOutput, error)
	StartImageScanRequest(*ecr.StartImageScanInput) (*request.Request, *ecr.StartImageScanOutput)

	StartLifecyclePolicyPreview(*ecr.StartLifecyclePolicyPreviewInput) (*ecr.StartLifecyclePolicyPreviewOutput, error)
	StartLifecyclePolicyPreviewWithContext(aws.Context, *ecr.StartLifecyclePolicyPreviewInput, ...request.Option) (*ecr.StartLifecyclePolicyPreviewOutput, error)
	StartLifecyclePolicyPreviewRequest(*ecr.StartLifecyclePolicyPreviewInput) (*request.Request, *ecr.StartLifecyclePolicyPreviewOutput)

	TagResource(*ecr.TagResourceInput) (*ecr.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *ecr.TagResourceInput, ...request.Option) (*ecr.TagResourceOutput, error)
	TagResourceRequest(*ecr.TagResourceInput) (*request.Request, *ecr.TagResourceOutput)

	UntagResource(*ecr.UntagResourceInput) (*ecr.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *ecr.UntagResourceInput, ...request.Option) (*ecr.UntagResourceOutput, error)
	UntagResourceRequest(*ecr.UntagResourceInput) (*request.Request, *ecr.UntagResourceOutput)

	UploadLayerPart(*ecr.UploadLayerPartInput) (*ecr.UploadLayerPartOutput, error)
	UploadLayerPartWithContext(aws.Context, *ecr.UploadLayerPartInput, ...request.Option) (*ecr.UploadLayerPartOutput, error)
	UploadLayerPartRequest(*ecr.UploadLayerPartInput) (*request.Request, *ecr.UploadLayerPartOutput)

	WaitUntilImageScanComplete(*ecr.DescribeImageScanFindingsInput) error
	WaitUntilImageScanCompleteWithContext(aws.Context, *ecr.DescribeImageScanFindingsInput, ...request.WaiterOption) error

	WaitUntilLifecyclePolicyPreviewComplete(*ecr.GetLifecyclePolicyPreviewInput) error
	WaitUntilLifecyclePolicyPreviewCompleteWithContext(aws.Context, *ecr.GetLifecyclePolicyPreviewInput, ...request.WaiterOption) error
}

var _ ECRAPI = (*ecr.ECR)(nil)
