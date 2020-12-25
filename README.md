# test-new-ci

## workfllow 原則

|name|env|db|
|-|-|-|
|LOCAL|本地|本地|
|DEV|測試|測試資料庫|
|STAGING|上線前測試 A/B |正式資料庫| 
|PROD|正式|正式資料庫|

Git branch 對照:

- main -> prod
- pr (dev->main) -> staging ( one time  just only one !)
- dev -> DEV
- pr (featrue->dev) -> local  

---

tool : istio / github action / kustomize

## TODO

- [x] 1 GKE with 3 StageMode
- [ ] cicd ( github action )
  - [ ] testing 單元測試
  - [ ] notify
  - [ ] notify button
- [ ] 提高 config 編輯性 ( josh )

demo Docker image:

- asia.gcr.io/cheep2workshop/zeki-k8s

## Doc

### DEVOPS 架構範例

```sh
├── deploy-k8s
│   ├── base
│   │   ├── configMap.yaml
│   │   ├── deploy.yaml
│   │   ├── kustomization.yaml
│   │   └── virtualservice.yaml
│   ├── init-istio
│   │   └── gateway.yaml
│   └── overlays
│       ├── dev
│       │   ├── config.yaml
│       │   └── kustomization.yaml
│       ├── prod
│       │   ├── config.yaml
│       │   └── kustomization.yaml
│       └── staging
│           ├── config.yaml
│           └── kustomization.yaml
```

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
...
```
