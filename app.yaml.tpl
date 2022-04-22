runtime: go116

handlers:
  - url: /.*
    script: auto
    secure: always

env_variables:
  APP_ENV: ${APP_ENV}
