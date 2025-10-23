package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"spurt-cms/graphql/controller"
	"spurt-cms/graphql/info"
	"spurt-cms/graphql/model"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spurtcms/team"
	"gorm.io/gorm"
)

func AuthMiddleware(ctx context.Context, _ interface{}, next graphql.Resolver) (interface{}, error) {

	c, ok := ctx.Value(controller.GinContext).(*gin.Context)

	if !ok {

		controller.ErrorLog.Printf("%v", info.ErrGinCtx)

		c.AbortWithStatus(http.StatusInternalServerError)

		return nil, info.ErrGinCtx
	}

	var graphqlSettings model.TblGraphqlSettings

	err := AuthenticateApiKey(c, &graphqlSettings)

	if err != nil {

		controller.ErrorLog.Printf("%v", err)

		c.AbortWithStatus(http.StatusBadRequest)

		return nil, err
	}

	tenantDetails, err := controller.NewTeamWP.UserDetails(team.Team{TenantId: graphqlSettings.TenantId})

	if err != nil {

		if err == gorm.ErrRecordNotFound {

			controller.ErrorLog.Printf("%v", info.ErrTenantId)

			c.AbortWithStatus(http.StatusBadRequest)

			return nil, err
		}

		controller.ErrorLog.Printf("%v", err)

		c.AbortWithStatus(http.StatusInternalServerError)

		return nil, err

	}
	userkey := c.GetHeader("Userkey")

	var userdet *model.Members

	if userkey != "" {

		fmt.Println("checkmemberlogin")

		userdet, err = AuthenticateUserLogin(c, userkey)

		if err != nil {

			controller.ErrorLog.Printf("%v", err)

			c.AbortWithStatus(http.StatusBadRequest)

			return nil, err
		}

	}

	fmt.Println(userdet, "userdetails")

	c.Set("userdetails", userdet)

	c.Set("tenantDetails", tenantDetails)

	return next(ctx)
}

func AuthenticateApiKey(c *gin.Context, graphqlSettings *model.TblGraphqlSettings) error {

	apiKey := c.GetHeader("ApiKey")

	err := model.Model.GetApiSettings(apiKey, graphqlSettings)

	if err != nil {

		switch {

		case err == gorm.ErrRecordNotFound && apiKey != "":

			return errors.New("invalid api key")

		case err == gorm.ErrRecordNotFound && apiKey == "":

			return errors.New("api key required")

		default:

			return err

		}
	}

	if graphqlSettings.Duration != "Unlimited" && time.Now().After(graphqlSettings.ExpiryTime) {

		return info.ErrExpireApiKey
	}

	return nil

}

func AuthenticateUserLogin(c *gin.Context, userkey string) (userdetails *model.Members, err error) {

	Claims := jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(userkey, Claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	var ErrorToken = errors.New("invalid token")
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println(err, "kkkkk")
			return &model.Members{}, ErrorToken
		}

		return &model.Members{}, ErrorToken
	}

	if !tkn.Valid {
		fmt.Println(tkn)
		return &model.Members{}, ErrorToken
	}
	usrid := Claims["member_id"]

	tenantid := Claims["tenant_id"]

	fmt.Println(usrid, tenantid, "detailsauth")

	useridint := int(usrid.(float64))

	tenantint := string(tenantid.(string))

	// user, _ := strconv.Atoi(useridint)

	c.Set("memberid", useridint)

	member, _ := controller.MemberInstance.GetMemberDetails(useridint, tenantint)

	return &model.Members{FirstName: member.FirstName, LastName: &member.LastName, Email: member.Email, Mobile: &member.MobileNo, ProfileImage: &member.ProfileImage, ProfileImagePath: &member.ProfileImagePath, Username: &member.Username, ID: &member.Id}, nil
}

func LimitRequestBodySize(limit int64) gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, limit)

		c.Next()

	}
}
