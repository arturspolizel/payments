package controller

import (
	"reflect"
	"testing"

	"github.com/arturspolizel/payments/mocks"
	"github.com/arturspolizel/payments/model"
	"github.com/xorcare/pointer"
)

type depFields struct {
	merchantRepo mocks.MerchantRepository
}

func TestMerchantController_Create(t *testing.T) {
	type args struct {
		merchant model.Merchant
	}

	tests := []struct {
		name string
		args args
		want uint

		on     func(*depFields)
		assert func(*depFields)
	}{
		{
			name: "Success",
			args: args{
				merchant: model.Merchant{
					Name:                    "Test",
					Active:                  true,
					MaximumTransactionValue: pointer.Uint(100),
				},
			},
			want: 1,
			on: func(df *depFields) {
				df.merchantRepo.Mock.On("Create", model.Merchant{
					Name:                    "Test",
					Active:                  true,
					MaximumTransactionValue: pointer.Uint(100),
				}).Return(uint(1))
			},
			assert: func(df *depFields) {
				df.merchantRepo.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			depFields := depFields{
				merchantRepo: *mocks.NewMerchantRepository(t),
			}

			tt.on(&depFields)
			defer tt.assert(&depFields)

			controller := NewMerchantController(&depFields.merchantRepo)

			if got := controller.Create(tt.args.merchant); got != tt.want {
				t.Errorf("MerchantController.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerchantController_Get(t *testing.T) {
	type args struct {
		id uint
	}

	tests := []struct {
		name string
		args args
		want model.Merchant

		on     func(*depFields)
		assert func(*depFields)
	}{
		{
			name: "Success",
			args: args{
				id: 1,
			},
			want: model.Merchant{
				ID: 1,
			},
			on: func(df *depFields) {
				df.merchantRepo.Mock.On("Get", uint(1)).Return(model.Merchant{
					ID: 1,
				})
			},
			assert: func(df *depFields) {
				df.merchantRepo.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			depFields := depFields{
				merchantRepo: *mocks.NewMerchantRepository(t),
			}

			tt.on(&depFields)
			defer tt.assert(&depFields)

			controller := NewMerchantController(&depFields.merchantRepo)
			if got := controller.Get(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MerchantController.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
