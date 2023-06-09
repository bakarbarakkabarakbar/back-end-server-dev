package admin

import (
	"back-end-server-dev/entities"
	"back-end-server-dev/repositories/mocks"
	"crypto/sha1"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestUseCase_CreateAdmin(t *testing.T) {
	type fields struct {
		adminRepo    mocks.AdminRepoInterface
		customerRepo mocks.CustomerRepoInterface
	}
	type args struct {
		admin *ActorParamWithPassword
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(ar *mocks.AdminRepoInterface, data *ActorParamWithPassword)
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, admin *ActorParamWithPassword) {
				var hash = sha1.New()
				hash.Write([]byte(admin.Password))
				ar.On("CreateAdmin", &entities.Actor{
					Id:         admin.Id,
					Username:   admin.Username,
					Password:   fmt.Sprintf("%x", hash.Sum(nil)),
					RoleId:     admin.RoleId,
					IsVerified: "false",
					IsActive:   "false",
					CreatedAt:  time.Now(),
					ModifiedAt: time.Now(),
				}).Return(
					errors.New("err CreateAdmin"))
			},
			args:    args{admin: &ActorParamWithPassword{}},
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, admin *ActorParamWithPassword) {
				var hash = sha1.New()
				hash.Write([]byte(admin.Password))
				ar.On("CreateAdmin", &entities.Actor{
					Id:         admin.Id,
					Username:   admin.Username,
					Password:   fmt.Sprintf("%x", hash.Sum(nil)),
					RoleId:     admin.RoleId,
					IsVerified: "false",
					IsActive:   "false",
					CreatedAt:  time.Now(),
					ModifiedAt: time.Now(),
				}).Return(
					nil)
			},
			args: args{admin: &ActorParamWithPassword{
				Id:       1,
				Username: "akbar",
				Password: "maulana",
				RoleId:   1,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.adminRepo, tt.args.admin)
			uc := UseCase{
				adminRepo:    &tt.fields.adminRepo,
				customerRepo: &tt.fields.customerRepo,
			}
			if err := uc.CreateAdmin(tt.args.admin); (err != nil) != tt.wantErr {
				t.Errorf("CreateAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_CreateCustomer(t *testing.T) {
	type fields struct {
		adminRepo    mocks.AdminRepoInterface
		customerRepo mocks.CustomerRepoInterface
	}
	type args struct {
		customer *CustomerParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(ar *mocks.AdminRepoInterface, customer *CustomerParam)
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, customer *CustomerParam) {
				ar.On("CreateCustomer", &entities.Customer{
					Id:         customer.Id,
					FirstName:  customer.FirstName,
					LastName:   customer.LastName,
					Email:      customer.Email,
					Avatar:     customer.Avatar,
					CreatedAt:  time.Now(),
					ModifiedAt: time.Now(),
				}).Return(errors.New("err CreateCustomer"))
			},
			args:    args{customer: &CustomerParam{}},
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, customer *CustomerParam) {
				ar.On("CreateCustomer", &entities.Customer{
					Id:         customer.Id,
					FirstName:  customer.FirstName,
					LastName:   customer.LastName,
					Email:      customer.Email,
					Avatar:     customer.Avatar,
					CreatedAt:  time.Now(),
					ModifiedAt: time.Now(),
				}).Return(nil)
			},
			args: args{customer: &CustomerParam{
				Id:        1,
				FirstName: "akbar",
				LastName:  "maulana",
				Email:     "akbar@example.com",
				Avatar:    "katara",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.adminRepo, tt.args.customer)
			uc := UseCase{
				adminRepo:    &tt.fields.adminRepo,
				customerRepo: &tt.fields.customerRepo,
			}
			if err := uc.CreateCustomer(tt.args.customer); (err != nil) != tt.wantErr {
				t.Errorf("CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_CreateRegisterAdmin(t *testing.T) {
	type fields struct {
		adminRepo    mocks.AdminRepoInterface
		customerRepo mocks.CustomerRepoInterface
	}
	type args struct {
		admin *RegisterApprovalParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(ar *mocks.AdminRepoInterface, admin *RegisterApprovalParam)
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, admin *RegisterApprovalParam) {
				ar.On("CreateRegisterAdmin", &entities.RegisterApproval{
					Id:           admin.Id,
					AdminId:      admin.AdminId,
					SuperAdminId: admin.SuperAdminId,
					Status:       "pending",
				}).Return(errors.New("err CreateRegisterAdmin"))
			},
			args:    args{admin: &RegisterApprovalParam{}},
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, admin *RegisterApprovalParam) {
				ar.On("CreateRegisterAdmin", &entities.RegisterApproval{
					Id:           admin.Id,
					AdminId:      admin.AdminId,
					SuperAdminId: admin.SuperAdminId,
					Status:       "pending",
				}).Return(nil)
			},
			args: args{admin: &RegisterApprovalParam{
				Id:           1,
				AdminId:      2,
				SuperAdminId: 1,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.adminRepo, tt.args.admin)
			uc := UseCase{
				adminRepo:    &tt.fields.adminRepo,
				customerRepo: &tt.fields.customerRepo,
			}
			if err := uc.CreateRegisterAdmin(tt.args.admin); (err != nil) != tt.wantErr {
				t.Errorf("CreateRegisterAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_GetAdminById(t *testing.T) {
	type fields struct {
		adminRepo    mocks.AdminRepoInterface
		customerRepo mocks.CustomerRepoInterface
	}
	type args struct {
		admin *ActorParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(ar *mocks.AdminRepoInterface, admin *ActorParam)
		args     args
		want     ActorParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, admin *ActorParam) {
				ar.On("GetAdminById", &admin.Id).Return(
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
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, admin *ActorParam) {
				ar.On("GetAdminById", &admin.Id).Return(
					entities.Actor{
						Id:         1,
						Username:   "akbar",
						Password:   "maulana",
						RoleId:     2,
						IsVerified: "true",
						IsActive:   "true",
						CreatedAt:  time.Now(),
						ModifiedAt: time.Now(),
					},
					nil)
			},
			args: args{admin: &ActorParam{Id: 1}},
			want: ActorParam{
				Id:         1,
				Username:   "akbar",
				RoleId:     2,
				IsVerified: "true",
				IsActive:   "true",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.adminRepo, tt.args.admin)
			uc := UseCase{
				adminRepo:    &tt.fields.adminRepo,
				customerRepo: &tt.fields.customerRepo,
			}
			got, err := uc.GetAdminById(tt.args.admin)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdminById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAdminById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetAdminsByUsername(t *testing.T) {
	type fields struct {
		adminRepo    mocks.AdminRepoInterface
		customerRepo mocks.CustomerRepoInterface
	}
	type args struct {
		admin *ActorParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(ar *mocks.AdminRepoInterface, admin *ActorParam)
		args     args
		want     []ActorParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, admin *ActorParam) {
				ar.On("GetAdminsByUsername", &admin.Username).Return(
					[]entities.Actor{},
					errors.New("err GetAdminsByUsername"))
			},
			args:    args{admin: &ActorParam{}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, admin *ActorParam) {
				ar.On("GetAdminsByUsername", &admin.Username).Return(
					[]entities.Actor{
						{
							Id:         1,
							Username:   "akbar",
							Password:   "maulana",
							RoleId:     1,
							IsVerified: "true",
							IsActive:   "true",
							CreatedAt:  time.Now(),
							ModifiedAt: time.Now(),
						},
						{
							Id:         2,
							Username:   "akbar",
							Password:   "maulana",
							RoleId:     2,
							IsVerified: "false",
							IsActive:   "false",
							CreatedAt:  time.Now(),
							ModifiedAt: time.Now(),
						},
					}, nil,
				)

			},
			args: args{admin: &ActorParam{Username: "akbar"}},
			want: []ActorParam{
				{
					Id:         1,
					Username:   "akbar",
					RoleId:     1,
					IsVerified: "true",
					IsActive:   "true",
				},
				{
					Id:         2,
					Username:   "akbar",
					RoleId:     2,
					IsVerified: "false",
					IsActive:   "false",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.adminRepo, tt.args.admin)
			uc := UseCase{
				adminRepo:    &tt.fields.adminRepo,
				customerRepo: &tt.fields.customerRepo,
			}
			got, err := uc.GetAdminsByUsername(tt.args.admin)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdminsByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAdminsByUsername() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetAllAdmins(t *testing.T) {
	type fields struct {
		adminRepo    mocks.AdminRepoInterface
		customerRepo mocks.CustomerRepoInterface
	}
	type args struct {
		page *uint
	}
	var page uint = 1
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(ar *mocks.AdminRepoInterface, page *uint)
		args     args
		want     []ActorParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, page *uint) {
				ar.On("GetAllAdmins", page).Return(
					[]entities.Actor{},
					errors.New("err GetAllAdmins"))
			},
			args:    args{page: nil},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, page *uint) {
				ar.On("GetAllAdmins", page).Return(
					[]entities.Actor{
						{
							Id:         1,
							Username:   "akbar",
							Password:   "maulana",
							RoleId:     1,
							IsVerified: "true",
							IsActive:   "true",
							CreatedAt:  time.Now(),
							ModifiedAt: time.Now(),
						},
						{
							Id:         2,
							Username:   "akbar",
							Password:   "maulana",
							RoleId:     2,
							IsVerified: "false",
							IsActive:   "false",
							CreatedAt:  time.Now(),
							ModifiedAt: time.Now(),
						},
					}, nil,
				)
			},
			args: args{page: &page},
			want: []ActorParam{
				{
					Id:         1,
					Username:   "akbar",
					RoleId:     1,
					IsVerified: "true",
					IsActive:   "true",
				},
				{
					Id:         2,
					Username:   "akbar",
					RoleId:     2,
					IsVerified: "false",
					IsActive:   "false",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.adminRepo, tt.args.page)
			uc := UseCase{
				adminRepo:    &tt.fields.adminRepo,
				customerRepo: &tt.fields.customerRepo,
			}
			got, err := uc.GetAllAdmins(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllAdmins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllAdmins() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetAllCustomers(t *testing.T) {
	type fields struct {
		adminRepo    mocks.AdminRepoInterface
		customerRepo mocks.CustomerRepoInterface
	}
	type args struct {
		page *uint
	}
	var page uint = 1
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(ar *mocks.AdminRepoInterface, page *uint)
		args     args
		want     []CustomerParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, page *uint) {
				ar.On("GetAllCustomers", page).Return(
					[]entities.Customer{}, errors.New("err GetAllCustomers"))
			},
			args:    args{page: &page},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, page *uint) {
				ar.On("GetAllCustomers", page).Return(
					[]entities.Customer{
						{
							Id:         1,
							FirstName:  "akbar",
							LastName:   "maulana",
							Email:      "akbar@example.com",
							Avatar:     "aang",
							CreatedAt:  time.Now(),
							ModifiedAt: time.Now(),
						},
						{
							Id:         2,
							FirstName:  "maulana",
							LastName:   "akbar",
							Email:      "maulana@example.com",
							Avatar:     "zuko",
							CreatedAt:  time.Now(),
							ModifiedAt: time.Now(),
						},
					}, nil)
			},
			args: args{page: &page},
			want: []CustomerParam{
				{
					Id:        1,
					FirstName: "akbar",
					LastName:  "maulana",
					Email:     "akbar@example.com",
					Avatar:    "aang",
				},
				{
					Id:        2,
					FirstName: "maulana",
					LastName:  "akbar",
					Email:     "maulana@example.com",
					Avatar:    "zuko",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.adminRepo, tt.args.page)
			uc := UseCase{
				adminRepo:    &tt.fields.adminRepo,
				customerRepo: &tt.fields.customerRepo,
			}
			got, err := uc.GetAllCustomers(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllCustomers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCustomers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetCustomerById(t *testing.T) {
	type fields struct {
		adminRepo    mocks.AdminRepoInterface
		customerRepo mocks.CustomerRepoInterface
	}
	type args struct {
		customer *CustomerParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(cr *mocks.CustomerRepoInterface, customer *CustomerParam)
		args     args
		want     CustomerParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(cr *mocks.CustomerRepoInterface, customer *CustomerParam) {
				cr.On("GetCustomerById", &customer.Id).Return(
					entities.Customer{}, errors.New("err GetCustomerById"))
			},
			args:    args{customer: &CustomerParam{}},
			want:    CustomerParam{},
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(cr *mocks.CustomerRepoInterface, customer *CustomerParam) {
				cr.On("GetCustomerById", &customer.Id).Return(
					entities.Customer{
						Id:         1,
						FirstName:  "akbar",
						LastName:   "maulana",
						Email:      "akbar@example.com",
						Avatar:     "katara",
						CreatedAt:  time.Now(),
						ModifiedAt: time.Now(),
					}, nil)
			},
			args: args{customer: &CustomerParam{
				Id: 1,
			}},
			want: CustomerParam{
				Id:        1,
				FirstName: "akbar",
				LastName:  "maulana",
				Email:     "akbar@example.com",
				Avatar:    "katara",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.customerRepo, tt.args.customer)
			uc := UseCase{
				adminRepo:    &tt.fields.adminRepo,
				customerRepo: &tt.fields.customerRepo,
			}
			got, err := uc.GetCustomerById(tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomerById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomerById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetCustomersByEmail(t *testing.T) {
	type fields struct {
		adminRepo    mocks.AdminRepoInterface
		customerRepo mocks.CustomerRepoInterface
	}
	type args struct {
		customer *CustomerParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(ar *mocks.AdminRepoInterface, customer *CustomerParam)
		args     args
		want     []CustomerParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, customer *CustomerParam) {
				ar.On("GetCustomersByEmail", customer.Email).Return(
					[]entities.Customer{}, errors.New("err GetCustomersByEmail"))
			},
			args:    args{customer: &CustomerParam{}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(cr *mocks.AdminRepoInterface, customer *CustomerParam) {
				cr.On("GetCustomersByEmail", &customer.Email).Return(
					[]entities.Customer{
						{
							Id:         1,
							FirstName:  "akbar",
							LastName:   "maulana",
							Email:      "akbar@example.com",
							Avatar:     "katara",
							CreatedAt:  time.Now(),
							ModifiedAt: time.Now(),
						},
						{
							Id:         2,
							FirstName:  "maulana",
							LastName:   "akbar",
							Email:      "maulana@example.com",
							Avatar:     "aang",
							CreatedAt:  time.Now(),
							ModifiedAt: time.Now(),
						},
					}, nil)
			},
			args: args{customer: &CustomerParam{
				Email: "@example.com",
			}},
			want: []CustomerParam{
				{
					Id:        1,
					FirstName: "akbar",
					LastName:  "maulana",
					Email:     "akbar@example.com",
					Avatar:    "katara",
				},
				{
					Id:        2,
					FirstName: "maulana",
					LastName:  "akbar",
					Email:     "maulana@example.com",
					Avatar:    "aang",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.adminRepo, tt.args.customer)
			uc := UseCase{
				adminRepo:    &tt.fields.adminRepo,
				customerRepo: &tt.fields.customerRepo,
			}
			got, err := uc.GetCustomersByEmail(tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomersByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomersByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetCustomersByName(t *testing.T) {
	type fields struct {
		adminRepo    mocks.AdminRepoInterface
		customerRepo mocks.CustomerRepoInterface
	}
	type args struct {
		customer *CustomerParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(ar *mocks.AdminRepoInterface, customer *CustomerParam)
		args     args
		want     []CustomerParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, customer *CustomerParam) {
				ar.On("GetCustomersByName", &customer.FirstName).Return(
					[]entities.Customer{}, errors.New("err GetCustomersByName"))
			},
			args:    args{customer: &CustomerParam{}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, customer *CustomerParam) {
				ar.On("GetCustomersByName", &customer.FirstName).Return(
					[]entities.Customer{
						{
							Id:         1,
							FirstName:  "akbar",
							LastName:   "maulana",
							Email:      "akbar@example.com",
							Avatar:     "aang",
							CreatedAt:  time.Time{},
							ModifiedAt: time.Time{},
						},
						{
							Id:         2,
							FirstName:  "maulana",
							LastName:   "akbar",
							Email:      "maulana@example.com",
							Avatar:     "zuko",
							CreatedAt:  time.Time{},
							ModifiedAt: time.Time{},
						},
					}, nil)
			},
			args: args{customer: &CustomerParam{FirstName: "akbar"}},
			want: []CustomerParam{
				{
					Id:        1,
					FirstName: "akbar",
					LastName:  "maulana",
					Email:     "akbar@example.com",
					Avatar:    "aang",
				},
				{
					Id:        2,
					FirstName: "maulana",
					LastName:  "akbar",
					Email:     "maulana@example.com",
					Avatar:    "zuko",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.adminRepo, tt.args.customer)
			uc := UseCase{
				adminRepo:    &tt.fields.adminRepo,
				customerRepo: &tt.fields.customerRepo,
			}
			got, err := uc.GetCustomersByName(tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomersByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomersByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_ModifyAdmin(t *testing.T) {
	type fields struct {
		adminRepo    mocks.AdminRepoInterface
		customerRepo mocks.CustomerRepoInterface
	}
	type args struct {
		admin *ActorParamWithPassword
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(ar *mocks.AdminRepoInterface, admin *ActorParamWithPassword)
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, admin *ActorParamWithPassword) {
				ar.On("GetAdminById", &admin.Id).Return(
					entities.Actor{}, errors.New("err GetAdminById"))
			},
			args:    args{admin: &ActorParamWithPassword{}},
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(ar *mocks.AdminRepoInterface, admin *ActorParamWithPassword) {
				var hash = sha1.New()
				hash.Write([]byte(admin.Password))
				ar.On("GetAdminById", &admin.Id).Return(
					entities.Actor{
						Id:         admin.Id,
						Username:   admin.Username,
						Password:   fmt.Sprintf("%x", hash.Sum(nil)),
						RoleId:     admin.RoleId,
						IsVerified: "true",
						IsActive:   "true",
						CreatedAt:  time.Now(),
						ModifiedAt: time.Now(),
					}, nil)
				ar.On("ModifyAdmin", &entities.Actor{
					Id:         admin.Id,
					Username:   admin.Username,
					Password:   fmt.Sprintf("%x", hash.Sum(nil)),
					RoleId:     admin.RoleId,
					IsVerified: "true",
					IsActive:   "true",
					CreatedAt:  time.Now(),
					ModifiedAt: time.Now(),
				}).Return(nil)
			},
			args: args{admin: &ActorParamWithPassword{
				Id:         1,
				Username:   "akbar",
				Password:   "maulana",
				RoleId:     1,
				IsVerified: "true",
				IsActive:   "true",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.adminRepo, tt.args.admin)
			uc := UseCase{
				adminRepo:    &tt.fields.adminRepo,
				customerRepo: &tt.fields.customerRepo,
			}
			if err := uc.ModifyAdmin(tt.args.admin); (err != nil) != tt.wantErr {
				t.Errorf("ModifyAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_ModifyCustomer(t *testing.T) {
	type fields struct {
		adminRepo    mocks.AdminRepoInterface
		customerRepo mocks.CustomerRepoInterface
	}
	type args struct {
		customer *CustomerParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(cr *mocks.CustomerRepoInterface, ar *mocks.AdminRepoInterface, customer *CustomerParam)
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(cr *mocks.CustomerRepoInterface, ar *mocks.AdminRepoInterface, customer *CustomerParam) {
				cr.On("GetCustomerById", &customer.Id).Return(
					entities.Customer{}, errors.New("err GetCustomerById"))
			},
			args:    args{customer: &CustomerParam{}},
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(cr *mocks.CustomerRepoInterface, ar *mocks.AdminRepoInterface, customer *CustomerParam) {
				cr.On("GetCustomerById", &customer.Id).Return(
					entities.Customer{
						Id:         1,
						FirstName:  "akbar",
						LastName:   "maulana",
						Email:      "akbar@example.com",
						Avatar:     "aang",
						CreatedAt:  time.Now(),
						ModifiedAt: time.Now(),
					}, nil)
				ar.On("ModifyCustomer", &entities.Customer{
					Id:         1,
					FirstName:  "akbar",
					LastName:   "maulana",
					Email:      "akbar@example.com",
					Avatar:     "aang",
					CreatedAt:  time.Now(),
					ModifiedAt: time.Now()}).Return(nil)
			},
			args: args{customer: &CustomerParam{
				Id:        1,
				FirstName: "akbar",
				LastName:  "maulana",
				Email:     "akbar@example.com",
				Avatar:    "aang",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo(&tt.fields.customerRepo, &tt.fields.adminRepo, tt.args.customer)
			uc := UseCase{
				adminRepo:    &tt.fields.adminRepo,
				customerRepo: &tt.fields.customerRepo,
			}
			if err := uc.ModifyCustomer(tt.args.customer); (err != nil) != tt.wantErr {
				t.Errorf("ModifyCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_RemoveCustomerById(t *testing.T) {
	type fields struct {
		adminRepo    mocks.AdminRepoInterface
		customerRepo mocks.CustomerRepoInterface
	}
	type args struct {
		customer *CustomerParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(cr *mocks.CustomerRepoInterface, ar *mocks.AdminRepoInterface, customer *CustomerParam)
		args     args
		want     CustomerParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Error: check when empty struct passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(cr *mocks.CustomerRepoInterface, ar *mocks.AdminRepoInterface, customer *CustomerParam) {
				cr.On("GetCustomerById", &customer.Id).Return(
					entities.Customer{}, errors.New("err GetCustomerById"))
			},
			args:    args{customer: &CustomerParam{}},
			want:    CustomerParam{},
			wantErr: true,
		},
		{
			name: "Success: correct param passed",
			fields: fields{
				adminRepo:    *mocks.NewAdminRepoInterface(t),
				customerRepo: *mocks.NewCustomerRepoInterface(t),
			},
			mockRepo: func(cr *mocks.CustomerRepoInterface, ar *mocks.AdminRepoInterface, customer *CustomerParam) {
				cr.On("GetCustomerById", &customer.Id).Return(
					entities.Customer{
						Id:         1,
						FirstName:  "akbar",
						LastName:   "maulana",
						Email:      "akbar@example.com",
						Avatar:     "aang",
						CreatedAt:  time.Now(),
						ModifiedAt: time.Now(),
					}, nil)
				ar.On("RemoveCustomerById", &customer.Id).Return(
					nil)
			},
			args: args{customer: &CustomerParam{Id: 1}},
			want: CustomerParam{
				Id:        1,
				FirstName: "akbar",
				LastName:  "maulana",
				Email:     "akbar@example.com",
				Avatar:    "aang",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.mockRepo(&tt.fields.customerRepo, &tt.fields.adminRepo, tt.args.customer)
		t.Run(tt.name, func(t *testing.T) {
			uc := UseCase{
				adminRepo:    &tt.fields.adminRepo,
				customerRepo: &tt.fields.customerRepo,
			}
			got, err := uc.RemoveCustomerById(tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveCustomerById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveCustomerById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
