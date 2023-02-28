# Skillathon demo app

It's basically a dummy app that has a few handlers useful for Kubernetes setup.

## Handlers

- `/health/liveness` and `/health/readiness`:
  - Those hadnlers imitate liveness and readiness health checks;
  - Can be used for liveness/readiness probes;
  - In many apps, there will be just a single `/healthz` endpoint that is used for both probes;
- `/configure/liveness`:
  - It makes `/health/liveness` start returning non-2xx status code, essentially imitating the app failure;
  - The second call to the same endpoint reverses behaviour;
- `/configure/readiness`:
  - It makes `/health/readiness` start returning non-2xx status code, essentially imitating the app failure;
  - The second call to the same endpoint reverses behaviour;
- `/metrics`:
  - This handler exposes a basic set of runtime-metrics + number of served requests;
  - Useful to play with ServiceMonitor CR (prometheus-operator);
- `/shutdown`:
  - A special handler for a `preStop` hook ([docs](https://kubernetes.io/docs/tasks/configure-pod-container/attach-handler-lifecycle-event/));
  - When a pod is being terminated, the main process in each container receives a `SIGTERM` signal. The best practice is to make sure on-the-fly requests are still handled while the new requests are not sent to the app. The first point is natively covered by Go's `(*http.Server).Shutdown(ctx context.Context)`, the second - by a network plugin installed in a cluster meaning the respective service endpoints are being taken out of load-balancing. When a cluster is big enough, it takes a little while before the routing is synchronized across all nodes. And here, optional `preStop` hook (HTTP or exec) comes into play. - It gets executed by `kubelet` before `SIGTERM` is sent while the pod is already being taken out of load-balancing. As `SIGTERM` is sent only after the hook is completed, by introducing a few seconds delay (e.g. through `time.Sleep`), we can make sure that the requests, which slipped through before the routing is fully synchronized, are still served.
- `/slow`:
  - Imitates a handler that returns data for a slow request to a database;
- `/`:
  - Just a "hello world"-like handler.

## OCI-image

At the time of writing, the app is published under `quay.io/weisdd/skillathon:0.1.0`, though it's not guranteed to stay there in the long run.
