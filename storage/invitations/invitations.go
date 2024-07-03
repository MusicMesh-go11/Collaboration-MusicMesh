package invitations

import (
	pb "MusicMesh/Collaboration-MusicMesh/generate/invitation"
	"database/sql"
)

type InvitationRepo struct {
	pb.UnimplementedInvitationServiceServer
	DB *sql.DB
}

func NewInvitationRepo(db *sql.DB) *InvitationRepo {
	return &InvitationRepo{DB: db}
}
