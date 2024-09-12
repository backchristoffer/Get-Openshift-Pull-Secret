# Fetch RH Pull-Secret 
### This is NOT supported by Red Hat - This was created for my own personal use. Use this code at your own risk.

Use REST API to get RH Pull-Secret \
Requires a Red Hat user account. \
Used to simplify the solution from KCS4844461 "How to download the pull secret from cloud.redhat.com/openshift/install/pull-secret using a REST API call?" \


Docs: 
* https://docs.openshift.com/container-platform/4.16/openshift_images/managing_images/using-image-pull-secrets.html 

### Access to Red Hat knowledgebase requires you to login with your Red Hat account. 
* https://access.redhat.com/solutions/4844461



### How to use. 
~~~
$ git clone backchristoffer/get_openshift_pull_secret
$ cd get_openshift_pull_secret
$ echo <offline token> > .env
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