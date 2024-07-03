package comments

import (
	pb "MusicMesh/Collaboration-MusicMesh/generate/comment"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (s *CommentRepo) Create(ctx context.Context, req *pb.Comment) (*pb.Void, error) {
	query := `INSERT INTO comments (composition_id, user_id, content, created_at) VALUES ($1, $2, $3, $4)`
	_, err := s.DB.ExecContext(ctx, query, req.CompositionId, req.UserId, req.Content, time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, fmt.Errorf("failed to create comment: %v", err)
	}
	return &pb.Void{}, nil
}

func (s *CommentRepo) GetById(ctx context.Context, req *pb.CommentID) (*pb.CommentRes, error) {
	query := `SELECT id, composition_id, user_id, content, created_at, updated_at FROM comments WHERE id = $1`
	comment := &pb.CommentRes{}
	err := s.DB.QueryRowContext(ctx, query, req.CompositionId).Scan(&comment.Id, &comment.CompositionId, &comment.UserId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("comment with ID %s not found", req.CompositionId)
	} else if err != nil {
		return nil, fmt.Errorf("failed to retrieve comment: %v", err)
	}
	return comment, nil
}

func (s *CommentRepo) Delete(ctx context.Context, req *pb.CommentID) (*pb.Void, error) {
	_, err := s.DB.ExecContext(ctx, `UPDATE comments SET 
		deleted_at = date_part('epoch', current_timestamp)::INT 
		WHERE id = $1 AND deleted_at = 0`, req.CompositionId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete comment: %v", err)
	}
	return &pb.Void{}, nil
}
