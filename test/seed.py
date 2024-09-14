import json
import requests

# Define the API endpoint where user creation requests are sent
endpoint = "http://localhost:8080/api/v1/sign-up"

# Load user data from the JSON file
with open('t.json', 'r') as file:
    users = json.load(file)

# Loop through each user and make a POST request to create a user
for user in users:
    # Prepare the payload for the POST request
    payload = {
        "name": user['name'],
        "password": user['password'],
        "age": user['age'],
        "gender": user['gender'],
        "location": {
            "latitude": user['location']['latitude'],
            "longitude": user['location']['longitude']
        },
        "interests": user['interests'],
        "preferences": {
            "min_age": user['preferences']['min_age'],
            "max_age": user['preferences']['max_age'],
            "preferred_gender": user['preferences']['preferred_gender'],
            "max_distance": user['preferences']['max_distance']
        }
    }

    # Make the POST request
    response = requests.post(endpoint, json=payload)

    # Check the response status
    if response.status_code == 201:
        print(f"User {user['name']} created successfully.")
    else:
        print(f"Failed to create user {user['name']}. Status Code: {response.status_code}, Response: {response.text}")
