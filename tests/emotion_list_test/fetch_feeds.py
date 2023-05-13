import yaml

cookies = ""
# 读取config.yaml中的qzone.cookies
yaml_str = ""
with open('config.yaml', 'r') as f:
    yaml_str = f.read()

config = yaml.load(yaml_str, Loader=yaml.FullLoader)
cookies = config['qzone']['cookies']

print(cookies)

def generate_gtk(skey):
    """生成gtk"""
    hash_val = 5381
    for i in range(len(skey)):
        hash_val += (hash_val << 5) + ord(skey[i])
    return str(hash_val & 2147483647)

# 把cookies转换成字典
cookies_dict = {}
for cookie in cookies.split(';'):

    cookie = cookie.strip()
    cookie = cookie.split('=')
    if len(cookie)<2:
        continue
    cookies_dict[cookie[0]] = cookie[1]

print(cookies_dict)

gtk = generate_gtk(cookies_dict['p_skey'])

import time
# 获取接口数据
url = "https://user.qzone.qq.com/proxy/domain/ic2.qzone.qq.com/cgi-bin/feeds/feeds3_html_more?uin={}&scope=0&view=1&daylist=&uinlist=&gid=&flag=1&filter=all&applist=all&refresh=0&aisortEndTime=0&aisortOffset=0&getAisort=0&aisortBeginTime=0&pagenum=1&externparam=&firstGetGroup=0&icServerTime=0&mixnocache=0&scene=0&begintime=0&count=10&dayspac=0&sidomain=qzonestyle.gtimg.cn&useutf8=1&outputhtmlfeed=1&rd=0.3436615045444116&usertime={}&windowId=0.9554080071832363&getob=1&g_tk={}".format(cookies_dict['uin'][1:],time.time()*1000, gtk)

import requests

resp = requests.get(
    url,
    cookies=cookies_dict
)

# print(resp.text)

# 提取所有nickname:'xxx',emoji中的xxx

import re

nicknames = re.findall(r"nickname:'(.*?)',emoji", resp.text)

print(nicknames)

# 提取所有,uin:'xxx',ouin中的xxx

uins = re.findall(r",uin:'(.*?)',ouin", resp.text)

print(uins)

# 提取所有html:'xxx',中的xxx

htmls = re.findall(r"html:'(.*?)',", resp.text)

print(htmls)

def rev_escape(s):
    """反转义"""
    # \x3C -> <
    # \/ -> /
    # \x22 -> "
    # &amp; -> &

    s = s.replace('\\x3C', '<')

    s = s.replace('\\/', '/')

    s = s.replace('\\x22', '"')

    s = s.replace('&amp;', '&')

    return s

htmls = [rev_escape(html) for html in htmls]

# 把htmls以json存到htmls.json中
import json

with open('htmls.json', 'w') as f:
    json.dump(htmls, f, ensure_ascii=False, indent=4)

