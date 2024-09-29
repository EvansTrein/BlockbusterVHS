from flask import render_template, request, redirect, flash
from config import app, db
from models import VhsTape, Client
from re import fullmatch


with app.app_context():
    db.create_all()


@app.route("/")
def home():
    return render_template("home_page.html")


@app.route("/vhs/<int:id>")
def vhs(id):
    vhs = VhsTape.query.get(id)
    return render_template("vhs_page.html", vhs=vhs)


@app.route("/all_vhstapes")
def all_vhstapes():
    all_vhstapes = VhsTape.query.order_by(VhsTape.year.desc()).all()
    return render_template("all_vhstape_page.html", all_vhstapes=all_vhstapes)


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
        elif len(title) > 100 or len(age_rating) > 4:
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
        count = request.form["count"]

        if all((title, year, age_rating, count)):
            vhs.title = title
            vhs.year = year
            vhs.age_rating = age_rating
            vhs.count = count
        else:
            flash("fields cannot be empty")
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

    try:
        db.session.delete(vhs)
        db.session.commit()
        return redirect("/all_vhstapes")
    except Exception as err:
        return f"There was a problem deleting that VHS tape >>>> {str(err)}"


@app.route("/clear_database", methods=["POST"])
def clear_database():
    db.session.query(VhsTape).delete()
    db.session.commit()

    return "The database has been cleared"


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

    try:
        db.session.delete(client)
        db.session.commit()
        return redirect("/all_clients")
    except Exception as err:
        return f"There was a problem deleting that client tape >>>> {str(err)}"


@app.route("/clear_database_clients", methods=["POST"])
def clear_database_clients():
    db.session.query(Client).delete()
    db.session.commit()

    return "The database has been cleared"


if __name__ == "__main__":
    app.run(debug=True)
