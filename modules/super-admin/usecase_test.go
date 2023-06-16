package super_admin

import (
	"errors"
	"reflect"
	"testing"
	"user-management-backend/entities"
	"user-management-backend/repositories/mocks"
)

func TestUseCase_GetActiveAdmins(t *testing.T) {
	type fields struct {
		superAdminRepo mocks.SuperAdminRepoInterface
		adminRepo      mocks.AdminRepoInterface
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(sar *mocks.SuperAdminRepoInterface)
		want     []ActorParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct received",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface) {
				sar.On("GetActiveAdmins").Return(
					nil,
					errors.New("err GetActiveAdmins"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface) {
				sar.On("GetActiveAdmins").Return([]entities.Actor{
					{
						Id:         1,
						Username:   "super-admin",
						Password:   "super-admin",
						RoleId:     1,
						IsVerified: "true",
						IsActive:   "true",
					},
					{
						Id:         2,
						Username:   "akbar",
						Password:   "akbar",
						RoleId:     1,
						IsVerified: "false",
						IsActive:   "true",
					},
				}, nil)
			},
			want: []ActorParam{
				{
					Id:         1,
					Username:   "super-admin",
					RoleId:     1,
					IsVerified: "true",
					IsActive:   "true",
				},
				{
					Id:         2,
					Username:   "akbar",
					RoleId:     1,
					IsVerified: "false",
					IsActive:   "true",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.superAdminRepo)
			uc := UseCase{
				sar: &tt.fields.superAdminRepo,
				ar:  &tt.fields.adminRepo,
			}
			got, err := uc.GetActiveAdmins()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetActiveAdmins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetActiveAdmins() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetApprovedAdmins(t *testing.T) {
	type fields struct {
		superAdminRepo mocks.SuperAdminRepoInterface
		adminRepo      mocks.AdminRepoInterface
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(sar *mocks.SuperAdminRepoInterface)
		want     []RegisterApprovalParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct received",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface) {
				sar.On("GetApprovedAdmins").Return(
					nil,
					errors.New("err GetApprovedAdmins"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface) {
				sar.On("GetApprovedAdmins").Return([]entities.RegisterApproval{
					{
						Id:           1,
						AdminId:      2,
						SuperAdminId: 1,
						Status:       "approved",
					},
					{
						Id:           2,
						AdminId:      3,
						SuperAdminId: 1,
						Status:       "approved",
					},
				}, nil)
			},
			want: []RegisterApprovalParam{
				{
					Id:           1,
					AdminId:      2,
					SuperAdminId: 1,
					Status:       "approved",
				},
				{
					Id:           2,
					AdminId:      3,
					SuperAdminId: 1,
					Status:       "approved",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.superAdminRepo)
			uc := UseCase{
				sar: &tt.fields.superAdminRepo,
				ar:  &tt.fields.adminRepo,
			}
			got, err := uc.GetApprovedAdmins()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApprovedAdmins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApprovedAdmins() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetPendingAdmins(t *testing.T) {
	type fields struct {
		superAdminRepo mocks.SuperAdminRepoInterface
		adminRepo      mocks.AdminRepoInterface
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(sar *mocks.SuperAdminRepoInterface)
		want     []RegisterApprovalParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct received",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface) {
				sar.On("GetPendingAdmins").Return(
					nil,
					errors.New("err GetPendingAdmins"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface) {
				sar.On("GetPendingAdmins").Return([]entities.RegisterApproval{
					{
						Id:           1,
						AdminId:      2,
						SuperAdminId: 1,
						Status:       "pending",
					},
					{
						Id:           2,
						AdminId:      3,
						SuperAdminId: 1,
						Status:       "pending",
					},
				}, nil)
			},
			want: []RegisterApprovalParam{
				{
					Id:           1,
					AdminId:      2,
					SuperAdminId: 1,
					Status:       "pending",
				},
				{
					Id:           2,
					AdminId:      3,
					SuperAdminId: 1,
					Status:       "pending",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.superAdminRepo)
			uc := UseCase{
				sar: &tt.fields.superAdminRepo,
				ar:  &tt.fields.adminRepo,
			}
			got, err := uc.GetPendingAdmins()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPendingAdmins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPendingAdmins() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetRegisterAdminById(t *testing.T) {
	type fields struct {
		superAdminRepo mocks.SuperAdminRepoInterface
		adminRepo      mocks.AdminRepoInterface
	}
	type args struct {
		register *RegisterApprovalParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(sar *mocks.SuperAdminRepoInterface,
			register *RegisterApprovalParam)
		args    args
		want    RegisterApprovalParam
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct received",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface, register *RegisterApprovalParam) {
				sar.On("GetRegisterAdminById", &register.Id).Return(
					entities.RegisterApproval{},
					errors.New("err GetRegisterAdminById"))
			},
			args:    args{register: &RegisterApprovalParam{}},
			want:    RegisterApprovalParam{},
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface, register *RegisterApprovalParam) {
				sar.On("GetRegisterAdminById", &register.Id).Return(
					entities.RegisterApproval{
						Id:           register.Id,
						AdminId:      2,
						SuperAdminId: 1,
						Status:       "pending",
					},
					nil)
			},
			args: args{register: &RegisterApprovalParam{
				Id: 2,
			}},
			want: RegisterApprovalParam{
				Id:           2,
				AdminId:      2,
				SuperAdminId: 1,
				Status:       "pending",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.superAdminRepo, tt.args.register)
			uc := UseCase{
				sar: &tt.fields.superAdminRepo,
				ar:  &tt.fields.adminRepo,
			}
			got, err := uc.GetRegisterAdminById(tt.args.register)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRegisterAdminById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRegisterAdminById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetRejectedAdmin(t *testing.T) {
	type fields struct {
		superAdminRepo mocks.SuperAdminRepoInterface
		adminRepo      mocks.AdminRepoInterface
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(sar *mocks.SuperAdminRepoInterface)
		want     []RegisterApprovalParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct received",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface) {
				sar.On("GetRejectedAdmins").Return(
					nil,
					errors.New("err GetRejectedAdmins"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface) {
				sar.On("GetRejectedAdmins").Return([]entities.RegisterApproval{
					{
						Id:           1,
						AdminId:      2,
						SuperAdminId: 1,
						Status:       "rejected",
					},
					{
						Id:           2,
						AdminId:      3,
						SuperAdminId: 1,
						Status:       "rejected",
					},
				}, nil)
			},
			want: []RegisterApprovalParam{
				{
					Id:           1,
					AdminId:      2,
					SuperAdminId: 1,
					Status:       "rejected",
				},
				{
					Id:           2,
					AdminId:      3,
					SuperAdminId: 1,
					Status:       "rejected",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.superAdminRepo)
			uc := UseCase{
				sar: &tt.fields.superAdminRepo,
				ar:  &tt.fields.adminRepo,
			}
			got, err := uc.GetRejectedAdmin()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRejectedAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRejectedAdmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetVerifiedAdmins(t *testing.T) {
	type fields struct {
		superAdminRepo mocks.SuperAdminRepoInterface
		adminRepo      mocks.AdminRepoInterface
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(sar *mocks.SuperAdminRepoInterface)
		want     []ActorParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct received",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface) {
				sar.On("GetVerifiedAdmins").Return(
					nil,
					errors.New("err GetVerifiedAdmins"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface) {
				sar.On("GetVerifiedAdmins").Return(
					[]entities.Actor{
						{
							Id:         1,
							Username:   "akbar",
							Password:   "akbar",
							RoleId:     2,
							IsVerified: "true",
							IsActive:   "false",
						},
						{
							Id:         2,
							Username:   "super-admin",
							Password:   "super-admin",
							RoleId:     1,
							IsVerified: "true",
							IsActive:   "true",
						},
					},
					nil)
			},
			want: []ActorParam{
				{
					Id:         1,
					Username:   "akbar",
					RoleId:     2,
					IsVerified: "true",
					IsActive:   "false",
				},
				{
					Id:         2,
					Username:   "super-admin",
					RoleId:     1,
					IsVerified: "true",
					IsActive:   "true",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.superAdminRepo)
			uc := UseCase{
				sar: &tt.fields.superAdminRepo,
				ar:  &tt.fields.adminRepo,
			}
			got, err := uc.GetVerifiedAdmins()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVerifiedAdmins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVerifiedAdmins() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_ModifyRegisterAdminById(t *testing.T) {
	type fields struct {
		superAdminRepo mocks.SuperAdminRepoInterface
		adminRepo      mocks.AdminRepoInterface
	}
	type args struct {
		register *RegisterApprovalParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(sar *mocks.SuperAdminRepoInterface,
			register *RegisterApprovalParam)
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct received",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface,
				register *RegisterApprovalParam) {
				sar.On("GetRegisterAdminById", &register.Id).Return(
					entities.RegisterApproval{},
					errors.New("err GetRegisterAdminById"))
			},
			args:    args{register: &RegisterApprovalParam{}},
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface,
				register *RegisterApprovalParam) {
				sar.On("GetRegisterAdminById", &register.Id).Return(
					entities.RegisterApproval{
						Id:           1,
						AdminId:      2,
						SuperAdminId: 1,
						Status:       "pending",
					},
					errors.New("err GetRegisterAdminById"))
				sar.On("ModifyRegisterAdminById", &entities.RegisterApproval{
					Id:           1,
					AdminId:      2,
					SuperAdminId: 1,
					Status:       register.Status,
				}).Return(
					nil)
			},
			args: args{register: &RegisterApprovalParam{
				Id:           1,
				AdminId:      2,
				SuperAdminId: 1,
				Status:       "approved",
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.superAdminRepo, tt.args.register)
			uc := UseCase{
				sar: &tt.fields.superAdminRepo,
				ar:  &tt.fields.adminRepo,
			}
			if err := uc.ModifyRegisterAdminById(tt.args.register); (err != nil) != tt.wantErr {
				t.Errorf("ModifyRegisterAdminById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_ModifyStatusAdminById(t *testing.T) {
	type fields struct {
		superAdminRepo mocks.SuperAdminRepoInterface
		adminRepo      mocks.AdminRepoInterface
	}
	type args struct {
		actor *ActorParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(ar *mocks.AdminRepoInterface, actor *ActorParam)
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct received",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, actor *ActorParam) {
				ar.On("GetAdminById", &actor.Id).Return(
					entities.Actor{},
					errors.New("err GetAdminById"))
			},
			args:    args{actor: &ActorParam{}},
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, actor *ActorParam) {
				ar.On("GetAdminById", &actor.Id).Return(
					entities.Actor{
						Id:         actor.Id,
						Username:   "akbar",
						Password:   "maulana",
						RoleId:     1,
						IsVerified: "true",
						IsActive:   "true",
					},
					nil)
				ar.On("ModifyAdmin", &entities.Actor{
					Id:         actor.Id,
					Username:   "akbar",
					Password:   "maulana",
					RoleId:     1,
					IsVerified: "true",
					IsActive:   "true",
				}).Return(
					nil)
			},
			args: args{actor: &ActorParam{
				Id:         1,
				Username:   "akbar",
				RoleId:     1,
				IsVerified: "true",
				IsActive:   "true",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.adminRepo, tt.args.actor)
			uc := UseCase{
				sar: &tt.fields.superAdminRepo,
				ar:  &tt.fields.adminRepo,
			}
			if err := uc.ModifyStatusAdminById(tt.args.actor); (err != nil) != tt.wantErr {
				t.Errorf("ModifyStatusAdminById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_RemoveAdminById(t *testing.T) {
	type fields struct {
		superAdminRepo mocks.SuperAdminRepoInterface
		adminRepo      mocks.AdminRepoInterface
	}
	type args struct {
		admin *ActorParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(sar *mocks.SuperAdminRepoInterface,
			ar *mocks.AdminRepoInterface,
			actor *ActorParam)
		args    args
		want    ActorParam
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct received",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface, ar *mocks.AdminRepoInterface, actor *ActorParam) {
				ar.On("GetAdminById", &actor.Id).Return(
					entities.Actor{},
					errors.New("err GetAdminById"))
			},
			args:    args{admin: &ActorParam{}},
			want:    ActorParam{},
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface, ar *mocks.AdminRepoInterface, actor *ActorParam) {
				ar.On("GetAdminById", &actor.Id).Return(
					entities.Actor{
						Id:         1,
						Username:   "akbar",
						Password:   "maulana",
						RoleId:     1,
						IsVerified: "true",
						IsActive:   "true",
					},
					nil)
				sar.On("RemoveAdminById", &actor.Id).Return(
					nil)
			},
			args: args{admin: &ActorParam{
				Id: 1,
			}},
			want: ActorParam{
				Id:         1,
				Username:   "akbar",
				RoleId:     1,
				IsVerified: "true",
				IsActive:   "true",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.superAdminRepo, &tt.fields.adminRepo, tt.args.admin)
			uc := UseCase{
				sar: &tt.fields.superAdminRepo,
				ar:  &tt.fields.adminRepo,
			}
			got, err := uc.RemoveAdminById(tt.args.admin)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveAdminById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAdminById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_RemoveRegisterAdminById(t *testing.T) {
	type fields struct {
		superAdminRepo mocks.SuperAdminRepoInterface
		adminRepo      mocks.AdminRepoInterface
	}
	type args struct {
		register *RegisterApprovalParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(sar *mocks.SuperAdminRepoInterface,
			register *RegisterApprovalParam)
		args    args
		want    RegisterApprovalParam
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct received",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface,
				register *RegisterApprovalParam) {
				sar.On("GetRegisterAdminById", &register.Id).Return(
					entities.RegisterApproval{},
					errors.New("err GetRegisterAdminById"))
			},
			args:    args{register: &RegisterApprovalParam{}},
			want:    RegisterApprovalParam{},
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				superAdminRepo: *mocks.NewSuperAdminRepoInterface(t),
				adminRepo:      *mocks.NewAdminRepoInterface(t),
			},
			mockRepo: func(sar *mocks.SuperAdminRepoInterface,
				register *RegisterApprovalParam) {
				sar.On("GetRegisterAdminById", &register.Id).Return(
					entities.RegisterApproval{
						Id:           1,
						AdminId:      2,
						SuperAdminId: 1,
						Status:       "approved",
					},
					nil)
				sar.On("RemoveRegisterAdminById", &register.Id).Return(
					nil)
			},
			args: args{register: &RegisterApprovalParam{Id: 1}},
			want: RegisterApprovalParam{
				Id:           1,
				AdminId:      2,
				SuperAdminId: 1,
				Status:       "approved",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.superAdminRepo, tt.args.register)
			uc := UseCase{
				sar: &tt.fields.superAdminRepo,
				ar:  &tt.fields.adminRepo,
			}
			got, err := uc.RemoveRegisterAdminById(tt.args.register)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveRegisterAdminById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveRegisterAdminById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
