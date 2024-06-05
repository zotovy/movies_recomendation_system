from fastapi import FastAPI
from pydantic import BaseModel
from typing import List
from movie_recommender import movieRecommender
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=['*'],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


class Ratings(BaseModel):
    movie_id: int
    rating: int


def ratings_to_dict(ratings: List[Ratings]):
    d = dict()
    for rating in ratings:
        d[rating.movie_id] = rating.rating
    return d


@app.post('/recommend')
async def recommend(ratings: List[Ratings]):
    ratings_dict = ratings_to_dict(ratings)
    distances, similar_users = await movieRecommender.recommend(ratings_dict)
    return {
        'distances': distances,
        'similar_users': similar_users
    }
