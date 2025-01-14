// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Fms.Outputs
{

    [OutputType]
    public sealed class PolicySecurityServicePolicyData
    {
        /// <summary>
        /// Details about the service that are specific to the service type, in JSON format. For service type `SHIELD_ADVANCED`, this is an empty string. Examples depending on `type` can be found in the [AWS Firewall Manager SecurityServicePolicyData API Reference](https://docs.aws.amazon.com/fms/2018-01-01/APIReference/API_SecurityServicePolicyData.html).
        /// </summary>
        public readonly string? ManagedServiceData;
        /// <summary>
        /// The service that the policy is using to protect the resources. Valid values are `WAFV2`, `WAF`, `SHIELD_ADVANCED`, `SECURITY_GROUPS_COMMON`, `SECURITY_GROUPS_CONTENT_AUDIT`, and `SECURITY_GROUPS_USAGE_AUDIT`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private PolicySecurityServicePolicyData(
            string? managedServiceData,

            string type)
        {
            ManagedServiceData = managedServiceData;
            Type = type;
        }
    }
}
