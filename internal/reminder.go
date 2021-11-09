package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"go.uber.org/zap"
)

type Reminder struct {
	sync.RWMutex
	reminders map[string]ReminderPayload
}

func NewReminder() Reminder {
	return Reminder{reminders: make(map[string]ReminderPayload)}
}

func (s *Reminder) Add(payload ReminderPayload) error {
	s.Lock()
	defer s.Unlock()

	s.reminders[uuid.NewString()] = payload

	return nil
}

func (s *Reminder) Run() {
	s.Lock()
	defer s.Unlock()

	now := time.Now().UTC()

	for id, payload := range s.reminders {
		if payload.GetReminderDate().Before(now) {
			err := s.SendSMS(payload)
			if err != nil {
				zap.L().Error("reminder-me-scheduler-sms", zap.Error(err), zap.Reflect("payload", payload))
			}

			delete(s.reminders, id)
		}
	}
}

func (s *Reminder) SendSMS(payload ReminderPayload) error {
	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		AccountSid: os.Getenv("TWILIO_ACCOUNT_SID"),
		Username:   os.Getenv("TWILIO_ACCOUNT_SID"),
		Password:   os.Getenv("TWILIO_TOKEN"),
	})

	params := &openapi.CreateMessageParams{}
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetTo(payload.GetPhoneNumber())
	params.SetBody(fmt.Sprintf("Olá %s, não esqueça: '%s'", payload.Name, payload.Message))

	resp, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		zap.L().Error("send-sms", zap.Error(err), zap.Reflect("payload", payload))
		return err
	}

	response, _ := json.Marshal(*resp)

	zap.L().Info("send-sms", zap.ByteString("response", response))

	return nil
}
