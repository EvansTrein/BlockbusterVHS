import requests
import json
from bs4 import BeautifulSoup


header = {
    "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
    "Accept-Language": "en-US",
}

link = "https://www.imdb.com/chart/top/"


def download(num):
    all_films = []
    answer = {"error": False, "error_text": "", "films": None}
    try:
        responce = requests.get(link, headers=header).text
        soup = BeautifulSoup(responce, "lxml")
        data = soup.find("script", id="__NEXT_DATA__")
        data_json = json.loads(data.text)

        for el in data_json["props"]["pageProps"]["pageData"]["chartTitles"]["edges"]:
            film = {"title": "", "year": "", "age_rating": "", "count": num}

            film["title"] = el["node"]["titleText"]["text"]
            film["year"] = str(el["node"]["releaseYear"]["year"])
            rating = el["node"]["certificate"]
            film["age_rating"] = (
                el["node"]["certificate"]["rating"] if rating else "Not Rated"
            )

            all_films.append(film)

        answer["films"] = all_films
        return answer
    except Exception as er:
        answer["error"] = True
        answer["error_text"] = str(er)
        return answer
