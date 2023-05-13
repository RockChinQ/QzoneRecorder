import json

# 读取htmls.json
htmls = []

with open('htmls.json', 'r', encoding='utf-8') as f:
    htmls = json.load(f)

print(htmls[3])