# whitelist-domain-public-ip-aws
AWS Lambda for whitelisting the public IP Ranges of any external domain in AWS Security Group Outbound rules.

## Build Lambda

**Linux**
```
# Get dependency
go get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
```

```
# Compile and zip
GOOS=linux go build lambda.go && zip go_lambda.zip lambda
```

**Windows**
```
# Get dependency
go get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
```

```
# Compile and zip
$env:GOOS = "linux"
$env:CGO_ENABLED = "0"
$env:GOARCH = "amd64"
go build -o lambda lambda.go; ~\Go\Bin\build-lambda-zip.exe -output go_lambda.zip lambda
```

### Lambda Settings

**Environment variables**
|  Key | Value  | Description |
|---   |---     |---          |
| awsRegion  | eu-central-1 | AWS Region |
| securityGroupIDs  | sg-00ffabccebd5efda2,sg-041c5e7daf95e16a3 | Comma separated list of Security groups (no spaces) |
| domainNames  |  hub.docker.com,helm.nginx.com | Comma separated list of Domain Names (no spaces) |

**Change Handler entrypoint**
Change Handler from the default "hello" to "lambda"

**Lambda Execution role**

Create standard Lambda execution role and add ec2:AuthorizeSecurityGroupEgress write permissions to it.

Example policy for "AuthorizeSecurityGroupEgress" access : 
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": "ec2:AuthorizeSecurityGroupEgress",
            "Resource": "arn:aws:ec2:*:722377226063:security-group/*"
        }
    ]
}
```

**Lambda trigger**
EventBridge time-based trigger example - scheduled expressions

Add Trigger > EventBridge > Create Rule > whitelist-external-ip-lambda > Scheduled expression
```
# Run every hour
cron(0 * * * ? *)

# Run every 30 min
cron(0/30 * * * ? *)
```


**Local test vars**
```
# In Go Code
domainNames = []string{"hub.docker.com", "helm.nginx.com"}
securityGroupIDs = []string{"sg-00ffabccebd5efda2", "sg-041c5e7daf95e16a3"}
awsRegion = "eu-central-1"

# Env vars - Linux
export domainNames = "hub.docker.com,helm.nginx.com"
export securityGroupIDs = "sg-00ffabccebd5efda2,sg-041c5e7daf95e16a3"
export awsRegion = "eu-central-1"

# Env vars - Windows
$env:domainNames = "hub.docker.com,helm.nginx.com"
$env:securityGroupIDs = "sg-00ffabccebd5efda2,sg-041c5e7daf95e16a3"
$env:awsRegion = "eu-central-1"
```

#### Errors : 
You haven't changed Handler entrypoint in Lambda settings:
```
{
  "errorMessage": "fork/exec /var/task/hello: no such file or directory",
  "errorType": "PathError"
}
```

Missing IAM permissions to change security group egress:
```
{
  "errorMessage": "updateEgressErr: awsUpdateSg: UnauthorizedOperation: You are not authorized to perform this operation. Encoded authorization failure message: ...\n\tstatus code: 403, request id: 2795cff9-f02f-498b-ac23-616be92e2676",
  "errorType": "errorString"
}
```

Maximum amount of rules per security group exceeed: 
```
{
  "errorMessage": "updateEgressErr: awsUpdateSg: RulesPerSecurityGroupLimitExceeded: The maximum number of rules per security group has been reached.\n\tstatus code: 400, request id: ed01a938-5270-49bc-b6dd-2712ce224383",
  "errorType": "errorString"
}
```
AWS has a limit of 60 inbound/outbound rules per Security Group - https://docs.amazonaws.cn/en_us/vpc/latest/userguide/amazon-vpc-limits.html (you can request AWS to increase this limit)
