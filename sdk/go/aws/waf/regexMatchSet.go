// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a WAF Regex Match Set Resource
type RegexMatchSet struct {
	s *pulumi.ResourceState
}

// NewRegexMatchSet registers a new resource with the given unique name, arguments, and options.
func NewRegexMatchSet(ctx *pulumi.Context,
	name string, args *RegexMatchSetArgs, opts ...pulumi.ResourceOpt) (*RegexMatchSet, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["regexMatchTuples"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["regexMatchTuples"] = args.RegexMatchTuples
	}
	s, err := ctx.RegisterResource("aws:waf/regexMatchSet:RegexMatchSet", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegexMatchSet{s: s}, nil
}

// GetRegexMatchSet gets an existing RegexMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegexMatchSet(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RegexMatchSetState, opts ...pulumi.ResourceOpt) (*RegexMatchSet, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
		inputs["regexMatchTuples"] = state.RegexMatchTuples
	}
	s, err := ctx.ReadResource("aws:waf/regexMatchSet:RegexMatchSet", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegexMatchSet{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RegexMatchSet) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RegexMatchSet) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The name or description of the Regex Match Set.
func (r *RegexMatchSet) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The regular expression pattern that you want AWS WAF to search for in web requests,
// the location in requests that you want AWS WAF to search, and other settings. See below.
func (r *RegexMatchSet) RegexMatchTuples() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["regexMatchTuples"])
}

// Input properties used for looking up and filtering RegexMatchSet resources.
type RegexMatchSetState struct {
	// The name or description of the Regex Match Set.
	Name interface{}
	// The regular expression pattern that you want AWS WAF to search for in web requests,
	// the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples interface{}
}

// The set of arguments for constructing a RegexMatchSet resource.
type RegexMatchSetArgs struct {
	// The name or description of the Regex Match Set.
	Name interface{}
	// The regular expression pattern that you want AWS WAF to search for in web requests,
	// the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples interface{}
}