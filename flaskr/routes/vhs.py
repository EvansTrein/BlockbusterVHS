from flask import Blueprint, render_template, redirect, request, flash
from flaskr.config import app, db
from flaskr.db_models import VhsTape, Rental
from flaskr.services.database_validation import validateCreateVhs, validateUpdateVhs

route_vhs = Blueprint('route_vhs', __name__)

@app.route("/create_vhstape", methods=["GET", "POST"])
def create_vhstape(*args):
    if request.method == "POST":
        if len(args) == 0:
            data_in = {
                "key_in": "create",
                "title": request.form["title"],
                "year": request.form["year"],
                "age_rating": request.form["age_rating"],
                "count": request.form["count"]
            }
        else:
            data_in = {
                "key_in": "create",
                "title": args[0]["title"],
                "year": args[0]["year"],
                "age_rating": args[0]["age_rating"],
                "count": args[0]["count"]
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
        all_vhstapes = VhsTape.query.order_by(VhsTape.year.desc()).all()

    return render_template("vhs/all_vhstapes_page.html", all_vhstapes=all_vhstapes, **data)


@app.route("/vhs/<int:id>")
def vhs(id):
    vhs = VhsTape.query.get(id)
    return render_template("vhs/vhs_page.html", vhs=vhs)


@app.route("/vhs/<int:id>/update", methods=["GET", "POST"])
def update_vhstape(id):
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