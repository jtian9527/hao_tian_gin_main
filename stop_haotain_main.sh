ps -aux | grep hao_tian_gin_main | grep api |  grep -v grep  | awk '{print $2}' | xargs kill -9