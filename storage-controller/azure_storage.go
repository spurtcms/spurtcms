package storagecontroller

import (
	"log"
	"mime/multipart"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

var (
	AZURE_STORAGE_ACCOUNT = os.Getenv("AZURE_STORAGE_ACCOUNT")
	CONTAINER_NAME        = os.Getenv("CONTAINER_NAME")
)

func CreateAzureClient() *azblob.Client {

	url := "https://" + AZURE_STORAGE_ACCOUNT + ".blob.core.windows.net/"

	credential, err := azidentity.NewDefaultAzureCredential(nil)

	if err != nil {
		log.Println(err)
	}

	client, err := azblob.NewClient(url, credential, nil)

	if err != nil {
		log.Println(err)
	}

	return client
}

func UploadFileAzure(file multipart.File, fileHeader *multipart.FileHeader, filePath string) {

	// client := CreateAzureClient()

	// client.UploadBuffer(context.Background(), CONTAINER_NAME, fileHeader.Filename)

}
