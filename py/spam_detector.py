import numpy as np
from collections import defaultdict

class NaiveBayesSpamDetector:
    def __init__(self):
        # 存储词汇在垃圾邮件和正常邮件中的出现次数
        self.spam_words = defaultdict(int)
        self.ham_words = defaultdict(int)
        # 垃圾邮件和正常邮件的总数
        self.spam_count = 0
        self.ham_count = 0
        # 所有出现过的词汇
        self.vocabulary = set()
        # 平滑参数，避免概率为0
        self.alpha = 1

    def train(self, emails, labels):
        """
        训练模型
        emails: 邮件内容列表，每个元素是一个分词后的列表
        labels: 标签列表，1表示垃圾邮件，0表示正常邮件
        """
        for email, label in zip(emails, labels):
            if label == 1:
                self.spam_count += 1
                for word in email:
                    self.spam_words[word] += 1
                    self.vocabulary.add(word)
            else:
                self.ham_count += 1
                for word in email:
                    self.ham_words[word] += 1
                    self.vocabulary.add(word)

    def predict(self, email):
        """预测邮件是否为垃圾邮件"""
        # 计算先验概率：P(垃圾邮件)和P(正常邮件)
        total_emails = self.spam_count + self.ham_count
        p_spam = self.spam_count / total_emails
        p_ham = self.ham_count / total_emails

        # 计算P(邮件|垃圾邮件)和P(邮件|正常邮件)
        # 使用对数避免数值下溢
        p_email_spam = 0
        p_email_ham = 0

        # 词汇总数
        spam_word_total = sum(self.spam_words.values())
        ham_word_total = sum(self.ham_words.values())
        vocab_size = len(self.vocabulary)

        for word in email:
            # 计算P(词|垃圾邮件)，使用拉普拉斯平滑
            p_word_spam = (self.spam_words.get(word, 0) + self.alpha) / \
                         (spam_word_total + self.alpha * vocab_size)
            p_email_spam += np.log(p_word_spam)

            # 计算P(词|正常邮件)
            p_word_ham = (self.ham_words.get(word, 0) + self.alpha) / \
                        (ham_word_total + self.alpha * vocab_size)
            p_email_ham += np.log(p_word_ham)

        # 计算后验概率的对数（因为log(a*b) = log(a)+log(b)）
        p_spam_email = np.log(p_spam) + p_email_spam
        p_ham_email = np.log(p_ham) + p_email_ham

        # 返回1表示垃圾邮件，0表示正常邮件
        return 1 if p_spam_email > p_ham_email else 0

# 示例使用
if __name__ == "__main__":
    # 训练数据：邮件内容（已分词）和标签（1=垃圾邮件，0=正常邮件）
    train_emails = [
        ["免费", "中奖", "点击", "链接"],  # 垃圾邮件
        ["优惠", "折扣", "限时", "抢购"],  # 垃圾邮件
        ["账单", "支付", "请", "查收"],  # 正常邮件
        ["会议", "时间", "地点", "通知"],  # 正常邮件
        ["免费", "赠送", "礼品", "领取"],  # 垃圾邮件
        ["项目", "进度", "汇报", "附件"]   # 正常邮件
    ]
    train_labels = [1, 1, 0, 0, 1, 0]

    # 创建并训练模型
    detector = NaiveBayesSpamDetector()
    detector.train(train_emails, train_labels)

    # 测试邮件
    test_emails = [
        ["免费", "活动", "参加", "有礼"],  # 预期垃圾邮件
        ["工作", "安排", "明天", "讨论"],  # 预期正常邮件
        ["中奖", "信息", "点击", "领取"]   # 预期垃圾邮件
    ]

    # 预测并输出结果
    for i, email in enumerate(test_emails):
        result = detector.predict(email)
        print(f"测试邮件 {i+1}: {email}")
        print(f"预测结果: {'垃圾邮件' if result == 1 else '正常邮件'}\n")
