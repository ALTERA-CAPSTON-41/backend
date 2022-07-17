<p align="center">
  <h1 align="center">Clinic Management API</h1>

  <p align="center">
    REST API of Clinic Application
  </p>
</p>

<div align="center">
  
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![semantic-release: angular][semantic-badge]][semantic-url]
[![codecov][codecov-shield]][codecov-url]
  
</div>

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#deployment-guide">Deployment Guide</a></li>
      </ul>
    </li>
    <li><a href="#features">Features</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#contributors">Contributors</a></li>
  </ol>
</details>

## About The Project
Hospital has its own business process to serve the community for healthcare. They will need data that can improve their services. Digitalization of health record will be saved a lot of time to documentate and to look up service data and the patients.

This project aims to help digitize data from health facilities. The target of this project is level 1 and 2 health facilities. It is hoped that these health facilities can improve their services without having to think about archiving patient data and medical records.

However, the features in this project are limited to the main features. Not included in the service payment and inventory management module at the pharmacy, which should be used as another topic in the capstone project in this batch.

## Getting Started

### Prerequisites

When you're going to contribute or compile, you'll need at least:
  - Go 1.18+
  - MySQL 8.0.x
  - MongoDB 5.x

### Installation or Configure

```bash
# clone if you don't have the code base
$ git clone git@github.com:ALTERA-CAPSTON-41/backend.git

# tidy up and run
$ go mod tidy

$ go run main.go
```

### Deployment Guide

#### Self-Build
```sh
# clone if you don't have the code base
$ git clone git@github.com:ALTERA-CAPSTON-41/backend.git

# install the dependencies
$ go mod download

# compile 
$ go build -o main
```

After build, move with the `public` folder next to executable  file. Then, execute with `./main` command. Finally, configure the `.env` file to be configured trough your own MySQL database.

#### Docker Build 
```sh
# clone if you don't have the code base
$ git clone git@github.com:ALTERA-CAPSTON-41/backend.git

# build a docker image
$ docker build -t clinic-backend-api .

# create a container that running at port :19000
$ docker run -dp 19000:8000 --name clinic-backend-api clinic-backend-api

# copy .env file to your container
$ docker cp ./app.env clinic-backend-api:/app/.

# restart container
$ docker restart clinic-backend-api
```
If you decide to use docker method, you need to run docker daemon and create the .env file first. Then deployment can be done.

#### GHCR.io Image
```sh
# pull docker image at latest version
$ docker pull ghcr.io/altera-capston-41/backend:latest

# create a container that running at port :19000
$ docker run -dp 19000:8000 --name clinic-backend-api ghcr.io/altera-capston-41/backend:latest

# copy .env file to your container
$ docker cp ./app.env clinic-backend-api:/app/.

# restart container
$ docker restart clinic-backend-api
```
If you decide to use docker method, you need to run docker daemon and create the .env file first. Then deployment can be done.

## Features

### 📝 Medical Record
Medical records are recorded by doctors after observing patients. In the medical record, the doctor also records the diagnosis of the disease experienced by the patient.

Recording of diagnoses is carried out with ICD-10 standardization so that it is easier to read and document. Observation results including symptoms and suggestions can be documented in this feature.

Medical records can only be made by a doctor. However, admins and nurses will only get read-only access to medical records. Because, recording the diagnosis from the results of observations can only be done by a doctor.

### 🎫 Service Queue Manager
Queue is done before the observation by the doctor. Queue registration is done by the admin with the data source from the patient's ID card (aka. KTP). This queue registration is focused on outpatient services and referral patients.

This feature does not include the type of registration of patients with insurance or patients without it. In addition, it does not include financial management or cashier for service fees.

### 💊 Prescription
Drug prescriptions are written based on the diagnosis based on the observation by the doctor. This feature only focuses on recipe writing. Not included in the drug stock management feature.

## Licenses

All of this source code are licensed under the MIT licenses. Although this source code is related to the Alterra Academy.

## Contact

- Adhicitta Masran - <adhicittamasran@gmail.com>
- Hamdan Yuwafi - <yuwafi.hamdan365@gmail.com>

## Contributors

✨ Thanks goes to these wonderful people who speed up the development process: 

<!-- ALL-CONTRIBUTORS-LIST:START -->
<table>
    <tr>
        <td align="center">
            <a href="https://github.com/dhichii">
                <img src="https://avatars.githubusercontent.com/u/75155775?v=4?s=100" width="100px;" alt=""/>
                <br />
                <sub><b>Adhicitta Masran</b></sub>
            </a>
        </td>
        <td align="center">
            <a href="https://github.com/thisham">
                <img src="https://avatars.githubusercontent.com/u/59078748?v=4?s=100" width="100px;" alt=""/>
                <br />
                <sub><b>Hamdan Yuwafi</b></sub>
            </a>
        </td>
        <td align="center">
            <a href="https://github.com/hudabikhoir">
                <img src="https://avatars.githubusercontent.com/u/35209506?v=4?s=100" width="100px;" alt=""/>
                <br />
                <sub><b>Nur Huda Bikhoir</b></sub>
            </a>
        </td>
        <td align="center">
            <a href="https://github.com/bimbimprasetyoafif">
                <img src="https://avatars.githubusercontent.com/u/26946357?v=4?s=100" width="100px;" alt=""/>
                <br />
                <sub><b>Bimo Prasetyo Afif</b></sub>
            </a>
        </td>
    </tr>
    <tr>
      <td align="center">
        Main Contributor
      </td>
      <td align="center">
        Main Contributor
      </td>
      <td align="center">
        Mentor
      </td>
      <td align="center">
        Mentor
      </td>
    </tr>
</table>
<!-- ALL-CONTRIBUTORS-LIST:FINISH -->

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[stars-shield]: https://img.shields.io/github/stars/ALTERA-CAPSTON-41/backend.svg?style=for-the-badge
[stars-url]: https://github.com/ALTERA-CAPSTON-41/backend/stargazers
[issues-shield]: https://img.shields.io/github/issues/ALTERA-CAPSTON-41/backend.svg?style=for-the-badge
[issues-url]: https://github.com/ALTERA-CAPSTON-41/backend/issues
[license-shield]: https://img.shields.io/github/license/ALTERA-CAPSTON-41/backend.svg?style=for-the-badge
[license-url]: https://github.com/ALTERA-CAPSTON-41/backend/blob/master/LICENSE
[semantic-badge]: https://img.shields.io/badge/semantic--release-angular-e10079?style=for-the-badge&logo=semantic-release
[semantic-url]: https://github.com/semantic-release/semantic-release
[codecov-shield]: https://img.shields.io/codecov/c/gh/ALTERA-CAPSTON-41/backend?label=CODECOV&logo=codecov&style=for-the-badge&token=YL5V1QWP2I
[codecov-url]: https://codecov.io/gh/ALTERA-CAPSTON-41/backend
