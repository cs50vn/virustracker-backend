#!/usr/bin/env bash
#@echo off

export BUILD_TYPE=all

python3 scripts/generate-input-sheet.py $PWD linux $BUILD_TYPE
