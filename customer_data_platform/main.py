from typing import Union, Dict
from fastapi import FastAPI
from pydantic import BaseModel

PaymentData = Union[str, ]
class Payment(BaseModel):
    url: Union[str, None] = None
    webhookId: str
    data: Dict[str, str]


app = FastAPI()


@app.get("/")
async def root():
    return {"message": "Hello World"}

@app.post("/payment/")
async def payment(msg: Payment):
    return {"message": "message recieved!", "data": msg}
