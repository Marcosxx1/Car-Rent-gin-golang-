package authusecase

import (
	"testing"

	authdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/auth"
	testingmocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/testing-mocks"
	"github.com/gin-gonic/gin"
)


func TestLoginUseCase_Execute(t *testing.T) {
	mockContext := &gin.Context{} 
	mockUserRepo := &testingmocks.MockUserRepository{}
	useCase := NewLoginUseCase(mockContext, mockUserRepo)

	tests := []struct {
		name     string
		request  *authdto.LoginInputDTO
		want     string
		wantErr  bool
		errMsg   string
	}{
		{
			name: "Valid Login",
			request: &authdto.LoginInputDTO{
				Email:    "test@example.com",
				Password: "password123",
			},
			want:    "valid_token",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := useCase.Execute(tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("LoginUseCase.Execute() error = %v, wantErrMsg %v", err, tt.errMsg)
				return
			}
			if got != tt.want {
				t.Errorf("LoginUseCase.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
