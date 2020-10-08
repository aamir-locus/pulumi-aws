// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue.Outputs
{

    [OutputType]
    public sealed class MLTransformSchema
    {
        /// <summary>
        /// The type of data in the column.
        /// </summary>
        public readonly string? DataType;
        /// <summary>
        /// The name you assign to this ML Transform. It must be unique in your account.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private MLTransformSchema(
            string? dataType,

            string? name)
        {
            DataType = dataType;
            Name = name;
        }
    }
}