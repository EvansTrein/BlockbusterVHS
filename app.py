from flask import render_template, request, redirect, flash
from config import app, db
from models import VhsTape, Client, Rental
from re import fullmatch


with app.app_context():
    db.create_all()


@app.route("/")
def home():
    return render_template("home_page.html")


@app.route("/clear_database", methods=["POST"])
def clear_database():
    db.session.query(VhsTape).delete()
    db.session.query(Client).delete()
    db.session.query(Rental).delete()
    db.session.commit()

    return "ALL databases have been cleaned"


@app.route("/vhs/<int:id>")
def vhs(id):
    vhs = VhsTape.query.get(id)
    return render_template("vhs_page.html", vhs=vhs)


@app.route("/all_vhstapes")
def all_vhstapes():
    all_vhstapes = VhsTape.query.order_by(VhsTape.year.desc()).all()
    return render_template("all_vhstapes_page.html", all_vhstapes=all_vhstapes)


@app.route("/create_vhstape", methods=["GET", "POST"])
def create_vhstape():
    if request.method == "POST":
        title = request.form["title"]
        year = request.form["year"]
        age_rating = request.form["age_rating"]
        count = request.form["count"]

        # checking the conditions for a database record
        if VhsTape.query.filter_by(title=title).first():
            flash("There's already a tape like this")
            return render_template("create_vhstape_page.html")
        elif not all((title, year, age_rating, count)):
            flash("All fields must be filled in")
            return render_template("create_vhstape_page.html")
        elif len(title) > 100 or len(age_rating) > 10:
            flash(
                'The number of characters in the "name film" or "age rating" field is exceeded'
            )
            return render_template("create_vhstape_page.html")
        else:
            try:
                vhs = VhsTape(
                    title=title, year=year, age_rating=age_rating, count=count
                )
                db.session.add(vhs)
                db.session.commit()
                return redirect("/create_vhstape")
            except Exception as err:
                return f"There was an issue adding your VHS tape >>>> {str(err)}"

    else:
        return render_template("create_vhstape_page.html")


@app.route("/vhs/<int:id>/update", methods=["GET", "POST"])
def update_vhstape(id):
    vhs = VhsTape.query.get(id)
    if request.method == "POST":
        title = request.form["title"]
        year = request.form["year"]
        age_rating = request.form["age_rating"]
        count = int(request.form["count"])

        if count < vhs.issued_to_clients:
            flash("The total quantity may not be less than the sum of available and issued")
            return redirect(f"/vhs/{id}/update")

        if all((title, year, age_rating, count)):
            vhs.title = title
            vhs.year = year
            vhs.age_rating = age_rating
            vhs.count = count
            vhs.available_quantity = count - vhs.issued_to_clients
        else:
            flash("Fields cannot be empty")
            return redirect(f"/vhs/{id}/update")

        try:
            db.session.commit()
            return redirect(f"/vhs/{id}")
        except Exception as err:
            return f"failed to update:/n {str(err)}"

    else:
        return render_template("update_vhstape_page.html", vhs=vhs)


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



# ===========================================================================================================


@app.route("/create_client", methods=["GET", "POST"])
def create_client():
    if request.method == "POST":
        name = request.form["name"]
        age = request.form["age"]
        phone = request.form["phone"]

        # checking the conditions for a database record
        if Client.query.filter_by(name=name).first():
            flash("There's already a tape like this")
            return render_template("create_client_page.html")
        elif not all((name, age, phone)):
            flash("All fields must be filled in")
            return render_template("create_client_page.html")
        elif int(age) < 14:
            flash("The client cannot be under 14 years of age")
            return render_template("create_client_page.html")
        elif not fullmatch(r"\+\d{10,20}", phone):
            flash(
                'Incorrect phone number, the number starts with "+country code" and has 10 to 20 digits'
            )
            return render_template("create_client_page.html")
        else:
            try:
                new_client = Client(name=name, age=age, phone=phone)
                db.session.bind = "clients"
                db.session.add(new_client)
                db.session.commit()
                return redirect("/create_client")
            except Exception as err:
                return f"There was an issue adding your client >>>> {str(err)}"

    else:
        return render_template("create_client_page.html")


