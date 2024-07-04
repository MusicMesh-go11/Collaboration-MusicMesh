package comments

//func TestCommentRepo_Create(t *testing.T) {
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db: %v", err)
//	}
//	defer db.Close()
//
//	repo := &CommentRepo{DB: db}
//
//	testComment := &pb.Comment{
//		CompositionId: "362ff0b6-8b9b-499e-883a-fbc3b25897d8",
//		UserId:        "362ff0b6-8b9b-499e-883a-fbc3b25897d8",
//		Content:       "Test content",
//	}
//
//	_, err = repo.Create(context.Background(), testComment)
//	assert.NoError(t, err)
//
//	var count int
//	err = db.QueryRow("SELECT COUNT(*) FROM comments WHERE composition_id = $1 AND user_id = $2", testComment.CompositionId, testComment.UserId).Scan(&count)
//	if err != nil {
//		t.Fatalf("Failed to query comments table: %v", err)
//	}
//
//	assert.Equal(t, 1, count)
//}

//func TestCommentRepo_GetById(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db: %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository with the test database connection
//	repo := &CommentRepo{DB: db}
//
//	// Insert a test comment into the comments table
//	commentID := "7ac5eb0f-e842-476a-859e-05830fc821c0"
//	compositionID := "362ff0b6-8b9b-499e-883a-fbc3b25897d8"
//	userID := "362ff0b6-8b9b-499e-883a-fbc3b25897d8"
//	content := "Test content"
//
//	_, err = db.Exec(`INSERT INTO comments (composition_id, user_id, content)
//VALUES ($1, $2, $3)
//ON CONFLICT (comments_id) DO NOTHING`, compositionID, userID, content)
//	if err != nil {
//		t.Fatalf("Failed to insert test data: %v", err)
//	}
//
//	// Create a CommentID request
//	commentIDReq := pb.CommentID{CommentId: commentID}
//
//	// Call the GetById method
//	commentRes, err := repo.GetById(context.Background(), &commentIDReq)
//
//	// Assert the GetById method returned no error
//	assert.NoError(t, err)
//
//	// Verify the details of the comment
//	assert.Equal(t, commentID, commentRes.Id)
//	assert.Equal(t, compositionID, commentRes.CompositionId)
//	assert.Equal(t, userID, commentRes.UserId)
//	assert.Equal(t, content, commentRes.Content)
//}

//func TestCommentRepo_Delete(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db: %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository with the test database connection
//	repo := &CommentRepo{DB: db}
//
//	// Insert a test comment into the comments table
//	commentID := "362ff0b6-8b9b-499e-883a-fbc3b25897d8"
//	compositionID := "362ff0b6-8b9b-499e-883a-fbc3b25897d8"
//	userID := "362ff0b6-8b9b-499e-883a-fbc3b25897d8"
//	content := "Test content"
//
//	_, err = db.Exec(`INSERT INTO comments (comments_id, composition_id, user_id, content)
//VALUES ($1, $2, $3, $4)
//ON CONFLICT (comments_id) DO NOTHING`, commentID, compositionID, userID, content)
//	if err != nil {
//		t.Fatalf("Failed to insert test data: %v", err)
//	}
//
//	// Create a CommentID request
//	commentIDReq := pb.CommentID{CommentId: commentID}
//
//	// Call the Delete method
//	_, err = repo.Delete(context.Background(), &commentIDReq)
//
//	// Assert the Delete method returned no error
//	assert.NoError(t, err)
//
//	// Verify the comment was soft deleted in the database
//	var deletedAt int64
//	err = db.QueryRow("SELECT deleted_at FROM comments WHERE comments_id = $1", commentID).Scan(&deletedAt)
//	if err != nil {
//		t.Fatalf("Failed to query comments table: %v", err)
//	}
//
//	assert.NotEqual(t, 0, deletedAt)
//}
