// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CloudWatch Log Metric Filter resource.
type LogMetricFilter struct {
	s *pulumi.ResourceState
}

// NewLogMetricFilter registers a new resource with the given unique name, arguments, and options.
func NewLogMetricFilter(ctx *pulumi.Context,
	name string, args *LogMetricFilterArgs, opts ...pulumi.ResourceOpt) (*LogMetricFilter, error) {
	if args == nil || args.LogGroupName == nil {
		return nil, errors.New("missing required argument 'LogGroupName'")
	}
	if args == nil || args.MetricTransformation == nil {
		return nil, errors.New("missing required argument 'MetricTransformation'")
	}
	if args == nil || args.Pattern == nil {
		return nil, errors.New("missing required argument 'Pattern'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["logGroupName"] = nil
		inputs["metricTransformation"] = nil
		inputs["name"] = nil
		inputs["pattern"] = nil
	} else {
		inputs["logGroupName"] = args.LogGroupName
		inputs["metricTransformation"] = args.MetricTransformation
		inputs["name"] = args.Name
		inputs["pattern"] = args.Pattern
	}
	s, err := ctx.RegisterResource("aws:cloudwatch/logMetricFilter:LogMetricFilter", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LogMetricFilter{s: s}, nil
}

// GetLogMetricFilter gets an existing LogMetricFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogMetricFilter(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LogMetricFilterState, opts ...pulumi.ResourceOpt) (*LogMetricFilter, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["logGroupName"] = state.LogGroupName
		inputs["metricTransformation"] = state.MetricTransformation
		inputs["name"] = state.Name
		inputs["pattern"] = state.Pattern
	}
	s, err := ctx.ReadResource("aws:cloudwatch/logMetricFilter:LogMetricFilter", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LogMetricFilter{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *LogMetricFilter) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *LogMetricFilter) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The name of the log group to associate the metric filter with.
func (r *LogMetricFilter) LogGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["logGroupName"])
}

// A block defining collection of information
// needed to define how metric data gets emitted. See below.
func (r *LogMetricFilter) MetricTransformation() *pulumi.Output {
	return r.s.State["metricTransformation"]
}

// The name of the CloudWatch metric to which the monitored log information should be published (e.g. `ErrorCount`)
func (r *LogMetricFilter) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
// for extracting metric data out of ingested log events.
func (r *LogMetricFilter) Pattern() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["pattern"])
}

// Input properties used for looking up and filtering LogMetricFilter resources.
type LogMetricFilterState struct {
	// The name of the log group to associate the metric filter with.
	LogGroupName interface{}
	// A block defining collection of information
	// needed to define how metric data gets emitted. See below.
	MetricTransformation interface{}
	// The name of the CloudWatch metric to which the monitored log information should be published (e.g. `ErrorCount`)
	Name interface{}
	// A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
	// for extracting metric data out of ingested log events.
	Pattern interface{}
}

// The set of arguments for constructing a LogMetricFilter resource.
type LogMetricFilterArgs struct {
	// The name of the log group to associate the metric filter with.
	LogGroupName interface{}
	// A block defining collection of information
	// needed to define how metric data gets emitted. See below.
	MetricTransformation interface{}
	// The name of the CloudWatch metric to which the monitored log information should be published (e.g. `ErrorCount`)
	Name interface{}
	// A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
	// for extracting metric data out of ingested log events.
	Pattern interface{}
}