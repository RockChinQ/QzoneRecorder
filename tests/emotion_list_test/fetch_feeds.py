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
pagenum = 10
# 获取接口数据
begintime = 1683952141
# url = "https://user.qzone.qq.com/proxy/domain/ic2.qzone.qq.com/cgi-bin/feeds/feeds3_html_more?uin={}&scope=0&view=1&daylist=&uinlist=&gid=&flag=1&filter=all&applist=all&refresh=0&aisortEndTime=0&aisortOffset=0&getAisort=0&aisortBeginTime=0&pagenum={}&externparam=basetime=1684034841&pagenum={}&dayvalue=0&getadvlast=1&hasgetadv=10412150919^0^1684038055*10376731439^0^1684038055&lastentertime=0&LastAdvPos=1&UnReadCount=0&UnReadSum=8&LastIsADV=1&UpdatedFollowUins=&UpdatedFollowCount=0&LastRecomBrandID=&TRKPreciList=&firstGetGroup=0&icServerTime=0&mixnocache=0&scene=0&begintime=1684034841&count=10&dayspac=0&sidomain=qzonestyle.gtimg.cn&useutf8=1&outputhtmlfeed=1&rd=0.7917492881867267&usertime={}&windowId=0.6394508482669143&g_tk={}&g_tk={}".format(cookies_dict['uin'][1:],pagenum,pagenum,time.time()*1000, gtk,gtk)
url = "https://user.qzone.qq.com/proxy/domain/ic2.qzone.qq.com/cgi-bin/feeds/feeds3_html_more?uin={}&scope=0&view=1&daylist=&uinlist=&gid=&flag=1&filter=all&applist=all&refresh=0&aisortEndTime=0&aisortOffset=0&getAisort=0&aisortBeginTime=0&pagenum={}&externparam=basetime={}&pagenum={}&dayvalue=0&getadvlast=1&hasgetadv=10408302957^0^1684038110&lastentertime=0&LastAdvPos=1&UnReadCount=0&UnReadSum=18&LastIsADV=1&UpdatedFollowUins=&UpdatedFollowCount=0&LastRecomBrandID=&TRKPreciList=&firstGetGroup=0&icServerTime=0&mixnocache=0&scene=0&begintime={}&count=10&dayspac=0&sidomain=qzonestyle.gtimg.cn&useutf8=1&outputhtmlfeed=1&rd=0.9141480747320878&usertime={}&windowId=0.8010646234564336&g_tk={}&g_tk={}".format(cookies_dict['uin'][1:],pagenum,begintime, pagenum,begintime, time.time()*1000, gtk, gtk)
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

with open('htmls.json', 'w', encoding="utf-8") as f:
    json.dump(htmls, f, ensure_ascii=False, indent=4)

print("count:", len(htmls))

print(url)
