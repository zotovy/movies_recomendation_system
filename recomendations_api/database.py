import psycopg2
import pandas as pd
import time

from settings import settings


class Database:
    def __init__(self):
        self.__conn = self.__connect__()
        self.__cursor = self.__conn.cursor()

    @staticmethod
    def __connect__():
        start_time = time.time()
        timeout = 60  # seconds

        while True:
            try:
                connection = psycopg2.connect(
                    host=settings.database_host,
                    database=settings.database_db,
                    user=settings.database_user,
                    password=settings.database_password,
                    port=settings.database_port
                )
                print("Successfully connected to PostgreSQL")
                return connection
            except psycopg2.OperationalError as e:
                print(f"Failed to connect to PostgreSQL: {e}")
                if time.time() - start_time >= timeout:
                    print("Timeout reached, exiting.")
                    return None
                time.sleep(1)

    def get_all_ratings(self):
        return pd.read_sql("SELECT user_id, movie_id, rating FROM ratings", self.__conn)

    def get_movies_to_recommend(self, user_id: int):
        self.__cursor.execute(
            """
                SELECT DISTINCT movies.id
                FROM movies
                LEFT JOIN ratings ON movies.id = ratings.movie_id AND ratings.user_id = %s
                WHERE ratings.movie_id IS NULL;
           """,
            (user_id,),
        )

        data = self.__cursor.fetchall()
        return data


database = Database()
