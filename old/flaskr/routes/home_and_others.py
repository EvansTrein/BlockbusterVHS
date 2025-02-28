from flask import (
    Blueprint,
    render_template,
    jsonify,
    flash,
    request,
    send_from_directory,
)
from flaskr.config import app, db
from flaskr.db_models import Client, VhsTape, Rental
from flaskr.services.uploading_movies_to_database import download_films
from flaskr.routes.vhs import create_vhstape

# connect routers from this file and name them as follows
route_home_and_others = Blueprint("route_home_and_others", __name__)


@app.route("/")
def home():
    """
    Main page of the application. Show the total number of clients, films, rental records and the number of films in stock.
    :return: rendered template
    """
    data = {
        "total_rentals": Rental.query.count(),
        "total_clients": Client.query.count(),
        "total_films": VhsTape.query.count(),
        "total_movies_in_stock": db.session.query(
            db.func.sum(VhsTape.available_quantity)
        ).scalar(),
    }

    return render_template("home_page.html", **data)


@app.route("/clear_database", methods=["POST"])
def clear_database():
    """
    Deletes all records from all tables in the database (VhsTape, Client, Rental).

    :return: JSON with a redirect to the main page
    """
    db.session.query(VhsTape).delete()
    db.session.query(Client).delete()
    db.session.query(Rental).delete()
    db.session.commit()

    return jsonify({"redirect": "/"})


@app.route("/download_films", methods=["POST"])
def download():
    """
    Download films from the site IMDb.com and add them to the database
    :return: JSON with a redirect to the main page
    """
    number = int(request.get_data())
    res = download_films.download(number)
    if not res["error"]:
        for el in res["films"]:
            create_vhstape(el)

        return jsonify({"redirect": "/"})

    flash(res["error_text"])
    return jsonify({"redirect": "/"})


@app.route("/static/<path:filename>")
def serve_static(filename):
    """
    This function is necessary to serve static files from the static folder. It is not used explicitly,
    but Flask uses it to serve static files when the URL starts with '/static'.
    :param filename: The name of the file to be served
    :return: The file from the static folder
    """
    return send_from_directory(app.static_folder, filename)
