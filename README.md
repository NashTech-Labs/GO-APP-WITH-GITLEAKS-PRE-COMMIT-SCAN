
  

# HTTP SERVER USING GO WITH GITLEAK PRE-COMMIT HOOK

  

A simple REST API built with Go and the Gin web framework to manage a list of Users and giteask pre-commit hook.

  

## Endpoints

  

*  `GET /get_users` – Retrieve all users

*  `POST /add_user` – Add a new user

  

## Requirements

  

* Go 1.16+

* Gin web framework

  

## Run the Application

  

```bash

go  build && ./day1

```

  

Server runs at `http://localhost:8000`

## Working with gitleaks
To implemnt git leak in my application I have created a pre-commit hook that will detect the secret before the commit reducing the chance to push the secret values to source code repository by mistake. It will not allow user to commit their changes to source code as shown in image (/screenshot/git-leaks-scan-with-leak.png).

1. Install the gitleaks by running `./gitleaksinstall.sh`
2. I have created a pre-commit script using python (filename : pre-commit) 
3. Move pre-commit file to .git/hooks/ directory 
4. Now when ever you run git commit the hook will run  first resulting in detection of secrets.
