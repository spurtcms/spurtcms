package controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"spurt-cms/controllers"
	"spurt-cms/graphql/info"
	"spurt-cms/graphql/model"
	storagecontroller "spurt-cms/storage-controller"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spurtcms/member"
	"github.com/spurtcms/team"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func MemberRegister(ctx context.Context, memberData *model.MemberDetails, arguments *model.MemberArguments) (bool, error) {

	c, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {

		ErrorLog.Printf("%v", info.ErrGinCtx)

		c.AbortWithStatus(500)

		return false, info.ErrGinCtx
	}

	tenantId := memberData.TenantID
	createdBy := memberData.UserID
	fmt.Println("userid:", createdBy, tenantId)

	var (
		memberSettings member.TblMemberSetting
		err            error
	)

	memberSettings, err = MemberInstance.GetMemberSettings(tenantId)
	if err != nil {

		return false, err
	}

	if memberSettings.AllowRegistration == 0 {

		return false, info.ErrMemberRegisterPerm
	}

	var memberDetails member.TblMember

	if memberData.Mobile.IsSet() {

		memberDetails.MobileNo = *memberData.Mobile.Value()
	}

	if memberData.LastName.IsSet() {

		memberDetails.LastName = *memberData.LastName.Value()
	}

	if memberData.Password != "" {

		memberDetails.Password = memberData.Password
	}

	if memberData.Username.IsSet() {

		memberDetails.Username = *memberData.Username.Value()

		isExist, err := MemberInstance.CheckNameInMember(0, *memberData.Username.Value(), tenantId)
		if isExist || err == nil {
			err = errors.New("member already exists")

			c.AbortWithError(422, err)

			return isExist, err
		}
	}

	if memberData.Email != "" {

		memberDetails.Email = memberData.Email

		isExist, err := MemberInstance.CheckEmailInMember(0, memberData.Email, tenantId)

		if isExist || err == nil {

			err = errors.New("member already exists")

			c.AbortWithError(http.StatusBadRequest, err)

			return isExist, err
		}
	}

	memberDetails.FirstName = memberData.FirstName
	memberDetails.Username = strings.ToLower(memberData.FirstName)
	memberDetails.IsActive = 1

	memberss := member.MemberCreationUpdation{
		FirstName: memberDetails.FirstName,
		LastName:  memberDetails.LastName,
		Username:  memberDetails.Username,
		Email:     memberDetails.Email,
		Password:  memberDetails.Password,
		IsActive:  memberDetails.IsActive,
		CreatedBy: createdBy,
		TenantId:  tenantId,
	}

	members, _ := MemberInstance.CreateMember(memberss)

	MemberProfile := member.MemberprofilecreationUpdation{
		MemberId:  members.Id,
		CreatedBy: createdBy,
		TenantId:  tenantId,
	}

	merr := MemberInstance.CreateMemberProfile(MemberProfile)
	if merr != nil {
		ErrorLog.Printf("memberprofile create error: %s", merr)
	}

	return true, nil
}

