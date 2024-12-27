package info

import "errors"

var (
	ErrLoadEnv              = errors.New("failed to load env file")
	ErrGinCtx               = errors.New("failed to get gin context")
	ErrAuth                 = errors.New("unauthorized access")
	ErrReqApiKey            = errors.New("required api token to access graphql server")
	ErrExpireApiKey         = errors.New("graphql api token expired")
	ErrTenantId             = errors.New("tenant id is required")
	ErrMemberRegisterPerm   = errors.New("member register permission denied")
	ErrFetchTenantDetails   = errors.New("failed to get tenant Details")
	ErrGetImage             = errors.New("failed to retrieve the file")
	ErrGetAwsCreds          = errors.New("failed to retrieve the aws credentials")
	ErrDecodeImg            = errors.New("failed to parse the image data")
	ErrImageResize          = errors.New("failed to resize the image")
	ErrFetchCategoryDetails = errors.New("failed to get the category list")
	ErrReqMandatory         = errors.New("required mandatory fields")
	ErrNoRowsAffected       = errors.New("no rows affected")
	ErrUpdateViewCount      = errors.New("failed to update view count")
	ErrRecordNotFound       = errors.New("record not found")
)
