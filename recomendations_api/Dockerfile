FROM python:3.9

# Set the working directory in the container
WORKDIR /app
COPY requirements.txt .

# Install packages
RUN pip install --no-cache-dir -r requirements.txt

# Copy the local code to the container
COPY app.py .
COPY movie_recommender.py .
COPY database.py .
COPY settings.py .

# Expose the port the app runs on
EXPOSE 8000

# Command to run the application
CMD ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8000"]