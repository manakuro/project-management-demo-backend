runtime: custom

env: flex

automatic_scaling:
  min_num_instances: 1
  max_num_instances: 1

readiness_check:
  path: "/readiness_check"

handlers:
  - url: /.*
    script: auto
    secure: always

  - url: /favicon.ico
    static_files: favicon.ico
    upload: favicon.ico

resources:
  cpu: 0.5
  memory_gb: 0.5
  disk_size_gb: 0.5

env_variables:
  APP_ENV: ${APP_ENV}
