package secrets_cloudstorage

// cloud-storage-admin
// https://cloud.google.com/storage/docs/reference/libraries?hl=ja

const (
	GCS_Local = "./secrets/cloudstorage/plane/ogp-local.json"
	GCS_Stg   = "./secrets/cloudstorage/plane/ogp-stg.json"
	GCS_Prd   = "./secrets/cloudstorage/plane/ogp-prd.json"
)

type Buckets struct {
	Default string
	Avatar  string
	OgImage string
}

var Buckets_Local = Buckets{
	Avatar:  "ogp-local-avatar",
	OgImage: "ogp-local-ogimage",
}

var Buckets_Stg = Buckets{
	Avatar:  "ogp-stg-avatar",
	OgImage: "ogp-stg-ogimage",
}

var Buckets_Prd = Buckets{
	Avatar:  "ogp-prd-avatar",
	OgImage: "ogp-prd-ogimage",
}
