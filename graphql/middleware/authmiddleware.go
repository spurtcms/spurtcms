package middleware

import (
	"context"
	"errors"
	"net/http"
	"spurt-cms/graphql/controller"
	"spurt-cms/graphql/info"
	"spurt-cms/graphql/model"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
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

func LimitRequestBodySize(limit int64) gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, limit)

		c.Next()

	}
}
