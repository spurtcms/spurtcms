package storagecontroller

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"spurt-cms/models"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	AWSID     string
	AWSKEY    string
	AWSREGION string
	AWSBUCKET string
)

type S3Service struct {
}

func GetSelectedType() (storagetyp models.TblStorageType, err error) {

	storagetype, err := models.GetStorageValue()

	if err != nil {

		fmt.Println(err)

		return models.TblStorageType{}, err
	}

	return storagetype, nil

}

func SetS3value() {

	storagetype, _ := GetSelectedType()

	if storagetype.Aws != nil {

		AWSID = storagetype.Aws["AccessId"].(string)

		AWSKEY = storagetype.Aws["AccessKey"].(string)

		AWSREGION = storagetype.Aws["Region"].(string)

		AWSBUCKET = storagetype.Aws["BucketName"].(string)

	}

}

// create session
func CreateS3Session() (ses *s3.S3, err error) {

	SetS3value()

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(AWSREGION),
		Credentials: credentials.NewStaticCredentials(AWSID, AWSKEY, ""),
	})

	if err != nil {

		log.Println("Error creating session: ", err)

		return nil, err

	}

	svc := s3.New(sess)

	return svc, nil

}

func CreateS3Sess() *session.Session {

	SetS3value()

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region:      &AWSREGION,
			Credentials: credentials.NewStaticCredentials(AWSID, AWSKEY, ""),
		},
	))

	return sess
}

/*list all object from the bucket*/
func ListS3BucketWithPath(path string) (res *s3.ListObjectsV2Output, err error) {

	svc, _ := CreateS3Session()

	resp, lerr := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:    aws.String(AWSBUCKET),
		Prefix:    aws.String(path),
		Delimiter: aws.String("/"),
	})

	if lerr != nil {

		log.Println("Error list bucket:", lerr)

		return nil, lerr
	}

	return resp, nil
}

/*list all object from the bucket*/
func ListS3Bucket() (res *s3.ListObjectsV2Output, err error) {

	svc, _ := CreateS3Session()

	resp, lerr := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(AWSBUCKET),
	})

	if lerr != nil {

		log.Println("Error list bucket:", lerr)

		return nil, lerr
	}

	return resp, nil
}

/*upload files from s3 */
func UploadFileS3(file multipart.File, fileHeader *multipart.FileHeader, filePath string) error {

	sess := CreateS3Sess()

	filename := filePath + fileHeader.Filename

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(AWSBUCKET),
		Key:    aws.String(filename),
		Body:   file,
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}

	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location), result)

	return nil
}

func CreateFolderToS3(foldername string, folderpath string) error {

	if foldername != "" {

		svc, _ := CreateS3Session()

		// fmt.Println("inside folder create s3", folderpath+foldername)

		put, err := svc.PutObject(&s3.PutObjectInput{
			Bucket: aws.String(AWSBUCKET),
			Key:    aws.String(folderpath + foldername + "/"),
			Body:   bytes.NewReader(nil),
		})

		if err != nil {
			return fmt.Errorf("failed to create folder, %v", err)
		}

		fmt.Printf("create folder to, %s\n", put)

		return nil

	}

	return errors.New("foldername is empty can't create")
}

func DeleteS3Files(filename string) error {

	svc, _ := CreateS3Session()

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(AWSBUCKET),
		Key:    aws.String(filename),
	})

	if err != nil {

		fmt.Printf("Error deleting object %s: %s\n", filename, err)

		return err
	}

	fmt.Printf("Object %s deleted successfully!\n", filename)

	return nil
}

func GetObjectFromS3(key string) (*s3.GetObjectOutput, error) {

	svc, _ := CreateS3Session()

	rawObject, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(AWSBUCKET),
		Key:    aws.String(key),
	})

	if err != nil {

		fmt.Printf("Error get object %s: %s\n", key, err)

		return nil, err
	}

	return rawObject, nil
}

/*convert */
func ConvertS3ImagetoBase64(rawObject *s3.GetObjectOutput) (string, error) {

	// Read the image data from the response body
	imageData, err := io.ReadAll(rawObject.Body)

	if err != nil {

		return "", err
	}

	base64ImageData := base64.StdEncoding.EncodeToString(imageData)

	return base64ImageData, nil

}

func StoreS3Base64(base64Image string, filepath string, key string) error {

	// Decode base64 string into binary data
	imageData, err := base64.StdEncoding.DecodeString(base64Image)

	sess := CreateS3Sess()

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(AWSBUCKET),
		Key:    aws.String(filepath + key),
		Body:   strings.NewReader(string(imageData)),
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}

	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location), result)

	return nil

}
