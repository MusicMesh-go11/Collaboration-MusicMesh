package collaborations

import (
	pb "MusicMesh/Collaboration-MusicMesh/generate/collaboration"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type CollaborationRepo struct {
	pb.UnimplementedCollaborationServiceServer
	DB *sql.DB
}

func NewCollaborationRepo(db *sql.DB) *CollaborationRepo {
	return &CollaborationRepo{DB: db}
}
func (s *CollaborationServiceServer) Create(ctx context.Context, req *pb.Collaboration) (*pb.Void, error) {
    query := `INSERT INTO collaborations (composition_id, user_id, role, joined_at) VALUES ($1, $2, $3, $4)`
    _, err := s.db.ExecContext(ctx, query, req.CompositionId, req.UserId, req.Role, time.Now())
    if err != nil {
        return nil, errors.New(fmt.Sprintf("failed to create collaboration: %v", err))
    }
    return &pb.Void{}, nil
}
func (s *CollaborationServiceServer) GetById(ctx context.Context, req *pb.CollaborationID) (*pb.CollaborationRes, error) {
    query := `SELECT collaboration_id, composition_id, user_id, role, joined_at FROM collaborations WHERE collaboration_id = $1`
    row := s.db.QueryRowContext(ctx, query, req.CollaborationId)

    var collab pb.CollaborationRes
    err := row.Scan(&collab.Id, &collab.CompositionId, &collab.InviterId, &collab.Role, &collab.JoinedAt)
    if err == sql.ErrNoRows {
        return nil, errors.New(fmt.Sprintf("collaboration with ID %s not found", req.CollaborationId))
    } else if err != nil {
        return nil, errors.New(fmt.Sprintf("failed to retrieve collaboration: %v", err))
    }

    return &collab, nil
}

func (s *CollaborationServiceServer) GetByCompositionId(ctx context.Context, req *pb.CompositionID) (*pb.CollaborationRes, error) {
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

func (s *CollaborationServiceServer) Delete(ctx context.Context, req *pb.CollaborationID) (*pb.Void, error) {
    query := `DELETE FROM collaborations WHERE collaboration_id = $1`
    _, err := s.db.ExecContext(ctx, query, req.CollaborationId)
    if err != nil {
        return nil, errors.New(fmt.Sprintf("failed to delete collaboration: %v", err))
    }
    return &pb.Void{}, nil
}