package invitations

import (
	pb "MusicMesh/Collaboration-MusicMesh/generate/invitation"
	"database/sql"
	"fmt"
	"context"
	"errors"
	"time"
)

type InvitationRepo struct{
	pb.UnimplementedInvitationServiceServer
	DB * sql.DB
}

func NewInvitationRepo(db *sql.DB) *InvitationRepo{
	return &InvitationRepo{DB: db}
}
func (s *InvitationServiceServer) Create(ctx context.Context, req *pb.Invitation) (*pb.Void, error) {
    query := `INSERT INTO invitations (composition_id, inviter_id, invitee_id, status, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6)`
    now := time.Now()
    _, err := s.db.ExecContext(ctx, query, req.CompositionId, req.InviterId, req.InviteeId, req.Status, now, now)
    if err != nil {
        return nil, errors.New(fmt.Sprintf("failed to create invitation: %v", err))
    }
    return &pb.Void{}, nil
}

func (s *InvitationServiceServer) GetById(ctx context.Context, req *pb.InvitationID) (*pb.InvitationRes, error) {
    query := `SELECT id, composition_id, inviter_id, invitee_id, status, created_at, updated_at, deleted_at 
              FROM invitations WHERE id = $1`
    row := s.db.QueryRowContext(ctx, query, req.Id)

    var inv pb.InvitationRes
    err := row.Scan(&inv.Id, &inv.CompositionId, &inv.InviterId, &inv.InviteeId, &inv.Status, &inv.CreatedAt, &inv.UpdatedAt, &inv.DeletedAt)
    if err == sql.ErrNoRows {
        return nil, errors.New(fmt.Sprintf("invitation with ID %s not found", req.Id))
    } else if err != nil {
        return nil, errors.New(fmt.Sprintf("failed to retrieve invitation: %v", err))
    }

    return &inv, nil
}

func (s *InvitationServiceServer) Delete(ctx context.Context, req *pb.InvitationID) (*pb.Void, error) {
    query := `UPDATE invitations SET deleted_at = $1 WHERE id = $2`
    now := time.Now().Unix()
    _, err := s.db.ExecContext(ctx, query, now, req.Id)
    if err != nil {
        return nil, errors.New(fmt.Sprintf("failed to delete invitation: %v", err))
    }
    return &pb.Void{}, nil
}