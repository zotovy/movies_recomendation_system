from pydantic_settings import BaseSettings
from pydantic import Field


class Settings(BaseSettings):
    database_host: str = Field(default='localhost')
    database_db: str = Field(default='postgres')
    database_user: str = Field(default='postgres')
    database_password: str = Field(default='postgres')
    database_port: int = Field(default=5432)


settings = Settings()
print("Loaded envs:\n", settings.model_dump())
