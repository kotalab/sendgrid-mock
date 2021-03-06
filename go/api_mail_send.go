/*
 * SendGrid v3 API Documentation
 *
 * # The SendGrid Web API V3 Documentation  This is the entirety of the documented v3 endpoints. We have updated all the descriptions, parameters, requests, and responses.  ## Authentication   Every endpoint requires Authentication in the form of an Authorization Header:  Authorization: Bearer API_KEY
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	types "github.com/kotalab/sendgrid-mock/go/generated_type"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func POSTMailSend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	b := types.Body{}

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := debugSendMail(b)
	if err != nil {
		fmt.Printf("failed to send email. reason=%s", err.Error())
	}

	for key, values := range res.Headers {
		w.Header().Add(key, strings.Join(values, ";"))
	}

	w.WriteHeader(res.StatusCode)
	fmt.Fprintln(w, res.Body)
}

func debugSendMail(b types.Body) (*rest.Response, error) {
	m := mail.NewV3Mail()
	m.Subject = b.Subject
	for _, p := range b.Personalizations {
		ps := mail.NewPersonalization()

		for _, bcc := range p.Bcc {
			ps.AddBCCs(mail.NewEmail(bcc.Name, bcc.Email))
		}

		for _, cc := range p.Cc {
			ps.AddCCs(mail.NewEmail(cc.Name, cc.Email))
		}

		for _, to := range p.To {
			ps.AddTos(mail.NewEmail(to.Name, to.Email))
		}

		if p.CustomArgs != nil {
			ca := *p.CustomArgs
			c, _ := ca.(map[string]string)
			for key, value := range c {
				ps.SetCustomArg(key, value)
			}
		}

		if p.Headers != nil {
			h := *p.Headers
			headers, _ := h.(map[string]string)
			for key, value := range headers {
				ps.SetHeader(key, value)
			}
		}

		ps.SetSendAt(int(p.SendAt))

		if p.Substitutions != nil {
			s := *p.Substitutions
			sub, _ := s.(map[string]string)
			for key, value := range sub {
				ps.SetSubstitution(key, value)
			}
		}

		m.AddPersonalizations(ps)
	}

	for _, a := range b.Attachments {
		at := mail.NewAttachment()
		at.SetContent(a.Content)
		at.SetContentID(a.ContentId)
		at.SetDisposition(a.Disposition)
		at.SetFilename(a.Filename)
		at.SetType(a.Type_)

		m.AddAttachment(at)
	}

	for _, ca := range b.Categories {
		m.AddCategories(ca)
	}

	for _, content := range b.Content {
		c := mail.NewContent(content.Type_, content.Value)
		m.AddContent(c)
	}

	if b.Sections != nil {
		se := *b.Sections
		s, _ := se.(map[string]string)
		for key, value := range s {
			m.AddSection(key, value)
		}
	}

	if b.Asm != nil {
		asm := mail.NewASM()
		asm.SetGroupID(int(b.Asm.GroupId))
		var ho []int
		for _, g := range b.Asm.GroupsToDisplay {
			i := int(g)
			ho = append(ho, i)
		}
		asm.AddGroupsToDisplay(ho...)
		m.SetASM(asm)
	}
	m.SetBatchID(b.BatchId)

	if b.CustomArgs != nil {
		ca := *b.CustomArgs
		c, _ := ca.(map[string]string)
		for key, value := range c {
			m.SetCustomArg(key, value)
		}
	}

	f := mail.NewEmail(b.From.Name, b.From.Email)
	m.SetFrom(f)

	if b.Headers != nil {
		h := *b.Headers
		headers, _ := h.(map[string]string)
		for key, value := range headers {
			m.SetHeader(key, value)
		}
	}

	m.SetIPPoolID(b.IpPoolName)

	ms := mail.NewMailSettings()
	ms.SetSandboxMode(mail.NewSetting(true))

	if b.MailSettings != nil {
		if b.MailSettings.Bcc != nil {
			bcc := mail.NewBCCSetting()
			bcc.SetEmail(b.MailSettings.Bcc.Email)
			bcc.SetEnable(b.MailSettings.Bcc.Enable)
			ms.SetBCC(bcc)
		}

		if b.MailSettings.BypassListManagement != nil {
			ms.SetBypassListManagement(mail.NewSetting(b.MailSettings.BypassListManagement.Enable))
		}

		if b.MailSettings.Footer != nil {
			fs := mail.NewFooterSetting()
			fs.SetEnable(b.MailSettings.Footer.Enable)
			fs.SetHTML(b.MailSettings.Footer.Html)
			fs.SetText(b.MailSettings.Footer.Text)
			ms.SetFooter(fs)
		}

		if b.MailSettings.SpamCheck != nil {
			ss := mail.NewSpamCheckSetting()
			ss.SetEnable(b.MailSettings.SpamCheck.Enable)
			ss.SetPostToURL(b.MailSettings.SpamCheck.PostToUrl)
			ss.SetSpamThreshold(int(b.MailSettings.SpamCheck.Threshold))
			ms.SetSpamCheckSettings(ss)
		}
	}
	m.SetMailSettings(ms)

	if b.ReplyTo != nil {
		rt := mail.NewEmail(b.ReplyTo.Name, b.ReplyTo.Email)
		m.SetReplyTo(rt)
	}

	m.SetSendAt(int(b.SendAt))

	m.SetTemplateID(b.TemplateId)

	if b.TrackingSettings != nil {
		trs := mail.NewTrackingSettings()

		if b.TrackingSettings.ClickTracking != nil {
			cts := mail.NewClickTrackingSetting()
			cts.SetEnable(b.TrackingSettings.ClickTracking.Enable)
			cts.SetEnableText(b.TrackingSettings.ClickTracking.EnableText)
			trs.SetClickTracking(cts)
		}

		if b.TrackingSettings.Ganalytics != nil {
			gas := mail.NewGaSetting()
			gas.SetEnable(b.TrackingSettings.Ganalytics.Enable)
			gas.SetCampaignContent(b.TrackingSettings.Ganalytics.UtmContent)
			gas.SetCampaignMedium(b.TrackingSettings.Ganalytics.UtmMedium)
			gas.SetCampaignName(b.TrackingSettings.Ganalytics.UtmCampaign)
			gas.SetCampaignSource(b.TrackingSettings.Ganalytics.UtmSource)
			gas.SetCampaignTerm(b.TrackingSettings.Ganalytics.UtmTerm)
			trs.SetGoogleAnalytics(gas)
		}

		if b.TrackingSettings.OpenTracking != nil {
			ots := mail.NewOpenTrackingSetting()
			ots.SetEnable(b.TrackingSettings.OpenTracking.Enable)
			ots.SetSubstitutionTag(b.TrackingSettings.OpenTracking.SubstitutionTag)
			trs.SetOpenTracking(ots)
		}

		if b.TrackingSettings.SubscriptionTracking != nil {
			sts := mail.NewSubscriptionTrackingSetting()
			sts.SetSubstitutionTag(b.TrackingSettings.SubscriptionTracking.SubstitutionTag)
			sts.SetEnable(b.TrackingSettings.SubscriptionTracking.Enable)
			sts.SetText(b.TrackingSettings.SubscriptionTracking.Text)
			sts.SetHTML(b.TrackingSettings.SubscriptionTracking.Html)
			trs.SetSubscriptionTracking(sts)
		}
		m.SetTrackingSettings(trs)
	}

	cl := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	return cl.Send(m)
}
