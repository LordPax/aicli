#!/bin/bash

[ "$EUID" -ne 0 ] && echo "Please run as root" && exit 1

function installFunc() {
	[ ! -f aicli ] && echo "aicli not found" && exit 1
	echo "Installing aicli to /usr/bin/aicli"
	install -Dm 755 aicli /usr/bin/aicli
}

function uninstall() {
	echo "Uninstalling aicli from /usr/bin/aicli"
	rm -f /usr/bin/aicli
}

case "$1" in
	help) echo "Usage: $0 {install|uninstall}" ;;
	uninstall) uninstall ;;
	*) installFunc ;;
esac
