// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The ECS container definition data source allows access to details of
// a specific container within an AWS ECS service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ecs.GetContainerDefinition(ctx, &ecs.GetContainerDefinitionArgs{
// 			TaskDefinition: aws_ecs_task_definition.Mongo.Id,
// 			ContainerName:  "mongodb",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetContainerDefinition(ctx *pulumi.Context, args *GetContainerDefinitionArgs, opts ...pulumi.InvokeOption) (*GetContainerDefinitionResult, error) {
	var rv GetContainerDefinitionResult
	err := ctx.Invoke("aws:ecs/getContainerDefinition:getContainerDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainerDefinition.
type GetContainerDefinitionArgs struct {
	// The name of the container definition
	ContainerName string `pulumi:"containerName"`
	// The ARN of the task definition which contains the container
	TaskDefinition string `pulumi:"taskDefinition"`
}

// A collection of values returned by getContainerDefinition.
type GetContainerDefinitionResult struct {
	ContainerName string `pulumi:"containerName"`
	// The CPU limit for this container definition
	Cpu int `pulumi:"cpu"`
	// Indicator if networking is disabled
	DisableNetworking bool `pulumi:"disableNetworking"`
	// Set docker labels
	DockerLabels map[string]string `pulumi:"dockerLabels"`
	// The environment in use
	Environment map[string]string `pulumi:"environment"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The docker image in use, including the digest
	Image string `pulumi:"image"`
	// The digest of the docker image in use
	ImageDigest string `pulumi:"imageDigest"`
	// The memory limit for this container definition
	Memory int `pulumi:"memory"`
	// The soft limit (in MiB) of memory to reserve for the container. When system memory is under contention, Docker attempts to keep the container memory to this soft limit
	MemoryReservation int    `pulumi:"memoryReservation"`
	TaskDefinition    string `pulumi:"taskDefinition"`
}
