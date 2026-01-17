import numpy as np
from sklearn.metrics.pairwise import cosine_similarity, linear_kernel
from sklearn.preprocessing import normalize

def calculate_mean_similarity(cluster_vectors: np.ndarray, new_vector: np.ndarray) -> float:
    """
    计算聚类向量与新向量的平均余弦相似度（兼容cosine_similarity/linear_kernel）
    前提：输入已做L2归一化
    """
    # 确保输入为2维矩阵
    cluster_vectors = np.atleast_2d(cluster_vectors)
    new_vector = np.atleast_2d(new_vector)
    
    # 校验维度匹配
    if cluster_vectors.shape[1] != new_vector.shape[1]:
        raise ValueError(f"维度不匹配：聚类向量{cluster_vectors.shape}，新向量{new_vector.shape}")
    # 校验空数组
    if cluster_vectors.size == 0:
        return 0.0
    
    # 方法1：cosine_similarity（自动归一化，但若输入已归一化，结果等价于linear_kernel）
    sim_cosine = cosine_similarity(new_vector, cluster_vectors)
    mean_cosine = sim_cosine.mean()
    
    # 方法2：linear_kernel（输入已归一化时，结果与cosine_similarity一致）
    sim_linear = linear_kernel(new_vector, cluster_vectors)
    mean_linear = sim_linear.mean()
    
    return mean_cosine, mean_linear  # 返回两种方法的结果，用于对比


def test_similarity_consistency():
    """测试L2归一化后，cosine_similarity和linear_kernel结果一致"""
    # 测试场景1：基础场景（3个3维归一化向量）
    print("=== 测试场景1：基础3维向量 ===")
    # 构造L2归一化的聚类向量
    cluster_vectors = np.array([
        [1/np.sqrt(3), 1/np.sqrt(3), 1/np.sqrt(3)],  # L2范数=1
        [0, 1/np.sqrt(2), 1/np.sqrt(2)],              # L2范数=1
        [1, 0, 0]                                     # L2范数=1
    ])
    # 构造L2归一化的新向量
    new_vector = np.array([0.577, 0.577, 0.577])  # 等价于[1/√3,1/√3,1/√3]，L2范数=1
    #  “归一化向量” 存在计算误差
    # 你写的新向量：[0.577, 0.577, 0.577]
    #new_vector = np.array([0.577, 0.577, 0.577])
    # 实际1/√3 ≈ 0.577350269，你只保留了3位小数，导致L2范数≠1：
    #print(np.linalg.norm(new_vector, ord=2))  # 计算得≈0.999649 ≠ 1

    mean_cosine, mean_linear = calculate_mean_similarity(cluster_vectors, new_vector)
    print(f"cosine均值：{mean_cosine:.6f}")
    print(f"linear均值：{mean_linear:.6f}")
    print(f"结果是否一致：{np.isclose(mean_cosine, mean_linear)}\n")

    # 测试场景2：单向量场景（1个5维归一化向量）
    print("=== 测试场景2：单向量5维 ===")
    cluster_vectors = np.array([[0.2, 0.4, 0.6, 0.4, 0.4]])  # L2范数=√(0.04+0.16+0.36+0.16+0.16)=√1=1
    new_vector = np.array([0.2, 0.4, 0.6, 0.4, 0.4])
    mean_cosine, mean_linear = calculate_mean_similarity(cluster_vectors, new_vector)
    print(f"cosine均值：{mean_cosine:.6f}")
    print(f"linear均值：{mean_linear:.6f}")  # linear_kernel直接算点积：0.2²+0.4²+0.6²+0.4²+0.4² = 0.88，所以均值 = 0.88；
                                            # 这是典型的 "未真正归一化导致结果偏离"。

    print(f"结果是否一致：{np.isclose(mean_cosine, mean_linear)}\n")

    # 测试场景3：高维随机归一化向量（10个128维）
    print("=== 测试场景3：高维随机向量（128维） ===")
    # 随机生成向量并L2归一化
    cluster_vectors = np.random.randn(10, 128)
    cluster_vectors = normalize(cluster_vectors, norm='l2')  # 确保L2归一化
    new_vector = np.random.randn(128)
    new_vector = normalize(new_vector.reshape(1, -1), norm='l2').flatten()
    
    mean_cosine, mean_linear = calculate_mean_similarity(cluster_vectors, new_vector)
    print(f"cosine均值：{mean_cosine:.6f}")
    print(f"linear均值：{mean_linear:.6f}")
    print(f"结果是否一致：{np.isclose(mean_cosine, mean_linear)}\n")

   

if __name__ == "__main__":
    test_similarity_consistency()


