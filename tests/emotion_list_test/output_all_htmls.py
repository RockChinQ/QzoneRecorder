# 创建文件夹html_dumps
import os
import json

if not os.path.exists('html_dumps'):

    os.mkdir('html_dumps')


htmls = []

with open('htmls.json', 'r', encoding='utf-8') as f:

    htmls = json.load(f)

# 在每个前面加上<meta charset="utf-8" />,输出到html_dumps中
for i, html in enumerate(htmls):
    with open('html_dumps/{}.html'.format(i), 'w', encoding='utf-8') as f:

        f.write('<meta charset="utf-8" />\n')

        f.write(html)