As a Caasp developer, I need to learn GO language. While I have an opportunity to learn GO language during HeckWeek, I want to combine GO, IAAS, and public Clouds(AWS). By this project, not only learning GO language but also advancing to utilize essential GO libraries and creating a GO project can be achieved. In addition, I need to familiarize AWS api to deploy nodes, stop the nodes, delete nodes, and upload images.

I can deploy nodes in cloud easily by this project. If I can create available nodes in AWS by this projects. Manually I can try to cluster the available AWS nodes for Kubernetes nodes.

``` 
Prerequisites: 
1. Learning GO language 
2. Learning AWS

Goal of this project in hackweek18:
1. deploying nodes in AWS using GO language

Post Project(if feasible): 
1. Manually trying to add AWS nodes in kubenetes cluster 
```

### Install asw libraries
```
go get -u github.com/aws/aws-sdk-go/aws  
go get -u github.com/aws/aws-sdk-go/aws/session  
go get -u github.com/aws/aws-sdk-go/service/ec2  

```
### In order to access aws, you need to add the following in .bashrc file
```
### AWS Access Through Go Libraries
export AWS_REGION=<your region>
export AWS_ACCESS_KEY_ID=<your access id>
export AWS_SECRET_ACCESS_KEY=<your secret access key>
```
source ~/.bashrc
How to find the id and key: https://help.bittitan.com/hc/en-us/articles/115008255268-How-do-I-find-my-AWS-Access-Key-and-Secret-Access-Key-


