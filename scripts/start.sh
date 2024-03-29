#!usr/bin/env bash
# 只要出错就退出
set -o errexit

# 只要有未定义的变量就退出
# set -o nounset

# 只要管道中任何一个命令失败就退出
set -o pipefail

# 开启shell调试模式，打印命令和输出
set -x

# 导入公用脚本，并检查是否执行成功
source service.sh
if [ $? -ne 0 ]; then
    echo "Failed to load service.sh"
fi

# 检查当前执行脚本的用户
current_user=$(whoami)
if [[ "${current_user}" != "$user" ]]; then
    log "The current user ${current_user} is not ${user} as expected"
    exit 1
fi

log "Begin to build the package"
source build.sh

log "Build completed! Begin to run!"
# 启动前检查进程是否存在
if ps aux | grep -v grep | grep "$pkg_name" > /dev/null; then
    log "Old process $pkg_name is running. Killing the pid..."
    
    # 获取进程ID并杀死进程
    process_id=$(ps aux | grep -v grep | grep "$pkg_name" | awk '{print $2}')
    kill -9 "$process_id"
    
    log "Process $pkg_name killed."
fi

# 在后台启动进程
nohup ${pkg_name} 2>&1 | awk '{ print strftime("%Y-%m-%d %H:%M:%S|SYSTEM|"), $0; fflush(); }' >> $log_path &
