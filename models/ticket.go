package models

import (
	"gorm.io/gorm"
)

type TicketStatus string

const (
	StatusPending  TicketStatus = "Pending"
	StatusApproved TicketStatus = "Approved"
)

type TicketCheckStatus string

const (
	CheckStatusNew     TicketCheckStatus = "NewTicket"
	CheckStatusChecked TicketCheckStatus = "Checked"
)

type Ticket struct {
	gorm.Model

	TicketName        string            `json:"ticket_name"`
	TicketType        string            `json:"ticket_type"`
	TicketStatus      TicketStatus      `json:"ticket_status"`
	TicketCheckStatus TicketCheckStatus `json:"ticket_check_status"`

	Products []Product `json:"products"`
}
