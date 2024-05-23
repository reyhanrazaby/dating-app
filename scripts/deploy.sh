echo "Building image..."
docker build . -t reyhanrazaby/dating-app

echo "Running container..."
docker run -d -p 4545:4545 --name reyhan-dating-app reyhanrazaby/dating-app