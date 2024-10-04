import requests

# link = 'https://www.imdb.com/chart/top/'
link = 'https://browser-info.ru/'
responce = requests.get(link)
print(responce)

