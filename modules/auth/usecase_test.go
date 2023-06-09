package auth

import (
	"back-end-server-dev/entities"
	"back-end-server-dev/repositories/mocks"
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestUseCase_CreateActorSession(t *testing.T) {
	type fields struct {
		authRepo mocks.AuthRepoInterface
	}
	type args struct {
		account *ActorSessionParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(ar *mocks.AuthRepoInterface, data *ActorSessionParam)
		args     args
		want     error
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name:   "Error: check when empty struct passed",
			fields: fields{authRepo: *mocks.NewAuthRepoInterface(t)},
			mockRepo: func(ar *mocks.AuthRepoInterface, data *ActorSessionParam) {
				ar.On("CreateActorSession", &entities.ActorSession{
					CreatedAt: time.Now(),
					ExpiresAt: time.Now().Add(time.Hour * 1),
				}).Return(
					errors.New("err CreateActorSession"))
			},
			args:    args{account: &ActorSessionParam{}},
			want:    errors.New("err CreateActorSession"),
			wantErr: true,
		},
		{
			name:   "Success: correct param passed",
			fields: fields{authRepo: *mocks.NewAuthRepoInterface(t)},
			mockRepo: func(ar *mocks.AuthRepoInterface, data *ActorSessionParam) {
				ar.On("CreateActorSession", &entities.ActorSession{
					Id:        1,
					UserId:    1,
					Token:     "TokenJWT",
					CreatedAt: time.Now(),
					ExpiresAt: time.Now().Add(time.Hour * 1),
				}).Return(nil)
			},
			args: args{account: &ActorSessionParam{
				Id:      1,
				ActorId: 1,
				Token:   "TokenJWT",
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.authRepo, tt.args.account)
			uc := UseCase{
				authRepo: &tt.fields.authRepo,
			}
			if err := uc.CreateActorSession(tt.args.account); (err != nil) != tt.wantErr {
				t.Errorf("CreateActorSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_GetCredentialByUsername(t *testing.T) {
	type fields struct {
		authRepo mocks.AuthRepoInterface
	}
	type args struct {
		account *CredentialParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(ar *mocks.AuthRepoInterface, data *CredentialParam)
		args     args
		want     CredentialParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name:   "Error: check when empty struct passed",
			fields: fields{authRepo: *mocks.NewAuthRepoInterface(t)},
			mockRepo: func(ar *mocks.AuthRepoInterface, data *CredentialParam) {
				ar.On("GetActorByUsername", &data.username).Return(
					entities.Actor{},
					errors.New("err GetActorByUsername"))
			},
			args:    args{account: &CredentialParam{}},
			want:    CredentialParam{},
			wantErr: true,
		},
		{
			name:   "Success: correct param passed",
			fields: fields{authRepo: *mocks.NewAuthRepoInterface(t)},
			mockRepo: func(ar *mocks.AuthRepoInterface, data *CredentialParam) {
				ar.On("GetActorByUsername", &data.username).Return(
					entities.Actor{
						Id:         1,
						Username:   data.username,
						Password:   "super-admin",
						RoleId:     1,
						IsVerified: "true",
						IsActive:   "true",
						CreatedAt:  time.Now(),
						ModifiedAt: time.Now(),
					},
					nil)
			},
			args: args{account: &CredentialParam{
				username: "super-admin",
			}},
			want: CredentialParam{
				id:       1,
				username: "super-admin",
				password: "super-admin",
				roleId:   1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.authRepo, tt.args.account)
			uc := UseCase{
				authRepo: &tt.fields.authRepo,
			}
			got, err := uc.GetCredentialByUsername(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCredentialByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCredentialByUsername() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetLastActorSessionByToken(t *testing.T) {
	type fields struct {
		authRepo mocks.AuthRepoInterface
	}
	type args struct {
		account *ActorSessionParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(ar *mocks.AuthRepoInterface, data *ActorSessionParam)
		args     args
		want     ActorSessionParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name:   "Error: check when empty struct passed",
			fields: fields{authRepo: *mocks.NewAuthRepoInterface(t)},
			mockRepo: func(ar *mocks.AuthRepoInterface, data *ActorSessionParam) {
				ar.On("GetLastActorSessionByToken", &data.Token).Return(
					entities.ActorSession{},
					errors.New("err GetLastActorSessionByToken"))
			},
			args:    args{account: &ActorSessionParam{}},
			want:    ActorSessionParam{},
			wantErr: true,
		},
		{
			name:   "Success: correct param passed",
			fields: fields{authRepo: *mocks.NewAuthRepoInterface(t)},
			mockRepo: func(ar *mocks.AuthRepoInterface, data *ActorSessionParam) {
				ar.On("GetLastActorSessionByToken", &data.Token).Return(
					entities.ActorSession{
						Id:        1,
						UserId:    1,
						Token:     data.Token,
						CreatedAt: time.Now(),
						ExpiresAt: time.Now(),
					},
					nil)
			},
			args: args{account: &ActorSessionParam{
				Token: "TokenJWT",
			}},
			want: ActorSessionParam{
				Id:      1,
				ActorId: 1,
				Token:   "TokenJWT",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.authRepo, tt.args.account)
			uc := UseCase{
				authRepo: &tt.fields.authRepo,
			}
			got, err := uc.GetLastActorSessionByToken(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLastActorSessionByToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLastActorSessionByToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
