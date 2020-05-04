#!/usr/bin/env bash
#@echo off

#Current dir
#Host type
#Build type: release, debug

export BUILD_TYPE=$1

python3 scripts/build.py $PWD linux $BUILD_TYPE

