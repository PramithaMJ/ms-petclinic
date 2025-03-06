from fastapi import FastAPI
from uuid import uuid4
import requests
import logging

# Set up FastAPI application
app = FastAPI()

# Register with Eureka
def register_with_eureka():
    eureka_url = "http://localhost:8761/eureka/apps/owner-service"
    app_data = {
        "instance": {
            "hostName": "localhost",
            "app": "owner-service",
            "ipAddr": "127.0.0.1",
            "status": "UP",
            "port": {"$": 8086, "@enabled": "true"},
            "vipAddress": "owner-service",
            "secureVipAddress": "owner-service"
        }
    }
    headers = {'Content-Type': 'application/json'}
    response = requests.post(eureka_url, json=app_data, headers=headers)
    if response.status_code == 200:
        logging.info("Successfully registered with Eureka")
    else:
        logging.error(f"Failed to register with Eureka: {response.status_code}")

register_with_eureka()

# Define API endpoints
@app.get("/owners")
async def get_owner():
    owner_id = str(uuid4())
    return {"owner_id": owner_id, "name": "John Doe", "email": "john.doe@example.com"}

@app.post("/owners")
async def create_owner(owner_name: str, owner_email: str):
    owner_id = str(uuid4())
    return {"owner_id": owner_id, "name": owner_name, "email": owner_email}
