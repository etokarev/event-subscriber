package main

import (
	"fmt"

	"github.com/Bestowinc/protoss/gen/go/proto/core"
	pas_schema "github.com/Bestowinc/protoss/gen/go/proto/pas/schema"
	"github.com/golang/protobuf/proto"
)

var (
	bestowApis = "type.bestow.co/"
)

// PolicyProtoEvents is the full list of the policy event messages.
var PolicyProtoEvents = []proto.Message{
	&pas_schema.BillingMethodChanged{},
	&pas_schema.CancellationReversed{},
	&pas_schema.CollateralAssignmentAdded{},
	&pas_schema.CollateralAssignmentReleased{},
	&pas_schema.CommissionRateChanged{},
	&pas_schema.ContingentBeneficiariesChanged{},
	&pas_schema.ContractAdjusted{},
	&pas_schema.CreditCardChanged{},
	&pas_schema.DeathClaimConfirmed{},
	&pas_schema.DeathClaimInitiated{},
	&pas_schema.DeathClaimInitiatedReversed{},
	&pas_schema.DisputeClosed{},
	&pas_schema.DisputeOpened{},
	&pas_schema.DocumentsLinked{},
	&pas_schema.FinalizeScheduleBindPolicyFailed{},
	&pas_schema.FreelookCancelled{},
	&pas_schema.LapseProcessed{},
	&pas_schema.UnlapseProcessed{},
	&pas_schema.NewPolicyApproved{},
	&pas_schema.NewPolicyDeclined{},
	&pas_schema.NewPolicyImported{},
	&pas_schema.NoteAdded{},
	&pas_schema.NoteDeleted{},
	&pas_schema.NotificationAdded{},
	&pas_schema.NotificationUpdated{},
	&pas_schema.OwnerChanged{},
	&pas_schema.PaymentFailed{},
	&pas_schema.PaymentMissed{},
	&pas_schema.PaymentProcessed{},
	&pas_schema.PaymentSuspended{},
	&pas_schema.PaymentsAdjusted{},
	&pas_schema.PolicyActionCreated{},
	&pas_schema.PolicyActionReverted{},
	&pas_schema.PolicyBindScheduled{},
	&pas_schema.PolicyBound{},
	&pas_schema.PolicyCancelled{},
	&pas_schema.PolicyCorrectionApplied{},
	&pas_schema.PolicyDocumentsAccepted{},
	&pas_schema.PolicyDocumentsChanged{},
	&pas_schema.PolicyDocumentsGenerated{},
	&pas_schema.PolicyExpired{},
	&pas_schema.PolicyGraceExtended{},
	&pas_schema.PolicyGraced{},
	&pas_schema.PolicyLapsed{},
	&pas_schema.PolicyPendingCancelled{},
	&pas_schema.PolicyPendingUnlapsed{},
	&pas_schema.PolicyReinstated{},
	&pas_schema.PolicyReplacementAdded{},
	&pas_schema.PolicyRescinded{},
	&pas_schema.PolicyStaled{},
	&pas_schema.PolicyStateUpdated{},
	&pas_schema.PolicySuspended{},
	&pas_schema.PolicyUngraced{},
	&pas_schema.PolicyUnlapsed{},
	&pas_schema.PolicyUnsuspended{},
	&pas_schema.PolicyVoided{},
	&pas_schema.PrimaryBeneficiariesChanged{},
	&pas_schema.RefundProcessed{},
	&pas_schema.ScheduleBindPolicyExpired{},
	&pas_schema.SecondaryContactChanged{},
	&pas_schema.SecondaryContactReset{},
	&pas_schema.SpecialInstructionsChanged{},
	&pas_schema.SubscriptionCancelled{},
	&pas_schema.SubscriptionChanged{},
	&pas_schema.PolicyAlertAdded{},
}

type EventEnvelopeUnpacker struct {
	typeLookup map[string]proto.Message
}

// NewEventEnvelopeUnpackerDesc constructs a new NewEventEnvelopeUnpackerDesc object. The type
// lookup map is populated from the types of the passed in messages.
func NewUnpacker(messages []proto.Message) (*EventEnvelopeUnpacker, error) {
	typeLookup := make(map[string]proto.Message, len(messages))
	for _, v := range messages {
		typeLookup[bestowApis+proto.MessageName(v)] = v
	}

	return &EventEnvelopeUnpacker{typeLookup: typeLookup}, nil
}

func (b *EventEnvelopeUnpacker) Unpack(payload *core.EventEnvelope) (proto.Message, error) {

	if val, ok := b.typeLookup[payload.TypeUrl]; ok {
		message := proto.Clone(val)
		if err := proto.Unmarshal(payload.Value, message); err != nil {
			return nil, fmt.Errorf("could not unmarshal EventEnvelope with typeUrl: %v", payload.TypeUrl)
		}
		return message, nil
	}

	return nil, fmt.Errorf("unknown typeUrl: %v", payload.TypeUrl)
}
