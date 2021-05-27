#!/bin/sh

# PROVIDE: shumfyid
# REQUIRE: DAEMON
# KEYWORD: shutdown

. /etc/rc.subr

name="shumfyid"

rcvar="${name}_enable"
load_rc_config ${name}
: ${shumfyid_enable:=NO}

pidfile="/var/run/${name}.pid"
logfile="/var/log/${name}.log"
procname="/usr/local/bin/${name}"

command="/usr/sbin/daemon"
command_args="-p ${pidfile} -o ${logfile} ${procname}"

run_rc_command "$1"
