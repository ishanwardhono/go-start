package module

import (
	mock_repo "app/database/repo/mock"
	"app/entity"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUserModel_InsertUser(t *testing.T) {
	ctl := gomock.NewController(t)
	mockRepo := mock_repo.NewMockUserRepo(ctl)

	tests := []struct {
		name    string
		user    entity.User
		wantErr bool
		runMock func()
	}{
		{
			name: "test insert user",
			user: entity.User{
				Name:  "ishan",
				Email: "ishan@gmail.com",
			},
			wantErr: false,
			runMock: func() {
				mockRepo.EXPECT().InsertUser(gomock.Any()).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.runMock != nil {
				tt.runMock()
			}
			m := &UserModel{
				repo: mockRepo,
			}
			if err := m.InsertUser(tt.user); (err != nil) != tt.wantErr {
				t.Errorf("UserModel.InsertUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
