# LLM powered Chatbot

## About

A project to provides API to enable LLM powered chat functionality and also serve capability of upload and downloading image.


## Tech Stack

- Server built on GoLang using Gorilla Mux and Cohere AI package
- LLM model used: Cohere.ai | API : https://docs.cohere.com/reference/generate
- Image is stored locally in memory (to keep things simple)

#
**NOTE: Please put your Cohere API Key before running this project**


## To build and start the project in dev mode locally
_First, create a file with name (& ext) keys.txt  in keys/ directory and put your <Cohere-api-key> into it_ <br><br>

Run the following cmd from the main/ directory : 

1. ```go run main.go```


![](/screeshots/preview.png)
<br>
_More Screenshots present in screenshots dir inside this project_

## API Documentation:

### Endpoints:

GET
```http request
/getAllMessages?pNo={P_NO}&pSize={P_SIZE}
```
**Description** : Fetches all the messages with metadata like role, message id, message type in reverse chronological order <br>
**Query Parameters** : pNo, pSize denoting page number and page size for paginated response <br>
**Response** : List of messages with metadata in reverse chornological order <br>  

### Sample Request
```shell script
curl --location --request GET 'localhost:8000/getAllMessages?pNo=1&pSize=10' \
--header 'Content-Type: application/json'
```

### Sample Response
```json
{
    "status": true,
    "data": {
        "messages": [
            {
                "role": "user",
                "type": "image/png",
                "id": "f970d70b-1224-4b20-908b-afceefce96b4"
            },
            {   "role": "assistant",
                "type": "text",
                "id": "7b81050c-5854-4e2b-a9a8-58f0f715a7e9"
            },
            {
                "role": "user",
                "type": "text",
                "id": "202921cd-0eac-4e43-8223-5542bd797234"
            }
        ]
    }
}
```

GET
```http request
/image/{id}
```
**Description** : Downloads the image with given image id  <br>
**Path Parameters** : id -> gives id for the corresponding image to fetch <br>
**Response** : Image file <br>  

### Sample Request
```shell script
curl --location --request GET 'localhost:8000/image/e8b38dee-ab6c-476b-8117-0cfbe40e747b' \
--header 'Content-Type: application/json'
```

### Sample Response


POST
```http request
/image
```
**Description** : Uploads the image <br>
**Response** : Success/Failure denoting response string <br>  

### Sample Request and Response
![Screenshot 2023-08-29 at 10 39 46 PM](https://github.com/avminus/image-chatbot/assets/35273797/82a3dda2-dd7c-43f8-b755-d9bb12898d93)

POST
```http request
/message
```
**Description** : Sends message as chat to the llm powered bot to recieve a response of text <br>
**Response** : Reply from LLM bot for the given prompt <br>  

### Sample Request
```shell script
curl --location --request POST 'localhost:8000/message' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": "name top 5 countries with highest GDPs"
}'

```
### Request Body
```json
{
    "message" : "what are the three laws of motion"
}
```

### Sample Response
```json
{
    "status": true,
    "data": [
        {
            "message": " The countries with the highest GDPs are the United States, China, India, Germany, and Japan. Here are some more countries with high GDPs:\n\n- Italy\n- Canada\n- Mexico\n- South Korea\n- Brazil\n\nThese countries have high GDPs because they have a large number of industries, high population, and high consumption."
        }
    ]
}
```
