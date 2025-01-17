package resources

import (
	"context"
	"fmt"

	"github.com/andrewderr/cloud-nuke-a1/config"
	"github.com/andrewderr/cloud-nuke-a1/logging"
	"github.com/andrewderr/cloud-nuke-a1/report"
	"github.com/andrewderr/cloud-nuke-a1/telemetry"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/gruntwork-io/go-commons/errors"
	commonTelemetry "github.com/gruntwork-io/go-commons/telemetry"
)

// Returns a list of strings of ACM ARNs
func (a *ACM) getAll(c context.Context, configObj config.Config) ([]*string, error) {

	params := &acm.ListCertificatesInput{}

	acmArns := []*string{}
	err := a.Client.ListCertificatesPages(params,
		func(page *acm.ListCertificatesOutput, lastPage bool) bool {
			for i := range page.CertificateSummaryList {
				logging.Debug(fmt.Sprintf("Found ACM %s with domain name %s",
					*page.CertificateSummaryList[i].CertificateArn, *page.CertificateSummaryList[i].DomainName))
				if a.shouldInclude(page.CertificateSummaryList[i], configObj) {
					logging.Debug(fmt.Sprintf(
						"Including ACM %s", *page.CertificateSummaryList[i].CertificateArn))
					acmArns = append(acmArns, page.CertificateSummaryList[i].CertificateArn)
				} else {
					logging.Debug(fmt.Sprintf(
						"Skipping ACM %s", *page.CertificateSummaryList[i].CertificateArn))
				}
			}

			return !lastPage
		},
	)
	if err != nil {
		return nil, errors.WithStackTrace(err)
	}

	return acmArns, nil
}

func (a *ACM) shouldInclude(acm *acm.CertificateSummary, configObj config.Config) bool {
	if acm == nil {
		return false
	}

	if acm.InUse != nil && *acm.InUse {
		logging.Debug(fmt.Sprintf("ACM %s is in use", *acm.CertificateArn))
		return false
	}

	shouldInclude := configObj.ACM.ShouldInclude(config.ResourceValue{
		Name: acm.DomainName,
		Time: acm.CreatedAt,
	})
	logging.Debug(fmt.Sprintf("shouldInclude result for ACM: %s w/ domain name: %s, time: %s, and config: %+v",
		*acm.CertificateArn, *acm.DomainName, acm.CreatedAt, configObj.ACM))
	return shouldInclude
}

// Deletes all ACMs
func (a *ACM) nukeAll(arns []*string) error {
	if len(arns) == 0 {
		logging.Debugf("No ACMs to nuke in region %s", a.Region)
		return nil
	}

	logging.Debugf("Deleting all ACMs in region %s", a.Region)

	deletedCount := 0
	for _, acmArn := range arns {
		params := &acm.DeleteCertificateInput{
			CertificateArn: acmArn,
		}

		_, err := a.Client.DeleteCertificate(params)
		if err != nil {
			logging.Debugf("[Failed] %s", err)
			telemetry.TrackEvent(commonTelemetry.EventContext{
				EventName: "Error Nuking ACM",
			}, map[string]interface{}{
				"region": a.Region,
			})
		} else {
			deletedCount++
			logging.Debugf("Deleted ACM: %s", *acmArn)
		}

		e := report.Entry{
			Identifier:   *acmArn,
			ResourceType: "ACM",
			Error:        err,
		}
		report.Record(e)
	}

	logging.Debugf("[OK] %d ACM(s) terminated in %s", deletedCount, a.Region)
	return nil
}
