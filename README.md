# Movie Recommendation System

This is a demo recommendation system that can be implemented in many services

This repository contains a movie recommendation system developed using various technologies including GoLang for the
API, FastAPI for the recommendation API, and PostgreSQL for database management. The recommendation model is trained
using datasets from TMDB (The Movie Database) and The Movies Dataset. The machine learning model is implemented using
Surprise library with the SVD algorithm. Additionally, the system includes a fast and intelligent movie autocomplete
feature.

## Features:

1. Movies recommendations using user-based collaborative filtering
2. Fast fuzzy and intelligent autocomplete

## Components:

1. **API (GoLang)**:
    - This component serves as the backend API for the movie recommendation system.
    - It handles user requests and interacts with the recommendation engine and database.

2. **Recommendation API (FastAPI)**:
    - Used for interaction with ML model

3. **PostgreSQL**
    - PostgreSQL is used to store movie data and user interactions
    - It stores information about movies, user ratings, and other relevant data for the recommendation engine.

4. **Machine Learning Model (Surprise - SVD Algorithm):**
    - The core recommendation engine is powered by a machine learning model built using the Surprise library.
    - The SVD (Singular Value Decomposition) algorithm is employed to make movie recommendations to users.

#### Datasets

1. [TMDB 5000 Movie Dataset](https://www.kaggle.com/datasets/tmdb/tmdb-movie-metadata): A comprehensive dataset
   containing information about movies, including metadata, cast, crew
1. [The Movies Dataset](https://www.kaggle.com/datasets/rounakbanik/the-movies-dataset): Provides a wide range of
   movie-related information for building a robust recommendation engine, includes movies ratings

## Usage:
Clone the repository:

```bash
git clone https://github.com/zotovy/movies_recomendation_system.git
```

Run via docker:
```bash
cd delpoyments && docker-compose -f docker-compose.dev.yaml up
```

Setup database:
```bash
bash scripts/sh/setup_database.sh
```

### TODO-List
- [ ] Web application
- [ ] REST endpoints to manage user ratings
- [ ] MLFlow to manage relevance of SVD model
- [ ] Item-based recommendations