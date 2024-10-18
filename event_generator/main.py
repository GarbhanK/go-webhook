import datetime as dt
import json
import os
import uuid
import random

from fastapi import FastAPI
import redis

app = FastAPI()

redis_address = os.getenv("REDIS_ADDRESS", "")
host, port = redis_address.split(":")
port = int(port)

# Create a connection to the Redis server
redis_connection = redis.StrictRedis(host=host, port=port)

def get_payment():
    random_payload = {
        'url': os.getenv("WEBHOOK_ADDRESS", ""),
        'webhookId': uuid.uuid4().hex,
        'data': {
            'id': uuid.uuid4().hex,
            'payment': f"PY-{''.join((random.choice('abcdxyzpqr').capitalize() for i in range(5)))}",
            'event': random.choice(["accepted", "completed", "canceled"]),
            'created': dt.datetime.now().strftime("%d/%m/%Y, %H:%M:%S"),
        }
    }
    return random_payload

@app.get("/")
async def root():
    return {"message": "Hello World"}

@app.get("/payment")
async def payment():
    webhook_payload_json = get_payment()

    # publish json string to 'payments' channel in Redis
    # redis_connection.publish('payments', webhook_payload_json)

    return webhook_payload_json
