import requests

# link = 'https://www.imdb.com/chart/top/'
link = 'https://vk.com/'
responce = requests.get(link)
print(responce)

