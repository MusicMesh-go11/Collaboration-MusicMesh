package collaborations

//func TestCollaborationRepo_Create(t *testing.T) {
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db: %v", err)
//	}
//	defer db.Close()
//
//	repo := &CollaborationRepo{DB: db}
//
//	testCollaboration := &pb.Collaboration{
//		CompositionId: "362ff0b6-8b9b-499e-883a-fbc3b25897d8",
//		UserId:        "362ff0b6-8b9b-499e-883a-fbc3b25897d8",
//		Role:          "owner",
//		JoinedAt:      "2024-07-04T15:04:05Z",
//	}
//
//	_, err = repo.Create(context.Background(), testCollaboration)
//	assert.NoError(t, err)
//
//	var count int
//	err = db.QueryRow("SELECT COUNT(*) FROM collaborations WHERE composition_id = $1 AND user_id = $2", testCollaboration.CompositionId, testCollaboration.UserId).Scan(&count)
//	if err != nil {
//		t.Fatalf("Failed to query collaborations table: %v", err)
//	}
//
//	assert.Equal(t, 1, count)
//}

//	func TestUserRepo_Get(t *testing.T) {
//		// Connect to the test database
//		db, err := postgres.Conn()
//		if err != nil {
//			t.Fatalf("Failed to open sql db, %v", err)
//		}
//		defer db.Close()
//
//		// Initialize the repository with the test database connection
//		repo := &UserRepo{DB: db}
//
//		//// Insert a test user into the users table if not already present
//		//userID := "0313bb44-dac6-4a9f-9e37-865419f93adb"
//		//_, err = db.Exec(`INSERT INTO users (user_id, user_name, email, password, created_at, updated_at, deleted_at)
//		//	VALUES ($1, 'test_user', 'test@example.com', 'password', now(), now(), 0)
//		//	ON CONFLICT (user_id) DO NOTHING`, userID)
//		//if err != nil {
//		//	t.Fatalf("Failed to insert test data, %v", err)
//		//}
//
//		// Define the filter request to fetch all users
//		filterRequest := &pb.FilterRequest{
//			Query: "SELECT user_id, user_name, email, password, created_at FROM users",
//			Arr:   nil, // No parameters needed for this query
//		}
//
//		// Call the Get method
//		usersRes, err := repo.Get(context.Background(), filterRequest)
//
//		// Assert the Get method returned no error
//		assert.NoError(t, err)
//
//		// Verify the number of users returned
//		assert.Equal(t, 1, len(usersRes.Users))
//
//		// Verify the details of the user
//		assert.Equal(t, "0313bb44-dac6-4a9f-9e37-865419f93adb", usersRes.Users[0].UserID)
//		assert.Equal(t, "test_user", usersRes.Users[0].UserName)
//		assert.Equal(t, "test@example.com", usersRes.Users[0].Email)
//	}

//func TestCollaborationRepo_GetById(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db: %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository with the test database connection
//	repo := &CollaborationRepo{DB: db}
//
//	// Insert a test collaboration into the collaborations table
//	collaborationID := "0313bb44-dac6-4a9f-9e37-865419f93adb"
//	compositionID := "362ff0b6-8b9b-499e-883a-fbc3b25897d8"
//	userID := "362ff0b6-8b9b-499e-883a-fbc3b25897d8"
//	role := "owner"
//
//	_, err = db.Exec(`INSERT INTO collaborations (collaborations_id, composition_id, user_id, role)
//		VALUES ($1, $2, $3, $4)
//		ON CONFLICT (collaborations_id) DO NOTHING`, collaborationID, compositionID, userID, role)
//	if err != nil {
//		t.Fatalf("Failed to insert test data: %v", err)
//	}
//
//	// Create a CollaborationID request
//	collaborationIDReq := pb.CollaborationID{CollaborationId: collaborationID}
//
//	// Call the GetById method
//	collaborationRes, err := repo.GetById(context.Background(), &collaborationIDReq)
//
//	// Assert the GetById method returned no error
//	assert.NoError(t, err)
//
//	// Verify the details of the collaboration
//	assert.Equal(t, collaborationID, collaborationRes.Id)
//	assert.Equal(t, compositionID, collaborationRes.CompositionId)
//	assert.Equal(t, userID, collaborationRes.UserId)
//	assert.Equal(t, role, collaborationRes.Role)
//}

//func TestCollaborationRepo_Update(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db: %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository with the test database connection
//	repo := &CollaborationRepo{DB: db}
//
//	// Insert a test collaboration into the collaborations table if not already present
//	collaborationID := "0313bb44-dac6-4a9f-9e37-865419f93adb"
//	compositionID := "362ff0b6-8b9b-499e-883a-fbc3b25897d8"
//	userID := "362ff0b6-8b9b-499e-883a-fbc3b25897d8"
//	role := "owner"
//
//	_, err = db.Exec(`INSERT INTO collaborations (collaborations_id, composition_id, user_id, role)
//		VALUES ($1, $2, $3, $4)
//		ON CONFLICT (collaborations_id) DO NOTHING`, collaborationID, compositionID, userID, role)
//	if err != nil {
//		t.Fatalf("Failed to insert test data: %v", err)
//	}
//
//	// Define the updated collaboration data
//	updatedCollaboration := &pb.CollaborationRes{
//		Id:            collaborationID,
//		CompositionId: "362ff0b6-8b9b-499e-883a-fbc3b25897d8",
//		UserId:        "362ff0b6-8b9b-499e-883a-fbc3b25897d8",
//		Role:          "viewer",
//	}
//
//	// Call the Update method
//	_, err = repo.Update(context.Background(), updatedCollaboration)
//
//	// Assert the Update method returned no error
//	assert.NoError(t, err)
//
//	// Verify the collaboration details were updated correctly by querying the database
//	var compositionIDUpdated, userIDUpdated, roleUpdated string
//	err = db.QueryRow("SELECT composition_id, user_id, role FROM collaborations WHERE collaborations_id = $1", collaborationID).Scan(&compositionIDUpdated, &userIDUpdated, &roleUpdated)
//	if err != nil {
//		t.Fatalf("Failed to query collaborations table: %v", err)
//	}
//
//	// Assert that the collaboration details were updated correctly
//	assert.Equal(t, "362ff0b6-8b9b-499e-883a-fbc3b25897d8", compositionIDUpdated)
//	assert.Equal(t, "362ff0b6-8b9b-499e-883a-fbc3b25897d8", userIDUpdated)
//	assert.Equal(t, "viewer", roleUpdated)
//}
