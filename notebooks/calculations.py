import pandas as pd
from sklearn.neighbors import NearestNeighbors
from scipy.sparse import csr_matrix
import pickle
import json

df = pd.read_csv('../datasets/movies-dataset/ratings.csv')

# Создание уникальных списков пользователей и фильмов
user_list = df['userId'].unique()
movie_list = df['movieId'].unique()

# Создание словарей для преобразования id в индексы
user_to_index = {user: index for index, user in enumerate(user_list)}
movie_to_index = {movie: index for index, movie in enumerate(movie_list)}

# Преобразование userId и movieId в индексы
df['user_index'] = df['userId'].map(user_to_index)
df['movie_index'] = df['movieId'].map(movie_to_index)

# Создание разреженной матрицы
num_users = len(user_list)
num_movies = len(movie_list)
matrix = csr_matrix((df['rating'], (df['user_index'], df['movie_index'])), shape=(num_users, num_movies))
print((num_users, num_movies))

knn = NearestNeighbors(metric='cosine', algorithm='brute', n_neighbors=21)
knn.fit(matrix)

# Save
knnPickle = open('knn.pickle', 'wb')
pickle.dump(knn, knnPickle)
knnPickle.close()

with open('movie_to_index.json', 'w') as f:
    movie_to_index_json = {str(movie): index for index, movie in enumerate(movie_list)}
    json.dump(movie_to_index_json, f)
