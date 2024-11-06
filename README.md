# Tufin Client

This repository contains a Go client called `tufin` for deploying and managing a simple Kubernetes cluster with a WordPress application.

## Prerequisites

Before running the `tufin` client, ensure you have the following installed:

* **Go:**  A recent version of Go (e.g., 1.18 or later).
* **k3d:**  A lightweight wrapper for `k3s` for creating Kubernetes clusters. 
  * Install it using Homebrew

```bash
brew install k3d
```

* kubectl: The Kubernetes command-line tool. You can usually install this with k3d.


1. Clone this git repo
```bash
git clone https://github.com/shay79il/Tufin
```

2. Build the tufin client:
```bash
go build -o tufin
```

3. Create k3s cluster:
```bash
./tufin cluster 
```

4. Deploy resources:
`This command deploys a MySQL and a WordPress pod in the cluster.`
```bash
./tufin deploy 
```

5. Get status:
`This command displays a table with the names and statuses of the pods in the default namespace.`
```bash
./tufin status
```


6. Access the WordPress application at 
`This command forwards port 8080 on your local machine to port 80 of the WordPress service.`
```bash
kubectl port-forward svc/wordpress 8080:80
```

7. Open browser and go to  [http://localhost:8080](http://localhost:8080)



### Persistent Storage (Optional)
We could use persistent storage for the MySQL database. 


### License
This project is licensed under the MIT License. 

