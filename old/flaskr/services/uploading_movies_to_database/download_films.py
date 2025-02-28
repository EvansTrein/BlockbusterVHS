import requests
import json
from bs4 import BeautifulSoup


header = {
    "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
    "Accept-Language": "en-US",
}

link = "https://www.imdb.com/chart/top/"


def download(num):
    """
    Downloads films from IMDB top chart and returns them in a list

    On the IMDB site this 250 movies lie inside the <script> tag, in json format, the data is taken from there, it's not parsing the site fully.

    The function takes one argument, the number of films to download.

    Returns a dictionary with two keys: "error" and "error_text". If the download is successful,
    "error" is False and "error_text" is an empty string. If the download fails, "error" is True and
    "error_text" is the error message.

    The downloaded films are stored in a list and returned as the value of the "films" key in the
    returned dictionary. Each film is a dictionary with the following keys: "title", "year",
    "age_rating", "poster_image", and "count".

    :param num: The number of films to download
    :type num: int
    :return: A dictionary with the downloaded films
    :rtype: dict
    """
    all_films = []
    answer = {"error": False, "error_text": "", "films": None}
    try:
        responce = requests.get(link, headers=header).text
        soup = BeautifulSoup(responce, "lxml")
        data = soup.find("script", id="__NEXT_DATA__")
        data_json = json.loads(data.text)

        for el in data_json["props"]["pageProps"]["pageData"]["chartTitles"]["edges"]:
            film = {
                "title": "",
                "year": "",
                "age_rating": "",
                "poster_image": "",
                "count": num,
            }

            film["title"] = el["node"]["titleText"]["text"]
            film["year"] = str(el["node"]["releaseYear"]["year"])
            rating = el["node"]["certificate"]
            film["age_rating"] = (
                el["node"]["certificate"]["rating"] if rating else "Not Rated"
            )
            film["poster_image"] = el["node"]["primaryImage"]["url"]

            all_films.append(film)

        answer["films"] = all_films
        return answer
    except Exception as er:
        answer["error"] = True
        answer["error_text"] = str(er)
        return answer
