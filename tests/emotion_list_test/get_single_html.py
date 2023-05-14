import json

# 读取htmls.json
htmls = []

with open('htmls.json', 'r', encoding='utf-8') as f:
    htmls = json.load(f)

index = 2
print(htmls[index])

# 写入到single_html.html

with open('single_html.html', 'w', encoding='utf-8') as f:
    f.write(htmls[index])

