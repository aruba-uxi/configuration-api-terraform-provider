#!/usr/bin/env bash

set -e

os="linux"
arch="amd64"
repo="HPE/terraform-provider-hpeuxi"
linux_hpeuxi_dir="${HOME}/.local/share/terraform/plugins/registry.terraform.io/aruba-uxi/hpeuxi"

get_latest_release () {
  local release_url="https://api.github.com/repos/${repo}/releases/latest"
  curl -sL "$release_url" | grep -Po '"tag_name": "\K.*?(?=")'
}

download_and_extract () {
  local dest_dir="${linux_hpeuxi_dir}/${version_number}/${os}_${arch}/"
  local hpeuxi_zip="terraform-provider-hpeuxi_${version_number}_${os}_${arch}.zip"
  local hpeuxi_dl_url="https://github.com/${repo}/releases/download/${VERSION}/${hpeuxi_zip}"

  mkdir -p "$dest_dir" && cd "$dest_dir"
  curl -sL "$hpeuxi_dl_url" -o "$hpeuxi_zip" && \
    unzip -u "$hpeuxi_zip" && \
    rm -f "$hpeuxi_zip"
}

VERSION=${VERSION:=$(get_latest_release)}
version_number=${VERSION//v}
download_and_extract