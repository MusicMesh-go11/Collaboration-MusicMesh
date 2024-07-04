package invitations

import (
	pb "MusicMesh/Collaboration-MusicMesh/generate/invitation"
	"context"
	"database/sql"
	"fmt"
)

func (s *InvitationRepo) Create(ctx context.Context, req *pb.Invitation) (*pb.Void, error) {
	query := `INSERT INTO invitations (composition_id, inviter_id, invitee_id, status) VALUES ($1, $2, $3, $4)`
	_, err := s.DB.ExecContext(ctx, query, req.CompositionId, req.InviterId, req.InviteeId, req.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to create invitation: %v", err)
	}
	return &pb.Void{}, nil
}

func (s *InvitationRepo) GetById(ctx context.Context, req *pb.InvitationID) (*pb.InvitationRes, error) {
	query := `SELECT id, composition_id, inviter_id, invitee_id, status FROM invitations WHERE id = $1`
	invite := &pb.InvitationRes{}
	err := s.DB.QueryRowContext(ctx, query, req.Id).Scan(&invite.Id, &invite.CompositionId, &invite.InviterId, &invite.InviteeId, &invite.Status, &invite.CreatedAt, &invite.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("invitation with ID %s not found", req.Id)
	} else if err != nil {
		return nil, fmt.Errorf("failed to retrieve invitation: %v", err)
	}
	return invite, nil
}

func (s *InvitationRepo) Delete(ctx context.Context, req *pb.InvitationID) (*pb.Void, error) {
	_, err := s.DB.ExecContext(ctx, `UPDATE invitations SET 
		deleted_at = date_part('epoch', current_timestamp)::INT 
		WHERE id = $1 AND deleted_at = 0`, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete invitation: %v", err)
	}
	return &pb.Void{}, nil
}
