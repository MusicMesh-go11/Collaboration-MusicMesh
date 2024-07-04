package collaborations

import (
	pb "MusicMesh/Collaboration-MusicMesh/generate/collaboration"
	"context"
	"database/sql"
	"fmt"
)

func (s *CollaborationRepo) Create(ctx context.Context, req *pb.Collaboration) (*pb.Void, error) {
	query := `INSERT INTO collaborations (composition_id, user_id, role) VALUES ($1, $2, $3)`
	_, err := s.DB.ExecContext(ctx, query, req.CompositionId, req.UserId, req.Role)
	if err != nil {
		return nil, fmt.Errorf("failed to create collaboration: %v", err)
	}
	return &pb.Void{}, nil
}

func (s *CollaborationRepo) GetById(ctx context.Context, req *pb.CollaborationID) (*pb.CollaborationRes, error) {
	query := `SELECT collaborations_id, composition_id, user_id, role, joined_at FROM collaborations WHERE collaborations_id = $1`
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

func (s *CollaborationRepo) Update(ctx context.Context, in *pb.CollaborationRes) (*pb.Void, error) {
	_, err := s.DB.Exec("UPDATE collaborations SET composition_id = $1, user_id = $2, role = $3 WHERE collaborations_id = $4",
		in.CompositionId, in.UserId, in.Role, in.Id)
	return &pb.Void{}, err
}
