package models

import "fmt"

type Guild struct {
	Id                       string `json:"id"`
	Name                     string `json:"name"`
	ApproximateMemberCount   int    `json:"approximate_member_count"`
	ApproximatePresenceCount int    `json:"approximate_presence_count"`
}

func (g *Guild) String() string {
	return fmt.Sprintf(
		"%s\n%s\n%d\n%d\n",
		g.Id,
		g.Name,
		g.ApproximateMemberCount,
		g.ApproximatePresenceCount,
	)
}
