#!/usr/bin/env zsh
# -*- coding: utf-8 -*-
  # shellcheck shell=bash
  # shellcheck source=/dev/null
  # shellcheck disable=2178,2155

. $(which ansi_colors.sh)

setopt NO_ALL_EXPORT

TEMPLATE_DIR=~/Documents/coding/cc_template

ZDOTDIR="${ZDOTDIR:=~/.dotfiles}"


TEMPLATE_DIR=~/go/src/github.com/skeptycal/gorepo/template
REPO_PATH=$PWD

REPO_NAME=${REPO_PATH##*/}

COPY_DIR=${TEMPLATE_DIR}/copy
AUTO_DIR=${TEMPLATE_DIR}/automated
MOD_DIR=${TEMPLATE_DIR}/mod

export SIMPLE_BACKUP_SUFFIX="bak"
COPY_OPTIONS="-al --strip-trailing-slashes"


cp -lr ${COPY_DIR}/* $REPO


setopt ALL_EXPORT
