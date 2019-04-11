/*
 * SendGrid v3 API Documentation
 *
 * # The SendGrid Web API V3 Documentation  This is the entirety of the documented v3 endpoints. We have updated all the descriptions, parameters, requests, and responses.  ## Authentication   Every endpoint requires Authentication in the form of an Authorization Header:  Authorization: Bearer API_KEY
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Allows you to track whether a recipient clicked a link in your email.
type MailsendTrackingSettingsClickTracking struct {

	// Indicates if this setting is enabled.
	Enable bool `json:"enable,omitempty"`

	// Indicates if this setting should be included in the text/plain portion of your email.
	EnableText bool `json:"enable_text,omitempty"`
}
