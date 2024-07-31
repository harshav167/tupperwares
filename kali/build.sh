# docker buildx create --name mybuilder --use
# docker buildx inspect --bootstrap
docker buildx build --platform linux/arm64 -t kali-linux-arm --load . &
