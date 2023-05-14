

import json
def retrieve_info_from_html(htmls: str):
    """从html中提取信息"""
    # 提取发表人uin和nick
    # nameCard_12345678\">RockChinQ</a> 中的12345678和RockChinQ
    import re
    
    uin_nick = re.findall(r"nameCard_(\d+)\">(.*?)</a>", htmls)

    print(uin_nick)

    # 提取文字内容
    # class=\"f-info\">This is content here!</div><div class=\"qz_summary wupfeed\" 中的This is content here!

    contents = re.findall(r"class=\"f-info\">(.*?)</div><div class=\"qz_summary wupfeed\"", htmls)

    print(contents)

    # 提取绝对时间戳
    # data-abstime=\"1683901591\" 中的1683901591

    abstimes = re.findall(r"data-abstime=\"(\d+)\"", htmls)

    print(abstimes)

    # 提取图片
    # src=\"http://a1.qpic.cnxxx style=中的http://a1.qpic.cnxxx
    
    pattern = re.compile(r"src=\"(http://a1.*?)\"")

    images = pattern.findall(htmls)

    # 把/替换成/
    images = [image.replace('/', '/') for image in images]

    # 把&amp;替换成&
    images = [image.replace('&amp;', '&') for image in images]

    # 提取每个image中的&ek=1前的字符串
    images = [image.split('&ek=1')[0] for image in images]

    print(json.dumps(images, indent=4, ensure_ascii=False))

    # 提取tid
    # data-tid=\"abcd1234\" 中的abcd1234
    # 并且长度为24

    tids = re.findall(r"data-tid=\"([a-z\d]{24})\"", htmls)

    print(tids)


# 读取htmls.json
htmls = None
with open('htmls.json', 'r', encoding='utf-8') as f:
    htmls = json.load(f)

retrieve_info_from_html(htmls[1])
