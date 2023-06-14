package customers

import (
	"back-end-server-dev/entities"
	"back-end-server-dev/repositories/mocks"
	"back-end-server-dev/utils/orm"
	"errors"
	"reflect"
	"testing"
)

type Connection struct {
	orm orm.ObjectRelationalMappingInterface
}

func TestUseCase_GetCustomerByEmail(t *testing.T) {
	type fields struct {
		customerRepo mocks.CustomerRepoInterface
	}
	type args struct {
		customer *CustomerParam
	}

	tests := []struct {
		name     string
		fields   fields
		mockRepo func(cr *mocks.CustomerRepoInterface, data *CustomerParam)
		args     args
		want     CustomerParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name:   "Error: check when empty struct passed",
			fields: fields{customerRepo: *mocks.NewCustomerRepoInterface(t)},
			mockRepo: func(cr *mocks.CustomerRepoInterface, data *CustomerParam) {
				cr.On("GetCustomerByEmail", &data.Email).Return(entities.Customer{}, errors.New("err GetCustomerByEmail"))
			},
			args:    args{customer: &CustomerParam{}},
			want:    CustomerParam{},
			wantErr: true,
		},
		{
			name:   "Error: check when empty param passed",
			fields: fields{customerRepo: *mocks.NewCustomerRepoInterface(t)},
			mockRepo: func(cr *mocks.CustomerRepoInterface, data *CustomerParam) {
				cr.On("GetCustomerByEmail", &data.Email).Return(entities.Customer{}, errors.New("err GetCustomerByEmail"))
			},
			args:    args{customer: &CustomerParam{Email: ""}},
			want:    CustomerParam{},
			wantErr: true,
		},
		{
			name:   "Error: check when wrong param passed",
			fields: fields{customerRepo: *mocks.NewCustomerRepoInterface(t)},
			mockRepo: func(cr *mocks.CustomerRepoInterface, data *CustomerParam) {
				cr.On("GetCustomerByEmail", &data.Email).Return(entities.Customer{}, errors.New("err GetCustomerByEmail"))
			},
			args:    args{customer: &CustomerParam{Id: 1}},
			want:    CustomerParam{},
			wantErr: true,
		},
		{
			name:   "Success: correct param passed",
			fields: fields{customerRepo: *mocks.NewCustomerRepoInterface(t)},
			mockRepo: func(cr *mocks.CustomerRepoInterface, data *CustomerParam) {
				cr.On("GetCustomerByEmail", &data.Email).Return(entities.Customer{
					Id:        1,
					FirstName: "akbar",
					LastName:  "maulana",
					Email:     "akbar@example.com",
					Avatar:    "katara",
				}, nil)
			},
			args: args{customer: &CustomerParam{Email: "akbar@example.com"}},
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
				cr: &tt.fields.customerRepo,
			}
			got, err := uc.GetCustomerByEmail(tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomerByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomerByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetCustomerById(t *testing.T) {
	type fields struct {
		customerRepo mocks.CustomerRepoInterface
	}
	type args struct {
		customer *CustomerParam
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo func(cr *mocks.CustomerRepoInterface, data *CustomerParam)
		args     args
		want     CustomerParam
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name:   "Error: check when empty struct passed",
			fields: fields{customerRepo: *mocks.NewCustomerRepoInterface(t)},
			mockRepo: func(cr *mocks.CustomerRepoInterface, data *CustomerParam) {
				cr.On("GetCustomerById", &data.Id).Return(
					entities.Customer{}, errors.New("err GetCustomerById"))
			},
			args:    args{customer: &CustomerParam{}},
			want:    CustomerParam{},
			wantErr: true,
		},
		{
			name:   "Error: check when empty param passed",
			fields: fields{customerRepo: *mocks.NewCustomerRepoInterface(t)},
			mockRepo: func(cr *mocks.CustomerRepoInterface, data *CustomerParam) {
				cr.On("GetCustomerById", &data.Id).Return(
					entities.Customer{}, errors.New("err GetCustomerById"))
			},
			args:    args{customer: &CustomerParam{Id: 0}},
			want:    CustomerParam{},
			wantErr: true,
		},
		{
			name:   "Error: check when wrong param passed",
			fields: fields{customerRepo: *mocks.NewCustomerRepoInterface(t)},
			mockRepo: func(cr *mocks.CustomerRepoInterface, data *CustomerParam) {
				cr.On("GetCustomerById", &data.Id).Return(
					entities.Customer{}, errors.New("err GetCustomerById"))
			},
			args:    args{customer: &CustomerParam{Email: "akbar@example.com"}},
			want:    CustomerParam{},
			wantErr: true,
		},
		{
			name:   "Success: correct param passed",
			fields: fields{customerRepo: *mocks.NewCustomerRepoInterface(t)},
			mockRepo: func(cr *mocks.CustomerRepoInterface, data *CustomerParam) {
				cr.On("GetCustomerById", &data.Id).Return(entities.Customer{
					Id:        1,
					FirstName: "akbar",
					LastName:  "maulana",
					Email:     "akbar@example.com",
					Avatar:    "katara",
				}, nil)
			},
			args: args{customer: &CustomerParam{Id: 1}},
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
				cr: &tt.fields.customerRepo,
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
