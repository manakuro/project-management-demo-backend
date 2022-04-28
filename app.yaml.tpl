runtime: go116
main: ./cmd/app

handlers:
  - url: /.*
    script: auto
    secure: always

env_variables:
  APP_ENV: ${APP_ENV}