func MemberLogin(ctx context.Context, memberData *model.MemberSignin) (*model.MemberCheckLoginResponse, error) {

	c, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {

		ErrorLog.Printf("%v", info.ErrGinCtx)

		c.AbortWithStatus(500)

		return nil, info.ErrGinCtx
	}

	Email := memberData.Email
	Password := memberData.Password
	UserId := memberData.UserID
	TenantId := memberData.TenantID

	token, pass, email, _ := model.Model.CheckMemberLogin(Email, Password, UserId, TenantId)
	if !email {
		return &model.MemberCheckLoginResponse{
			Email:   false,
			Message: "Login failed: invalid credentials",
			Success: false,
		}, nil
	}
	if !pass {
		return &model.MemberCheckLoginResponse{
			Email:    true,
			Password: false,
			Message:  "Login failed: invalid credentials",
			Success:  false,
		}, nil
	}
	member, _ := MemberInstance.GetMemberAndProfileData(0, Email, 0, "", TenantId)

	var firstn = strings.ToUpper(member.FirstName[:1])

	var lastn string
	if member.LastName != "" {
		lastn = strings.ToUpper(member.LastName[:1])
	}
	member.NameString = firstn + lastn

	member.NameString = firstn + lastn
	return &model.MemberCheckLoginResponse{Email: true, Password: true, Message: "Login successful", Token: token, Success: true,
		MemberDetails: &model.Members{Email: member.Email, FirstName: member.FirstName, LastName: &member.LastName, ProfileImage: &member.ProfileImage, ProfileImagePath: &member.ProfileImagePath, Mobile: &member.MobileNo, NameString: &member.NameString, ID: &member.Id}}, nil
}
func MembersList(ctx context.Context, filter *model.Filter) (*model.MembersDetails, error) {

	c, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {

		ErrorLog.Printf("%v", info.ErrGinCtx)

		return &model.MembersDetails{}, info.ErrGinCtx
	}

	tenantDetails, err := GetTenantDetails(c)

	if err != nil {

		ErrorLog.Printf("%v", info.ErrFetchTenantDetails)

		c.AbortWithStatus(500)

		return &model.MembersDetails{}, info.ErrFetchTenantDetails
	}

	var (
		limit, offset int
		keyword       string
	)

	if filter != nil {

		if filter.Limit.IsSet() && filter.Limit.Value() != nil {

			limit = *filter.Limit.Value()
		}

		if filter.Offset.IsSet() && filter.Offset.Value() != nil {

			offset = *filter.Offset.Value()
		}

		if filter.Keyword.IsSet() && filter.Keyword.Value() != nil {

			keyword = *filter.Keyword.Value()
		}

	}

	input := model.MembersListReq{
		Limit:    limit,
		Offset:   offset,
		Keyword:  keyword,
		TenantId: tenantDetails.TenantId,
		Count:    true,
	}

	members, count, err := model.Model.MembersList(input)

	if err != nil {

		if err == gorm.ErrRecordNotFound {

			c.AbortWithStatus(400)

			return &model.MembersDetails{}, err
		}

		ErrorLog.Printf("%v", err)

		c.AbortWithStatus(500)

		return &model.MembersDetails{}, err
	}

	var memberlist []model.Members

	for _, member := range members {

		modifiedOn := member.ModifiedOn

		modifiedBy := member.ModifiedBy

		member := model.Members{
			ID:         member.ID,
			FirstName:  member.FirstName,
			LastName:   member.LastName,
			Email:      member.Email,
			Mobile:     member.Mobile,
			IsActive:   member.IsActive,
			CreatedOn:  member.CreatedOn,
			CreatedBy:  member.CreatedBy,
			IsDeleted:  member.IsDeleted,
			ModifiedOn: modifiedOn,
			ModifiedBy: modifiedBy,
			TenantID:   member.TenantID,
		}

		memberlist = append(memberlist, member)
	}

	return &model.MembersDetails{MembersList: memberlist, Count: int(count)}, nil
}

//Forgot Password API//

