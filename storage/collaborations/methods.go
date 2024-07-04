package collaborations

import (
	pb "MusicMesh/Collaboration-MusicMesh/generate/collaboration"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (s *CollaborationRepo) Create(ctx context.Context, req *pb.Collaboration) (*pb.Void, error) {
	query := `INSERT INTO collaborations (composition_id, user_id, role, joined_at) VALUES ($1, $2, $3, $4)`
	_, err := s.DB.ExecContext(ctx, query, req.CompositionId, req.UserId, req.Role, time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, fmt.Errorf("failed to create collaboration: %v", err)
	}
	return &pb.Void{}, nil
}

func (s *CollaborationRepo) GetById(ctx context.Context, req *pb.CollaborationID) (*pb.CollaborationRes, error) {
	query := `SELECT collaboration_id, composition_id, user_id, role, joined_at FROM collaborations WHERE collaboration_id = $1`
	collab := &pb.CollaborationRes{}
	err := s.DB.QueryRowContext(ctx, query, req.CollaborationId).Scan(&collab.Id, &collab.CompositionId, &collab.UserId, &collab.Role, &collab.JoinedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("collaboration with ID %s not found", req.CollaborationId)
	} else if err != nil {
		return nil, fmt.Errorf("failed to retrieve collaboration: %v", err)
	}
	return collab, nil
}

func (s *CollaborationRepo) GetByCompositionId(ctx context.Context, req *pb.CompositionID) (*pb.CollaborationResList, error) {
	query := `SELECT collaboration_id, composition_id, user_id, role, joined_at FROM collaborations WHERE composition_id = $1`
	rows, err := s.DB.QueryContext(ctx, query, req.CompositionId)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve collaborations: %v", err)
	}
	defer rows.Close()

	var collaborations []*pb.CollaborationRes
	for rows.Next() {
		var collab pb.CollaborationRes
		err := rows.Scan(&collab.Id, &collab.CompositionId, &collab.UserId, &collab.Role, &collab.JoinedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan collaboration: %v", err)
		}
		collaborations = append(collaborations, &collab)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to process collaborations: %v", err)
	}

	return &pb.CollaborationResList{Collaborations: collaborations}, nil
}

func (s *CollaborationRepo) Delete(ctx context.Context, req *pb.CollaborationID) (*pb.Void, error) {
	_, err := s.DB.ExecContext(ctx, `UPDATE collaborations SET 
		deleted_at = date_part('epoch', current_timestamp)::INT 
		WHERE collaboration_id = $1 AND deleted_at = 0`, req.CollaborationId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete collaboration: %v", err)
	}
	return &pb.Void{}, nil
}
