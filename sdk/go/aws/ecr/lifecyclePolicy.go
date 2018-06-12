// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecr

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an ECR lifecycle policy.
type LifecyclePolicy struct {
	s *pulumi.ResourceState
}

// NewLifecyclePolicy registers a new resource with the given unique name, arguments, and options.
func NewLifecyclePolicy(ctx *pulumi.Context,
	name string, args *LifecyclePolicyArgs, opts ...pulumi.ResourceOpt) (*LifecyclePolicy, error) {
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	if args == nil || args.Repository == nil {
		return nil, errors.New("missing required argument 'Repository'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["policy"] = nil
		inputs["repository"] = nil
	} else {
		inputs["policy"] = args.Policy
		inputs["repository"] = args.Repository
	}
	inputs["registryId"] = nil
	s, err := ctx.RegisterResource("aws:ecr/lifecyclePolicy:LifecyclePolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LifecyclePolicy{s: s}, nil
}

// GetLifecyclePolicy gets an existing LifecyclePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLifecyclePolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LifecyclePolicyState, opts ...pulumi.ResourceOpt) (*LifecyclePolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["policy"] = state.Policy
		inputs["registryId"] = state.RegistryId
		inputs["repository"] = state.Repository
	}
	s, err := ctx.ReadResource("aws:ecr/lifecyclePolicy:LifecyclePolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LifecyclePolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *LifecyclePolicy) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *LifecyclePolicy) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The policy document. This is a JSON formatted string. See more details about [Policy Parameters](http://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html#lifecycle_policy_parameters) in the official AWS docs.
func (r *LifecyclePolicy) Policy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policy"])
}

// The registry ID where the repository was created.
func (r *LifecyclePolicy) RegistryId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["registryId"])
}

// Name of the repository to apply the policy.
func (r *LifecyclePolicy) Repository() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["repository"])
}

// Input properties used for looking up and filtering LifecyclePolicy resources.
type LifecyclePolicyState struct {
	// The policy document. This is a JSON formatted string. See more details about [Policy Parameters](http://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html#lifecycle_policy_parameters) in the official AWS docs.
	Policy interface{}
	// The registry ID where the repository was created.
	RegistryId interface{}
	// Name of the repository to apply the policy.
	Repository interface{}
}

// The set of arguments for constructing a LifecyclePolicy resource.
type LifecyclePolicyArgs struct {
	// The policy document. This is a JSON formatted string. See more details about [Policy Parameters](http://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html#lifecycle_policy_parameters) in the official AWS docs.
	Policy interface{}
	// Name of the repository to apply the policy.
	Repository interface{}
}