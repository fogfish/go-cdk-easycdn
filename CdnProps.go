package easycdn

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

type CdnProps struct {
	Site *string `field:"required" json:"site" yaml:"site"`
	TlsCertificateArn *string `field:"required" json:"tlsCertificateArn" yaml:"tlsCertificateArn"`
	Bucket awss3.Bucket `field:"optional" json:"bucket" yaml:"bucket"`
	HttpVersion awscloudfront.HttpVersion `field:"optional" json:"httpVersion" yaml:"httpVersion"`
}

