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
)

func POSTMailSend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	b := &Body{}

	if err := json.NewDecoder(r.Body).Decode(b); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("subject=%s\n", b.Subject)
	for _, c := range b.Content {
		fmt.Printf("type=%s, value=%s\n", c.Type_, c.Value)
	}

	for _, p := range b.Personalizations {
		fmt.Printf("personalized subject=%s\n", p.Subject)

		for _, t := range p.To {
			fmt.Printf("send to email=%s, name=%s\n", t.Email, t.Name)
		}

		for _, t := range p.Cc {
			fmt.Printf("cc to email=%s, name=%s\n", t.Email, t.Name)
		}

		for _, t := range p.Bcc {
			fmt.Printf("bcc to email=%s, name=%s\n", t.Email, t.Name)
		}
	}
	fmt.Printf("send from email=%s, name=%s\n", b.From.Email, b.From.Name)
	fmt.Printf("reply to email=%s, name=%s\n", b.ReplyTo.Email, b.ReplyTo.Name)

	enc := json.NewEncoder(w)
	if err := enc.Encode(b); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}