| acm                         | ACM                          | ✅ (Domain Name)                       | ✅ (Created Time)                    | ❌    |
| acmpca                      | ACMPCA                       | ❌                                     | ✅ (LastStateChange or Created Time) | ❌    |
| ami                         | AMI                          | ✅ (Image Name)                        | ✅ (Creation Time)                   | ❌    |
| apigateway                  | APIGateway                   | ✅ (API Name)                          | ✅ (Created Time)                    | ❌    |
| apigatewayv2                | APIGatewayV2                 | ✅ (API Name)                          | ✅ (Created Time)                    | ❌    |
| accessanalyzer              | AccessAnalyzer               | ✅ (Analyzer Name)                     | ✅ (Created Time)                    | ❌    |
| backup-vault                | BackupVault                  | ✅ (Backup Vault Name)                 | ✅ (Created Time)                    | ❌    |
| cloudwatch-alarm            | CloudWatchAlarm              | ✅ (Alarm Name)                        | ✅ (AlarmConfigurationUpdated Time)  | ❌    |
| cloudwatch-dashboard        | CloudWatchDashboard          | ✅ (Dashboard Name)                    | ✅ (LastModified Time)               | ❌    |
| cloudwatch-loggroup         | CloudWatchLogGroup           | ✅ (Log Group Name)                    | ✅ (Creation Time)                   | ❌    |
| cloudtrail                  | CloudtrailTrail              | ✅ (Trail Name)                        | ❌                                   | ❌    |
| codedeploy-application      | CodeDeployApplications       | ✅ (Application Name)                  | ✅ (Creation Time)                   | ❌    |
| config-recorders            | ConfigServiceRecorder        | ✅ (Recorder Name)                     | ❌                                   | ❌    |
| config-rules                | ConfigServiceRule            | ✅ (Rule Name)                         | ❌                                   | ❌    |
| rds-subnet-group            | DBSubnetGroups               | ✅ (DB Subnet Group Name)              | ❌                                   | ❌    |
| dynamodb                    | DynamoDB                     | ✅ (Table Name)                        | ✅ (Creation Time)                   | ❌    |
| ec2-dedicated-hosts         | EC2DedicatedHosts            | ✅ (EC2 Name Tag)                      | ✅ (Allocation Time)                 | ❌    |
| ec2-dhcp-option             | EC2DhcpOption                | ❌                                     | ❌                                   | ❌    |
| ecr                         | ECRRepository                | ✅ (Repository Name)                   | ✅ (Creation Time)                   | ❌    |
| ecscluster                  | ECSCluster                   | ✅ (Cluster Name)                      | ❌                                   | ❌    |
| ecsserv                     | ECSService                   | ✅ (Service Name)                      | ✅ (Creation Time)                   | ❌    |
| elb                         | ELBv1                        | ✅ (Load Balancer Name)                | ✅ (Created Time)                    | ❌    |
| elbv2                       | ELBv2                        | ✅ (Load Balancer Name)                | ✅ (Created Time)                    | ❌    |
| efs                         | ElasticFileSystem            | ✅ (File System Name)                  | ✅ (Creation Time)                   | ❌    |
| elasticache                 | Elasticache                  | ✅ (Cluster ID & Replication Group ID) | ✅ (Creation Time)                   | ❌    |
| elasticacheparametergroups  | ElasticacheParameterGroups   | ✅ (Parameter Group Name)              | ❌                                   | ❌    |
| elasticachesubnetgroups     | ElasticacheSubnetGroups      | ✅ (Subnet Group Name)                 | ❌                                   | ❌    |
| guardduty                   | GuardDuty                    | ❌                                     | ✅ (Created Time)                    | ❌    |
| iam-group                   | IAMGroups                    | ✅ (Group Name)                        | ✅ (Creation Time)                   | ❌    |
| iam-role                    | IAMRoles                     | ✅ (Role Name)                         | ✅ (Creation Time)                   | ❌    |
| iam-service-linked-role     | IAMServiceLinkedRoles        | ✅ (Service Linked Role Name)          | ✅ (Creation Time)                   | ❌    |
| kmscustomerkeys             | KMSCustomerKeys              | ✅ (Key Name)                          | ✅ (Creation Time)                   | ❌    |
| kinesis-stream              | KinesisStream                | ✅ (Stream Name)                       | ❌                                   | ❌    |
| lambda                      | LambdaFunction               | ✅ (Function Name)                     | ✅ (Last Modified Time)              | ❌    |
| lc                          | LaunchConfiguration          | ✅ (Launch Configuration Name)         | ✅ (Created Time)                    | ❌    |
| lt                          | LaunchTemplate               | ✅ (Launch Template Name)              | ✅ (Created Time)                    | ❌    |
| macie-member                | MacieMember                  | ❌                                     | ✅ (Creation Time)                   | ❌    |
| msk-cluster                 | MskCluster                   | ✅ (Cluster Name)                      | ✅ (Creation Time)                   | ❌    |
| oidcprovider                | OIDCProvider                 | ✅ (Provider URL)                      | ✅ (Creation Time)                   | ❌    |
| opensearchdomain            | OpenSearchDomain             | ✅ (Domain Name)                       | ✅ (First Seen Tag Time)             | ❌    |
| redshift                    | Redshift                     | ✅ (Cluster Identifier)                | ✅ (Creation Time)                   | ❌    |
| snstopic                    | SNS                          | ✅ (Topic Name)                        | ✅ (First Seen Tag Time)             | ❌    |
| sqs                         | SQS                          | ✅ (Queue Name)                        | ✅ (Creation Time)                   | ❌    |
| sagemaker-notebook-smni     | SageMakerNotebook            | ✅ (Notebook Instnace Name)            | ✅ (Creation Time)                   | ❌    |
| secretsmanager              | SecretsManagerSecrets        | ✅ (Secret Name)                       | ✅ (Last Accessed or Creation Time)  | ❌    |
| security-hub                | SecurityHub                  | ❌                                     | ✅ (Created Time)                    | ❌    |
| snap                        | Snapshots                    | ❌                                     | ✅ (Creation Time)                   | ✅    |
| transit-gateway             | TransitGateway               | ❌                                     | ✅ (Creation Time)                   | ❌    |
| transit-gateway-route-table | TransitGatewayRouteTable     | ❌                                     | ✅ (Creation Time)                   | ❌    |
| transit-gateway-attachment  | TransitGatewaysVpcAttachment | ❌                                     | ✅ (Creation Time)                   | ❌    |
| vpc                         | VPC                          | ✅ (EC2 Name Tag)                      | ✅ (First Seen Tag Time)             | ❌    |
