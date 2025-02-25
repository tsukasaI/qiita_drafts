macOS でcolimaをセットアップするメモ

# colima & docker installation
brew install colima docker

# compose installation
DOCKER_CONFIG=${DOCKER_CONFIG:-$HOME/.docker}
mkdir -p $DOCKER_CONFIG/cli-plugins
curl -SL https://github.com/docker/compose/releases/download/v2.33.1/docker-compose-darwin-aarch64 -o $DOCKER_CONFIG/cli-plugins/docker-compose

chmod +x $DOCKER_CONFIG/cli-plugins/docker-compose


# buildx installation
DOCKER_CONFIG=${DOCKER_CONFIG:-$HOME/.docker}
curl -SL https://github.com/docker/buildx/releases/download/v0.21.1/buildx-v0.21.1.darwin-arm64 -o $DOCKER_CONFIG/cli-plugins/docker-buildx

chmod +x $DOCKER_CONFIG/cli-plugins/docker-buildx

docker buildx install
