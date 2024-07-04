package comments

import (
	pb "MusicMesh/Collaboration-MusicMesh/generate/comment"
	"database/sql"
)

type CommentRepo struct {
	pb.UnimplementedCommentServiceServer
	DB *sql.DB
}

func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{DB: db}
}
