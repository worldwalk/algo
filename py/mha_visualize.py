# 导入必要的库
import sys
# 移除可能导致冲突的trpc-agent目录
if '/Users/shikaiyuan/code/testone/tpilot/trpc-agent/trpc_agent' in sys.path:
    sys.path.remove('/Users/shikaiyuan/code/testone/tpilot/trpc-agent/trpc_agent')
print(sys.path)
import torch
import torch.nn as nn
import matplotlib.pyplot as plt

def main():
    # ===================== 1. 初始化多头注意力（MHA）层 =====================
    # embed_dim=512：输入/输出的特征维度（d_model）
    # num_heads=8：多头数量，需满足 embed_dim % num_heads == 0
    # batch_first=True：输入输出形状为 [batch_size, seq_len, embed_dim]
    mha = nn.MultiheadAttention(embed_dim=512, num_heads=8, batch_first=True)
    print(f"MHA层的num_heads参数：{mha.num_heads}")  # 应输出8
    print(f"MHA层的embed_dim参数：{mha.embed_dim}")  # 应输出512
    print(f"是否满足embed_dim % num_heads == 0：{512 % 8 == 0}")  # 应输出True

    # ===================== 2. 构造模拟输入（词嵌入+位置编码的模拟） =====================
    # batch_size=1，seq_len=5（5个token），embed_dim=512
    x = torch.randn(1, 5, 512)  # 形状：[1, 5, 512]

    # ===================== 3. 执行多头自注意力计算 =====================
    # Q=K=V=x：自注意力（查询、键、值用同一个输入）
    # 返回：attn_output（MHA最终输出）、attn_weights（各head注意力权重）
    attn_output, attn_weights = mha(x, x, x, average_attn_weights=False)
    print(f"Original attn_weights shape: {attn_weights.shape}")

    # ===================== 4. 处理注意力权重，准备可视化 =====================
    # 去除batch_size维度（batch_size=1无意义），形状变为 [8,5,5]
    attn_weights = attn_weights.squeeze(0)  # [num_heads, seq_len, seq_len]
    print(f"after attn_weights shape: {attn_weights.shape}")

    # ===================== 5. 可视化每个head的注意力权重 =====================
    # 创建2行4列子图（8个head对应8个子图），画布大小16×8
    fig, axes = plt.subplots(2, 4, figsize=(16, 8))
    axes = axes.flatten()  # 展平为一维数组，方便循环

    # 遍历每个head，绘制热力图
    for i in range(8):
        # 取出第i个head的权重，分离计算图并转numpy（matplotlib仅支持numpy）
        head_weights = attn_weights[i].detach().numpy()
        
        # 绘制热力图：颜色越深，注意力权重越大
        im = axes[i].imshow(head_weights, cmap='Blues')
        
        # 设置子图标题
        axes[i].set_title(f'Head {i+1}')
        
        # 添加颜色条（显示权重数值对应关系）
        plt.colorbar(im, ax=axes[i])

    # 自动调整子图间距，避免重叠
    plt.tight_layout()
    # 显示画布
    plt.show()

# 程序入口
if __name__ == "__main__":
    main()