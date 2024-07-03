package comments

import (
	pb "MusicMesh/Collaboration-MusicMesh/generate/comment"
	"database/sql"
	"fmt"
	"context"
	"errors"
	"time"
	
)

type CommentRepo struct {
	pb.UnimplementedCommentServiceServer
	DB *sql.DB
}

func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{DB: db}
}

func (s *CommentServiceServer) Create(ctx context.Context, req *pb.Comment) (*pb.Void, error) {
    query := `INSERT INTO comments (composition_id, user_id, content, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5)`
    now := time.Now()
    _, err := s.db.ExecContext(ctx, query, req.CompositionId, req.UserId, req.Content, now, now)
    if err != nil {
        return nil, errors.New(fmt.Sprintf("failed to create comment: %v", err))
    }
    return &pb.Void{}, nil
}

func (s *CommentServiceServer) GetById(ctx context.Context, req *pb.CommentID) (*pb.CommentRes, error) {
    query := `SELECT id, composition_id, user_id, content, created_at, updated_at FROM comments WHERE id = $1`
    row := s.db.QueryRowContext(ctx, query, req.CompositionId)

    var comment pb.CommentRes
    err := row.Scan(&comment.Id, &comment.CompositionId, &comment.UserId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
    if err == sql.ErrNoRows {
        return nil, errors.New(fmt.Sprintf("comment with ID %s not found", req.CompositionId))
    } else if err != nil {
        return nil, errors.New(fmt.Sprintf("failed to retrieve comment: %v", err))
    }

    return &comment, nil
}

func (s *CommentServiceServer) Delete(ctx context.Context, req *pb.CommentID) (*pb.Void, error) {
    query := `DELETE FROM comments WHERE id = $1`
    _, err := s.db.ExecContext(ctx, query, req.CompositionId)
    if err != nil {
        return nil, errors.New(fmt.Sprintf("failed to delete comment: %v", err))
    }
    return &pb.Void{}, nil
}