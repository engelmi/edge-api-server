# Kubernetes Edge API Server

## Development

List available make commands:
```bash
$ make help

Usage:
  make <target>

General
  help             Display this help.

Development
  fmt              Run go fmt against code.
  vet              Run go vet against code.
  test             Run tests.

Build
  build-manager    Build manager binary.
  build-apiserver  Build manager binary.
  container-build-manager  Build container images with the manager.
  container-build-apiserver  Build container images with the apiserver.

Deployment
  install          Install CRDs into the K8s cluster specified in ~/.kube/config.
  uninstall        Uninstall CRDs from the K8s cluster specified in ~/.kube/config. Call with ignore-not-found=true to ignore resource not found errors during deletion.
  deploy           Deploy controller to the K8s cluster specified in ~/.kube/config.
  undeploy         Undeploy controller from the K8s cluster specified in ~/.kube/config. Call with ignore-not-found=true to ignore resource not found errors during deletion.
  kustomize        Download kustomize locally if necessary.
  envtest          Download envtest-setup locally if necessary.
```

When generating code, OpenAPI violations are tracked in [./known-openapi-violations.list](./known-openapi-violations.list). In order to update the list run code generation like this:
```bash
UPDATE_API_KNOWN_VIOLATIONS=true make controller-gen
```
