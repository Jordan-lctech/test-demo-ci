# test-new-ci

istio / github action / kustomize

## TODO

DEV: 測試 STAGING: 準測試 PROD 正式

- [ ] cicd ( github action )
  - [ ] testing 單元測試
  - [ ] notify
  - [ ] notify button
- [ ] 提高 config 編輯性
- [ ] 1 GKE with 3 StageMode

demo Docker image:

- asia.gcr.io/cheep2workshop/zeki-k8s

## Doc

### DEVOPS 架構範例

demo architectur

### demo APP 說明

webserver ( .../hello ) -> print env mode

## local golang

```sh
go run .
# RUN_MODE = ""
```

## local docker

```sh
docker build -t test .
docker run -p8080:8080 --rm test
# RUN_MODE = "docker"
```

## k8s dev

```sh
kubectl apply -k deploy-k8s/overlays/dev
# RUN_MODE = "k8s-dev"
```
