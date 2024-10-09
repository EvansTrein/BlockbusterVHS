from flask import Blueprint, render_template, redirect, request, flash
from flaskr.config import app, db
from flaskr.db_models import Client, VhsTape, Rental
from flaskr.services.database_validation import validateCreateClient, validateUpdateClient
from .rental import create_rental

route_client = Blueprint('route_client', __name__)

@app.route("/create_client", methods=["GET", "POST"])
def create_client():
    if request.method == "POST":
        data_in = {
            "key_in": "create",
            "name": request.form["name"],
            "age": request.form["age"],
            "phone": request.form["phone"],
        }

        valid_res = validateCreateClient(data_in)
        if valid_res["error"]:
            flash(valid_res["error_text"])
            return render_template("client/create_client_page.html")

        try:
            new_client = Client(
                name=data_in["name"], age=data_in["age"], phone=data_in["phone"]
            )
            db.session.bind = "clients"
            db.session.add(new_client)
            db.session.commit()
            return redirect("/create_client")
        except Exception as err:
            return f"There was an issue adding your client >>>> {str(err)}"

    else:
        return render_template("client/create_client_page.html")


@app.route("/all_clients")
def all_clients():
    total = Client.query.count()
    min_age = request.args.get("min_age", type=int)
    max_age = request.args.get("max_age", type=int)
    data = {"total": total}

    if min_age and max_age:
        all_clients = Client.query.filter(Client.age.between(min_age, max_age)).all()
        count = Client.query.filter(Client.age.between(min_age, max_age)).count()
        data["count"] = count
    elif min_age:
        all_clients = (
            Client.query.filter(Client.age >= min_age).order_by(Client.age.asc()).all()
        )
        count = Client.query.filter(Client.age >= min_age).count()
        data["count"] = count
    elif max_age:
        all_clients = (
            Client.query.filter(Client.age <= max_age).order_by(Client.age.asc()).all()
        )
        count = Client.query.filter(Client.age <= max_age).count()
        data["count"] = count
    else:
        all_clients = Client.query.order_by(Client.name.desc()).all()
        for client in all_clients:
            count_rentals = Rental.query.filter(Rental.client_name == client.name).count()
            client.rental_count = count_rentals

    return render_template("client/all_clients_page.html", all_clients=all_clients, **data)


@app.route("/client/<int:id>", methods=["GET", "POST"])
def client(id):
    client = Client.query.get(id)
    all_clients_films = Rental.query.filter(Rental.client_name == client.name).all()
    count_films = Rental.query.filter(Rental.client_name == client.name).count()

    if request.method == "POST":
        fiml_id = int(request.form["fiml_id"])
        total_rentals = Rental.query.count()

        if not VhsTape.query.filter_by(id_num=fiml_id).first():
            flash("A movie with that ID does not exist")
            return redirect(f"/client/{id}")
        elif VhsTape.query.filter_by(id_num=fiml_id).first().available_quantity == 0:
            flash("Out of stock")
            return redirect(f"/client/{id}")
        else:
            create_rental(id, fiml_id)
            if total_rentals == Rental.query.count():
                flash("This customer has already rented this movie")

        return redirect(f"/client/{id}")
    else:
        return render_template(
            "client/client_page.html",
            client=client,
            all_clients_films=all_clients_films,
            count_films=count_films,
        )


@app.route("/client/<int:id>/update", methods=["GET", "POST"])
def update_client(id):
    client = Client.query.get(id)
    if request.method == "POST":
        data_in = {
            "key_in": "update",
            "obj_bd": client,
            "name": client.name,
            "age": request.form["age"],
            "phone": request.form["phone"],
        }

        valid_res = validateUpdateClient(data_in)
        if valid_res:
            flash(valid_res)
            return redirect(f"/client/{id}/update")

        try:
            client.age = data_in["age"]
            client.phone = data_in["phone"]
            db.session.commit()
            return redirect(f"/client/{id}")
        except Exception as err:
            return f"failed to update:/n {str(err)}"

    else:
        return render_template("client/update_client_page.html", client=client)


@app.route("/client/<int:id>/delete")
def delete_client(id):
    client = Client.query.get_or_404(id)
    existing_rental = Rental.query.filter_by(client_name=client.name).first()

    if existing_rental:
        flash("You can't delete a client who has a movie")
        return redirect("/all_clients")

    try:
        db.session.delete(client)
        db.session.commit()
        return redirect("/all_clients")
    except Exception as err:
        return f"There was a problem deleting that client >>>> {str(err)}"