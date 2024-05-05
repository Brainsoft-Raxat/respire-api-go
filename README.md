# Respire API Go
Our project, "**respire.**", aims to address the issue of smoking addiction by fostering a collaborative and supportive environment around users. Recognizing that quitting smoking is a challenging endeavor that benefits greatly from encouragement and accountability, our project leverages the power of community and technology to create a unique support system for its users.

This repository is a backend app for our android app. 

## Table of Contents

- [Respire API Go](#respire-api-go)
  - [Table of Contents](#table-of-contents)
    - [Link to the Mobile app's  source code:](#link-to-the-mobile-apps--source-code)
  - [Overview](#overview)
    - [Demo Video](#demo-video)
  - [Technologies](#technologies)
  - [UN Sustainable Development Goals And Targets](#un-sustainable-development-goals-and-targets)
  - [Core functionalities](#core-functionalities)
    - [Smoking Habbits Tracker](#smoking-habbits-tracker)
    - [Social Networking and Engaging in Cooperative Challenges](#social-networking-and-engaging-in-cooperative-challenges)
    - [Personalized AI Generated Daily Recommendations to Cope with Cravings](#personalized-ai-generated-daily-recommendations-to-cope-with-cravings)
  - [Installation](#installation)
    - [Prerequisites](#prerequisites)
    - [Clone the repository](#clone-the-repository)
    - [Install dependencies](#install-dependencies)
  - [Usage](#usage)
    - [Build and run the application](#build-and-run-the-application)
  - [Configuration](#configuration)
  - [Swagger documentation](#swagger-documentation)

### Link to the Mobile app's  source code:
https://github.com/aidanakalimbekova/respire-mobile

## Overview

### Demo Video
![respire._pages-to-jpg-0001.jpg](ReadmeContent/respire._pages-to-jpg-0001.jpg)
> Youtube Video Link: [https://www.youtube.com/watch?v=XZaXcsDsGV4](https://youtu.be/B7IVdjwAz60?si=q3IRj3_KTUFKOEeb)

___

## Technologies
<div align="center">
	<img height="60" src="https://user-images.githubusercontent.com/25181517/189716855-2c69ca7a-5149-4647-936d-780610911353.png" alt="Firebase" title="Firebase" />
	<img height="60" src="https://user-images.githubusercontent.com/63765620/228299869-36a40db1-c608-45cb-8fd4-fbe45425ecb2.png" alt="Android" title="Android" />
	<img height="60" src="https://user-images.githubusercontent.com/63765620/228302531-4822866b-d460-4741-9185-958f17fce9f7.png" alt="Google Cloud" title="Google Cloud" />
    <img height="60" src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png" alt="Golang" title="Golang" />
    <img height="70" src="https://thirdeyedata.ai/wp-content/uploads/2021/07/VertexAI-512-color.png" alt="Vertex AI" title="Vertex AI" style="margin-left: 5px;" />
    <img height="70" src="https://media.licdn.com/dms/image/D4E12AQHQP9J275Q_uA/article-cover_image-shrink_600_2000/0/1700940849777?e=2147483647&v=beta&t=m0HEQrukIOqU4fe1K9M19PaHq3UbvEubLzeIH1shcSc" alt="LangChain" title="LangChain" style="margin-left: 2px;"/> 
    <img height="59" src="https://miro.medium.com/v2/resize:fit:793/0*RTW5byy6eH_eSWTP.png" alt="Chroma" title="Chroma" style="margin-left: 2px;"/> 
</div>

## UN Sustainable Development Goals And Targets

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/c/c4/Sustainable_Development_Goal_3.png/1200px-Sustainable_Development_Goal_3.png" alt="SDG" style="width:200px;height:200px;">

## Core functionalities
### Smoking Habbits Tracker
<div align="center">
	<img height="400" src="ReadmeContent/1.png" alt="Tracker" title="Tracker" />
    <img height="400" src="ReadmeContent/2.png" alt="Add Pill" title="Ad Pill" />
	<img height="400" src="ReadmeContent/5.png" alt="Add Pill" title="Ad Pill" />
</div>

### Social Networking and Engaging in Cooperative Challenges
 <div align="center">
	<img height="400" src="ReadmeContent/12.png" alt="Tracker" title="Tracker" />
    <img height="400" src="ReadmeContent/6.png" alt="Add Pill" title="Ad Pill" />
	<img height="400" src="ReadmeContent/11.png" alt="Add Pill" title="Ad Pill" />
</div>

### Personalized AI Generated Daily Recommendations to Cope with Cravings
<div align="center">
	<img height="400" src="ReadmeContent/8.png" alt="Tracker" title="Tracker" />
    <img height="400" src="ReadmeContent/9.png" alt="Add Pill" title="Ad Pill" />
</div>

## Installation
### Prerequisites

- Go (version 1.21.4)

### Clone the repository

```bash
git clone https://github.com/Brainsoft-Raxat/respire-api-go.git
cd repository
```
### Install dependencies
```bash
go get ./...
```

## Usage
### Build and run the application
```bash
go build ./cmd/app/main.go -o app .
./app
```

By default, the application will run on localhost:8080.

## Configuration
- Edit config.yaml according to your configs.
- When local development set ***env*** to 'local' and ***host*** to localhost.
- Feel free to set any port you want.
- Download **Firebase SDK JSON** configuration and put it into project root directory. In our case file is named as "quitsmoke-20141-firebase-adminsdk-ugo14-c5730ea21d.json"
  
## Swagger documentation
https://respire-api-go-jc4tvs5hja-ey.a.run.app/swagger//index.html#/users/get_user_search

