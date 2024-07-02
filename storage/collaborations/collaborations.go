package collaborations

import (
	pb "MusicMesh/Collaboration-MusicMesh/generate/collaboration"
	"database/sql"
)

type CollaborationRepo struct {
	pb.UnimplementedCollaborationServiceServer
	DB *sql.DB
}

func NewCollaborationRepo(db *sql.DB) *CollaborationRepo {
	return &CollaborationRepo{DB: db}
}
