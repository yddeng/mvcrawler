ps -ef | grep httpSev | grep -v grep | awk '{print $2}' | xargs kill -9