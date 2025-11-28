.PHONY: help sqlc goose up status down clean

RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[1;33m
BLUE=\033[0;34m
PURPLE=\033[0;35m
CYAN=\033[0;36m
NC=\033[0m # No Color

# 默认目标
help:
	@echo "${BLUE}Usage:${NC}"
	@echo "  ${GREEN}make sqlc${NC}     - 使用 sqlc 根据 SQL 查询生成 Go 代码"
	@echo "  ${GREEN}make goose${NC}    - 运行 goose 数据库迁移工具"
	@echo "  ${GREEN}make up${NC}       - 执行所有待处理的数据库迁移"
	@echo "  ${GREEN}make status${NC}   - 查看当前数据库迁移状态"
	@echo "  ${GREEN}make down${NC}     - 回滚最后一次数据库迁移"
	@echo "  ${GREEN}make clean${NC}    - 清理生成的代码文件"

# 使用 sqlc 生成代码
sqlc:
	@echo "${BLUE}正在使用 sqlc 根据 SQL 查询生成 Go 代码...${NC}"
	sqlc generate
	@echo "${GREEN}代码生成完成！${NC}"

# 应用所有可用的迁移
up:
	@echo "${BLUE}正在执行所有待处理的数据库迁移...${NC}"
	goose up
	@echo "${GREEN}所有数据库迁移已成功应用！${NC}"

# 显示迁移状态
status:
	@echo "${BLUE}正在查看数据库迁移状态...${NC}"
	goose status

# 回滚最近一次迁移
down:
	@echo "${BLUE}正在回滚最后一次数据库迁移...${NC}"
	goose down
	@echo "${GREEN}数据库迁移回滚成功！${NC}"

# 清理生成的文件
clean:
	@echo "${YELLOW}正在清理生成的代码文件...${NC}"
	rm -f db/*.go
	@echo "${GREEN}代码文件清理完成！${NC}"