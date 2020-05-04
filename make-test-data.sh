#!/usr/bin/env bash
#@echo off

export BUILD_TYPE=all

python3 scripts/make-test-data.py $PWD linux $BUILD_TYPE
