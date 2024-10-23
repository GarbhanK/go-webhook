
curl -H 'Content-Type: application/json' \
      -d '{ "url":"foo","webhookId":"bar", "data": {"id": "a1b2x34d", "payment": "paypaypay", "event": "accepted", "created": "todayrightnowtoday"} }' \
      -X POST \
      http://localhost:8000/payment/


