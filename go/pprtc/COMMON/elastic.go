package elasticupload

import (
	"TSM/common/elasticsearchuploader"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/olivere/elastic/v7"
	"github.com/sha1sum/aws_signing_client"
)

type elasticSearch struct {
	signer  *v4.Signer
	domain  string
	region  string
	service string
}

func NewElasticeSearch(sign *v4.Signer, domain string, region string, es string) *elasticSearch {
	return &elasticSearch{
		signer:  sign,
		domain:  domain,
		region:  region,
		service: es,
	}
}

// UploadJsonDoc - Uploads Json Document To Elastic - 2021-08-04
// just to understand index = db, id = document id, jsonDoc = json document object
func (svc *elasticSearch) UploadJsonDoc(index string, id string, jsonDoc string) (err error) {

	endpoint := svc.domain + "/" + index + "/" + "_doc" + "/" + id

	convertedJsonDoc := strings.NewReader(jsonDoc)

	// An HTTP client for sending the request
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, endpoint, convertedJsonDoc)
	if err != nil {
		return elasticsearchuploader.ErrCreateReq
	}

	// You can probably infer Content-Type programmatically, but here, we just say that it's JSON
	req.Header.Add("Content-Type", "application/json")

	// Sign the request, send it, and print the response
	_, err = svc.signer.Sign(req, convertedJsonDoc, svc.service, svc.region, time.Now())
	if err != nil {
		return elasticsearchuploader.ErrSigning
	}

	resp, err := client.Do(req)
	fmt.Println("resp :: ", resp)
	fmt.Println("err :: ", err)
	if err != nil {
		return elasticsearchuploader.ErrWhileUploadingTo
	}

	return nil
}

// BulkUploadJsonDoc - Uploads Whole String Document To Elastic - 2021-08-12
// just to understand index = db, id = document id, jsonDoc = json document object
func (svc *elasticSearch) BulkUploadJsonDoc(index string, elasticDoc string) (err error) {

	endpoint := svc.domain + "/_bulk"

	convertedElasticDoc := strings.NewReader(elasticDoc)

	// An HTTP client for sending the request
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, endpoint, convertedElasticDoc)
	if err != nil {
		return elasticsearchuploader.ErrCreateReq
	}

	// You can probably infer Content-Type programmatically, but here, we just say that it's JSON
	req.Header.Add("Content-Type", "application/json")

	// Sign the request, send it, and print the response
	_, err = svc.signer.Sign(req, convertedElasticDoc, svc.service, svc.region, time.Now())
	if err != nil {
		return elasticsearchuploader.ErrSigning
	}

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return elasticsearchuploader.ErrWhileBulkUploading
	}

	return nil
}

func (svc *elasticSearch) BulkUpload(ctx context.Context, bulkRequests []elastic.BulkableRequest) (err error) {
	client, err := svc.getESCLient()
	if err != nil {
		fmt.Println("Common BulkUpload 1:", err)
		return
	}
	defer client.Stop()

	bulkRequest := client.Bulk()

	for _, req := range bulkRequests {
		bulkRequest.Add(req)
	}

	_, err = bulkRequest.Do(ctx)
	if err != nil {
		fmt.Println("Common BulkUpload 2", err)
		return
	}
	fmt.Println("Error Response From ES AWS", err)
	//fmt.Println(resp)
	return
}

func (svc *elasticSearch) getESCLient() (*elastic.Client, error) {

	awsClient, err := aws_signing_client.New(svc.signer, nil, svc.service, svc.region)
	if err != nil {
		return nil, err
	}

	return elastic.NewClient(
		elastic.SetURL(svc.domain),
		elastic.SetScheme("https"),
		elastic.SetSniff(false),
		elastic.SetHttpClient(awsClient),
		elastic.SetHealthcheck(false),
	)
}

// DeleteJsonDoc - Deletes Json Document From Elastic - 2021-08-11
// just to understand index = db, id = document id, jsonDoc = json document object
func (svc *elasticSearch) DeleteJsonDoc(index string, id string) (err error) {

	endpoint := svc.domain + "/" + index + "/" + "_doc" + "/" + id

	// An HTTP client for sending the request
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return elasticsearchuploader.ErrCreateReq
	}

	// You can probably infer Content-Type programmatically, but here, we just say that it's JSON
	req.Header.Add("Content-Type", "application/json")

	// Sign the request, send it, and print the response
	_, err = svc.signer.Sign(req, nil, svc.service, svc.region, time.Now())
	if err != nil {
		return elasticsearchuploader.ErrSigning
	}

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return elasticsearchuploader.ErrWhileDeletingFrom
	}

	return nil
}
