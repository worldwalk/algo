def coinChange_backtrack(coins, amount):
    """
    回溯法求解凑零钱：返回所有可能的组合（无重复）
    :param coins: 硬币面额数组
    :param amount: 目标金额
    :return: 所有组合构成的二维列表，无解返回[]
    """
    res = []  # 存储最终所有有效组合
    path = []  # 存储当前递归的组合路径

    def backtrack(start, remain_amount):
        # 递归终止条件1：凑出目标金额 → 记录组合
        if remain_amount == 0:
            res.append(path.copy())
            return
        # 递归终止条件2：金额超了 → 剪枝，直接返回
        if remain_amount < 0:
            return
        
        # 遍历硬币：从start开始，避免重复组合（[1,2]和[2,1]算一种）
        # 关键点：循环从start开始而不是从0开始的原因
        # 1. 避免生成重复组合：例如[1,2]和[2,1]都能凑成3，但它们是相同的组合
        # 2. 确保组合的顺序是递增的，即每个后续选择的硬币索引不小于当前索引
        # 3. 同时允许同一硬币多次使用（因为下一次递归传入的start还是i，不是i+1）
        for i in range(start, len(coins)):
            coin = coins[i]
            path.append(coin)  # 选择当前硬币
            # 递归：剩余金额减少，起始下标不变（硬币可重复使用）
            # 如果改为i+1，则表示每个硬币只能使用一次
            backtrack(i, remain_amount - coin)
            path.pop()  # 回溯：撤销选择
    
    if amount == 0:
        return [[]]
    # 排序后剪枝效果更好（可选，不影响正确性）
    coins.sort()
    backtrack(0, amount)
    return res if res else []

def coinChange_dp(coins, amount):
    """
    动态规划求解凑零钱：返回最少硬币个数，无解返回-1
    :param coins: 硬币面额数组
    :param amount: 目标金额
    :return: 最少硬币数 / -1
    """
    # 步骤1：初始化dp数组，无穷大表示初始不可达
    INF = float('inf')
    dp = [INF] * (amount + 1)
    dp[0] = 0  # 基准：金额0需要0枚硬币

    # 步骤2：遍历所有金额，填充dp数组
    for i in range(1, amount + 1):
        # 遍历所有硬币，尝试更新dp[i]
        for coin in coins:
            if coin <= i:  # 硬币面额不能超过当前金额
                dp[i] = min(dp[i], dp[i - coin] + 1)

    # 步骤3：返回结果，判断是否可达
    return dp[amount] if dp[amount] != INF else -1
def coinChange_dp_with_path(coins, amount):
    """
    动态规划：返回（最少硬币数，最优组合），无解返回(-1, [])
    """
    INF = float('inf')
    dp = [INF] * (amount + 1)
    dp[0] = 0
    path = [None] * (amount + 1)  # 记录每个金额最后选的硬币

    for i in range(1, amount + 1):
        for coin in coins:
            if coin <= i and dp[i - coin] + 1 < dp[i]:
                dp[i] = dp[i - coin] + 1
                path[i] = coin  # 更新最后选择的硬币

    # 回溯获取最优组合
    best_combination = []
    if dp[amount] != INF:
        current = amount
        while current > 0:
            coin = path[current]
            best_combination.append(coin)
            current -= coin

    return (dp[amount], best_combination) if dp[amount] != INF else (-1, [])

# 调用示例
coins = [1,2,5]
amount =11
min_num, comb = coinChange_dp_with_path(coins, amount)
print(f"最少硬币数：{min_num}，最优组合：{comb}")
# 输出：最少硬币数：3，最优组合：[1,5,5]
