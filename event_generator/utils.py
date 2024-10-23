import redis
import os


def get_redis_connection() -> redis.Redis:
    redis_address = os.getenv("REDIS_ADDRESS")
    
    if redis_address is None:
        redis_address = "localhost:6379"
    host, port = redis_address.split(":")

    # Create a connection to the Redis server
    redis_connection = redis.Redis(host=host, port=int(port))

    return redis_connection
