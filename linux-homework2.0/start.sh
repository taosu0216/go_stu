#!/bin/bash

# 获取所有用户列表，并过滤系统自带用户
user_list=$(awk -F: '$3 >= 1000 {print $1}' /etc/passwd)

# 遍历每个用户
for user in $user_list; do
    history_file="/home/$user/.bash_history"

    # 检查历史记录文件是否存在
    if [ -f "$history_file" ]; then
        echo "用户: $user"
        echo "最近的2条命令:"

        # 获取最近的2条命令并写入新文件
        tail -n 2 "$history_file" > "/tmp/recent_commands_$user.txt"
        cat "/tmp/recent_commands_$user.txt"

        echo "========================"
    else
        echo "用户: $user"
        echo "历史记录文件不存在"
        echo "========================"
    fi
done