#!/usr/bin/env bash

# Should be run only after all manifests have been deleted
export SERVICE_PRINCIPAL_ID=$(az ad sp list --all  --output tsv \
    --query "[?displayName=='containerazure-gcp-upbound'].{id:id}")
az ad sp delete --id "${SERVICE_PRINCIPAL_ID}"

export APPLICATION_ID=$(az ad app list --all \
 --query "[?displayName=='containerazure-gcp-upbound'].appId" \
 --output tsv)

az ad app create --id "${APPLICATION_ID}"
