package resources

import (
	"context"
	"regexp"
	"testing"

	"github.com/andrewderr/cloud-nuke-a1/config"
	"github.com/andrewderr/cloud-nuke-a1/telemetry"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/aws/aws-sdk-go/service/configservice/configserviceiface"
	"github.com/stretchr/testify/require"
)

type mockedConfigServiceRule struct {
	configserviceiface.ConfigServiceAPI
	DescribeConfigRulesOutput            configservice.DescribeConfigRulesOutput
	DeleteConfigRuleOutput               configservice.DeleteConfigRuleOutput
	DeleteRemediationConfigurationOutput configservice.DeleteRemediationConfigurationOutput
}

func (m mockedConfigServiceRule) DescribeConfigRulesPages(input *configservice.DescribeConfigRulesInput, fn func(*configservice.DescribeConfigRulesOutput, bool) bool) error {
	fn(&m.DescribeConfigRulesOutput, true)
	return nil
}

func (m mockedConfigServiceRule) DeleteConfigRule(input *configservice.DeleteConfigRuleInput) (*configservice.DeleteConfigRuleOutput, error) {
	return &m.DeleteConfigRuleOutput, nil
}

func (m mockedConfigServiceRule) DeleteRemediationConfiguration(input *configservice.DeleteRemediationConfigurationInput) (*configservice.DeleteRemediationConfigurationOutput, error) {
	return &m.DeleteRemediationConfigurationOutput, nil
}

func TestConfigServiceRule_GetAll(t *testing.T) {
	telemetry.InitTelemetry("cloud-nuke", "")
	t.Parallel()

	testName1 := "test-rule-1"
	testName2 := "test-rule-2"
	csr := ConfigServiceRule{
		Client: mockedConfigServiceRule{
			DescribeConfigRulesOutput: configservice.DescribeConfigRulesOutput{
				ConfigRules: []*configservice.ConfigRule{
					{ConfigRuleName: aws.String(testName1), ConfigRuleState: aws.String("ACTIVE")},
					{ConfigRuleName: aws.String(testName2), ConfigRuleState: aws.String("ACTIVE")},
				},
			},
		},
	}

	tests := map[string]struct {
		configObj config.ResourceType
		expected  []string
	}{
		"emptyFilter": {
			configObj: config.ResourceType{},
			expected:  []string{testName1, testName2},
		},
		"nameExclusionFilter": {
			configObj: config.ResourceType{
				ExcludeRule: config.FilterRule{
					NamesRegExp: []config.Expression{{
						RE: *regexp.MustCompile(testName1),
					}}},
			},
			expected: []string{testName2},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			names, err := csr.getAll(context.Background(), config.Config{
				ConfigServiceRule: tc.configObj,
			})

			require.NoError(t, err)
			require.Equal(t, tc.expected, aws.StringValueSlice(names))
		})
	}
}

func TestConfigServiceRule_NukeAll(t *testing.T) {
	telemetry.InitTelemetry("cloud-nuke", "")
	t.Parallel()

	csr := ConfigServiceRule{
		Client: mockedConfigServiceRule{
			DeleteConfigRuleOutput:               configservice.DeleteConfigRuleOutput{},
			DeleteRemediationConfigurationOutput: configservice.DeleteRemediationConfigurationOutput{},
		},
	}

	err := csr.nukeAll([]string{"test"})
	require.NoError(t, err)
}
