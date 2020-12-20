# Golang Webhook Filter

Nanoservice like a standalone Lambda function based on Golang to receive and filter data from a webhook of a DAM system. Filter requests and trigger action only on certain criteria.

Comes with a mock-api endpoint which is supposed to simulate an external API that gets forwarded the filtered results in production and does not exist locally.


```
curl --location --request POST 'http://localhost:3000/webhook' \
--header 'Content-Type: application/json' \
--data-raw '{"AssetId": "87c23cwqDD2111", "metadata": {"folderPath": "/Client/Watanga Group Holding SE & Co. KGaA/my-image-name-jpg", "cf_approvalState_client1": "Approved", "cf_assetType": {"value": "Content Image"}}}'
```

## Build local Docker file
```
docker build -t dam-webhook-go .
docker run -p 3001:3000 -e APP_ENV=production --name dam-webhook dam-webhook-go 
```