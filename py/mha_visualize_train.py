# 导入必要的库
import sys
# 移除可能导致冲突的trpc-agent目录
if '/Users/shikaiyuan/code/testone/tpilot/trpc-agent/trpc_agent' in sys.path:
    sys.path.remove('/Users/shikaiyuan/code/testone/tpilot/trpc-agent/trpc_agent')
print(sys.path)
import torch
import torch.nn as nn
import torch.optim as optim
import matplotlib.pyplot as plt

# 定义一个包含MHA层的简单模型类
class MHA_Model(nn.Module):
    def __init__(self):
        super(MHA_Model, self).__init__()
        # 多头注意力层
        self.mha = nn.MultiheadAttention(embed_dim=512, num_heads=8, batch_first=True)
        # 输出层（用于分类任务）
        self.fc = nn.Linear(512, 10)  # 假设是10分类任务
    
    def forward(self, x):
        # MHA前向传播
        attn_output, attn_weights = self.mha(x, x, x, average_attn_weights=False)
        # 取最后一个token的输出用于分类
        last_token_output = attn_output[:, -1, :]
        # 全连接层输出分类结果
        output = self.fc(last_token_output)
        return output, attn_weights

# 训练函数
def train_model():
    # ===================== 1. 初始化模型 =====================
    model = MHA_Model()
    print(f"MHA层的num_heads参数：{model.mha.num_heads}")  # 应输出8
    print(f"MHA层的embed_dim参数：{model.mha.embed_dim}")  # 应输出512
    print(f"是否满足embed_dim % num_heads == 0：{512 % 8 == 0}")  # 应输出True

    # ===================== 2. 定义损失函数和优化器 =====================
    criterion = nn.CrossEntropyLoss()  # 交叉熵损失，适用于分类任务
    optimizer = optim.Adam(model.parameters(), lr=0.001)  # Adam优化器，学习率0.001

    # ===================== 3. 生成模拟训练数据 =====================
    # 生成100个样本，每个样本形状：[1, 5, 512]
    # batch_size=1，seq_len=5（5个token），embed_dim=512
    train_data = torch.randn(100, 1, 5, 512)
    # 生成对应的标签（0-9的随机整数）
    train_labels = torch.randint(0, 10, (100,))

    # ===================== 4. 训练循环 =====================
    num_epochs = 5  # 训练轮数
    for epoch in range(num_epochs):
        model.train()  # 设置模型为训练模式
        running_loss = 0.0
        
        for i in range(len(train_data)):
            x = train_data[i]
            label = train_labels[i]
            
            # 梯度清零
            optimizer.zero_grad()
            
            # 前向传播
            output, _ = model(x)
            
            # 计算损失
            loss = criterion(output, label.unsqueeze(0))  # 添加batch维度
            
            # 反向传播
            loss.backward()
            
            # 更新参数
            optimizer.step()
            
            running_loss += loss.item()
        
        # 打印每轮训练的平均损失
        print(f"Epoch {epoch+1}/{num_epochs}, Loss: {running_loss/len(train_data):.4f}")
    
    print("训练完成！")
    return model

# 可视化函数
def visualize_attention(model):
    # 构造一个测试输入
    x = torch.randn(1, 5, 512)
    
    # 获取模型的注意力权重
    with torch.no_grad():  # 不需要计算梯度
        _, attn_weights = model(x)
    
    print(f"Original attn_weights shape: {attn_weights.shape}")
    
    # 处理注意力权重，准备可视化
    attn_weights = attn_weights.squeeze(0)  # 去除batch_size维度，形状变为 [8,5,5]
    print(f"after attn_weights shape: {attn_weights.shape}")
    
    # 可视化每个head的注意力权重
    fig, axes = plt.subplots(2, 4, figsize=(16, 8))
    axes = axes.flatten()  # 展平为一维数组，方便循环
    
    for i in range(8):
        # 取出第i个head的权重，转numpy
        head_weights = attn_weights[i].numpy()
        
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

def main():
    # 训练模型
    trained_model = train_model()
    
    # 可视化训练后模型的注意力权重
    print("\n--- 训练后模型的注意力权重可视化 ---\n")
    visualize_attention(trained_model)

# 程序入口
if __name__ == "__main__":
    main()