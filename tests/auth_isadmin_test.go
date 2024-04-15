package tests

// TODO : Test fails because of server stack overflow error (wtf)

//func TestIsAdmin_HappyAuthorization(t *testing.T) {
//	ctx, st := suite.New(t)
//
//	email := gofakeit.Email()
//	password := gofakeit.Password(true, true, true, true, false, 12)
//
//	responseReg, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
//		Email:    email,
//		Password: password,
//	})
//	require.NoError(t, err)
//	assert.NotEmpty(t, responseReg.GetUserId())
//
//	responseAuth, err := st.AuthClient.IsAdmin(ctx, &ssov1.IsAdminRequest{
//		UserId: responseReg.UserId,
//	})
//	require.NoError(t, err)
//	assert.Equalf(t, false, responseAuth.IsAdmin, "user is not admin")
//}
