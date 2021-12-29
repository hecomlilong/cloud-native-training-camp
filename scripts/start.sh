#!/bin/bash
arch=$(go env GOARCH)
./cmd/linux/$arch/mod$1
