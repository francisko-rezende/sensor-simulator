# Sensor Simulator

![]('./../repo-img.jpeg)

This repository contains the Go code I used to simulate sensors generating data for my other project, Connect Lab Ultra Plus.

It reads the db and creates a random measured value for each active sensor.

The repo also has a working Dockerfile to make the app more portable.

## Prerequisites

- Docker: [Install Docker](https://docs.docker.com/get-docker/)

## Getting Started

1. Make sure you have ran the migrations and seeded the db in your Connect Lab Ultra Plus copy, this app only makes sense as a complement to the main app.

2. **Clone the Repository:**

   Clone this repository to your local machine:

   ```bash
   git clone https://github.com/francisko-rezende/sensor-simulator.git
   cd sensor-simulator
   ```

3. Create a `.env` following the model in `.env.example`

4. Build your image using `docker build -t sensor-simulator .`

5. Run your image using `docker run sensor-simulator`

And you're done!

PS: Remember to also make any additional adjustments that might be specific to your setup. For instance, I was running my db on a different docker-compose so I had to either set a `--network` param in my `docker run` call or specify it in this app's `docker-compose.yml`, as you can see by checking out the file here in the repo.
