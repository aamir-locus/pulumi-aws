// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a WAF Byte Match Set Resource
type ByteMatchSet struct {
	s *pulumi.ResourceState
}

// NewByteMatchSet registers a new resource with the given unique name, arguments, and options.
func NewByteMatchSet(ctx *pulumi.Context,
	name string, args *ByteMatchSetArgs, opts ...pulumi.ResourceOpt) (*ByteMatchSet, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["byteMatchTuples"] = nil
		inputs["name"] = nil
	} else {
		inputs["byteMatchTuples"] = args.ByteMatchTuples
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("aws:waf/byteMatchSet:ByteMatchSet", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ByteMatchSet{s: s}, nil
}

// GetByteMatchSet gets an existing ByteMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetByteMatchSet(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ByteMatchSetState, opts ...pulumi.ResourceOpt) (*ByteMatchSet, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["byteMatchTuples"] = state.ByteMatchTuples
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("aws:waf/byteMatchSet:ByteMatchSet", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ByteMatchSet{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ByteMatchSet) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ByteMatchSet) ID() *pulumi.IDOutput {
	return r.s.ID
}

// Specifies the bytes (typically a string that corresponds
// with ASCII characters) that you want to search for in web requests,
// the location in requests that you want to search, and other settings.
func (r *ByteMatchSet) ByteMatchTuples() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["byteMatchTuples"])
}

// The name or description of the Byte Match Set.
func (r *ByteMatchSet) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering ByteMatchSet resources.
type ByteMatchSetState struct {
	// Specifies the bytes (typically a string that corresponds
	// with ASCII characters) that you want to search for in web requests,
	// the location in requests that you want to search, and other settings.
	ByteMatchTuples interface{}
	// The name or description of the Byte Match Set.
	Name interface{}
}

// The set of arguments for constructing a ByteMatchSet resource.
type ByteMatchSetArgs struct {
	// Specifies the bytes (typically a string that corresponds
	// with ASCII characters) that you want to search for in web requests,
	// the location in requests that you want to search, and other settings.
	ByteMatchTuples interface{}
	// The name or description of the Byte Match Set.
	Name interface{}
}