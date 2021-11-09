package handlers

import (
	"ROles-Based/gen/models"
)
// toUserDomain converts gen to domain model
//func toUserDomain(r *getting_role.GetroleParams) *domain.UserRole {
//	return &domain.UserRole{
//		Designation:     r.Designation,
//		ApiName:            r.APIName,
//	}
//}

// toUserGen converts domain to gen model
func toUserGen(success int,s string) *models.Rolereturn{
	return &models.Rolereturn{
		Code: int64(success),
		Message: s,
	}
}

func toUserGen1(InternalServer int,IS string) *models.Rolereturn1{
	return &models.Rolereturn1{
		Code: int64(InternalServer),
		Message: IS,
	}
}

func toUserGen2(forbidden int,fb string) *models.Rolereturn2{
	return &models.Rolereturn2{
		Code: int64(forbidden),
		Message: fb,
	}
}

func toUserGen3(Badrequest int,br string) *models.Rolereturn3{
	return &models.Rolereturn3{
		Code: int64(Badrequest),
		Message: br,
	}
}

