#!/bin/bash

# Docker image repository
FRONTEND_REPO="khalt00/noitu:frontend"
BACKEND_REPO="khalt00/noitu:backend"

# Build and push Frontend
echo "Building Frontend..."
docker build -t $FRONTEND_REPO ./noitu-fe
if [ $? -ne 0 ]; then
  echo "Frontend build failed!"
  exit 1
fi

echo "Pushing Frontend to $FRONTEND_REPO..."
docker push $FRONTEND_REPO
if [ $? -ne 0 ]; then
  echo "Frontend push failed!"
  exit 1
fi

# Build and push Backend
echo "Building Backend..."
docker build -t $BACKEND_REPO ./noitu-be
if [ $? -ne 0 ]; then
  echo "Backend build failed!"
  exit 1
fi

echo "Pushing Backend to $BACKEND_REPO..."
docker push $BACKEND_REPO
if [ $? -ne 0 ]; then
  echo "Backend push failed!"
  exit 1
fi

echo "Build and push completed for both Frontend and Backend."
