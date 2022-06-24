# Integration Tests

## Configuration
You can specify the environment you wish to run the tests against by specifying configuration values in the `.env` in the root of this project. 


```bash
  API_ADDRESS         = ""
  API_ADDRESS_LEGACY  = ""
  API_TOKEN           = ""
  
  IDS_ADDRESS         = ""
  IDS_CLIENT_ID       = ""
  IDS_CLIENT_SECRET   = ""
  IDS_SCOPE           = ""
```

## Usage
Integration tests are run with the following command in the root directory of the project:

```
task integration_test:run
```

This will run all the tests within the integration test suite.

## Creating Tests
These are the steps for adding a new test to the integration test suite.


#### Creating the test directory

Create a directory for the resource you wish to test within the `/test/integration/directory` directory. 

```shell
$ ls -l resources
total 0
drwxr-xr-x   6 user  staff  192 Apr  1 00:00 customer
drwxr-xr-x   8 user  staff  256 Apr  1 00:00 customer_user
drwxr-xr-x   8 user  staff  256 Apr  1 00:00 dns
drwxr-xr-x   8 user  staff  256 Apr  1 00:00 edgecname
drwxr-xr-x   8 user  staff  256 Apr  1 00:00 origin
drwxr-xr-x  10 user  staff  320 Apr  1 00:00 rules_engine_policy
drwxr-xr-x   9 user  staff  288 Apr  1 00:00 waf
```
_a listing of integration tests at the time this documentation was created._


### Creating the tests
Within the newly-created directory, perform the following steps:

   1. **Request Parameter Files**: Create a json file for each request...
 ```shell
 $ ls -l 
 total 0
 drwxr-xr-x   6 user  staff  192 Apr  1 00:00 add_custom_ruleset.json
 drwxr-xr-x   8 user  staff  256 Apr  1 00:00 delete_custom_ruleset.json
 drwxr-xr-x   8 user  staff  256 Apr  1 00:00 get_all_custom_rulesets.json
 drwxr-xr-x   8 user  staff  256 Apr  1 00:00 get_custom_ruleset.json
 drwxr-xr-x   8 user  staff  256 Apr  1 00:00 update_custom_ruleset.json
 ```

   2. **Include Taskfile File**: create a `Taskfile.yaml` file. This should execute the necessary tests being sure to clean up after itself and to leave in a clean of state as possible. _An example of a `Taskfile.yaml` in `test/integration/waf/Taskfile.yaml`.
```Taskfile
# https://taskfile.dev
version: '3'


output: prefixed

vars:
  JSON_FILES: "test/integration/waf"


tasks:
  access_AddCustomRuleSet:
    dir: ../..
    cmds:
      - ./cli waf access AddCustomRuleSet -i='{{.JSON}}'
      #- ./cli waf access AddCustomRuleSet -f={{.JSON_PATH}}
    vars:
      JSON_PATH: "{{.JSON_FILES}}/add_custom_ruleset.json"
      JSON:
        sh: cat {{.JSON_PATH}}

  access_DeleteCustomRuleSet:
    dir: ../..
    cmds:
      - ./cli waf access DeleteCustomRuleSet -i='{{.JSON}}'
      #- ./cli waf access DeleteCustomRuleSet -f={{.JSON_PATH}}
    vars:
      JSON_PATH: "{{.JSON_FILES}}/delete_custom_ruleset.json"
      JSON:
        sh: cat {{.JSON_PATH}}

  access_GetAllCustomRuleSets:
    dir: ../..
    cmds:
      - ./cli waf access GetAllCustomRuleSets -i='{{.JSON}}'
      #- ./cli waf access GetAllCustomRuleSets -f={{.JSON_PATH}}
    vars:
      JSON_PATH: "{{.JSON_FILES}}/get_all_custom_rulesets.json"
      JSON:
        sh: cat {{.JSON_PATH}}

  access_GetCustomRuleSet:
    dir: ../..
    cmds:
      - ./cli waf access GetCustomRuleSet -i='{{.JSON}}'
      #- ./cli waf access GetCustomRuleSet -f={{.JSON_PATH}}
    vars:
      JSON_PATH: "{{.JSON_FILES}}/get_custom_ruleset.json"
      JSON:
        sh: cat {{.JSON_PATH}}

  access_UpdateCustomRuleSet:
    dir: ../..
    cmds:
      - ./cli waf access UpdateCustomRuleSet -i='{{.JSON}}'
      #- ./cli waf access UpdateCustomRuleSet -f={{.JSON_PATH}}
    vars:
      JSON_PATH: "{{.JSON_FILES}}/update_custom_ruleset.json"
      JSON:
        sh: cat {{.JSON_PATH}}
  run:
    desc: "execute the waf test suite"
    cmds:
      - task: access_AddCustomRuleSet
      - task: access_GetAllCustomRuleSets
      - task: access_DeleteCustomRuleSet
      - task: access_GetCustomRuleSet
      - task: access_UpdateCustomRuleSet

  default:
    cmds:
      - task: run
```