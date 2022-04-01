<img align="right" width="159px" src="https://jackboxgames.b-cdn.net/wp-content/uploads/2019/10/No-Text.png" />

A [Jackbox Games](https://www.jackboxgames.com/) demo for a basic application developed in [Go](https://go.dev/) and [Vue](https://vuejs.org/).

## Table of Contents

- [Background](#background)
- [Install](#install)
- [Example](#example)
- [Issues](#issues)
- [Maintainers](#maintainers)

## Background

This simple application consists of a Go server and a Vue frontend. The application allows users to register for the site, login to the site, and view their dashboard.

The user can:
- Register for the site by providing their Email and Name.
- Login with their email.
- View their dashboard.

My goal with this app was mainly to show my apptitude for working on applications and back end services. The vue app is based off the sample app. I am comfortable working on front end but am much more so working on the backend.

Thank you for looking over this project. I appreciate your time.


## Install

Installation of the project is very simple.
1. Clone the repository to your local machine. - `git clone https://github.com/mca312/jackbox.git`
1. In the `jackbox` directory. Run `make start`
1. In a browser, navigate to [http://localhost:8447](http://localhost:8447)

The `make start` command is running `docker compose up -d`. This will build both the Go and Vue app in respective containers and start them. The Go app runs on port 8448 and the web app will run on 8447.

Two test users are automatically created with the email `angellotti@gmail.com` and `fake@user.com`.  These are setup in the `config.yml` file.

`make stop` will stop the docker containers by running `docker compose down`.
`make build` will build the docker images without starting the containers.
`make tests` will run the Go tests.

## Example

Note: There is no real authentication. I determine logged in by if the local storage has a valid user id. (I only have so much time in my life). If the user id is set, you can go to the dashboard. If you go to the dashboard with an id that hasn't been registered it will unset it and redirect to the home/login page.

After installing, navigate to [http://localhost:8447](http://localhost:8447).

- Try to log in with any email address. Server side validation will check if it exists
  - test@email.com returns `User does not exist`
- Log in with angellotti@gmail.com.
  - This will bring you to the dashboard, which only displays the users name and email address. This is sample data.
- Logout
  - This will clear the stored user id.
- Register a new user.
  - Click register and provide a name and new email address. Email addresses are unique and are server side validated.
  - Email address are validated to contain at least one @.
  - Names are validated to not be blank.

## Issues

This app was developed quickly so a few things that may seems weird.
- You can register an email as `email@`. As long as it has only one @, it works.
- The web app is a single page app, but if you refresh it on the `/dashboard` page, it will 404. Nginx isn't set up for it and my goal wasn't to spend time on it.
- To keep everything in one repo, I put the webapp in it's own directory. In a real world scenario, it would be in it's own repository.
- The dashboard only consists of the users name and email. It's a sad dashboard. :(

## Maintainers

[Michael Angellotti](https://github.com/mca312)
