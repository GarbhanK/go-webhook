import uuid
import random
import datetime as dt
import os

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

def get_song():
    random_payload = {
        'url': os.getenv("WEBHOOK_ADDRESS", ""),
        'webhookId': uuid.uuid4().hex,
        'data': {
            'id': uuid.uuid4().hex,
            'song_title': random.choice(["Southern Nights", "So Into You", "Song 2", "Gee Baby"]),
            'event': random.choice(["played", "paused", "add_favourite", "remove_favourite"]),
            'created': dt.datetime.now().strftime("%d/%m/%Y, %H:%M:%S"),
        }
    }
    return random_payload
