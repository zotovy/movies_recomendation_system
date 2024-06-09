from fastapi import FastAPI
from pydantic import BaseModel
from typing import List
from movie_recommender import movieRecommender
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI(
    title="Recommendation API",
    description="API to interact with SVD user-movie model"
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=['*'],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


class MovieRecommendationItem(BaseModel):
    movie_id: int
    rating: float


@app.get('/recommend', tags=['ML'], summary='Returns list of movie recommendations for specific user')
async def recommend(user_id: int, amount: int = 10_000) -> List[MovieRecommendationItem]:
    data = movieRecommender.recommend(user_id)[:amount]
    return [
        MovieRecommendationItem(movie_id=item[0], rating=item[1]) for item in data
    ]


@app.post('/retrain', tags=['ML'], summary='Retrains model with actual data')
async def retrain():
    movieRecommender.retrain()
    return {
        "status": 200
    }


@app.get('/ping', tags=['Heath'], summary='Return pong if server is alive')
async def ping():
    return 'pong'
