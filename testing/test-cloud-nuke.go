package main

import (
	"context"
	"fmt"
	"os"
	"time"

	nuke_aws "github.com/andrewderr/cloud-nuke-a1/aws"
	"github.com/andrewderr/cloud-nuke-a1/config"
	"github.com/andrewderr/cloud-nuke-a1/telemetry"
)

func main() {
	fmt.Print("Starting cloud-nuke-a1...\n")
	// filterTag := "northstar-testing"
	// var filterRule config.FilterRule = config.FilterRule{
	// 	Tag: &filterTag,
	// }

	// var resourceType config.ResourceType = config.ResourceType{
	// 	IncludeRule: filterRule,
	// }

	// var kmsResourceType config.KMSCustomerKeyResourceType = config.KMSCustomerKeyResourceType{
	// 	ResourceType: resourceType,
	// }
	// fmt.Print("2...\n")

	// var customConfig config.Config = config.Config{
	// 	ACM:                             resourceType,
	// 	ACMPCA:                          resourceType,
	// 	AMI:                             resourceType,
	// 	APIGateway:                      resourceType,
	// 	APIGatewayV2:                    resourceType,
	// 	AccessAnalyzer:                  resourceType,
	// 	AutoScalingGroup:                resourceType,
	// 	BackupVault:                     resourceType,
	// 	CloudWatchAlarm:                 resourceType,
	// 	CloudWatchDashboard:             resourceType,
	// 	CloudWatchLogGroup:              resourceType,
	// 	CloudtrailTrail:                 resourceType,
	// 	CodeDeployApplications:          resourceType,
	// 	ConfigServiceRecorder:           resourceType,
	// 	ConfigServiceRule:               resourceType,
	// 	DBClusters:                      resourceType,
	// 	DBInstances:                     resourceType,
	// 	DBSubnetGroups:                  resourceType,
	// 	DynamoDB:                        resourceType,
	// 	EBSVolume:                       resourceType,
	// 	EC2:                             resourceType,
	// 	EC2DedicatedHosts:               resourceType,
	// 	EC2DHCPOption:                   resourceType,
	// 	EC2KeyPairs:                     resourceType,
	// 	ECRRepository:                   resourceType,
	// 	ECSCluster:                      resourceType,
	// 	ECSService:                      resourceType,
	// 	EKSCluster:                      resourceType,
	// 	ELBv1:                           resourceType,
	// 	ELBv2:                           resourceType,
	// 	ElasticFileSystem:               resourceType,
	// 	ElasticIP:                       resourceType,
	// 	Elasticache:                     resourceType,
	// 	ElasticacheParameterGroups:      resourceType,
	// 	ElasticacheSubnetGroups:         resourceType,
	// 	GuardDuty:                       resourceType,
	// 	IAMGroups:                       resourceType,
	// 	IAMPolicies:                     resourceType,
	// 	IAMRoles:                        resourceType,
	// 	IAMServiceLinkedRoles:           resourceType,
	// 	IAMUsers:                        resourceType,
	// 	KMSCustomerKeys:                 kmsResourceType,
	// 	KinesisStream:                   resourceType,
	// 	LambdaFunction:                  resourceType,
	// 	LambdaLayer:                     resourceType,
	// 	LaunchConfiguration:             resourceType,
	// 	LaunchTemplate:                  resourceType,
	// 	MacieMember:                     resourceType,
	// 	MSKCluster:                      resourceType,
	// 	NatGateway:                      resourceType,
	// 	OIDCProvider:                    resourceType,
	// 	OpenSearchDomain:                resourceType,
	// 	Redshift:                        resourceType,
	// 	RdsSnapshot:                     resourceType,
	// 	S3:                              resourceType,
	// 	SNS:                             resourceType,
	// 	SQS:                             resourceType,
	// 	SageMakerNotebook:               resourceType,
	// 	SecretsManagerSecrets:           resourceType,
	// 	SecurityHub:                     resourceType,
	// 	Snapshots:                       resourceType,
	// 	TransitGateway:                  resourceType,
	// 	TransitGatewayRouteTable:        resourceType,
	// 	TransitGatewaysVpcAttachment:    resourceType,
	// 	TransitGatewayPeeringAttachment: resourceType,
	// 	VPC:                             resourceType,
	// }

	config := config.Config{}

	regions := []string{"us-west-2"}
	includeAfter := time.Date(2023, time.November, 28, 0, 0, 0, 0, time.UTC)
	excludeAfter := time.Now()
	// var query *aws.Query = &aws.Query{
	// 	Regions:              regions,
	// 	ExcludeResourceTypes: []string{},
	// }

	query, err := nuke_aws.NewQuery(
		regions,
		[]string{},
		[]string{"all"},
		[]string{},
		&excludeAfter,
		&includeAfter,
		false,
	)
	if err != nil {
		fmt.Errorf("Error creating query: %v", err)
	}

	// accountResources, err := nuke_aws.InspectResources(query)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // You can call GetRegion to examine a single region's resources
	// usWest1Resources := accountResources.GetRegion("us-west-1")

	// // Then interrogate them with the new methods:

	// // Count the number of any resource type within the region
	// countOfEc2InUsWest1 := usWest1Resources.CountOfResourceType("ec2")

	// fmt.Printf("countOfEc2InUsWest1: %d\n", countOfEc2InUsWest1)
	// // countOfEc2InUsWest1: 2

	// fmt.Printf("usWest1Resources.ResourceTypePresent(\"ec2\"):%b\n", usWest1Resources.ResourceTypePresent("ec2"))
	// // usWest1Resources.ResourceTypePresent("ec2"): true

	// // Get all the resource identifiers for a given resource type
	// // In this example, we're only looking for ec2 instances
	// resourceIds := usWest1Resources.IdentifiersForResourceType("ec2")

	// fmt.Printf("resourceIds: %s", resourceIds)
	os.Setenv("DISABLE_TELEMETRY", "true")
	telemetry.InitTelemetry("cloud-nuke", "local-test-environment")
	ctx := context.Background()
	fmt.Print("Getting all resources...\n")
	resources, err := nuke_aws.GetAllResources(ctx, query, config)
	if err != nil {
		fmt.Errorf("Error when getting resources: %v", err)
	}

	fmt.Printf("Total number of resources: %v\n", resources.TotalResourceCount())

	for _, region := range resources.Resources {
		fmt.Printf("Region: %v\n", region)
		for _, resource := range region.Resources {
			var awsResource nuke_aws.AwsResources = *resource
			fmt.Printf("Resource: %v\n", awsResource.ResourceName())
			fmt.Printf("ResourceIdentifiers: %v\n", awsResource.ResourceIdentifiers())
			fmt.Printf("\n")
		}
	}

	// // Delete all resources
	// deletionErrs := aws.NukeAllResources(resources, regions)
	// if deletionErrs != nil {
	// 	fmt.Errorf("Errors during deletion: %v", deletionErrs)
	// }
}
