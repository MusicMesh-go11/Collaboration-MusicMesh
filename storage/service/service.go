package service

import (
    "context"
    "database/sql"

    pb "github.com/MusicMesh-go11/Collabration_Service/genproto/collabration"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/timestamppb"
)

type CollaborationService struct {
    pb.UnimplementedCollaborationServiceServer
    db *sql.DB
}

func NewCollaborationService(db *sql.DB) *CollaborationService {
    return &CollaborationService{db: db}
}

// Collaboration CRUD

func (s *CollaborationService) CreateCollaboration(ctx context.Context, req *pb.Collaboration) (*pb.Collaboration, error) {
    query := `INSERT INTO collaborations (composition_id, user_id, role) VALUES ($1, $2, $3) RETURNING id, joined_at`
    var id int32
    var joinedAt timestamppb.Timestamp
    err := s.db.QueryRowContext(ctx, query, req.CompositionId, req.UserId, req.Role).Scan(&id, &joinedAt)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to create collaboration: %v", err)
    }
    req.Id = id
    req.JoinedAt = &joinedAt
    return req, nil
}

func (s *CollaborationService) GetCollaboration(ctx context.Context, req *pb.GetCollaborationRequest) (*pb.Collaboration, error) {
    query := `SELECT composition_id, user_id, role, joined_at FROM collaborations WHERE id = $1`
    var collab pb.Collaboration
    var joinedAt timestamppb.Timestamp
    err := s.db.QueryRowContext(ctx, query, req.Id).Scan(&collab.CompositionId, &collab.UserId, &collab.Role, &joinedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, status.Errorf(codes.NotFound, "Collaboration not found")
        }
        return nil, status.Errorf(codes.Internal, "Failed to get collaboration: %v", err)
    }
    collab.Id = req.Id
    collab.JoinedAt = &joinedAt
    return &collab, nil
}

func (s *CollaborationService) UpdateCollaboration(ctx context.Context, req *pb.Collaboration) (*pb.Collaboration, error) {
    query := `UPDATE collaborations SET composition_id = $1, user_id = $2, role = $3 WHERE id = $4`
    _, err := s.db.ExecContext(ctx, query, req.CompositionId, req.UserId, req.Role, req.Id)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to update collaboration: %v", err)
    }
    return req, nil
}

func (s *CollaborationService) DeleteCollaboration(ctx context.Context, req *pb.DeleteCollaborationRequest) (*pb.DeleteCollaborationResponse, error) {
    query := `DELETE FROM collaborations WHERE id = $1`
    _, err := s.db.ExecContext(ctx, query, req.Id)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to delete collaboration: %v", err)
    }
    return &pb.DeleteCollaborationResponse{Success: true}, nil
}

// Invitation CRUD

func (s *CollaborationService) CreateInvitation(ctx context.Context, req *pb.Invitation) (*pb.Invitation, error) {
    query := `INSERT INTO invitations (composition_id, inviter_id, invitee_id, status) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
    var id int32
    var createdAt timestamppb.Timestamp
    err := s.db.QueryRowContext(ctx, query, req.CompositionId, req.InviterId, req.InviteeId, req.Status).Scan(&id, &createdAt)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to create invitation: %v", err)
    }
    req.Id = id
    req.CreatedAt = &createdAt
    return req, nil
}

func (s *CollaborationService) GetInvitation(ctx context.Context, req *pb.GetInvitationRequest) (*pb.Invitation, error) {
    query := `SELECT composition_id, inviter_id, invitee_id, status, created_at FROM invitations WHERE id = $1`
    var inv pb.Invitation
    var createdAt timestamppb.Timestamp
    err := s.db.QueryRowContext(ctx, query, req.Id).Scan(&inv.CompositionId, &inv.InviterId, &inv.InviteeId, &inv.Status, &createdAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, status.Errorf(codes.NotFound, "Invitation not found")
        }
        return nil, status.Errorf(codes.Internal, "Failed to get invitation: %v", err)
    }
    inv.Id = req.Id
    inv.CreatedAt = &createdAt
    return &inv, nil
}

func (s *CollaborationService) UpdateInvitation(ctx context.Context, req *pb.Invitation) (*pb.Invitation, error) {
    query := `UPDATE invitations SET composition_id = $1, inviter_id = $2, invitee_id = $3, status = $4 WHERE id = $5`
    _, err := s.db.ExecContext(ctx, query, req.CompositionId, req.InviterId, req.InviteeId, req.Status, req.Id)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to update invitation: %v", err)
    }
    return req, nil
}

func (s *CollaborationService) DeleteInvitation(ctx context.Context, req *pb.DeleteInvitationRequest) (*pb.DeleteInvitationResponse, error) {
    query := `DELETE FROM invitations WHERE id = $1`
    _, err := s.db.ExecContext(ctx, query, req.Id)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to delete invitation: %v", err)
    }
    return &pb.DeleteInvitationResponse{Success: true}, nil
}

// Comment CRUD

func (s *CollaborationService) CreateComment(ctx context.Context, req *pb.Comment) (*pb.Comment, error) {
    query := `INSERT INTO comments (composition_id, user_id, content) VALUES ($1, $2, $3) RETURNING id, created_at`
    var id int32
    var createdAt timestamppb.Timestamp
    err := s.db.QueryRowContext(ctx, query, req.CompositionId, req.UserId, req.Content).Scan(&id, &createdAt)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to create comment: %v", err)
    }
    req.Id = id
    req.CreatedAt = &createdAt
    return req, nil
}

func (s *CollaborationService) GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.Comment, error) {
    query := `SELECT composition_id, user_id, content, created_at FROM comments WHERE id = $1`
    var comment pb.Comment
    var createdAt timestamppb.Timestamp
    err := s.db.QueryRowContext(ctx, query, req.Id).Scan(&comment.CompositionId, &comment.UserId, &comment.Content, &createdAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, status.Errorf(codes.NotFound, "Comment not found")
        }
        return nil, status.Errorf(codes.Internal, "Failed to get comment: %v", err)
    }
    comment.Id = req.Id
    comment.CreatedAt = &createdAt
    return &comment, nil
}

func (s *CollaborationService) UpdateComment(ctx context.Context, req *pb.Comment) (*pb.Comment, error) {
    query := `UPDATE comments SET composition_id = $1, user_id = $2, content = $3 WHERE id = $4`
    _, err := s.db.ExecContext(ctx, query, req.CompositionId, req.UserId, req.Content, req.Id)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to update comment: %v", err)
    }
    return req, nil
}

func (s *CollaborationService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
    query := `DELETE FROM comments WHERE id = $1`
    _, err := s.db.ExecContext(ctx, query, req.Id)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to delete comment: %v", err)
    }
    return &pb.DeleteCommentResponse{Success: true}, nil
}