from flask import Blueprint, render_template, redirect, request, flash
from flaskr.config import app, db
from flaskr.db_models import VhsTape, Rental
from flaskr.services.database_validation import validateCreateRental

# connect routers from this file and name them as follows
route_rental = Blueprint('route_rental', __name__)

@app.route("/create_rental", methods=["GET", "POST"])
def create_rental(*args):
    """
    Creates a rental, given a client ID and a VHS tape ID.

    If the request method is POST, validates the input data, checks if the client and the VHS tape exist and if the VHS tape is available, and if everything is correct, creates a new rental and updates the VHS tape quantities.

    If the request method is GET, renders a template to input the client ID and the VHS tape ID.

    :param args: A tuple of two elements, the first is the client ID and the second is the VHS tape ID. If no arguments are given, they are taken from the request form.
    :return: Redirects to the same page if the request method is POST, renders a template if the request method is GET.
    """

    if request.method == "POST":
        data_in = {
            "client_id": request.form["client_id"] if len(args) == 0 else args[0],
            "vhs_tape_id": request.form["vhs_tape_id"] if len(args) == 0 else args[1],
        }

        # submit data for verification
        valid_res = validateCreateRental(data_in)
        if valid_res["error"]:
            flash(valid_res["error_text"])
            return render_template("rental/create_rental_page.html")

        try:
            vhs = VhsTape.query.get(data_in["vhs_tape_id"])
            rental = Rental(
                client_id=data_in["client_id"], vhs_tape_id=data_in["vhs_tape_id"]
            )
            db.session.bind = "rentals"
            db.session.add(rental)
            vhs.issued_to_clients += 1
            vhs.available_quantity -= 1
            db.session.commit()
            return redirect("/create_rental")
        except Exception as err:
            return f"There was a problem with the cassette issue >>>> {str(err)}"

    else:
        return render_template("rental/create_rental_page.html")


@app.route("/all_rentals")
def all_rentals():
    """
    Displays all rentals with optional filters by client name and film title.

    If client_name or film_title are provided as query parameters, it filters the rentals by that criteria and displays the total count of filtered rentals.
    Otherwise, all rentals are displayed.

    :return: rendered template with all rentals and total count
    """

    total = Rental.query.count()
    client_name = request.args.get("client_name")
    film_title = request.args.get("film_title")
    data = {"total": total}

    if client_name:
        all_issued = Rental.query.filter(
            Rental.client_name.like(f"%{client_name}%")
        ).all()
        count = Rental.query.filter(Rental.client_name == client_name).count()
        data["count"] = count
    elif film_title:
        all_issued = Rental.query.filter(Rental.title_vhs.like(f"%{film_title}%")).all()
        count = Rental.query.filter(Rental.title_vhs == film_title).count()
        data["count"] = count
    else:
        all_issued = Rental.query.all()

    return render_template("rental/all_rentals_page.html", all_issued=all_issued, **data)



@app.route("/rental/<int:id>/delete")
def delete_rental(id):
    """
    Deletes a rental with given id, returns user to the previous page if that was a client's page, otherwise returns to the page with all rentals.
    :param id: id of the rental to delete
    """

    rental = Rental.query.get_or_404(id)
    referrer_url = request.referrer

    try:
        vhs = VhsTape.query.get(rental.vhs_tape_id)
        vhs.issued_to_clients -= 1
        vhs.available_quantity += 1
        db.session.delete(rental)
        db.session.commit()
        if "client" in referrer_url:
            return redirect(referrer_url)
        return redirect("/all_rentals")
    except Exception as err:
        return f"There was a problem with deleting this lease >>>> {str(err)}"