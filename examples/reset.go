package main

import (
	"html/template"

	"github.com/Depado/hermes"
)

type reset struct {
}

func (r *reset) Name() string {
	return "reset"
}

func (r *reset) Email() hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Name: "Jon Snow",
			Intros: []template.HTML{
				"You have received this email because a password reset request for Hermes account was received.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Click the button below to reset your password:",
					Button: hermes.Button{
						Color: "#DC4D2F",
						Text:  "Reset your password",
						Link:  "https://hermes-example.com/reset-password?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
			Outros: []template.HTML{
				"If you did not request a password reset, no further action is required on your part.",
			},
			Signature: "Thanks",
		},
	}
}
