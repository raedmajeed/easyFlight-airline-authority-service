package utils

import (
	"fmt"
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	"time"
)

func WriteMessageToEmail(OTP, email string) *dom.EmailMessage {
	messageDescription := fmt.Sprintf("OTP to verify your email %v, ", OTP)

	return &dom.EmailMessage{
		Email:   email,
		Subject: fmt.Sprintf("OTP: %v, YOUR OTP TO VERIFY.", OTP),
		Content: messageDescription,
	}
}

type AirlineInitialCred struct {
	Email    string
	Password string
}

func SendAirlinePasswordEmail(email, password string) *dom.EmailMessage {
	messageDescription := fmt.Sprintf("Temporary password %v, reset your password using forgot password request", password)

	return &dom.EmailMessage{
		Email:   email,
		Subject: fmt.Sprintln("TEMPORARY LOGIN CRED."),
		Content: messageDescription,
	}
}

func SendAirlineFareSetFailure(email, flightNumber, airport string, date time.Time, id uint) *dom.EmailMessage {
	messageDescription := fmt.Sprintf("error setting fare for flight %v, charted to depart from %v on %v, chart Id is %v", flightNumber, airport, date, id)

	return &dom.EmailMessage{
		Email:   email,
		Subject: fmt.Sprintln("TEMPORARY LOGIN CRED."),
		Content: messageDescription,
	}
}
