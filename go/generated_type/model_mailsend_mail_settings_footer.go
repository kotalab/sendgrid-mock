/*
 * SendGrid v3 API Documentation
 *
 * # The SendGrid Web API V3 Documentation  This is the entirety of the documented v3 endpoints. We have updated all the descriptions, parameters, requests, and responses.  ## Authentication   Every endpoint requires Authentication in the form of an Authorization Header:  Authorization: Bearer API_KEY
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package generated_type

// The default footer that you would like included on every email.
type MailsendMailSettingsFooter struct {

	// Indicates if this setting is enabled.
	Enable bool `json:"enable,omitempty"`

	// The plain text content of your footer.
	Text string `json:"text,omitempty"`

	// The HTML content of your footer.
	Html string `json:"html,omitempty"`
}
