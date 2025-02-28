from flask import Blueprint, render_template, redirect, request, flash
from flaskr.config import app, db
from flaskr.db_models import Client, VhsTape, Rental
from flaskr.services.database_validation import (
    validateCreateClient,
    validateUpdateClient,
)
from .rental import create_rental

# connect routers from this file and name them as follows
route_client = Blueprint("route_client", __name__)


@app.route("/create_client", methods=["GET", "POST"])
def create_client():
    """
    Creates a client, given a name, age and phone number.

    If the request method is POST, validates the input data, checks if the client exists and if the data is correct, creates a new client and commits the changes to the database.

    If the request method is GET, renders a template to input the client data.

    :return: Redirects to the same page if the request method is POST, renders a template if the request method is GET.
    """
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
    """
    Displays all clients with optional filters by age range.

    If min_age and/or max_age are provided as query parameters, it filters the clients by that criteria and displays the total count of filtered clients.
    Otherwise, all clients are displayed.

    :return: rendered template with all clients and total count
    """
    total = Client.query.count()
    name = request.args.get("name", type=str)
    min_age = request.args.get("min_age", type=int)
    max_age = request.args.get("max_age", type=int)
    data = {"total": total}

    if name:
        all_clients = Client.query.filter(Client.name.ilike(f'%{name}%')).all()
        count = Client.query.filter(Client.name.ilike(f'%{name}%')).count()
        data["count"] = count
    elif min_age and max_age:
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
        all_clients = Client.query.all()
        for client in all_clients:
            count_rentals = Rental.query.filter(
                Rental.client_name == client.name
            ).count()
            client.rental_count = count_rentals

        all_clients.sort(key=lambda x: x.rental_count, reverse=True)

    return render_template(
        "client/all_clients_page.html", all_clients=all_clients, **data
    )


@app.route("/client/<int:id>", methods=["GET", "POST"])
def client(id):
    """
    Displays a client with its details and list of rented movies.

    If the request is a POST, it creates a rental for the client with the provided movie ID.
    It checks if the movie exists and if it is in stock.
    If the movie is already rented by the client, it does not create a new rental.

    :param id: the client id
    :return: rendered template with the client and its rented movies
    """
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
    """
    Updates a client with given id, given a name, age and phone number.

    If the request method is POST, validates the input data, checks if the client exists and if the data is correct, updates the client and commits the changes to the database.

    If the request method is GET, renders a template to input the client data.

    :param id: the client id
    :return: Redirects to the same page if the request method is POST, renders a template if the request method is GET.
    """
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
    """
    Deletes a client with given id.

    If the request method is POST, renders a template to confirm the deletion.

    If the request method is GET, checks if the client has any rentals, if not, deletes the client and commits the changes to the database.

    :param id: the client id
    :return: Redirects to the page with all clients if the request method is GET, renders a template if the request method is POST.
    """
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
