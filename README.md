# Fetch RH Pull-Secret 

Use REST API to get RH Pull-Secret. \
Requires a Red Hat user account. \
Used to simplify the solution from KCS4844461 "How to download the pull secret from cloud.redhat.com/openshift/install/pull-secret using a REST API call?" \


Docs: 
* https://docs.openshift.com/container-platform/4.16/openshift_images/managing_images/using-image-pull-secrets.html 

### Access to Red Hat knowledgebase requires you to login with your Red Hat account. 
* https://access.redhat.com/solutions/4844461



### How to use. 
* Get your offline token from https://console.redhat.com/openshift/token, please read https://access.redhat.com/solutions/4844461 for more information.
~~~
$ git clone backchristoffer/get_openshift_pull_secret
$ cd get_openshift_pull_secret
$ echo 'OFFLINE_ACCESS_TOKEN="offline token"' > .env OR export OFFLINE_ACCESS_TOKEN="offline token"
$ go build .
$ ./ocps
~~~
#### Output example
~~~
{
  "auths": {
    "cloud.openshift.com": {
      "auth": "<numbersandtext>",
      "email": "emai@email.com"
    },
    "quay.io": {
      "auth": "<numbersandtext>",
      "email": "emai@email.com"
    },
    "registry.connect.redhat.com": {
      "auth": "<numbersandtext>",
      "email": "emai@email.com"
    },
    "registry.redhat.io": {
      "auth": "<numbersandtext>",
      "email": "emai@email.com"
    }
  }
}
~~~
