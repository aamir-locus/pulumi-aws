# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetRegexPatternSetResult:
    """
    A collection of values returned by getRegexPatternSet.
    """
    def __init__(__self__, arn=None, description=None, id=None, name=None, regular_expressions=None, scope=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The Amazon Resource Name (ARN) of the entity.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        The description of the set that helps with identification.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if regular_expressions and not isinstance(regular_expressions, list):
            raise TypeError("Expected argument 'regular_expressions' to be a list")
        __self__.regular_expressions = regular_expressions
        """
        One or more blocks of regular expression patterns that AWS WAF is searching for. See Regular Expression below for details.
        """
        if scope and not isinstance(scope, str):
            raise TypeError("Expected argument 'scope' to be a str")
        __self__.scope = scope
class AwaitableGetRegexPatternSetResult(GetRegexPatternSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegexPatternSetResult(
            arn=self.arn,
            description=self.description,
            id=self.id,
            name=self.name,
            regular_expressions=self.regular_expressions,
            scope=self.scope)

def get_regex_pattern_set(name=None,scope=None,opts=None):
    """
    Retrieves the summary of a WAFv2 Regex Pattern Set.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.wafv2.get_regex_pattern_set(name="some-regex-pattern-set",
        scope="REGIONAL")
    ```


    :param str name: The name of the WAFv2 Regex Pattern Set.
    :param str scope: Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['scope'] = scope
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:wafv2/getRegexPatternSet:getRegexPatternSet', __args__, opts=opts).value

    return AwaitableGetRegexPatternSetResult(
        arn=__ret__.get('arn'),
        description=__ret__.get('description'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        regular_expressions=__ret__.get('regularExpressions'),
        scope=__ret__.get('scope'))