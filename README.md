LLM powered Chatbot with Image Upload Download Capability
About
A project to provides API to enable LLM powered chat functionality and also serve capability of upload and downloading image.

Tech Stack
Server built on GoLang using Gorilla Mux and Cohere AI package
LLM model used: Cohere.ai | API : https://docs.cohere.com/reference/generate
Image is stored locally in memory


To build and start the project in dev mode locally
First, create a file with name (& ext) keys.txt  in keys/ directory your cohere api key 

Run the following cmd from the main/ directory :

go run main.go


API Documentation:

GET                 /getAllMessages?pNo={p_no}&pSize={p_size}
Description:        Fetches all the messages with metadata like role, message id, message type in reverse chronological order
Query Parameters:   pNo, pSize denoting page number and page size for paginated response
Response:           List of messages with metadata in reverse chornological order

Sample Request:     /getAllMessages?pNo=1&pSize=10
Sample Response:
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

GET                 /image/{id}
Description:        Downloads the image with given image id 
Path Parameters:    id -> gives id for the corresponding image to fetch
Response:           image file

Sample Request:     /image/e8b38dee-ab6c-476b-8117-0cfbe40e747b
Sample Response:    


POST                /image
Description         Uploads the image 
Response:           Success/Failure message

Sample Request:     /image
Sample Response:    "Image Uploaded Successfully!"


POST                /message
Description:        Sends message as chat to the llm powered bot to recieve a response of text
Response:           Reply from LLM bot

Sample Request:     /message
Request Body:
{
    "message" : "what are the three laws of motion"
}

Sample Response:

{
    "status": true,
    "data": [
        {
            "message": " The three laws of motion are the fundamental principles of how objects move and interact with each other. These laws are:\n\n1. The law of acceleration: This law states that an object will accelerate in the direction of its velocity. This means that if an object is moving in a straight line, it will accelerate in that line.\n\n2. The law of action-reaction: This law states that an object will always move in response to an external force acting on it. This means that if an object is being acted on by a force, it will move in response to that force.\n\n3. The law of conservation of momentum: This law states that the total momentum of a system is always conserved. This means that the total momentum of an object and its surrounding objects will always remain the same.\n\nThese laws are fundamental to understanding how objects move and interact with each other. They are used to explain a wide range of phenomena, from the behavior of planets and stars to the operation of machines and devices."
        }
    ]
}