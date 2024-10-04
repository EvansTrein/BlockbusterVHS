from flask import Blueprint, render_template, jsonify, flash
from flaskr.config import app, db
from flaskr.db_models import Client, VhsTape, Rental
from flaskr.services.uploading_movies_to_database import download_films
from flaskr.routes.vhs import create_vhstape

route_home_and_others = Blueprint('route_home_and_others', __name__)

@app.route("/")
def home():
    data = {
        "total_rentals": Rental.query.count(),
        "total_clients": Client.query.count(),
        "total_films": VhsTape.query.count(),
        "total_movies_in_stock": db.session.query(db.func.sum(VhsTape.available_quantity)).scalar(),
    }

    return render_template("home_page.html", **data)


@app.route("/clear_database", methods=["POST"])
def clear_database():
    db.session.query(VhsTape).delete()
    db.session.query(Client).delete()
    db.session.query(Rental).delete()
    db.session.commit()

    return jsonify({'redirect': '/'})


@app.route("/download_films", methods=["POST"])
def download():
    res = download_films.download()
    if not res['error']:
        for el in res['films']:
            create_vhstape(el)

        return jsonify({'redirect': '/'})
    
    flash(res['error_text'])
    return jsonify({'redirect': '/'})