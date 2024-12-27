package controller

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path"
	"spurt-cms/graphql/info"
	logPkg "spurt-cms/graphql/logger"
	"spurt-cms/graphql/model"
	"spurt-cms/models"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nfnt/resize"
	newauth "github.com/spurtcms/auth"
	"github.com/spurtcms/categories"
	chn "github.com/spurtcms/channels"
	"github.com/spurtcms/member"
	"github.com/spurtcms/team"
	role "github.com/spurtcms/team-roles"
)

type Key string

type Header string

var (
	GinContext           Key
	ErrorLog             *log.Logger
	NewAuth              *newauth.Auth
	NewRole              *role.PermissionConfig
	NewRoleWP            *role.PermissionConfig
	ChannelConfig        *chn.Channel
	ChannelConfigWP      *chn.Channel
	NewTeam              *team.Teams
	NewTeamWP            *team.Teams
	MemberInstance       *member.Member
	MemberAuthInstance   *member.Member
	CategoryInstance     *categories.Categories
	CategoryAuthInstance *categories.Categories
	MaxChunkLength       int = 8388573 //max size of a string response
)

func init() {

	if err := godotenv.Load(); err != nil {

		ErrorLog.Printf("%v", info.ErrLoadEnv)

		log.Fatal(err)
	}

	ErrorLog = logPkg.ErrorLog()

	AuthConfig()

	GetTeamInstance()

	GetTeamInstanceWithoutPermission()

	GetChannelInstance()

	GetChannelInstanceWithoutPermission()

	GetMemberInstance()

	GetMemberAuthInstance()

	GetCategoryInstance()

	GetCategoryAuthInstance()
}

// AuthCofing
func AuthConfig() *newauth.Auth {

	NewAuth = newauth.AuthSetup(newauth.Config{
		SecretKey: os.Getenv("JWT_SECRET"),
		DB:        models.DB,
		// DataBaseType: os.Getenv("DATABASE_TYPE"),
	})

	return NewAuth
}

// channel config
func GetChannelInstance() *chn.Channel {

	ChannelConfig = chn.ChannelSetup(chn.Config{
		DB:               models.DB,
		AuthEnable:       false,
		PermissionEnable: true,
		Auth:             NewAuth,
		// DataBaseType:     os.Getenv("DATABASE_TYPE"),
	})

	return ChannelConfig
}

// channel config without permission
func GetChannelInstanceWithoutPermission() *chn.Channel {

	ChannelConfigWP = chn.ChannelSetup(chn.Config{
		DB:               models.DB,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             NewAuth,
		// DataBaseType:     os.Getenv("DATABASE_TYPE"),
	})

	return ChannelConfigWP
}

// TeamConfig
func GetTeamInstance() *team.Teams {

	NewTeam = team.TeamSetup(team.Config{
		DB:               models.DB,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             NewAuth,
		// DataBaseType:     os.Getenv("DATABASE_TYPE"),
	})

	return NewTeam
}

// TeamConfig
func GetTeamInstanceWithoutPermission() *team.Teams {

	NewTeamWP = team.TeamSetup(team.Config{
		DB:               models.DB,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             NewAuth,
		// DataBaseType:     os.Getenv("DATABASE_TYPE"),
	})

	return NewTeamWP
}

// Member Instance without Auth
func GetMemberInstance() *member.Member {

	MemberInstance = member.MemberSetup(member.Config{
		DB:               models.DB,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             NewAuth,
	})

	return MemberInstance
}

func GetMemberAuthInstance() *member.Member {

	MemberAuthInstance = member.MemberSetup(member.Config{
		DB:               models.DB,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             NewAuth,
	})

	return MemberAuthInstance
}

func GetTenantDetails(c *gin.Context) (team.TblUser, error) {

	tenantData, ok := c.Get("tenantDetails")

	if !ok {

		return team.TblUser{}, info.ErrFetchTenantDetails
	}

	return tenantData.(team.TblUser), nil
}

func GetCategoryInstance() *categories.Categories {

	CategoryInstance = categories.CategoriesSetup(categories.Config{
		DB:               models.DB,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             NewAuth,
	})

	return CategoryInstance
}

