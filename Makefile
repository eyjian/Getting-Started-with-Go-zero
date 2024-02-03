# Written by yijian on 2024/02/03

.PHONY: tidy fetch

tidy:
	go mod tidy

fetch: # 强制用远程仓库的覆盖本地，运行时需指定分支名，如：make fetch branch=main
	git fetch --all&&git reset --hard origin/$$branch