@app.route("/all_clients")
def all_clients():
    all_clients = Client.query.order_by(Client.name.desc()).all()
    return render_template("all_clients_page.html", all_clients=all_clients)


@app.route("/client/<int:id>")
def client(id):
    client = Client.query.get(id)
    return render_template("client_page.html", client=client)


@app.route("/client/<int:id>/update", methods=["GET", "POST"])
def update_client(id):
    client = Client.query.get(id)
    if request.method == "POST":
        name = request.form["name"]
        age = request.form["age"]
        phone = request.form["phone"]

        if all((name, age, phone)):
            client.name = name
            client.age = age
            client.phone = phone
        else:
            flash("fields cannot be empty")
            return redirect(f"/client/{id}/update")

        try:
            db.session.commit()
            return redirect(f"/client/{id}")
        except Exception as err:
            return f"failed to update:/n {str(err)}"

    else:
        return render_template("update_client_page.html", client=client)


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


# ===========================================================================================================


@app.route("/all_rentals")
def all_rentals():
    total = Rental.query.count()
    client_name = request.args.get("client_name")
    film_title = request.args.get("film_title")
    data = {'total': total}

    if client_name:
        all_issued = Rental.query.filter(Rental.client_name.like(f"%{client_name}%")).all()
        count = Rental.query.filter(Rental.client_name == client_name).count()
        data['count'] = count
    elif film_title:
        all_issued = Rental.query.filter(Rental.title_vhs.like(f"%{film_title}%")).all()
        count = Rental.query.filter(Rental.title_vhs == film_title).count()
        data['count'] = count
    else:
        all_issued = Rental.query.all()

    return render_template("all_rentals_page.html", all_issued=all_issued, **data)


@app.route("/create_rental", methods=["GET", "POST"])
def create_rental():
    if request.method == "POST":
        client_id = request.form["client_id"]
        vhs_tape_id = request.form["vhs_tape_id"]
        client = Client.query.get(client_id)
        vhs = VhsTape.query.get(vhs_tape_id)
        existing_rental = Rental.query.filter_by(title_vhs=vhs.title, client_name=client.name).first()

        if (client is None) and (vhs is None):
            flash("The movie and client with those ID's do not exist")
            return render_template("create_rental_page.html")
        elif not client:
            flash("Client with this ID does not exist")
            return render_template("create_rental_page.html")
        elif not vhs:
            flash("A movie with that ID does not exist")
            return render_template("create_rental_page.html")
        elif existing_rental:
            flash("This customer has already rented this movie")
            return render_template("create_rental_page.html")
        elif vhs.available_quantity == 0:
            flash("Out of stock")
            return render_template("create_rental_page.html")
        else:
            try:
                rental = Rental(client_id=client_id, vhs_tape_id=vhs_tape_id)
                db.session.bind = "rentals"
                db.session.add(rental)
                vhs.issued_to_clients += 1
                vhs.available_quantity -= 1
                db.session.commit()
                return redirect("/create_rental")
            except Exception as err:
                return f"There was a problem with the cassette issue >>>> {str(err)}"
    
    else:
        return render_template("create_rental_page.html")
    

@app.route("/rental/<int:id>/delete")
def delete_rental(id):
    rental = Rental.query.get_or_404(id)

    try:
        vhs = VhsTape.query.get(rental.vhs_tape_id)
        vhs.issued_to_clients -= 1
        vhs.available_quantity += 1
        db.session.delete(rental)
        db.session.commit()
        return redirect("/all_rentals")
    except Exception as err:
        return f"There was a problem with deleting this lease >>>> {str(err)}"


if __name__ == "__main__":
    app.run(debug=True)