func ForgotPassword(ctx context.Context, memberdetails *model.MemberInfo) (*model.ForgotPasswordResponse, error) {

	_, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {

		ErrorLog.Printf("%v", info.ErrGinCtx)

		return &model.ForgotPasswordResponse{}, info.ErrGinCtx
	}

	member, err := MemberInstance.GetMemberAndProfileData(0, memberdetails.Email, 0, "", memberdetails.TenantID)

	id := member.Id
	token, _ := CreateTokenWithExpireTime(id, memberdetails.TenantID)

	fmt.Println(token, "token")

	if err != nil {

		ErrorLog.Printf("Send Link forget password error: %s", err)
		return &model.ForgotPasswordResponse{Message: "you are not registered with us", Success: false}, fmt.Errorf("you are not registered with us")
	}

	var url_prefix = os.Getenv("BASE_URL")

	var resetpassurl = memberdetails.URL

	linkedin := os.Getenv("LINKEDIN")
	facebook := os.Getenv("FACEBOOK")
	twitter := os.Getenv("TWITTER")
	youtube := os.Getenv("YOUTUBE")
	insta := os.Getenv("INSTAGRAM")
	data := map[string]interface{}{
		"fname":         member.Username,
		"resetpassword": resetpassurl + "/auth/change-password?token=" + token,
		"restpassurl":   resetpassurl + "/auth/change-password",
		"admin_logo":    url_prefix + "public/img/SpurtCMSlogo.png",
		"fb_logo":       url_prefix + "public/img/email-icons/facebook.png",
		"linkedin_logo": url_prefix + "public/img/email-icons/linkedin.png",
		"twitter_logo":  url_prefix + "public/img/email-icons/x.png",
		"youtube_logo":  url_prefix + "public/img/email-icons/youtube.png",
		"insta_log":     url_prefix + "public/img/email-icons/instagram.png",
		"facebook":      facebook,
		"instagram":     insta,
		"youtube":       youtube,
		"linkedin":      linkedin,
		"twitter":       twitter,
	}
	var wg sync.WaitGroup

	wg.Add(1)

	Chan := make(chan string, 1)

	go controllers.ForgetPasswordEmail(Chan, &wg, data, member.Email, "Forgot Password", "")

	close(Chan)

	return &model.ForgotPasswordResponse{Message: "Mail sent successfully!", Success: true}, nil
}

//Reset Password API//

func ResetPassword(ctx context.Context, memberinfo *model.MemberResetpassInfo) (bool, error) {

	_, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {

		ErrorLog.Printf("%v", info.ErrGinCtx)

		return false, info.ErrGinCtx
	}
	newpass := memberinfo.NewPassword

	confrimpass := memberinfo.ConfrimPassword

	token := memberinfo.Token

	Claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ErrorLog.Printf("Invalid token change password error: %s", err)
			return false, err
		}
	}

	memberIDStr, ok := Claims["member_id"].(string)
	if !ok {
		fmt.Println("member_id is not a string")
		return false, err
	}

	memberid, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("Error converting member_id:", err)
		return false, err
	}

	TenantidStr, ok := Claims["tenant_id"].(string)
	if !ok {
		return false, err
	}

	Tenantid := TenantidStr

	getuserinfo, _ := NewTeamWP.GetTenantDetails(Tenantid)

	if newpass == confrimpass {

		err := MemberInstance.MemberPasswordUpdate(newpass, confrimpass, "", memberid, getuserinfo.Id, Tenantid)
		if err != nil {
			ErrorLog.Printf("set new password error: %s", err)
			return false, err
		}

	} else {
		return false, fmt.Errorf("password don't mismatch")
	}

	return true, nil
}

