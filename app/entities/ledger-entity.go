package entities

import (
	types "bank/app/types"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Quota struct {
	Total int64
	Used  int64
}

type LedgerEntity struct {
	AccountId int64
	Version   int64
	Timestamp int64
	Quota     Quota
	EventType string
	Events    []types.LedgerEvent
}

type ILedgerEntity interface {
	ApplyEvents(events []types.LedgerEvent) error
	AddQuota(AssetType string, amount int64) (string, error)
	ReduceQuota(AssetType string, amount int64) error
	UseQuota(AssetType string, amount int64) error
}

func (l *LedgerEntity) AddQuota(AssetType string, amount int64, reason string) (string, error) {
	l.Version = l.Version + 1

	a := types.LedgerEvent{}
	a.AccountId = l.AccountId
	a.Amount = amount
	a.AssetType = AssetType
	a.Reason = reason
	a.EventName = "GRANT"
	a.TransactionId = uuid.New().String()
	a.Timestamp = time.Now().Unix()
	a.Version = l.Version

	l.Events = append(l.Events, a)

	l.ApplyEvents(l.Events)

	return a.TransactionId, nil
}

func (l *LedgerEntity) ReduceQuota(AssetType string, amount int64, reason string) (string, error) {
	l.Version = l.Version + 1

	a := types.LedgerEvent{}
	a.AccountId = l.AccountId
	a.Amount = amount
	a.AssetType = AssetType
	a.Reason = reason
	a.EventName = "DEDUCT"
	a.TransactionId = uuid.New().String()
	a.Timestamp = time.Now().Unix()
	a.Version = l.Version

	l.Events = append(l.Events, a)

	l.ApplyEvents(l.Events)

	return a.TransactionId, nil
}

func (l *LedgerEntity) UseQuota(AssetType string, amount int64, reason string) (string, error) {
	l.Version = l.Version + 1

	a := types.LedgerEvent{}
	a.AccountId = l.AccountId
	a.Amount = amount
	a.AssetType = AssetType
	a.Reason = reason
	a.EventName = "USE"
	a.TransactionId = uuid.New().String()
	a.Timestamp = time.Now().Unix()
	a.Version = l.Version

	l.Events = append(l.Events, a)

	err := l.ApplyEvents(l.Events)

	if err != nil {
		return "", err
	}

	return a.TransactionId, nil
}

func (l *LedgerEntity) GetEvents() []types.LedgerEvent {
	return l.Events
}

func (l *LedgerEntity) ApplyEvents(events []types.LedgerEvent) error {
	for _, v := range events {
		switch eventName := v.EventName; eventName {
		case "GRANT":
			l.Version = v.Version
			quota := l.Quota
			quota.Total = quota.Total + v.Amount

			l.Quota = quota

			break
		case "DEDUCT":
			l.Version = v.Version
			currentQuotaAvailable := l.Quota.Total - l.Quota.Used

			if (currentQuotaAvailable - v.Amount) >= 0 {
				quota := l.Quota
				quota.Total = quota.Total - v.Amount

				l.Quota = quota
			} else {
				panic("Operation forbidden, quota cannot be in negative balance")
			}
			break
		case "USE":
			l.Version = v.Version
			currentQuotaAvailable := l.Quota.Total - l.Quota.Used

			if (currentQuotaAvailable - v.Amount) >= 0 {
				quota := l.Quota
				quota.Used += v.Amount

				l.Quota = quota
			} else {
				return errors.New("not enough response to completed this operation")
			}
			break
		default:
			fmt.Printf("%s.\n", eventName)
		}
	}

	return nil
}
