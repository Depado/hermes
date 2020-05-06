package main

import (
	"html/template"

	"github.com/Depado/hermes"
)

type inviteCode struct {
}

func (w *inviteCode) Name() string {
	return "invite_code"
}

func (w *inviteCode) Email() hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Name: "Jon Snow",
			Intros: []template.HTML{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Please copy your invite code:",
					InviteCode:   "123456",
				},
			},
			Outros: []template.HTML{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
}
