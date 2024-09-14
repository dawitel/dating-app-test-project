import json
import random
import requests

# Generate random user data
def generate_random_user():
    names = ["John Doe", "Jane Smith", "Alice Johnson", "Bob Brown", "Emily Davis", "Michael Wilson", "Sarah Miller", "David Martinez"]
    interests = ["hiking", "cooking", "music", "coding", "traveling", "reading", "sports", "gaming"]
    genders = ["male", "female"]
    
    user = {
        "name": random.choice(names),
        "password": "the_most_secure_password_on_earth",
        "age": random.randint(18, 40),
        "gender": random.choice(genders),
        "location": {
            "latitude": round(random.uniform(-90, 90), 4),
            "longitude": round(random.uniform(-180, 180), 4)
        },
        "interests": random.sample(interests, k=random.randint(1, len(interests))),
        "preferences": {
            "min_age": random.randint(18, 30),
            "max_age": random.randint(31, 40),
            "preferred_gender": random.choice(genders),
            "max_distance": random.randint(1, 100)
        }
    }
    return user

# Endpoint URL
url = "http://localhost:8080/api/v1/sign-up"

# Send 100 user data
for _ in range(100):
    user_data = generate_random_user()
    response = requests.post(url, json=user_data)
    print(f"Status Code: {response.status_code}, Response: {response.json()}")
