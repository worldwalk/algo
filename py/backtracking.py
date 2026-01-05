# 经典回溯问题：给定数组[2,3,6,7]和目标值7，找到所有和为 7 的组合（无重复组合）。
def combination_sum(candidates, target):
    result = []  # 存储最终解
    # 回溯函数：path=当前组合，start=起始索引（避免重复组合），current_sum=当前和
    def backtrack(path, start, current_sum):
        # 1. 终止条件1：找到有效解
        if current_sum == target:
            result.append(path.copy())
            return
        # 2. 终止条件2：剪枝（当前和超过目标，无需继续）
        if current_sum > target:
            return
        # 3. 遍历所有可能的分支
        for i in range(start, len(candidates)):
            # 选择当前元素
            path.append(candidates[i])
            # 递归深入（start=i：避免重复组合，如[2,3]和[3,2]视为同一组合）
            backtrack(path, i, current_sum + candidates[i])
            # 回溯：撤销选择（核心！回退到上一步，尝试下一个分支）
            path.pop()
    # 初始化回溯
    backtrack([], 0, 0)
    return result

# 测试
print(combination_sum([2,3,6,7], 7))  # 输出：[[2,2,3], [7]]

