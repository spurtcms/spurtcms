package storagecontroller

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"spurt-cms/models"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	AWSID       string
	AWSKEY      string
	AWSREGION   string
	AWSBUCKET   string
	Tanentid    = os.Getenv("TENANT_ID")
	TenantId, _ = strconv.Atoi(Tanentid)
)

type S3Service struct {
}

func GetSelectedType() (storagetyp models.TblStorageType, err error) {

	storagetype, err := models.GetStorageValue(TenantId)

	if err != nil {

		fmt.Println(err)

		return models.TblStorageType{}, err
	}

	return storagetype, nil

}

func SetS3value() {

	// storagetype, _ := GetSelectedType()

	// if storagetype.Aws != nil {

	// 	AWSID = storagetype.Aws["AccessId"].(string)

	// 	AWSKEY = storagetype.Aws["AccessKey"].(string)

	// 	AWSREGION = storagetype.Aws["Region"].(string)

	// 	AWSBUCKET = storagetype.Aws["BucketName"].(string)

	// }

	AWSID = os.Getenv("AWS_ACCESS_KEY_ID")
	AWSKEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
	AWSREGION = os.Getenv("AWS_DEFAULT_REGION")
	AWSBUCKET = os.Getenv("AWS_BUCKET")

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
		MaxKeys:   aws.Int64(50),
	})

	if lerr != nil {

		log.Println("Error list bucket:", lerr)

		return nil, lerr
	}

	return resp, nil
}

/*list all object from the bucket*/
func LoadMoreListS3BucketWithPath(path string, continuationToken string) (res *s3.ListObjectsV2Output, err error) {

	svc, _ := CreateS3Session()

	resp, lerr := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:            aws.String(AWSBUCKET),
		Prefix:            aws.String(path),
		Delimiter:         aws.String("/"),
		MaxKeys:           aws.Int64(50),
		ContinuationToken: aws.String(continuationToken),
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

	fmt.Println("fileName", filename)

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

	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))

	return nil
}

/*upload files from s3 */
func UploadCropImageS3(fileName string, filePath string, imagebyte []byte) error {

	sess := CreateS3Sess()

	filename := filePath

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(AWSBUCKET),
		Key:    aws.String(filename),
		Body:   strings.NewReader(string(imagebyte)),
		ACL:    aws.String("public-read"),
	})

	if err != nil {

		if aerr, ok := err.(awserr.Error); ok {

			return fmt.Errorf("failed to upload file, %v", aerr.Message())
		}
	}
	fmt.Println("file uploaded to", aws.StringValue(&result.Location), result)

	return nil
}

func CreateFolderToS3(foldername string, folderpath string) (folderPath string, err error) {

	if foldername != "" {

		svc, _ := CreateS3Session()

		// fmt.Println("inside folder create s3", folderpath+foldername)

		put, err := svc.PutObject(&s3.PutObjectInput{
			Bucket: aws.String(AWSBUCKET),
			Key:    aws.String(folderpath + foldername),
			Body:   bytes.NewReader(nil),
		})

		if err != nil {
			return "", fmt.Errorf("failed to create folder, %v", err)
		}

		fmt.Printf("create folder to, %s\n", put)

		var s3Path = folderPath + foldername + "/"

		return s3Path, nil

	}

	return "", errors.New("foldername is empty can't create")
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

	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))

	return nil

}

/*list all object from the bucket*/
func ListS3BucketWithPath1(path string) (res *s3.ListObjectsV2Output, err error) {

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

func RenameFileS3(oldFolderName, newFolderName string) error {
	svc, err := CreateS3Session()
	if err != nil {
		log.Fatalf("Failed to create S3 session: %v", err)
		return err
	}

	listInput := &s3.ListObjectsV2Input{
		Bucket: aws.String(AWSBUCKET),
		Prefix: aws.String(oldFolderName),
	}

	result, err := svc.ListObjectsV2(listInput)
	if err != nil {
		log.Fatalf("Failed to list objects: %v", err)
		return err
	}

	for _, item := range result.Contents {
		oldKey := *item.Key
		newKey := strings.Replace(oldKey, oldFolderName, newFolderName, 1)

		_, err := svc.CopyObject(&s3.CopyObjectInput{
			Bucket:     aws.String(AWSBUCKET),
			CopySource: aws.String(AWSBUCKET + "/" + oldKey),
			Key:        aws.String(newKey),
		})

		if err != nil {
			log.Printf("Failed to copy object %s to %s: %v", oldKey, newKey, err)
			return err
		}
	}

	for _, item := range result.Contents {
		_, err := svc.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(AWSBUCKET),
			Key:    item.Key,
		})

		if err != nil {
			log.Printf("Error deleting object %s: %s\n", *item.Key, err)
			return err
		}
	}

	return nil
}
func DeleteS3FolderAndContents(folderName string) error {
	svc, err := CreateS3Session()
	if err != nil {
		return fmt.Errorf("failed to create S3 session: %v", err)
	}

	listInput := &s3.ListObjectsV2Input{
		Bucket: aws.String(AWSBUCKET),
		Prefix: aws.String(folderName),
	}

	result, err := svc.ListObjectsV2(listInput)
	if err != nil {
		return fmt.Errorf("failed to list objects: %v", err)
	}

	for _, item := range result.Contents {
		_, err := svc.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(AWSBUCKET),
			Key:    item.Key,
		})
		if err != nil {
			return fmt.Errorf("error deleting object %s: %s", *item.Key, err)
		}
		fmt.Printf("Deleted object: %s\n", *item.Key)
	}

	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(AWSBUCKET),
		Key:    aws.String(folderName),
	})
	if err != nil {
		return fmt.Errorf("error deleting folder %s: %s", folderName, err)
	}

	fmt.Printf("Deleted folder: %s\n", folderName)
	return nil
}
