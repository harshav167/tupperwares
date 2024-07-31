# Use Kali Linux rolling as the base image
FROM kalilinux/kali-rolling

# Set environment to noninteractive to avoid prompts during installation
ENV DEBIAN_FRONTEND=noninteractive

# Run apt update and install the kali-linux-large metapackage
RUN apt-get update && \
    apt-get -y install kali-linux-headless && \
    apt-get -y install neovim tmux nala && \
    apt-get -y install kali-linux-large

# Reset environment back to interactive
ENV DEBIAN_FRONTEND=interactive

# Any additional customizations or configurations can be added here

# The CMD directive specifies the command to run when the container starts
CMD ["/bin/bash"]
