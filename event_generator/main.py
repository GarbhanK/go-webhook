import json

from fastapi import FastAPI

import event_generator.events as events
import event_generator.utils as utils


app = FastAPI()

redis_connection = utils.get_redis_connection()

@app.get("/")
async def root():
    return {"message": "Hello World"}


@app.get("/payment")
async def payment():
    webhook_payload_json = json.dumps(events.get_payment())

    # publish json string to 'payments' channel in Redis
    redis_connection.publish('payments', webhook_payload_json)

    return webhook_payload_json


@app.get("/song")
async def song():
    webhook_payload_json = json.dumps(events.get_song())

    # publish json string to 'songs' channel in Redis
    redis_connection.publish('songs', webhook_payload_json)

    return webhook_payload_json


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
