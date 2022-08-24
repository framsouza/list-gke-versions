# List GKE version

This code will print the current version of the masters and nodes running on GKE.

### Usage
First, make sure you have the `GOOGLE_APPLICATION_CREDENTIALS` environment variable set as per the google docs.

It will require one argument which is `project`.

You can choose to build the binary by running:

```
go build -o list-gke-version
```

and run it as:

```
./list-gke-version -project=<PROJECTNAME>
```

or without buidling the binary

```
go run main.go -project=<PROJECTNAMEs>
```

### Output expected

```
CLUSTER NAME                             MASTER VERSION            NODE VERSION
mygke-apps-global-eu                   1.21.12-gke.2200          1.21.12-gke.2200          
test-gke                               1.21.12-gke.2200          1.21.12-gke.2200          
mygke-apps-cd-stag                     1.21.13-gke.900           1.21.13-gke.900           
mygke-apps-cd-prod                     1.21.12-gke.2200          1.21.12-gke.2200          
mygke-apps-prod                        1.21.12-gke.2200          1.21.12-gke.2200          
mygke-apps-devs-aac                    1.21.13-gke.900           1.21.13-gke.900           
mygke-apps-es                          1.21.13-gke.900           1.21.13-gke.900           
mygke-apps-global                      1.21.12-gke.2200          1.21.12-gke.2200          
mygke-apps-nick                        1.22.10-gke.600           1.22.10-gke.600           
mygke-apps-nick                        1.21.12-gke.2200          1.16.15-gke.4901          
mygke-apps-rele                        1.21.12-gke.2200          1.21.12-gke.2200          
mygke-apps-stage                       1.21.12-gke.2200          1.21.12-gke.2200          
mygke-apps-ops                         1.21.12-gke.2200          1.21.12-gke.2200          
mygke-apps-web                         1.21.12-gke.2200          1.21.12-gke.2200          
mygke-apps-au                          1.21.12-gke.2200          1.21.12-gke.2200          
mygke-apps-jarpy                       1.21.13-gke.900           1.21.13-gke.900 %     
```