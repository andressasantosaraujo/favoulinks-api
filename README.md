#Favoulinks
Favoulinks is an application where a user can store their favourites bookmarks for easy access in the future. Each bookmark should have a title, clickable URL and an optional category. Using Favoulinks users can create a new bookmark, read bookmarks created, update exists bookmarks and delete bookmarks stored.

#Technologies
This is the Favoulinks server-side repository. The backend was build in AWS Lambda using Golang, Amazon DynamoDB, Amazon API Gateway and The AWS Serverless Application Model (SAM) to define the application and model it using YAML.

#Installation
   * [Create an AWS account](https://aws.amazon.com/)
   * [Configure AWS credentials](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html)
   * [Install the AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

#Usage
1. Build `template.yaml` to compile any dependencies that you have in the application and moves all the files into the `.aws-sam/build` folder :
```bash
    sam build
```
2. Package and deploy the application using the default arguments :
```bash 
    sam deploy --guided
```
3. Once deploy finish the output will show on the console :
```bash 
Key                 FavoulinksApi
Description         Favoulinks API Gateway endpoint URL for Prod stage
Value               https://x8to3u904e.execute-api.us-east-2.amazonaws.com/Prod/favoulinks/
```
   * Use the URL on `Value` to access the endpoint
# Endpoints
* GET :
    * 
     ```bash 
        curl --location --request GET 'https://uri-example.amazonaws.com/Prod/favoulinks/'
    ```
    * 
    ```bash 
        curl --location --request GET 'https://uri-example.amazonaws.com/Prod/favoulinks/?url=bbc.co.uk'
     ```
* POST :
```bash 
    curl --location --request POST 'https://uri-example.amazonaws.com/Prod/favoulinks/' \
            --header 'Content-Type: application/json' \
            --data-raw '{
                "title": "BBC",
                "url": "bbc.co.uk",
                "category": "News"
            }'
```
* PUT : 
```bash 
    curl --location --request PUT 'https://uri-example.amazonaws.com/Prod/favoulinks/' \
     --header 'Content-Type: application/json' \
     --data-raw '{
        "title": "BBC",
        "url": "bbc.co.uk",
        "category": "Everything"
     }'
```
* DELETE :
```bash 
    curl --location --request DELETE 'https://uri-example.amazonaws.com/Prod/favoulinks/?url=bbc.co.uk'
```
