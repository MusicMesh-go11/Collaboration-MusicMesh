package collaborations

import (
	pb "MusicMesh/Collaboration-MusicMesh/generate/collaboration"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type CollaborationRepo struct {
	pb.UnimplementedCollaborationServiceServer
	DB *sql.DB
}

func NewCollaborationRepo(db *sql.DB) *CollaborationRepo {
	return &CollaborationRepo{DB: db}
}
func (s *CollaborationRepo) Create(ctx context.Context, req *pb.Collaboration) (*pb.Void, error) {
	query := `INSERT INTO collaborations (composition_id, user_id, role, joined_at) VALUES ($1, $2, $3, $4)`
	_, err := s.DB.ExecContext(ctx, query, req.CompositionId, req.UserId, req.Role, time.Now())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to create collaboration: %v", err))
	}
	return &pb.Void{}, nil
}
func (s *CollaborationRepo) GetById(ctx context.Context, req *pb.CollaborationID) (*pb.CollaborationRes, error) {
	query := `SELECT * FROM collaborations WHERE collaboration_id = $1`
	collab := &pb.CollaborationRes{}
	s.DB.QueryRow(query, req.CollaborationID).Scan(&collab.Id, &collab.CompositionId, &collab.Us)

	return &collab, nil
}

func (s *CollaborationRepo) GetByCompositionId(ctx context.Context, req *pb.CompositionID) (*pb.CollaborationRes, error) {
	query := `SELECT collaboration_id, composition_id, user_id, role, joined_at FROM collaborations WHERE composition_id = $1`
	row := s.db.QueryRowContext(ctx, query, req.CompositionId)

	var collab pb.CollaborationRes
	err := row.Scan(&collab.Id, &collab.CompositionId, &collab.InviterId, &collab.Role, &collab.JoinedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New(fmt.Sprintf("collaboration with composition ID %s not found", req.CompositionId))
	} else if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to retrieve collaboration: %v", err))
	}

	return &collab, nil
}

func (s *CollaborationRepo) Delete(ctx context.Context, req *pb.CollaborationID) (*pb.Void, error) {
	query := `DELETE FROM collaborations WHERE collaboration_id = $1`
	_, err := s.db.ExecContext(ctx, query, req.CollaborationId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to delete collaboration: %v", err))
	}
	return &pb.Void{}, nil
}
