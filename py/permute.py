def permute(nums):
    """
    基础回溯法：数字全排列（无重复数字）
    :param nums: 输入无重复数字数组
    :return: 所有排列构成的二维列表
    """
    res = []          # 存储最终所有排列结果
    path = []         # 存储当前递归的路径（单次排列）
    used = [False] * len(nums)  # 标记数组：记录数字是否被选中
    
    def backtrack():
        # 递归终止条件：路径长度等于原数组长度 → 记录排列
        if len(path) == len(nums):
            res.append(path.copy())  # 注意：必须拷贝，否则会被后续修改
            return
        
        # 遍历所有数字，尝试选择
        for i in range(len(nums)):
            if not used[i]:  # 跳过已被选中的数字
                used[i] = True       # 标记：当前数字被选中
                path.append(nums[i]) # 将数字加入路径
                backtrack()          # 递归：继续选择下一个数字
                path.pop()           # 回溯：撤销选择，移出路径
                used[i] = False      # 回溯：取消标记
    
    backtrack()
    return res

def permute_swap(nums):
    """
    优化回溯法：原地交换实现全排列（无额外空间开销，最优）
    :param nums: 输入无重复数字数组
    :return: 所有排列构成的二维列表
    """
    res = []
    n = len(nums)
    
    def backtrack(start):
        # 终止条件：起始索引到末尾 → 记录当前排列
        if start == n:
            res.append(nums.copy())  # 拷贝当前数组（已完成排列）
            return
        
        # 从start开始，遍历所有未选数字
        for i in range(start, n):
            nums[start], nums[i] = nums[i], nums[start]  # 交换：选当前数字到start位置
            backtrack(start + 1)                         # 递归：处理下一个位置
            nums[start], nums[i] = nums[i], nums[start]  # 回溯：交换回原位置
    
    backtrack(0)
    return res


# ========== 测试用例 ==========
if __name__ == "__main__":
    nums1 = [1, 2, 3]
    print(f"数组 {nums1} 的全排列：\n{permute(nums1)}")
    
    nums2 = [0, 1]
    print(f"\n数组 {nums2} 的全排列：\n{permute(nums2)}")

    nums = [1,2,3]
    print(f"原地交换版-数组 {nums} 的全排列：\n{permute_swap(nums)}")
