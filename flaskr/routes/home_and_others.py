from flask import Blueprint, render_template
from flaskr.config import app, db
from flaskr.db_models import Client, VhsTape, Rental

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

    return "ALL databases have been cleaned"