func CreateTokenWithExpireTime(userid int, tenantid string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["member_id"] = (userid)
	atClaims["tenant_id"] = tenantid

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

// social Login //
func MemberSocialLogin(ctx context.Context, input *model.SocialLoginInput) (*model.SocialLoginResponse, error) {

	c, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {
		ErrorLog.Printf("%v", info.ErrGinCtx)
		c.AbortWithError(500, nil)
	}

	apiKey := c.GetHeader("ApiKey")

	tenantdetails, _ := model.Model.GettenantByapikey(apiKey)

	fmt.Println(apiKey, tenantdetails, "tenantdetailsss")

	userinfo, _ := model.Model.GetUserbyTenantId(tenantdetails.TenantId)

	memberinfo, _ := MemberInstance.GetMemberAndProfileData(0, input.Email, 0, "", tenantdetails.TenantId)

	if memberinfo.Id != 0 && memberinfo.IsActive == 0 {
		message := "This account is inactive. Please contact Admin."
		return &model.SocialLoginResponse{Message: message}, fmt.Errorf("account is inactive")
	}

	var response model.SocialLoginResponse
	var memid int
	var username string
	var lastname string
	var password string

	if memberinfo.Email != input.Email {

		if input.Username.IsSet() && input.Username.Value() != nil {

			username = *input.Username.Value()
		}
		if input.LastName.IsSet() && input.LastName.Value() != nil {

			lastname = *input.LastName.Value()
		}
		if input.Password.IsSet() && input.Password.Value() != nil {

			password = *input.Password.Value()
		}

		membergrp, err := model.Model.GetmemberGroup("default-group", tenantdetails.TenantId)

		fmt.Println(membergrp, "membergrbb")

		memberss := member.MemberCreationUpdation{
			FirstName:   input.FirstName,
			LastName:    lastname,
			Username:    username,
			Email:       input.Email,
			Password:    password,
			IsActive:    1,
			CreatedBy:   userinfo.Id,
			TenantId:    tenantdetails.TenantId,
			StorageType: "aws",
		}

		members, err := MemberInstance.CreateMember(memberss)

		if err != nil {
			fmt.Println("member creation error")
			return &model.SocialLoginResponse{Message: "member Creation error"}, err

		}

		MemberProfile := member.MemberprofilecreationUpdation{
			MemberId:  members.Id,
			CreatedBy: userinfo.Id,
			TenantId:  tenantdetails.TenantId,
		}

		merr := MemberInstance.CreateMemberProfile(MemberProfile)
		if merr != nil {
			ErrorLog.Printf("memberprofile create error: %s", merr)
			return &model.SocialLoginResponse{Message: "memberprofile create error"}, merr
		}

		memid = members.Id

	} else {
		memid = memberinfo.Id // Existing user's ID
	}

	successMessage := "Successfully Logged In"
	token, _ := CreateTokenWithExpireTime(memid, tenantdetails.TenantId)
	response.Token = token
	response.Message = successMessage
	response.Userid = memid

	return &response, nil
}

//Member Profile Details//

func MemberProfileDetails(ctx context.Context, id *int) (*model.Members, error) {

	c, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {
		ErrorLog.Printf("%v", info.ErrGinCtx)
		c.AbortWithError(500, nil)
	}
	var memberid int
	var memdet model.Members

	if id != nil {

		memberid = *id
	}
	if memberid == 0 {

		err := errors.New("unauthorized access")

		c.AbortWithError(http.StatusUnauthorized, err)

		return &model.Members{}, err

	}
	tenantData, ok := c.Get("tenantDetails")

	tenantdetails := tenantData.(team.TblUser)

	memberdetails, err := MemberInstance.GetMemberAndProfileData(memberid, "", 0, "", tenantdetails.TenantId)
	if err != nil {

		return &model.Members{}, err
	}
	var firstn = strings.ToUpper(memberdetails.FirstName[:1])
	var lastn string
	if memberdetails.LastName != "" {
		lastn = strings.ToUpper(memberdetails.LastName[:1])
	}
	memberdetails.NameString = firstn + lastn

	memberdetails.NameString = firstn + lastn
	memdet = model.Members{
		ID:               &memberdetails.Id,
		FirstName:        memberdetails.FirstName,
		LastName:         &memberdetails.LastName,
		Username:         &memberdetails.Username,
		ProfileImage:     &memberdetails.ProfileImage,
		ProfileImagePath: &memberdetails.ProfileImagePath,
		Email:            memberdetails.Email,
		IsActive:         &memberdetails.IsActive,
		Mobile:           &memberdetails.MobileNo,
		GroupID:          &memberdetails.MemberGroupId,
		NameString:       &memberdetails.NameString,
	}

	return &memdet, nil
}

//Update Memberprofile//

func UpdateMemberProfile(ctx context.Context, input *model.UpdateMember) (*model.UpdatememberResponse, error) {
	c, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {
		ErrorLog.Printf("%v", info.ErrGinCtx)
		c.AbortWithError(500, nil)
	}
	if input.ID == 0 {
		return &model.UpdatememberResponse{Message: "Invalid member ID"}, fmt.Errorf("member ID cannot be zero")
	}
	var imageName, imagePath, lastname, Password, hasspass string

	tenantData, _ := c.Get("tenantDetails")
	tenantDetails := tenantData.(team.TblUser)
	member, err := MemberInstance.GetMemberDetails(input.ID, tenantDetails.TenantId)

	if err != nil {

		return &model.UpdatememberResponse{Message: "Member not found or deleted"}, fmt.Errorf("member with ID %d not found: %v", input.ID, err)
	}

	if input.RemoveImage.IsSet() && *input.RemoveImage.Value() {
		imageName = ""
		imagePath = ""
	} else if input.ProfileImage.IsSet() && *input.ProfileImage.Value() != "" {

		var err error
		var imageByte []byte
		imageName, imagePath, imageByte, err = controllers.ConvertBase64toByte(*input.ProfileImage.Value(), "member")

		if err != nil {
			return &model.UpdatememberResponse{Message: "Failed to image upload"}, err
		}

		imagePath = tenantDetails.S3FolderName + imagePath

		if uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte); uerr != nil {
			return &model.UpdatememberResponse{Message: "Failed to image upload"}, uerr
		}
	} else {

		imageName = member.ProfileImage
		imagePath = member.ProfileImagePath
	}

	fmt.Println("imagename", imageName, imagePath)

	if input.LastName.IsSet() && input.LastName.Value() != nil {
		lastname = *input.LastName.Value()
	}

	if input.Password.IsSet() && input.Password.Value() != nil {

		Password = *input.Password.Value()
	}

	// if input.Username.IsSet() && input.Username.Value() != nil {
	// 	Username = *input.Username.Value()
	// }

	if input.Email != "" {

		isExist, err := MemberInstance.CheckEmailInMember(input.ID, input.Email, tenantDetails.TenantId)

		if isExist || err == nil {

			fmt.Println("isexit", isExist, err)

			err = errors.New("Email already exists")

			c.AbortWithError(http.StatusBadRequest, err)

			return &model.UpdatememberResponse{Message: "Email Already Exists"}, err
		}
	}

	if input.Mobile != "" {

		isExist, err := MemberInstance.CheckNumberInMember(input.ID, input.Mobile, tenantDetails.TenantId)

		if isExist || err == nil {

			err = errors.New("Mobile Number already exists")

			c.AbortWithError(http.StatusBadRequest, err)

			return &model.UpdatememberResponse{Message: "Mobile Number Already Exists"}, err
		}
	}
	if Password != "" {

		hasspass = hashingPassword(Password)

	} else {

		hasspass = member.Password
	}

	updateMap := map[string]interface{}{
		"first_name":         input.FirstName,
		"last_name":          lastname,
		"mobile_no":          input.Mobile,
		"email":              input.Email,
		"password":           hasspass,
		"profile_image":      imageName,
		"profile_image_path": imagePath,
		"is_active":          1,
		"storage_type":       "aws",
	}

	err2 := MemberInstance.MemberFlexibleUpdate(updateMap, input.ID, tenantDetails.Id, tenantDetails.TenantId)

	if err2 != nil {
		return &model.UpdatememberResponse{Message: "User Update error"}, err2
	}

	nmember, _ := MemberInstance.GetMemberDetails(input.ID, tenantDetails.TenantId)

	return &model.UpdatememberResponse{
		Message: "User Updated Successfully",
		MemberDetails: &model.Members{
			FirstName:        nmember.FirstName,
			LastName:         &nmember.LastName,
			Email:            nmember.Email,
			Mobile:           &nmember.MobileNo,
			ProfileImage:     &nmember.ProfileImage,
			ProfileImagePath: &nmember.ProfileImagePath,
			Username:         &nmember.Username,
			ID:               &nmember.Id,
		},
	}, nil
}
func hashingPassword(pass string) string {

	passbyte, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	if err != nil {

		panic(err)

	}

	return string(passbyte)
}