func GetCategoryAuthInstance() *categories.Categories {
	CategoryAuthInstance = categories.CategoriesSetup(categories.Config{
		DB:               models.DB,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             NewAuth,
	})

	return CategoryAuthInstance
}

func ImageResize(c *gin.Context) {

	fileName := c.Query("name")

	filePath := c.Query("path")

	extension := path.Ext(fileName)

	var storageType model.TblStorageTypes

	err := model.GetStorageType(models.DB, &storageType)

	if err != nil {

		fmt.Println(err)

		c.AbortWithError(500, fmt.Errorf("%v-%v", info.ErrGetAwsCreds, err))

		return
	}

	var byteData []byte

	rawObject, err := GetObjectFromS3(storageType.Aws, strings.TrimSuffix(filePath, "/")+"/"+fileName)

	if err != nil {

		fmt.Println(err)

		c.AbortWithError(500, fmt.Errorf("%v-%v", info.ErrGetImage, err))

		return
	}

	buf := new(bytes.Buffer)

	buf.ReadFrom(rawObject.Body)

	byteData = buf.Bytes()

	extType := strings.Trim(extension, ".")

	if c.Query("width") == "" || c.Query("height") == "" {

		if extType == "svg" {

			extType = "svg+xml"
		}

		c.Data(200, "image/"+extType, byteData)

		return
	}

	width, _ := strconv.ParseUint(c.Query("width"), 10, 64)

	height, _ := strconv.ParseUint(c.Query("height"), 10, 64)

	Image, _, err := image.Decode(bytes.NewReader(byteData))

	if err != nil {

		fmt.Println(err)

		c.AbortWithError(500, fmt.Errorf("%v-%v", info.ErrDecodeImg, err))

		return
	}

	newImage := resize.Resize(uint(width), uint(height), Image, resize.Lanczos3)

	if extension == ".png" {

		err = png.Encode(c.Writer, newImage)

		if err != nil {

			fmt.Println(err)

			c.AbortWithError(500, fmt.Errorf("%v-%v", info.ErrImageResize, err))

			return
		}
	}

	if extension == ".jpeg" || extension == ".jpg" {

		err = jpeg.Encode(c.Writer, newImage, nil)

		if err != nil {

			fmt.Println(err)

			c.AbortWithError(500, fmt.Errorf("%v-%v", info.ErrImageResize, err))

			return
		}

	}

}

func GetObjectFromS3(AwsCredentials map[string]interface{}, key string) (*s3.GetObjectOutput, error) {

	session := CreateAwsSession(AwsCredentials)

	s3Svc := CreateS3Session(session)

	awsBucket := AwsCredentials["BucketName"].(string)

	rawObject, err := s3Svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(awsBucket),
		Key:    aws.String(key),
	})

	if err != nil {

		fmt.Printf("Error get object %s: %s\n", key, err)

		return nil, err
	}

	return rawObject, nil
}

// create session
func CreateS3Session(awsSession *session.Session) (ses *s3.S3) {

	svc := s3.New(awsSession)

	return svc

}

func CreateAwsSession(AwsCredentials map[string]interface{}) *session.Session {

	var awsId, awsKey, awsRegion string

	if AwsCredentials != nil {

		awsId = AwsCredentials["AccessId"].(string)

		awsKey = AwsCredentials["AccessKey"].(string)

		awsRegion = AwsCredentials["Region"].(string)

	}

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region:      &awsRegion,
			Credentials: credentials.NewStaticCredentials(awsId, awsKey, ""),
		},
	))

	return sess
}

func FetchChunkData(fullLengthData string) []string {

	var chunks []string

	dataLen := utf8.RuneCountInString(fullLengthData)

	for i := 0; i < dataLen; i += MaxChunkLength {

		var chunk string

		switch {

		case i+MaxChunkLength > dataLen:

			chunk = fullLengthData[i:dataLen]

		default:

			chunk = fullLengthData[i : i+MaxChunkLength]

		}

		chunks = append(chunks, chunk)

	}

	return chunks
}
