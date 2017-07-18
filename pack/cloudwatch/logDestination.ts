// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class LogDestination extends lumi.NamedResource implements LogDestinationArgs {
    public /*out*/ readonly arn: string;
    public readonly logDestinationName?: string;
    public readonly roleArn: string;
    public readonly targetArn: string;

    constructor(name: string, args: LogDestinationArgs) {
        super(name);
        this.logDestinationName = args.logDestinationName;
        if (lumirt.defaultIfComputed(args.roleArn, "") === undefined) {
            throw new Error("Property argument 'roleArn' is required, but was missing");
        }
        this.roleArn = args.roleArn;
        if (lumirt.defaultIfComputed(args.targetArn, "") === undefined) {
            throw new Error("Property argument 'targetArn' is required, but was missing");
        }
        this.targetArn = args.targetArn;
    }
}

export interface LogDestinationArgs {
    readonly logDestinationName?: string;
    readonly roleArn: string;
    readonly targetArn: string;
}
