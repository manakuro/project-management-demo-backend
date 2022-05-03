runtime: go116
main: ./cmd/app

automatic_scaling:
  min_idle_instances: automatic
  max_idle_instances: 1
  min_pending_latency: 3000ms
  max_pending_latency: automatic
  max_instances: 2

env: standard
instance_class: F1

handlers:
  - url: /.*
    script: auto
    secure: always

  - url: /favicon.ico
    static_files: favicon.ico
    upload: favicon.ico

env_variables:
  APP_ENV: ${APP_ENV}
