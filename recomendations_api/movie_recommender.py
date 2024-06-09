import pickle
from surprise import Reader, Dataset, SVD
from typing import List, Tuple, Optional

from database import database


class MovieRecommender:
    def __init__(self):
        self._model: Optional[SVD] = None
        self.__load_model()

    def __load_model(self):
        self._model = pickle.load(open('./static/svd_model.pkl', 'rb'))

    def retrain(self):
        # Грузим данные
        ratings = database.get_all_ratings()
        reader = Reader(line_format='user item rating', rating_scale=(1, 5))
        data = Dataset.load_from_df(ratings, reader=reader)

        full_train_set = data.build_full_trainset()

        # Обучение модели с лучшими параметрами на обучающей выборке
        model = SVD(n_epochs=35, lr_all=0.01, reg_all=0.1)
        model.fit(full_train_set)

        # Экспорт модели
        with open('./static/svd_model.pkl', 'wb') as f:
            pickle.dump(model, f)

        self._model = model

    def recommend(self, user_id: int) -> List[Tuple[int, float]]:
        """
        Выдает рекомендации по уже выставленным рекомендациям пользователя

        :param user_id: id пользователя
        :return: массив со значениями (movie_id, rating)
        """
        movies = database.get_movies_to_recommend(user_id)
        predictions = [(movie_id[0], self._model.predict(user_id, movie_id[0]).est) for movie_id in movies]
        top_movies = sorted(predictions, key=lambda x: x[1], reverse=True)
        return top_movies


movieRecommender = MovieRecommender()
