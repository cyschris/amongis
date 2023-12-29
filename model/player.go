package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/peterstace/simplefeatures/geom"
)

type PlayerModel struct {
	Name      string     `json:"name" required:"true"`
	Id        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	Role      string     `json:"role"`
	Room      string     `json:"room"`
	Status    string     `json:"status"`
	Location  geom.Point `json:"location" required:"true"`
}

func (p *PlayerModel) Validate() error {
	if p.CreatedAt.IsZero() {
		p.CreatedAt = time.Now()
	}
	if p.Name == "" {
		return fmt.Errorf("empty user name")
	}
	if p.Role != "killer" && p.Role != "player" {
		return fmt.Errorf("invalid role")
	}

	return nil
}

type ActionModel struct {
	ActionType string `json:"action_type" required:"true"`
	Target     string `json:"target" required:"true"`
}


