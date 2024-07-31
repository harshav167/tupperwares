# docker buildx create --name mybuilder --use
# docker buildx inspect --bootstrap
# docker buildx build --platform linux/arm64 -t kali-linux-arm --load . &
docker buildx build --platform linux/amd64,linux/arm64 --push -t ghcr.io/harshav167/tupperwares/kali:latest .
