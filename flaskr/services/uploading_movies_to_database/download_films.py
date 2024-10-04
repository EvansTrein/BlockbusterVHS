import requests
from bs4 import BeautifulSoup

header = {'user-agent': 'magic', 'Accept-Language': 'en-US'}

link = 'https://www.imdb.com/chart/top/'

def download():
    answer = {"error": False, "error_text": "", "films": None}
    try:
        responce = requests.get(link, headers=header).text
        soup = BeautifulSoup(responce, 'lxml')
        block = soup.find('ul', class_='ipc-metadata-list ipc-metadata-list--dividers-between sc-a1e81754-0 dHaCOW compact-list-view ipc-metadata-list--base')
        all_top_250 = block.find_all('li', class_='ipc-metadata-list-summary-item sc-10233bc-0 TwzGn cli-parent')

        all_films = []
        for one_element in all_top_250:
            film = {'title': '', 'year': '', 'age_rating': '', 'count': '20'}
            film['title'] = ''.join(x for x in one_element.find('h3', class_='ipc-title__text').text if x.isalpha() or x == ' ')
            data = one_element.find('div', class_='sc-ab348ad5-7 cqgETV cli-title-metadata')
            film['year'] = data.find_all('span')[0].text
            film['age_rating'] = data.find_all('span')[2].text
            all_films.append(film)
        
        answer['films'] = all_films
        return answer
    except Exception as er:
        answer['error'] = True
        answer['error_text'] = str(er)
        return answer


# download()
