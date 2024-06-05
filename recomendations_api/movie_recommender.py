import pickle
import numpy as np
import json
from scipy.sparse import csr_matrix


class MovieRecommender:
    def __init__(self):
        self.__load_model()
        self.__load_movie_ratings()
        self.num_movies = 45115

    def __load_model(self):
        self._model = pickle.load(open('static/knn.pickle', 'rb'))

    def __load_movie_ratings(self):
        with open('static/movie_to_index.json') as json_file:
            self.movie_to_index = json.load(json_file)

    def __create_new_row(self, movie_ratings):
        """
        Создает новую строку для разреженной матрицы на основе переданных рейтингов фильмов.

        :param movie_ratings: Словарь, где ключ - movieId, значение - рейтинг
        :return: csr_matrix (разреженная матрица) с одной строкой
        """
        row = np.zeros(self.num_movies)

        for movie_id, rating in movie_ratings.items():
            if str(movie_id) in self.movie_to_index:
                index = int(self.movie_to_index[str(movie_id)])
                row[index] = rating

        return csr_matrix(row)

    async def recommend(self, ratings):
        """
        Выдает рекомендации по уже выставленным рекомендациям пользователя

        :param ratings: Словарь, где ключ - movieId, значение - рейтинг
        :return: distances, user_ids
        """
        user = self.__create_new_row(ratings)
        distances, similar_users = self._model.kneighbors(user, n_neighbors=21)
        return distances[0][1:].tolist(), similar_users[0][1:].tolist()


movieRecommender = MovieRecommender()
