// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iam
{
    public static class GetPolicyDocument
    {
        /// <summary>
        /// Generates an IAM policy document in JSON format for use with resources that expect policy documents such as `aws.iam.Policy`.
        /// 
        /// Using this data source to generate policy documents is *optional*. It is also valid to use literal JSON strings in your configuration or to use the `file` interpolation function to read a raw JSON policy document from a file.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Example with Both Source and Override Documents
        /// 
        /// You can also combine `source_json` and `override_json` in the same document.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var source = Output.Create(Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
        ///         {
        ///             Statements = 
        ///             {
        ///                 new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
        ///                 {
        ///                     Sid = "OverridePlaceholder",
        ///                     Actions = 
        ///                     {
        ///                         "ec2:DescribeAccountAttributes",
        ///                     },
        ///                     Resources = 
        ///                     {
        ///                         "*",
        ///                     },
        ///                 },
        ///             },
        ///         }));
        ///         var @override = Output.Create(Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
        ///         {
        ///             Statements = 
        ///             {
        ///                 new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
        ///                 {
        ///                     Sid = "OverridePlaceholder",
        ///                     Actions = 
        ///                     {
        ///                         "s3:GetObject",
        ///                     },
        ///                     Resources = 
        ///                     {
        ///                         "*",
        ///                     },
        ///                 },
        ///             },
        ///         }));
        ///         var politik = Output.Tuple(source, @override).Apply(values =&gt;
        ///         {
        ///             var source = values.Item1;
        ///             var @override = values.Item2;
        ///             return Output.Create(Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
        ///             {
        ///                 SourceJson = source.Json,
        ///                 OverrideJson = @override.Json,
        ///             }));
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// `data.aws_iam_policy_document.politik.json` will evaluate to:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPolicyDocumentResult> InvokeAsync(GetPolicyDocumentArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyDocumentResult>("aws:iam/getPolicyDocument:getPolicyDocument", args ?? new GetPolicyDocumentArgs(), options.WithVersion());
    }


    public sealed class GetPolicyDocumentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// IAM policy document whose statements with non-blank `sid`s will override statements with the same `sid` from documents assigned to the `source_json`, `source_policy_documents`, and `override_policy_documents` arguments. Non-overriding statements will be added to the exported document.
        /// </summary>
        [Input("overrideJson")]
        public string? OverrideJson { get; set; }

        [Input("overridePolicyDocuments")]
        private List<string>? _overridePolicyDocuments;

        /// <summary>
        /// List of IAM policy documents that are merged together into the exported document. In merging, statements with non-blank `sid`s will override statements with the same `sid` from earlier documents in the list. Statements with non-blank `sid`s will also override statements with the same `sid` from documents provided in the `source_json` and `source_policy_documents` arguments.  Non-overriding statements will be added to the exported document.
        /// </summary>
        public List<string> OverridePolicyDocuments
        {
            get => _overridePolicyDocuments ?? (_overridePolicyDocuments = new List<string>());
            set => _overridePolicyDocuments = value;
        }

        /// <summary>
        /// ID for the policy document.
        /// </summary>
        [Input("policyId")]
        public string? PolicyId { get; set; }

        /// <summary>
        /// IAM policy document used as a base for the exported policy document. Statements with the same `sid` from documents assigned to the `override_json` and `override_policy_documents` arguments will override source statements.
        /// </summary>
        [Input("sourceJson")]
        public string? SourceJson { get; set; }

        [Input("sourcePolicyDocuments")]
        private List<string>? _sourcePolicyDocuments;

        /// <summary>
        /// List of IAM policy documents that are merged together into the exported document. Statements defined in `source_policy_documents` or `source_json` must have unique `sid`s. Statements with the same `sid` from documents assigned to the `override_json` and `override_policy_documents` arguments will override source statements.
        /// </summary>
        public List<string> SourcePolicyDocuments
        {
            get => _sourcePolicyDocuments ?? (_sourcePolicyDocuments = new List<string>());
            set => _sourcePolicyDocuments = value;
        }

        [Input("statements")]
        private List<Inputs.GetPolicyDocumentStatementArgs>? _statements;

        /// <summary>
        /// Configuration block for a policy statement. Detailed below.
        /// </summary>
        public List<Inputs.GetPolicyDocumentStatementArgs> Statements
        {
            get => _statements ?? (_statements = new List<Inputs.GetPolicyDocumentStatementArgs>());
            set => _statements = value;
        }

        /// <summary>
        /// IAM policy document version. Valid values are `2008-10-17` and `2012-10-17`. Defaults to `2012-10-17`. For more information, see the [AWS IAM User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_version.html).
        /// </summary>
        [Input("version")]
        public string? Version { get; set; }

        public GetPolicyDocumentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicyDocumentResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Standard JSON policy document rendered based on the arguments above.
        /// </summary>
        public readonly string Json;
        public readonly string? OverrideJson;
        public readonly ImmutableArray<string> OverridePolicyDocuments;
        public readonly string? PolicyId;
        public readonly string? SourceJson;
        public readonly ImmutableArray<string> SourcePolicyDocuments;
        public readonly ImmutableArray<Outputs.GetPolicyDocumentStatementResult> Statements;
        public readonly string? Version;

        [OutputConstructor]
        private GetPolicyDocumentResult(
            string id,

            string json,

            string? overrideJson,

            ImmutableArray<string> overridePolicyDocuments,

            string? policyId,

            string? sourceJson,

            ImmutableArray<string> sourcePolicyDocuments,

            ImmutableArray<Outputs.GetPolicyDocumentStatementResult> statements,

            string? version)
        {
            Id = id;
            Json = json;
            OverrideJson = overrideJson;
            OverridePolicyDocuments = overridePolicyDocuments;
            PolicyId = policyId;
            SourceJson = sourceJson;
            SourcePolicyDocuments = sourcePolicyDocuments;
            Statements = statements;
            Version = version;
        }
    }
}
