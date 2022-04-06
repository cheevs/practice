## Overview
This is a simple practice application that receives an http post request 
and responds with payload, headers, environment variables and custom nginx response headers

Nginx is configured to return two custom headers in the response.

* `X-Request-Start` - contains the timestamp
* `X-URI` - contains the path

Nginx is configured with two config maps, `ingress-nginx-controller` and `custom-headers`

HTTP POST response
```json 
{
  "payload": {}, //payload of the request
  "headers": {}, //headers of the request
  "environmentVariables": [] // environment variable containing "PRACTICE"
}
```

### Kubernetes Resources
The application consists of a deployment, service and ingress resources.
The ingress specifies the path of `/` on port `80`to route all traffic to the service `practice`.

### Requirements
minikube v1.23.2+

### Minikube configuration
Minikube needs to be started with the vm option to allow for the nginx ingress addon to be installed

`minikube start --vm=true --driver=hyperkit`

Next enable the minikube addon.
Note: There can be errors if on a minikube version less than 1.23.2

`minikube addons enable ingress`

### Installation
The docker build uses a multi stage build for compiling the application and creating the optimized image

`docker build . -t practice:1.0`

Loading the image directly into minikube registry

`minikube image load practice:1.0`

Applying the resources into the default namespace

`kubectl --context minikube apply -f kubernetes-resources/`

### Example Usage
```bash
curl -v --location --request POST $(minikube ip):80/foo/bar \
--header 'Content-Type: application/json' \
--data-raw '{
    "foo": "bar"
}'
```

Response: 

Notice the additional headers: `X-Request-Start: t=1649212463.284` `X-URI: /foo/bar`

```
Note: Unnecessary use of -X or --request, POST is already inferred.
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying 192.168.64.4...
* TCP_NODELAY set
* Connected to 192.168.64.4 (192.168.64.4) port 80 (#0)
> POST / HTTP/1.1
> Host: 192.168.64.4
> User-Agent: curl/7.64.1
> Accept: */*
> Content-Type: application/json
> Content-Length: 20
>
} [20 bytes data]
* upload completely sent off: 20 out of 20 bytes
< HTTP/1.1 200 OK
< Date: Wed, 06 Apr 2022 02:34:23 GMT
< Content-Type: application/json
< Content-Length: 751
< Connection: keep-alive
< X-Request-Start: t=1649212463.284
< X-URI: /
<
{ [751 bytes data]
100   771  100   751  100    20   366k  10000 --:--:-- --:--:-- --:--:--  376k
* Connection #0 to host 192.168.64.4 left intact
* Closing connection 0
{
  "payload": {
    "foo": "bar"
  },
  "headers": {
    "Accept": [
      "*/*"
    ],
    "Content-Length": [
      "20"
    ],
    "Content-Type": [
      "application/json"
    ],
    "User-Agent": [
      "curl/7.64.1"
    ],
    "X-Forwarded-For": [
      "192.168.64.1"
    ],
    "X-Forwarded-Host": [
      "192.168.64.4"
    ],
    "X-Forwarded-Port": [
      "80"
    ],
    "X-Forwarded-Proto": [
      "http"
    ],
    "X-Forwarded-Scheme": [
      "http"
    ],
    "X-Real-Ip": [
      "192.168.64.1"
    ],
    "X-Request-Id": [
      "06599088d2aa0ec7f1ed192ffbd3df17"
    ],
    "X-Scheme": [
      "http"
    ]
  },
  "environmentVariables": [
    "PRACTICE_FOO=FOO",
    "PRACTICE_BAR=BAR",
    "PRACTICE_PORT_80_TCP_PORT=80",
    "PRACTICE_SERVICE_HOST=10.102.237.77",
    "PRACTICE_SERVICE_PORT=80",
    "PRACTICE_SERVICE_PORT_HTTP=80",
    "PRACTICE_PORT=tcp://10.102.237.77:80",
    "PRACTICE_PORT_80_TCP_ADDR=10.102.237.77",
    "PRACTICE_PORT_80_TCP_PROTO=tcp",
    "PRACTICE_PORT_80_TCP=tcp://10.102.237.77:80"
  ]
}
```