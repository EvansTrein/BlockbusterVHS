from flask import Blueprint, render_template, redirect, request, flash
from flaskr.config import app, db
from flaskr.db_models import VhsTape, Rental
from flaskr.services.database_validation import validateCreateVhs, validateUpdateVhs

# connect routers from this file and name them as follows
route_vhs = Blueprint("route_vhs", __name__)


@app.route("/create_vhstape", methods=["GET", "POST"])
def create_vhstape(*args):
    """
    Route to create a new VHS tape in the database.

    This route is used to add a new VHS tape to the database.
    The form data is validated using the validateCreateVhs function.

    GET parameters:
        None

    POST parameters:
        key_in: str, "create" or "create_with_film_data"
        title: str, title of the VHS tape
        year: int, year of the VHS tape
        age_rating: str, age rating of the VHS tape
        count: int, count of the VHS tape
        poster_image: str, link to the poster image of the VHS tape

    Returns a rendered template with a form for adding a new VHS tape on GET.
    On POST, validates the data and adds the new VHS tape to the database if valid.
    If invalid, flashes an error message and renders the form again.
    """
    if request.method == "POST":
        if len(args) == 0:
            data_in = {
                "key_in": "create",
                "title": request.form["title"],
                "year": request.form["year"],
                "age_rating": request.form["age_rating"],
                "count": request.form["count"],
                "poster_image": "../static/images/stock_img_film.jpg",
            }
        else:
            data_in = {
                "key_in": "create",
                "title": args[0]["title"],
                "year": args[0]["year"],
                "age_rating": args[0]["age_rating"],
                "count": args[0]["count"],
                "poster_image": args[0]["poster_image"],
            }

        valid_res = validateCreateVhs(data_in)
        if valid_res["error"]:
            flash(valid_res["error_text"])
            return render_template("vhs/create_vhstape_page.html")

        try:
            vhs = VhsTape(
                title=data_in["title"],
                year=data_in["year"],
                age_rating=data_in["age_rating"],
                count=data_in["count"],
                poster_image=data_in["poster_image"],
            )
            db.session.add(vhs)
            db.session.commit()
            return redirect("/create_vhstape")
        except Exception as err:
            return f"There was an issue adding your VHS tape >>>> {str(err)}"

    else:
        return render_template("vhs/create_vhstape_page.html")


@app.route("/all_vhstapes")
def all_vhstapes():
    """
    Route to show all VHS tapes in the database, with optional filters
    by film title, year of release, or age rating.

    GET parameters:
        id_num: int, filter by film id
        film_title: str, filter by film title
        year_of_release: int, filter by year of release
        age_rating: str, filter by age rating

    Returns a rendered template with a list of all VHS tapes, and
    the total count of VHS tapes.
    """
    total = VhsTape.query.count()
    film_id = request.args.get("id_num", type=int)
    film_title = request.args.get("film_title")
    year_of_release = request.args.get("year_of_release", type=int)
    age_rating = request.args.get("age_rating")
    data = {"total": total}

    if film_id:
        all_vhstapes = VhsTape.query.filter(VhsTape.id_num == film_id).all()
    elif film_title:
        all_vhstapes = VhsTape.query.filter(VhsTape.title.like(f"%{film_title}%")).all()
        count = VhsTape.query.filter(VhsTape.title.like(f"%{film_title}%")).count()
        data["count"] = count
    elif year_of_release:
        all_vhstapes = VhsTape.query.filter(VhsTape.year == year_of_release).all()
        count = VhsTape.query.filter(VhsTape.year == year_of_release).count()
        data["count"] = count
    elif age_rating:
        all_vhstapes = VhsTape.query.filter(VhsTape.age_rating == age_rating).all()
        count = VhsTape.query.filter(VhsTape.age_rating == age_rating).count()
        data["count"] = count
    else:
        all_vhstapes = VhsTape.query.order_by(VhsTape.id_num.asc()).all()

    return render_template(
        "vhs/all_vhstapes_page.html", all_vhstapes=all_vhstapes, **data
    )


@app.route("/vhs/<int:id>")
def vhs(id):
    """
    Shows a single VHS tape by given id.

    Parameters:
        id (int): The id of the VHS tape to show.

    Returns:
        A rendered template with information about the VHS tape.
    """
    vhs = VhsTape.query.get(id)
    return render_template("vhs/vhs_page.html", vhs=vhs)


@app.route("/vhs/<int:id>/update", methods=["GET", "POST"])
def update_vhstape(id):
    """
    Route to update a VHS tape in the database.

    Parameters:
        id (int): The id of the VHS tape to update.

    GET parameters:
        None

    POST parameters:
        year: int, the new year of the VHS tape
        age_rating: str, the new age rating of the VHS tape
        count: int, the new count of the VHS tape

    Returns a rendered template with a form for updating a VHS tape on GET.
    On POST, validates the data and updates the VHS tape in the database if valid.
    If invalid, flashes an error message and renders the form again.
    """
    vhs = VhsTape.query.get(id)
    if request.method == "POST":
        data_in = {
            "key_in": "update",
            "obj_bd": vhs,
            "title": vhs.title,
            "year": request.form["year"],
            "age_rating": request.form["age_rating"],
            "count": request.form["count"],
        }

        valid_res = validateUpdateVhs(data_in)
        if valid_res:
            flash(valid_res)
            return redirect(f"/vhs/{id}/update")

        try:
            vhs.year = data_in["year"]
            vhs.age_rating = data_in["age_rating"]
            vhs.count = data_in["count"]
            vhs.available_quantity = int(data_in["count"]) - vhs.issued_to_clients
            db.session.commit()
            return redirect(f"/vhs/{id}")
        except Exception as err:
            return f"failed to update:/n {str(err)}"

    else:
        return render_template("vhs/update_vhstape_page.html", vhs=vhs)


@app.route("/vhs/<int:id>/delete")
def delete_vhstape(id):
    """
    Deletes a VHS tape given its ID.

    Checks if there is an existing rental for the VHS tape being deleted.
    If there is, flashes an error message and redirects to the all VHS tapes page.
    If there isn't, deletes the VHS tape from the database and redirects to the all VHS tapes page.
    If there is a problem deleting the VHS tape, flashes an error message with the exception message.
    """
    vhs = VhsTape.query.get_or_404(id)
    existing_rental = Rental.query.filter_by(title_vhs=vhs.title).first()

    if existing_rental:
        flash("You can't delete a movie that's been given away")
        return redirect("/all_vhstapes")

    try:
        db.session.delete(vhs)
        db.session.commit()
        return redirect("/all_vhstapes")
    except Exception as err:
        return f"There was a problem deleting that VHS tape >>>> {str(err)}"
