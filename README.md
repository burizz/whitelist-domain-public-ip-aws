# whitelist-external-public-ip-aws
AWS Lambda for whitelisting the public IP Ranges of any external domain in AWS Security Group Outbound rules.

## Build Lambda

Build on Linux : 
```
# Get dependency
go get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
```

```
# Compile and zip
GOOS=linux go build lambda.go && zip go_lambda.zip lambda
```

Build on Windows : 
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

#### Vars

Lambda Env Vars :
|  Key | Value  | Description |
|---   |---     |---          |
| awsRegion  | eu-central-1 | AWS Region |
| securityGroupIDs  | sg-041c5e7daf95e16a3,sg-041c5e7daf95e16a3 | Comma separated list of Security groups (no spaces) |
| domainNames  |  hub.docker.com,helm.nginx.com | Comma separated list of Domain Names (no spaces) |

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

#### IAM Permissions
Created standard Lambda execution role and add ec2:AuthorizeSecurityGroupEgress write permissions.

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