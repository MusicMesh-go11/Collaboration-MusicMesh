package main

import (
	"MusicMesh/Collaboration-MusicMesh/generate/collaboration"
	comment2 "MusicMesh/Collaboration-MusicMesh/generate/comment"
	"MusicMesh/Collaboration-MusicMesh/generate/invitation"
	"MusicMesh/Collaboration-MusicMesh/storage/collaborations"
	"MusicMesh/Collaboration-MusicMesh/storage/comments"
	"MusicMesh/Collaboration-MusicMesh/storage/invitations"
	"MusicMesh/Collaboration-MusicMesh/storage/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listner, err := net.Listen("tcp", ":5053")
	if err != nil {
		panic(err)
	}
	defer listner.Close()

	db, err := postgres.Conn()
	if err != nil {
		panic(err)
	}

	collab := collaborations.NewCollaborationRepo(db)
	comment := comments.NewCommentRepo(db)
	invit := invitations.NewInvitationRepo(db)
	grpcServer := grpc.NewServer()

	collaboration.RegisterCollaborationServiceServer(grpcServer, collab)
	comment2.RegisterCommentServiceServer(grpcServer, comment)
	invitation.RegisterInvitationServiceServer(grpcServer, invit)

	log.Println("Listening on :5053...")
	panic(grpcServer.Serve(listner))
}
