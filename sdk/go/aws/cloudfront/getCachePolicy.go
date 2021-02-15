// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Example Usage
//
// The following example below creates a CloudFront cache policy.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudfront"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "example-policy"
// 		_, err := cloudfront.LookupCachePolicy(ctx, &cloudfront.LookupCachePolicyArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupCachePolicy(ctx *pulumi.Context, args *LookupCachePolicyArgs, opts ...pulumi.InvokeOption) (*LookupCachePolicyResult, error) {
	var rv LookupCachePolicyResult
	err := ctx.Invoke("aws:cloudfront/getCachePolicy:getCachePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCachePolicy.
type LookupCachePolicyArgs struct {
	// The identifier for the cache policy.
	Id *string `pulumi:"id"`
	// A unique name to identify the cache policy.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getCachePolicy.
type LookupCachePolicyResult struct {
	// A comment to describe the cache policy.
	Comment string `pulumi:"comment"`
	// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	DefaultTtl int `pulumi:"defaultTtl"`
	// The current version of the cache policy.
	Etag string  `pulumi:"etag"`
	Id   *string `pulumi:"id"`
	// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MaxTtl int `pulumi:"maxTtl"`
	// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MinTtl int     `pulumi:"minTtl"`
	Name   *string `pulumi:"name"`
	// The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
	ParametersInCacheKeyAndForwardedToOrigins []GetCachePolicyParametersInCacheKeyAndForwardedToOrigin `pulumi:"parametersInCacheKeyAndForwardedToOrigins"`
}