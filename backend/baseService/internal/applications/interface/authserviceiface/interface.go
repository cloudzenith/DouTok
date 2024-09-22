package authserviceiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/verificationcode"
)

//go:generate mockgen -source=interface.go -destination=interface_mock.go -package=authserviceiface AuthService
type AuthService interface {
	CreateVerificationCode(ctx context.Context, bits, expireTime int64) (*verificationcode.VerificationCode, error)
	ValidateVerificationCode(ctx context.Context, code *verificationcode.VerificationCode) (bool, error)
}
