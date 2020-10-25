# whitelist-external-public-ip-aws
AWS Lambda for whitelisting the public IP Ranges of any external domain in AWS Security Group Outbound rules.

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
