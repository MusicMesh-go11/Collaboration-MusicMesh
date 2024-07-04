package invitations

//func TestInvitationRepo_Create(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db: %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository with the test database connection
//	repo := &InvitationRepo{DB: db}
//
//	// Create a test invitation
//	testInvitation := &pb.Invitation{
//		CompositionId: "362ff0b6-8b9b-499e-883a-fbc3b25897d8",
//		InviterId:     "362ff0b6-8b9b-499e-883a-fbc3b25897d8",
//		InviteeId:     "362ff0b6-8b9b-499e-883a-fbc3b25897d8",
//		Status:        "accepted",
//	}
//
//	// Call the Create method
//	_, err = repo.Create(context.Background(), testInvitation)
//	assert.NoError(t, err)
//
//	// Verify the invitation was created in the database
//	var count int
//	err = db.QueryRow("SELECT COUNT(*) FROM invitations WHERE composition_id = $1 AND inviter_id = $2 AND invitee_id = $3",
//		testInvitation.CompositionId, testInvitation.InviterId, testInvitation.InviteeId).Scan(&count)
//	if err != nil {
//		t.Fatalf("Failed to query invitations table: %v", err)
//	}
//
//	assert.Equal(t, 1, count)
//}
//
//func TestInvitationRepo_GetById(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db: %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository with the test database connection
//	repo := &InvitationRepo{DB: db}
//
//	// Insert a test invitation into the invitations table
//	invitationID := "39e84b6a-e64a-4c1c-97f5-f21ce7668bed"
//	compositionID := "362ff0b6-8b9b-499e-883a-fbc3b25897d8"
//	inviterID := "362ff0b6-8b9b-499e-883a-fbc3b25897d8"
//	inviteeID := "362ff0b6-8b9b-499e-883a-fbc3b25897d8"
//	status := "accepted"
//
//	_, err = db.Exec(`INSERT INTO invitations (composition_id, inviter_id, invitee_id, status)
//		VALUES ($1, $2, $3, $4) ON CONFLICT (invitations_id) DO NOTHING`,
//		compositionID, inviterID, inviteeID, status)
//	if err != nil {
//		t.Fatalf("Failed to insert test data: %v", err)
//	}
//
//	// Create an InvitationID request
//	invitationIDReq := pb.InvitationID{Id: invitationID}
//
//	// Call the GetById method
//	invitationRes, err := repo.GetById(context.Background(), &invitationIDReq)
//
//	// Assert the GetById method returned no error
//	assert.NoError(t, err)
//
//	// Verify the details of the invitation
//	assert.Equal(t, invitationID, invitationRes.Id)
//	assert.Equal(t, compositionID, invitationRes.CompositionId)
//	assert.Equal(t, inviterID, invitationRes.InviterId)
//	assert.Equal(t, inviteeID, invitationRes.InviteeId)
//	assert.Equal(t, status, invitationRes.Status)
//}
//
//func TestInvitationRepo_Delete(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db: %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository with the test database connection
//	repo := &InvitationRepo{DB: db}
//
//	// Insert a test invitation into the invitations table
//	invitationID := "362ff0b6-8b9b-499e-883a-fbc3b25897d8"
//	compositionID := "362ff0b6-8b9b-499e-883a-fbc3b25897d8"
//	inviterID := "362ff0b6-8b9b-499e-883a-fbc3b25897d8"
//	inviteeID := "362ff0b6-8b9b-499e-883a-fbc3b25897d8"
//	status := "declined"
//
//	_, err = db.Exec(`INSERT INTO invitations (invitations_id, composition_id, inviter_id, invitee_id, status)
//		VALUES ($1, $2, $3, $4, $5)
//		ON CONFLICT (invitations_id) DO NOTHING`,
//		invitationID, compositionID, inviterID, inviteeID, status)
//	if err != nil {
//		t.Fatalf("Failed to insert test data: %v", err)
//	}
//
//	// Create an InvitationID request
//	invitationIDReq := pb.InvitationID{Id: invitationID}
//
//	// Call the Delete method
//	_, err = repo.Delete(context.Background(), &invitationIDReq)
//
//	// Assert the Delete method returned no error
//	assert.NoError(t, err)
//
//	// Verify the invitation was soft deleted in the database
//	var deletedAt int64
//	err = db.QueryRow("SELECT deleted_at FROM invitations WHERE invitations_id = $1", invitationID).Scan(&deletedAt)
//	if err != nil {
//		t.Fatalf("Failed to query invitations table: %v", err)
//	}
//
//	assert.NotEqual(t, 0, deletedAt)
//}
