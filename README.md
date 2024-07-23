# S3 Lambda Trigger

An AWS lambda listening to events from S3 and running Datamin pipelines using the uploaded file content as input.

![GitHub branch check runs](https://img.shields.io/github/check-runs/datamin-io/s3-lambda-trigger/main?color=green)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/datamin-io/s3-lambda-trigger?color=blue)
<a href="https://github.com/datamin-io/ylem?tab=Apache-2.0-1-ov-file">![Static Badge](https://img.shields.io/badge/license-Apache%202.0-blue)</a>
<a href="https://datamin.io" target="_blank">![Static Badge](https://img.shields.io/badge/website-datamin.io-blue)</a>
<a href="https://docs.datamin.io" target="_blank">![Static Badge](https://img.shields.io/badge/documentation-docs.datamin.io-blue)</a>
<a href="https://join.slack.com/t/datamincommunity/shared_invite/zt-2nawzl6h0-qqJ0j7Vx_AEHfnB45xJg2Q" target="_blank">![Static Badge](https://img.shields.io/badge/community-join%20Slack-blue)</a>

## How it works

After installation, the lambda listens to S3 events about created objects (i.e. uploaded files). If the file path matches the configured expression, the lambda reads their contents and runs the configured pipelines for each file.

## Installation

### Pre-requisites

Create an OAuth client for the lambda here: https://app.datamin.io/api-clients and copy the client ID and client secret key.

### Method 1: From Zip archive

Follow [this guide](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html) to install the lambda from a zip archive.

### Method 2: From Datamin container registry

1. Navigate to AWS Lambda → Functions section.
2. Click "Create function".
3. Choose "Container image".
4. Enter function name.
5. Enter container image URI: `241485570393.dkr.ecr.eu-central-1.amazonaws.com/datamin-integration:latest`
6. Architecture: `x86_64`.
7. Execution role: pick `Create a new role with basic Lambda permissions`.

## Configuration

First, make sure that the lambda has read access to objects in all S3 buckets you plan to use.

Second, follow [this guide](https://docs.aws.amazon.com/lambda/latest/dg/with-s3-example.html) to set up a S3 trigger for the lambda.

Then, you need to configure several environment variables for the lamba to work.

Navigate to Configuration → Environment variables section of the lambda.

| Variable | Description |
|--|-|
| **DTMN_API_CLIENT_ID** | OAuth client ID. Create a new client here: https://app.datamin.io/api-clients |
| **DTMN_API_CLIENT_SECRET** | OAuth client secret |
| **DTMN_S3_MAPPING** | Mapping of path expressions to pipeline UUIDs. See below. |

### Mapping

The format is:

`<path expression #1>`:`<pipeline UUID #1>,<pipeline UUID #2>`;`<path expression #2>`:`<pipeline UUID #3>`

**Path expression**: A Glob-like expression to match the full object path against, including the bucket name.

If the object path matches the expression, lambda will read the object contents and send them as input to all pipelines which UUIDs are enumerated in the config.

**Pipeline UUIDs:** a comma-separated list of pipelines to run if the object path matches the expression.

#### Mapping example
- `my-bucket-name/logs/*.json:18d27bb8-f886-4072-9165-485baa3e332d;my-bucket-name/*/daily.json:2f0cfe9e-3571-49a4-9103-cbd9d3d58c2b`
